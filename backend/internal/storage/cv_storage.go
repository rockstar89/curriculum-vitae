package storage

import (
	"cv-backend/internal/models"
	"fmt"
	"io"
	"time"

	"gorm.io/gorm"
)

// CVStorage handles CV file operations in the database
type CVStorage struct {
	db *gorm.DB
}

// NewCVStorage creates a new CVStorage instance
func NewCVStorage() *CVStorage {
	return &CVStorage{
		db: GetDB(),
	}
}

// UploadCV saves a CV file to database (replaces existing CV)
func (cs *CVStorage) UploadCV(file io.Reader, originalName string, fileSize int64, contentType string) (*models.CVFile, error) {
	// Read file data
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %w", err)
	}

	// Start transaction
	tx := cs.db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", tx.Error)
	}
	defer tx.Rollback()

	// Mark all existing files as not current
	if err := tx.Model(&models.CVFile{}).Where("is_current = ?", true).Updates(map[string]interface{}{
		"is_current": false,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return nil, fmt.Errorf("failed to update existing files: %w", err)
	}

	// Insert new CV file
	cvFile := models.CVFile{
		FileName:     "current_cv.pdf",
		OriginalName: originalName,
		ContentType:  contentType,
		FileSize:     fileSize,
		FileData:     fileData,
		IsCurrent:    true,
	}

	if err := tx.Create(&cvFile).Error; err != nil {
		return nil, fmt.Errorf("failed to insert CV file: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &cvFile, nil
}

// GetCurrentCV returns the current CV file metadata
func (cs *CVStorage) GetCurrentCV() (*models.CVFile, error) {
	var cvFile models.CVFile
	err := cs.db.Where("is_current = ?", true).
		Order("created_at DESC").
		First(&cvFile).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No CV found
		}
		return nil, fmt.Errorf("failed to get current CV: %w", err)
	}

	return &cvFile, nil
}

// GetCurrentCVWithData returns the current CV with file data
func (cs *CVStorage) GetCurrentCVWithData() (*models.CVFile, []byte, error) {
	var cvFile models.CVFile
	
	err := cs.db.Where("is_current = ?", true).
		Order("created_at DESC").
		First(&cvFile).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, nil // No CV found
		}
		return nil, nil, fmt.Errorf("failed to get current CV with data: %w", err)
	}

	return &cvFile, cvFile.FileData, nil
}

// DeleteCV deletes the current CV file
func (cs *CVStorage) DeleteCV() error {
	result := cs.db.Where("is_current = ?", true).Delete(&models.CVFile{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete CV: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no CV found to delete")
	}

	return nil
}

// GetStats returns statistics about CV files
func (cs *CVStorage) GetStats() (int, int64, error) {
	type stats struct {
		FileCount int64 `gorm:"column:file_count"`
		TotalSize int64 `gorm:"column:total_size"`
	}

	var result stats
	err := cs.db.Model(&models.CVFile{}).
		Select("COUNT(*) as file_count, COALESCE(SUM(file_size), 0) as total_size").
		Where("is_current = ?", true).
		Scan(&result).Error

	if err != nil {
		return 0, 0, fmt.Errorf("failed to get stats: %w", err)
	}

	return int(result.FileCount), result.TotalSize, nil
}