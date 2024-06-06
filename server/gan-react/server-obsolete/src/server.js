const http = require('http');
const app = require('./app');
const { setupWebSocket } = require('./websocket');

const PORT = process.env.PORT || 3000;
const server = http.createServer(app);

// Setup WebSocket and attach it to the server
setupWebSocket(server);

server.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
