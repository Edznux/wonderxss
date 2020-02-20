import React from "react";
import axios from "axios";

import {
  API_PAYLOADS,
  URL_PAYLOAD,
  API_INJECTIONS,
  API_ALIASES,
} from "../helpers/constants";
import {
  Container,
  Button,
  Input,
  Grid,
  Dialog,
  DialogTitle,
  DialogActions,
  DialogContent,
  TextField,
} from "@material-ui/core";
import EnhancedTable from "./Table";
import "../styles/InjectionsList.css";
import PayloadSelector from "./PayloadSelector";
import SRISelector from "./SRISelector";
import CustomCheckbox from "./CustomCheckbox";
import CrossOriginSelector from "./CrossOriginSelector";

const REPLACE_URL_TAG = "##URL##";
const REPLACE_SRI_TAG = "##SRI_HASH##";
const REPLACE_CROSSORIGIN = "##CROSSORIGIN##";

const SRIKinds = ["sha256", "sha384", "sha512"];
const CROSSORIGIN = ["anonymous", "use-credentials"];

export default class InjectionsList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      headCells: [
        {
          id: "ID",
          field: "id",
          numeric: false,
          disablePadding: false,
          label: "id",
          ellipsis: true,
          hidden: true,
        },
        {
          id: "Name",
          field: "name",
          numeric: false,
          disablePadding: false,
          label: "Name",
          ellipsis: true,
        },
        {
          id: "Injection",
          field: "formatedContent",
          numeric: false,
          disablePadding: false,
          label: "Injection",
        },
        {
          id: "Created_At",
          field: "created_at",
          numeric: false,
          disablePadding: false,
          label: "Created At",
          ellipsis: true,
        },
      ],
      newTitle: "",
      newContent: "",
      currentAlias: "",
      aliasesOrPayloadsIDs: [""],
      useSubdomain: false,
      useHTTPS: true,
      SRIKind: SRIKinds[0],
      injections: [],
      crossorigin: "",
      openDialog: false,
    };
  }
  setCurrentAlias = event => {
    let s = event.target.value;
    if (!s) {
      s = "";
    }
    this.setState({ currentAlias: s });
  };
  setCurrentSRI = event => {
    let s = event.target.value;
    let found = SRIKinds.indexOf(s);
    if (found > -1) {
      this.setState({ SRIKind: SRIKinds[found] }, this.formatInjections);
    } else {
      this.setState({ SRIKind: SRIKinds[0] }, this.formatInjections);
    }
  };
  setCurrentCrossOrigin = event => {
    let selected = event.target.value;
    let found = CROSSORIGIN.indexOf(selected);
    if (found > -1) {
      this.setState({ crossorigin: CROSSORIGIN[found] }, this.formatInjections);
    } else {
      this.setState({ crossorigin: CROSSORIGIN[0] }, this.formatInjections);
    }
  };
  toggleSubdomain = event => {
    this.setState({ useSubdomain: event.target.checked }, this.formatInjections);
  };
  toggleHTTPS = event => {
    this.setState({ useHTTPS: event.target.checked }, this.formatInjections);
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
        let tmp = this.state.aliasesOrPayloadsIDs;
        console.log("Payloads: ", rows.data);
        rows.data.map(row => {
          return tmp.push([row.id, row.name]);
        });
        this.setState({
          aliasesOrPayloadsIDs: tmp,
        }, this.formatInjections);
      });

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
        let tmp = this.state.aliasesOrPayloadsIDs;
        console.log("Aliases: ", rows.data);
        rows.data.map(row => {
          // Don't push alias ID, only the alias name to the selector.
          // We don't refer to the payload using the alias ID, only its alias (and the payload id itself)
          return tmp.push([row.alias, row.alias]);
        });
        this.setState({
          aliasesOrPayloadsIDs: tmp,
        }, this.formatInjections);
      });

    axios
      .get(API_INJECTIONS)
      .then(res => {
        if (res.status !== 200) {
          throw new Error("Couldn't load injections");
        } else {
          return res.data;
        }
      })
      .then(rows => {
        // rows.data.map(injection => {
        //   injection.formatedContent = this.formatInjection(injection.content);
        //   return injection;
        // });
        this.setState({
          injections: rows.data,
        }, this.formatInjections);
      });
  }
  formatInjections = () =>{
    this.state.injections.map(injection => {
      injection.formatedContent = this.formatInjection(injection.content);
      return injection;
    });
  }
  formatInjection = injection => {
    console.log("injection:", injection);
    let url = "";
    let res = "";
    let proto = "http" + (this.state.useHTTPS ? "s" : "") + ":";
    if (this.state.useSubdomain) {
      url =
        proto + "//" + this.state.currentAlias + "." + window.location.hostname; // + ":" + window.location.port;
    } else {
      url = proto + URL_PAYLOAD + this.state.currentAlias;
    }
    res = injection.replace(REPLACE_URL_TAG, url);
    res = res.replace(REPLACE_SRI_TAG, this.state.SRIKind);
    res = res.replace(REPLACE_CROSSORIGIN, this.state.crossorigin);
    console.log("res:", res);
    return res;
  };

  newInjection = event => {
    let injections = this.state.injections;
    injections.push({
      name: this.state.newTitle,
      content: this.state.newContent,
    });
    this.setState({ injections: injections });
    axios
      .post(API_INJECTIONS, {
        name: this.state.newTitle,
        content: this.state.newContent,
      })
      .then(res => {
        console.log("OK, saved injection", res);
      });
    this.handleClickOpenClose();
  };

  handleClickOpenClose = () => {
    this.setState({
      openDialog: !this.state.openDialog,
    });
  };

  render() {
    return (
      <Container>
        <Grid container spacing={3}>
          <Grid item xs={6}>
            <PayloadSelector Payloads={["test?"]} Labels={["Test Labels 1"]}></PayloadSelector>
          </Grid>
          <Grid item xs={6}>
            <SRISelector SRIKinds={SRIKinds} setCurrentSRI={this.setCurrentSRI}></SRISelector>
          </Grid>
          <Grid item xs={6}>
            <CrossOriginSelector CROSSORIGIN={CROSSORIGIN} setCurrentCrossOrigin={this.setCurrentCrossOrigin}></CrossOriginSelector>
          </Grid>
          <Grid item xs={6}>
            <CustomCheckbox labelText={"Use subdomain"} onChange={this.toggleSubdomain}></CustomCheckbox>
          </Grid>
          <Grid item xs={6}>
            <CustomCheckbox labelText={"Use HTTPS:"} onChange={this.toggleHTTPS} checked={true}></CustomCheckbox>
          </Grid>
          <Grid item xs={12}>
            {this.state.injections.length > 0 ? (
              <EnhancedTable
                headCells={this.state.headCells}
                data={this.state.injections}
                isDeleteButtonEnabled={true}
              ></EnhancedTable>
            ) : (
              <center>
                No injections found. <br />
                You can start using the app by creating a new injection.
              </center>
            )}
          </Grid>
          <Grid item xs={12}>
            <Button
              className="submit-button"
              variant="outlined"
              color="primary"
              onClick={this.handleClickOpenClose}
            >
              Create new injection
            </Button>
            <Dialog
              open={this.state.openDialog}
              onClose={this.handleClickOpenClose}
              aria-labelledby="form-dialog-title"
            >
              <DialogTitle id="form-dialog-title">New injection</DialogTitle>

              <DialogContent>
                <small>
                  Replacement list: {REPLACE_CROSSORIGIN}, {REPLACE_URL_TAG},{" "}
                  {REPLACE_SRI_TAG}
                </small>
                <Grid container spacing={3}>
                  <Grid item xs={12}>
                    <Input
                      type="text"
                      placeholder="Name"
                      className="input"
                      onChange={event =>
                        this.setState({ newTitle: event.target.value })
                      }
                    ></Input>
                  </Grid>
                  <Grid item xs={12}>
                    <TextField
                      type="text"
                      placeholder="Injection"
                      className="input"
                      onChange={event =>
                        this.setState({ newContent: event.target.value })
                      }
                      multiline
                    />
                  </Grid>
                </Grid>
              </DialogContent>
              <DialogActions>
                <Button onClick={this.handleClickOpenClose} color="primary">
                  Cancel
                </Button>
                <Button onClick={this.newInjection} color="primary">
                  Create
                </Button>
              </DialogActions>
            </Dialog>
          </Grid>
        </Grid>
      </Container>
    );
  }
}
