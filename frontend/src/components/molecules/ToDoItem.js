import React from 'react';
import { IconButton } from '../atoms/IconButton';
export const ToDoItem = ({ label }) => {
    return (
        <div>
            <IconButton />
            <label>{label}</label>
            <IconButton />
        </div>
    );
};

ToDoItem.defaultProps = {
    label: "[item]"
}