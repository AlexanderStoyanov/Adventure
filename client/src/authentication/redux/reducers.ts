import { 
    signUpState,
    AuthActionTypes,
    SIGN_UP_REQUEST, 
    SIGN_UP_SUCCESS, 
    SIGN_UP_FAILURE,

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
        default:
            return state;
    }
}