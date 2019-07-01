import React from 'react';
import TextField from '../../common/textField';
import ErrorEntity from '../../common/model/errorEntity';
import { authState } from '../redux/types';
import { withRouter } from "react-router-dom";
import { userSignUpRequest, userLoginRequest } from '../redux/actions';

interface Props {
    userSignUpRequest: typeof userSignUpRequest
    userLoginRequest: typeof userLoginRequest
    auth: authState
    history: any
}

interface State {
    email: string
    password: string
    errors: ErrorEntity
    timeout: boolean
}

class Form extends React.Component<Props, State> {
    constructor(props: Readonly<Props>) {
        super(props);
        this.state = {
            email: '',
            password: '',
            errors: {},
            timeout: false,
        };

        this.onFieldChange = this.onFieldChange.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }


    private onFieldChange(fieldName: string, value: string) {
        this.setState({ ...this.state, [fieldName]: value });
    }

    private isValid() {
        if (this.state.email === '' || this.state.password === '') {
            return false;
        }
        return true;
    }

    private async onSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        if (this.isValid()) {
            const { email, password } = this.state;
            await this.props.userLoginRequest({email, password});
            if(this.props.auth.user) {
                this.props.history.push('/secret');
            }
        }
    }

    render() {
        const { errors } = this.state;
        return (
            <form onSubmit={this.onSubmit} >

                <TextField
                    error={errors.email}
                    label="Email"
                    onChange={this.onFieldChange}
                    //onBlur={this.checkUserExists}
                    value={this.state.email}
                    field="email"
                    type="text"
                />

                <TextField
                    error={errors.password}
                    label="Password"
                    onChange={this.onFieldChange}
                    value={this.state.password}
                    field="password"
                    type="password"
                />
                
                <button disabled={this.state.timeout} type="submit" className="btn btn-primary">Submit</button>
            </form>
        );
    }
}

export default withRouter<any>(Form);