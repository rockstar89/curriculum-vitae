package models

import (
	"time"
)

// User represents a user in the database
type User struct {
	ID                 uint      `gorm:"primarykey" json:"id"`
	Username           string    `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash       string    `gorm:"not null" json:"password_hash"`
	FirstLogin         bool      `gorm:"default:true" json:"first_login"`
	LoginCount         int       `gorm:"default:0" json:"login_count"`
	LastPasswordChange time.Time `gorm:"autoCreateTime" json:"last_password_change"`
	LastLoginAt        time.Time `gorm:"autoCreateTime" json:"last_login_at"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CVFile represents a CV file stored in the database
type CVFile struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	FileName     string    `gorm:"column:filename;not null" json:"fileName"`
	OriginalName string    `gorm:"column:original_name;not null" json:"originalName"`
	ContentType  string    `gorm:"column:content_type;not null;default:'application/pdf'" json:"contentType"`
	FileSize     int64     `gorm:"column:file_size;not null" json:"fileSize"`
	FileData     []byte    `gorm:"column:file_data;not null" json:"-"` // Don't include in JSON
	IsCurrent    bool      `gorm:"column:is_current;default:true;index" json:"isCurrent"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime;index" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

// TableName sets the table name for CVFile
func (CVFile) TableName() string {
	return "cv_files"
}