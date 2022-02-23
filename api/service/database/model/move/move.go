package move

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: new insertions should insert into the Pokemon_Move table as well.
// This should also be done in the Pokemon table

type Move struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}

func (m Model) FindAll() ([]Move, error) {
	sqlStatement := "SELECT * FROM Moves"
	resp, err := m.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []Move{}

	for resp.Next() {
		var respMove Move
		err := resp.Scan(
			&respMove.Id,
			&respMove.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, respMove)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m Model) FindByID(id string) (*Move, error) {
	sqlStatement := "SELECT * FROM Moves WHERE id=?"
	resp := m.db.QueryRow(sqlStatement, id)

	var respMove Move
	err := resp.Scan(
		&respMove.Id,
		&respMove.Name,
	)
	if err != nil {
		return nil, err
	}
	return &respMove, nil
}

func (m Model) FindByName(name string) (*Move, error) {
	sqlStatement := "SELECT * FROM Moves WHERE name=?"
	resp := m.db.QueryRow(sqlStatement, name)

	var respMove Move
	err := resp.Scan(
		&respMove.Id,
		&respMove.Name,
	)
	if err != nil {
		return nil, err
	}
	return &respMove, nil
}

func (m Model) Insert(gen *Move) (sql.Result, error) {
	sqlStatement := "INSERT INTO Moves (name) Values(?)"
	result, err := m.db.Exec(sqlStatement, gen.Name)
	return result, err
}

func (m Model) UpdateByID(gen *Move) (sql.Result, error) {
	sqlStatement := "UPDATE Moves SET name=? WHERE id=?"
	result, err := m.db.Exec(sqlStatement, gen.Name, gen.Id)
	return result, err
}

func (m Model) DeleteByName(name string) error {
	sqlStatement := "DELETE FROM Moves WHERE name=?"
	_, err := m.db.Exec(sqlStatement, name)
	return err
}
