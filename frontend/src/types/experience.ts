export interface Experience {
  id: number;
  title: string;
  company: string;
  industry: string;
  period: {
    en: string;
    sr: string;
  };
  location: {
    en: string;
    sr: string;
  };
  description: {
    en: string;
    sr: string;
  };
  technologies: string[];
}