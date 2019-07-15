export const EVENTS_LOAD_REQUEST = 'EVENTS_LOAD_REQUEST';
export const EVENTS_LOAD_SUCCESS = 'EVENTS_LOAD_SUCCESS';
export const EVENTS_LOAD_FAILURE = 'EVENTS_LOAD_FAILURE';

interface Event {
    name: string
    peopleJoined: string
}

export interface eventsData {
}

export interface signUpState {
}

interface eventsLoadRequestAction {
    type: typeof EVENTS_LOAD_REQUEST
}

interface eventsLoadSuccessAction {
    type: typeof EVENTS_LOAD_SUCCESS,
    payload: { events: Array<Event> }
}

interface eventsLoadFailureAction {
    type: typeof EVENTS_LOAD_FAILURE,
    payload: { error: string },
}

export type EventsActionTypes = 
eventsLoadRequestAction | 
eventsLoadSuccessAction | 
eventsLoadFailureAction