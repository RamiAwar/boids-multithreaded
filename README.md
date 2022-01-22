# Multithreaded Boids Simulation in Go

This is an exercise in goroutines and memory sharing with RWMutexes. Each boid updates itself on a separate goroutine, that runs every 5ms. In the run below, 1000 different goroutines are running concurrently and then the boid positions and headings are rendered once a frame.

The way neighbors are looked up here is by checking all other boids, which is very inefficient. This caps out at a maximum boid count of 1500 (90% CPU utilization reached at this point). For a better implementation, keep reading.

![1000 Boids, O(n2) Lookup](1000_boids.gif)


In the implementation below, a spatial hash grid was used to find the nearest boids. This works by hashing boids into cells in a grid, based on their position. The position vector {x, y} would be the key in this hash map, and the hashmap is updated accordingly with correct locking/unlocking. This implementation caps out at 2000 (80% CPU utilization). Pushing any further makes it lag, which might be due to memory limitations. To be explored.

![1300 Boids, Spatial Lookup, Low Cohesion](1300_boids_spatialhash_low_cohesion.gif)

