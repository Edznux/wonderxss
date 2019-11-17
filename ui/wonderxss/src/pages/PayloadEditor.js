import React from 'react'
import AceEditor from "react-ace";

import "ace-builds/src-noconflict/mode-javascript";
import "ace-builds/src-noconflict/theme-github";
import { API_PAYLOADS } from '../helpers/constants';


class PayloadEditor extends React.Component {
    
    constructor(props) {
        super(props)
        this.state = {
        }
    };
    render() {
        const createPayload = () => {
            var payloadName = document.getElementById("payload-name").value;
            var payload = this.refs.aceEditor.editor.getValue()
            if(payloadName === ""){
                alert("please provide a payload name.")
                return
            }
            
            console.log("Payload name:", payloadName)
            console.log("Payload content:");
            console.log(payload);

            fetch(API_PAYLOADS, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ name: payloadName, content: payload })
            })
            .then(res => res.json())
            .then(res => console.log(res));
        };
        return (
            <div>
                <h1>Payload Editor</h1>
                <input id="payload-name" type="text" placeholder="Payload name"></input>
                <AceEditor
                    ref="aceEditor"
                    mode="javascript"
                    theme="github"
                    name="editor"
                    editorProps={{ $blockScrolling: true}}
                />,
                <button onClick={createPayload}>Create paylaod</button>
            </div>
        )
    }
}
export default PayloadEditor