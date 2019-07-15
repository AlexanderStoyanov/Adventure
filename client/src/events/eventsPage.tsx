import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch } from 'redux';
import { AppState } from '../rootReducer';
import Events from './events';
import { loadEventList } from './redux/actions';

interface Props {
    loadEventList: () => void
}


class EventsPage extends React.Component<Props> {

    componentDidMount() {
        this.props.loadEventList()
    }

    render() {
        return(
            <div className="container">
                <div className="row mt-5">
                    <div className="col-md-12 mx-auto">
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
        loadEventList,
    },
    dispatch
)

export default connect(mapStateToProps, mapDispatchToProps)(EventsPage);