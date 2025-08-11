import React from 'react';
import { Experience } from '../types/experience';
import { useLanguage } from '../contexts/LanguageContext';
import experiencesData from '../data/experiences.json';

const ExperienceComponent: React.FC = () => {
  const { t, language } = useLanguage();
  
  const experiences: Experience[] = experiencesData;

  return (
    <section className="experience" id="experience">
      <div className="container">
        <h2 className="section-title">{t('experience.title')}</h2>
        <div className="experience-timeline">
          {experiences.map((exp: Experience) => (
            <div key={exp.id} className="experience-item">
              <div className="experience-dot"></div>
              <div className="experience-content">
                <div className="experience-header">
                  <h3 className="experience-title">{exp.title}</h3>
                  <span className="experience-period">{exp.period[language]}</span>
                </div>
                <div className="experience-company">
                  <strong>{exp.company}</strong> • {exp.industry}
                </div>
                <div className="experience-location">{exp.location[language]}</div>
                <p className="experience-description">{exp.description[language]}</p>
                <div className="experience-tech">
                  {exp.technologies.map((tech: string) => (
                    <span key={tech} className="tech-tag">{tech}</span>
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default ExperienceComponent;