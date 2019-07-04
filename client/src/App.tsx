import React from 'react';
import { Route, Switch } from "react-router-dom";

import Register from './authentication/register/register';
import Login from './authentication/login/login';
import Maps from './maps/maps';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <main>
          <Switch>
            <Route path={`/register`} component={Register} />
            <Route path={`/`} component={Maps} />
          </Switch>
        </main>
      </div>
    );
  }
}

export default App;
