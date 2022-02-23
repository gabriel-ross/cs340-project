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
        <CardTitle tag="h5">Add Type</CardTitle>
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
          </Form>
        </CardBody>
        <Button>Submit</Button>
      </CardBody>
    </Card>
  );
}

export default AddTypes;
