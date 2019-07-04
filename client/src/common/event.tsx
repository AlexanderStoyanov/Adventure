import React from 'react';

interface Props {
    lat: number
    lng: number
    text: string
}


class Event extends React.Component<Props> {
    static style = {
        width: 50,
        height: 50,
        background: 'red',
    }

    render() {
        return (
            <div className='border border-dark rounded-circle' style={Event.style}>
                <span className='h-100 d-flex justify-content-center align-items-center'>
                    bla bla
                </span>
            </div>
        )
    }
}

export default Event;