# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Add require for Javonet
```
require 'javonet-ruby-sdk'
```

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

  ### Code
  ```ruby
  called_runtime = Javonet.in_memory.python
  ```

### Load python module to your app
You can load a custom library by calling:

  ### Code
  ```ruby
  called_runtime.load_library(".")
  ```

### Access Robot Class
You now need to get that Class from loaded module

  ### Code
  ```ruby
  class_name = 'robot-connector.Robot'
  robotClass = called_runtime.get_type(class_name).execute
  ```

### Invoke Solve Method

  ### Code
  ```ruby
  robotClass.invoke_static_method('solve').execute
  ```

### Run your ruby code from Terminal

  ### Code
  ```bash
  ruby ./main.rb
  ```
