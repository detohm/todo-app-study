import { faSquare } from '@fortawesome/free-regular-svg-icons';
import { faCheckSquare, faTrash } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { IToDoItem } from './ToDoItem.interface';
import styles from './ToDoItem.module.css';

export const ToDoItem = ({ id, index, text, isCompleted, onComplete, onDelete }: IToDoItem) => {

    const handleCompleteClick = (e: React.MouseEvent<SVGSVGElement>) => {
        e.preventDefault();
        onComplete(id, index);
    };

    const handleDeleteClick = (e: React.MouseEvent<SVGSVGElement>) => {
        e.preventDefault();
        onDelete(id);
    };

    return (
        <div className={styles.item}>
            <FontAwesomeIcon
                icon={isCompleted ? faCheckSquare : faSquare}
                className={styles['complete-icon']}
                onClick={handleCompleteClick} />
            <div className={styles.label}>{text}</div>
            <FontAwesomeIcon
                icon={faTrash}
                className={styles['trash-icon']}
                onClick={handleDeleteClick} />
        </div>
    );
};