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

function AddPokemon() {
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
            <FormGroup row>
              <Label for="secondaryType" sm={2}>
                Secondary Type
              </Label>
              <Col sm={10}>
                <Input id="secondaryType" name="secondaryType" type="select">
                  <option>None</option>
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
            <FormGroup row>
              <Label for="generation" sm={2}>
                Generation
              </Label>
              <Col sm={10}>
                <Input id="generation" name="generation" type="select">
                  <option>I</option>
                  <option>II</option>
                  <option>III</option>
                  <option>IV</option>
                  <option>V</option>
                  <option>VI</option>
                  <option>VII</option>
                  <option>VIII</option>
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
