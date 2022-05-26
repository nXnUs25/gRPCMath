# gRPCPrims
Primes RPC Server Streaming API

## Primes API
In this exercise, implement a Primes RPC Server Streaming API:

The function takes a Request message that has one integer, and returns a stream of Responses that represent the prime number decomposition of that number (see below for the algorithm).

#### Example:
The client will send one number (120) and the server will respond with a stream of (2,2,2,3,5), because 120=2*2*2*3*5 

Algorithm (pseudo code):

```c
k = 2
N = 210
while N > 1:
    if N % k == 0:   // if k evenly divides into N
        print k      // this is a factor
        N = N / k    // divide N by k so that we have the rest of the number left.
    else:
        k = k + 1
```
