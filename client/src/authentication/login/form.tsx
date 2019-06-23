import React from 'react';
import TextField from '../../common/textField';
import ErrorEntity from '../../common/model/errorEntity';

interface Props {
    userSignUpRequest: any
    userLoginRequest: any
}

interface State {
    username: string;
    password: string;
    errors: ErrorEntity;
    timeout: boolean;
}

class Form extends React.Component<Props, State> {
    constructor(props: Readonly<Props>) {
        super(props);
        this.state = {
            username: '',
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
        if (this.state.username === '' || this.state.password === '') {
            return false;
        }
        return true;
    }

    private onSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        if (this.isValid()) {
            const { username, password } = this.state;
            this.props.userLoginRequest({username, password});
        }
    }

    render() {
        const { errors } = this.state;
        return (
            <form onSubmit={this.onSubmit} >

                <TextField
                    error={errors.username}
                    label="Username"
                    onChange={this.onFieldChange}
                    //onBlur={this.checkUserExists}
                    value={this.state.username}
                    field="username"
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

export default Form;