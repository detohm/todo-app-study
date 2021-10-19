import React from 'react';
import { ToDoItem } from '../molecules/ToDoItem';

const style = {
    border: "1px solid red"
};


const ToDoList = ({ toDos, onItemDelete }) => {
    return (
        <div style={style}>
            {toDos.map((toDo) => (<ToDoItem key={toDo.id} id={toDo.id} text={toDo.text} onDelete={onItemDelete} />))}
        </div>
    );
};

ToDoList.defaultProps = {
    toDos: []
}

export default ToDoList;