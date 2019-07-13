import React from 'react';

interface Props {
    eventName: string
    peopleJoined: string
}

const EventEntry: React.FunctionComponent<Props> = ({ eventName, peopleJoined }) => {
    return (
        <li className="list-group-item d-flex justify-content-between align-items-center">
            {eventName}
            <span className="badge badge-primary badge-pill">{peopleJoined}</span>
        </li>
    );
};

export default EventEntry;