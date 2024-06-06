const client = require('./client');
const clientInstance = new client();

class Example {
    static connect(){
        clientInstance.connect();
    }

    static run() {
        console.log("example.run started");
        try{
            //setTimeout(function() {
                console.log("inside then");
                clientInstance.send(JSON.stringify({ payload: "sendInstruction" }));
              //}, 1000);
            //clientInstance.connect();
            
            }
            catch(err) {
                console.error('Error connecting to WebSocket server:', err);
            };
        console.log("example.run finished")
    }
}

module.exports = Example;
