import React from 'react'
import axios from 'axios';
import AceEditor from "react-ace";
import { Container } from '@material-ui/core';
import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/theme-github";
import { API_COLLECTORS } from "../helpers/constants"
import test from './test';


export default class LootViewer extends React.Component {

    constructor(props) {
        super(props)

        const { id } = this.props.match.params;
        this.state = {
            id: id,
            currentPayload: test.value,
            currentType: test.type,
        }
    };

    componentDidMount() {
        const { id } = this.props.match.params;
        axios.get(API_COLLECTORS + '/' + id).then(res => {
            if (res.status !== 200) {
                throw new Error("Couldn't load collector " + id)
            } else {
                return res.data
            }
        }).then((rows) => {
            this.setState({
                currentPayloads: rows.data.data.value,
                currentType: rows.data.data.type,
            })
        });
    }

    renderEditor = (payload) => {
        return <AceEditor
            width="100%"
            ref="aceEditor"
            mode="javascript"
            theme="github"
            name="editor"
            editorProps={{ $blockScrolling: true }}
            onChange={(data) => { this.setState({ "currentPayload": data }) }}
            value={ payload }
        />;
    }

    renderImage = (payload) => {
        return payload;
    }

    renderDefault = (type) => {
        return type
    }

    displayType = (type, payload) => {
        switch(type) {
            case 'code': 
                return this.renderEditor(payload);
            case 'image':
                return this.renderImage(payload);
            default:
                return this.renderDefault(type);
        }
    }
    
    render() {
        const { id, currentPayload, currentType } = this.state;

        let display = this.displayType(currentType, currentPayload);

        return (
            <Container className="container">
                <h1>Loot {id} </h1>
                {display}
            </Container>
        );
    }
}
