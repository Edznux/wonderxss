import React from 'react'
import { Container, Button, Input } from '@material-ui/core';
import { URL_OTP_REGISTER } from '../helpers/constants';
import { getLoginFromJWT, setAuthToken } from '../helpers/auth';
import QRCode from "qrcode.react";
import axios from 'axios';

export default class OTPPage extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      user: "",
      token: "",
      secret: "",
      qrcode: ""
    }
  }
  registerOTP = () => {
    axios.post(URL_OTP_REGISTER, {
      token: this.state.token,
      secret: this.state.secret,
    }).then(res => {
      if (res.status !== 200) {
        throw new Error("Couldn't save OTP token")
      } else {
        return res.data
      }
    }).then((res) => {
        console.log(res)
        if(res.data) {
            setAuthToken(res.data)
        }
    });
  }
  componentDidMount() {
    this.setState({user: getLoginFromJWT()})

    axios.get(URL_OTP_REGISTER).then(res => {
      if (res.status !== 200) {
        throw new Error("Couldn't load payloads")
      } else {
        return res.data
      }
    }).then((res) => {
      this.setState({ secret: res.data.secret})
    });
  }
  render() {
    return (
      <Container className="container">
        <QRCode value={`otpauth://totp/${this.state.user}?secret=${this.state.secret}&issuer=WonderXSS`} />
        <Input onChange={(event) => this.setState({ token: event.target.value })}></Input>
        <Button onClick={this.registerOTP}>Validate</Button>
      </Container>
    )
  }
}