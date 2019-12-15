import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import { Route, Link, BrowserRouter as Router, Switch, Redirect} from 'react-router-dom'
import App from './pages/App';
import Payloads from './pages/Payloads';
import Login from './pages/Login';
import Logout from './pages/Logout';
import PayloadEditor from './pages/PayloadEditor';
import NotFound from './pages/NotFound';
import { isLoggedIn } from './helpers/auth';
import { Box } from '@material-ui/core';
import Aliases from './pages/Aliases';

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


const PrivateRoute = ({ component: Component, ...rest }) => (
    <Route {...rest} render={(props) => (
        isLoggedIn() === true
            ? <Component {...props} />
            : <Redirect to='/login' />
    )} />
)


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
            {
                !isLoggedIn() && <Link to="/login">Login</Link>
            }
            {
                isLoggedIn() && <Link to="/logout">Logout</Link>
            }
        </Box>
    </div>
    <Switch>
        <Route path="/login" component={Login} />
        <PrivateRoute exact path="/" component={App} />
        <PrivateRoute path="/payloads" component={Payloads} />
        <PrivateRoute path="/editor" component={PayloadEditor} />
        <PrivateRoute path="/aliases" component={Aliases} />
        <PrivateRoute path="/logout" component={Logout} />
        <Route component={NotFound} />
    </Switch>
</Router>
);

ReactDOM.render(routing, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
