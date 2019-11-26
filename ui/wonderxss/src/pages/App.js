import React from 'react';
import './App.css';
import InjectionsList from '../components/InjectionsList';
import Alert from '../components/Alert';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          WonderXSS
        </p>
      </header>
      <Alert></Alert>
      <div id="container">
        <InjectionsList></InjectionsList>
      </div>
    </div>
  );
}

export default App;
