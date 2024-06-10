# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Install Javonet NodeJs package
```java
mvn dependency:get -Dartifact='com.javonet:javonet-java-sdk:2.4.3'
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```java
Javonet.activate("your-license-key");
```

### Javonet needs to be imported as any other dependency.
```java
import com.javonet.sdk.*;
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

```java
Javonet.inMemory().python();
```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```java
  RuntimeContext pythonRuntime = Javonet.inMemory().python();
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  ```java
  calledRuntime.loadLibrary(resourceDirectory)
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```java
  calledRuntime.loadLibrary(".");
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```java
  const className = "FileName.ClassName"
  called_runtime.get_type(class_name).execute()
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```java
  const className = "robot-connector.Robot"
  InvocationContext calledRuntimeType = calledRuntime.getType(className).execute();
  ```
</details>

### Invoke Solve Method

  ```java
  calledRuntimeType.invokeStaticMethod(methodName).execute();
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```java
  calledRuntimeType.invokeStaticMethod("solve").execute();
  ```
</details>

### Run your Java code from Terminal


<details>
  <summary>Help me</summary>
  
  ### Code
  ```bash
  java ./Main.java
  ```
</details>