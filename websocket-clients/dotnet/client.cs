using System.Net.WebSockets;
using System.Text;
using System.Text.Json;

class Program
{
    static async Task Main(string[] args)
    {
        Uri serverUri = new Uri("ws://localhost:3000");

        using var ws = new ClientWebSocket();
        await ws.ConnectAsync(serverUri, CancellationToken.None);
        Console.WriteLine("Connected to WebSocket server at ws://localhost:3000");

        // Send a message to the WebSocket server
        var messageObject = new { payload = "sendInstruction" };
        string messageToSend = JsonSerializer.Serialize(messageObject);
        await SendMessageAsync(ws, messageToSend);

        // Receive a message from the WebSocket server
        string messageReceived = await ReceiveMessageAsync(ws);
        Console.WriteLine($"Message received: {messageReceived}");

        // Close the WebSocket connection
        await ws.CloseAsync(WebSocketCloseStatus.NormalClosure, "Closing", CancellationToken.None);
        Console.WriteLine("WebSocket connection closed");
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