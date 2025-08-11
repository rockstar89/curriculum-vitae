package storage

import (
	"database/sql"
	"fmt"
	"io"
	"time"
)

// CVFile represents a CV file stored in the database
type CVFile struct {
	ID           int       `json:"id"`
	FileName     string    `json:"fileName"`
	OriginalName string    `json:"originalName"`
	FileSize     int64     `json:"fileSize"`
	ContentType  string    `json:"contentType"`
	IsCurrent    bool      `json:"isCurrent"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// CVStorage handles CV file operations in the database
type CVStorage struct {
	db *sql.DB
}

// NewCVStorage creates a new CVStorage instance
func NewCVStorage() *CVStorage {
	return &CVStorage{
		db: GetDB(),
	}
}

// UploadCV saves a CV file to database (replaces existing CV)
func (cs *CVStorage) UploadCV(file io.Reader, originalName string, fileSize int64, contentType string) (*CVFile, error) {
	// Read file data
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}

	// Start transaction
	tx, err := cs.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Mark all existing files as not current
	_, err = tx.Exec("UPDATE cv_files SET is_current = false, updated_at = CURRENT_TIMESTAMP")
	if err != nil {
		return nil, fmt.Errorf("failed to update existing files: %w", err)
	}

	// Insert new CV file
	var cvFile CVFile
	err = tx.QueryRow(`
		INSERT INTO cv_files (filename, original_name, content_type, file_size, file_data, is_current, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, filename, original_name, content_type, file_size, is_current, created_at, updated_at
	`, "current_cv.pdf", originalName, contentType, fileSize, fileData).Scan(
		&cvFile.ID, &cvFile.FileName, &cvFile.OriginalName, 
		&cvFile.ContentType, &cvFile.FileSize, &cvFile.IsCurrent,
		&cvFile.CreatedAt, &cvFile.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert CV file: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &cvFile, nil
}

// GetCurrentCV returns the current CV file metadata
func (cs *CVStorage) GetCurrentCV() (*CVFile, error) {
	var cvFile CVFile
	err := cs.db.QueryRow(`
		SELECT id, filename, original_name, content_type, file_size, is_current, created_at, updated_at
		FROM cv_files 
		WHERE is_current = true 
		ORDER BY created_at DESC 
		LIMIT 1
	`).Scan(
		&cvFile.ID, &cvFile.FileName, &cvFile.OriginalName,
		&cvFile.ContentType, &cvFile.FileSize, &cvFile.IsCurrent,
		&cvFile.CreatedAt, &cvFile.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No CV found
		}
		return nil, fmt.Errorf("failed to get current CV: %w", err)
	}

	return &cvFile, nil
}

// GetCurrentCVWithData returns the current CV with file data
func (cs *CVStorage) GetCurrentCVWithData() (*CVFile, []byte, error) {
	var cvFile CVFile
	var fileData []byte
	
	err := cs.db.QueryRow(`
		SELECT id, filename, original_name, content_type, file_size, is_current, created_at, updated_at, file_data
		FROM cv_files 
		WHERE is_current = true 
		ORDER BY created_at DESC 
		LIMIT 1
	`).Scan(
		&cvFile.ID, &cvFile.FileName, &cvFile.OriginalName,
		&cvFile.ContentType, &cvFile.FileSize, &cvFile.IsCurrent,
		&cvFile.CreatedAt, &cvFile.UpdatedAt, &fileData,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil // No CV found
		}
		return nil, nil, fmt.Errorf("failed to get current CV with data: %w", err)
	}

	return &cvFile, fileData, nil
}

// DeleteCV deletes the current CV file
func (cs *CVStorage) DeleteCV() error {
	result, err := cs.db.Exec("DELETE FROM cv_files WHERE is_current = true")
	if err != nil {
		return fmt.Errorf("failed to delete CV: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no CV found to delete")
	}

	return nil
}

// GetStats returns statistics about CV files
func (cs *CVStorage) GetStats() (int, int64, error) {
	var fileCount int
	var totalSize int64

	err := cs.db.QueryRow(`
		SELECT COUNT(*), COALESCE(SUM(file_size), 0) 
		FROM cv_files 
		WHERE is_current = true
	`).Scan(&fileCount, &totalSize)

	if err != nil {
		return 0, 0, fmt.Errorf("failed to get stats: %w", err)
	}

	return fileCount, totalSize, nil
}