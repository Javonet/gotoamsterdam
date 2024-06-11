# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Install Javonet Java package in project pom.xml
```java
<dependency>
    <groupId>com.javonet</groupId>
    <artifactId>javonet-java-sdk</artifactId>
    <version>2.4.3</version>
</dependency>
```

Then run in terminal
```
mvn clean install
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```java
Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA");
```

### Javonet needs to be imported as any other dependency.
```java
import com.javonet.sdk.*;
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

  
  ### Code
  ```java
  RuntimeContext pythonRuntime = Javonet.inMemory().python();
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  
  ### Code
  ```java
  pythonRuntime.loadLibrary(".");
  ```


### Access Robot Class
You now need to get that Class from loaded module

  ### Code
  ```java
  String className = "robot-connector.Robot";
  InvocationContext pythonRuntimeType = pythonRuntime.getType(className).execute();
  ```


### Invoke Solve Method


  ### Code
  ```java
  pythonRuntimeType.invokeStaticMethod("solve").execute();
  ```


### Run your Java code from Terminal


  ### Code
  ```bash
  mvn exec:java
  ```
