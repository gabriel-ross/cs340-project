import { useState, useEffect } from "react";
import NavBar from "./components/Navbar";
import AddPokemonForm from "./components/AddPokemon";
import EditButton from "./components/EditPokemon";
import DeleteButton from "./components/DeleteButton";
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
import { useSearchParams } from "react-router-dom";
import axios from "axios";

function App() {
  let [searchParams, setSearchParams] = useSearchParams();
  let [query, setQuery] = useState(searchParams.get("query") || "");
  const [pokemon, setPokemon] = useState(null);
  const [filteredPokemon, setFilteredPokemon] = useState(null);
  const [types, setTypes] = useState(null);
  const [generations, setGenerations] = useState(null);

  useEffect(() => {
    axios.get("/pokemon").then((response) => {
      setFilteredPokemon(response.data);
      setPokemon(response.data);
    });
    axios.get("/types").then((response) => {
      setTypes(response.data);
    });
    axios.get("/generations").then((response) => {
      setGenerations(response.data);
    });
  }, []);

  const handleChange = (e) => {
    const value = e.target.value;
    setQuery(value);
  };
  const handleSubmit = async (event) => {
    event.preventDefault();
    let result = pokemon.filter(function (poke) {
      console.log(poke.name.toLowerCase() === query.toLowerCase());
      return (
        poke.name.toLowerCase() === query.toLowerCase() ||
        poke.primaryType.toLowerCase() === query.toLowerCase() ||
        (poke.secondaryType &&
          poke.secondaryType.toLowerCase() === query.toLowerCase())
      );
    });
    setFilteredPokemon(result);
  };

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
            <AddPokemonForm types={types} generations={generations} />
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
            <div className="mt-4">
              <Form onSubmit={handleSubmit}>
                <Row xs="2">
                  <Col>
                    <Input
                      bsSize="sm"
                      type="search"
                      value={query}
                      onChange={handleChange}
                      placeholder="Search by Pokémon or by Type..."
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
          </Col>
        </Row>
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
                  {filteredPokemon.length > 0 &&
                    filteredPokemon.map((poke, id) => (
                      <tr key={id}>
                        <th scope="row">{poke.id}</th>
                        <td>{poke.name}</td>
                        <td>{poke.primaryType}</td>
                        <td>{poke.secondaryType}</td>
                        <td>{poke.generation}</td>
                        <td>
                          <div>
                            <EditButton
                              id={poke.id}
                              name={poke.name}
                              primaryType={poke.primaryType}
                              secondaryType={poke.secondaryType}
                              generation={poke.generation}
                              types={types}
                              generations={generations}
                            />{" "}
                            <DeleteButton route={`/pokemon/${poke.id}`} />
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
