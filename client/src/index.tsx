import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import { createStore, applyMiddleware, compose } from 'redux';
import { rootReducer } from './rootReducer';
import firebase from 'firebase';

import App from './App';
import * as serviceWorker from './serviceWorker';
import { firebaseConfig } from './config';

firebase.initializeApp(firebaseConfig);

const store = createStore(
    rootReducer,
    compose(
        applyMiddleware(thunk),
        (window as any).__REDUX_DEVTOOLS_EXTENSION__ ? (window as any).__REDUX_DEVTOOLS_EXTENSION__() : (f: any) => f
    ) 
);

ReactDOM.render((
    <Provider store={store}>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </Provider>), document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
