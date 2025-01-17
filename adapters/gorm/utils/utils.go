package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var schemaCache = &sync.Map{}

func DbError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}

func GetDbFields(db *gorm.DB, obj interface{}) ([]string, []string, error) {
	s, err := schema.Parse(obj, schemaCache, db.NamingStrategy)
	if err != nil {
		return nil, nil, err
	}

	var dataFields []string
	for _, field := range s.DBNames {
		if !slices.Contains(s.PrimaryFieldDBNames, field) {
			dataFields = append(dataFields, field)
		}
	}

	if len(s.PrimaryFieldDBNames) == 0 {
		return nil, nil, fmt.Errorf("no primary key found in struct")
	}

	if len(dataFields) == 0 {
		return nil, nil, fmt.Errorf("no column found in struct")
	}

	return s.PrimaryFieldDBNames, dataFields, nil
}

/*
func GetDbFields(obj interface{}) ([]string, []string, error) {
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
		tag := field.Tag.Get("gorm")

		if strings.Contains(tag, "-") {
			continue
		}

		tagSetting := schema.ParseTagSetting(tag, ";")
		if strings.Contains(tag, "primary_key") {
			dbFieldName, ok := tagSetting["COLUMN"]
			if !ok {
				dbFieldName = strcase.ToSnake(field.Name)
			}
			pkFields = append(pkFields, dbFieldName)
		} else {
			dbFieldName, ok := tagSetting["COLUMN"]
			if !ok {
				dbFieldName = strcase.ToSnake(field.Name)
			}
			dataFields = append(dataFields, dbFieldName)
		}
	}

	if len(pkFields) == 0 {
		return nil, nil, fmt.Errorf("no primary key found in struct")
	}

	if len(dataFields) == 0 {
		return nil, nil, fmt.Errorf("no column found in struct")
	}

	return pkFields, dataFields, nil
}
*/
