import React from "react";
import { Container } from "@material-ui/core";
import { setAuthToken } from "../helpers/auth";

export default class Lougout extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      login: "",
      password: "",
      error: false,
    };
  }
  componentDidMount() {
    setAuthToken("");
    this.props.history.push(`/login`);
  }
  render() {
    return <Container></Container>;
  }
}
