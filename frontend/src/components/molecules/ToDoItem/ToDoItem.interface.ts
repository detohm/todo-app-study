export interface IToDoItem {
    id: number;
    index: number;
    isCompleted: boolean;
    text: string;

    onComplete: (id: number, index: number) => void;
    onDelete: (id: number) => void;
};