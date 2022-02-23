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

function AddPokemonMoves() {
  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Pokémon Moves</CardTitle>
        <CardBody>
          <Form>
            <FormGroup row>
              <Label for="pokemonName" sm={2}>
                Pokémon Name
              </Label>
              <Col sm={10}>
                <Input
                  id="name"
                  name="pokemonName"
                  placeholder="Pokemon Name"
                  type="text"
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="moveName" sm={2}>
                Move Name
              </Label>
              <Col sm={10}>
                <Input
                  id="name"
                  name="moveName"
                  placeholder="Move Name"
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

export default AddPokemonMoves;
