import {Client} from '@heroiclabs/nakama-js';
import {WebSocket} from 'ws';
Object.assign(global, { WebSocket: WebSocket });

console.log(WebSocket);

var client = new Client("defaultkey", "127.0.0.1", 7350);

console.log(client);

var randomString = function() {
    let str = ""
    for (let i = 0; i < 13; i++) {
        str += Math.floor(10 * Math.random())
    }
    return str
}

var device_id = randomString();
var username = device_id;

var create = true
var session = await client.authenticateDevice(device_id, create, username)
console.log("auth success", session);

const account = await client.getAccount(session);
console.log("account is", account)

const secure = false; // Enable if server is run with an SSL certificate
const trace = false;
const socket = client.createSocket(secure, trace);
socket.ondisconnect = (evt) => {
    console.info("Disconnected", evt);
};

var appearOnline = true
session = await socket.connect(session, appearOnline);

console.log("session is", session)

let result = await client.listMatches(session)

let dec = new TextDecoder()
socket.onmatchdata = (result) => {
    console.log("tx receipt: ", dec.decode(result.data));
}

let match_id = result.matches[0].match_id
let match = await socket.joinMatch(match_id)
console.log("i joined the match: ", match)
