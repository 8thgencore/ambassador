import React from 'react';
import './App.css';
import Users from './pages/Users';
import Login from './pages/Login';
import Register from './pages/Register';
import { BrowserRouter, Route } from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Route path={'/'} exact  component={Users}></Route>
        <Route path={'/login'} component={Login}></Route>
        <Route path={'/register'} component={Register}></Route>
        <Route path={'/users'} exact  component={Users}></Route>
      </BrowserRouter>
    </div>
  );
}

export default App;
