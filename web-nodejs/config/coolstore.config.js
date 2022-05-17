var config = {
    MY_SERVICE: "web",
    GATEWAY_SERVICE: "gateway",
    SECURE_GATEWAY_SERVICE: "gateway",
    API_ENDPOINT: "",
    SECURE_API_ENDPOINT: "",
    SSO_ENABLED: process.env.SSO_URL ? true : false
};

if (process.env.COOLSTORE_GW_ENDPOINT != null) {
    config.API_ENDPOINT = process.env.COOLSTORE_GW_ENDPOINT;
} else if (process.env.COOLSTORE_GW_SERVICE != null) {
    config.GATEWAY_SERVICE = process.env.COOLSTORE_GW_SERVICE;
}


if (process.env.SECURE_COOLSTORE_GW_ENDPOINT != null) {
    config.SECURE_API_ENDPOINT = process.env.SECURE_COOLSTORE_GW_ENDPOINT;
} else if (process.env.SECURE_COOLSTORE_GW_SERVICE != null) {
    config.SECURE_GATEWAY_SERVICE = process.env.SECURE_COOLSTORE_GW_SERVICE;
}

console.log("Using API_ENDPOINT " + config.API_ENDPOINT);
console.log("Using GATEWAY_SERVICE " + config.GATEWAY_SERVICE);
console.log("Using SECURE_GATEWAY_SERVICE " + config.SECURE_GATEWAY_SERVICE);
console.log("Using SECURE_API_ENDPOINT " + config.SECURE_API_ENDPOINT);
console.log("Using SSO_ENABLED " + config.SSO_ENABLED);
module.exports = config;
