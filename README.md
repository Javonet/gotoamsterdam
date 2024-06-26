# Javonet - Speed of Integration (GoTo Amsterdam)

## Install Environments on Player's Machine
Java - https://www.oracle.com/java/technologies/downloads/

## One Script Run (ToDo)
To run everything on a user machine, you just need to run `player.bat`.  
To run on the server side, you just need to run `server.bat`.

## Just Scripts
On player machines, change the IP address in the environment variable to the target server.
For dotnet, you have to Configure NuGet to use Local Source with fixed Javonet
`dotnet nuget add source .\player\nuget-source -n LocalSource`

## React PWA

Run with the command `npm run ./gan-react-server`.

This is a page through which we will establish a connection with the device via Bluetooth (here you need to manually initialize it once by pressing the button in the top right corner).

To run this, the node version has to be 14.21.3 (it won't work with newer versions).
The best way is to install nvm (Node Version Manager) and setup this for gan-react-server folder.
`nvm install 14.21.3`
``

## Websockets
Communication is based on websockets.  
To send commands to the page and receive commands from game users, we need to set up a Websocket server.

### Websocket Server
In this case, it's a simple Node.js server.  
Run it on the machine that will connect to the robot with the command 
`npm install ws`
`node ./server/websocket-server/server.js`

### Websocket Clients
These are 2 websocket clients, both running on each player's machine.  
Start them by running `dotnet run ./clients/dotnet` and `python ./clients/python-client.py`.

## GUI UI
The starting point for each player is a simple GUI written in Python.  
Start this by running `python ./gui/gui-main.py`.

GUI was written in Python using Figma and Tkinter Designer which exports Figma to Python code.
https://github.com/ParthJadhav/Tkinter-Designer
https://www.figma.com/design/FePKixJUTvrp16nVWjVhSp/Javonet-Speed-of-Integration-challenge?node-id=0-1&t=VOtIb4waSXcqttce-1

A significant upgrade could be achieved by using Unity for a nicer feel and look that is closer to game style.

## Code Starter
When the GUI is fired, each player can choose their preferred technology.  
When they choose one, we will run a `.bat` file located in a folder dedicated to each programming language under `code-starters`.

Each folder will have a subfolder ending with "_0", e.g., `code-starters/dotnet/dotnet_0`.  
Every time a player chooses a specific programming language, the `.bat` file will create a copy of that folder, increment the suffix, and call `code .` to open it in VS Code.

### 0 Folder
This folder should contain a starting point for the given language and the function to activate the robot.

In each "0" folder, you will find a `player-manual.md` file. It should be rendered by pressing `Ctrl+Shift+V`.
