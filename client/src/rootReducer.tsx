import { combineReducers } from 'redux';
import authentication from './authentication/redux/reducers';


export const rootReducer = combineReducers({
    auth: authentication,
});

export type AppState = ReturnType<typeof rootReducer>