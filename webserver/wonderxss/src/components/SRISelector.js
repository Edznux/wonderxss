import React from "react";
import {
  Select,
  FormLabel,
} from "@material-ui/core";

export default class SRISelector extends React.Component {
    constructor(props) {
        super(props);
        this.onChange = this.onChange.bind(this);
        this.state = {}
    }
    onChange(e) {
      this.props.setCurrentSRI(e)
    }
    render = () => {
        return(
            <FormLabel>
              <span className="input-text-label">SRI Type:</span>
              <Select onChange={this.onChange}>
                {
                  this.props.SRIKinds.map(sri => {
                    return <option value={sri}>{sri}</option>;
                  })
                }
              </Select>
            </FormLabel>
        )
    }
}