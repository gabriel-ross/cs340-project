import axios from "axios";
import React, { useState } from "react";
import {
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Col,
  Input,
  Form,
  FormGroup,
  Label,
  FormText,
} from "reactstrap";

function EditPokemon({
  id,
  name,
  primaryType,
  secondaryType,
  generation,
  types,
  generations,
}) {
  const [modal, setModal] = React.useState(false);
  const toggle = () => setModal(!modal);

  const [data, setData] = useState({
    id: id,
    name: name,
    primaryType: primaryType,
    secondaryType: secondaryType,
    generation: generation,
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
      secondaryType: data.secondaryType,
      generation: data.generation,
    };
    await axios.patch(`/pokemon/${id}`, requestData).then((response) => {
      window.location.reload();
    });
  };
  return (
    <React.Fragment>
      <Button color="primary" outline size="sm" onClick={toggle}>
        Edit
      </Button>
      <Modal size="lg" isOpen={modal} toggle={toggle}>
        <ModalHeader toggle={toggle}>Edit Pokemon</ModalHeader>
        <ModalBody>
          <Form id="editPokemon" onSubmit={handleSubmit}>
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
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button color="primary" form="editPokemon" onClick={handleSubmit}>
            Update
          </Button>{" "}
          <Button onClick={toggle}>Cancel</Button>
        </ModalFooter>
      </Modal>
    </React.Fragment>
  );
}

export default EditPokemon;
