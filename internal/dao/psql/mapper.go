package psql

import (
	"fmt"
	"github.com/calebtraceyco/mind-your-business-api/external"
	"reflect"
	"strings"
)

type MapperI interface {
	NewUserExec(request *external.ApiRequest) string
}

type Mapper struct{}

func (m Mapper) NewUserExec(request *external.ApiRequest) string {
	columns, values := parseStructToSlices(request.Payload.Request.User)
	return fmt.Sprintf(InsertExec, "users", strings.Join(columns, ", "), strings.Join(values, ", "))
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
			default:
			}
		}
	}

	return tags, values
}
