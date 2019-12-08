import React from 'react';
import axios from 'axios';

import EnhancedTable from './Table';
import { API_PAYLOADS } from "../helpers/constants"


export default class PayloadsTable extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            payloads: [],
            headCells : [
                { id: 'ID', numeric: false, disablePadding: true, label: 'ID', ellipsis:true },
                { id: 'Name', numeric: false, disablePadding: false, label: 'Name' },
                { id: 'Content', numeric: false, disablePadding: false, label: 'Content', ellipsis: true  },
                { id: 'Hash', numeric: false, disablePadding: false, label: 'Hash', ellipsis: true  },
                { id: 'Created_At', numeric: false, disablePadding: false, label: 'Created At', ellipsis: true  },
            ],
        }
    };

    componentDidMount(){
        axios.get(API_PAYLOADS).then(res => {
            if (res.status !== 200){
                throw new Error("Couldn't load payloads")
            }else{
                return res.data
            }
        }).then((rows) => {
            console.log(rows.data)
            let tmp = [];
            rows.data.map((row) => {
                return tmp.push([
                    row.id,
                    row.name,
                    row.content,
                    row.hash,
                    row.created_at,
                ])
            })
            this.setState({
                payloads: tmp
            })
        });
    }
    render(){
        return (
        <div className="Payloads">
            <span>Payloads:</span>
            <EnhancedTable headCells={this.state.headCells} data={this.state.payloads}></EnhancedTable>
        </div>
        );
    };
}