import React from 'react';
import './App.css';
import 'rsuite/dist/styles/rsuite-default.css';
import 'rsuite/lib/styles/themes/dark/index.less';
// import Layout from  "./layouts/Layout"
import RouteManager from './auth/RouteManager';


function App() {
  return (
    <div className="app">
      <RouteManager/>
    </div>
  );
}

export default App;
