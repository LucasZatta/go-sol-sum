# GO Max Contiguous Subarray Sum
This is a simple GO API implemented using Gin Gonic. Its purpose is to solve the Max contiguous subarray sum problem.

## Project Structure:

```
├── controllers
│   └──subarraysum.go  \\Main logic accessed by the router
├── models
│   └──maxsum.go  \\POST request structure and Algorithms implementation
├── router
│   └──router.go  \\Router and paths
├── Dockerfile    \\Dockerfile configurations(ENV, commands, dependencies installation)
├── go.mod        \\Go dependencies file
├── go.sum
├── main.go        
```
## Installation
Approach 1: Inside the project file, run the command `docker build -t golang-maxsum`, this will buid a docker image containing the project using the dockerfile configurations.
Once the build is finished, run the command `docker run -p 8080:8000 golang-maxsum`. This will run the image with specified ports parameter.

Approach 2: If you rather not clone this repository, just pull the DockerHub repository containing the project image.
`docker pull lucasszatta/go-maxsum`
After pulling the image, run it using the command: `docker run -p 8080:8000 lucasszatta/golang-maxsum`

## REST API
Solve Max subarray sum problem
### Request
  `POST /maxsum`
  Body Example: { "list": [-2, 3, 5, -1, 4, -5] }
  
  Response Example: { "result": 11 }
  
## Problem

The max subarray sum problem is a pretty straight forward one, given an one dimensional list of n element, the task is to find a contiguous subarray with the largest sum. The optimal solution for this problem (O(n)) is achieved using Kadane's algorithm.

### Kadane's Algorithm

The basic idea is to keep track of the highest subset sum yet, and compare it with the next subset sums and replace it if necessary. Iterating all the indexes and calculating the sum of every possible subarray ending in that index.

Instead of going through every possible subset, Kadane proved, through contradiction, that the max subset of a given index *n* is either the element at *n* itself or the max subset sum ending at the previous index *n-1* + element at *n*. Therefore we can obtain the max subset sum ending in a certain index doing a simple comparrison between two values, achieving a linear solution to the problem.
