import { useState, useEffect } from "react";
import NavBar from "../components/Navbar";
import AddPokemonMovesForm from "../components/AddPokemonMoves";
import DeleteButton from "../components/DeleteButton";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  UncontrolledCollapse,
} from "reactstrap";
import axios from "axios";

function PokemonMoves() {
  const [moves, setMoves] = useState(null);
  const [pokemon, setPokemon] = useState(null);
  const [pokeMoves, setpokeMoves] = useState(null);

  useEffect(() => {
    axios.get("/pokemon/moves").then((response) => {
      setpokeMoves(response.data);
    });
    axios.get("/moves").then((response) => {
      setMoves(response.data);
    });
    axios.get("/pokemon").then((response) => {
      setPokemon(response.data);
    });
  }, []);

  console.log(pokeMoves);

  return (
    <div className="App">
      <NavBar />
      <Container className="p-4">
        <h1>
          Pokémon Moves{" "}
          <Button color="primary" size="sm" className="mx-3" id="toggler">
            Add Pokémon Moves Form
          </Button>
        </h1>
        <div>
          <UncontrolledCollapse toggler="#toggler">
            <AddPokemonMovesForm moves={moves} pokemon={pokemon} />
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
            <Table className="mt-3" hover responsive>
              <thead>
                <tr>
                  <th>Pokémon ID</th>
                  <th>Move ID</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {pokeMoves &&
                  pokeMoves.map((pokeMove, id) => (
                    <tr key={id}>
                      <td>{pokeMove.pokemon_id}</td>
                      <td>{pokeMove.move_id}</td>
                      <td>
                        <div>
                          <DeleteButton
                            route={`/pokemon/${pokeMove.pokemon_id}/moves/${pokeMove.move_id}`}
                          />
                        </div>
                      </td>
                    </tr>
                  ))}
              </tbody>
            </Table>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default PokemonMoves;
