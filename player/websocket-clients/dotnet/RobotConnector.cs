using System.Net.WebSockets;
using System.Text;
using System.Text.Json;

class Robot
{
    public static void Solve(){
        var task = SolveAsync();
        task.Wait();
    }

    public static async Task SolveAsync(){
        //Use env var?
        var ipAddress = "192.168.68.105";
        Uri serverUri = new Uri($"ws://{ipAddress}:3000");

        using var ws = new ClientWebSocket();
        await ws.ConnectAsync(serverUri, CancellationToken.None);
        Console.WriteLine($"Connected to Robot at {DateTime.Now}");

        // Send a message to the Robot
        await SendMessageAsync(ws, Environment.MachineName);

        // Receive a message from the Robot
        string messageReceived = await ReceiveMessageAsync(ws);
        //Console.WriteLine($"Message received: {messageReceived}");

        // Close the WebSocket connection
        await ws.CloseAsync(WebSocketCloseStatus.NormalClosure, "Closing", CancellationToken.None);
        Console.WriteLine("Connection closed");
    }
    private static async Task SendMessageAsync(ClientWebSocket ws, string message)
    {
        byte[] bytesToSend = Encoding.UTF8.GetBytes(message);
        var buffer = new ArraySegment<byte>(bytesToSend);
        await ws.SendAsync(buffer, WebSocketMessageType.Text, true, CancellationToken.None);
        Console.WriteLine($"Message sent: {message}");
    }
    private static async Task<string> ReceiveMessageAsync(ClientWebSocket ws)
    {
        var buffer = new byte[1024];
        var result = await ws.ReceiveAsync(new ArraySegment<byte>(buffer), CancellationToken.None);
        return Encoding.UTF8.GetString(buffer, 0, result.Count);
    }
}