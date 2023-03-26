package user

import (
	"encoding/json"
	"fmt"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type MapperI interface {
	MapUserExec(user *models.User) string
}

type Mapper struct{}

func (m Mapper) MapUserExec(user *models.User) string {
	detailCols, detailValues := parseStructToSlices(user)

	return fmt.Sprintf(psql.InsertExec, "users", strings.Join(detailCols, ", "), strings.Join(detailValues, ", "))
}

func parseStructToSlices(obj any) ([]string, []string) {
	var tags, values []string

	obj = dereferencePointer(obj)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	numFields := t.NumField()

	for i := 0; i < numFields; i++ {
		// 'db' struct tag field = db column name
		field := v.Field(i)

		if field.IsValid() {
			if tag := t.Field(i).Tag.Get(DatabaseStructTag); tag != "" && !strings.EqualFold(tag, "id") {
				switch field.Kind() {
				case reflect.String:
					if str := field.String(); str != "" {
						tags = append(tags, tag)
						values = append(values, wrapInSingleQuotes(field.String()))
					}
				case reflect.Struct:
					if data, err := json.Marshal(field.Interface()); err == nil {
						tags = append(tags, tag)
						values = append(values, wrapInSingleQuotes(string(data)))
					}
					//t, v := parseStructToSlices(field)
					log.Infof("Struct field missed from exec mapping: %s", tag)
				default:
				}
			} else {
				continue
			}
		}
	}

	return tags, values
}

func mapFields(request any, daoModel any) error {
	requestValue := reflect.ValueOf(request).Elem()
	daoModelValue := reflect.ValueOf(daoModel).Elem()

	for i := 0; i < requestValue.NumField(); i++ {
		requestField := requestValue.Field(i)
		daoModelField := daoModelValue.FieldByName(requestValue.Type().Field(i).Name)

		if daoModelField.IsValid() && daoModelField.CanSet() {
			switch requestField.Kind() {
			case reflect.String:
				if str := requestField.String(); str != "" {
					//tags = append(tags, tag)
					//values = append(values, wrapInSingleQuotes(field.String()))
				}
			case reflect.Struct:

				nestedModel := reflect.New(daoModelField.Type())
				if err := mapFields(requestField.Addr().Interface(), nestedModel.Interface()); err != nil {
					return err
				}
				daoModelField.Set(nestedModel.Elem())
			default:
				daoModelField.Set(requestField)
			}
		}

		if requestField.Kind() == reflect.Struct {
			nestedModel := reflect.New(daoModelField.Type())
			if err := mapFields(requestField.Addr().Interface(), nestedModel.Interface()); err != nil {
				return err
			}
			daoModelField.Set(nestedModel.Elem())
		} else {
			daoModelField.Set(requestField)
		}
	}

	return nil
}

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
