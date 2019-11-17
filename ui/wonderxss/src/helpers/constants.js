// DEBUG (FIXME)
// const URL = window.location.protocol + "//" + window.location.hostname + ":" + window.location.port
const URL = "https://" + window.location.hostname

export const API_BASE = URL + "/api/v1"
export const API_PAYLOADS = API_BASE + "/payloads"
export const API_ALIASES = API_BASE + "/aliases"
export const API_COLLECTORS = API_BASE + "/collectors"
export const API_EXECUTIONS = API_BASE + "/executions"