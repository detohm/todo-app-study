import React, { useState } from 'react';

const style = {
    border: "1px solid blue",
    ["margin-bottom"]: "24px"
};

const inputStyle = {
    width: "280px"
};

const ToDoForm = ({ onSubmit, placeHolderText }) => {

    const [value, setValue] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        if (value === "") {
            return;
        }
        onSubmit(value);
        setValue("");
    };

    return (
        <form onSubmit={handleSubmit} style={style} >
            <input style={inputStyle} type="text" value={value} placeholder={placeHolderText} onChange={(e) => setValue(e.target.value)} />
        </form>
    )
};

ToDoForm.defaultProps = {
    placeHolderText: "input here !"
};

export default ToDoForm;