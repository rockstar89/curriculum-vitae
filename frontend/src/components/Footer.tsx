import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useLanguage } from '../contexts/LanguageContext';

const Footer: React.FC = () => {
  const location = useLocation();
  const isHomePage = location.pathname === '/';
  const { t } = useLanguage();

  return (
    <footer className="footer" id="contact">
      <div className="container">
        <div className="footer-content">
          <div className="footer-section">
            <h3>{t('footer.name')}</h3>
            <p>{t('footer.collab')}</p>
            <p>{t('footer.description')}</p>
          </div>
          <div className="footer-section">
            <h4>{t('footer.quickLinks')}</h4>
            <Link to="/">{t('header.home')}</Link>
            {isHomePage ? (
              <>
                <a href="#experience">{t('header.experience')}</a>
                <a href="#skills">{t('header.skills')}</a>
                <a href="#education">{t('header.education')}</a>
                <a href="#contact">{t('header.contact')}</a>
              </>
            ) : (
              <>
                <Link to="/#experience">{t('header.experience')}</Link>
                <Link to="/#skills">{t('header.skills')}</Link>
                <Link to="/#education">{t('header.education')}</Link>
                <Link to="/#contact">{t('header.contact')}</Link>
              </>
            )}
          </div>
          <div className="footer-section">
            <h4>{t('footer.contact')}</h4>
            <div className="contact-info">
              <p>{t('footer.email')}</p>
              <p>{t('footer.address')}</p>
              <div className="social-links">
                <a href="https://linkedin.com/in/nenadmih" target="_blank" rel="noopener noreferrer">LinkedIn</a>
                <a href="https://github.com/nenadmihajlovic" target="_blank" rel="noopener noreferrer">GitHub</a>
              </div>
            </div>
          </div>
        </div>
        <div className="footer-bottom">
          <p>&copy; 2025 {t('footer.name')}. {t('footer.rights')}</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;