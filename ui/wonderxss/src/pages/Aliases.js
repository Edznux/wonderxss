import React from 'react'
import { Container, TextField, Select } from '@material-ui/core';
import EnhancedTable from '../components/Table';
import { API_ALIASES } from '../helpers/constants';
import axios from 'axios';

export default class Aliases extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      newAlias: "",
      headCells : [
        { id: 'ID', numeric: false, disablePadding: true, label: 'ID', ellipsis: true },
        { id: 'Name', numeric: false, disablePadding: false, label: 'Name' },
        { id: 'Payload', numeric: false, disablePadding: false, label: 'Content', ellipsis: true },
        { id: 'Created_At', numeric: false, disablePadding: false, label: 'Created At', ellipsis: true },
      ],
      aliases: []
    }
  } 

  componentDidMount() {
    axios.get(API_ALIASES).then(res => {
      if (res.status !== 200) {
        throw new Error("Couldn't load payloads")
      } else {
        return res.data
      }
    }).then((rows) => {
      console.log(rows.data)
      let tmp = [];
      rows.data.map((row) => {
        return tmp.push([
          row.id,
          row.name,
          row.payload,
          row.created_at,
        ])
      })
      this.setState({
        aliases: tmp
      })
    });
  }
  render() {
    return (
      <Container className="container">
        Aliases list:
        <EnhancedTable headCells={this.state.headCells} data={this.state.aliases}></EnhancedTable>
        Alias : <TextField
          className="alias-field"
          type="text"
          onChange={(event) => this.setState({ newAlias: event.target.value })}
        />
        Payload : 
        <Select>
          <option value="a">a</option>
        </Select>
        <input type="submit" value="Submit" />
      </Container>
    )
  }
}