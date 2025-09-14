# Claude Context: Todo Management Application

## Project Overview
A full-stack Todo/Task Management application with React TypeScript frontend and Go backend. Features a clean, modern UI with priority-based visual indicators, CRUD operations, and real-time updates.

## Architecture

### Frontend (`frontend/`)
- **Framework**: React 18 with TypeScript
- **Build Tool**: Vite
- **Styling**: Custom CSS (removed Tailwind due to PostCSS conflicts)
- **HTTP Client**: Axios
- **Icons**: Lucide React
- **Development Server**: http://localhost:5173

### Backend (`backend/`)
- **Language**: Go
- **Framework**: Gin web framework
- **CORS**: gin-contrib/cors
- **Storage**: In-memory (for simplicity)
- **Server Port**: http://localhost:8080

## Key Features

### Todo Management
- ✅ Create, read, update, delete todos
- ✅ Toggle completion status
- ✅ Priority levels (Low, Medium, High) with visual indicators
- ✅ Edit todos inline
- ✅ Filter by status (All, Active, Completed)
- ✅ Progress tracking display

### Visual Enhancements
- ✅ Priority-based left border colors (Red/Yellow/Green)
- ✅ Priority icons in badges (AlertCircle/ChevronUp/Circle)
- ✅ Color-coded priority badges
- ✅ Compact icon buttons with proper hover states
- ✅ Loading states and error handling
- ✅ Responsive design

## Directory Structure

```
openchoreo-app-sample/
├── frontend/                 # React frontend application
│   ├── src/
│   │   ├── components/      # React components
│   │   │   ├── AddTodo.tsx  # Add new todo form component
│   │   │   └── TodoItem.tsx # Individual todo card component
│   │   ├── services/        # API service layer
│   │   │   └── api.ts       # Axios-based API client
│   │   ├── types/           # TypeScript type definitions
│   │   │   └── index.ts     # Todo-related interfaces
│   │   ├── App.tsx          # Main application component
│   │   ├── main.tsx         # Application entry point
│   │   └── index.css        # Global styles (custom CSS)
│   ├── .env                 # Environment variables
│   ├── .env.example         # Environment template
│   ├── package.json         # Frontend dependencies
│   ├── Dockerfile          # Frontend container configuration
│   ├── .dockerignore       # Docker build optimization
│   └── nginx.conf          # Nginx configuration for production
├── backend/                 # Go backend application
│   ├── main.go             # Main application file with all endpoints
│   ├── go.mod              # Go module file
│   ├── go.sum              # Go dependencies
│   ├── Dockerfile          # Backend container configuration
│   └── .dockerignore       # Docker build optimization
├── docker-compose.yml      # Multi-service orchestration
├── README.md               # Project documentation
├── DOCKER.md               # Docker deployment guide
└── CLAUDE.md               # This context file
```

## API Endpoints

### Backend REST API (`http://localhost:8080`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/v1/todos` | Get all todos |
| GET | `/api/v1/todos/:id` | Get specific todo |
| POST | `/api/v1/todos` | Create new todo |
| PUT | `/api/v1/todos/:id` | Update todo |
| DELETE | `/api/v1/todos/:id` | Delete todo |
| PUT | `/api/v1/todos/:id/toggle` | Toggle completion status |

### Todo Data Structure
```go
type Todo struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    Priority    string    `json:"priority"` // "low", "medium", "high"
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Development Setup

### Running the Application

#### Option 1: Docker (Recommended for Production)

1. **Using Docker Compose** (simplest):
   ```bash
   docker-compose up --build
   ```
   - Frontend: http://localhost:3000
   - Backend: http://localhost:8080

2. **Individual Docker containers**:
   ```bash
   # Backend
   cd backend
   docker build -t todo-backend .
   docker run -p 8080:8080 todo-backend
   
   # Frontend
   cd frontend
   docker build -t todo-frontend .
   docker run -p 3000:8080 -e VITE_API_URL=http://localhost:8080 todo-frontend
   ```

#### Option 2: Local Development

1. **Start Backend**:
   ```bash
   cd backend
   go run main.go
   ```
   Backend runs on http://localhost:8080

2. **Start Frontend**:
   ```bash
   cd frontend
   npm run dev
   ```
   Frontend runs on http://localhost:5173

### Environment Configuration
- Frontend uses `VITE_API_URL` environment variable
- Default backend URL: `http://localhost:8080`
- See `.env.example` for configuration template

