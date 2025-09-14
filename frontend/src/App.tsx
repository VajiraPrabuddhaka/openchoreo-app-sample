import { useState, useEffect } from 'react';
import { CheckSquare, Filter } from 'lucide-react';
import { todoAPI } from './services/api';
import type { Todo, CreateTodo, UpdateTodo } from './types';
import TodoItem from './components/TodoItem';
import AddTodo from './components/AddTodo';

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [filter, setFilter] = useState<'all' | 'active' | 'completed'>('all');

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      setLoading(true);
      const data = await todoAPI.getAllTodos();
      setTodos(data);
    } catch (err) {
      setError('Failed to load todos. Please make sure the backend server is running.');
      console.error('Error fetching todos:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleAddTodo = async (newTodo: CreateTodo) => {
    try {
      const created = await todoAPI.createTodo(newTodo);
      setTodos([...todos, created]);
    } catch (err) {
      setError('Failed to add todo');
      console.error('Error adding todo:', err);
    }
  };

  const handleToggleTodo = async (id: number) => {
    try {
      const updated = await todoAPI.toggleTodo(id);
      setTodos(todos.map(todo => todo.id === id ? updated : todo));
    } catch (err) {
      setError('Failed to update todo');
      console.error('Error toggling todo:', err);
    }
  };

  const handleUpdateTodo = async (id: number, updatedTodo: UpdateTodo) => {
    try {
      const updated = await todoAPI.updateTodo(id, updatedTodo);
      setTodos(todos.map(todo => todo.id === id ? updated : todo));
    } catch (err) {
      setError('Failed to update todo');
      console.error('Error updating todo:', err);
    }
  };

  const handleDeleteTodo = async (id: number) => {
    try {
      await todoAPI.deleteTodo(id);
      setTodos(todos.filter(todo => todo.id !== id));
    } catch (err) {
      setError('Failed to delete todo');
      console.error('Error deleting todo:', err);
    }
  };

  const getFilteredTodos = () => {
    switch (filter) {
      case 'active':
        return todos.filter(todo => !todo.completed);
      case 'completed':
        return todos.filter(todo => todo.completed);
      default:
        return todos;
    }
  };

  const filteredTodos = getFilteredTodos();
  const completedCount = todos.filter(todo => todo.completed).length;
  const totalCount = todos.length;

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading todos...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-4xl mx-auto py-8 px-4">
        <header className="text-center mb-8">
          <div className="flex items-center justify-center gap-3 mb-4">
            <CheckSquare className="h-8 w-8 text-blue-500" />
            <h1 className="text-3xl font-bold text-gray-900">Todo Manager</h1>
          </div>
          <p className="text-gray-600">Stay organized and get things done</p>
          
          {totalCount > 0 && (
            <div className="mt-4 text-sm text-gray-500">
              {completedCount} of {totalCount} tasks completed
            </div>
          )}
        </header>

        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg">
            <p className="text-red-700">{error}</p>
            <button 
              onClick={() => setError(null)}
              className="mt-2 text-red-600 underline text-sm"
            >
              Dismiss
            </button>
          </div>
        )}

        <div className="space-y-6">
          <AddTodo onAdd={handleAddTodo} />

          {todos.length > 0 && (
            <div className="flex items-center gap-4 bg-white p-4 rounded-lg shadow-sm">
              <Filter className="h-4 w-4 text-gray-500" />
              <span className="text-sm text-gray-600">Filter:</span>
              <div className="flex gap-2">
                {(['all', 'active', 'completed'] as const).map((filterType) => (
                  <button
                    key={filterType}
                    onClick={() => setFilter(filterType)}
                    className={`px-3 py-1 text-sm rounded-full capitalize ${
                      filter === filterType
                        ? 'bg-blue-100 text-blue-700 border border-blue-200'
                        : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                    } transition-colors`}
                  >
                    {filterType}
                    {filterType === 'active' && ` (${todos.filter(t => !t.completed).length})`}
                    {filterType === 'completed' && ` (${completedCount})`}
                  </button>
                ))}
              </div>
            </div>
          )}

          <div className="space-y-4">
            {filteredTodos.length === 0 ? (
              <div className="text-center py-12">
                <CheckSquare className="h-16 w-16 text-gray-300 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  {filter === 'all' ? 'No todos yet' : `No ${filter} todos`}
                </h3>
                <p className="text-gray-500">
                  {filter === 'all' 
                    ? 'Add your first todo to get started!' 
                    : `You have no ${filter} todos at the moment.`
                  }
                </p>
              </div>
            ) : (
              filteredTodos.map((todo) => (
                <TodoItem
                  key={todo.id}
                  todo={todo}
                  onToggle={handleToggleTodo}
                  onUpdate={handleUpdateTodo}
                  onDelete={handleDeleteTodo}
                />
              ))
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
