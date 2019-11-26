import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import { Route, Link, BrowserRouter as Router, Switch } from 'react-router-dom'
import App from './pages/App';
import Payloads from './pages/Payloads';
import PayloadEditor from './pages/PayloadEditor';
import NotFound from './pages/NotFound';

window.ws = new WebSocket("wss://localhost:4443/ws")


const routing = (
<Router>
    <ul>
        <li>
            <Link to="/">Home</Link>
        </li>
        <li>
            <Link to="/payloads">Payloads</Link>
        </li>
        <li>
            <Link to="/editor">Payload Editor</Link>
        </li>
        <li>
            <Link to="/aliases">Aliases</Link>
        </li>
    </ul>
    <Switch>
        <Route exact path="/" component={App} />
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
