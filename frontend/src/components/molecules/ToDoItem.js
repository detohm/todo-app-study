import React from 'react';
import { IconButton } from '../atoms/IconButton';
export const ToDoItem = ({ text, id, onDelete }) => {

    return (
        <div>
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