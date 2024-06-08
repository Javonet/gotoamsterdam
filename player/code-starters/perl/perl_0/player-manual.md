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
Javonet->activate("your-license-key");
```

### Javonet needs to be imported as any other dependency.
```perl
use aliased 'Javonet::Javonet' => 'Javonet';
```

### Creating Python context
In order to run Python modules, we need to create it's context.
You can do that by invoking this in memory (*Tip: we also offer remote connections!)

```perl
Javonet->in_memory()->python();
```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```perl
  my $python_runtime = Javonet->in_memory()->python();
  ```
</details>

### Load python module to your app
You can load a custom library by calling:
  ```perl
  $python_runtime->load_library($library_path);
  ```

<details>
  <summary>Help me</summary>
  
  ### Code
  ```perl
  $python_runtime->load_library(".");
  ```
</details>

### Access Robot Class
You now need to get that Class from loaded module
  ```perl
  my $class_name = "FileName.ClassName";
  $python_runtime->get_type($class_name)->execute();
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```perl
  my $class_name = "robot-connector.Robot";
  my $python_type = $python_runtime->get_type($class_name)->execute();
  ```
</details>

### Invoke Solve Method

  ```perl
  $python_type->invoke_static_method(methodName)->execute();
  ```
<details>
  <summary>Help me</summary>
  
  ### Code
  ```perl
  c$python_type->invoke_static_method("solve")->execute();
  ```
</details>

### Run your perl code from Terminal


<details>
  <summary>Help me</summary>
  
  ### Code
  ```bash
  perl ./main.pl
  ```
</details>