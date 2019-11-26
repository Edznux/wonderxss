import React from 'react'

class Alert extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            alertMsg = ""
        }
    }
    componentDidMount() {
        window.ws.onmessage = function (msg) {
            var jsonMsg = JSON.parse(msg)
            alert(msg)
            this.state.alertMsg = jsonMsg.content
        }
    }
    render() {
        return (
            <div>
                Payload triggered!
                <span>
                    {this.state.alertMsg}
                </span>
            </div>
        )
    }
}
export default Alert