# RedditChallenge296-GO
Challenge is to write out the twelve days of christmas with minimalistic coding.  While I imagine my code can be shortened, the use of mapping in this project was a good experience.

Things that I might be able to improve upon:
```
1. Iterate through maps without the assistance of an array.  Maps are not indexed by number and therefore cannot be referenced by number; however, a for each styled statement may allow for map iteration.  To investigate.
2.  I suspect that ordinal numbers are already created in go; I could potentially remove the ordinal number array.
```
I haven't tried mapping integers yet, but I am curious as to whether or not they retain their value type when placed in a map.  I have noticed some languages(c in arduino??) do not always retain the data type of integers when placed in arrays, etc.


Below you will find the code of Reddit user skeeto for this challenge in C:
```
#include <stdio.h>

int
main(void)
{
    char *days[] = {"first", "second", "third", "four", "fif", "six",
                    "seven", "eigh", "nin", "ten", "eleven", "twelf"};
    for (int i = 0; i < 12; i++) {
        printf("On the %s%s day of Christmas\n", days[i], "th" + (i < 3) * 2);
        puts("my true love sent to me:");
        switch (i) {
            case 11: puts("twelve Drummers Drumming");
            case 10: puts("eleven Pipers Piping");
            case 9:  puts("ten Lords a Leaping");
            case 8:  puts("nine Ladies Dancing");
            case 7:  puts("eight Maids a Milking");
            case 6:  puts("seven Swans a Swimming");
            case 5:  puts("six Geese a Laying");
            case 4:  puts("five Golden Rings");
            case 3:  puts("four Calling Birds");
            case 2:  puts("three French Hens");
            case 1:  puts("two Turtle Doves");
            case 0:  printf("%sa Partridge in a Pear Tree\n\n", "and " + !i*4);
        }
    }
}
```

I thought this was some nifty code because it uses the implicit numeric values of true and false to determine whether the day should have a 'th' on the end.
The code is also an example of how pointers are used (*).
