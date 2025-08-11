# CV Landing Page - Full Stack Application

A modern CV landing page built with React (TypeScript) frontend and Go backend, featuring CV upload management, PostgreSQL database, and Docker containerization.

## 🚀 Features

- **Frontend (React + TypeScript + SCSS)**
  - Modern, responsive CV landing page
  - Professional minimalistic design
  - TypeScript for type safety
  - SCSS with organized architecture (7-1 pattern)
  - Admin authentication system
  - CV upload and management interface

- **Backend (Go)**
  - RESTful API with Gin framework
  - JWT authentication
  - PostgreSQL integration with GORM
  - File upload with validation
  - CV version management
  - Health check endpoints

- **Database (PostgreSQL)**
  - Complete CV file storage (binary data + metadata)
  - Version control for uploaded CVs
  - Automatic migrations with GORM
  - No file system dependencies

- **Deployment**
  - Docker containerization
  - docker-compose for easy local development
  - Production-ready configuration

## 🏗️ Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  React Frontend │────│   Go Backend     │────│   PostgreSQL    │
│  (TypeScript)   │    │  (Gin + GORM)    │    │    Database     │
│  Port: 3000     │    │  Port: 8080      │    │  Port: 5432     │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

## 🛠️ Tech Stack

### Frontend
- **React 18** with TypeScript
- **React Router** for navigation
- **SASS/SCSS** with organized architecture
- **Responsive Design** with mobile-first approach

### Backend
- **Go 1.21** with Gin web framework
- **GORM** for database operations
- **JWT** for authentication
- **PostgreSQL** database
- **File upload** with validation

### Infrastructure
- **Docker** & **Docker Compose**
- **Nginx** for frontend serving
- **pgAdmin** for database management

## 🚦 Quick Start

### Prerequisites
- Docker and Docker Compose
- Git

### 1. Clone Repository
```bash
git clone <repository-url>
cd claude-cv-landing
```

### 2. Environment Setup

Copy environment files and update with your values:

```bash
# Backend environment
cp backend/.env.example backend/.env
# Edit backend/.env with your configuration
```

### 3. Start with Docker Compose

**For full application (production-like):**
```bash
docker-compose up --build
```

**For development (database only):**
```bash
# Start only PostgreSQL and pgAdmin
docker-compose -f docker-compose.dev.yml up -d

# Run frontend in development mode
npm install
npm start

# Run backend in development mode
cd backend
go mod tidy
go run cmd/main.go
```

### 4. Access Applications

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **pgAdmin**: http://localhost:5050 (admin@example.com / admin123)

## 📚 API Documentation

### Authentication
```bash
# Login
POST /api/login
Content-Type: application/json
{
  "username": "admin",
  "password": "your-password"
}

# Verify token
GET /api/verify
Authorization: Bearer <token>
```

### CV Management
```bash
# Upload CV (Protected)
POST /api/upload-cv
Authorization: Bearer <token>
Content-Type: multipart/form-data
# Form field: cv (PDF file)

# Download current CV (Public)
GET /api/download-cv

# Get CV info (Protected)
GET /api/cv-info
Authorization: Bearer <token>

# List all CVs (Protected)
GET /api/cv-list
Authorization: Bearer <token>

# Set current CV (Protected)
PUT /api/cv/{id}/set-current
Authorization: Bearer <token>

# Delete CV (Protected)
DELETE /api/cv/{id}
Authorization: Bearer <token>
```

## 🏗️ Development

### Frontend Development
```bash
# Install dependencies
npm install

# Start development server
npm start

# Build for production
npm run build

# Run tests
npm test
```

### Backend Development
```bash
cd backend

# Install dependencies
go mod tidy

# Run in development mode
go run cmd/main.go

# Run with live reload (using air)
go install github.com/cosmtrek/air@latest
air

# Build for production
go build -o main cmd/main.go
```

### Database Operations
```bash
# Connect to database
docker exec -it cv-postgres psql -U postgres -d cv_backend

# Backup database
docker exec cv-postgres pg_dump -U postgres cv_backend > backup.sql

# Restore database
docker exec -i cv-postgres psql -U postgres cv_backend < backup.sql
```

## 📁 Project Structure

```
├── src/                          # Frontend source
│   ├── components/              # React components
│   ├── hooks/                   # Custom React hooks
│   ├── services/               # API services
│   ├── styles/                 # SCSS styles (7-1 architecture)
│   │   ├── abstracts/          # Variables, mixins
│   │   ├── base/              # Reset, typography
│   │   ├── components/        # Component styles
│   │   ├── layout/           # Header, container styles
│   │   └── pages/            # Page-specific styles
│   ├── types/                # TypeScript type definitions
│   ├── utils/               # Utility functions
│   └── constants/          # Application constants
├── backend/                  # Go backend
│   ├── cmd/                 # Application entry points
│   ├── internal/           # Private application code
│   │   ├── auth/          # JWT authentication
│   │   ├── database/     # Database connection
│   │   ├── handlers/    # HTTP handlers
│   │   ├── middleware/ # HTTP middleware
│   │   ├── models/    # Database models
│   │   └── services/ # Business logic
│   ├── Dockerfile        # Backend Docker configuration
│   └── .env.example     # Environment variables template
├── docker-compose.yml        # Production Docker setup
├── docker-compose.dev.yml   # Development Docker setup
├── Dockerfile              # Frontend Docker configuration
└── nginx.conf             # Nginx configuration
```

## 🔧 Environment Variables

### Backend (.env)
```bash
# Server Configuration
PORT=8080
GIN_MODE=release

# Authentication
JWT_SECRET=your-super-secret-jwt-key
ADMIN_USERNAME=admin
ADMIN_PASSWORD=your-secure-password

# CORS
CORS_ORIGIN=http://localhost:3000

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=cv_backend
DB_SSLMODE=disable
```

## 🚢 Deployment

### Production with Docker
```bash
# Build and start all services
docker-compose up --build -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Health Checks
- Backend: http://localhost:8080/health
- Database: Check using pgAdmin or psql

## 🔒 Security Considerations

1. **Change default passwords** in production
2. **Use strong JWT secrets**
3. **Enable HTTPS** in production
4. **Configure proper CORS** origins
5. **Regular security updates**
6. **File upload validation**
7. **Rate limiting** (consider adding)

## 🐛 Troubleshooting

### Common Issues

**Database connection failed:**
```bash
# Check if PostgreSQL is running
docker-compose ps

# Check database logs
docker-compose logs postgres
```

**Frontend can't connect to backend:**
```bash
# Verify backend is running
curl http://localhost:8080/health

# Check CORS configuration
# Verify API_BASE_URL in frontend
```

**File upload issues:**
```bash
# Check database connection
docker-compose -f docker-compose.curriculum-vitae.dev.yml ps

# Check file size limits (10MB default)
# Verify file type validation (PDF only)
# Check database storage space
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## ✨ Acknowledgments

- Built with modern web technologies
- Follows industry best practices
- Production-ready architecture
- Comprehensive documentation