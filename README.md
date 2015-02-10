# CS370 Assignment 2 README
I was interested in learning the programming language Go, so I learned it over the past few days and have completed this assignment in Go. I am willing to display what I have done on my machine or install it on a school computer if you are not able to get it working.

### Setup:
1. Download and install Go from https://golang.org/dl/ or using $ brew install go.

### To Run:
1. Run from the src directory: $ go run scanner.go < test.pn
2. Or run from the src directory: $ go build scanner.go
3. If $ go build scanner.go, then run ./scanner < test.pn

### Petri Net Format:
place 5
transition 3
P0 : T1 : 2
P1 : T1 : 3
…
(1,1,0,0,0)