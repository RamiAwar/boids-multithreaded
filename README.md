# Multithreaded Boids Simulation in Go

This is an exercise in goroutines and memory sharing with RWMutexes. Each boid updates itself on a separate goroutine, that runs every 5ms. In the run below, 1000 different goroutines are running concurrently and then the boid positions and headings are rendered once a frame.

![1000 Boids, O(n2) Lookup](1000_boids.gif)
