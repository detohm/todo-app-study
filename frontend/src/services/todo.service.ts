import axios, { AxiosResponse } from "axios";

const baseURL = "/api/v1";
export interface IToDo {
    id: number;
    description: string;
    isCompleted: boolean;
    isDeleted: boolean;
};

export const getToDos = async (): Promise<AxiosResponse<IToDo[]>> => {
    return await axios.create({ baseURL: baseURL })
        .get<IToDo[]>(`/todos`);
};

export const createToDos = async (data: IToDo):
    Promise<AxiosResponse<number>> => {
    return await axios
        .create({ baseURL: baseURL }).post<number>(`/todo`, data);
}

export const completeToDo = async (id: number): Promise<AxiosResponse> => {
    return await axios.create({ baseURL: baseURL })
        .patch(`/todo/${id}/complete`);
}

export const deleteToDo = async (id: number): Promise<AxiosResponse> => {
    return await axios.create({ baseURL: baseURL })
        .delete(`/todo/${id}`);
}
