import axios, { AxiosResponse } from "axios";

const baseURL = "http://localhost:3000/api/v1";

export interface IToDo {
    id: number;
    description: string;
    isCompleted: boolean;
    isDeleted: boolean;
};

export const getToDos = async (): Promise<AxiosResponse<IToDo[]>> => {
    return await axios.create({ baseURL: baseURL }).get<IToDo[]>(`/todos`);
};