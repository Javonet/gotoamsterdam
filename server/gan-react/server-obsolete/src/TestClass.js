const { broadcastMessage } = require('./websocket');

class TestClass{
    static multiplyByTwo(a) {
        return 2 * a;
    }

    multiplyTwoNumbers(a, b) {
        return a * b
    }
}

module.exports = TestClass