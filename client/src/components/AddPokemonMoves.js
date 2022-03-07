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

function AddPokemonMoves({pokemon, moves}) {
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
                <Input id="pokemon" name="pokemon" type="select">
                  {pokemon && pokemon.map((poke, id) => (
                    <option value={poke.name} key={id}>{poke.name}</option>
                  ))}
                  </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="moveName" sm={2}>
                Move Name
              </Label>
              <Col sm={10}>
                 <Input id="move" name="move" type="select">
                  {moves && moves.map((move, id) => (
                    <option value={move.name} key={id}>{move.name}</option>
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

export default AddPokemonMoves;
