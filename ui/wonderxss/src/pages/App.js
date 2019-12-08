import React from 'react';
import './App.css';
import InjectionsList from '../components/InjectionsList';
import Alert from '../components/Alert';
import { Container } from '@material-ui/core';

function App() {
  return (
    <div className="App">
      <Alert></Alert>
      <Container className="container">
        <InjectionsList></InjectionsList>
      </Container>
    </div>
  );
}

export default App;
