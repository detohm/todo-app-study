import React, { useEffect, useState } from 'react';
import styles from './App.module.css';
import ToDoForm from "./components/organisms/ToDoForm/ToDoForm";
import ToDoList from "./components/organisms/ToDoList/ToDoList";
import { getToDos, IToDo } from './services/todo.service';

const App = () => {

  const [toDos, setToDos] = useState<IToDo[]>([]);

  useEffect(() => {

    getToDos().then((res) => {
      if (res.data)
        setToDos(res.data);
    }).catch((e) => {
      console.log(e);
    });

  }, []);

  const handleSubmit = (newToDoText: string) => {
    console.log(newToDoText)
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
