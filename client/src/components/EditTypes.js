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
} from "reactstrap";

function EditTypes({ id, name }) {
  const [modal, setModal] = React.useState(false);
  const toggle = () => setModal(!modal);

  const [data, setData] = useState({
    name: name,
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
    };
    await axios.patch(`/types/${id}`, requestData).then((response) => {
      window.location.reload();
    });
  };
  return (
    <React.Fragment>
      <Button color="primary" outline size="sm" onClick={toggle}>
        Edit
      </Button>
      <Modal isOpen={modal} toggle={toggle}>
        <ModalHeader toggle={toggle}>Edit Types</ModalHeader>
        <ModalBody>
          <Form id="editTypes" onSubmit={handleSubmit}>
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
                  value={data.name}
                  onChange={handleChange}
                />
              </Col>
            </FormGroup>
          </Form>
        </ModalBody>
        <ModalFooter>
          <Button color="primary" form="editTypes" onClick={handleSubmit}>
            Update
          </Button>{" "}
          <Button onClick={toggle}>Cancel</Button>
        </ModalFooter>
      </Modal>
    </React.Fragment>
  );
}

export default EditTypes;
