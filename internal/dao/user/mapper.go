package user

import (
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
	cols, values := parseStructToSlices(user.Detail)

	return fmt.Sprintf(psql.InsertExec, "users", strings.Join(cols, ", "), strings.Join(values, ", "))
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
			tag := t.Field(i).Tag.Get(DatabaseStructTag)

			switch field.Kind() {
			case reflect.String:
				if str := field.String(); str != "" {
					tags = append(tags, tag)
					values = append(values, wrapInSingleQuotes(field.String()))
				}
			case reflect.Struct:
				//t, v := parseStructToSlices(field)
				log.Infof("Struct field missed from exec mapping: %s", tag)
			default:
			}
		}
	}

	return tags, values
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
