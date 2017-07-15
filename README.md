### Usage:
* run main.go to start the server, this will register a battle (only 1 battle can run for now)
* issue commands by POST body, for example:
```
curl -X POST -d '{"player":"me","command":{"source":0,"target":1,"move":0}}' "localhost:8080/move"
```
* once 2 commands are issued, the turn is evaluated and both players receive turn resolution

### Status:
Current data is Gen V (from http://pokemonessentials.wikia.com/wiki)
GenVI is TODO