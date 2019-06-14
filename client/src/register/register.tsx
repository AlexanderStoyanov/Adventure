import React from 'react';
import Form from './form';

interface Props {
    
}

class Register extends React.Component<Props> {
    render() {
        return (
            <div className="container">
                <div className="row mt-5">
                    <div className="col-md-4 mx-auto">
                        <Form/>
                    </div>
                </div>
            </div>
        );
    }
}

export default Register;