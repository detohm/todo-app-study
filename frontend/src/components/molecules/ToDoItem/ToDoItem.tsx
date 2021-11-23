import React from 'react';
import { IconButton } from '../../atoms/IconButton/IconButton';
import { IToDoItem } from './ToDoItem.interface';

const style = {
    border: "1px solid yellow",
    padding: "10px"
};

export const ToDoItem = ({ text, id, onDelete }: IToDoItem) => {

    return (
        <div style={style}>
            <IconButton id={id} label="complete" />
            <label>{text}</label>
            <IconButton id={id} label="delete" onClick={onDelete} />
        </div>
    );
};

ToDoItem.defaultProps = {
    text: "[item]",
    id: null
}