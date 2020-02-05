import React from "react";
import {
  Container,
  TextField,
  Select,
  InputLabel,
  Button,
  Grid,
  Input,
  Dialog,
  DialogTitle,
  DialogActions,
  DialogContent
} from "@material-ui/core";

import EnhancedTable from "../components/Table";
import { API_ALIASES, API_PAYLOADS } from "../helpers/constants";
import axios from "axios";
import "./Aliases.css";

export default class Aliases extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      currentPayload: "",
      currentAlias: "",
      headCells: [
        {
          id: "ID",
          field: "id",
          numeric: false,
          disablePadding: true,
          label: "id",
          hidden: true,
        },
        {
          id: "Name",
          field: "alias",
          numeric: false,
          disablePadding: true,
          label: "Name",
        },
        {
          id: "Payload",
          field: "payload_id",
          numeric: false,
          disablePadding: true,
          label: "Content",
          ellipsis: true,
        },
        {
          id: "Created_At",
          field: "created_at",
          numeric: false,
          disablePadding: true,
          label: "Created At",
          ellipsis: true,
        },
      ],
      aliases: [],
      payloads: [],
    };
  }
  setCurrentPayload = event => {
    this.setState({ currentPayload: event.target.value });
  };
  formatPayloadContent = payloadID => {
    // FIXME
    // This is sooooo ugly
    // In a perfect world, we should just have a map[payloadID]payload
    // This way, formating with the payloadID would only be this.state.payloads[id].name
    let fmt = "";
    let res = this.state.payloads.filter(function(payload) {
      return payload.id === payloadID;
    })[0];
    if (res) {
      fmt = `[${res.name}] ${res.content.slice(0, 50)}`;
    }
    return fmt;
  };
  createAlias = () => {
    axios
      .post(API_ALIASES, {
        alias: this.state.currentAlias,
        payload_id: this.state.currentPayload,
      })
      .then(res => {
        console.log(res);
      });
  };
  handleClickOpenClose = () => {
    this.setState({
      openDialog: !this.state.openDialog,
    });
  };

  componentDidMount() {
    axios
      .get(API_PAYLOADS)
      .then(res => {
        if (res.status !== 200) {
          throw new Error("Couldn't load payloads");
        } else {
          return res.data;
        }
      })
      .then(rows => {
        console.log("rows.data payload: ", rows.data);
        this.setState({ payloads: rows.data });
        axios
          .get(API_ALIASES)
          .then(res => {
            if (res.status !== 200) {
              throw new Error("Couldn't load payloads");
            } else {
              return res.data;
            }
          })
          .then(rows => {
            rows.data.map(row => {
              row.payload_id = this.formatPayloadContent(row.payload_id);
              return row;
            });
            this.setState({
              aliases: rows.data,
            });
          });
      });
  }
  render() {
    return (
      <Container className="container">
        <h1>Aliases</h1>
        {
          this.state.aliases.length > 0 ?
          <EnhancedTable
            headCells={this.state.headCells}
            data={this.state.aliases}
            isDeleteButtonEnabled={true}
          ></EnhancedTable>
          : 
          <div>No aliases found</div>
        }
        <Button
          className="submit-button"
          variant="outlined"
          color="primary"
          onClick={this.handleClickOpenClose}
        >
        Create new alias
        </Button>
        <Dialog
          open={this.state.openDialog}
          onClose={this.handleClickOpenClose}
          aria-labelledby="form-dialog-title"
        >
          <DialogTitle id="form-dialog-title">New alias</DialogTitle>

          <DialogContent>
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <Input
                  type="text"
                  placeholder="Alias"
                  className="input"
                  onChange={event =>
                    this.setState({ currentAlias: event.target.value })
                  }
                ></Input>
              </Grid>
              <Grid item xs={12}>
                <Select onChange={this.setCurrentPayload} style={{width:"100%"}}>
                  {this.state.payloads.map(payload => {
                    return <option value={payload.id}>{payload.name}</option>;
                  })}
                </Select>
              </Grid>
            </Grid>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClickOpenClose} color="primary">
              Cancel
            </Button>
            <Button onClick={this.createAlias} color="primary">
              Create
            </Button>
          </DialogActions>
        </Dialog>
      </Container>
    );
  }
}
