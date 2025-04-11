# Translate a repo from one language (Clojure) to another (Golang).

This is attempted with multiple AI's, each with it's own directory:

* Grok
* Claude
* ChatGPT

## The Project - Tic Tac Toe

This repo was created by Alex Root-Roach during his Clean Coders Studio apprenticeship:

https://github.com/arootroatch/tic-tac-toe-clojure/tree

### Release 1 

The first release of TTT has a console interface where the human can play against an unbeatable computer player.

https://github.com/arootroatch/tic-tac-toe-clojure/tree/5a339a00bc23ca8dd9b4c14501a26dd4b4df8711

Talking notes:
 * Statistical model
 * Making assumptions
 * Piece by piece looses connectivity
 * Claude console
 * Simple implementation, with elaborate prompt, surprising good.
 * Layers of abstraction, needs hand-holding
 * Handles feedback well, but needs lots of feedback
 * Novelties are not handled well
 * Credits, ran ouy, daily limit, Enterprise edition
 * Libraries introduce huge contextual input, costly
 * Copies paradigms, not idiomatic.  Needs instruction
 * Claude seems to know it needs to change other files

Conclusion:
* Use claude to speed up translation
* Claude is a companion that needs close supervision
* Consider investing in Acceptance Tests prior to translation
  * Maybe use AI to generate temporary, disposable Acceptance Tests
* Frontend?