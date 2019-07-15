import {
    eventsData,
    EventsActionTypes,
    eventsState,
    EVENTS_LOAD_REQUEST,
    EVENTS_LOAD_SUCCESS,
    EVENTS_LOAD_FAILURE,
} from "./types";

const initialState: eventsData = {
    loading: false,
};

export default (state = initialState, action: EventsActionTypes): eventsState => {
    switch (action.type) {
        case EVENTS_LOAD_REQUEST:
            return {
                loading: true,
            };
        case EVENTS_LOAD_SUCCESS:
            return {
                ...state,
                loading: false,
                error: null,
                list: action.payload.events,
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