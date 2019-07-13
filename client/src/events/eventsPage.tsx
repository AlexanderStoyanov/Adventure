import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch } from 'redux';
import { AppState } from '../rootReducer';
import Events from './events';

interface Props {

}

class EventsPage extends React.Component<Props> {
    render() {
        return(
            <div className="container">
                <div className="row mt-5">
                    <div className="col-md-4 mx-auto">
                        <Events 
                            
                        />
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state: AppState) => ({
    auth: state.auth,
})

const mapDispatchToProps = (dispatch: Dispatch) => bindActionCreators(
    {

    },
    dispatch
)

export default connect(mapStateToProps, mapDispatchToProps)(EventsPage);