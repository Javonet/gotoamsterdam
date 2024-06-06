//This works by executing node .\example.js
const client = require('./client');

const clientInstance = new client();
clientInstance.connect()
  .then(() => {
    clientInstance.send(JSON.stringify({payload: "sendInstruction"}));
  })
  .catch((err) => {
    console.error('Error connecting to WebSocket server:', err);
  });
