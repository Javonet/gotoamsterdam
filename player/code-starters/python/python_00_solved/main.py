from javonet.sdk import Javonet

def main():
    Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA")
    called_runtime = Javonet.in_memory().netcore()

    library_path = "./RobotConnector.dll"
    called_runtime.load_library(library_path)

    called_runtime_type = called_runtime.get_type("Robot").execute()
    called_runtime_type.invoke_static_method("Solve").execute()

if __name__ == "__main__":
    main()