### Fan-out

Process of starting multiple goroutines to handle input from the pipeline

Requirements for stage:

- it doesn't rely on values that the stage had calculated before
- it takes a long time to run

### Fan-in

Process of combining multiple results in one channel
