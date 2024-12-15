# Boids: Birds + Android, MULTITHREADING in golang

Boids is a multithreaded simulation that treats each boid (bird-like entity) as an independent thread. The project explores different configurations of multithreading, both without Inter-Process Communication (IPC) and with IPC, to simulate swarms.

This example leverages the concept of flocking behavior, where multiple boids interact within a constrained 2D environment. The goal is to study multithreading concepts while visualizing the interactions of autonomous agents in a simple simulation.

## Features

- **Multithreaded Boid Simulation:** Each boid is represented as a thread.
- **Swarm Dynamics:** Simulates collective behavior of boids in a confined space.
- **IPC and Non-IPC Configurations:** Demonstrates multithreading with and without Inter-Process Communication.
- **Visualization:** Uses the Ebiten library for rendering the simulation in real-time.
- **Configurable:** Adjustable number of boids, screen dimensions, and simulation parameters.

## Getting Started

### Prerequisites

- Go (version 1.19 or later)
- Ebiten library: Install via:
  ```bash
  go get github.com/hajimehoshi/ebiten/v2
  ```

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/genGit963/mutlithreading-concureny-golang
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the simulation:
   ```bash
   go run .
   ```

### Usage

#### Configuration

Modify constants in the `main.go` file to customize the simulation:

- `screenWidth, screenHeight`: Dimensions of the simulation window.
- `boidCount`: Number of boids in the simulation.

#### Visualization

- **Green dots**: Represent individual boids.
- The simulation window updates in real-time as boids move and interact.

## Code Structure

- `main.go`: Contains the entry point, game loop, and visualization logic.
- `boids.go`: Defines the `Boid` struct and related behaviors (movement, threading).
- `vector2d.go`: Provides utility functions for 2D vector operations.

## Examples

### Without IPC

Each boid operates independently in its own thread without communication with other boids. This mode demonstrates simple multithreading.

### With IPC

Boids communicate through shared memory or message-passing mechanisms, allowing for swarm behaviors such as alignment, cohesion, and separation.

## License

This project is licensed under the ISC License. See the LICENSE file for details.

## Code of Conduct

We welcome contributions from everyone. By participating in this project, you agree to abide by our Code of Conduct, which fosters a welcoming and inclusive environment for all contributors.

---
