package user

import (
	"encoding/json"
	"fmt"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
	"github.com/calebtraceyco/mind-your-business-api/internal/dao/psql"
	log "github.com/sirupsen/logrus"
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
	var fieldNames, values []string

	fieldMap := make(map[string]any)

	obj = dereferencePointer(obj)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		// get the field name for struct value with 'db' tag
		fieldName := t.Field(i).Tag.Get(DatabaseStructTag)

		if field.IsValid() && fieldName != "" && !slices.Contains(excludedTags, fieldName) {
			switch field.Kind() {
			case reflect.String:
				fieldMap = handleFieldString(field, fieldName, fieldMap)
			case reflect.Struct:
				fieldNames, values = handleFieldStruct(field, fieldName, fieldNames, values)
			case reflect.Slice:
				fieldNames, values = handleFieldSlice(field, fieldName, fieldNames, values)
			default:
				continue
			}
		}
	}

	return strings.Join(fieldNames, ", "), strings.Join(values, ", ")
}

func handleFieldString(field reflect.Value, fieldName string, fieldMap map[string]any) map[string]any {
	//fieldNames = append(fieldNames, fieldName)
	if value, found := fieldMap[fieldName]; found {
		value = strings.Join([]string{value.(string), wrapInSingleQuotes(field.String())}, ", ")
	}
	//values = append(values, wrapInSingleQuotes(field.String()))
	return fieldMap
}

// handleFieldStruct turns a nested struct into a string to be stored as jsonb in postgres
// TODO this seems bad since theres no validation for the json value being stored
func handleFieldStruct(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	if data, err := json.Marshal(field.Interface()); err == nil {
		tags = append(tags, tag)
		values = append(values, wrapInSingleQuotes(string(data)))
	} else if err != nil {
		log.Errorf("handleFieldStruct: error: %v", err)
	}
	return tags, values
}

// handleFieldSlice is a mess
func handleFieldSlice(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	for i := 0; i < field.Len(); i++ {
		subValue := field.Index(i)
		if subValue.IsValid() && subValue.Kind() == reflect.String {
			tags = append(tags, tag)
			values = append(values, wrapInSingleQuotes(subValue.String()))
		}
	}
	return []string{tag}, values
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

var excludedTags = []string{"id", "created_at", "updated_at"}
