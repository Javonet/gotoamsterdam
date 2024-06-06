const WebSocket = require('ws');

class Client {
  constructor() {
    this.ws = new WebSocket('ws://localhost:3000');
  }
  
  connect() {
    console.log(this.ws);
    this.ws.send(JSON.stringify({ payload: "sendInstruction" }));
  }

  send(message) {
    if (!this.ws) {
      throw new Error('WebSocket connection is not established');
    }
    this.ws.send(message);
  }
}

module.exports = Client;