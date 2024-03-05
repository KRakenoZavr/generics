package something

import (
	"fmt"
	"reflect"
)

type Post struct {
	id    int
	value string
}

type User struct {
	id    int
	value string
}

type SomeV interface {
	Post | User
}

// type Model = User

// type IModel interface {
// 	get() SomeV
// }

// type Summable interface {
// 	int | int8
// }

// func sum[T Summable](a, b T) T {
// 	return a + b
// }

type IGen[T SomeV] interface {
	getSelf() interface{}
	getSource() T
}

type Gen[T SomeV] struct {
	self   interface{}
	source T
}

func (g *Gen[T]) getSelf() interface{} {
	return g.self
}

func (g *Gen[T]) getSource() T {
	return g.source
}

func length(v interface{}) int {
	return reflect.ValueOf(v).FieldByName("ids").Len()
}

func update[T SomeV](fieldName string, data T, val string) T {

	var kek IGen[T]

	kek = &Gen[T]{self: data, source: data}

	reflect.ValueOf(kek.getSource()).FieldByName(fieldName)

	return kek.getSource()
}

func main() {
	asd := Post{
		id:    1,
		value: "kek",
	}

	gen := Gen[Post]{
		self:   asd,
		source: asd,
	}

	// kek := sum(1, 2)

	fmt.Println(gen)
	// fmt.Println(gen, kek)
}
