import NavBar from "./components/Navbar";
import AddPokemonForm from "./components/AddPokemon";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  Input,
  Form,
  UncontrolledCollapse
} from "reactstrap";
import { Link } from "react-router-dom";

function App() {
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
            <AddPokemonForm />
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
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
                <tr>
                  <th scope="row">1</th>
                  <td>Bulbasaur</td>
                  <td>Grass</td>
                  <td>Poison</td>
                  <td>I</td>
                  <td>
                    <div>
                      <Link to="/PokemonMoves?query=Bulbasaur">
                      <Button color="primary" outline size="sm">
                        View Moves
                      </Button></Link>{" "}
                      <Button color="primary" outline size="sm">
                        Edit
                      </Button>{" "}
                      <Button color="primary" outline size="sm">
                        Delete
                      </Button>
                    </div>
                  </td>
                </tr>
                <tr>
                  <th scope="row">2</th>
                  <td>Ivysaur</td>
                  <td>Grass</td>
                  <td>Poison</td>
                  <td>I</td>
                  <td>
                    <div>
                      <Link to="/PokemonMoves?query=Ivysaur">
                      <Button color="primary" outline size="sm">
                        View Moves
                      </Button></Link>{" "}
                      <Button color="primary" outline size="sm">
                        Edit
                      </Button>{" "}
                      <Button color="primary" outline size="sm">
                        Delete
                      </Button>
                    </div>
                  </td>
                </tr>
                <tr>
                  <th scope="row">3</th>
                  <td>Venusaur</td>
                  <td>Grass</td>
                  <td>Poison</td>
                  <td>I</td>
                  <td>
                    <div>
                      <Link to="/PokemonMoves?query=Venusaur">
                      <Button color="primary" outline size="sm">
                        View Moves
                      </Button></Link>{" "}
                      <Button color="primary" outline size="sm">
                        Edit
                      </Button>{" "}
                      <Button color="primary" outline size="sm">
                        Delete
                      </Button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </Table>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
