import React from 'react';
import axios from 'axios';
import TextField from '@material-ui/core/TextField';
import { Container, InputLabel } from '@material-ui/core';


import { setAuthToken } from "../helpers/auth";
import './Login.css';

class Login extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            login: "",
            password: "",
            error: false
        }
    };
    handleSubmit(event) {
        event.preventDefault();
        axios({
            url: '/login',
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data: "login=" + this.state.login + "&password=" + this.state.password,
        })
            .then(res => {
                let data = res.data
                if (data.data) {
                    setAuthToken(data.data)
                    this.props.history.push(`/`)
                } else {
                    this.setState({ error: true })
                }
            })
            .catch(err => {
                this.setState({ error: true })
            })
    }
    render() {
        return (
            <Container>
                <div id="error" hidden={!this.state.error}>Bad credentials</div>
                <div id="login-form">
                    {/* <FormGroup onSubmit={(event) => this.handleSubmit(event)}> */}

                    <form className="login-flex-container" onSubmit={(event) => this.handleSubmit(event)}>
                        <InputLabel className="login-field">
                            <span className="login-text">Username:</span>
                            <TextField
                                className="login-field"
                                hintText="Enter your username"
                                floatingLabelText="login"
                                onChange={(event) => this.setState({ login: event.target.value })}
                                classes="login-field"
                            />
                        </InputLabel>
                        <InputLabel className="login-field">
                            <span className="login-text">Password:</span>
                            <TextField
                                className="login-field"
                                type="password"
                                hintText="Enter your password"
                                floatingLabelText="password"
                                classes="login-field"
                                onChange={(event) => this.setState({ password: event.target.value })}
                            />
                        </InputLabel>
                        <input type="submit" value="Submit" className="login-field" />
                        {/* </FormGroup> */}
                    </form>
                </div>
            </Container>
        )
    }
}
export default Login