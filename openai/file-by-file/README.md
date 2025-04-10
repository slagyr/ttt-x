# Translate a repo from one language (Clojure) to another (Golang) with a file by file approach.

The hope for this experiment is that by minimizing information exposed to the LLM at a given time, it would incentivize
more direct line-by-line translation that would maintain the same logic as the original tic-tac-toe implementation.

Ultimately, I found this approach to be even less ideal than the `whole-repository` approach. Not only was ChatGPT unable
to maintain the logic through the translation, the files also no longer worked with each other. This vastly increased
the amount of errors that needed to be manually fixed, with no real extra benefit in return.

I was hopeful that the AI would at least recognize the minimax implementation but that was not true. First it made an
AI that makes moves randomly. In response I told it `That last file you translated is missing information on how the bot works, what do you think it is and can you fix it?`
In response ChatGPT gave a slightly more advanced bot but still was unable to parse that minimax was the implementation
being used. Instead, the approach was to make a bot that just blocked a winning move against the player, or took a winning
move if it saw one.

One perk was I found test files to be slightly more detailed, but still lacked good translation. Some tests were made up,
others didn't honor the same test set up, and some tests were excluded entirely. I also had to fix a syntax error
because ChatGPT didn't know how to translate `with-out-str` very well