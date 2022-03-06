package pokemon

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gabriel-ross/cs340-project/server/storage/model/move"
	_ "github.com/go-sql-driver/mysql"
)

type Pokemon struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	PrimaryType   string `json:"primaryType"`
	SecondaryType string `json:"secondaryType,omitempty"`
	Generation    string `json:"generation"`
}

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}

// !: the secType code is repeated several times and will be required every time
// !: a pokemon get is unmarshaled.

// todo: add a method for converting a SQL select result to a Pokemon struct. This should solve both issues above

// FindAll queries all entries in Model.DB and returns
// a slice of Pokemon
func (m Model) FindAll() ([]Pokemon, error) {
	sqlQuery := `SELECT Pokemon.id, Pokemon.name, t1.name AS primary_type, t2.name AS secondary_type, Generations.name FROM Pokemon
	JOIN Types AS t1 ON Pokemon.primary_type=t1.id
	LEFT JOIN Types AS t2 ON Pokemon.secondary_type=t2.id
	JOIN Generations ON Pokemon.generation=Generations.id`
	resp, err := m.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []Pokemon{}

	for resp.Next() {
		var respPokemon Pokemon
		var secType sql.NullString
		err := resp.Scan(
			&respPokemon.Id,
			&respPokemon.Name,
			&respPokemon.PrimaryType,
			&secType,
			&respPokemon.Generation,
		)
		if err != nil {
			return nil, err
		}
		if secType.Valid {
			respPokemon.SecondaryType = secType.String
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
	sqlQuery := `SELECT Pokemon.id, Pokemon.name, t1.name AS primary_type, t2.name AS secondary_type, Generations.name FROM Pokemon
	JOIN Types AS t1 ON Pokemon.primary_type=t1.id
	LEFT JOIN Types AS t2 ON Pokemon.secondary_type=t2.id
	JOIN Generations ON Pokemon.generation=Generations.id
	WHERE Pokemon.id=?`
	resp := m.db.QueryRow(sqlQuery, id)

	var respPokemon Pokemon
	var secType sql.NullString
	err := resp.Scan(
		&respPokemon.Id,
		&respPokemon.Name,
		&respPokemon.PrimaryType,
		&secType,
		&respPokemon.Generation,
	)
	if err != nil {
		return nil, err
	}
	if secType.Valid {
		respPokemon.SecondaryType = secType.String
	}
	return &respPokemon, nil
}

// return any pokemon that match the params in the passed in pokemon
func (m Model) Find(filters map[string]string) ([]Pokemon, error) {

	var sqlQuery strings.Builder
	sqlQuery.WriteString(`SELECT Pokemon.id, Pokemon.name, t1.name AS primary_type, t2.name AS secondary_type, Generations.name FROM Pokemon
	JOIN Types AS t1 ON Pokemon.primary_type=t1.id
	LEFT JOIN Types AS t2 ON Pokemon.secondary_type=t2.id
	JOIN Generations ON Pokemon.generation=Generations.id `)

	count := len(filters)

	// build where
	var where strings.Builder
	where.WriteString("WHERE ")

	for key, val := range filters {
		switch key {
		case "type":
			fmt.Fprintf(&where, "t1.name='%s' OR t2.name='%s' ", val, val)
		case "generation":
			fmt.Fprintf(&where, "Generations.name='%s' ", val)
		}
		if count--; count > 0 {
			sqlQuery.WriteString("AND ")
		}
	}

	sqlQuery.WriteString(where.String())

	resp, err := m.db.Query(sqlQuery.String())
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	result := []Pokemon{}

	for resp.Next() {
		var respPokemon Pokemon
		var secType sql.NullString
		err := resp.Scan(
			&respPokemon.Id,
			&respPokemon.Name,
			&respPokemon.PrimaryType,
			&secType,
			&respPokemon.Generation,
		)
		if err != nil {
			return nil, err
		}
		if secType.Valid {
			respPokemon.SecondaryType = secType.String
		}
		result = append(result, respPokemon)
	}
	if err = resp.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (m Model) Insert(pk *Pokemon) (sql.Result, error) {
	sqlQuery := "INSERT INTO Pokemon Values(?, ?, (SELECT id FROM Types WHERE name=?), (SELECT id FROM Types WHERE name=?), (SELECT id FROM Generations WHERE name=?))"
	result, err := m.db.Exec(sqlQuery, pk.Id, pk.Name, pk.PrimaryType, m.wrapPossibleNull(pk.SecondaryType), pk.Generation)
	return result, err
}

func (m Model) wrapPossibleNull(val string) *sql.NullString {
	return &sql.NullString{
		String: val,
		Valid:  val != "",
	}
}

func (m Model) Update(pk *Pokemon) (sql.Result, error) {
	sqlQuery := "UPDATE Pokemon SET name=?, primary_type=(SELECT id FROM Types WHERE name=?), secondary_type=(SELECT id FROM Types WHERE name=?), generation=(SELECT id FROM Generations WHERE name=?) WHERE id=?"
	result, err := m.db.Exec(sqlQuery, pk.Name, pk.PrimaryType, pk.SecondaryType, pk.Generation, pk.Id)
	return result, err
}

func (m Model) DeleteByID(id string) error {
	sqlQuery := "DELETE FROM Pokemon WHERE id=?"
	_, err := m.db.Exec(sqlQuery, id)
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
