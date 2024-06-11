# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

<!-- ### Install Javonet NodeJs package
```cpp
gem install javonet-cpp-sdk
``` -->

### Javonet needs to be imported as any other dependency.
```cpp
#include "Javonet.h"
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```cpp
Javonet::Activate("p5XB-z7MN-Tp9a-d3NH-y4GA");
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

  ### Code
  ```cpp
  auto pythonRuntime = Javonet::InMemory()->Python();
  ```


### Load python module to your app
You can load a custom library by calling:

  ### Code
  ```cpp
  calledRuntime->LoadLibrary(".");
  ```


### Access Robot Class
You now need to get that Class from loaded module

  ### Code
  ```cpp
  auto className = "TestClass.TestClass";
  auto calledRuntimeType = calledRuntime->GetType(className)->Execute();
  ```

### Invoke Solve Method


  ### Code
  ```cpp
  calledRuntimeType->InvokeStaticMethod("solve")->Execute();
  ```

### Run your cpp code from Terminal


  ### Code
  ```bash
  cpp ./main.pl
  ```
