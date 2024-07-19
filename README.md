# GRPC Client-Server template in Go

This repository introduces a `proto3` file designed to simulate RPC communications in a Proof of Stake (PoS) Blockchain environment.
The `proto3` file facilitates several key functionalities, including client-server interactions for subscribing, fetching block information, adding transaction data, and block baking.

### Features

1. **Client Subscription and Reputation Assignment**
    - Clients can subscribe to the server and are assigned a UUID and a reputation score upon registration.
    - Registration involves sending an empty body request and receiving a UUID and reputation score in response.

2. **Fetching Last Block Information**
    - Clients can call the server to fetch the last block information.

3. **Adding Transaction Data**
    - Clients can call the server to add transaction data.

4. **Block Baking Process**
    - Clients subscribe to the server to participate in block baking.
    - The server randomly selects a client every 30 seconds to bake a block.
    - The chosen client receives a notification and upon sending back a confirmation, the block is mined, and the client's reputation increases.

### Code Generation

This repository includes the Go code generated from the `proto3` file for both the client and the server.
The client code includes a `Flag.parse()` function, allowing the user to choose which RPC method to invoke at the program's start.

### Implementation Details

1. **Proto3 File**
    - Defines the RPC methods and messages required for the described functionalities.
    
2. **Go Client and Server Code**
    - Implements the client and server logic based on the proto3 definitions.
    - The client includes a `Flag.parse()` function for selecting RPC methods.

### Usage

A Makefile is included in the repository to facilitate the compilation and execution of the client and server code.
The following commands can be used to run the client and server code:

- **Server**
    - The server handles subscriptions, block information requests, transaction data addition, and the block baking process.

- **Client**
    - The client can be started with different flags to call the appropriate RPC method.
        - `go run client.go -lastblock`
        - `go run client.go -register`
        - `go run client.go -subscribe -uuid <uuid>`
        - `go run client.go -addtransaction -data <data> -uuid <uuid>`
        - `go run client.go -bakeblock -uuid <uuid>`
        - `go run client.go -confirmblock -uuid <uuid>`