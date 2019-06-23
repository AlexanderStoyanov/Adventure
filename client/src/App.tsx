import React from 'react';
import { Route, Switch } from "react-router-dom";

import Register from './authentication/register/register';
import Login from './authentication/login/login';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <main>
          <Switch>
            <Route path={`/register`} component={Register} />
            <Route path={`/`} component={Login} />
          </Switch>
        </main>
      </div>
    );
  }
}

export default App;