## Key Components

### TodoItem.tsx
- Individual todo card with priority indicators
- Inline editing capability
- Priority-based visual styling (borders, colors, icons)
- Compact action buttons (edit/delete)

### AddTodo.tsx
- Expandable form for creating new todos
- Priority selection dropdown
- Form validation and submission

### App.tsx
- Main application container
- State management for todos
- API integration
- Filtering and display logic

## Styling Approach

### Custom CSS Instead of Tailwind
- **Issue**: PostCSS plugin conflicts with Tailwind CSS
- **Solution**: Created comprehensive custom CSS with all required classes
- **Benefit**: No build tool conflicts, full control over styling

### Priority Visual Indicators
```css
/* High Priority */
.border-l-red-500 { border-left-color: #ef4444; }
.bg-red-100 { background-color: #fee2e2; }
.text-red-800 { color: #991b1b; }

/* Medium Priority */
.border-l-yellow-500 { border-left-color: #eab308; }
.bg-yellow-100 { background-color: #fef3c7; }
.text-yellow-800 { color: #92400e; }

/* Low Priority */
.border-l-green-500 { border-left-color: #22c55e; }
.bg-green-100 { background-color: #dcfce7; }
.text-green-800 { color: #166534; }
```

### Button Styling
```css
.icon-button {
  width: 2rem !important;
  height: 2rem !important;
  padding: 0.375rem !important;
  cursor: pointer !important;
}

button {
  cursor: pointer;
}
```

## Known Issues & Solutions

### 1. Tailwind CSS PostCSS Conflicts
- **Problem**: `@tailwindcss/postcss` plugin compatibility issues
- **Solution**: Removed Tailwind, implemented equivalent custom CSS
- **Status**: ✅ Resolved

### 2. Module Import/Export Issues
- **Problem**: TypeScript interface import errors
- **Solution**: Created centralized types directory with proper exports
- **Status**: ✅ Resolved

### 3. Button Hover States
- **Problem**: No pointer cursor on button hover
- **Solution**: Added `cursor: pointer` to all buttons and icon buttons
- **Status**: ✅ Resolved

## Recent Enhancements

### UI/UX Improvements
1. **Compact Icon Buttons**: Reduced padding from `p-2` to `p-1`, fixed button sizing
2. **Priority Visual Indicators**: Added left border colors, priority icons, enhanced badges
3. **Cursor States**: Added pointer cursor to all interactive elements
4. **Color-coded Priority System**: Red (High), Yellow (Medium), Green (Low)

### Technical Improvements
1. **Type Safety**: Centralized TypeScript interfaces
2. **Error Handling**: Comprehensive error states and user feedback
3. **Loading States**: Loading spinners and state management
4. **Code Organization**: Separated concerns with services and components

### Docker Containerization
1. **Multi-stage Builds**: Optimized image sizes (Backend: ~15-20MB, Frontend: ~25-30MB)
2. **Production Ready**: Non-root user execution, health checks, security headers
3. **Service Orchestration**: Docker Compose with proper dependencies and networking
4. **Development Workflow**: Single-command deployment with `docker-compose up --build`

## Future Enhancement Opportunities

### Functionality
- [ ] User authentication and authorization
- [ ] Database persistence (PostgreSQL/MongoDB)
- [ ] Due dates and reminders
- [ ] Categories/tags for todos
- [ ] Search and advanced filtering
- [ ] Drag-and-drop reordering
- [ ] Bulk operations (delete multiple, mark all complete)

