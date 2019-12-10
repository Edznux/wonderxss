import React from 'react'
import PayloadsTable from "../components/PayloadsTable"
import { Container } from '@material-ui/core';

class Payloads extends React.Component {
    render() {
        return (
            <Container>
                <h1>Payloads</h1>
                <PayloadsTable></PayloadsTable>
            </Container>
        )
    }
}
export default Payloads