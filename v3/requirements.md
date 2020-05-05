## Requirement

The `task` contract should fulfill the following requirements :
- it should provide a _proof of computation_ either through a _proof of stake_ or a _proof of work_ (probably not suitable).

## Limitation

Ethereum smart contract or any kind of fully decentralized app has some inherent limitation:
- No data is _private_ since all data written or modified has to be pushed to the network.
- No-one is trustworthy.

## Proof of computation requirements

A proof of computation should be resilient against such attack :
- A client pushing random value without doing any computation.
- (if result verification by another peer is used) a verifying client reading data already written by another peer and copying it.