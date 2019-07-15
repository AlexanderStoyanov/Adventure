import React from 'react';
import EventEntry from "./eventEntry";
import { EventEntity } from "./redux/types";

interface Props {
    events: {
        loading: boolean
        error?: string | null
        list?: Array<EventEntity>
    }
}


class Events extends React.Component<Props> {
    render() {
        var eventEntries = null;
        if (this.props.events.list) {
            eventEntries = this.props.events.list.map(event => {
                return <EventEntry
                    key={event.Name}
                    event={event}
                />;
            });
        }
        return (
            <ul className="list-group">
                {eventEntries}
            </ul>
        )
    }
}

export default Events;