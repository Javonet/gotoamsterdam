# gotoamsterdam

## One Script Run (ToDo)
Aby odpalić wszystko on a user machrine, you just need to run 'player.bat'
To run on the server side, you just need to run 'server.bat'

## Just scripts
Na maszynach playerów, Zmień adres IP w env variable na serwer docelowy

## React PWA
Odpalamy  poleceniem 'npm run ./gan-react-server'

Jest to stronka dzięki której nawiążemy połączenie z kostką za pomocą Bluetooth (tutaj trzeba raz ręcznie zainicjalizować przyciskiem w prawym górnym rogu)

## Websockety
Komunikacja opiera się o websockety.
Aby wysyałć komendy do stronki i odbierać polecenia od użytkowników gry, musimy postawić Websocket server.

### Websocket Server
W tym przypadku jest to prosty NodeJs serwer.
Odpalamy go na maszynie która będzie łączyć się z robotem poleceniem 'node ./server/server.js'

### Websocket Clients
These are 2 websocket clients, both sitting on each player's machine.
You should start them by running 'dotnet run ./clients/dotnet' and 'python ./clients/python-client.py

## GUI UI
The starting point for each player is simple GUI written in python.
We start this by running 'python ./gui/gui-main.py'

## Code Starter
When GUI is fired, each player can choose their preffered technology.
When they choose one, we will run a .bat file, that is located in a folder dedicated to each programming language under 'code-starters'

Each folder will have first folder ending with "_0" i.e. 'code-starters/dotnet/dotnet_0'.
Every time player chooses a specific programming language, the .bat file will create a copy of that folder, will increment the suffix and call 'code .' opennning VS Code.


### 0 Folder
It's a folder that should contain a starting point for given language and activating robot function to call.

In each 0 folder you will find a player-manual.md file. It should be rendered by pressing "Ctrl+Shift+V"

P


