# Todo Management Application

A full-stack Todo/Task Management application with **two frontend implementations**: React Client-Side Rendering (CSR) and Go Server-Side Rendering (SSR), both sharing the same Go backend.

## Features

- ✅ Create, read, update, and delete todos
- ✅ Mark todos as complete/incomplete
- ✅ Priority levels (Low, Medium, High) with visual indicators
- ✅ Filter todos by status (All, Active, Completed)
- ✅ Clean and responsive UI
- ✅ Docker containerization with multi-architecture support
- ✅ Environment variable configuration
- ✅ Health checks and monitoring

## Architecture Options

Choose between two frontend implementations:

### Option 1: Client-Side Rendering (CSR)
- **Frontend**: React 18 with TypeScript
- **Pros**: Rich interactivity, modern development experience
- **Best for**: Interactive web applications, SPAs

### Option 2: Server-Side Rendering (SSR) 
- **Frontend**: Go with HTML templates
- **Pros**: Faster initial load, better SEO, simpler deployment
- **Best for**: Content-focused applications, better performance

## Technology Stack

### Frontend Options

#### CSR Frontend (`frontend/`)
- **React 18** with TypeScript
- **Vite** for development and build
- **Custom CSS** (removed Tailwind due to PostCSS conflicts)
- **Axios** for API calls
- **Lucide React** for icons

#### SSR Frontend (`frontend-ssr-go/`)
- **Go** with Gorilla Mux router
- **HTML templates** with Go templating
- **Custom CSS** matching React styling
- **HTTP client** for API integration
- **Lucide icons** via CDN

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

#### Option 1: React CSR Frontend

1. **Start the Backend Server**
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```
   The backend will start on `http://localhost:8080`

2. **Start the React Frontend**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```
   The React frontend will start on `http://localhost:5173`

3. **Open your browser** and navigate to `http://localhost:5173`

#### Option 2: Go SSR Frontend

1. **Start the Backend Server** (same as above)
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```

2. **Start the Go SSR Frontend**
   ```bash
   cd frontend-ssr-go
   go mod tidy
   API_URL=http://localhost:8080 PORT=3001 go run main.go
   ```
   The Go SSR frontend will start on `http://localhost:3001`

3. **Open your browser** and navigate to `http://localhost:3001`

#### Docker Deployment

**For React CSR:**
```bash
docker-compose up --build
```
- React Frontend: http://localhost:3000
- Backend API: http://localhost:8080

**For Go SSR:**
```bash
docker-compose -f docker-compose-ssr.yml up --build
```
- Go SSR Frontend: http://localhost:3001
- Backend API: http://localhost:8080

## Configuration

### Environment Variables

#### React CSR Frontend
- Copy `frontend/.env.example` to `frontend/.env`
- Set `VITE_API_URL` to your backend URL (default: `http://localhost:8080`)

#### Go SSR Frontend
- `API_URL`: Backend API URL (default: `http://localhost:8080`)
- `PORT`: Frontend server port (default: `3001`)

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

#### React CSR Frontend
```bash
cd frontend
npm run dev      # Start development server
npm run build    # Build for production
npm run preview  # Preview production build
```

#### Go SSR Frontend
```bash
cd frontend-ssr-go
go run main.go   # Start development server
go build         # Build for production
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
├── frontend/                    # React CSR frontend
│   ├── src/
│   │   ├── components/         # React components
│   │   │   ├── AddTodo.tsx     # Add todo form
│   │   │   └── TodoItem.tsx    # Todo item card
│   │   ├── services/           # API service layer
│   │   │   └── api.ts          # Axios HTTP client
│   │   ├── types/              # TypeScript definitions
│   │   │   └── index.ts        # Todo interfaces
│   │   ├── App.tsx             # Main React component
│   │   ├── main.tsx            # Application entry point
│   │   └── index.css           # Custom CSS styles
│   ├── .env.example            # Environment template
│   ├── package.json            # Dependencies
│   ├── Dockerfile              # Container build
│   └── nginx.conf              # Production config
├── frontend-ssr-go/            # Go SSR frontend
│   ├── templates/              # HTML templates
│   │   └── index.html          # Main template
│   ├── static/                 # Static assets
│   │   └── css/style.css       # CSS styles
│   ├── main.go                 # HTTP server
│   ├── go.mod                  # Go module
│   └── Dockerfile              # Container build
├── backend/                    # Shared Go API backend
│   ├── main.go                 # RESTful API server
│   ├── go.mod                  # Go module
│   ├── go.sum                  # Dependencies
│   └── Dockerfile              # Container build
├── docker-compose.yml          # CSR deployment
├── docker-compose-ssr.yml      # SSR deployment
├── CLAUDE.md                   # Development context
├── DOCKER.md                   # Docker deployment guide
└── README.md                   # Project documentation
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
- Clean, modern design with custom CSS
- Responsive layout that works on desktop and mobile
- Loading states and error handling
- Progress indicator showing completed vs total tasks
- Priority-based visual indicators (colors, icons, borders)
- Form validation and user feedback

### Performance Comparison

| Feature | React CSR | Go SSR |
|---------|-----------|--------|
| **Initial Load** | Slower (JS bundle) | ⚡ Faster (pre-rendered) |
| **SEO** | Limited | 🏆 Excellent |
| **Interactivity** | Rich client-side | Traditional forms |
| **Deployment** | Build process required | Single binary |
| **Caching** | Client-side | Server/CDN friendly |
| **Development** | Modern tooling | Simple templates |

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