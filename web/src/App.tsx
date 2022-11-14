import React from 'react';
import { HelmetProvider } from 'react-helmet-async';
import { Route, Routes } from 'react-router-dom';
import Home from '@/pages/Home';

const App: React.FC = () => (
  <div className="App">
    <HelmetProvider>
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </HelmetProvider>
  </div>
);

export default App;
