# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Install Javonet GoLang package

Javonet is already added to your project move to next step!

### Javonet needs to be imported as any other dependency.
```go
Javonet "javonet.com/javonet"
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
calledRuntime, _ := Javonet.InMemory().Python()
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
  calledRuntime.LoadLibrary(".")
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```go
  calledRuntime.LoadLibrary(".")
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```go
  calledRuntimeType, _ := calledRuntime.GetType("PythonFileName.Class").Execute()
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```go
  calledRuntimeType, _ := calledRuntime.GetType("robot-connector.Robot").Execute()
  ```
</details>

### Invoke Solve Method

  ```cpp
  calledRuntimeType.InvokeStaticMethod(MethodName).Execute()
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```cpp
  calledRuntimeType.InvokeStaticMethod("solve").Execute()
  ```
</details>

### Run your cpp code from Terminal


<details>
  <summary>Help me</summary>
  
  ### Code
  ```bash
  go get
  go run ./main.go - runs an app
  ```
</details>