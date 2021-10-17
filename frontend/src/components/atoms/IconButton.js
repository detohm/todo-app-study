import React from 'react';

export const IconButton = ({ label, id, onClick }) => {

    const handleClick = (e) => {
        e.preventDefault();
        onClick(id);
    }
    return <button onClick={handleClick}>{label}</button>;

};

IconButton.defaultProps = {
    label: "",
    id: null,
    onClick: () => { },
};