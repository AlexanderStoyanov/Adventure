import React from 'react';
import { EventEntity } from "./redux/types";

interface Props {
    event: EventEntity
}

const EventEntry: React.FunctionComponent<Props> = ({ event }) => {
    return (
        <li className="list-group-item d-flex justify-content-between align-items-center">
            {event.Name}
            <span className="badge badge-primary badge-pill">{event.PeopleJoined}</span>
        </li>
    );
};

export default EventEntry;