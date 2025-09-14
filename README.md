# Todo Management Application

A full-stack Todo/Task Management application built with React (TypeScript) frontend and Go backend.

## Features

- ✅ Create, read, update, and delete todos
- ✅ Mark todos as complete/incomplete
- ✅ Priority levels (Low, Medium, High)
- ✅ Filter todos by status (All, Active, Completed)
- ✅ Clean and responsive UI with Tailwind CSS
- ✅ Real-time updates
- ✅ Environment variable configuration for backend URL

## Technology Stack

### Frontend (`frontend/`)
- **React 18** with TypeScript
- **Vite** for development and build
- **Tailwind CSS** for styling
- **Axios** for API calls
- **Lucide React** for icons

### Backend (`backend/`)
- **Go** with Gin web framework
- **CORS** support for cross-origin requests
- **RESTful API** design
- **In-memory storage** (for simplicity)

## Getting Started

### Prerequisites
- Node.js (v18 or higher)
- Go (v1.19 or higher)
- npm or yarn

### Running the Application

1. **Start the Backend Server**
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```
   The backend will start on `http://localhost:8080`

2. **Start the Frontend Development Server**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```
   The frontend will start on `http://localhost:5173`

3. **Open your browser** and navigate to `http://localhost:5173`

## Configuration

### Environment Variables

The frontend uses environment variables to configure the backend URL:

- Copy `frontend/.env.example` to `frontend/.env`
- Set `VITE_API_URL` to your backend URL (default: `http://localhost:8080`)

### API Endpoints

The backend exposes the following REST API endpoints:

- `GET /health` - Health check
- `GET /api/v1/todos` - Get all todos
- `GET /api/v1/todos/:id` - Get a specific todo
- `POST /api/v1/todos` - Create a new todo
- `PUT /api/v1/todos/:id` - Update a todo
- `DELETE /api/v1/todos/:id` - Delete a todo
- `PUT /api/v1/todos/:id/toggle` - Toggle todo completion status

## Development

### Frontend Development
```bash
cd frontend
npm run dev      # Start development server
npm run build    # Build for production
npm run preview  # Preview production build
```

### Backend Development
```bash
cd backend
go run main.go   # Start development server
go build         # Build for production
```

## Project Structure

```
openchoreo-app-sample/
├── frontend/                 # React frontend application
│   ├── src/
│   │   ├── components/      # Reusable React components
│   │   │   ├── AddTodo.tsx
│   │   │   └── TodoItem.tsx
│   │   ├── services/        # API service layer
│   │   │   └── api.ts
│   │   ├── App.tsx          # Main application component
│   │   ├── main.tsx         # Application entry point
│   │   └── index.css        # Global styles
│   ├── .env.example         # Environment variables template
│   ├── package.json         # Frontend dependencies
│   └── tailwind.config.js   # Tailwind CSS configuration
├── backend/                 # Go backend application
│   ├── main.go             # Main application file
│   ├── go.mod              # Go module file
│   └── go.sum              # Go dependencies
└── README.md               # Project documentation
```

## Features Overview

### Todo Management
- **Add Todos**: Click "Add New Todo" to create new tasks
- **Edit Todos**: Click the edit icon on any todo to modify it
- **Complete Todos**: Click the checkbox to mark todos as complete
- **Delete Todos**: Click the trash icon to remove todos
- **Priority Levels**: Set priority as Low, Medium, or High with color-coded badges

### Filtering
- **All**: View all todos
- **Active**: View only incomplete todos
- **Completed**: View only completed todos

### UI Features
- Clean, modern design with Tailwind CSS
- Responsive layout that works on desktop and mobile
- Loading states and error handling
- Progress indicator showing completed vs total tasks

## Next Steps

To enhance this application further, consider:

- Adding user authentication
- Implementing data persistence (database)
- Adding due dates and reminders
- Categories and tags for todos
- Search functionality
- Drag and drop reordering
- Dark mode support
- Offline support with PWA features