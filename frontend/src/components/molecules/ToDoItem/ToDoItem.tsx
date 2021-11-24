import { faCheck, faSquare, faTrash } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { IToDoItem } from './ToDoItem.interface';

const style = {
    border: "1px solid yellow",
    padding: "10px"
};

export const ToDoItem = ({ id, index, text, isCompleted, onComplete, onDelete }: IToDoItem) => {

    const handleCompleteClick = (e: React.MouseEvent<SVGSVGElement>) => {
        e.preventDefault();
        console.log("id:" + id + ",index:" + index);
        onComplete(id, index);
    };

    const handleDeleteClick = (e: React.MouseEvent<SVGSVGElement>) => {
        e.preventDefault();
        onDelete(id);
    };

    return (
        <div style={style}>
            <FontAwesomeIcon
                icon={isCompleted ? faCheck : faSquare}
                onClick={handleCompleteClick} />
            <label>{text}</label>
            <FontAwesomeIcon
                icon={faTrash}
                onClick={handleDeleteClick} />
        </div>
    );
};