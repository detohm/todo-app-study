import React, { useState } from 'react';
import { IToDoForm } from './ToDoForm.interface';
import styles from './ToDoForm.module.css';
const inputStyle = {
    width: "280px"
};

const ToDoForm = ({
    onSubmit,
    placeHolderText = "input here !"
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
            <input style={inputStyle} type="text" value={value} placeholder={placeHolderText} onChange={(e) => setValue(e.target.value)} />
        </form>
    )
};

ToDoForm.defaultProps = {
    placeHolderText: "input here !"
};

export default ToDoForm;