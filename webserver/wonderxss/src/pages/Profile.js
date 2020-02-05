import React from "react";
import { Container } from "@material-ui/core";
import { Link } from "react-router-dom";
import OTPRegister from "../components/OTP";
import UserInfo from "../components/UserInfo";
import { Redirect } from "react-router-dom";

import "./Profile.css";

export default class Profile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      currentPage: "",
    };
  }

  componentDidMount() {
    this.setState({ currentPage: this.props.match.params.subpage });
  }
  selectPage = () => {
    switch (this.state.currentPage) {
      case "/profile/otp":
        return <OTPRegister></OTPRegister>;
      case "/profile/info":
        return <UserInfo></UserInfo>;
      default:
        this.setState({ currentPage: "/profile/info" });
        return <Redirect to="/profile/info" />;
    }
  };
  render() {
    return (
      <Container className="container profile-container">
      <div class="profile-sidebar">
        <div class="profile-sidebar-item">
          <Link
            to="/profile/info"
            onClick={event =>
              this.setState({ currentPage: event.target.pathname })
            }
            >
            Info
          </Link>
        </div>
        <div class="profile-sidebar-item">
          <Link
            to="/profile/otp"
            onClick={event =>
              this.setState({ currentPage: event.target.pathname })
            }
            >
            Security & 2FA
          </Link>
        </div>
      </div>
      <div class="profile-body">
        <h1>Profile</h1>
        {this.selectPage()}
      </div>
      </Container>
    );
  }
}
