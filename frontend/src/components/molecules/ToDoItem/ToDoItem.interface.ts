export interface IToDoItem {
    text: string;
    id: number;
    onDelete: (id: number) => void;
};