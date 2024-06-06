const WebSocket = require('ws');

let wss;

const setupWebSocket = (server) => {
    wss = new WebSocket.Server({ server });

    wss.on('connection', (ws) => {
        console.log('Client connected');

        ws.on('message', (message) => {
            console.log('Received:', message);
            ws.send(`Echo: ${message}`);
        });

        ws.on('close', () => {
            console.log('Client disconnected');
        });
    });
};

const broadcastMessage = (message) => {
    if (!wss) {
        console.error('WebSocket server is not initialized');
        return;
    }
    wss.clients.forEach((client) => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(message);
        }
    });
};

module.exports = { setupWebSocket, broadcastMessage };
