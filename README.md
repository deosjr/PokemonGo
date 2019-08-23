### Usage:
* main.go runs singleplayer by default, this will let you test a 1v1 battle on the command line.

* run main.go with the -serve flag to start as server
* multiple games can be created by calling /game, join a game with /join
* issue commands by POST body, for example:
```
curl -X POST -d '{"game":"id","player":"me","command":{"source":0,"target":1,"move":0}}' "localhost:8080/move"
```
* once 2 commands are issued, the turn is evaluated and both players receive turn resolution

### Status:
Current data is Gen V (from http://pokemonessentials.wikia.com/wiki)
GenVI is TODO
