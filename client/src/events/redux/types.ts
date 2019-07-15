export const EVENTS_LOAD_REQUEST = 'EVENTS_LOAD_REQUEST';
export const EVENTS_LOAD_SUCCESS = 'EVENTS_LOAD_SUCCESS';
export const EVENTS_LOAD_FAILURE = 'EVENTS_LOAD_FAILURE';

export interface EventEntity {
    Name: string
    PeopleJoined: string
}

export interface eventsData {
    loading: boolean
}

export interface eventsState {
    loading: boolean
    error?: string | null
    list?: Array<EventEntity>
}

interface eventsLoadRequestAction {
    type: typeof EVENTS_LOAD_REQUEST
}

interface eventsLoadSuccessAction {
    type: typeof EVENTS_LOAD_SUCCESS,
    payload: { events: Array<EventEntity> }
}

interface eventsLoadFailureAction {
    type: typeof EVENTS_LOAD_FAILURE,
    payload: { error: string },
}

export const CREATE_EVENT_REQUEST = 'CREATE_EVENT_REQUEST';
export const CREATE_EVENT_SUCCESS = 'CREATE_EVENT_SUCCESS';
export const CREATE_EVENT_FAILURE = 'CREATE_EVENT_FAILURE';

export type EventsActionTypes = 
eventsLoadRequestAction | 
eventsLoadSuccessAction | 
eventsLoadFailureAction