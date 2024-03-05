package entity

import (
	"database/sql"
	"fmt"
)

type IEntity[T any] interface {
	Get(int) ([]T, error)
	Put(T) int
}

type Entity[T any] struct {
	Data  []T
	db    *sql.DB
	table string
}

func (e *Entity[T]) Get(idx int) ([]T, error) {
	data := make([]T, 0)

	rows, err := e.db.Query(fmt.Sprintf("SELECT * FROM %s", e.table))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp T
		err := rows.Scan(tmp)
		if err != nil {
			return nil, err
		}
		data = append(data, tmp)
	}

	return data, nil
}

func (e *Entity[T]) Put(data T) int {
	e.Data = append(e.Data, data)
	return len(e.Data)
}

func NewEntity[T any]() *Entity[T] {
	return &Entity[T]{
		Data: make([]T, 0),
	}
}
