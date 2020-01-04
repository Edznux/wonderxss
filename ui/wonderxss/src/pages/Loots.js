import React from 'react'
import axios from 'axios';
import { Container } from '@material-ui/core';
import EnhancedTable from '../components/Table';
import { API_COLLECTORS } from "../helpers/constants"


export default class Loots extends React.Component {
    constructor(props) {
        super(props)

        // Test values.
        const dateObject = new Date();
        const date = dateObject.getMonth()+1 + '/' + dateObject.getDay() + '/' + dateObject.getFullYear();

        this.state = {
            payloads: [ 
                {'PayloadID': 1, 'link': 'loot/1', 'CreatedAt': date}
            ],
            headCells: [
                { id: 'ID', field: "PayloadID", numeric: false, disablePadding: true, label: 'ID', ellipsis: true},
                { id: 'Link', field: "link", numeric: false, disablePadding: false, label: 'Link' },
                { id: 'Date', field: "CreatedAt", numeric: false, disablePadding: false, label: 'Date', ellipsis: true },
            ],
        }
    };

    componentDidMount() {
        axios.get(API_COLLECTORS).then(res => {
            if (res.status !== 200) {
                throw new Error("Couldn't load collectors")
            } else {
                return res.data
            }
        }).then((rows) => {
            this.setState({
                payloads: rows.data
            })
        });
    }

    render() {
        return (
            <Container>
                <h1>Loots</h1>
                <div className="Payloads">
                    <EnhancedTable headCells={this.state.headCells} data={this.state.payloads} isDeleteButtonEnabled={true}></EnhancedTable>
                </div>
            </Container>
        )
    }
}
