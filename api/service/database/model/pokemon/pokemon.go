package pokemon

import (
	"database/sql"
	"fmt"
	"strings"

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
func (p Model) FindAll() ([]Pokemon, error) {
	sqlStatement := "SELECT * FROM Pokemon"
	resp, err := p.db.Query(sqlStatement)
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
func (p Model) FindByID(id string) (*Pokemon, error) {
	sqlStatement := "SELECT * FROM Pokemon WHERE id=?"
	resp := p.db.QueryRow(sqlStatement, id)

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
func (p Model) Find(query map[string]string) ([]Pokemon, error) {
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

	resp, err := p.db.Query(sqlStatement.String())
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

// TODO: primary and secondary type need to query the types table for
// the id's of the types
func (p Model) Insert(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "INSERT INTO Pokemon Values(?, ?, ?, ?, ?)"
	result, err := p.db.Exec(sqlStatement, pk.Id, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation)
	return result, err
}

func (p Model) UpdateByID(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "UPDATE Pokemon SET name=?, primary_type=?, secondary_type=?, generation=? WHERE id=?"
	result, err := p.db.Exec(sqlStatement, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.Id)
	return result, err
}

func (p Model) UpdateByName(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "UPDATE Pokemon SET id=?, primary_type=?, secondary_type=?, generation=? WHERE name=?"
	result, err := p.db.Exec(sqlStatement, pk.Id, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.Name)
	return result, err
}

func (p Model) DeleteByID(id string) error {
	sqlStatement := "DELETE FROM Pokemon WHERE id=?"
	_, err := p.db.Exec(sqlStatement, id)
	return err
}
