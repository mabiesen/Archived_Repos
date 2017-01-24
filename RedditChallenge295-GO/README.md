# RedditChallenge295-GO
Challenge was as follows: input string and desired output string are provided.  The input string is morphed letter by letter and printed to the terminal until the input matches desired output.

The only major difficulty presented by this challenge was the manner in which golang handles strings.  Strings in golang are composed of bytes and do not readily convert to letters.  Additionally, strings in golang are immutable(per online resources).  These difficulties were overcome with a function that converts strings to runes(character reference is how I would describe a rune at this point), replaces the desired runes, converts back to a string, and returns the string to a new variable in the main function.


https://www.reddit.com/r/dailyprogrammer/comments/5hy8sm/20161212_challenge_295_easy_letter_by_letter/?st=ix7wcwnc&sh=a27957c6
