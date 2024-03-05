package something

import "reflect"

// type Some[T any] interface {
// 	New() T
// }

type DBPart interface {
	int | string | bool
}

type Part[T any] struct {
	self   interface{}
	source T
}

type IPart[T any] interface {
	getSelf() interface{}
	getSource() T
}

func (p *Part[T]) GetSelf() interface{} {
	return p.self
}

func (p *Part[T]) GetSource() T {
	return p.source
}

func (p *Part[T]) GetKeyValue() map[string]interface{} {

	r := reflect.ValueOf(p.GetSelf()).Elem()
	rt := r.Type()

	var keyMap = make(map[string]interface{})

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldName := field.Name

		sourceValue := reflect.ValueOf(p.GetSource()).FieldByName(fieldName)

		keyMap[fieldName] = sourceValue
	}

	return keyMap
}

func setKeyValue[T any](self interface{}, source T) T {
	rrs := reflect.ValueOf(self)

	r := reflect.ValueOf(self).Elem()
	rt := r.Type()

	rvS := reflect.ValueOf(source)
	// rtS := r.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldName := field.Name

		reflectValue := rvS.FieldByName(fieldName)

		if reflectValue.IsZero() {
			continue
		}

		rvS.SetMapIndex(reflect.ValueOf(fieldName), reflect.ValueOf(rrs.MapIndex(reflect.ValueOf(fieldName))))

		// sourceValue := reflect.ValueOf(source).FieldByName(fieldName)

		// keyMap[fieldName] = sourceValue
	}

	// return keyMap

}

func NewPart[T any](self interface{}, source T) Part[T] {
	asd := new(T)

	return Part[T]{
		self:   self,
		source: source,
	}
}
