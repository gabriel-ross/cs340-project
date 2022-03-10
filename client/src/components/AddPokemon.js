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
  FormText,
  Label,
} from "reactstrap";

function AddPokemon({ types, generations }) {
  const [data, setData] = useState({
    id: "",
    name: "",
    primaryType: "Bug",
    secondaryType: "None",
    generation: "I",
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
      id: data.id,
      name: data.name,
      primaryType: data.primaryType,
      secondaryType: data.secondaryType === "None" ? "" : data.secondaryType,
      generation: data.generation,
    };
    await axios.post(`/pokemon/${data.id}`, requestData).then((response) => {
      console.log(response);
      window.location.reload();
    });
  };

  return (
    <Card color="light">
      <CardBody>
        <CardTitle tag="h5">Add Pokémon</CardTitle>
        <CardBody>
          <Form onSubmit={handleSubmit}>
            <FormGroup row>
              <Label for="name" sm={2}>
                PokeDex #
              </Label>
              <Col sm={10}>
                <Input
                  id="id"
                  name="id"
                  placeholder="PokéDex Number"
                  type="number"
                  value={data.id}
                  onChange={handleChange}
                />
                <FormText>
                  We recommend adding a Pokémon by its{" "}
                  <a href="https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_National_Pok%C3%A9dex_number">
                    correct Pokedex id
                  </a>
                  .
                </FormText>
              </Col>
            </FormGroup>
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
                  value={data.name}
                  onChange={handleChange}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="primaryType" sm={2}>
                Primary Type
              </Label>
              <Col sm={10}>
                <Input
                  id="primaryType"
                  name="primaryType"
                  type="select"
                  value={data.primaryType}
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
            <FormGroup row>
              <Label for="secondaryType" sm={2}>
                Secondary Type
              </Label>
              <Col sm={10}>
                <Input
                  id="secondaryType"
                  name="secondaryType"
                  type="select"
                  value={data.secondaryType}
                  onChange={handleChange}
                >
                  <option value="None">None</option>
                  {types &&
                    types.map((type, id) => (
                      <option value={type.name} key={id}>
                        {type.name}
                      </option>
                    ))}
                </Input>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="generation" sm={2}>
                Generation
              </Label>
              <Col sm={10}>
                <Input
                  id="generation"
                  name="generation"
                  type="select"
                  value={data.generation}
                  onChange={handleChange}
                >
                  {generations &&
                    generations.map((generation, id) => (
                      <option value={generation.name} key={id}>
                        {generation.name}
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

export default AddPokemon;
