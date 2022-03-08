import  {useState, useEffect } from "react";
import NavBar from "../components/Navbar";
import AddPokemonMovesForm from "../components/AddPokemonMoves";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  Form,
  Input,
  UncontrolledCollapse,
} from "reactstrap";
import {
  useSearchParams
} from "react-router-dom";
import axios from "axios";

function PokemonMoves() {
  let [searchParams, setSearchParams] = useSearchParams();
  let [query, setQuery] = useState(
    searchParams.get("query")
  );

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

  console.log(pokeMoves)

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
            <AddPokemonMovesForm moves={moves} pokemon={pokemon}/>
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
                      <div className="mt-4">
              <Form>
                <Row xs="2">
                  <Col>
                    <Input
                      bsSize="sm"
                      type="search"
                      value={query}
                      placeholder="Search by Pokémon or Move name..."
                    />
                  </Col>
                  <Col>
                    <Button color="primary" outline size="sm">
                      Search
                    </Button>
                  </Col>
                </Row>
              </Form>
            </div>
            <Table className="mt-3" hover responsive>
              <thead>
                <tr>
                  <th>Pokémon</th>
                  <th>Move</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {pokeMoves &&
                  pokeMoves.map((pokeMove, id) => (
                    <tr key={id}>
                  <td>{pokeMove.PkID}</td>
                  <td>{pokeMove.MvID}</td>
                  <td>
                    <div>
                      <Button color="primary" outline size="sm">
                        Edit
                      </Button>{" "}
                      <Button color="primary" outline size="sm">
                        Delete
                      </Button>
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
