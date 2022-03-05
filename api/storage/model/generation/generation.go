package generation

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Generation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}

func (m Model) FindAll() ([]Generation, error) {
	sqlStatement := "SELECT * FROM Generations"
	resp, err := m.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []Generation{}

	for resp.Next() {
		var respGeneration Generation
		err := resp.Scan(
			&respGeneration.Id,
			&respGeneration.Name,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, respGeneration)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m Model) FindByID(id string) (*Generation, error) {
	sqlStatement := "SELECT * FROM Generations WHERE id=?"
	resp := m.db.QueryRow(sqlStatement, id)

	var respGeneration Generation
	err := resp.Scan(
		&respGeneration.Id,
		&respGeneration.Name,
	)
	if err != nil {
		return nil, err
	}
	return &respGeneration, nil
}

func (m Model) FindByName(name string) (*Generation, error) {
	sqlStatement := "SELECT * FROM Generations WHERE name=?"
	resp := m.db.QueryRow(sqlStatement, name)

	var respGeneration Generation
	err := resp.Scan(
		&respGeneration.Id,
		&respGeneration.Name,
	)
	if err != nil {
		return nil, err
	}
	return &respGeneration, nil
}

func (m Model) Insert(gen *Generation) (sql.Result, error) {
	sqlStatement := "INSERT INTO Generations (name) Values(?)"
	result, err := m.db.Exec(sqlStatement, gen.Name)
	return result, err
}

func (m Model) Update(gen *Generation) (sql.Result, error) {
	sqlStatement := "UPDATE Generations SET name=? WHERE id=?"
	result, err := m.db.Exec(sqlStatement, gen.Name, gen.Id)
	return result, err
}

func (m Model) DeleteByID(id string) error {
	sqlStatement := "DELETE FROM Generations WHERE id=?"
	_, err := m.db.Exec(sqlStatement, id)
	return err
}

func (m Model) DeleteByName(name string) error {
	sqlStatement := "DELETE FROM Generations WHERE name=?"
	_, err := m.db.Exec(sqlStatement, name)
	return err
}
