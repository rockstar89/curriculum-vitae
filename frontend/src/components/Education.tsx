import React from 'react';
import { useLanguage } from '../contexts/LanguageContext';
import type { Education as EducationType, Language as LanguageType, Hobby as HobbyType } from '../types/education';
import educationData from '../data/education.json';
import languagesData from '../data/languages.json';
import hobbiesData from '../data/hobbies.json';

const Education: React.FC = () => {
  const { t, language } = useLanguage();
  
  const education = educationData as EducationType[];
  const languages = languagesData as LanguageType[];
  const hobbies = hobbiesData as HobbyType[];

  return (
    <section className="education" id="education">
      <div className="container">
        <h2 className="section-title">{t('education.title')}</h2>
        
        <div className="education-content">
          <div className="education-section">
            <h3>{t('education.section')}</h3>
            <div className="education-items">
              {education.map(edu => (
                <div key={edu.id} className={`education-item ${edu.type}`}>
                  <div className="education-header">
                    <h4>{edu.degree[language]}</h4>
                    <span className="education-period">{edu.period}</span>
                  </div>
                  <div className="education-institution">{edu.institution[language]}</div>
                  <div className="education-location">{edu.location[language]}</div>
                  <p className="education-description">{edu.description[language]}</p>
                </div>
              ))}
            </div>
          </div>

          <div className="languages-section">
            <h3>{t('education.languages')}</h3>
            <div className="languages-grid">
              {languages.map(lang => (
                <div key={lang.name[language]} className="language-item">
                  <div className="language-header">
                    <span className="language-name">{lang.name[language]}</span>
                    <span className="language-level">{lang.level[language]}</span>
                  </div>
                  <div className="language-bar">
                    <div 
                      className="language-progress" 
                      style={{width: `${lang.proficiency}%`}}
                    ></div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        <div className="hobbies-section">
          <h3>{t('education.hobbies')}</h3>
          <div className="hobbies-grid">
            {hobbies.map(hobby => (
              <div key={hobby.name[language]} className="hobby-item">
                <div className="hobby-category">{hobby.category[language]}</div>
                <span className="hobby-name">{hobby.name[language]}</span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default Education;