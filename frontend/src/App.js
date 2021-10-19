import React, { useEffect, useState } from 'react';
import ToDoForm from "./components/organisms/ToDoForm";
import ToDoList from "./components/organisms/ToDoList";
import "./style.scss";

const style = {
  width: "300px",
  margin: "0 auto",
  position: "relative",
  top: 60
};

const App = () => {

  const [toDos, setToDos] = useState([]);

  // TODO - remove mock data to retrieve from server
  useEffect(() => {
    setToDos([
      { text: "read books", id: 1 },
      { text: "practise programming", id: 2 },
      { text: "work out", id: 3 }
    ]);
  }, []);

  const handleSubmit = (newToDoText) => {
    setToDos((prevToDos) => [
      ...prevToDos,
      {
        text: newToDoText,
        id: Date.now() // TODO - change to persistance layer
      }
    ]);
  };

  const handleDelete = (id) => {
    setToDos(toDos.filter((item) => item.id !== id));
  };

  return (
    <div className="app" style={style} >
      <ToDoForm onSubmit={handleSubmit} />
      <ToDoList toDos={toDos} onItemDelete={handleDelete} />
    </div>
  );
}

export default App;
