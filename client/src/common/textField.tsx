import React from 'react';
import classnames from 'classnames';

interface Props {
    field: string;
    value: string;
    label: string;
    error?: string;
    type: string;
    onChange: (fieldName: string, value: string) => void;
    //onBlur?: ChangeEventHandler;
}

const TextField: React.FunctionComponent<Props> = (props) => {
    return (
        <div className="form-group">
            <label className="control-label">{props.label}</label>
            <input
                value={props.value}
                onChange={onChangeInput(props)}
                //onBlur={props.onBlur}
                name={props.field}
                type={props.type}
                className={classnames('form-control', { 'is-invalid': props.error })}
            ></input>
            {props.error && <span className="help-block">{props.error}</span>}
        </div>
    );
}

const onChangeInput = (props: Props) => (event: React.ChangeEvent<HTMLInputElement>) => {
    props.onChange(event.target.name, event.target.value);
}


export default TextField;