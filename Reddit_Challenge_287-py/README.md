# Reddit_Challenge_287-py

https://www.reddit.com/r/dailyprogrammer/comments/56tbds/20161010_challenge_287_easy_kaprekars_routine/

Find kaprekar solution to 4 digit integers in another file

The kaprekar hypothesis suggests that for any given 4 digit number that is NOT comprised of all the same number( ie. 1111,0000) we can arrive at the number 6174 through the following process:
1. Arrange the 4 digits as ascending
2. Arrange the 4 digits as descending
3. Take the ascending digits(as a 4 digit integer) and subtract the descending digits.  If the product is not 6174, the process is repeated and eventually you will reach 6174.

NOTE: in the event that a product is less than four digits, prepend 0's to bring it back to four digits (ex 154 is 0154).

The solution contained in this repository works as expected.  One improvement would be to log the number and number of iterations required to obtain a solution.  This could be accomplished through making listarry a global variable that is referenced either when a solution is found or no solution is possible.
