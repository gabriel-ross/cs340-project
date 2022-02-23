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

function AddTypes() {
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
                  <option>Normal</option>
                  <option>Fighting</option>
                  <option>Flying</option>
                  <option>Poison</option>
                  <option>Ground</option>
                  <option>Rock</option>
                  <option>Bug</option>
                  <option>Ghost</option>
                  <option>Steel</option>
                  <option>Fire</option>
                  <option>Water</option>
                  <option>Grass</option>
                  <option>Electric</option>
                  <option>Psychic</option>
                  <option>Ice</option>
                  <option>Dragon</option>
                  <option>Dark</option>
                  <option>Fairy</option>
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
