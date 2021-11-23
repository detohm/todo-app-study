import { IToDo } from "../../../services/todo.service";

export interface ITodoList {
    toDos: IToDo[];
    onItemDelete: (id: number) => void;
};