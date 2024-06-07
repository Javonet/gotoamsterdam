using System;

namespace JavonetSpeedOfIntegration
{
    using Javonet.Netcore.Sdk;
    //using Javonet.Netcore.Sdk;
    class JavonetSpeedOfIntegration{
        static void Main(string[] args)
        {
            Javonet.Activate("p5XB-z7MN-Tp9a-d3NH-y4GA");

            var calledRuntime = Javonet.InMemory().Python();

            var libraryPath = "C:/Local/GoToAmsterdam/player/code-starters/dotnet/dotnet_00_solved";
            calledRuntime.LoadLibrary(libraryPath);

            var calledRuntimeType = calledRuntime.GetType("robot-connector.Robot").Execute();
            calledRuntimeType.InvokeStaticMethod("solve").Execute();

            Console.WriteLine("Done");
        }
    }
}