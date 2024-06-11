# How to play

Your aim is to run a code written in Python.
There is a static method `solve` that is in the `Robot` class in `robot-connector.py` file.

You license key is: "p5XB-z7MN-Tp9a-d3NH-y4GA"

### Install Javonet NodeJs package
```perl
cpanm Javonet::Javonet
```

### Javonet license key activation
Javonet needs to be activated first. Activation must be called only once at the start-up of an application.

```perl
Javonet->activate("p5XB-z7MN-Tp9a-d3NH-y4GA");
```

### Javonet needs to be imported as any other dependency.
```perl
use aliased 'Javonet::Javonet' => 'Javonet';
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

  ### Code
  ```perl
  my $python_runtime = Javonet->in_memory()->python();
  ```


### Load python module to your app
You can load a custom library by calling:

  ### Code
  ```perl
  $python_runtime->load_library(".");
  ```


### Access Robot Class
You now need to get that Class from loaded module
  
  ### Code
  ```perl
  my $class_name = "robot-connector.Robot";
  my $python_type = $python_runtime->get_type($class_name)->execute();
  ```

### Invoke Solve Method
  
  ### Code
  ```perl
  $python_type->invoke_static_method("solve")->execute();
  ```

### Run your perl code from Terminal

  ### Code
  ```bash
  perl ./main.pl
  ```