### UI/UX
- [ ] Dark mode support
- [ ] Mobile-first responsive design improvements
- [ ] Keyboard shortcuts
- [ ] Accessibility enhancements (ARIA labels, screen reader support)
- [ ] Animation and transitions
- [ ] Rich text editor for descriptions

### Technical
- [ ] Unit and integration tests
- [ ] PWA capabilities (offline support)
- [ ] Real-time updates with WebSockets
- [ ] API rate limiting and caching
- [ ] CI/CD pipeline setup
- [ ] Kubernetes deployment manifests
- [ ] Database persistence (PostgreSQL/MongoDB)
- [ ] Redis caching layer

## Docker Deployment

### Container Architecture

#### Frontend Container
- **Base Images**: Node.js 24 (build) → Nginx 1.25 (runtime)
- **Features**: Multi-stage build, SPA routing, gzip compression, security headers
- **Size**: ~25-30MB optimized
- **Port**: 8080 (mapped to 3000 in docker-compose)
- **Health Check**: `/health` endpoint

#### Backend Container  
- **Base Images**: Go 1.21 (build) → Alpine (runtime)
- **Features**: Static binary, non-root user, minimal attack surface
- **Size**: ~15-20MB optimized  
- **Port**: 8080
- **Health Check**: `/health` endpoint

### Docker Configuration Files

1. **`backend/Dockerfile`**: Go multi-stage build with Alpine runtime
2. **`frontend/Dockerfile`**: Node.js build + Nginx serving
3. **`frontend/nginx.conf`**: Custom Nginx config with SPA support
4. **`docker-compose.yml`**: Service orchestration
5. **`.dockerignore`**: Build optimization for both services

### Key Docker Features

- ✅ **Multi-stage builds** for optimized image sizes
- ✅ **Security**: Non-root user execution in both containers
- ✅ **Health checks** for container monitoring
- ✅ **Production ready** with proper error handling
- ✅ **SPA routing** support with Nginx fallback
- ✅ **Static asset caching** and gzip compression
- ✅ **Security headers** (CSP, XSS protection, etc.)

### Docker Commands

```bash
# Quick start
docker-compose up --build

# Background mode
docker-compose up -d

# View logs
docker-compose logs -f

# Stop and cleanup
docker-compose down

# Individual service builds
docker-compose build backend
docker-compose build frontend
```

See `DOCKER.md` for comprehensive Docker deployment guide.

## Development Commands

### Frontend
```bash
npm run dev      # Start development server
npm run build    # Build for production
npm run preview  # Preview production build
```

### Backend
```bash
go run main.go   # Start development server
go build         # Build for production
go mod tidy      # Clean up dependencies
```

## Dependencies

### Frontend Key Dependencies
- `react@^18.2.0` - React framework
- `typescript@^5.0.0` - TypeScript support
- `vite@^7.1.5` - Build tool
- `axios@^1.6.0` - HTTP client
- `lucide-react@^0.263.1` - Icons

### Backend Key Dependencies
- `github.com/gin-gonic/gin` - Web framework
- `github.com/gin-contrib/cors` - CORS middleware

## Important Notes for Future Development

1. **No Tailwind CSS**: Project uses custom CSS due to PostCSS conflicts
2. **In-Memory Storage**: Backend uses in-memory storage, data resets on restart
3. **CORS Enabled**: Backend allows all origins for development
4. **Environment Variables**: Use `.env` file for frontend configuration
5. **Type Safety**: Always update TypeScript interfaces when changing data structures
6. **Priority System**: Maintain consistency in priority colors and icons across components

## Testing the Application

### Basic Flow
1. Start both backend and frontend servers
2. Navigate to http://localhost:5173
3. Create a new todo with different priority levels
4. Test editing, completion toggle, and deletion
5. Verify filtering functionality (All/Active/Completed)
6. Check visual priority indicators (borders, colors, icons)

### API Testing
```bash
# Health check
curl http://localhost:8080/health

# Get all todos
curl http://localhost:8080/api/v1/todos

# Create new todo
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Todo","description":"Test Description","priority":"high"}'
```

This context should provide all necessary information for future development and enhancements of the Todo Management application.