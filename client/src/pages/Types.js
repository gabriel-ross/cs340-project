import { useState, useEffect } from "react";
import NavBar from "../components/Navbar";
import AddTypesForm from "../components/AddTypes";
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

function Types() {
  const [types, setTypes] = useState(null);

  useEffect(() => {
    axios.get("/types").then((response) => {
      setTypes(response.data);
    });
  }, []);
  return (
    <div className="App">
      <NavBar />
      <Container className="p-4">
        <h1>
          Types{" "}
          <Button color="primary" size="sm" className="mx-3" id="toggler">
            Add Types Form
          </Button>
        </h1>
        <div>
          <UncontrolledCollapse toggler="#toggler">
            <AddTypesForm />
          </UncontrolledCollapse>
        </div>
        <Row>
          <Col>
            <Table className="mt-3" hover responsive>
              <thead>
                <tr>
                  <th>ID #</th>
                  <th>Name</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {types &&
                  types.map((type, id) => (
                    <tr key={id}>
                      <th scope="row">{type.id}</th>
                      <td>{type.name}</td>
                      <td>
                        <div>
                          <Button color="primary" outline size="sm">
                            Edit
                          </Button>{" "}
                          <DeleteButton route={`/types/${type.id}`} />
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

export default Types;
