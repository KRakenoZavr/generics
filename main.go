// package main

// import "fmt"

// type Number interface {
// 	int64 | float64
// }

// func main() {

// 	category := Category{
// 		ID:   1,
// 		Name: "Go Generics",
// 		Slug: "go-generics",
// 	}
// 	// create cache for blog.Category struct
// 	cc := New[Category]()
// 	// add category to cache
// 	cc.Set(category.Slug, category)

// 	// create a new post
// 	post := Post{
// 		ID: 1,
// 		Categories: []Category{
// 			{ID: 1, Name: "Go Generics", Slug: "go-generics"},
// 		},
// 		Title: "Generics in Golang structs",
// 		Slug:  "generics-in-golang-structs",
// 	}
// 	// create cache for blog.Post struct
// 	cp := New[Post]()
// 	// add post to cache
// 	cp.Set(post.Slug, post)

// 	fmt.Println(cp)

// 	// ints := map[string]int64{
// 	// 	"first":  34,
// 	// 	"second": 12,
// 	// }

// 	// floats := map[string]float64{
// 	// 	"first":  35.98,
// 	// 	"second": 26.99,
// 	// }

// 	// fmt.Printf("Generic Sums: %v and %v\n",
// 	// 	SumIntsOrFloats[string, int64](ints),
// 	// 	SumIntsOrFloats[string, float64](floats))
// }

// func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
// 	var s V
// 	for _, v := range m {
// 		s += v
// 	}
// 	return s
// }

// type Category struct {
// 	ID   int32
// 	Name string
// 	Slug string
// }

// type Post struct {
// 	ID          int32
// 	Categories  []Category
// 	Title       string
// 	Description string
// 	Slug        string
// }

// type cacheable interface {
// 	Category | Post
// }

// type cache[T cacheable] struct {
// 	data map[string]T
// }

// func (c *cache[T]) Set(key string, value T) {
// 	c.data[key] = value
// }

// func (c *cache[T]) Get(key string) (v T) {
// 	if v, ok := c.data[key]; ok {
// 		return v
// 	}

// 	return
// }

// func New[T cacheable]() *cache[T] {
// 	c := cache[T]{}
// 	c.data = make(map[string]T)

// 	return &c
// }

package main

import (
	"database/sql"
	"fmt"
	"generic/repo"

	_ "github.com/mattn/go-sqlite3"
)

// type Post struct {
// 	id    string
// 	value int
// }

// type IPostService interface {
// 	entity.IEntity[Post]
// 	pop()
// }

// type PostService struct {
// 	*entity.Entity[Post]
// }

// func (p *PostService) pop() {
// 	p.Entity.Data = p.Entity.Data[:len(p.Entity.Data)]
// }

// func NewPost() *PostService {
// 	p := &PostService{
// 		Entity: entity.NewEntity[Post](),
// 	}

// 	return p
// }

type Post struct {
	ID    int    `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
	// Kek   *int
}

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

func main() {

	// var p IPostService = NewPost()

	// p.pop()

	// asd := Post{ID: 1, Value: "str"}

	// asd1 := new(Post)

	db, err := sql.Open("sqlite3", "ad.db?_foreign_keys=on")
	if err != nil {
		fmt.Println(err)
		return
	}

	postRepo := repo.NewRepo[Post](db, nil, "Post")

	userRepo := repo.NewRepo[User](db, nil, "User")

	users, err := userRepo.Get()
	posts, err := postRepo.Get()

	// pp, err := partial.New(&posts[0])

	// wpek := 1

	// asd := &Post{
	// 	ID:    0,
	// 	Value: "",
	// 	Kek:   &wpek,
	// }

	kek := struct {
		Value string
	}{
		// ID: 0,
		Value: "d",
	}

	// fmt.Printf("asd %+v\n", kek)

	// ads := struct {
	// 	ID string
	// }{
	// 	ID: "12",
	// }

	postRepo.Change("3", kek)

	// pp.New(kek, Post{})

	// pp, err := partial.New[Post](&kek)

	// fmt.Printf("qwe %+v\n", pp)

	// postRepo.Change("1", pp)

	// for _, l := range users {
	// 	l.
	// }

	// data := Post{
	// 	ID:    3,
	// 	Value: "kek",
	// }

	// err = repo.Create(data)
	// fmt.Println(posts)
	// fmt.Println(users)
	// fmt.Println(err)

	// a, err := postRepo.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(a)
}
