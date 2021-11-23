import React from 'react';
import { IIconButton } from './IconButton.interface';

export const IconButton = ({ label, id, onClick }: IIconButton) => {

    const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        if (onClick) {
            onClick(id);
        }
    }
    return <button onClick={handleClick}>{label}</button>;

};
