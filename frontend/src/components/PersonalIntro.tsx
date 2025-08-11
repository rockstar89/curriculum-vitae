import React from 'react';
import { apiService } from '../services/api';
import { downloadFile, openInNewTab } from '../utils';
import { EXTERNAL_LINKS } from '../constants';
import { useLanguage } from '../contexts/LanguageContext';

const PersonalIntro: React.FC = () => {
  const { t } = useLanguage();
  
  const handleDownloadCV = (): void => {
    downloadFile(apiService.getDownloadCvUrl(), 'Nenad_Mihajlovic_CV.pdf');
  };

  const handleViewProjects = (): void => {
    openInNewTab(EXTERNAL_LINKS.GITHUB);
  };

  return (
    <section className="intro" id="home">
      <div className="container">
        <div className="intro-content">
          <div className="intro-text">
            <h1 className="intro-title">
              {t('intro.greeting')} {t('intro.name')}
            </h1>
            <h2 className="intro-role">{t('intro.title')}</h2>
            <p className="intro-subtitle">

            </p>
            <p className="intro-description">
              {t('intro.description')}
            </p>
            <div className="intro-actions">
              <button 
                className="btn btn-primary"
                onClick={handleDownloadCV}
              >
                {t('intro.downloadCV')}
              </button>
              <button 
                className="btn btn-secondary"
                onClick={handleViewProjects}
              >
                {t('intro.viewProjects')}
              </button>
            </div>
          </div>
          <div className="intro-image">
            <div className="profile-placeholder">
              <span>{t('intro.nameAcronym')}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default PersonalIntro;