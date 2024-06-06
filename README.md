# gotoamsterdam

Odpalamy gan-react-server poleceniem 'npm run'

Jest to stronka dzięki której nawiążemy połączenie z kostką za pomocą Bluetooth (tutaj trzeba raz ręcznie zainicjalizować przyciskiem w prawym górnym rogu)

## Websockety
Komunikacja opiera się o websockety.
Aby wysyałć komendy do stronki i odbierać polecenia od użytkowników gry, musimy postawić Websocket server.

### Server
W tym przypadku jest to prosty NodeJs serwer.
Odpalamy go na maszynie która będzie łączyć się z robotem poleceniem 'node ./server/server.js'

### Clients
These are 2 clients, both sitting on each player's machine.
You should start them by running 'dotnet run ./clients/dotnet' and 'python ./clients/python-client.py