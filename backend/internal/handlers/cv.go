package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"cv-backend/internal/storage"

	"github.com/gin-gonic/gin"
)

type CVHandler struct {
	cvStorage *storage.CVStorage
}

func NewCVHandler() *CVHandler {
	return &CVHandler{
		cvStorage: storage.NewCVStorage(),
	}
}

// UploadCV handles CV file upload (protected endpoint)
func (h *CVHandler) UploadCV(c *gin.Context) {
	// Get the uploaded file
	file, header, err := c.Request.FormFile("cv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	// Validate file type
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".pdf") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF files are allowed"})
		return
	}

	// Validate file size (10MB limit)
	if header.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB limit"})
		return
	}

	// Get content type
	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/pdf"
	}

	// Upload CV using file storage
	cvFile, err := h.cvStorage.UploadCV(file, header.Filename, header.Size, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save CV: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "CV uploaded successfully",
		"filename":    cvFile.FileName,
		"originalName": cvFile.OriginalName,
		"size":        cvFile.FileSize,
		"uploadedAt":  cvFile.CreatedAt,
	})
}

// DownloadCV serves the current CV file (public endpoint)
func (h *CVHandler) DownloadCV(c *gin.Context) {
	// Get current CV from file storage with data
	cvFile, data, err := h.cvStorage.GetCurrentCVWithData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current CV"})
		return
	}

	if cvFile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No CV found"})
		return
	}

	// Set headers for file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=Nenad_Mihajlovic_CV.pdf")
	c.Header("Content-Type", cvFile.ContentType)
	c.Header("Content-Length", strconv.FormatInt(cvFile.FileSize, 10))

	// Serve the file data
	c.Data(http.StatusOK, cvFile.ContentType, data)
}


// ViewCV serves the current CV for viewing in browser (public endpoint)
func (h *CVHandler) ViewCV(c *gin.Context) {
	// Get current CV from file storage with data
	cvFile, data, err := h.cvStorage.GetCurrentCVWithData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current CV"})
		return
	}

	if cvFile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No CV found"})
		return
	}

	// Set headers for inline viewing
	c.Header("Content-Type", cvFile.ContentType)
	c.Header("Content-Length", strconv.FormatInt(cvFile.FileSize, 10))
	c.Header("Content-Disposition", "inline; filename=Nenad_Mihajlovic_CV.pdf")

	// Serve the file data for viewing
	c.Data(http.StatusOK, cvFile.ContentType, data)
}

// GetCVInfo returns information about the current CV (protected endpoint)
func (h *CVHandler) GetCVInfo(c *gin.Context) {
	cvFile, err := h.cvStorage.GetCurrentCV()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current CV"})
		return
	}

	if cvFile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No CV found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists":        true,
		"name":         cvFile.OriginalName,
		"size":         cvFile.FileSize,
		"lastModified": cvFile.UpdatedAt.Format("2006-01-02 15:04:05"),
		"uploadedAt":   cvFile.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}



// DeleteCV deletes the current CV file (protected endpoint)
func (h *CVHandler) DeleteCV(c *gin.Context) {
	err := h.cvStorage.DeleteCV()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "CV deleted successfully",
	})
}

// GetStats returns statistics about CV files (protected endpoint)
func (h *CVHandler) GetStats(c *gin.Context) {
	fileCount, totalSize, err := h.cvStorage.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get statistics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fileCount": fileCount,
		"totalSize": totalSize,
		"totalSizeMB": float64(totalSize) / 1024 / 1024,
	})
}