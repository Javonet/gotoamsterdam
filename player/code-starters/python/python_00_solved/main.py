from javonet.sdk import Javonet
import time

def main():
    print("Python started")

    Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA")
    called_runtime = Javonet.in_memory().nodejs()

    library_path = "C:\\Local\\gan-scrambler-master\\server\\src\\new\\example.js"
    class_name = 'example'

    called_runtime.load_library(library_path)
    called_runtime_type = called_runtime.get_type(class_name).execute()
    response = called_runtime_type.invoke_static_method("connect").execute()
    time.sleep(3)
    
    # response = called_runtime_type.invoke_static_method("run").execute()
    time.sleep(5)

    print("Python done")

if __name__ == "__main__":
    main()