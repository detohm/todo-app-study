import React, { useState } from 'react';

const ToDoForm = ({ onSubmit }) => {

    const [value, setValue] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(value);
        setValue("");
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" value={value} onChange={(e) => setValue(e.target.value)} />
        </form>
    )
};


export default ToDoForm;