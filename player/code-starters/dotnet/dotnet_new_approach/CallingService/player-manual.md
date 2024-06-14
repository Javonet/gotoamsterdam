# How to play

Your aim is to run static method solve that is in the Robot class written in Python.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Add Javonet Nuget package
```c#
dotnet add package Javonet.Netcore.Sdk -s https://api.nuget.org/v3/index.json

//or for .NET Framework applications
dotnet add package Javonet.Clr.Sdk -s https://api.nuget.org/v3/index.json
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```c#
Javonet.Netcore.Sdk.Javonet.Activate("p5XB-z7MN-Tp9a-d3NH-y4GA");
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)


  ### Code
  ```c#
  var calledRuntime = Javonet.Netcore.Sdk.Javonet.InMemory().Python();
  ```

### Load python module to your app
You can load a custom library by calling:

  ### Code
  ```c#
  calledRuntime.LoadLibrary("PythonRobotModule");
  ```

### Access Robot Class
You now need to get that Class from loaded module
  ```c#
  var calledRuntimeType = calledRuntime.GetType("robot-connector.Robot").Execute();
  ```

### Invoke Solve Method

  ```c#
  var response = calledRuntimeType.InvokeStaticMethod("solve").Execute();
  ```