import React from 'react';
import GoogleMapReact from 'google-map-react';
import { apiKey } from '../config';
import Event from '../common/event';

interface Props {
}

class Maps extends React.Component<Props> {

    render() {
        return (
            <div style={{ height: '100vh', width: '100%' }}>
                <GoogleMapReact
                    bootstrapURLKeys={{ key: apiKey }}
                    defaultCenter={{
                        lat: 40.71,
                        lng: -74.00
                    }}
                    defaultZoom={15}
                >
                    <Event
                        lat={40.71}
                        lng={-74.00}
                        text="My Marker is hereeeee"
                    />
                </GoogleMapReact>
            </div>
        )
    }
}

export default Maps;