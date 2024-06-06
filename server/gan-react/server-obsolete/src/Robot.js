const { broadcastMessage } = require('./websocket');

class Robot{
    static solve() {
        if (broadcastMessage) {
            broadcastMessage(JSON.stringify({payload: "sendInstruction"}));
            console.log("Instruction sent");
        } else {
            console.log("Error");
            return;
        }
    };
}

module.exports = Robot