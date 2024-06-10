# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Install Javonet NodeJs package
```ruby
gem install javonet-ruby-sdk
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```ruby
Javonet.activate('your-license-key')
```

### Javonet needs to be imported as any other dependency.
```ruby
Javonet.activate('p5XB-z7MN-Tp9a-d3NH-y4GA')
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

```ruby
Javonet.in_memory.python
```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```ruby
  called_runtime = Javonet.in_memory.python
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  ```ruby
  called_runtime.load_library(library_path)
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```ruby
  called_runtime.load_library(".")
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```ruby
  class_name = 'FileName.ClassName'
  called_runtime.get_type(class_name).execute
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```ruby
  class_name = 'robot-connector.Robot'
  called_runtime_type = called_runtime.get_type(class_name).execute
  ```
</details>

### Invoke Solve Method

  ```ruby
  called_runtime_type.invoke_static_method(methodName).execute
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```ruby
  called_runtime_type.invoke_static_method('solve').execute
  ```
</details>

### Run your ruby code from Terminal


<details>
  <summary>Help me</summary>
  
  ### Code
  ```bash
  ruby ./main.pl
  ```
</details>