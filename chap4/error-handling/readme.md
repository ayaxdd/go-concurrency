### General rule

Concurrent processes should send their errors to another part of program that has complete info about the state of program

Goroutine's produced errors should be tightly coupled with result type and passed along through the same lines of communication (channels)
