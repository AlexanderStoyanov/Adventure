import isEmpty from 'lodash/isEmpty';
import ErrorEntity from './model/errorEntity';

interface Input {
    username: string;
    email: string;
    password: string;
}

export default function validateInput(input: Input) {
    let errors: ErrorEntity = {};

    if (input.username === '') {
        errors.username = 'Username is required';
    }
    if (input.email === '') {
        errors.email = 'Email is required';
    }
    if (input.password === '') {
        errors.password = 'Password is required';
    }

    return {
        errors,
        isValid: isEmpty(errors)
    };
}