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

function AddTypes({types}) {
  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Move</CardTitle>
        <CardBody>
          <Form>
            <FormGroup row>
              <Label for="name" sm={2}>
                Name
              </Label>
              <Col sm={10}>
                <Input
                  id="name"
                  name="name"
                  placeholder="Type Name"
                  type="text"
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="moveType" sm={2}>
                Move Type
              </Label>
              <Col sm={10}>
                  <Input id="moveType" name="moveType" type="select">
                  {types && types.map((type, id) => (
                    <option value={type.name} key={id}>{type.name}</option>
                  ))}
                </Input>
              </Col>
            </FormGroup>
          </Form>
        </CardBody>
        <Button>Submit</Button>
      </CardBody>
    </Card>
  );
}

export default AddTypes;
