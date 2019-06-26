import axios from 'axios';
import { ThunkAction } from 'redux-thunk';
import { AppState } from '../../rootReducer';
import { Action } from 'redux';

import request from '../../common/actions/request';
import receive from '../../common/actions/receive';
import error from '../../common/actions/error';
import { 
    SIGN_UP_REQUEST, 
    SIGN_UP_SUCCESS, 
    SIGN_UP_FAILURE,
    signUpData,
    LOGIN_REQUEST,
    LOGIN_SUCCESS,
    LOGIN_FAILURE,
    loginData,
} from "./types";


//Sign up attempt
export const userSignUpRequest = (signUpData: signUpData): ThunkAction<void, AppState, null, Action<string>> => {
    return async dispatch => {
        dispatch(request(SIGN_UP_REQUEST));
        try {
            let res = await axios.post('//localhost:8080/auth/register', signUpData);
            dispatch(receive(SIGN_UP_SUCCESS, res.data));
        } catch (err) {
            dispatch(error(SIGN_UP_FAILURE, err.message));
        }
    };
}

//Login attempt
export const userLoginRequest = (loginData: loginData): ThunkAction<void, AppState, null, Action<string>> => {
    return async dispatch => {
        dispatch(request(LOGIN_REQUEST));
        try {
            let res = await axios.post('//localhost:8080/auth/login', loginData);
            dispatch(receive(LOGIN_SUCCESS, res.data));
        } catch (err) {
            dispatch(error(LOGIN_FAILURE, err.message));
        }
    };
}