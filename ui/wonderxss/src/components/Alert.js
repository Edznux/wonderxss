import React from 'react'

class Alert extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            alertMsg: ""
        }
    }
    componentDidMount() {
        console.log("Mount alert")
        window.ws.onmessage = function (msg) {
            console.log("Recieved message !", msg)
            var jsonMsg = JSON.parse(msg)
            alert(msg)
            this.setState({"alertMsg":jsonMsg.content})
        }
    }
    render() {
        return (
            <div>
                <span>
                    {this.state.alertMsg}
                </span>
            </div>
        )
    }
}
export default Alert