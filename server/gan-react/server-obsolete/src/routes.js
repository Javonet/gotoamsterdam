const express = require('express');
const { broadcastMessage } = require('./websocket');
const router = express.Router();
const robot = require("./robot");

router.post('/send-message', (req, res) => {
    
    const { message } = req.body;
    if(robot.solve()){
        res.status(200).send({ status: 'Message sent to all clients' });
    } else {
        res.status(500).send({ status: 'WebSocket server not initialized' });
    }
});

module.exports = router;
