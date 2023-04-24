### Port scanner

Localhost port scanner writen in Go.
It checks all 65,535 possible ports and prints out the open ones.

It uses goroutines to execute it in parallel, creating 100 workers to distribute the overload.

Reference:
- [Worker Pools pattern](https://gobyexample.com/worker-pools)
- [Range over Channels](https://gobyexample.com/range-over-channels)
