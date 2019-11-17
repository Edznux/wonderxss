import React from 'react';
import './App.css';
import Payloads from './components/Payloads';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          WonderXSS
        </p>
      </header>
      
      <Payloads></Payloads>
    </div>
  );
}

export default App;
