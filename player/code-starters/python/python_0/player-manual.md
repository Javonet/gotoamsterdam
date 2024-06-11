# How to play

Your aim is to run static method solve that is in the Robot class written in C#.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Add Javonet Nuget package
```python
python -m pip install javonet-python-sdk
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```python
Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA")
```

### Javonet needs to be imported as any other dependency.
```python
from javonet.sdk import Javonet
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)


  
  ### Code
  ```python
  netcore_runtime = Javonet.in_memory().netcore()
  ```


### Load python module to your app
You can load a custom library by calling:
  
  ### Code
  ```python
  library_path = "./RobotConnector.dll"
  called_runtime.load_library(library_path)
  ```


### Access Robot Class
You now need to get that Class from loaded module
  
  ### Code
  ```python
  called_runtime_type = called_runtime.get_type("Robot").execute()
  ```

### Invoke Solve Method

  
  ### Code
  ```python
  called_runtime_type.invoke_static_method("Solve").execute()
  ```
### Run python code

  ### Code
  ```bash
    python main.py
  ```