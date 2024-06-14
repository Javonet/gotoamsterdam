using System;
using System.Net.Http.Json;
using System.Text.Json;
using System.Text.Json.Serialization;
using Javonet.Netcore.Sdk;

namespace JavonetSpeedOfIntegration
{

    class MainClass 
    {
        private static HttpClient httpClient = new HttpClient();
        private static Javonet.Netcore.Sdk.RuntimeContext calledRuntime=null;
        private static InvocationContext calledRuntimeType;
        static void Main(string[] args)
        {
            Javonet.Netcore.Sdk.Javonet.Activate("p5XB-z7MN-Tp9a-d3NH-y4GA");

            calledRuntime = Javonet.Netcore.Sdk.Javonet.Tcp(new Javonet.Netcore.Utils.TcpConnectionData("127.0.0.1",8080)).Netcore();
            calledRuntime.LoadLibrary(".\\bin\\Debug\\net8.0\\RobotConnector.dll");
            calledRuntimeType = calledRuntime.GetType("robotConnector.Robot").Execute();

            //Can you solve the Rubik's Cube?
            
            //This is your app
            //1. Open player-manual.md
            //2. Press Ctrl+K,V to get MD file formatted
            //3. Let's try, follow guide and win!

            Console.WriteLine("Let's Start!");


            JavonetApproach();
            WebApiApproach();

        }

    public static void JavonetApproach() {
            //This will be auto generated into strongly typed interface like Robot.Compute()
            var javonetResult = (int)calledRuntimeType.InvokeStaticMethod("Compute",4,5).Execute().GetValue();
    }

    class ResultRecord {
        public int result;
    }
    public static void WebApiApproach() {
          // using ()
          //  {
                var response = httpClient
                    .GetAsync("http://localhost:5229/compute").GetAwaiter().GetResult();
                var result = response.Content.ReadAsStringAsync().GetAwaiter().GetResult();
                var r =JsonSerializer.Deserialize<ResultRecord>(result);
                // Do stuff...
                //Console.WriteLine("Web API Result: "+result);
          //  }
    }
    }

}