import axios from "axios";

// Set default auth
export function setAuthToken(JWTToken) {
  if (JWTToken) {
    localStorage.setItem("jwt", JWTToken);
    axios.defaults.headers.common["Authorization"] = "Bearer " + JWTToken;
  } else {
    localStorage.removeItem("jwt");
    delete axios.defaults.headers.common["Authorization"];
  }
}
export function decodeJWT() {
  let jwt = localStorage.getItem("jwt");
  let splitted = jwt.split(".");
  let data, decoded;

  if (splitted.length !== 3) {
    return false;
  }
  data = splitted[1];

  try {
    decoded = JSON.parse(atob(data));
  } catch (e) {
    console.log(e);
    // If the json parse is crashing, expire the token.
    return false;
  }
  return decoded;
}

export function getLoginFromJWT() {
  return decodeJWT()["user_name"];
}
export function getUserIDFromJWT() {
  return decodeJWT()["user_id"];
}

export function isLoggedIn() {
  if (window.DEBUG) {
    return true;
  }
  let jwt = localStorage.getItem("jwt");
  if (isTokenExpired(jwt)) {
    return false;
  } else {
    setAuthToken(jwt);
    return true;
  }
}

export function isTokenExpired(JWTToken) {
  let decoded, data;
  if (JWTToken === null) {
    return true;
  }

  try {
    data = JWTToken.split(".")[1];
    decoded = JSON.parse(atob(data));
  } catch (e) {
    console.log(e);
    // If the json parse is crashing, expire the token.
    return true;
  }
  // check if the token is still "valid" (only check for non expiration)
  if (decoded.exp && decoded.exp < Math.floor(Date.now() / 1000)) {
    return true;
  }
  return false;
}
