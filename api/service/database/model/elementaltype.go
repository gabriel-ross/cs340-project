package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// type referred to as ElementalType to avoid conflict with Go "type" keyword
type ElementalType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ElementalTypeModel struct {
	db *sql.DB
}

func NewElementalTypeModel(db *sql.DB) *ElementalTypeModel {
	return &ElementalTypeModel{db: db}
}

func (t ElementalTypeModel) FindAll() ([]ElementalType, error) {
	sqlStatement := "SELECT * FROM Types"
	resp, err := t.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []ElementalType{}

	for resp.Next() {
		var respElementalType ElementalType
		err := resp.Scan(
			&respElementalType.Id,
			&respElementalType.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, respElementalType)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (t ElementalTypeModel) FindByID(id string) (*ElementalType, error) {
	sqlStatement := "SELECT * FROM Types WHERE id=?"
	resp := t.db.QueryRow(sqlStatement, id)

	var respElementalType ElementalType
	err := resp.Scan(
		&respElementalType.Id,
		&respElementalType.Name,
	)
	if err != nil {
		return nil, err
	}
	return &respElementalType, nil
}

func (t ElementalTypeModel) InsertType(elementalType *ElementalType) (sql.Result, error) {
	sqlStatement := "INSERT INTO Types (name) Values(?)"
	result, err := t.db.Exec(sqlStatement, elementalType.Name)
	return result, err
}

func (t ElementalTypeModel) UpdateTypeByID(et *ElementalType) (sql.Result, error) {
	sqlStatement := "UPDATE Types SET name=? WHERE id=?"
	result, err := t.db.Exec(sqlStatement, et.Name, et.Id)
	return result, err
}

func (t ElementalTypeModel) DeleteTypeByName(name string) error {
	sqlStatement := "DELETE FROM Pokemon WHERE name=?"
	_, err := t.db.Exec(sqlStatement, name)
	return err
}
