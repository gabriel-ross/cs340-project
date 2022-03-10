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

function AddMoves({ types }) {
  const [data, setData] = useState({
    name: "",
    type: "Bug",
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
      type: data.type,
    };
    await axios.post("/moves", requestData).then((response) => {
      console.log(response);
      window.location.reload();
    });
  };
  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Move</CardTitle>
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
                  placeholder="Move Name"
                  type="text"
                  value={data.name}
                  onChange={handleChange}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="type" sm={2}>
                Move Type
              </Label>
              <Col sm={10}>
                <Input
                  id="type"
                  name="type"
                  type="select"
                  value={data.type}
                  onChange={handleChange}
                >
                  {types &&
                    types.map((type, id) => (
                      <option value={type.name} key={id}>
                        {type.name}
                      </option>
                    ))}
                </Input>
              </Col>
            </FormGroup>
            <Button>Submit</Button>
          </Form>
        </CardBody>
      </CardBody>
    </Card>
  );
}

export default AddMoves;
