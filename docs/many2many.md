## Questions

### What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?
  
I think: that the program will probably close immediately, since the channel being used to communicate everything is closed immediately, ending the program.

In actuality: that's what happened, although I didn't expect the program to panic.

### What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?

I think: the program will crash much like the scenario above, just after one producer finishes instead of immediately.

In actuality: yeah...

### What happens if you remove the statement `close(ch)` completely?

I think: the program will probably throw a panic about a deadlock once the main routine comes to an end.

In actuality: It did not, which I suppose is because the main routine never actually goes to sleep at the end. Can't see any difference from when close(ch) is included after after wgp.Wait()

### What happens if you increase the number of consumers from 2 to 4?

I think: that almost all of the data sent will be received, with some missed occasionally given random variance in RandomSleep().

In actuality: I realize now that the base program actually usually consumes all of the data sent, I just read the console wrong. Whoops. The real difference is that the program is about twice as fast.

### Can you be sure that all strings are printed before the program
  stops?
  
I think: that since the program doesn't actually wait for all the consumers to finish, once all the producers are done then the program can end. I believe this means that the program will sometimes skip one or two of the last messages.

In actuality: I was wrong? After a lot of test runs it has yet to miss a message, but maybe I'm just reading them wrong. Maybe luck is on my side.

