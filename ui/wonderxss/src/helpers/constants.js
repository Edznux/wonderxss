// DEBUG (FIXME)
// export const URL = window.location.protocol + "//" + window.location.hostname + ":" + window.location.port
export const URL = "https://" + window.location.hostname
export const URL_LOGIN = URL + "/login"
export const URL_PAYLOAD = URL + "/p/"

export const API_BASE = URL + "/api/v1"
export const API_PAYLOADS = API_BASE + "/payloads"
export const API_ALIASES = API_BASE + "/aliases"
export const API_COLLECTORS = API_BASE + "/collectors"
export const API_EXECUTIONS = API_BASE + "/executions"