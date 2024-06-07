import asyncio
import websockets
from datetime import datetime
import socket

class Robot:
    def solve(self):
        async def connect_to_websocket(ip_address: str, port: int):
            uri = f"ws://{ip_address}:{port}"
            async with websockets.connect(uri) as websocket:
                hostname = socket.gethostname()
                print(f"Connected to Robot at {datetime.now()}")

                # Example of sending a message
                await websocket.send(hostname)
                print("Command sent")

        # Replace with the desired IP address and port
        ip_address = "192.168.68.105"
        port = 3000

        asyncio.get_event_loop().run_until_complete(connect_to_websocket(ip_address, port))


