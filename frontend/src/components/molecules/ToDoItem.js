import React from 'react';
import { IconButton } from '../atoms/IconButton';

const style = {
    border: "1px solid yellow",
    padding: "10px"
};

export const ToDoItem = ({ text, id, onDelete }) => {

    return (
        <div style={style}>
            <IconButton label="complete" />
            <label>{text}</label>
            <IconButton id={id} label="delete" onClick={onDelete} />
        </div>
    );
};

ToDoItem.defaultProps = {
    text: "[item]",
    id: null
}