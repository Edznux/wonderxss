import React from 'react'
import { Container, TextField, Select } from '@material-ui/core';
import EnhancedTable from '../components/Table';
import { API_ALIASES, API_PAYLOADS } from '../helpers/constants';
import axios from 'axios';

export default class Aliases extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      currentPayload: "",
      currentAlias: "",
      headCells : [
        { id: 'ID', numeric: false, disablePadding: true, label: 'ID', ellipsis: true },
        { id: 'Name', numeric: false, disablePadding: false, label: 'Name' },
        { id: 'Payload', numeric: false, disablePadding: false, label: 'Content', ellipsis: true },
        { id: 'Created_At', numeric: false, disablePadding: false, label: 'Created At', ellipsis: true },
      ],
      aliases: [],
      payloads: []
    }
  } 
  setCurrentPayload = (event) =>{
    this.setState({"currentPayload": event.target.value})
  }
  formatPayloadContent = (payloadID) => {
    // FIXME
    // This is sooooo ugly
    // In a perfect world, we should just have a map[payloadID]payload
    // This way, formating with the payloadID would only be this.state.payloads[id].name
    let fmt = ""
    let res = this.state.payloads.filter(function(payload){
      return payload.id === payloadID
    })[0]
    if (res){
      fmt = `[${res.name}] ${res.content.slice(0, 50)}`
    }
    return fmt
  }
  createAlias = () => {    
    axios.post(API_ALIASES, {
      alias: this.state.currentAlias,
      payload_id: this.state.currentPayload
    })
    .then(res => {
      console.log(res)
    });
  }
  componentDidMount() {
    axios.get(API_PAYLOADS).then(res => {
      if (res.status !== 200) {
        throw new Error("Couldn't load payloads")
      } else {
        return res.data
      }
    }).then((rows) => {
      console.log("rows.data payload: ", rows.data)
      this.setState({ payloads: rows.data})
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
            row.alias,
            this.formatPayloadContent(row.payload_id),
            row.created_at,
          ])
        })
        this.setState({
          aliases: tmp
        })
      });
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
          onChange={(event) => this.setState({ currentAlias: event.target.value })}
        />
        Payload : 
        <Select onChange={this.setCurrentPayload}>
          {
            this.state.payloads.map((payload) => {
              return <option value={payload.id}>{payload.name}</option>
            })
          }
        </Select>
        <input type="submit" value="Submit" onClick={this.createAlias} />
      </Container>
    )
  }
}