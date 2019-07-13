import React from 'react';
import EventEntry from "./eventEntry";

interface Props {
}


class Event extends React.Component<Props> {


    render() {
        let eventEntries =
            <EventEntry
                eventName="event1"
                peopleJoined="100"
            />;
        return (
            <ul className="list-group">
                {eventEntries}
            </ul>
        )
    }
}

export default Event;