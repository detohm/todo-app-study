import React from 'react';
import { ToDoItem } from '../../molecules/ToDoItem/ToDoItem';
import { ITodoList } from './ToDoList.interface';
import styles from './ToDoList.module.css';

const ToDoList = ({ toDos, onItemDelete }: ITodoList) => {
    return (
        <div className={styles.list}>
            {toDos.map((toDo) =>
            (<ToDoItem
                key={toDo.id}
                id={toDo.id}
                text={toDo.description}
                onDelete={onItemDelete}
            />))}
        </div>
    );
};

export default ToDoList;