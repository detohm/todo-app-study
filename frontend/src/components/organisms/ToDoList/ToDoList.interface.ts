import { IToDo } from "../../../services/todo.service";

export interface ITodoList {
    toDos: IToDo[];
    onItemComplete: (id: number, index: number) => void;
    onItemDelete: (id: number) => void;
};