package repo

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type DBModel[T any] struct{}

type IRepo[T any] interface {
	GetById(id string) (T, error)
	Get() ([]T, error)
	Create(data T) error
	Change(id string, data T) error
	Delete(id string) error
}

type Repo[T any] struct {
	db     *sql.DB
	logger *log.Logger
	table  string
}

func NewRepo[T any](db *sql.DB, logger *log.Logger, table string) *Repo[T] {
	return &Repo[T]{
		db, logger, table,
	}
}

func (r *Repo[T]) GetById(id string) (T, error) {
	var tmp T

	row := r.db.QueryRow(fmt.Sprintf("%s WHERE id=", r.selectStmt()), id)
	err := row.Scan(tmp)

	return tmp, err
}

func (r *Repo[T]) Get() ([]T, error) {
	data := make([]T, 0)

	rows, err := r.db.Query(r.selectStmt())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var tmp T

		s := reflect.ValueOf(&tmp).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)

		if err != nil {
			return nil, err
		}

		data = append(data, tmp)
	}

	return data, nil
}

func (rp *Repo[T]) Create(data T) error {

	str := fmt.Sprintf("INSERT INTO %s", rp.table)

	strColumns := "("
	strValues := " VALUES ("

	r := reflect.ValueOf(&data).Elem()
	rt := r.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		dbFieldName := field.Tag.Get("db")

		rv := reflect.ValueOf(&data)
		value := reflect.Indirect(rv).FieldByName(field.Name)

		switch value.Type().Name() {
		case "int":
			strValues += fmt.Sprintf("%d", value.Int())
		case "string":
			strValues += fmt.Sprintf("'%s'", value.String())
		case "float":
			strValues += fmt.Sprintf("%f", value.Float())
		case "bool":
			strValues += fmt.Sprintf("%t", value.Bool())
		}

		strColumns += dbFieldName

		if i != rt.NumField()-1 {
			strColumns += ","
			strValues += ","
		}
	}

	strColumns += ")"
	strValues += ")"

	sqlStatement := str + strColumns + strValues + ";"

	query, err := rp.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	tx, err := rp.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec()
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit()
}

func (rp *Repo[T]) Change(id string, data T) error {

	fmt.Printf("%+v\n", data)

	// fmt.Println(data)

	str := fmt.Sprintf("UPDATE %s SET", rp.table)

	strStmt := ""

	r := reflect.ValueOf(&data).Elem()
	rt := r.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		dbFieldName := field.Tag.Get("db")

		rv := reflect.ValueOf(&data)
		value := reflect.Indirect(rv).FieldByName(field.Name)

		// fmt.Println("val", value.)

		strStmt += fmt.Sprintf("%s=", dbFieldName)

		switch value.Type().Name() {
		case "int":
			strStmt += fmt.Sprintf("%d", value.Int())
		case "string":
			strStmt += fmt.Sprintf("'%s'", value.String())
		case "float":
			strStmt += fmt.Sprintf("%f", value.Float())
		case "bool":
			strStmt += fmt.Sprintf("%t", value.Bool())
		}

		if i != rt.NumField()-1 {
			strStmt += ","
		}
	}

	sqlStatement := fmt.Sprintf("%s %s;", str, strStmt)

	query, err := rp.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	tx, err := rp.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec()
	if err != nil {
		rp.logger.Println("doing rollback")
		rp.logger.Println(err.Error())
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *Repo[T]) Delete(id string) error {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=$1", r.table)

	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(query).Exec(id)
	if err != nil {
		r.logger.Println("doing rollback")
		r.logger.Println(err.Error())
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Private functions
func (r *Repo[T]) selectStmt() string {
	return fmt.Sprintf("SELECT * FROM %s", r.table)
}

// type Entity[T any] struct {
// 	Data  []T
// 	db    *sql.DB
// 	table string
// }

// func (e *Entity[T]) Get(idx int) ([]T, error) {
// 	data := make([]T, 0)

// 	rows, err := e.db.Query(fmt.Sprintf("SELECT * FROM %s", e.table))
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var tmp T
// 		err := rows.Scan(tmp)
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = append(data, tmp)
// 	}

// 	return data, nil
// }

// func (e *Entity[T]) Put(data T) int {
// 	e.Data = append(e.Data, data)
// 	return len(e.Data)
// }

// func NewEntity[T any]() *Entity[T] {
// 	return &Entity[T]{
// 		Data: make([]T, 0),
// 	}
// }
