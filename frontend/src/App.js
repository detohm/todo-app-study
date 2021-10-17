import React, { useState } from 'react';
import ToDoForm from "./components/organisms/ToDoForm";
import ToDoList from "./components/organisms/ToDoList";
import "./style.scss";

const App = () => {

  const [toDos, setToDos] = useState([]);

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
    alert(id);
    setToDos(toDos.filter((item) => item.id !== id));
  };

  return (
    <div className="app">
      <ToDoForm onSubmit={handleSubmit} />
      <ToDoList toDos={toDos} onItemDelete={handleDelete} />
    </div>
  );
}

export default App;
