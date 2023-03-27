package user

import (
	"encoding/json"
	"github.com/calebtraceyco/mind-your-business-api/external/models"
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
	mapStruct := m.parseStructToSlices(user)
	log.Infoln(mapStruct)
	//
	//for name, value := range mapStruct {
	//	//TODO parse values and
	//	//log.Infoln("field name: %v\n value: %v", name, value)
	//}
	return "", nil
	//return fmt.Sprintf(psql.InsertExec, "users", columnNames, columnValues), nil
}

type MapFields map[string]any

type MapStruct struct {
	Type   reflect.Type
	Value  reflect.Value
	Obj    any
	Fields MapFields
}

func (m Mapper) parseStructToSlices(obj any) MapFields {
	var fieldNames, values []string

	mapStruct := &MapStruct{
		Type:   reflect.TypeOf(obj),
		Value:  reflect.ValueOf(obj),
		Obj:    m.dereferencePointer(obj),
		Fields: make(map[string]any),
	}
	//obj = m.dereferencePointer(obj)
	//t := reflect.TypeOf(obj)
	//v := reflect.ValueOf(obj)

	for i := 0; i < mapStruct.Type.NumField(); i++ {
		field := mapStruct.Value.Field(i)
		// get the field name for struct value with 'db' tag
		fieldName := mapStruct.Type.Field(i).Tag.Get(DatabaseStructTag)

		if field.IsValid() && fieldName != "" && !slices.Contains(excludedTags, fieldName) {
			switch field.Kind() {
			case reflect.String:
				m.handleFieldString(fieldName, field, mapStruct)
			case reflect.Struct:
				fieldNames, values = m.handleFieldStruct(field, fieldName, fieldNames, values)
			case reflect.Slice:
				fieldNames, values = m.handleFieldSlice(field, fieldName, fieldNames, values)
			default:
				continue
			}
		}
	}

	return mapStruct.Fields
}

func parseField() {

}

func (m Mapper) handleFieldString(fieldName string, field reflect.Value, mapStruct *MapStruct) {
	//fieldNames = append(fieldNames, fieldName)
	if value, found := mapStruct.Fields[fieldName]; found {
		value = strings.Join([]string{value.(string), m.wrapInSingleQuotes(field.String())}, ", ")
	} else {
		log.Errorf("handleFieldString: field not found: %s", fieldName)
	}
	//values = append(values, wrapInSingleQuotes(field.String()))
}

// handleFieldStruct turns a nested struct into a string to be stored as jsonb in postgres
// TODO this seems bad since theres no validation for the json value being stored
func (m Mapper) handleFieldStruct(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	if data, err := json.Marshal(field.Interface()); err == nil {
		tags = append(tags, tag)
		values = append(values, m.wrapInSingleQuotes(string(data)))
	} else if err != nil {
		log.Errorf("handleFieldStruct: error: %v", err)
	}
	return tags, values
}

// handleFieldSlice is a mess
func (m Mapper) handleFieldSlice(field reflect.Value, tag string, tags, values []string) ([]string, []string) {
	for i := 0; i < field.Len(); i++ {
		subValue := field.Index(i)
		if subValue.IsValid() && subValue.Kind() == reflect.String {
			tags = append(tags, tag)
			values = append(values, m.wrapInSingleQuotes(subValue.String()))
		}
	}
	return []string{tag}, values
}

func (m Mapper) dereferencePointer(obj any) any {
	if reflect.ValueOf(obj).Kind() == reflect.Pointer {
		obj = reflect.ValueOf(obj).Elem().Interface()
	}
	return obj
}

func (m Mapper) wrapInSingleQuotes(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "\\'") + "'"
}

const DatabaseStructTag = "db"

var excludedTags = []string{"id", "created_at", "updated_at"}
