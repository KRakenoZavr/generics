package pp

import (
	"fmt"
	"reflect"
)

type A map[string]any

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

func New[T any](instance any, source T) (*Partial[T], error) {

	var instanceSource any

	t := reflect.TypeOf(instanceSource)
	v := reflect.New(t)
	initializeStruct(t, v.Elem())
	c := v.Interface().(T)
	// c.Meta.Properties["color"] = "red"          // map was already made!
	// c.Meta.Users = append(c.Meta.Users, "srid") // so was the slice.
	fmt.Println(v.Interface())
	fmt.Println(c)

	// ri := reflect.ValueOf(&instance).Elem()
	// rti := ri.Type()

	rs := reflect.ValueOf(&source).Elem()
	rts := rs.Type()

	for i := 0; i < rts.NumField(); i++ {
		field := rts.Field(i)

		// dbFieldName := field.Tag.Get("db")

		rvs := reflect.ValueOf(&source)
		value := reflect.Indirect(rvs).FieldByName(field.Name)

		ris := reflect.ValueOf(instance)
		// ival := reflect.Indirect(ris).FieldByName(field.Name)
		ival := ris.FieldByName(field.Name)

		fmt.Println(ival.IsValid(), value)

		// instanceSource

		// fmt.Println("val", value.)

		// strStmt += fmt.Sprintf("%s=", dbFieldName)

		// switch value.Type().Name() {
		// case "int":
		// 	strStmt += fmt.Sprintf("%d", value.Int())
		// case "string":
		// 	strStmt += fmt.Sprintf("'%s'", value.String())
		// case "float":
		// 	strStmt += fmt.Sprintf("%f", value.Float())
		// case "bool":
		// 	strStmt += fmt.Sprintf("%t", value.Bool())
		// }

		// if i != rt.NumField()-1 {
		// 	strStmt += ","
		// }
	}

	return &Partial[T]{
		Instance: instance,
		Source:   source,
	}, nil
}

// func New[T any](subjectPtr *T) (model Partial[T], err error) {
// 	if err != nil {
// 		return model, err
// 	}

// 	base := *subjectPtr
// 	model = Partial[T]{
// 		FieldNames: []string{},
// 		apply: func(thing T) *T {
// 			return &thing
// 		},
// 	}

// model = model.Add(func(subject *T) []string {
// 	fieldNames := []string{}
// 	subjectType := reflect.TypeOf(subject).Elem()
// 	for idx := 0; idx < subjectType.NumField(); idx++ {
// 		field := subjectType.Field(idx)

// 		fieldNames = append(fieldNames, field.Name)
// 		reflect.ValueOf(subject).Elem().FieldByIndex([]int{idx}).Set(
// 			reflect.ValueOf(base).FieldByIndex([]int{idx}),
// 		)
// 	}

// 	return fieldNames
// })

// 	return model, nil
// }

type Partial[T any] struct {
	Instance       any
	Source         T
	SourceInstance any
}
