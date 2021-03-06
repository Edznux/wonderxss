import React from "react";
import AceEditor from "react-ace";
import axios from "axios";
import { Container, Input, Button, TextField } from "@material-ui/core";
import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/theme-github";

import "./PayloadEditor.css";

import { API_PAYLOADS, API_ALIASES } from "../helpers/constants";

class PayloadEditor extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      currentAlias: "",
      currentPayload: "",
      currentPayloadName: "",
      currentContentType: "",
    };
  }
  createAlias = (payload_id, alias) => {
    console.log(payload_id, alias);
    if (!payload_id || !alias) {
      console.log("missing payload ID or alias name, not creating alias");
      return;
    }
    axios
      .post(API_ALIASES, {
        alias: alias,
        payload_id: payload_id,
      })
      .then(res => {
        console.log(res);
      });
  };
  createPayload = () => {
    if (this.state.currentPayloadName === "") {
      alert("please provide a payload name.");
      return;
    }

    axios
      .post(API_PAYLOADS, {
        name: this.state.currentPayloadName,
        content: this.state.currentPayload,
        content_type: this.state.currentContentType,
      })
      .then(res => {
        let payload_id = "";
        console.log(res.data.data);
        if (res.data.data && !res.data.error) {
          payload_id = res.data.data.ID;
          if (this.state.currentAlias !== "") {
            this.createAlias(payload_id, this.state.currentAlias);
          }
        }
      });
  };
  render() {
    return (
      <Container className="container">
        <h1>Payload editor</h1>
        <AceEditor
          width="100%"
          ref="aceEditor"
          mode="javascript"
          theme="github"
          name="editor"
          editorProps={{ $blockScrolling: true }}
          onChange={data => {
            this.setState({ currentPayload: data });
          }}
          value={this.state.currentPayload}
        />
        <Input
          type="text"
          placeholder="Payload name"
          onChange={event => {
            this.setState({ currentPayloadName: event.target.value });
          }}
        ></Input>
        <Input
          type="text"
          placeholder="short-alias"
          onChange={event => {
            this.setState({ currentAlias: event.target.value });
          }}
        ></Input>
        <TextField
          className="input-text"
          type="text"
          placeholder="Content Type"
          onChange={event => {
            this.setState({ currentContentType: event.target.value });
          }}
        ></TextField>

        <Button onClick={this.createPayload} className="input-text">
          Create payload
        </Button>
      </Container>
    );
  }
}
export default PayloadEditor;
