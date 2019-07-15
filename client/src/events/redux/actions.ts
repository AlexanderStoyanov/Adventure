import axios from 'axios';
import { ThunkAction } from 'redux-thunk';
import { AppState } from '../../rootReducer';
import { Action } from 'redux';

import request from '../../common/actions/request';
import receive from '../../common/actions/receive';
import error from '../../common/actions/error';
import {
    EVENTS_LOAD_REQUEST,
    EVENTS_LOAD_SUCCESS,
    EVENTS_LOAD_FAILURE,
    //eventsData,
} from "./types";


//Sign up attempt
export const userSignUpRequest = (): ThunkAction<void, AppState, null, Action<string>> => {
    return async dispatch => {
        dispatch(request(EVENTS_LOAD_REQUEST));
        try {
            let res = await axios.post('//localhost:8080/auth/register');
            dispatch(receive(EVENTS_LOAD_SUCCESS, res.data));
        } catch (err) {
            dispatch(error(EVENTS_LOAD_FAILURE, err.message));
        }
    };
}