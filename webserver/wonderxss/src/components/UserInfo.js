import React from "react";
import { API_USER } from "../helpers/constants";
import { getUserIDFromJWT } from "../helpers/auth";
import axios from "axios";

export default class UserInfo extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: "",
      userId: "",
      twoFactorEnabled: false,
      createdAt: "",
    };
  }

  componentDidMount() {
    axios
      .get(API_USER + "/" + getUserIDFromJWT())
      .then(res => {
        if (res.status !== 200) {
          throw new Error("Couldn't load payloads");
        } else {
          return res.data;
        }
      })
      .then(res => {
        this.setState({
          user: res.data.username,
          userId: res.data.id,
          twoFactorEnabled: res.data.two_factor_enabled,
          createdAt: res.data.created_at,
        });
      });
  }
  render() {
    return (
      <div>
        <h2>{this.state.user}</h2>
        <div>
          2FA Enabled:
          {this.state.twoFactorEnabled ? (
            <span style={{ color: "green" }}> YES</span>
          ) : (
            <span style={{ color: "red" }}> NO</span>
          )}
          <br />
          Created at: {this.state.createdAt} <br />
        </div>
      </div>
    );
  }
}
