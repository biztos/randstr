# randstr

A utility for generating random strings in source code.

This program generates a string of "cryptographically secure" random bytes
within the ASCII displayable range of 32-126, excluding characters that could
make quoting difficult in the target language.

Supported target languages include:

- go - Google's Go language, aka Golang (in which this utility is written).
- perl - The Perl language, in all its glory!
- js - Javascript, used all over the Interwebs.

## TODO

- Modularize
- Test
- Make other useful noise data: numbers, names?
- Support other human languages, abugidas, what?
