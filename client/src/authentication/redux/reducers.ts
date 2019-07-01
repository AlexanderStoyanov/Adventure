import {
    signUpState,
    AuthActionTypes,
    SIGN_UP_REQUEST,
    SIGN_UP_SUCCESS,
    SIGN_UP_FAILURE,
    LOGIN_REQUEST,
    LOGIN_SUCCESS,
    LOGIN_FAILURE,

} from './types';

const initialState: signUpState = {
};

export default (state = initialState, action: AuthActionTypes): signUpState => {
    switch (action.type) {
        case SIGN_UP_REQUEST:
            return {
                loading: true,
            };
        case SIGN_UP_SUCCESS:
            return {
                loading: false,
                error: null,
            };
        case SIGN_UP_FAILURE:
            return {
                loading: false,
                error: action.payload.error
            };
        case LOGIN_REQUEST:
            return {
                loading: true,
            };
        case LOGIN_SUCCESS:
            return {
                loading: false,
                error: null,
                user: action.payload,
            };
        case LOGIN_FAILURE:
            return {
                loading: false,
                error: action.payload.error
            };
        default:
            return state;
    }
}