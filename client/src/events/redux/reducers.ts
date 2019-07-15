import {
    eventsData,
    EventsActionTypes,
    signUpState,
    EVENTS_LOAD_REQUEST,
    EVENTS_LOAD_SUCCESS,
    EVENTS_LOAD_FAILURE,
} from "./types";

const initialState: eventsData = {
};

export default (state = initialState, action: EventsActionTypes): signUpState => {
    switch (action.type) {
        case EVENTS_LOAD_REQUEST:
            return {
                loading: true,
            };
        case EVENTS_LOAD_SUCCESS:
            return {
                loading: false,
                error: null,
                events: action.payload,
            };
        case EVENTS_LOAD_FAILURE:
            return {
                loading: false,
                error: action.payload.error
            };
        default:
            return state;
    }
}