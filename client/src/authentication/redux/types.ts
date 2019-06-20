export interface userData {
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

export type AuthActionTypes = userSignUpRequestAction | userSignUpSuccessAction | userSignUpFailureAction