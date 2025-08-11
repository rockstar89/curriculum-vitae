// Authentication types
export interface LoginCredentials {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  error?: string;
}

// CV types
export interface CVInfo {
  name: string;
  size: number;
  lastModified: string;
}

export interface UploadResponse {
  success: boolean;
  error?: string;
  message?: string;
}

// Component props
export interface LoginProps {
  onLogin: (token: string) => void;
}

export interface AdminProps {
  isAuthenticated: boolean;
  onLogout: () => void;
}

// Form event types
export interface FormChangeEvent {
  target: {
    name: string;
    value: string;
  };
}

export interface FileSelectEvent {
  target: {
    files: FileList | null;
  };
}