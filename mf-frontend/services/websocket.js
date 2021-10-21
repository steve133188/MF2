import { w3cwebsocket as W3CWebSocket } from "websocket";

const url = process.env.WS_URL

export const client = new W3CWebSocket('ws://127.0.0.1:8000');


