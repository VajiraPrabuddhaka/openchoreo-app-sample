import axios from 'axios';
import type { Todo, CreateTodo, UpdateTodo } from '../types';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: `${API_BASE_URL}/api/v1`,
  headers: {
    'Content-Type': 'application/json',
  },
});

export type { Todo, CreateTodo, UpdateTodo };

export const todoAPI = {
  getAllTodos: async (): Promise<Todo[]> => {
    const response = await api.get('/todos');
    return response.data as Todo[];
  },

  getTodoById: async (id: number): Promise<Todo> => {
    const response = await api.get(`/todos/${id}`);
    return response.data as Todo;
  },

  createTodo: async (todo: CreateTodo): Promise<Todo> => {
    const response = await api.post('/todos', todo);
    return response.data as Todo;
  },

  updateTodo: async (id: number, todo: UpdateTodo): Promise<Todo> => {
    const response = await api.put(`/todos/${id}`, todo);
    return response.data as Todo;
  },

  deleteTodo: async (id: number): Promise<void> => {
    await api.delete(`/todos/${id}`);
  },

  toggleTodo: async (id: number): Promise<Todo> => {
    const response = await api.put(`/todos/${id}/toggle`);
    return response.data as Todo;
  },
};