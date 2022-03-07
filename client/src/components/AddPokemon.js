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

function AddPokemon({ types, generations }) {
  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Pokémon</CardTitle>
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
                  placeholder="Pokémon Name"
                  type="text"
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="primaryType" sm={2}>
                Primary Type
              </Label>
              <Col sm={10}>
                <Input id="primaryType" name="primaryType" type="select">
                  {types && types.map((type, id) => (
                    <option value={type.name} key={id}>{type.name}</option>
                  ))}
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="secondaryType" sm={2}>
                Secondary Type
              </Label>
              <Col sm={10}>
                <Input id="secondaryType" name="secondaryType" type="select">
                  <option value="None">None</option>
                  {types && types.map((type, id) => (
                    <option value={type.name} key={id}>{type.name}</option>
                  ))}
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="generation" sm={2}>
                Generation
              </Label>
              <Col sm={10}>
                <Input id="generation" name="generation" type="select">
                  {generations && generations.map((generation, id) => (
                    <option value={generation.name} key={id}>{generation.name}</option>
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

export default AddPokemon;
