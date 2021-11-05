import React from 'react';
import './App.css';
import { BrowserRouter } from 'react-router-dom';
import { Router } from './router/Router';

function App() {
  return (
    <>
    <BrowserRouter>
    <p>APP</p>
      <Router/>
    </BrowserRouter>
    </>
  );
}

export default App;
