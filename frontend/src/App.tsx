import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './styles/main.scss';
import Header from './components/Header';
import PersonalIntro from './components/PersonalIntro';
import Experience from './components/Experience';
import Skills from './components/Skills';
import Education from './components/Education';
import Footer from './components/Footer';
import Admin from './components/Admin';
import Login from './components/Login';
import { useAuth } from './hooks/useAuth';
import { ROUTES } from './constants';
import { LanguageProvider } from './contexts/LanguageContext';

const HomePage: React.FC = () => {
  return (
    <>
      <PersonalIntro />
      <Experience />
      <Skills />
      <Education />
      <Footer />
    </>
  );
};

const App: React.FC = () => {
  const { isAuthenticated, isLoading, login, logout } = useAuth();

  if (isLoading) {
    return (
      <div className="App">
        <div style={{ padding: '2rem', textAlign: 'center' }}>
          Loading...
        </div>
      </div>
    );
  }

  return (
    <LanguageProvider>
      <Router>
        <div className="App">
          <Header />
          <Routes>
            <Route path={ROUTES.HOME} element={<HomePage />} />
            <Route path={ROUTES.LOGIN} element={<Login onLogin={login} />} />
            <Route path={ROUTES.ADMIN} element={<Admin isAuthenticated={isAuthenticated} onLogout={logout} />} />
          </Routes>
        </div>
      </Router>
    </LanguageProvider>
  );
};

export default App;