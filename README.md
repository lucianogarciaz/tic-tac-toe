# Tic Tac Toe
The rules of the tic tac toe game are the following:

* a game is over when all fields are taken
* a game is over when all fields in a column are taken by a player
* a game is over when all fields in a row are taken by a player
* a game is over when all fields in a diagonal are taken by a player
* a player can take a field if not already taken
* players take turns taking fields until the game is over
* X player always goes first
* there are two player in the game (X and O)

Initial state: having an empty 3x3 board, a flag "open" that indicates it, and two players: X and O.

Given a game that is open
When a player takes a field
Then the field is taken
