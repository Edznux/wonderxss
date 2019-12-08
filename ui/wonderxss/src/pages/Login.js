import React from 'react';
import axios from 'axios';
import TextField from '@material-ui/core/TextField';


import { setAuthToken } from "../helpers/auth";
import './Login.css';

class Login extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            login:"",
            password:"",
            error:false
        }
    };
    handleSubmit(event){ 
        event.preventDefault();
        axios({
            url: '/login',
            method: "POST",
            headers: {
                "Content-Type":"application/x-www-form-urlencoded"
            },
            data: "login="+ this.state.login+"&password="+ this.state.password,
        })
        // .then(res => res)
        .then(res => {
            console.log(res)
            let data = res.data
            if (data.data) {
                setAuthToken(data.data)
                // redirectHome()
            } else {
                this.setState({error:true})
            }
        })
    }
    render() {
        return (
            <div>
                <div id="bad-credentials" hidden={!this.state.error}>Bad credentials</div>
                <form onSubmit={(event) => this.handleSubmit(event)}>
                    Username: 
                    <TextField
                        hintText="Enter your username"
                        floatingLabelText="login"
                        onChange={(event) => this.setState({ login: event.target.value })}
                    />
                    Password:
                    <TextField
                        type="password"
                        hintText="Enter your password"
                        floatingLabelText="password"
                        onChange={(event) => this.setState({ password: event.target.value })}
                    />
                    <input type="submit" value="Submit" />
                </form>
            </div>
        )
    }
}
export default Login