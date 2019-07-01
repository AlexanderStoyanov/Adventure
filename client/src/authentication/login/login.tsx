import React from 'react';
import { bindActionCreators, Dispatch } from 'redux';
import Form from './form';
import { connect } from 'react-redux';
import { AppState } from '../../rootReducer';
import { signUpState, authState } from '../redux/types';
import { userSignUpRequest, userLoginRequest } from '../redux/actions';

interface Props {
    userSignUpRequest: typeof userSignUpRequest
    userLoginRequest: typeof userLoginRequest
    signUp: signUpState
    auth: authState
    history: any
}

class Login extends React.Component<Props> {
    render() {
        const { ...props } = this.props;
        return (
            <div className="container">
                <div className="row mt-5">
                    <div className="col-md-4 mx-auto">
                        <Form
                            {...props}
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
        userSignUpRequest,
        userLoginRequest
    },
    dispatch
)

export default connect(mapStateToProps, mapDispatchToProps)(Login);