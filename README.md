
# gRPCPrims
Primes RPC Server Streaming API

## Sum API
In this exercise, the goal is to implement a Sum RPC Unary API in a MathService:

The function takes a Request message that has two integers, and returns a Response that represents the sum of them.

#### Example:

The client will send two numbers (3 and 10) and the server will respond with (13)

## Primes API
In this exercise, implement a Primes RPC Server Streaming API:

The function takes a Request message that has one integer, and returns a stream of Responses that represent the prime number decomposition of that number (see below for the algorithm).

#### Example:
The client will send one number (120) and the server will respond with a stream of (2,2,2,3,5), because 120=2*2*2*3*5 

Algorithm (pseudo code):

---
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
---

## Avg API
In this exercise, implement a Avg RPC Client Streaming API in a MathService:

The function takes a stream of Request message that has one integer, and returns a Response with a double that represents the computed average

#### Example:

The client will send a stream of number (1,2,3,4) and the server will respond with (2.5), because (1+2+3+4)/4 = 2.5 

## Max API

In this exercise, the goal is to implement a Max RPC Bi-Directional Streaming API in a MathService:

The function takes a stream of Request message that has one integer, and returns a stream of Responses that represent the current maximum between all these integers

#### Example:

The client will send a stream of number (1,5,3,6,2,20) and the server will respond with a stream of (1,5,6,20)
