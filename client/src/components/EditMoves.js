import axios from "axios";
import React, { useState } from "react";
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, Col, Input, Form, FormGroup, Label } from "reactstrap";

function EditMoves({id, name, type, types}) {
  const [modal, setModal] = React.useState(false);
  const toggle = () => setModal(!modal);

  const [data, setData] = useState({
    name: name,
    type: type
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
      name: data.name,
      type: data.type
    };
    await axios.patch(`/moves/${id}`, requestData).then((response) => {
      window.location.reload()
    });
  };
  return (
    <React.Fragment>
      <Button color="primary" outline size="sm" onClick={toggle}>
        Edit
      </Button>
      <Modal isOpen={modal} toggle={toggle}>
        <ModalHeader toggle={toggle}>Edit Move</ModalHeader>
        <ModalBody>
          <Form id="editMoves" onSubmit={handleSubmit}>
            <FormGroup row>
              <Label for="name" sm={2}>
                Name
              </Label>
              <Col sm={10}>
                <Input
                  id="name"
                  name="name"
                  placeholder="Move Name"
                  type="text"
                  value={data.name}
                  onChange={handleChange}
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="type" sm={2}>
                Move Type
              </Label>
              <Col sm={10}>
                <Input
                  id="type"
                  name="type"
                  type="select"
                  value={data.type}
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
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button color="primary" form="editMoves" onClick={handleSubmit}>
            Update
          </Button>{" "}
          <Button onClick={toggle}>Cancel</Button>
        </ModalFooter>
      </Modal>
    </React.Fragment>
  );
}

export default EditMoves;
