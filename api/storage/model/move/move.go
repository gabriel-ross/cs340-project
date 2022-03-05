package move

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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

// todo: implement filtering queries

func (m Model) FindAll() ([]Move, error) {
	sqlQuery := "SELECT * FROM Moves"
	resp, err := m.db.Query(sqlQuery)
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
			&respMove.Type,
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
	sqlQuery := "SELECT * FROM Moves WHERE id=?"
	resp := m.db.QueryRow(sqlQuery, id)

	var respMove Move
	err := resp.Scan(
		&respMove.Id,
		&respMove.Name,
		&respMove.Type,
	)
	if err != nil {
		return nil, err
	}
	return &respMove, nil
}

func (m Model) FindByName(name string) (*Move, error) {
	sqlQuery := "SELECT * FROM Moves WHERE name=?"
	resp := m.db.QueryRow(sqlQuery, name)

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

func (m Model) Insert(move *Move) (sql.Result, error) {
	sqlQuery := "INSERT INTO Moves (name, type) Values(?, ?)"
	result, err := m.db.Exec(sqlQuery, move.Name, move.Type)
	return result, err
}

func (m Model) Update(move *Move) (sql.Result, error) {
	sqlQuery := "UPDATE Moves SET name=?, type=? WHERE id=?"
	result, err := m.db.Exec(sqlQuery, move.Name, move.Type, move.Id)
	return result, err
}

func (m Model) DeleteByID(id string) error {
	sqlQuery := "DELETE FROM Moves WHERE id=?"
	_, err := m.db.Exec(sqlQuery, id)
	return err
}

func (m Model) DeleteByName(name string) error {
	sqlQuery := "DELETE FROM Moves WHERE name=?"
	_, err := m.db.Exec(sqlQuery, name)
	return err
}
