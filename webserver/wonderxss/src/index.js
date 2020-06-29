import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import * as serviceWorker from "./serviceWorker";
import {
  Route,
  Link,
  BrowserRouter as Router,
  Switch,
  Redirect,
} from "react-router-dom";
import App from "./pages/App";
import Payloads from "./pages/Payloads";
import Login from "./pages/Login";
import ProfilePage from "./pages/Profile";
import Loots from './pages/Loots';
import LootViewer from './pages/LootViewer';
import Logout from "./pages/Logout";
import PayloadEditor from "./pages/PayloadEditor";
import NotFound from "./pages/NotFound";
import { isLoggedIn } from "./helpers/auth";
import { Box } from "@material-ui/core";
import Aliases from "./pages/Aliases";

window.ws = new WebSocket("wss://localhost/ws");
window.ws.onerror = function(event) {
  console.error("WebSocket error observed:", event);
};
window.ws.onclose = function(event) {
  console.error("WebSocket closed:", event);
};
window.ws.onopen = function(event) {
  console.error("WebSocket open:", event);
};

const PrivateRoute = ({ component: Component, ...rest }) => (
  <Route
    {...rest}
    render={props =>
      isLoggedIn() === true ? (
        <Component {...props} />
      ) : (
        <Redirect to="/login" />
      )
    }
  />
);

const routing = (
  <Router>
    <div className="navigation">
      <div id="menu">
        <Box flexGrow={1}>
          <Link to="/" style={{ width: "100%", display: "block" }}>
            Home
          </Link>
        </Box>
        <Box flexGrow={1}>
          <Link to="/loots">
            Loots
          </Link>
        </Box>
        <Box flexGrow={1}>
          <Link to="/payloads" style={{ width: "100%", display: "block" }}>
            Payloads
          </Link>
        </Box>
        <Box flexGrow={1}>
          <Link to="/editor" style={{ width: "100%", display: "block" }}>
            Payload Editor
          </Link>
        </Box>
        <Box flexGrow={1}>
          <Link to="/aliases" style={{ width: "100%", display: "block" }}>
            Aliases
          </Link>
        </Box>
        {isLoggedIn() && (
          <Box flexGrow={1}>
            <Link to="/profile" style={{ width: "100%", display: "block" }}>
              Profile
            </Link>
          </Box>
        )}
        <Box flexGrow={1}>
          {!isLoggedIn() && (
            <Link to="/login" style={{ width: "100%", display: "block" }}>
              Login
            </Link>
          )}
          {isLoggedIn() && (
            <Link to="/logout" style={{ width: "100%", display: "block" }}>
              Logout
            </Link>
          )}
        </Box>
      </div>
    </div>
    <Switch>
      <Route exact path="/login" component={Login} />
      <PrivateRoute exact path="/" component={App} />
      <PrivateRoute path="/payloads" component={Payloads} />
      <PrivateRoute path="/profile/:subpage" component={ProfilePage} />
      <PrivateRoute path="/profile" component={ProfilePage} />
      <PrivateRoute path="/editor" component={PayloadEditor} />
      <PrivateRoute path="/aliases" component={Aliases} />
      <PrivateRoute path="/logout" component={Logout} />
      <PrivateRoute path="/loots/:id" component={LootViewer} />
      <PrivateRoute path="/loots" component={Loots} />
      <Route component={NotFound} />
    </Switch>
  </Router>
);

ReactDOM.render(routing, document.getElementById("root"));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
