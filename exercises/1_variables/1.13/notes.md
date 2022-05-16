## With values such as int, bool, and string, when you pass them to a function, Go makes a copy of the value, and it's the copy that's used in the function. This copying means that a change that's made to the value in the function doesn't affect the value that you used when calling the function.

- CONS: copying uses up more and more memory as values get passed from function to function

The heap allows the value to exist until no part of your software has a pointer to it anymore. Go reclaims these values in what it calls its garbage collection process. This garbage collection happens periodically in the background, and you don't need to worry about it.

Having a pointer to a value means that a value is put on the heap

Working out whether a value needs to be put on the heap is called escape analysis

Beyond performance, you can use pointers to change your code's design. Sometimes, using pointers allows a cleaner interface and simplifies your code.
