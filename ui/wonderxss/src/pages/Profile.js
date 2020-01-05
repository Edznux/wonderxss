import React from 'react'
import { Container } from '@material-ui/core';
import { Link } from "react-router-dom";
import OTPRegister from '../components/OTP';
import UserInfo from '../components/UserInfo';
import { Redirect } from 'react-router-dom';

export default class Profile extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      currentPage:""
    }
  } 

  componentDidMount() {
    this.setState({currentPage:  this.props.match.params.subpage})
  }
  selectPage = () =>{
    switch(this.state.currentPage){
      case "/profile/otp":
        return <OTPRegister></OTPRegister>
      case "/profile/info":
        return <UserInfo></UserInfo>
      default:
        this.setState({currentPage:"/profile/info" })
        return <Redirect to="/profile/info"/>
    }
  }
  render() {
    return (
      <Container className="container">
        <h1>Profile</h1>
        <Link to="/profile/info" onClick={(event) => this.setState({currentPage: event.target.pathname }) }>Info</Link>
        <Link to="/profile/otp" onClick={(event) => this.setState({currentPage: event.target.pathname }) }>Security & 2FA</Link>
        {
          this.selectPage()
        }

      </Container>
    )
  }
}