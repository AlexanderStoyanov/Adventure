import React from 'react';
import { Route, Switch } from "react-router-dom";

import Register from './register/register';

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <main>
          <Switch>
            <Route exact path={`/`} component={Register} />
          </Switch>
        </main>
      </div>
    );
  }
}

export default App;
