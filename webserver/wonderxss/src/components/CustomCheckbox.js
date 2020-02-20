import React from "react";
import {
  Checkbox,
  FormLabel,
} from "@material-ui/core";

export default class CustomCheckbox extends React.Component {
    constructor(props) {
        super(props);
        this.state = {}
    }
    render = () => {
        return(
            <FormLabel>
              <span class="input-text-label">{this.props.labelText}</span>
              <Checkbox
                inputProps={{ "aria-label": this.props.labelText }}
                onChange={this.props.onChange}
                labelStyle={{ color: "white" }}
                iconStyle={{ fill: "white" }}
                defaultChecked={this.props.checked}
              />
            </FormLabel>
        )
    }
}