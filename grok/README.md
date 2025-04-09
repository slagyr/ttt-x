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

### Attempt #2

Given the prompt:

    Hmm...  There are quite a few differences between your implementation and the original.  One is the format of the output.  Second, and this is a big one, in the original the computer plays against the human.  How did you miss this?

Grok added a computer player but it is super simplistic, taking the first move available.  The output remains different.


### Reset Attempt using "Think"

Prompt:
```
I would like your help to convert a relatively small repository of Clojure code to Golang.  
Please preserve all functionality of the code.  Do make any assumptions about the behavior of the code.  
There are speclj tests that need to be converted as well.  
Here is the repo link: https://github.com/arootroatch/tic-tac-toe-clojure/tree/5a339a00bc23ca8dd9b4c14501a26dd4b4df8711
```

Result (Thinking for 2:12):

* The resulting code has much more structure than the first attempts, however it differs from the original.
* There are a couple compile errors.  Fixable within a few min.
* Tests pass
* The UI is different!!! It doesn't even print the board!!!
* The computer beat me, mostly because I couldn't see the board.
* Minimax seems to work on first glance

There is not a direct correlation between Go code and Clojure code.  Grok took a much closer look at the Clojure 
implementation, but it did not do direct conversion of the code.  It seems to have formulated a summary of the behavior
and then recreated it from scratch, taking many creative liberties.  This is NOT an exact translation.


