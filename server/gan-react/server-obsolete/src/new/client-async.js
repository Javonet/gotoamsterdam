const WebSocket = require('ws');

class client {
  constructor() {
    this.ws = null;
  }

  connect() {
    return new Promise((resolve, reject) => {
      const ws = new WebSocket('ws://localhost:3000');

      ws.on('open', () => {
        console.log('Connected to WebSocket server');
        this.ws = ws;
        resolve();
      });

      ws.on('error', (err) => {
        console.error('Failed to connect to WebSocket server:', err);
        reject(err);
      });
    });
  }

  send(message) {
    if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
      console.error('WebSocket connection is not open');
      return;
    }

    this.ws.send(message);
  }
}

module.exports = client;
