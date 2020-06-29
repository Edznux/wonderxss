// DEBUG (FIXME)
// export const URL = window.location.protocol + "//" + window.location.hostname + ":" + window.location.port
export const URL = "//" + window.location.hostname;
export const URL_LOGIN = URL + "/login";
export const URL_PAYLOAD = URL + "/p/";
export const URL_OTP_REGISTER = URL + "/otp/new";
export const URL_OTP = URL + "/otp";

export const API_BASE = URL + "/api/v1";
export const API_USER = API_BASE + "/users";
export const API_PAYLOADS = API_BASE + "/payloads";
export const API_INJECTIONS = API_BASE + "/injections";
export const API_ALIASES = API_BASE + "/aliases";
export const API_LOOTS = API_BASE + "/loots";
export const API_EXECUTIONS = API_BASE + "/executions";
export const API_COLLECTORS = API_BASE + "/collectors";
