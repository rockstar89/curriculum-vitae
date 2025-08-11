import { useState, useEffect, useCallback } from 'react';
import { apiService } from '../services/api';
import { STORAGE_KEYS } from '../constants';

interface UseAuthReturn {
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (token: string) => void;
  logout: () => void;
}

export const useAuth = (): UseAuthReturn => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  const logout = useCallback((): void => {
    localStorage.removeItem(STORAGE_KEYS.ADMIN_TOKEN);
    localStorage.removeItem(STORAGE_KEYS.ADMIN_USERNAME);
    setIsAuthenticated(false);
  }, []);

  const verifyToken = useCallback(async (token: string): Promise<void> => {
    try {
      const isValid = await apiService.verifyToken(token);
      if (isValid) {
        setIsAuthenticated(true);
      } else {
        // Token invalid, clear storage
        logout();
      }
    } catch (error) {
      console.error('Token verification failed:', error);
      logout();
    } finally {
      setIsLoading(false);
    }
  }, [logout]);

  useEffect(() => {
    const token = localStorage.getItem(STORAGE_KEYS.ADMIN_TOKEN);
    if (token) {
      verifyToken(token);
    } else {
      setIsLoading(false);
    }
  }, [verifyToken]);

  const login = (token: string): void => {
    setIsAuthenticated(true);
  };

  return {
    isAuthenticated,
    isLoading,
    login,
    logout,
  };
};