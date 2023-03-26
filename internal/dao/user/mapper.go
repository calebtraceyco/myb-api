package user

import (
	"encoding/json"
	"fmt"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
	"golang.org/x/exp/slices"
	"reflect"
	"strings"
)

type MapperI interface {
	MapUserExec(user *models.User) (string, error)
}

type Mapper struct{}

const MapUserExecError = "MapUserExec: user is nil"

func (m Mapper) MapUserExec(user *models.User) (string, error) {
	if reflect.ValueOf(user).IsNil() {
		panic(MapUserExecError)
	}
	columnNames, columnValues := parseStructToSlices(user)
	return fmt.Sprintf(psql.InsertExec, "users", columnNames, columnValues), nil
}

func parseStructToSlices(obj any) (string, string) {
	var tags, values []string

	obj = dereferencePointer(obj)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	numFields := t.NumField()

	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get(DatabaseStructTag)

		kind := field.Kind()

		if field.IsValid() && tag != "" && !slices.Contains(excludedTags, tag) {
			switch kind {
			case reflect.String:
				tags, values = handleFieldString(field, tag, tags, values)
			case reflect.Struct:
				tags, values = handleFieldStruct(field, tag, tags, values)
			case reflect.Slice:
				tags, values = handleFieldSlice(field, tag, tags, values)
			default:
			}
		} else {
			continue
		}
	}

	return strings.Join(tags, ", "), strings.Join(values, ", ")
}

func handleFieldString(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	tags = append(tags, tag)
	values = append(values, wrapInSingleQuotes(field.String()))
	return tags, values
}

func handleFieldStruct(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	if data, err := json.Marshal(field); err == nil {
		tags = append(tags, tag)
		values = append(values, wrapInSingleQuotes(string(data)))
	} else if err != nil {
		panic(err)
	}
	return tags, values
}

func handleFieldSlice(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	fieldType := reflect.TypeOf(field)

	for i := 0; i < fieldType.NumField(); i++ {
		subField := reflect.ValueOf(field)
		subValue := subField.Field(i)

		if subField.IsValid() {
			switch subValue.Kind() {
			case reflect.String:
				values = append(values, wrapInSingleQuotes(subValue.String()))
			}
		}
	}

	return []string{tag}, values
}

//
//func mapFields(request any, daoModel any) error {
//	requestValue := reflect.ValueOf(request).Elem()
//	daoModelValue := reflect.ValueOf(daoModel).Elem()
//
//	for i := 0; i < requestValue.NumField(); i++ {
//		requestField := requestValue.Field(i)
//		daoModelField := daoModelValue.FieldByName(requestValue.Type().Field(i).Name)
//
//		if daoModelField.IsValid() && daoModelField.CanSet() {
//			switch requestField.Kind() {
//			case reflect.String:
//				if str := requestField.String(); str != "" {
//					//tags = append(tags, tag)
//					//values = append(values, wrapInSingleQuotes(field.String()))
//				}
//			case reflect.Struct:
//
//				nestedModel := reflect.New(daoModelField.Type())
//				if err := mapFields(requestField.Addr().Interface(), nestedModel.Interface()); err != nil {
//					return err
//				}
//				daoModelField.Set(nestedModel.Elem())
//			default:
//				daoModelField.Set(requestField)
//			}
//		}
//
//		if requestField.Kind() == reflect.Struct {
//			nestedModel := reflect.New(daoModelField.Type())
//			if err := mapFields(requestField.Addr().Interface(), nestedModel.Interface()); err != nil {
//				return err
//			}
//			daoModelField.Set(nestedModel.Elem())
//		} else {
//			daoModelField.Set(requestField)
//		}
//	}
//
//	return nil
//}

func dereferencePointer(obj any) any {
	if reflect.ValueOf(obj).Kind() == reflect.Pointer {
		obj = reflect.ValueOf(obj).Elem().Interface()
	}
	return obj
}

func wrapInSingleQuotes(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "\\'") + "'"
}

const DatabaseStructTag = "db"

var excludedTags = []string{"id", "created_at", "updated_at"}
