export interface signUpData {
    username: string
    email: string
    password: String
}

export interface signUpState {
}

export const SIGN_UP_REQUEST = 'SIGN_UP_REQUEST';
export const SIGN_UP_SUCCESS = 'SIGN_UP_SUCCESS';
export const SIGN_UP_FAILURE = 'SIGN_UP_FAILURE';

interface userSignUpRequestAction {
    type: typeof SIGN_UP_REQUEST
}

interface userSignUpSuccessAction {
    type: typeof SIGN_UP_SUCCESS,
}

interface userSignUpFailureAction {
    type: typeof SIGN_UP_FAILURE,
    payload: { error: string },
}


export interface loginData {
    username: string
    password: String
}

export interface loginState {
}

export const LOGIN_REQUEST = 'LOGIN_REQUEST';
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS';
export const LOGIN_FAILURE = 'LOGIN_FAILURE';

interface userLoginRequestAction {
    type: typeof LOGIN_REQUEST
}

interface userLoginSuccessAction {
    type: typeof LOGIN_SUCCESS,
}

interface userLoginFailureAction {
    type: typeof LOGIN_FAILURE,
    payload: { error: string },
}

export type AuthActionTypes = 
userSignUpRequestAction | 
userSignUpSuccessAction | 
userSignUpFailureAction |
userLoginRequestAction |
userLoginSuccessAction |
userLoginFailureAction