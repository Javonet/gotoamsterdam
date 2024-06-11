# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

<!-- ### Install Javonet NodeJs package
```go
gem install javonet-go-sdk
``` -->

### Javonet needs to be imported as any other dependency.
```go
import Javonet "javonet.com/javonet"
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```go
Javonet.ActivateWithCredentials("p5XB-z7MN-Tp9a-d3NH-y4GA")
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

```go
Javonet.InMemory().Python()
```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```go
  calledRuntime, _ := Javonet.InMemory().Python()
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  ```go
  calledRuntime.LoadLibrary(libraryPath)
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```go
  alledRuntime.LoadLibrary(".")
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```go
  auto className = "TestClass.TestClass";
  calledRuntime->GetType(className)->Execute();
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```go
  className := "Robot.Robot"
  calledRuntimeType, _ := calledRuntime.GetType(className).Execute()
  ```
</details>

### Invoke Solve Method

  ```go
  calledRuntimeType.InvokeStaticMethod(methodName).Execute()
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```go
  calledRuntimeType.InvokeStaticMethod("solve").Execute()
  ```
</details>

### Run your go code from Terminal


<details>
  <summary>Help me</summary>
  
  ### Code
  ```bash
  go ./main.pl
  ```
</details>