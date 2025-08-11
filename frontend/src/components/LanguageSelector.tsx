import React from 'react';
import { useLanguage } from '../contexts/LanguageContext';

const LanguageSelector: React.FC = () => {
  const { language, setLanguage } = useLanguage();

  return (
    <div className="language-selector">
      <button
        onClick={() => setLanguage('en')}
        className={`lang-btn ${language === 'en' ? 'active' : ''}`}
        title="English"
      >
        EN
      </button>
      <span className="lang-separator">|</span>
      <button
        onClick={() => setLanguage('sr')}
        className={`lang-btn ${language === 'sr' ? 'active' : ''}`}
        title="Српски"
      >
        СР
      </button>
    </div>
  );
};

export default LanguageSelector;