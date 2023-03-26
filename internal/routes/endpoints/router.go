package endpoints

import (
	"github.com/calebtraceyco/mind-your-business-api/internal/facade"
	"net/http"
	"reflect"
)

type RouterI interface {
	NewUserHandler(service facade.ServiceI) http.HandlerFunc
	Health() http.HandlerFunc
}

type Router struct {
	//Service facade.ServiceI
}

//
//func buildUserSearchQuery(searchFields map[string]any) (string, any) {
//	var query strings.Builder
//	var args []interface{}
//
//	query.WriteString("SELECT id, name, email, age FROM users WHERE ")
//
//	i := 0
//	for field, value := range searchFields {
//		if i > 0 {
//			query.WriteString(" AND ")
//		}
//
//		query.WriteString(fmt.Sprintf("%s = $%d", field, i+1))
//		args = append(args, value)
//
//		i++
//	}
//
//	return query.String(), args
//}

func mapFields(request interface{}, daoModel interface{}) error {
	requestValue := reflect.ValueOf(request).Elem()
	daoModelValue := reflect.ValueOf(daoModel).Elem()

	for i := 0; i < requestValue.NumField(); i++ {
		requestField := requestValue.Field(i)
		daoModelField := daoModelValue.FieldByName(requestValue.Type().Field(i).Name)

		if daoModelField.IsValid() && daoModelField.CanSet() {
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
	}

	return nil
}
