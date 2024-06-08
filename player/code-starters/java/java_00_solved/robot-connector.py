import asyncio
import websockets
from datetime import datetime
import socket

class Robot:
    def __init__(self):
        pass

    async def connect(self, ip_address: str, port: int):
        uri = f"ws://{ip_address}:{port}"
        async with websockets.connect(uri) as websocket:
            hostname = socket.gethostname()
            print(f"Connected to Robot at {datetime.now()}")

            # Send Message
            await websocket.send(hostname)
            print("Command sent")

    @staticmethod
    def solve():
        print(f"Connected to Robot at {datetime.now()}")
        ip_address = "192.168.68.105"
        port = 3000
        robot = Robot()
        asyncio.get_event_loop().run_until_complete(robot.connect(ip_address, port))
