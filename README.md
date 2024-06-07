# Javonet - Speed of Integration (GoTo Amsterdam)

## One Script Run (ToDo)
To run everything on a user machine, you just need to run `player.bat`.  
To run on the server side, you just need to run `server.bat`.

## Just Scripts
On player machines, change the IP address in the environment variable to the target server.

## React PWA
Run with the command `npm run ./gan-react-server`.

This is a page through which we will establish a connection with the device via Bluetooth (here you need to manually initialize it once by pressing the button in the top right corner).

## Websockets
Communication is based on websockets.  
To send commands to the page and receive commands from game users, we need to set up a Websocket server.

### Websocket Server
In this case, it's a simple Node.js server.  
Run it on the machine that will connect to the robot with the command `node ./server/server.js`.

### Websocket Clients
These are 2 websocket clients, both running on each player's machine.  
Start them by running `dotnet run ./clients/dotnet` and `python ./clients/python-client.py`.

## GUI UI
The starting point for each player is a simple GUI written in Python.  
Start this by running `python ./gui/gui-main.py`.

## Code Starter
When the GUI is fired, each player can choose their preferred technology.  
When they choose one, we will run a `.bat` file located in a folder dedicated to each programming language under `code-starters`.

Each folder will have a subfolder ending with "_0", e.g., `code-starters/dotnet/dotnet_0`.  
Every time a player chooses a specific programming language, the `.bat` file will create a copy of that folder, increment the suffix, and call `code .` to open it in VS Code.

### 0 Folder
This folder should contain a starting point for the given language and the function to activate the robot.

In each "0" folder, you will find a `player-manual.md` file. It should be rendered by pressing `Ctrl+Shift+V`.
