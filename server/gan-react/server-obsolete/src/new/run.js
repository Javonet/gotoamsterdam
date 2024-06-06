const Example = require('./example');
Example.connect();
setTimeout(function() {
    Example.run();
}, 1000);
