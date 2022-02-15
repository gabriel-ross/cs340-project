import NavBar from "../components/Navbar";
import AddPokemonMovesForm from "../components/AddPokemonMoves";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  UncontrolledCollapse,
} from "reactstrap";

function PokemonMoves() {
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
            <AddPokemonMovesForm />
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
            <Table className="mt-3" hover responsive>
              <thead>
                <tr>
                  <th>ID #</th>
                  <th>Pokémon #</th>
                  <th>Move #</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <th scope="row">1</th>
                  <td>1</td>
                  <td>1</td>
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
                <tr>
                  <th scope="row">2</th>
                  <td>1</td>
                  <td>2</td>
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
                <tr>
                  <th scope="row">3</th>
                  <td>2</td>
                  <td>3</td>
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
              </tbody>
            </Table>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default PokemonMoves;
