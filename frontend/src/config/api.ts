export const API_CONFIG = {
  BASE_URL: process.env.REACT_APP_API_URL || 'http://localhost:8080',
  ENDPOINTS: {
    LOGIN: '/api/login',
    VERIFY: '/api/verify',
    CV_INFO: '/api/cv-info',
    UPLOAD_CV: '/api/upload-cv',
    DOWNLOAD_CV: '/api/download-cv',
    VIEW_CV: '/api/view-cv',
    DELETE_CV: '/api/cv',
    CHANGE_PASSWORD: '/api/change-password',
    STATS: '/api/cv-stats',
  }
};

// Helper function to build full URLs
export const buildApiUrl = (endpoint: string): string => {
  return `${API_CONFIG.BASE_URL}${endpoint}`;
};