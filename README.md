# msq

Durable message queue

Single lookup daemon

Multiple topic servers

Each topic exists only on one topic server

Each topic server has multiple topics

E.g. distributed time series blob storage

## Consumer

Reads a stream of byte arrays

Each consumer has its own offset stored in the topic server marking the id of
the message the consumer has last read

When a consumer reconnects it will resume reading from the last message offset
that was successfully acknowledged

## Producer

Connects to the lookup server to find the addresses of topic servers to publish to

Can publish to many topics

Single writer thread

Multiple reader threads

## Gateway

Performs admin functions

- Create and delete topics

- Move topics between nodes

## Log Files

Logs are composed of a metadata file and a series of segment files

Each segment file 
