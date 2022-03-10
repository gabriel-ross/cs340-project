import { useState } from "react";
import axios from "axios";
import {
  Col,
  Button,
  Input,
  Form,
  Card,
  CardBody,
  CardTitle,
  FormGroup,
  Label,
} from "reactstrap";

function AddGenerations() {
  const [data, setData] = useState({
    name: "",
  });

  const handleChange = (e) => {
    const value = e.target.value;
    setData({
      ...data,
      [e.target.name]: value,
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const requestData = {
      name: data.name,
    };
    await axios.post("/generations", requestData).then((response) => {
      window.location.reload()
    });
  };
  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Generation</CardTitle>
        <CardBody>
          <Form onSubmit={handleSubmit}>
            <FormGroup row>
              <Label for="name" sm={2}>
                Name
              </Label>
              <Col sm={10}>
                <Input
                  id="name"
                  name="name"
                  placeholder="Generation Name"
                  type="text"
                  value={data.name}
                  onChange={handleChange}
                />
              </Col>
            </FormGroup>
            <Button>Submit</Button>
          </Form>
        </CardBody>

      </CardBody>
    </Card>
  );
}

export default AddGenerations;
