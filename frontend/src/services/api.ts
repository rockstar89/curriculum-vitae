import { LoginCredentials, LoginResponse, CVInfo, UploadResponse } from '../types';
import { API_CONFIG, buildApiUrl } from '../config/api';

class ApiService {
  private getAuthHeaders(): HeadersInit {
    const token = localStorage.getItem('cvAdminToken');
    return {
      'Authorization': token ? `Bearer ${token}` : '',
    };
  }

  async login(credentials: LoginCredentials): Promise<LoginResponse> {
    const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.LOGIN), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    const data = await response.json();
    
    if (!response.ok) {
      throw new Error(data.error || 'Login failed');
    }

    return data;
  }

  async verifyToken(token: string): Promise<boolean> {
    try {
      const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.VERIFY), {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      return response.ok;
    } catch (error) {
      console.error('Token verification failed:', error);
      return false;
    }
  }

  async getCvInfo(): Promise<CVInfo | null> {
    try {
      const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.CV_INFO), {
        headers: this.getAuthHeaders(),
      });

      if (response.ok) {
        return await response.json();
      }
      return null;
    } catch (error) {
      console.error('Failed to fetch CV info:', error);
      return null;
    }
  }

  async uploadCv(file: File): Promise<UploadResponse> {
    const formData = new FormData();
    formData.append('cv', file);

    const token = localStorage.getItem('cvAdminToken');
    const headers: HeadersInit = {};
    
    // Only set Authorization header, let browser set Content-Type for FormData
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }

    const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.UPLOAD_CV), {
      method: 'POST',
      headers,
      body: formData,
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || 'Upload failed');
    }

    return data;
  }

  getDownloadCvUrl(): string {
    return buildApiUrl(API_CONFIG.ENDPOINTS.DOWNLOAD_CV);
  }
}

export const apiService = new ApiService();
export default apiService;