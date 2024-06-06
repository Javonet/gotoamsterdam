const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 3000 });

// Function to broadcast a message to all connected clients
function broadcastMessage(message) {
  wss.clients.forEach(function each(client) {
    if (client.readyState === WebSocket.OPEN) {
      client.send(message);
    }
  });
}

wss.on('connection', function connection(ws) {
  console.log('A new client connected');

  ws.on('message', function incoming(message) {
    console.log('received: %s', message);
    
    // Broadcast the received message to all connected clients
    robotMessage = JSON.stringify({ payload: 'sendInstruction'})
    broadcastMessage(message);
  });

  ws.send(JSON.stringify({ payload: "You've solved Rubik's Cube" }));
  //'Welcome to the WebSocket server!');
});

console.log('WebSocket server is running on ws://localhost:3000');
