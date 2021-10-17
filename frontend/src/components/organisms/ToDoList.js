import React from 'react';
import { ToDoItem } from '../molecules/ToDoItem';

const ToDoList = ({ toDos, onItemDelete }) => {
    return (
        <div>
            {toDos.map((toDo) => (<ToDoItem key={toDo.id} id={toDo.id} text={toDo.text} onDelete={onItemDelete} />))}
        </div>
    );
};

ToDoList.defaultProps = {
    toDos: []
}

export default ToDoList;