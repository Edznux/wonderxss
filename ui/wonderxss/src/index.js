import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import { Route, Link, BrowserRouter as Router, Switch } from 'react-router-dom'
import App from './pages/App';
import Payloads from './pages/Payloads';
import Login from './pages/Login';
import PayloadEditor from './pages/PayloadEditor';
import NotFound from './pages/NotFound';
import { setAuthToken } from './helpers/auth';
import { Box } from '@material-ui/core';

window.ws = new WebSocket("wss://localhost/ws")
window.ws.onerror = function(event){
    console.error("WebSocket error observed:", event);
}
window.ws.onclose = function(event){
    console.error("WebSocket closed:", event);
}
window.ws.onopen = function(event){
    console.error("WebSocket open:", event);
}

setAuthToken(localStorage.getItem("jwt"))

const routing = (
<Router>
    <div id="menu">
        <Box flexGrow={1}>
            <Link to="/">Home</Link>
        </Box>
        <Box flexGrow={1}>
            <Link to="/payloads">Payloads</Link>
        </Box>
        <Box flexGrow={1}>
            <Link to="/editor">Payload Editor</Link>
        </Box>
        <Box flexGrow={1}>
            <Link to="/aliases">Aliases</Link>
        </Box>
            <Box flexDirection="row-reverse" flexGrow={1}>
            <Link to="/login">Login</Link>
        </Box>
    </div>
    <Switch>
        <Route exact path="/" component={App} />
        <Route path="/login" component={Login} />
        <Route path="/payloads" component={Payloads} />
        <Route path="/editor" component={PayloadEditor} />
        <Route path="/alias" component={Payloads} />
        <Route component={NotFound} />
    </Switch>
</Router>
);

ReactDOM.render(routing, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
