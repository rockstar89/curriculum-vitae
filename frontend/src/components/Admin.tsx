import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AdminProps, CVInfo, UploadResponse } from '../types';
import { useLanguage } from '../contexts/LanguageContext';
import { buildApiUrl, API_CONFIG } from '../config/api';

const Admin: React.FC<AdminProps> = ({ isAuthenticated, onLogout }) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploadStatus, setUploadStatus] = useState<string>('');
  const [isUploading, setIsUploading] = useState<boolean>(false);
  const [cvInfo, setCvInfo] = useState<CVInfo | null>(null);
  
  // Password change state
  const [showPasswordChange, setShowPasswordChange] = useState<boolean>(false);
  const [passwordForm, setPasswordForm] = useState({
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  });
  const [passwordStatus, setPasswordStatus] = useState<string>('');
  const [isChangingPassword, setIsChangingPassword] = useState<boolean>(false);
  
  const navigate = useNavigate();
  const { t } = useLanguage();

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>): void => {
    const file = event.target.files?.[0];
    if (file && file.type === 'application/pdf') {
      setSelectedFile(file);
      setUploadStatus('');
    } else {
      setUploadStatus(t('admin.selectPDFError') || 'Please select a PDF file');
      setSelectedFile(null);
    }
  };

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }
    fetchCVInfo();
  }, [isAuthenticated, navigate]);

  const fetchCVInfo = async (): Promise<void> => {
    try {
      const token = localStorage.getItem('cvAdminToken');
      if (!token) return;

      const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.CV_INFO), {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const data: CVInfo = await response.json();
        setCvInfo(data);
      }
    } catch (error) {
      console.error('Failed to fetch CV info:', error);
    }
  };

  const handleUpload = async (): Promise<void> => {
    if (!selectedFile) {
      setUploadStatus(t('admin.selectFileFirst') || 'Please select a file first');
      return;
    }

    setIsUploading(true);
    setUploadStatus(t('admin.uploading'));

    try {
      const token = localStorage.getItem('cvAdminToken');
      if (!token) {
        setUploadStatus('No authentication token found');
        return;
      }

      const formData = new FormData();
      formData.append('cv', selectedFile);

      const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.UPLOAD_CV), {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
        body: formData,
      });

      const data: UploadResponse = await response.json();

      if (response.ok) {
        setUploadStatus(t('admin.uploadSuccess') || 'CV uploaded successfully!');
        setSelectedFile(null);
        
        const fileInput = document.getElementById('cv-upload') as HTMLInputElement;
        if (fileInput) fileInput.value = '';
        
        fetchCVInfo();
      } else {
        setUploadStatus(data.error || 'Upload failed');
      }
    } catch (error) {
      console.error('Upload error:', error);
      setUploadStatus(`Upload failed: ${error instanceof Error ? error.message : 'Please check your connection'}`);
    } finally {
      setIsUploading(false);
    }
  };

  const handleDownloadCurrent = (): void => {
    window.open(buildApiUrl(API_CONFIG.ENDPOINTS.DOWNLOAD_CV), '_blank');
  };

  const handleDeleteCV = async (): Promise<void> => {
    if (!window.confirm(t('admin.deleteConfirm'))) {
      return;
    }

    try {
      const token = localStorage.getItem('cvAdminToken');
      if (!token) {
        setUploadStatus('No authentication token found');
        return;
      }

      const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.DELETE_CV), {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      const data = await response.json();

      if (response.ok) {
        setUploadStatus('CV deleted successfully!');
        setCvInfo(null);
      } else {
        setUploadStatus(data.error || 'Failed to delete CV');
      }
    } catch (error) {
      setUploadStatus('Failed to delete CV. Please check your connection.');
    }
  };

  const handleLogout = (): void => {
    localStorage.removeItem('cvAdminToken');
    localStorage.removeItem('cvAdminUsername');
    onLogout();
    navigate('/login');
  };

  const handlePasswordChange = async (): Promise<void> => {
    if (passwordForm.newPassword !== passwordForm.confirmPassword) {
      setPasswordStatus('New passwords do not match');
      return;
    }

    if (passwordForm.newPassword.length < 8) {
      setPasswordStatus('New password must be at least 8 characters long');
      return;
    }

    setIsChangingPassword(true);
    setPasswordStatus('Changing password...');

    try {
      const token = localStorage.getItem('cvAdminToken');
      if (!token) {
        setPasswordStatus('No authentication token found');
        return;
      }

      const response = await fetch(buildApiUrl(API_CONFIG.ENDPOINTS.CHANGE_PASSWORD), {
        method: 'PUT',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          current_password: passwordForm.currentPassword,
          new_password: passwordForm.newPassword,
        }),
      });

      const data = await response.json();

      if (response.ok) {
        setPasswordStatus('Password changed successfully!');
        setPasswordForm({
          currentPassword: '',
          newPassword: '',
          confirmPassword: ''
        });
        setTimeout(() => {
          setShowPasswordChange(false);
          setPasswordStatus('');
        }, 2000);
      } else {
        setPasswordStatus(data.error || 'Password change failed');
      }
    } catch (error) {
      setPasswordStatus('Password change failed. Please check your connection.');
    } finally {
      setIsChangingPassword(false);
    }
  };

  const handlePasswordFormChange = (field: string, value: string): void => {
    setPasswordForm(prev => ({
      ...prev,
      [field]: value
    }));
    setPasswordStatus('');
  };

  return (
    <div className="admin-container">
      <div className="admin-content">
        <div className="admin-header">
          <div>
            <h1 className="admin-title">{t('admin.title')}</h1>
            <p className="admin-description">{t('admin.description')}</p>
          </div>
          <div className="admin-actions">
            <button 
              onClick={() => setShowPasswordChange(!showPasswordChange)} 
              className="btn btn-outline"
            >
              {t('admin.changePassword')}
            </button>
            <button onClick={handleLogout} className="btn btn-secondary">
              {t('admin.logout')}
            </button>
          </div>
        </div>
        
        <div className="upload-section">
          <div className="upload-area">
            <div className="upload-icon">📄</div>
            <h3>{t('admin.uploadNew')}</h3>
            <p>{t('admin.selectFile')}</p>
            
            <input
              id="cv-upload"
              type="file"
              accept=".pdf"
              onChange={handleFileSelect}
              className="file-input"
            />
            
            {selectedFile && (
              <div className="selected-file">
                <span>{t('admin.selectedFile')} {selectedFile.name}</span>
                <span className="file-size">({(selectedFile.size / 1024 / 1024).toFixed(2)} {t('common.mb')})</span>
              </div>
            )}
            
            <button 
              onClick={handleUpload}
              disabled={!selectedFile || isUploading}
              className={`btn ${!selectedFile || isUploading ? 'btn-disabled' : 'btn-primary'}`}
            >
              {isUploading ? t('admin.uploading') : t('admin.uploadCV')}
            </button>
            
            {uploadStatus && (
              <div className={`upload-status ${uploadStatus.includes('success') ? 'success' : ''}`}>
                {uploadStatus}
              </div>
            )}
          </div>
          
          <div className="current-cv-section">
            <h3>{t('admin.currentCV')}</h3>
            {cvInfo ? (
              <div className="cv-info">
                <p><strong>{t('admin.file')}</strong> {cvInfo.name}</p>
                <p><strong>{t('admin.size')}</strong> {(cvInfo.size / 1024 / 1024).toFixed(2)} {t('common.mb')}</p>
                <p><strong>{t('admin.lastUpdated')}</strong> {cvInfo.lastModified}</p>
                <div style={{ display: 'flex', gap: '0.75rem', justifyContent: 'center', marginTop: '1rem' }}>
                  <button onClick={handleDownloadCurrent} className="btn btn-secondary">
                    {t('admin.downloadCV')}
                  </button>
                  <button onClick={handleDeleteCV} className="btn btn-outline" style={{ color: '#dc2626', borderColor: '#dc2626' }}>
                    {t('admin.deleteCV')}
                  </button>
                </div>
              </div>
            ) : (
              <p>{t('admin.noCVUploaded')}</p>
            )}
          </div>
        </div>

        {showPasswordChange && (
          <div className="password-change-section">
            <h3>{t('password.title')}</h3>
            <div className="password-form">
              <div className="form-group">
                <label htmlFor="current-password">{t('password.current')}</label>
                <input
                  id="current-password"
                  type="password"
                  value={passwordForm.currentPassword}
                  onChange={(e) => handlePasswordFormChange('currentPassword', e.target.value)}
                  className="form-input"
                />
              </div>
              
              <div className="form-group">
                <label htmlFor="new-password">{t('password.new')}</label>
                <input
                  id="new-password"
                  type="password"
                  value={passwordForm.newPassword}
                  onChange={(e) => handlePasswordFormChange('newPassword', e.target.value)}
                  className="form-input"
                />
              </div>
              
              <div className="form-group">
                <label htmlFor="confirm-password">{t('password.confirm')}</label>
                <input
                  id="confirm-password"
                  type="password"
                  value={passwordForm.confirmPassword}
                  onChange={(e) => handlePasswordFormChange('confirmPassword', e.target.value)}
                  className="form-input"
                />
              </div>
              
              <div className="password-actions">
                <button
                  onClick={handlePasswordChange}
                  disabled={!passwordForm.currentPassword || !passwordForm.newPassword || !passwordForm.confirmPassword || isChangingPassword}
                  className={`btn ${!passwordForm.currentPassword || !passwordForm.newPassword || !passwordForm.confirmPassword || isChangingPassword ? 'btn-disabled' : 'btn-primary'}`}
                >
                  {isChangingPassword ? t('password.changing') : t('password.change')}
                </button>
                <button
                  onClick={() => {
                    setShowPasswordChange(false);
                    setPasswordForm({ currentPassword: '', newPassword: '', confirmPassword: '' });
                    setPasswordStatus('');
                  }}
                  className="btn btn-secondary"
                >
                  {t('password.cancel')}
                </button>
              </div>
              
              {passwordStatus && (
                <div className={`password-status ${passwordStatus.includes('success') ? 'success' : ''}`}>
                  {passwordStatus}
                </div>
              )}
            </div>
          </div>
        )}
        
        <div className="admin-instructions">
          <h3>{t('admin.instructions')}</h3>
          <ul>
            <li>{t('admin.instruction1')}</li>
            <li>{t('admin.instruction2')}</li>
            <li>{t('admin.instruction3')}</li>
            <li>{t('admin.instruction4')}</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Admin;