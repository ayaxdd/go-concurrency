##### without buffering

zeros - take 3 zeros
short - sleep 1 sec
long - sleep 4 sec

| time (s) | zeros         | short                  | long                  |
| -------- | ------------- | ---------------------- | --------------------- |
| 0        | zeros <- zero | short <- zeros sleep 1 |                       |
| 0        | zeros <- zero | sleeping (full)        |                       |
| 1        | zeros <- zero | short <- zeros sleep 1 | long <- short sleep 4 |
| 2        | close         | sleeping (full)        | 3                     |
| 3        | close         | sleeping (full)        | 2                     |
| 4        | close         | sleeping (full)        | 1                     |
| 5        | close         | sleeping (full)        | long <- short sleep 4 |
| 6        | close         | short <- zeros sleep 1 | 3                     |
| 7        | close         | sleeping (full)        | 2                     |
| 8        | close         | sleeping (full)        | 1                     |
| 9        | close         | close (0 rem)          | long <- short sleep 4 |
| 10       | close         | close                  | 3                     |
| 11       | close         | close                  | 2                     |
| 12       | close         | close                  | 1                     |
| 13       | close         | close                  | 0                     |

##### with buffering

zeros - take 3 zeros
short - sleep 1 sec
buffer - add short (cap 2)
long - sleep 4 sec

| time (s) | zeros         | short                     | buffer              | long                     |
| -------- | ------------- | ------------------------- | ------------------- | ------------------------ |
| 0        | zeros <- zero | sleep 1                   | 0/2                 |                          |
| 1        | zeros <- zero | short <- zeros && sleep 1 | 0/2                 | long <- short && sleep 4 |
| 2        | zeros <- zero | short <- zeros && sleep 1 | 1/2 buffer <- short | 3                        |
| 3        | closed        | short <- zeros && close   | 2/2 buffer <- short | 2                        |
| 4        | closed        | closed                    | 2/2                 | 1                        |
| 5        | closed        | closed                    | 1/2                 | long <- short && sleep 4 |
| 6        | closed        | closed                    | 1/2                 | 3                        |
| 7        | closed        | closed                    | 1/2                 | 2                        |
| 8        | closed        | closed                    | 1/2                 | 1                        |
| 9        | closed        | closed                    | 0/2                 | long <- short && sleep 4 |
| 10       | closed        | closed                    | 0/2                 | 3                        |
| 11       | closed        | closed                    | 0/2                 | 2                        |
| 12       | closed        | closed                    | 0/2                 | 1                        |
| 13       | closed        | closed                    | 0/2                 | closed                   |

### Little's Law

L = λW, where:

- L = the avg number of units in the system (stages)
- λ = the avg arrival rate of units (stage/time)
- W = the avg time a unit spends in the system (time for 1 unit)

Applies only to so-called stable systems

Stable system is one which the rate that work enters the pipeline, or ingress, is equal to the rate in which it exits the system, or egress
