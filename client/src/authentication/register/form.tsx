import React from 'react';
import TextField from '../../common/textField';
import validateInput from '../../common/validateInput';
import ErrorEntity from '../../common/model/errorEntity';

interface Props {
    userSignUpRequest: any
}

interface State {
    username: string
    email: string
    password: string
    errors: ErrorEntity
    timeout: boolean
}

class Form extends React.Component<Props, State> {
    constructor(props: Readonly<Props>) {
        super(props);
        this.state = {
            username: '',
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
        const { errors, isValid } = validateInput(this.state);
        this.setState({ errors });
        return isValid;
    }

    private onSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        if (this.isValid()) {
            const { username, email, password } = this.state;
            //this.setState({ timeout: true });
            this.props.userSignUpRequest({username, email, password});
            // this.setState({ errors: {}, timeout: true });
            // this.props.userSignUpRequest(this.state).then(
            //     (res: any) => {
            //         this.props.addFlashMessage({
            //             type: 'success',
            //             text: 'Account created successfully'
            //         });
            //         localStorage.setItem('token', res.data.token);
            //         //this.props.history.push("/forum");
            //     },
            //     (err: any) => {
            //         this.setState({ errors: err.response.data, timeout: false });
                    
            //     }
            // );
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
                    error={errors.email}
                    label="Email"
                    onChange={this.onFieldChange}
                    value={this.state.email}
                    field="email"
                    type="email"
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