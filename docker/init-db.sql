-- Initialize Curriculum Vitae Database

-- Create database if it doesn't exist (this is handled by POSTGRES_DB env var)
-- The database 'curriculum_vitae' will be created automatically

-- Set timezone
SET timezone = 'UTC';

-- Create extensions if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create CV files table for persistent storage
CREATE TABLE IF NOT EXISTS cv_files (
    id SERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    original_name VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL DEFAULT 'application/pdf',
    file_size BIGINT NOT NULL,
    file_data BYTEA NOT NULL,
    is_current BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster queries
CREATE INDEX IF NOT EXISTS idx_cv_files_current ON cv_files(is_current) WHERE is_current = true;
CREATE INDEX IF NOT EXISTS idx_cv_files_created_at ON cv_files(created_at);

-- Create users table for database-based user storage
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_login BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_password_change TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    login_count INTEGER DEFAULT 0,
    last_login_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster user lookups
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Print confirmation
SELECT 'Curriculum Vitae database initialized successfully!' as status;