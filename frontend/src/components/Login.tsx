import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { LoginCredentials, LoginProps } from '../types';
import { apiService } from '../services/api';
import { STORAGE_KEYS, ROUTES } from '../constants';

const Login: React.FC<LoginProps> = ({ onLogin }) => {
  const [credentials, setCredentials] = useState<LoginCredentials>({ username: '', password: '' });
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      const data = await apiService.login(credentials);
      
      localStorage.setItem(STORAGE_KEYS.ADMIN_TOKEN, data.token);
      localStorage.setItem(STORAGE_KEYS.ADMIN_USERNAME, credentials.username);
      onLogin(data.token);
      navigate(ROUTES.ADMIN);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Connection failed. Make sure the backend server is running.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const { name, value } = e.target;
    setCredentials(prevCredentials => ({
      ...prevCredentials,
      [name]: value,
    }));
  };

  return (
    <div className="login-container">
      <div className="login-content">
        <h1 className="login-title">Admin Login</h1>
        <p className="login-description">Access the CV management system</p>
        
        <form onSubmit={handleSubmit} className="login-form">
          <div className="form-group">
            <label htmlFor="username">Username</label>
            <input
              type="text"
              id="username"
              name="username"
              value={credentials.username}
              onChange={handleChange}
              required
              className="form-input"
              disabled={isLoading}
            />
          </div>
          
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              value={credentials.password}
              onChange={handleChange}
              required
              className="form-input"
              disabled={isLoading}
            />
          </div>
          
          {error && (
            <div className="error-message">
              {error}
            </div>
          )}
          
          <button
            type="submit"
            className={`btn ${isLoading ? 'btn-disabled' : 'btn-primary'}`}
            disabled={isLoading}
          >
            {isLoading ? 'Logging in...' : 'Login'}
          </button>
        </form>
        
        <div className="login-info">
          <p>Secure access for CV management and uploads</p>
        </div>
      </div>
    </div>
  );
};

export default Login;