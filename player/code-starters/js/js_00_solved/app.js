const { Javonet } = require('javonet-nodejs-sdk/lib/sdk/Javonet')

Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA")

let calledRuntime = Javonet.inMemory().python()
const className = "robot-connector.Robot"

calledRuntime.loadLibrary('.')
let calledRuntimeType = calledRuntime.getType(className).execute()

calledRuntimeType.invokeStaticMethod("solve").execute()