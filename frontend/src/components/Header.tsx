import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import ThemeToggle from './ThemeToggle';
import LanguageSelector from './LanguageSelector';
import { useLanguage } from '../contexts/LanguageContext';

const Header: React.FC = () => {
  const location = useLocation();
  const isHomePage = location.pathname === '/';
  const { t } = useLanguage();

  return (
    <header className="header">
      <div className="container">
        <div className="nav">
          <div className="logo">
            <h1>{t('header.name')}</h1>
            <span className="tagline">{t('header.jobTitle')}</span>
          </div>
          <nav className="nav-menu">
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
            <LanguageSelector />
            <ThemeToggle />
          </nav>
        </div>
      </div>
    </header>
  );
};

export default Header;