import React from "react";
import {
  Select,
  FormLabel,
} from "@material-ui/core";

export default class PayloadSelector extends React.Component {
    constructor(props) {
        super(props);
        console.log(this.props)
        this.state = {}
    }
    formatPayloadName = (payloadID, payloadName) => {
        console.log(`formatPayloadName(${payloadID}, ${payloadName})`);
        if (payloadID && payloadName) {
          return payloadID.slice(0, 8) + `... (${payloadName})`;
        }
        return "";
    }
    render = () => {
        return(
            <FormLabel>
                <span className="input-text-label">Payload ID and/or alias:</span>
                <Select onChange={console.log("fixme")}>
                {this.props.Payloads.map(aop => {
                    return (
                        <option value={aop[0]}>
                            {this.formatPayloadName(aop[0], aop[1])}
                        </option>
                    );
                })}
                </Select>
            </FormLabel>
        )
    }
}