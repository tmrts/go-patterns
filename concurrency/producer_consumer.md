# Producer Consumer

What is the producer-consumer pattern?
> In computing, the producer–consumer problem (also known as the bounded-buffer problem) is a classic example of a multi-process synchronization problem. The problem describes two processes, the producer and the consumer, who share a common, fixed-size buffer used as a queue. The producer's job is to generate data, put it into the buffer, and start again. At the same time, the consumer is consuming the data (i.e., removing it from the buffer), one piece at a time. The problem is to make sure that the producer won't try to add data into the buffer if it's full and that the consumer won't try to remove data from an empty buffer. --- from [wikipedia](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem)

## Implementaion

More information can be found in [producer_consumer package](producer_consumer).

## Example

An example can be found in [producer_consumer.go](producer_consumer.go).
