require("babel-polyfill") // To fix errors.

const { CurrentTimerPromiseClient } = require("./currenttimer_grpc_web_pb.js");
const { Empty } = require('google-protobuf/google/protobuf/empty_pb.js');
const { GrpcWebClientBase } = require("grpc-web");

function requestLoggerWrapper(fn) {
    return function() {
        let data = {
            method: arguments[0],
            payload: arguments[1],
            metadata: arguments[2],
        }

        console.info("GRPC request: ", data)

        return fn.apply(this, arguments);
    }
}

class CurrentTime {
    constructor() {
        this.service = new CurrentTimerPromiseClient("http://localhost:8081");
    }

    async getTime() {
        let request = new Empty
        let metadata = {"custom-header": "custom-value"}
        return await this.service.getCurrentTime(request, metadata)
    }
}

GrpcWebClientBase.prototype.rpcCall = requestLoggerWrapper(GrpcWebClientBase.prototype.rpcCall)
GrpcWebClientBase.prototype.unaryCall = requestLoggerWrapper(GrpcWebClientBase.prototype.unaryCall)

function convertTimestampToDate(ts) {
    return new Date(ts.seconds * 1000 + ts.nanos / 10e6)
}

setInterval(() => {
    let ct = new CurrentTime()

    ct.getTime()
        .then(response => {
            log("current time: " + convertTimestampToDate(response.toObject().timestamp))
        })
        .catch(e => error(e))
}, 1000)

function log(message) {
    console.info(message)
    document.getElementById("log").innerText = "Log: " + JSON.stringify(message)
}

function error(err) {
    console.error(err)
    document.getElementById("log").innerText = "Error: " + JSON.stringify(err)
}
