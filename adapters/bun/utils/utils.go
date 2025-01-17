package utils

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dreamph/dbre"
	"github.com/iancoleman/strcase"
	"github.com/uptrace/bun"
)

const (
	strNewLine = '\''
	strAt      = '@'
)

func DbError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}

type NamedParameterQuery struct {
	parsedParameters []interface{}
	originalQuery    string
	parsedQuery      string
}

func NewNamedParameterQuery(queryText string, params []interface{}) *NamedParameterQuery {
	ret := &NamedParameterQuery{
		originalQuery: queryText,
	}

	ret.parse(params)

	return ret
}

func (n *NamedParameterQuery) parse(params []interface{}) {
	var revisedBuilder bytes.Buffer
	var parameterBuilder bytes.Buffer
	positions := make(map[string][]int)

	queryText := n.originalQuery
	positionIndex := 0

	mapIsSliceParams := make(map[string]bool)
	for _, param := range params {
		arg, ok := param.(sql.NamedArg)
		if ok {
			mapIsSliceParams[arg.Name] = IsSlice(arg.Value)
		}
	}

	for i := 0; i < len(queryText); {
		character, width := utf8.DecodeRuneInString(queryText[i:])
		i += width

		// if it's a colon, do not write to builder, but grab name
		if character == strAt {
			for {
				character, width = utf8.DecodeRuneInString(queryText[i:])
				i += width

				if unicode.IsLetter(character) || unicode.IsDigit(character) {
					parameterBuilder.WriteRune(character)
				} else {
					break
				}
			}

			// add to positions
			parameterName := parameterBuilder.String()
			position := positions[parameterName]
			positions[parameterName] = append(position, positionIndex)
			positionIndex++

			if mapIsSliceParams[parameterName] {
				revisedBuilder.WriteString("(?)")
			} else {
				revisedBuilder.WriteString("?")
			}

			parameterBuilder.Reset()

			if width <= 0 {
				break
			}
		}

		// otherwise write.
		revisedBuilder.WriteRune(character)

		// if it's a quote, continue writing to builder, but do not search for parameters.
		if character == strNewLine {
			for {
				character, width = utf8.DecodeRuneInString(queryText[i:])
				i += width
				revisedBuilder.WriteRune(character)

				if character == strNewLine {
					break
				}
			}
		}
	}

	n.parsedQuery = revisedBuilder.String()
	n.parsedParameters = make([]interface{}, positionIndex)

	for _, param := range params {
		arg, ok := param.(sql.NamedArg)
		if ok {
			for _, index := range positions[arg.Name] {
				if IsSlice(arg.Value) {
					n.parsedParameters[index] = bun.In(arg.Value)
				} else {
					n.parsedParameters[index] = arg.Value
				}
			}
		}
	}
}

func (n *NamedParameterQuery) GetParsedQuery() string {
	return n.parsedQuery
}

func (n *NamedParameterQuery) GetParsedParameters() []interface{} {
	return n.parsedParameters
}

func BuildWhereCause(objPtr interface{}) *[]dbre.WhereCause {
	where := dbre.NewWhereBuilder().Where("1 = 1")

	objElem := reflect.ValueOf(objPtr).Elem()
	objType := reflect.TypeOf(objPtr).Elem()
	for i := 0; i < objElem.NumField(); i++ {
		if !objElem.Field(i).IsZero() {
			varValue := objElem.Field(i).Interface()

			bunTag := objType.Field(i).Tag.Get("bun")
			list := strings.Split(bunTag, ",")
			dbName := GetArrayValueByIndex(list, 0)
			if dbName != "" {
				where.Where(dbName+" = ?", varValue)
			} else {
				dbName = objElem.Type().Field(i).Name
				where.Where(strcase.ToSnake(dbName)+" = ?", varValue)
			}
		}
	}

	return where.WhereCauses()
}

func GetArrayValueByIndex(list []string, index int) string {
	if len(list) < index {
		return ""
	}
	return list[index]
}

func IsSlice(value interface{}) bool {
	typeOf := reflect.TypeOf(value)
	return typeOf.Kind() == reflect.Slice
}

func GenerateSetExpressions(fields []string, pk []string, dbType string) ([]string, []string, error) {
	var expressions []string

	if len(pk) == 0 {
		return nil, nil, fmt.Errorf("primary key fields are required")
	}

	for _, fieldName := range fields {
		if slices.Contains(pk, fieldName) {
			continue
		}

		expression, err := GenerateSetExpression(fieldName, dbType)
		if err != nil {
			return nil, nil, err
		}
		expressions = append(expressions, expression)
	}

	// Return the set expressions along with the provided primary key fields
	return expressions, pk, nil
}

func GenerateSetExpression(fieldName, dbType string) (string, error) {
	switch dbType {
	case "pg": // PostgreSQL
		return fmt.Sprintf("%s = EXCLUDED.%s", fieldName, fieldName), nil
	case "mysql": // MySQL
		return fmt.Sprintf("%s = VALUES(%s)", fieldName, fieldName), nil
	case "mssql": // Microsoft SQL Server
		return fmt.Sprintf("target.%s = source.%s", fieldName, fieldName), nil
	case "sqlite": // SQLite
		return fmt.Sprintf("%s = excluded.%s", fieldName, fieldName), nil
	case "oracle": // Oracle
		return fmt.Sprintf("target.%s = source.%s", fieldName, fieldName), nil
	default:
		return "", fmt.Errorf("unsupported database type: %s", dbType)
	}
}

func GetDbFields(db bun.IDB, obj interface{}) ([]string, []string, error) {
	tables := db.Dialect().Tables()
	table := tables.Get(reflect.TypeOf(obj))
	if table == nil {
		return nil, nil, fmt.Errorf("table not found")
	}

	if len(table.PKs) == 0 {
		return nil, nil, fmt.Errorf("no primary key found in struct")
	}

	if len(table.DataFields) == 0 {
		return nil, nil, fmt.Errorf("no column found in struct")
	}

	var pkFields []string
	var dataFields []string
	for _, field := range table.Fields {
		if field.IsPK {
			pkFields = append(pkFields, field.Name)
		} else {
			dataFields = append(dataFields, field.Name)
		}
	}

	return pkFields, dataFields, nil
}

/*func GetDbFields2(obj interface{}) ([]string, []string, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("expected a struct, got %s", val.Kind())
	}

	var pkFields []string
	var dataFields []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if field.Type == reflect.TypeOf(bun.BaseModel{}) {
			continue
		}

		tag := field.Tag.Get("bun")
		list := strings.Split(tag, ",")
		dbName := GetArrayValueByIndex(list, 0)
		if dbName == "" {
			dbName = strcase.ToSnake(field.Name)
		}
		if strings.Contains(tag, "pk") {
			pkFields = append(pkFields, dbName)
		} else {
			dataFields = append(dataFields, dbName)
		}
	}

	if len(pkFields) == 0 {
		return nil, nil, fmt.Errorf("no primary key found in struct")
	}

	if len(dataFields) == 0 {
		return nil, nil, fmt.Errorf("no column found in struct")
	}

	return pkFields, dataFields, nil
}*/
