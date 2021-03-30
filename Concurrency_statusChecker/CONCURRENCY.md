# Concurrency with go routines and channels

Concurrency is not parallelism

concurrency means we are doing multiple things at the same time, so for go routines to run at the same time in one core

but parallelism requires multiple cores/cpu to run the go routines at the same time

## Channels

before the introduction of channels, we have main go routines and child go routines, the main routine is when your main fn start/compile. Chile routines are one you can initiate with "go fn()" for concurrency. But if you just initiate child routines without communicating with the main routine, the main routine won't wait the child to finish its task and just exits.

Therefore channels are here to create a communication between main and child routines to establish bidirectional communication between routines, and to get your data before program exiting too early.

<-
