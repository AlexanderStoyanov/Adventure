import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators, Dispatch } from 'redux';
import { AppState } from '../rootReducer';
import Events from './events';
import { loadEventList, createEvent } from './redux/actions';
import { EventEntity } from "./redux/types";
import TextField from '../common/textField';
import ErrorEntity from '../common/model/errorEntity';

interface Props {
    events: {
        loading: boolean
        error?: string | null
        list?: Array<EventEntity>
    }
    loadEventList: () => void
    createEvent: (eventname: string) => void
}

interface State {
    event: EventEntity
    eventname: string
    errors: ErrorEntity
    timeout: boolean
}


class EventsPage extends React.Component<Props, State> {
    constructor(props: Readonly<Props>) {
        super(props);
        this.state = {
            event: {
                Name: '',
                PeopleJoined: '0',
            },
            eventname: '',
            errors: {},
            timeout: false,
        };

        this.onFieldChange = this.onFieldChange.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }

    async componentDidMount() {
        await this.props.loadEventList()
    }

    private onFieldChange(fieldName: string, value: string) {
        this.setState({ ...this.state, [fieldName]: value });
    }

    private onSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        const { eventname } = this.state;
        this.props.createEvent(eventname);
    }

    render() {
        const { ...props } = this.props;
        return (
            <div className="container">
                <div className="row mt-5">
                    <div className="col-md-12 mx-auto">
                        <form className="mb-5" onSubmit={this.onSubmit} >

                            <TextField
                                error={this.state.errors.eventname}
                                label="Event name"
                                onChange={this.onFieldChange}
                                value={this.state.eventname}
                                field="eventname"
                                type="text"
                            />

                            <button disabled={this.state.timeout} type="submit" className="btn btn-primary">Submit</button>
                        </form>
                        <Events
                            {...props}
                        />
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state: AppState) => ({
    events: state.events,
})

const mapDispatchToProps = (dispatch: Dispatch) => bindActionCreators(
    {
        loadEventList,
        createEvent,
    },
    dispatch
)

export default connect(mapStateToProps, mapDispatchToProps)(EventsPage);