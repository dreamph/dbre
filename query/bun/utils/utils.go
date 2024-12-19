package utils

import (
	"bytes"
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dreamph/dbre/query"
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
	var position []int
	var character rune
	var parameterName string
	var width int
	var positionIndex int
	positions := make(map[string][]int)

	queryText := n.originalQuery
	positionIndex = 0

	mapIsSliceParams := make(map[string]bool)
	for _, param := range params {
		arg, ok := param.(sql.NamedArg)
		if ok {
			mapIsSliceParams[arg.Name] = IsSlice(arg.Value)
		}
	}

	for i := 0; i < len(queryText); {
		character, width = utf8.DecodeRuneInString(queryText[i:])
		i += width

		// if it's a colon, do not write to builder, but grab name
		if character == strAt {
			for {
				character, width = utf8.DecodeRuneInString(queryText[i:])
				i += width

				if unicode.IsLetter(character) || unicode.IsDigit(character) {
					parameterBuilder.WriteString(string(character))
				} else {
					break
				}
			}

			// add to positions
			parameterName = parameterBuilder.String()
			position = positions[parameterName]
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
		revisedBuilder.WriteString(string(character))

		// if it's a quote, continue writing to builder, but do not search for parameters.
		if character == strNewLine {
			for {
				character, width = utf8.DecodeRuneInString(queryText[i:])
				i += width
				revisedBuilder.WriteString(string(character))

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

func BuildWhereCause(objPtr interface{}) *[]query.WhereCause {
	where := query.NewWhereBuilder().Where("1 = 1")

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
