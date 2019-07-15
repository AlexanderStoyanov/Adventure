import { combineReducers } from 'redux';
import authentication from './authentication/redux/reducers';
import events from './events/redux/reducers';


export const rootReducer = combineReducers({
    auth: authentication,
    events: events,
});

export type AppState = ReturnType<typeof rootReducer>