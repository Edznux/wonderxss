import React from "react";
import {
  Select,
  FormLabel,
} from "@material-ui/core";

export default class CrossOriginSelector extends React.Component {
    constructor(props) {
        super(props);
        this.onChange = this.onChange.bind(this);
        this.state = {}
    }
    onChange(e) {
      this.props.setCurrentCrossOrigin(e)
    }
    render = () => {
        return(
            <FormLabel>
              <span class="input-text-label">Crossorigin:</span>
              <Select onChange={this.onChange}>
                {this.props.CROSSORIGIN.map(crossorigin => {
                  return <option value={crossorigin}>{crossorigin}</option>;
                })}
              </Select>
            </FormLabel>
        )
    }
}