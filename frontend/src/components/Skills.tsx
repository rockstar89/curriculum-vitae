import React from 'react';
import { SkillCategory, Industry, Achievement } from '../types/skills';
import { useLanguage } from '../contexts/LanguageContext';
import skillsData from '../data/skills.json';
import industriesData from '../data/industries.json';
import achievementsData from '../data/achievements.json';

const Skills: React.FC = () => {
  const { t, language } = useLanguage();
  
  const skillCategories = skillsData as SkillCategory[];
  const industries = industriesData as Industry[];
  const achievements = achievementsData as Achievement[];

  return (
    <section className="skills" id="skills">
      <div className="container">
        <h2 className="section-title">{t('skills.title')}</h2>
        
        <div className="skills-content">
          <div className="technical-skills">
            <h3>{t('skills.technical')}</h3>
            <div className="skills-grid">
              {skillCategories.map(category => (
                <div key={category.id} className="skill-category">
                  <h4>{category.category[language]}</h4>
                  <div className="skills-list">
                    {category.skills.map(skill => (
                      <span key={skill} className="skill-tag">{skill}</span>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </div>

          <div className="industry-experience">
            <h3>{t('skills.industry')}</h3>
            <div className="industries-grid">
              {industries.map(industry => (
                <div key={industry.id} className="industry-card">
                  <div className="industry-category">{industry.category[language]}</div>
                  <div className="industry-info">
                    <h4>{industry.name[language]}</h4>
                    <span>{typeof industry.experience === 'string' ? industry.experience : industry.experience[language]}</span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        <div className="certifications">
          <h3>{t('education.certifications')}</h3>
          <div className="achievements-grid">
            {achievements.map(achievement => (
              <div key={achievement.id} className="achievement">
                <div className="achievement-number">{achievement.number}</div>
                <div className="achievement-text">{achievement.text[language]}</div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};

export default Skills;