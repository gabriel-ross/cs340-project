import  {useState } from "react";
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

function PokemonMoves() {
  let [searchParams, setSearchParams] = useSearchParams();
  let [query, setQuery] = useState(
    searchParams.get("query")
  );

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
                  <th>ID #</th>
                  <th>Pokémon</th>
                  <th>Move</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <th scope="row">1</th>
                  <td>Bulbasaur</td>
                  <td>Growl</td>
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
                  <td>Ivysaur</td>
                  <td>Poison Powder</td>
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
                  <td>Venusaur</td>
                  <td>Sleep Powder</td>
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
