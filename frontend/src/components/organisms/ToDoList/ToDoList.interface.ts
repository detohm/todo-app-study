import { IToDo } from "../../../App.interface";

export interface ITodoList {
    toDos: IToDo[];
    onItemDelete: (id: number) => void;
};