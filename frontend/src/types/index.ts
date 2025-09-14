export interface Todo {
  id: number;
  title: string;
  description: string;
  completed: boolean;
  priority: 'low' | 'medium' | 'high';
  created_at: string;
  updated_at: string;
}

export interface CreateTodo {
  title: string;
  description: string;
  priority: 'low' | 'medium' | 'high';
}

export interface UpdateTodo {
  title: string;
  description: string;
  priority: 'low' | 'medium' | 'high';
  completed: boolean;
}