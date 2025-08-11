# CV Landing Page - Full Stack Application

A modern CV landing page built with React (TypeScript) frontend and Go backend, featuring CV upload management, PostgreSQL database, and Docker containerization.

## 🚀 Quick Start

### Prerequisites
- Docker and Docker Compose
- Git

### 1. Clone Repository
```bash
git clone <repository-url>
cd curriculum-vitae
```

### 2. Start Full Application
```bash
cd docker
docker-compose up --build
```

### 3. Access Applications
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080  
- **pgAdmin**: http://localhost:5050

### 4. Admin Configuration
- Admin credentials are configured via environment variables
- Check `backend/.env.example` for required environment variables
- Copy and configure your own `.env` file with secure credentials

## 📁 Project Structure

```
curriculum-vitae/
├── frontend/                    # React TypeScript frontend
│   ├── src/                    # React components, hooks, services
│   ├── public/                 # Static assets
│   ├── Dockerfile              # Frontend container config
│   └── package.json            # Frontend dependencies
├── backend/                     # Go backend API
│   ├── cmd/                    # Application entry points
│   ├── internal/               # Business logic
│   ├── Dockerfile              # Backend container config
│   └── go.mod                  # Go dependencies
├── docker/                      # Docker orchestration
│   ├── docker-compose.yml      # Production setup
│   ├── docker-compose.dev.yml  # Development setup
│   └── init-db.sql             # Database initialization
├── scripts/                     # Utility scripts
│   └── start.sh                # Development startup script
└── docs/                        # Documentation
    └── detailed-setup.md       # Detailed setup instructions
```

## 🛠️ Development

### Frontend Development
```bash
cd frontend
npm install
npm start                       # http://localhost:3000
```

### Backend Development  
```bash
cd backend
go mod tidy
go run cmd/main.go             # http://localhost:8080
```

### Database Only (Development)
```bash
cd docker
docker-compose -f docker-compose.dev.yml up -d
```

## ✨ Features

- **Dark/Light Theme** with system preference detection
- **CV Upload & Management** with database storage
- **JWT Authentication** for admin access
- **Responsive Design** with professional styling
- **Docker Containerization** for easy deployment
- **PostgreSQL Database** with file storage
- **TypeScript Support** for type safety

## 📚 Documentation

- [Detailed Setup Guide](docs/detailed-setup.md)
- [API Documentation](docs/detailed-setup.md#-api-documentation)
- [Development Guide](docs/detailed-setup.md#-development)

---

Built with ❤️ using React, TypeScript, Go, PostgreSQL, and Docker.