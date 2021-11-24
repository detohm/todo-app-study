import React from 'react';
import { ToDoItem } from '../../molecules/ToDoItem/ToDoItem';
import { ITodoList } from './ToDoList.interface';
import styles from './ToDoList.module.css';

const ToDoList = ({ toDos, onItemComplete, onItemDelete }: ITodoList) => {
    return (
        <div className={styles.list}>
            {toDos.map((toDo, index) =>
            (<ToDoItem
                key={toDo.id}
                id={toDo.id}
                index={index}
                text={toDo.description}
                isCompleted={toDo.isCompleted}
                onComplete={onItemComplete}
                onDelete={onItemDelete}
            />))}
        </div>
    );
};

export default ToDoList;