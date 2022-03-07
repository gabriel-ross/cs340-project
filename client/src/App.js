import { useState, useEffect } from "react";
import NavBar from "./components/Navbar";
import AddPokemonForm from "./components/AddPokemon";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  UncontrolledCollapse,
} from "reactstrap";
import { Link } from "react-router-dom";
import axios from "axios";

function App() {
  const [pokemon, setPokemon] = useState(null);
  const [types, setTypes] = useState(null);
  const [generations, setGenerations] = useState(null);

  useEffect(() => {
    axios.get("/pokemon").then((response) => {
      setPokemon(response.data);
    });
    axios.get("/types").then((response) => {
      setTypes(response.data);
    });
    axios.get("/generations").then((response) => {
      setGenerations(response.data);
    });
  }, []);

  return (
    <div className="App">
      <NavBar />
      <Container className="p-4">
        <h1>
          Pokémon{" "}
          <Button color="primary" size="sm" className="mx-3" id="toggler">
            Add Pokémon Form
          </Button>
        </h1>
        <div>
          <UncontrolledCollapse toggler="#toggler">
            <AddPokemonForm types={types} generations={generations}/>
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
            {pokemon && (
              <Table className="mt-3" hover responsive>
                <thead>
                  <tr>
                    <th>ID #</th>
                    <th>Name</th>
                    <th>Primary Type</th>
                    <th>Secondary Type</th>
                    <th>Generation</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {pokemon.map((poke, id) => (
                    <tr key={id}>
                      <th scope="row">{poke.id}</th>
                      <td>{poke.name}</td>
                      <td>{poke.primaryType}</td>
                      <td>{poke.secondaryType}</td>
                      <td>{poke.generation}</td>
                      <td>
                        <div>
                          <Link to={`/PokemonMoves?query=${poke.name}`}>
                            <Button color="primary" outline size="sm">
                              View Moves
                            </Button>
                          </Link>{" "}
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
            )}
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
