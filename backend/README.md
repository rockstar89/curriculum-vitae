# CV Backend

Go backend for the CV landing page with authentication and file management.

## Setup

1. **Install Go** (version 1.21 or later)
2. **Install dependencies:**
   ```bash
   cd backend
   go mod tidy
   ```

3. **Configure environment:**
   ```bash
   cp .env.example .env
   # Edit .env with your settings
   ```

4. **Run the server:**
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Public Endpoints
- `GET /health` - Health check
- `POST /api/login` - Admin login
- `GET /api/download-cv` - Download current CV

### Protected Endpoints (require JWT token)
- `GET /api/verify` - Verify token
- `POST /api/upload-cv` - Upload new CV
- `GET /api/cv-info` - Get current CV info

## Authentication

1. Login with admin credentials to get JWT token
2. Include token in requests: `Authorization: Bearer <token>`
3. Token expires after 24 hours

## Environment Variables

- `PORT` - Server port (default: 8080)
- `JWT_SECRET` - JWT signing secret
- `ADMIN_USERNAME` - Admin username
- `ADMIN_PASSWORD` - Admin password
- `CORS_ORIGIN` - Allowed CORS origin
- `UPLOAD_DIR` - File upload directory

## File Upload

- Only PDF files allowed
- Max file size: 10MB
- Files stored with timestamp
- Current CV symlinked for easy access