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
Javonet.activate("your-license-key")
```

### Javonet needs to be imported as any other dependency.
```javascript
const { Javonet } = require('javonet-nodejs-sdk/lib/sdk/Javonet')
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

```javascript
Javonet.inMemory().python()
```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```javascript
  let pythonRuntime = Javonet.inMemory().python()
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  ```javascript
  calledRuntime.loadLibrary(resourceDirectory)
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```javascript
  resourceDirectory = "./"
  calledRuntime.loadLibrary(resourceDirectory)
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```javascript
  const className = "FileName.ClassName"
  called_runtime.get_type(class_name).execute()
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```javascript
  const className = "robot-connector.Robot"
  let calledRuntimeType = calledRuntime.getType(className).execute()
  ```
</details>

### Invoke Solve Method

  ```javascript
  called_runtime_type.invoke_static_method(method_name).execute()
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```javascript
  called_runtime_type.invoke_static_method("solve").execute()
  ```
</details>

### Run your JavaScript code from Terminal
Now you just need to run your app with NodeJs

<details>
  <summary>Help me</summary>
  
  ### Code
  ```bash
  node ./app.js
  ```
</details>