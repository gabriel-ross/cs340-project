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
  Label,
  Input,
  UncontrolledCollapse,
} from "reactstrap";
import axios from "axios";

function App() {
  const [pokemon, setPokemon] = useState(null);
  const [filteredPokemon, setFilteredPokemon] = useState(null);
  const [types, setTypes] = useState(null);
  const [generations, setGenerations] = useState(null);
  const [data, setData] = useState({
    type: "",
    generation: "",
  });

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
  const handleClear = (e) => {
    setFilteredPokemon(pokemon);
  };
  const handleFilter = (e) => {
    const value = e.target.value;
    setData({
      ...data,
      [e.target.name]: value,
    });
    handleSubmit(value);
  };
  const handleSubmit = async (value) => {
    await axios.get(`/pokemon/?type=${value}`).then((response) => {
      setFilteredPokemon(response.data);
    });
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
        <Row className="justify-content-end">
          <Label
            for="primaryType"
            style={{ textAlign: "right" }}
            sm={2}
          >
            Filter by Type
          </Label>
          <Col sm={2}>
            <Input
              id="type"
              name="type"
              type="select"
              value={data.type}
              onChange={handleFilter}
            >
              {types &&
                types.map((type, id) => (
                  <option value={type.name} key={id}>
                    {type.name}
                  </option>
                ))}
            </Input>
          </Col>
          <Col md="2">
            <Button
              color="primary"
              outline
              size="md"
              onClick={handleClear}
              block
            >
              Clear
            </Button>
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
                        </td>
                      </tr>
                    ))}
                  {filteredPokemon.length === 0 && (
                    <tr>
                      <th scope="row"></th>
                      <th>None found.</th>
<th></th>
                      <th></th>
                      <th></th>
                      <th></th>
                    </tr>
                  )}
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
