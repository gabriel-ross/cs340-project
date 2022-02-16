import NavBar from "../components/Navbar";
import AddMovesForm from "../components/AddMoves";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  UncontrolledCollapse,
} from "reactstrap";

function Moves() {
  return (
    <div className="App">
      <NavBar />
      <Container className="p-4">
        <h1>
          Moves{" "}
          <Button color="primary" size="sm" className="mx-3" id="toggler">
            Add Moves Form
          </Button>
        </h1>
        <div>
          <UncontrolledCollapse toggler="#toggler">
            <AddMovesForm />
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
            <Table className="mt-3" hover responsive>
              <thead>
                <tr>
                  <th>ID #</th>
                  <th>Name</th>
                  <th>Type</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <th scope="row">1</th>
                  <td>Growl</td>
                  <td>Normal</td>
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
                  <td>Poison Powder</td>
                  <td>Poison</td>
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
                  <td>Sleep Powder</td>
                  <td>Grass</td>
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

export default Moves;
