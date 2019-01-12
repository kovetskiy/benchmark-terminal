# benchmark-terminal

Simple benchmark tool for measuring terminal performance for personal purposes.

See [testcases.sh](testcases.sh) and [suite.sh](suite.sh), it just runs
terminal and makes it print specified number of lines in given interval of
time.

# Some results
```
terminal: rxvt-unicode
line length: 151
18631627 lines per 30.000096963s (30.00 seconds)
speed: 621052.23 l/s
---
terminal: xterm
line length: 151
1352116 lines per 30.002044144s (30.00 seconds)
speed: 45067.46 l/s
---
terminal: alacritty
line length: 151
17986687 lines per 30.000099413s (30.00 seconds)
speed: 599554.25 l/s
---
terminal: kitty
line length: 151
6483412 lines per 30.002921788s (30.00 seconds)
speed: 216092.69 l/s
---
```
