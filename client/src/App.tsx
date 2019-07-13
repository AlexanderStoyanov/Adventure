import React from 'react';
import { Route, Switch } from "react-router-dom";

import Register from './authentication/register/register';
import Login from './authentication/login/login';
import EventsPage from './events/eventsPage';

//import Maps from './maps/maps';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <main>
          <Switch>
            <Route path={`/login`} component={Login} />
            <Route path={`/register`} component={Register} />
            <Route path={`/`} component={EventsPage} />
          </Switch>
        </main>
      </div>
    );
  }
}

export default App;
