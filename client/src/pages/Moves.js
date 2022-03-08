import { useState, useEffect } from "react";
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
import axios from "axios";

function Moves() {
  const [moves, setMoves] = useState(null);
  const [types, setTypes] = useState(null);

  useEffect(() => {
    axios.get("/moves").then((response) => {
      setMoves(response.data);
    });
    axios.get("/types").then((response) => {
      setTypes(response.data);
    });
  }, []);
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
            <AddMovesForm types={types} />
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
                {moves &&
                  moves.map((move, id) => (
                    <tr key={id}>
                      <th scope="row">{move.id}</th>
                      <td>{move.name}</td>
                      <td>{move.type}</td>
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

export default Moves;
