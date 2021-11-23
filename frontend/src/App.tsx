import React, { useEffect, useState } from 'react';
import styles from './App.module.css';
import ToDoForm from "./components/organisms/ToDoForm/ToDoForm";
import ToDoList from "./components/organisms/ToDoList/ToDoList";
import { createToDos, deleteToDo, getToDos, IToDo } from './services/todo.service';

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
    let todo: IToDo = {
      id: 0,
      description: newToDoText,
      isCompleted: false,
      isDeleted: false
    };

    createToDos(todo).then((res) => {
      if (res.data) {
        setToDos((prev: IToDo[]) => {
          todo.id = res.data
          return [...prev, todo];
        });
      }
    }).catch((e) => {
      console.log(e);
    });
  };

  const handleDelete = (id: number) => {
    deleteToDo(id).then((res) => {
      setToDos(toDos.filter((item) => item.id !== id));
    }).catch((e) => {
      console.log(e);
    });
  };

  return (
    <div className={styles.app} >
      <ToDoForm onSubmit={handleSubmit} />
      <ToDoList toDos={toDos} onItemDelete={handleDelete} />
    </div>
  );
}

export default App;
