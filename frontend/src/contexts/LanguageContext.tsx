import React, { createContext, useContext, useState, useEffect } from 'react';

export type Language = 'en' | 'sr';

interface LanguageContextType {
  language: Language;
  setLanguage: (lang: Language) => void;
  t: (key: string) => string;
}

const LanguageContext = createContext<LanguageContextType | undefined>(undefined);

export const useLanguage = () => {
  const context = useContext(LanguageContext);
  if (!context) {
    throw new Error('useLanguage must be used within a LanguageProvider');
  }
  return context;
};

interface Translations {
  [key: string]: {
    [key: string]: string;
  };
}

const translations: Translations = {
  en: {
    // Header
    'header.home': 'Home',
    'header.experience': 'Experience',
    'header.skills': 'Skills',
    'header.education': 'Education',
    'header.contact': 'Contact',
    'header.downloadCV': 'Download CV',
    'header.admin': 'Admin',
    'header.name': 'Nenad Mihajlović',
    'header.jobTitle': '.NET Software Engineer',
    
    // Personal Intro
    'intro.greeting': 'Hi, I\'m',
    'intro.name': 'Nenad Mihajlović',
    'intro.title': 'Senior Full Stack Developer',
    'intro.description': 'A seasoned software engineer with 10+ years of experience leading distributed teams and architecting scalable solutions across healthcare, fintech, and enterprise platforms. From building microservices on Azure to mentoring teams of 8+ developers, I blend technical expertise in .NET, Angular, and cloud technologies with a passion for music, sports, and continuous learning. Currently crafting next-generation solutions at PwC while jamming on guitar in my downtime.',
    'intro.downloadCV': 'Download CV',
    'intro.viewProjects': 'View Projects',
    'intro.nameAcronym': 'NM',
    'intro.experience': 'Belgrade, Serbia • 10+ Years Experience',

    // About Section
    'about.title': 'About Me',
    'about.description': 'I am a seasoned Full Stack Developer with extensive experience in building scalable web applications and leading development teams. My expertise spans across modern JavaScript frameworks, cloud technologies, and agile methodologies.',
    
    // Experience
    'experience.title': 'Professional Experience',
    
    // Skills
    'skills.title': 'Skills & Expertise',
    'skills.technical': 'Technical Skills',
    'skills.industry': 'Industry Experience',
    
    // Education
    'education.title': 'Education & Languages',
    'education.section': 'Education',
    'education.languages': 'Languages',
    'education.hobbies': 'Interests & Hobbies',
    'education.certifications': 'Achievements & Highlights',
    
    // Admin
    'admin.title': 'CV Management',
    'admin.description': 'Upload and manage your CV file',
    'admin.changePassword': 'Change Password',
    'admin.logout': 'Logout',
    'admin.uploadNew': 'Upload New CV',
    'admin.selectFile': 'Select a PDF file to upload as your new CV',
    'admin.selectedFile': 'Selected:',
    'admin.uploadCV': 'Upload CV',
    'admin.uploading': 'Uploading...',
    'admin.currentCV': 'Current CV',
    'admin.file': 'File:',
    'admin.size': 'Size:',
    'admin.lastUpdated': 'Last Updated:',
    'admin.downloadCV': 'Download CV',
    'admin.deleteCV': 'Delete CV',
    'admin.deleteConfirm': 'Are you sure you want to delete the current CV? This action cannot be undone.',
    'admin.noCVUploaded': 'No CV uploaded yet',
    'admin.instructions': 'Instructions',
    'admin.instruction1': 'Only PDF files are accepted',
    'admin.instruction2': 'Maximum file size: 10MB',
    'admin.instruction3': 'The uploaded CV will replace the current one',
    'admin.instruction4': 'Visitors can download the CV from the main page',
    'admin.selectPDFError': 'Please select a PDF file',
    'admin.selectFileFirst': 'Please select a file first',
    'admin.uploadSuccess': 'CV uploaded successfully!',
    
    // Password Change
    'password.title': 'Change Password',
    'password.current': 'Current Password:',
    'password.new': 'New Password:',
    'password.confirm': 'Confirm New Password:',
    'password.change': 'Change Password',
    'password.changing': 'Changing...',
    'password.cancel': 'Cancel',
    
    // Login
    'login.title': 'Admin Login',
    'login.description': 'Access the CV management dashboard',
    'login.username': 'Username',
    'login.password': 'Password',
    'login.signIn': 'Sign In',
    'login.signingIn': 'Signing in...',
    'login.info': 'This is a secure admin panel for managing CV content.',
    
    // Footer
    'footer.contact': 'Contact Information',
    'footer.quickLinks': 'Quick Links',
    'footer.professional': 'Professional',
    'footer.rights': 'All rights reserved.',
    'footer.name': 'Nenad Mihajlović',
    'footer.address': 'Belgrade, Serbia',
    'footer.email': 'nenad.89.mihajlovic@gmail.com',
    'footer.collab': 'Available for exciting opportunities and collaborations',
    'footer.description': '.NET Software Engineer and tech enthusiast based in Belgrade, Serbia',
    
    // Common
    'common.loading': 'Loading...',
    'common.error': 'Error',
    'common.success': 'Success',
    'common.mb': 'MB',
  },
  sr: {
    // Header
    'header.home': 'Почетна',
    'header.experience': 'Искуство',
    'header.skills': 'Вештине',
    'header.education': 'Образовање',
    'header.contact': 'Контакт',
    'header.downloadCV': 'Преузми CV',
    'header.admin': 'Админ',
    'header.name': 'Ненад Михајловић',
    'header.jobTitle': '.NET Software Engineer',

    // Personal Intro
    'intro.greeting': 'Здраво, ја сам',
    'intro.name': 'Ненад Михајловић',
    'intro.title': 'Senior Full Stack Developer',
    'intro.description': 'Искусни софтверски инжењер са 10+ година искуства у вођењу дистрибуираних тимова и архитектурисању скалабилних решења кроз здравствене, финансијске и enterprise платформе. Од изградње микросервиса на Azure-у до менторства тимова од 8+ програмера, комбинујем техничку експертизу у .NET, Angular и cloud технологијама са страшћу за музику, спорт и континуирано учење. Тренутно креирам решења нове генерације у PwC-у док свирам гитару у слободно време.',
    'intro.downloadCV': 'Преузми CV',
    'intro.viewProjects': 'Погледај пројекте',
    'intro.nameAcronym': 'НМ',
    'intro.experience': 'Београд, Србија • 10+ година искуства',
    
    // About Section
    'about.title': 'О мени',
    'about.description': 'Ја сам искусни Full Stack програмер са великим искуством у изградњи скалабилних веб апликација и вођењу развојних тимова. Моја експертиза обухвата модерне JavaScript framework-е, cloud технологије и агилне методологије.',
    
    // Experience
    'experience.title': 'Професионално искуство',
    
    // Skills
    'skills.title': 'Вештине и експертиза',
    'skills.technical': 'Техничке вештине',
    'skills.industry': 'Индустријско искуство',
    
    // Education
    'education.title': 'Образовање и језици',
    'education.section': 'Образовање',
    'education.languages': 'Језици',
    'education.hobbies': 'Интересовања и хобији',
    'education.certifications': 'Достигнућа и истакнути резултати',
    
    // Admin
    'admin.title': 'Управљање CV-ом',
    'admin.description': 'Отпремите и управљајте вашим CV фајлом',
    'admin.changePassword': 'Промени лозинку',
    'admin.logout': 'Одјави се',
    'admin.uploadNew': 'Отпреми нови CV',
    'admin.selectFile': 'Изаберите PDF фајл за отпремање као ваш нови CV',
    'admin.selectedFile': 'Изабрано:',
    'admin.uploadCV': 'Отпреми CV',
    'admin.uploading': 'Отпремање...',
    'admin.currentCV': 'Тренутни CV',
    'admin.file': 'Фајл:',
    'admin.size': 'Величина:',
    'admin.lastUpdated': 'Последњи пут ажурирано:',
    'admin.downloadCV': 'Преузми CV',
    'admin.deleteCV': 'Обриши CV',
    'admin.deleteConfirm': 'Да ли сте сигурни да желите да обришете тренутни CV? Ова акција се не може опозвати.',
    'admin.noCVUploaded': 'Још увек није отпремљен CV',
    'admin.instructions': 'Инструкције',
    'admin.instruction1': 'Само PDF фајлови су дозвољени',
    'admin.instruction2': 'Максимална величина фајла: 10MB',
    'admin.instruction3': 'Отпремљени CV ће заменити тренутни',
    'admin.instruction4': 'Посетиоци могу преузети CV са главне стране',
    'admin.selectPDFError': 'Молимо изаберите PDF фајл',
    'admin.selectFileFirst': 'Молимо прво изаберите фајл',
    'admin.uploadSuccess': 'CV успешно отпремљен!',
    
    // Password Change
    'password.title': 'Промени лозинку',
    'password.current': 'Тренутна лозинка:',
    'password.new': 'Нова лозинка:',
    'password.confirm': 'Потврди нову лозинку:',
    'password.change': 'Промени лозинку',
    'password.changing': 'Мењање...',
    'password.cancel': 'Откажи',
    
    // Login
    'login.title': 'Админ пријава',
    'login.description': 'Приступите контролној табли за управљање CV-ом',
    'login.username': 'Корисничко име',
    'login.password': 'Лозинка',
    'login.signIn': 'Пријави се',
    'login.signingIn': 'Пријављивање...',
    'login.info': 'Ово је сигуран админ панел за управљање садржајем CV-а.',
    
    // Footer
    'footer.contact': 'Контакт информације',
    'footer.quickLinks': 'Брзе везе',
    'footer.professional': 'Професионално',
    'footer.rights': 'Сва права задржана.',
    'footer.name': 'Ненад Михајловић',
    'footer.address': 'Београд, Србија',
    'footer.email': 'nenad.89.mihajlovic@gmail.com',
    'footer.collab': 'Доступан за узбудљиве прилике и сарадње',
    'footer.description': '.NET софтверски инжењер и љубитељ технологије из Београда, Србија',

    // Common
    'common.loading': 'Учитавање...',
    'common.error': 'Грешка',
    'common.success': 'Успешно',
    'common.mb': 'MB',
  }
};

interface LanguageProviderProps {
  children: React.ReactNode;
}

export const LanguageProvider: React.FC<LanguageProviderProps> = ({ children }) => {
  const [language, setLanguageState] = useState<Language>(() => {
    const saved = localStorage.getItem('cv-language');
    return (saved as Language) || 'en';
  });

  const setLanguage = (lang: Language) => {
    setLanguageState(lang);
    localStorage.setItem('cv-language', lang);
    document.documentElement.lang = lang;
  };

  const t = (key: string): string => {
    return translations[language]?.[key] || key;
  };

  useEffect(() => {
    document.documentElement.lang = language;
  }, [language]);

  return (
    <LanguageContext.Provider value={{ language, setLanguage, t }}>
      {children}
    </LanguageContext.Provider>
  );
};