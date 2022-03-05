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
	sqlQuery := "SELECT * FROM Types"
	resp, err := m.db.Query(sqlQuery)
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
	sqlQuery := "SELECT * FROM Types WHERE id=?"
	resp := m.db.QueryRow(sqlQuery, id)

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
	sqlQuery := "SELECT * FROM Types WHERE name=?"
	resp := m.db.QueryRow(sqlQuery, name)

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
	sqlQuery := "INSERT INTO Types (name) Values(?)"
	result, err := m.db.Exec(sqlQuery, elementalType.Name)
	return result, err
}

func (m Model) Update(et *ElementalType) (sql.Result, error) {
	sqlQuery := "UPDATE Types SET name=? WHERE id=?"
	result, err := m.db.Exec(sqlQuery, et.Name, et.Id)
	return result, err
}

func (m Model) DeleteByID(id string) error {
	sqlQuery := "DELETE FROM Types WHERE id=?"
	_, err := m.db.Exec(sqlQuery, id)
	return err
}

func (m Model) DeleteByName(name string) error {
	sqlQuery := "DELETE FROM Types WHERE name=?"
	_, err := m.db.Exec(sqlQuery, name)
	return err
}
