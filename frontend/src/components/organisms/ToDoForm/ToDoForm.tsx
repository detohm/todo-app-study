import React, { useState } from 'react';
import { IToDoForm } from './ToDoForm.interface';
import styles from './ToDoForm.module.css';

const ToDoForm = ({
    onSubmit,
    placeHolderText = "Input here and press enter."
}: IToDoForm) => {

    const [value, setValue] = useState("");

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (value === "") {
            return;
        }
        onSubmit(value);
        setValue("");
    };

    return (
        <form onSubmit={handleSubmit} className={styles.form} >
            <input
                className={styles['input-text']}
                type="text"
                value={value}
                placeholder={placeHolderText}
                onChange={(e) => setValue(e.target.value)} />
        </form>
    )
};

export default ToDoForm;