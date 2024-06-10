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
Javonet.Activate("your-license-key");
```

### Javonet needs to be imported as any other dependency.
```c#
using Javonet.Netcore.Sdk; 

//or for .NET Framework applications 
using Javonet.Clr.Sdk for .NET Framework apps
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

```c#
Javonet.InMemory().Python();
```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```c#
  var calledRuntime = Javonet.InMemory().Python();
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  ```c#
  calledRuntime.LoadLibrary(pythonModuleFolder);
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```c#
  calledRuntime.LoadLibrary("PythonRobotModule");
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```c#
  var calledRuntimeType = calledRuntime.GetType("robot-connector.Robot").Execute();
  ```

### Invoke Solve Method

  ```c#
  var response = calledRuntimeType.InvokeStaticMethod("solve").Execute();
  ```