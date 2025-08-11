export interface Education {
  id: number;
  degree: {
    en: string;
    sr: string;
  };
  institution: {
    en: string;
    sr: string;
  };
  period: string;
  location: {
    en: string;
    sr: string;
  };
  description: {
    en: string;
    sr: string;
  };
  type: 'formal' | 'professional';
}

export interface Language {
  name: {
    en: string;
    sr: string;
  };
  level: {
    en: string;
    sr: string;
  };
  proficiency: number;
}

export interface Hobby {
  name: {
    en: string;
    sr: string;
  };
  category: {
    en: string;
    sr: string;
  };
}