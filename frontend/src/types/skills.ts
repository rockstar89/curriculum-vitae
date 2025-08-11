export interface SkillCategory {
  id: number;
  category: {
    en: string;
    sr: string;
  };
  skills: string[];
}

export interface Industry {
  id: number;
  name: {
    en: string;
    sr: string;
  };
  experience: string | {
    en: string;
    sr: string;
  };
  category: {
    en: string;
    sr: string;
  };
}

export interface Achievement {
  id: number;
  number: string;
  text: {
    en: string;
    sr: string;
  };
}