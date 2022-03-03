package pokemon

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gabriel-ross/cs340-project/server/service/database/model/move"
	_ "github.com/go-sql-driver/mysql"
)

type Pokemon struct {
	Id            string `json:"national-dex-id,string"`
	Name          string `json:"name"`
	PrimaryType   string `json:"primary-type"`
	SecondaryType string `json:"secondary-type"`
	Generation    string `json:"generation,string"`
}

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}

// FindAll queries all entries in Model.DB and returns
// a slice of Pokemon
func (m Model) FindAll() ([]Pokemon, error) {
	sqlStatement := "SELECT * FROM Pokemon"
	resp, err := m.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []Pokemon{}

	for resp.Next() {
		var respPokemon Pokemon
		err := resp.Scan(
			&respPokemon.Id,
			&respPokemon.Name,
			&respPokemon.PrimaryType,
			&respPokemon.SecondaryType,
			&respPokemon.Generation,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, respPokemon)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// FindByID queries Model.DB for an ID and returns a Pokemon
func (m Model) FindByID(id string) (*Pokemon, error) {
	sqlStatement := "SELECT * FROM Pokemon WHERE id=?"
	resp := m.db.QueryRow(sqlStatement, id)

	var respPokemon Pokemon
	err := resp.Scan(
		&respPokemon.Id,
		&respPokemon.Name,
		&respPokemon.PrimaryType,
		&respPokemon.SecondaryType,
		&respPokemon.Generation,
	)
	if err != nil {
		return nil, err
	}
	return &respPokemon, nil
}

// return any pokemon that match the params in the passed in pokemon
func (m Model) Find(query map[string]string) ([]Pokemon, error) {
	var sqlStatement strings.Builder
	sqlStatement.WriteString("SELECT * FROM Pokemon WHERE ")
	count := len(query)
	for param, val := range query {
		if val != "" {
			// search in both primary_type and secondary_type columns for the type
			// query param
			if param == "type" {
				fmt.Fprintf(&sqlStatement, "primary_type=%s OR secondary_type=%s ", val, val)
			} else {
				fmt.Fprintf(&sqlStatement, "%s=%s ", param, val)
			}
		}
		// concatenate AND between query params. should skip on last param
		if count--; count > 0 {
			sqlStatement.WriteString("AND ")
		}
	}

	resp, err := m.db.Query(sqlStatement.String())
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []Pokemon{}

	for resp.Next() {
		var respPokemon Pokemon
		err := resp.Scan(
			&respPokemon.Id,
			&respPokemon.Name,
			&respPokemon.PrimaryType,
			&respPokemon.SecondaryType,
			&respPokemon.Generation,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, respPokemon)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m Model) Insert(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "INSERT INTO Pokemon Values(?, ?, ?, ?, ?)"
	result, err := m.db.Exec(sqlStatement, pk.Id, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation)
	return result, err
}

func (m Model) UpdateByID(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "UPDATE Pokemon SET name=?, primary_type=?, secondary_type=?, generation=? WHERE id=?"
	result, err := m.db.Exec(sqlStatement, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.Id)
	return result, err
}

func (m Model) UpdateByName(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "UPDATE Pokemon SET id=?, primary_type=?, secondary_type=?, generation=? WHERE name=?"
	result, err := m.db.Exec(sqlStatement, pk.Id, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.Name)
	return result, err
}

func (m Model) DeleteByID(id string) error {
	sqlStatement := "DELETE FROM Pokemon WHERE id=?"
	_, err := m.db.Exec(sqlStatement, id)
	return err
}

func (m Model) FindAllMovesByPokemonID(id string) ([]move.Move, error) {
	sqlQuery := `SELECT Moves.id, Moves.name, Moves.type FROM Pokemon_Moves
	JOIN Moves ON Pokemon_Moves.move_id=Moves.id
	WHERE Pokemon_Moves.pokemon_id=?`
	resp, err := m.db.Query(sqlQuery, id)
	if err != nil {
		return nil, err
	}
	result := []move.Move{}

	for resp.Next() {
		var move move.Move
		err := resp.Scan(
			&move.Id,
			&move.Name,
			&move.Type,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, move)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m Model) InsertPokemonMove(pkid, mvid string) (sql.Result, error) {
	sqlQuery := `INSERT INTO Pokemon_Moves VALUES(?, ?)`
	result, err := m.db.Exec(sqlQuery, pkid, mvid)
	return result, err
}

func (m Model) DeletePokemonMove(pkid, mvid string) (sql.Result, error) {
	sqlQuery := `DELETE FROM Pokemon_Moves WHERE pokemon_id=? AND move_id=?`
	result, err := m.db.Exec(sqlQuery, pkid, mvid)
	return result, err
}
