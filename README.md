# TTT in Golang

Translated by grok 3 from Alex Root-Roatch's implementation: https://github.com/arootroatch/tic-tac-toe-clojure

## Basic Functionality

https://github.com/arootroatch/tic-tac-toe-clojure/tree/5a339a00bc23ca8dd9b4c14501a26dd4b4df8711

I asked Grok to translate the repo.  The result was quite different:
1. The way it printed the board:
```
Current board:
|   |
---------
|   |
---------
|   |  
Player x, enter position (0-8): 4

vs

1 2 3
4 5 6
7 8 9
Please enter your move (type 1-9 and hit enter):
3
```
2. The generated version was human vs human, whereas the original was human vs computer.  Big miss.
3. It assumed the use of `clojure.test` instead of notice the elaborate suite of `speclj` tests. 

It was almost like the AI tried to infer the intent of the original implementation without looking closely at the code.