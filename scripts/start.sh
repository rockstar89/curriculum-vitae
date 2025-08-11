#!/bin/bash

echo "🚀 Starting Curriculum Vitae Application Stack"
echo "=============================================="

# Start the database and pgAdmin
echo "📊 Starting PostgreSQL database and pgAdmin..."
docker-compose -f ../docker/docker-compose.dev.yml up -d

# Wait for database to be ready
echo "⏳ Waiting for database to be ready..."
sleep 15

# Check database status
echo "🔍 Checking database status..."
docker-compose -f ../docker/docker-compose.dev.yml ps

echo ""
echo "✅ Curriculum Vitae Stack Status:"
echo "================================="
echo "🗄️  PostgreSQL Database: Running on port 5432"
echo "🔧 pgAdmin:               http://localhost:5050"
echo "⚛️  React Frontend:       http://localhost:3000"
echo ""
echo "📝 Credentials:"
echo "==============="
echo "🔐 Admin Login:"
echo "   Username: admin"
echo "   Password: cvadmin2024"
echo ""
echo "🗄️  Database:"
echo "   Host: localhost:5432"
echo "   Database: curriculum_vitae_dev"
echo "   Username: cvadmin"
echo "   Password: cv2024secure"
echo ""
echo "🔧 pgAdmin:"
echo "   Email: admin@curriculum-vitae.local"
echo "   Password: pgadmin2024"
echo ""
echo "📋 Next Steps:"
echo "=============="
echo "1. To start the Go backend (after database is ready), run:"
echo "   cd backend"
echo "   go mod tidy"
echo "   go run cmd/main.go"
echo ""
echo "2. To start the React frontend, run:"
echo "   cd frontend"
echo "   npm start"
echo ""
echo "3. To start everything with Docker:"
echo "   docker-compose -f ../docker/docker-compose.yml up --build"
echo ""
echo "✨ PostgreSQL Storage:"
echo "======================"
echo "• CV files and user data are stored in PostgreSQL"
echo "• No file system dependencies - everything is in the database"
echo "• Simpler backup/restore - just backup the database"
echo "• Atomic transactions ensure data consistency"
echo "• Perfect for cloud deployments (Render, etc.)"
echo ""
echo "🎉 Your Curriculum Vitae application is ready!"