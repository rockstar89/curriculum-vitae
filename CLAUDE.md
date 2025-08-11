# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Frontend (React/TypeScript)
```bash
cd frontend
npm install           # Install dependencies
npm start            # Start development server (http://localhost:3000)
npm run build        # Build for production
npm test             # Run tests
```

### Backend (Go)
```bash
cd backend
go mod tidy          # Install dependencies
go run cmd/main.go   # Start development server (http://localhost:8080)
```

### Full Stack Development
```bash
# Start database only (for development)
cd docker
docker-compose -f docker-compose.dev.yml up -d

# Start full application stack
cd docker
docker-compose up --build

# Database development startup script
./scripts/start.sh   # Starts PostgreSQL + provides development instructions
```

### Testing and Linting
- Frontend uses React Testing Library (npm test)
- Backend uses standard Go testing (go test ./...)
- No specific linting commands configured - uses default React and Go tooling

### Deployment Commands
```bash
# Update dependencies
cd backend
go mod tidy

# Test with PostgreSQL locally
cd docker
docker-compose up --build

# Deploy to Render (automatic via git push with render.yaml)
git push origin main
```

## Architecture Overview

### Full Stack Structure
This is a CV landing page with admin functionality, split into:
- **Frontend**: React/TypeScript SPA with SCSS styling
- **Backend**: Go REST API with JWT authentication
- **Database**: PostgreSQL for all data storage (CV files, user data, metadata)
- **Deployment**: Docker containers with PostgreSQL database

### Backend Architecture (Go + Gin)
- **Entry Point**: `backend/cmd/main.go` - server initialization and routing
- **Handlers**: `backend/internal/handlers/` - HTTP request handlers
  - `auth.go`: Login, token verification, password management
  - `cv.go`: CV upload/download/delete operations
- **Storage Layer**: `backend/internal/storage/` - PostgreSQL-based storage
  - `cv_storage.go`: CV file operations (stored as BYTEA in PostgreSQL)
  - `user_storage.go`: User authentication and management
  - `db.go`: Database connection and initialization
- **Middleware**: CORS and JWT authentication middleware
- **Models**: Request/response structures in `backend/internal/models/`

### Frontend Architecture (React/TypeScript)
- **Routing**: React Router with routes: `/` (home), `/login`, `/admin`
- **State Management**: 
  - Context API for language/theme (no external state library)
  - Custom hooks pattern (`useAuth` for authentication)
- **Authentication**: JWT token stored in localStorage with auto-verification
- **Internationalization**: Built-in i18n system supporting English/Serbian
- **Styling**: SCSS with theme system (dark/light mode)
- **API Layer**: Centralized service class in `services/api.ts`

### Key Components Architecture
- **App.tsx**: Main router and authentication guard
- **Header**: Navigation with theme/language toggles
- **PersonalIntro**: Landing section with CV download CTA
- **Experience/Skills/Education**: Static content sections with JSON data sources
- **Admin**: Protected component for CV management
- **Login**: Authentication form with JWT handling

### Authentication Flow
1. Login stores JWT in localStorage
2. `useAuth` hook auto-verifies token on app load
3. Protected routes check authentication state
4. Token automatically passed to API calls via Authorization header
5. Backend validates JWT and extracts username for operations

### File Upload/Management System
- **PostgreSQL Storage**: All CV files stored as BYTEA in PostgreSQL database
- Only PDF files allowed (10MB limit)
- Single active CV per system (replaces previous on upload)
- Public endpoints for download/view, protected endpoints for management
- **Cloud-Ready**: No filesystem dependencies, perfect for Render/cloud deployments
- **Persistent**: Files survive server hibernation and restarts

### Environment Configuration
- **Backend**: Environment variables for JWT secret, admin credentials, CORS, database
  - `DATABASE_URL`: PostgreSQL connection string (for managed databases like Render)
  - Database connection variables: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`
- **Frontend**: REACT_APP_API_URL for backend URL configuration  
- **Docker**: PostgreSQL database with automatic initialization
- **Render**: Configured with free PostgreSQL plan for $0 cost deployment

### Development vs Production
- **Development**: Uses `docker-compose.dev.yml` with PostgreSQL + pgAdmin (database only)
- **Production**: Uses `docker-compose.yml` with PostgreSQL + backend + frontend
- **Local Development**: Start database with Docker, run backend/frontend independently

### Database Schema
- **PostgreSQL**: All data stored in database tables
- **cv_files**: Stores CV files as BYTEA with metadata (filename, size, content type)
- **users**: Stores user authentication data with bcrypt password hashing
- **Initialization**: Automatic schema creation via `docker/init-db.sql`

## Project-Specific Patterns

### Language/Translation System
- Translations stored in `frontend/src/contexts/LanguageContext.tsx`
- Access via `useLanguage()` hook and `t(key)` function
- Language preference persisted in localStorage
- Document language attribute updated automatically

### Theme System
- Dark/light mode toggle in header
- SCSS variables for theme management
- System preference detection on initial load
- Theme state managed at component level (not global context)

### API Centralization
- All API calls go through `frontend/src/services/api.ts`
- URL building via `frontend/src/config/api.ts`
- Consistent error handling and token management
- Type safety with TypeScript interfaces

### Data Structure
- JSON files in `frontend/src/data/` for static content (skills, education, etc.)
- TypeScript interfaces in `frontend/src/types/` for type definitions
- Centralized constants in `frontend/src/constants/`

## Admin Credentials
- Default username: `admin`
- Default password: `cvadmin2024`
- Configurable via environment variables: `ADMIN_USERNAME`, `ADMIN_PASSWORD`

## Database Tables
- **cv_files**: CV files stored as BYTEA with metadata
- **users**: User authentication data with bcrypt hashing
- **Built frontend**: `frontend/build/` (static files only)