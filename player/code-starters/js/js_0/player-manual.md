# How to play

Your aim is to run static method solve that is in the Robot class written in Python.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Install Javonet NodeJs package
```javascript
npm i javonet-nodejs-sdk
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```javascript
Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA")
```

### Javonet needs to be imported as any other dependency.
```javascript
const { Javonet } = require('javonet-nodejs-sdk/lib/sdk/Javonet')
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

  ### Code
  ```javascript
  let pythonRuntime = Javonet.inMemory().python()
  ```


### Load python module to your app
You can load a custom library by calling:

  ### Code
  ```javascript
  resourceDirectory = "./"
  pythonRuntime.loadLibrary(resourceDirectory)
  ```

  
  ### Code
  ```javascript
  const className = "robot-connector.Robot"
  let robotClass = calledRuntime.getType(className).execute()
  ```


### Invoke Solve Method

  ### Code
  ```javascript
  robotClass.invoke_static_method("solve").execute()
  ```

### Run your JavaScript code from Terminal


  ### Code
  ```bash
  node ./app.js
  ```
