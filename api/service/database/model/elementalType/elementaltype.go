package elementalType

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// type referred to as ElementalType to avoid conflict with Go "type" keyword
type ElementalType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}

func (m Model) FindAll() ([]ElementalType, error) {
	sqlStatement := "SELECT * FROM Types"
	resp, err := m.db.Query(sqlStatement)
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

func (m Model) FindByID(id string) (*ElementalType, error) {
	sqlStatement := "SELECT * FROM Types WHERE id=?"
	resp := m.db.QueryRow(sqlStatement, id)

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

func (m Model) FindByName(name string) (*ElementalType, error) {
	sqlStatement := "SELECT * FROM Types WHERE name=?"
	resp := m.db.QueryRow(sqlStatement, name)

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

func (m Model) Insert(elementalType *ElementalType) (sql.Result, error) {
	sqlStatement := "INSERT INTO Types (name) Values(?)"
	result, err := m.db.Exec(sqlStatement, elementalType.Name)
	return result, err
}

func (m Model) UpdateByID(et *ElementalType) (sql.Result, error) {
	sqlStatement := "UPDATE Types SET name=? WHERE id=?"
	result, err := m.db.Exec(sqlStatement, et.Name, et.Id)
	return result, err
}

func (m Model) DeleteByName(name string) error {
	sqlStatement := "DELETE FROM Types WHERE name=?"
	_, err := m.db.Exec(sqlStatement, name)
	return err
}
