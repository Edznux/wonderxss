import React from "react";
import axios from "axios";

import { Container } from "@material-ui/core";
import EnhancedTable from "../components/Table";
import { API_PAYLOADS } from "../helpers/constants";

class Payloads extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      payloads: [],
      headCells: [
        {
          id: "ID",
          field: "id",
          numeric: false,
          disablePadding: true,
          label: "ID",
          ellipsis: true,
        },
        {
          id: "Name",
          field: "name",
          numeric: false,
          disablePadding: false,
          label: "Name",
        },
        {
          id: "Content",
          field: "content",
          numeric: false,
          disablePadding: false,
          label: "Content",
          ellipsis: true,
        },
        {
          id: "ContentType",
          field: "content_type",
          numeric: false,
          disablePadding: false,
          label: "Content Type",
          ellipsis: true,
        },
        {
          id: "Hashes",
          field: "hashes",
          numeric: false,
          disablePadding: false,
          label: "Hashes",
          ellipsis: true,
        },
        {
          id: "Created_At",
          field: "created_at",
          numeric: false,
          disablePadding: false,
          label: "Created At",
          ellipsis: true,
        },
      ],
    };
  }
  formatSRI(hashes) {
    let res = [];
    for (var key in hashes) {
      if (hashes.hasOwnProperty(key)) {
        res.push(<li>{hashes[key]}</li>);
      }
    }
    return res;
  }
  componentDidMount() {
    axios
      .get(API_PAYLOADS)
      .then(res => {
        if (res.status !== 200) {
          throw new Error("Couldn't load payloads");
        } else {
          return res.data;
        }
      })
      .then(rows => {
        // This override the SRIHashes object to a HTML List.
        for (var row in rows.data) {
          rows.data[row].hashes = this.formatSRI(rows.data[row].hashes);
        }
        this.setState({
          payloads: rows.data,
        });
      });
  }
  render() {
    return (
      <Container className="container">
        <h1>Payloads</h1>
        <div className="Payloads">
          {
            this.state.payloads.length > 0 ?
            <EnhancedTable
              headCells={this.state.headCells}
              data={this.state.payloads}
              isDeleteButtonEnabled={true}
            ></EnhancedTable>
            : 
            <div>No payloads found</div>
          }
        </div>
      </Container>
    );
  }
}
export default Payloads;
