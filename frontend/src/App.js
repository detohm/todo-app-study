import React from 'react';
import { Header } from "./components/organisms/Header";
import { ToDoList } from "./components/organisms/ToDoList";
import "./style.scss";

function App() {
  return (
    <div className="app">
      <Header />
      <ToDoList />

    </div>
  );
}

export default App;
