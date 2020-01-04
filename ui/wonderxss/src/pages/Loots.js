import React from 'react'
import { Container } from '@material-ui/core';
import EnhancedTable from '../components/Table';


export default class Loots extends React.Component {
    constructor(props) {
        super(props)

        // Test values.
        const dateObject = new Date();
        const date = dateObject.getMonth()+1 + '/' + dateObject.getDay() + '/' + dateObject.getFullYear();

        this.state = {
            payloads: [ 
                {'id': 1, 'link': 'loot/1', 'date': date}
            ],
            headCells: [
                { id: 'ID', field: "id", numeric: false, disablePadding: true, label: 'ID', ellipsis: true},
                { id: 'Link', field: "link", numeric: false, disablePadding: false, label: 'Link' },
                { id: 'Date', field: "date", numeric: false, disablePadding: false, label: 'Date', ellipsis: true },
            ],
        }
    };

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
