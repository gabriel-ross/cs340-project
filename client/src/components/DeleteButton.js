import axios from "axios";
import { Button } from "reactstrap";

function DeleteButton({ route }) {
  const handleDelete = async (event) => {
    event.preventDefault();
    await axios.delete(route).then((response) => {
      window.location.reload()
    });
  };

  return (
    <Button color="primary" outline size="sm" onClick={handleDelete}>
      Delete
    </Button>
  );
}

export default DeleteButton;
