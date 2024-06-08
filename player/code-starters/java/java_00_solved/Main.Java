import com.javonet.sdk.*;

public class Main {
    public static void main(String []args) {
        Javonet.activate("p5XB-z7MN-Tp9a-d3NH-y4GA");
        System.out.println("Hello, World!");

        //------------------------------------------------
        //---------------Python---------------------------
        //------------------------------------------------
        RuntimeContext calledRuntime = Javonet.inMemory().python();
        calledRuntime.loadLibrary(".");
        InvocationContext calledRuntimeType = calledRuntime.getType("robot-connector.Robot").execute();
        calledRuntimeType.invokeStaticMethod("solve").execute();

        System.out.println("Done");

        //-------------------------------------------------
        //---------------NetCore---------------------------
        //-------------------------------------------------

        // RuntimeContext calledRuntime = Javonet.inMemory().netcore();
        // calledRuntime.loadLibrary("./RobotConnector.dll");
        // InvocationContext calledRuntimeType = calledRuntime.getType("Robot").execute();

        // calledRuntimeType.invokeStaticMethod("Solve").execute();

        // System.out.println("Done");
    }
}