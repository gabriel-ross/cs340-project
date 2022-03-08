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

function AddPokemonMoves({pokemon, moves}) {
  const [data, setData] = useState({
    pokemon: 1,
    move: 1
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
      pkid: data.pokemon,
      mvid: data.move
    };
    await axios.post("/pokemon/moves", requestData).then((response) => {
      console.log(response);
      window.location.reload();
    });
  };
  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Pokémon Moves</CardTitle>
        <CardBody>
          <Form onSubmit={handleSubmit}>
            <FormGroup row>
              <Label for="pokemonName" sm={2}>
                Pokémon Name
              </Label>
              <Col sm={10}>
                <Input id="pokemon" name="pokemon" type="select" value={data.pokemon}
                  onChange={handleChange}>
                  {pokemon && pokemon.map((poke, id) => (
                    <option value={poke.id} key={id}>{poke.name}</option>
                  ))}
                  </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="moveName" sm={2}>
                Move Name
              </Label>
              <Col sm={10}>
                 <Input id="move" name="move" type="select" value={data.move}
                  onChange={handleChange}>
                  {moves && moves.map((move, id) => (
                    <option value={move.id} key={id}>{move.name}</option>
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

export default AddPokemonMoves;
