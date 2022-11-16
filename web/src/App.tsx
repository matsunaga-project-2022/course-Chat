import React from 'react';
import { HelmetProvider } from 'react-helmet-async';
import { Route, Routes } from 'react-router-dom';
import Home from '@/pages/Home';
import Layout from '@/components/layout';
import { createTheme, Theme, ThemeProvider } from '@mui/material';

const App = () => {
  const theme: Theme = createTheme({});
  return (
    <div className="App">
      <HelmetProvider>
        <ThemeProvider theme={theme}>
          <Routes>
            <Route path="" element={<Layout />}>
              <Route path="/" element={<Home />} />
            </Route>
          </Routes>
        </ThemeProvider>
      </HelmetProvider>
    </div>
  );
};

export default App;
