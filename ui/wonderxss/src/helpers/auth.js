import axios from 'axios';

// Set default auth
export function setAuthToken (JWTToken){
    if (JWTToken) {
        localStorage.setItem("jwt", JWTToken) 
        axios.defaults.headers.common['Authorization'] = "Bearer " + JWTToken;
    } else {
        localStorage.setItem("jwt", "") 
        delete axios.defaults.headers.common['Authorization'];
    }
};

export function isTokenExpired(JWTToken){
    let decoded;
    let data = JWTToken.split(".")[1]
    try {
        decoded = JSON.parse(btoa(data))
    } catch(e){
        console.log(e)
        // If the json parse is crashing, expire the token.
        return true
    }
    // check if the token is still "valid" (only check for non expiration)
    if (decoded.exp && decoded.exp < Math.floor(Date.now() / 1000)){
        return true
    }
    return false
}