package model

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: queries need to properly convert between ID & name for
// primary and secondary type
type Pokemon struct {
	NDexId        int    `json:"national-dex-id,string"`
	Name          string `json:"name"`
	PrimaryType   string `json:"primary-type"`
	SecondaryType string `json:"secondary-type"`
	Generation    int    `json:"generation,string"`
}

type PokemonModel struct {
	db *sql.DB
}

func NewPokemonModel(db *sql.DB) *PokemonModel {
	return &PokemonModel{db: db}
}

// FindAll queries all entries in PokemonModel.DB and returns
// a slice of Pokemon
func (p PokemonModel) FindAll() ([]Pokemon, error) {
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
			&respPokemon.NDexId,
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

// FindByID queries PokemonModel.DB for an ID and returns a Pokemon
func (p PokemonModel) FindByID(id int) (*Pokemon, error) {
	sqlStatement := "SELECT * FROM Pokemon WHERE id=?"
	resp := p.db.QueryRow(sqlStatement, id)

	var respPokemon Pokemon
	err := resp.Scan(
		&respPokemon.NDexId,
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
func (p PokemonModel) Find(query map[string]string) ([]Pokemon, error) {
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
			&respPokemon.NDexId,
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
func (p PokemonModel) InsertPokemon(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "INSERT INTO Pokemon Values(?, ?, ?, ?, ?)"
	result, err := p.db.Exec(sqlStatement, pk.NDexId, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation)
	return result, err
}

func (p PokemonModel) UpdatePokemonByID(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "UPDATE Pokemon SET name=?, primary_type=?, secondary_type=?, generation=? WHERE id=?"
	result, err := p.db.Exec(sqlStatement, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.NDexId)
	return result, err
}

func (p PokemonModel) UpdatePokemonByName(pk *Pokemon) (sql.Result, error) {
	sqlStatement := "UPDATE Pokemon SET id=?, primary_type=?, secondary_type=?, generation=? WHERE name=?"
	result, err := p.db.Exec(sqlStatement, pk.NDexId, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.Name)
	return result, err
}

func (p PokemonModel) DeletePokemonByID(id string) error {
	sqlStatement := "DELETE FROM Pokemon WHERE id=?"
	_, err := p.db.Exec(sqlStatement, id)
	return err
}
