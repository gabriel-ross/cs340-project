import { useState, useEffect } from "react";
import NavBar from "../components/Navbar";
import AddGenerationsForm from "../components/AddGenerations";
import {
  Container,
  Row,
  Col,
  Table,
  Button,
  UncontrolledCollapse,
} from "reactstrap";
import axios from "axios";

function Generations() {
  const [generations, setGenerations] = useState(null);

  useEffect(() => {
    axios.get("/generations").then((response) => {
      setGenerations(response.data);
    });
  }, []);

  return (
    <div className="App">
      <NavBar />
      <Container className="p-4">
        <h1>
          Generations{" "}
          <Button color="primary" size="sm" className="mx-3" id="toggler">
            Add Generations Form
          </Button>
        </h1>
        <div>
          <UncontrolledCollapse toggler="#toggler">
            <AddGenerationsForm />
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
                {generations &&
                  generations.map((generation, id) => (
                    <tr key={id}>
                      <th scope="row">{generation.id}</th>
                      <td>{generation.name}</td>
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

export default Generations;
