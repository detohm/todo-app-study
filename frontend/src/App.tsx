import React, { useEffect, useState } from 'react';
import { IToDo } from './App.interface';
import styles from './App.module.css';
import ToDoForm from "./components/organisms/ToDoForm/ToDoForm";
import ToDoList from "./components/organisms/ToDoList/ToDoList";

const App = () => {

  const [toDos, setToDos] = useState<IToDo[]>([]);

  // TODO - remove mock data to retrieve from server
  useEffect(() => {
    setToDos([
      { text: "read books", id: 1 },
      { text: "practise programming", id: 2 },
      { text: "work out", id: 3 }
    ]);
  }, []);

  const handleSubmit = (newToDoText: string) => {
    setToDos((prevToDos) => [
      ...prevToDos,
      {
        text: newToDoText,
        id: Date.now() // TODO - change to persistance layer
      }
    ]);
  };

  const handleDelete = (id: number) => {
    setToDos(toDos.filter((item) => item.id !== id));
  };

  return (
    <div className={styles.app} >
      <ToDoForm onSubmit={handleSubmit} />
      <ToDoList toDos={toDos} onItemDelete={handleDelete} />
    </div>
  );
}

export default App;
