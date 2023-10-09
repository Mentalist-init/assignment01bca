# assignment01bca

The "assignment01bca" Go package is a simple implementation of a blockchain with basic functionalities for creating, displaying, changing, and verifying blocks. It also provides a hash calculation method.

## Table of Contents

- [Getting Started](#getting-started)
- [Usage](#usage)
- [Functions](#functions)
- [License](#license)

## Getting Started

To use this package in your Go project, you can follow the steps below:

1. Import the package in your Go code:

   ```go
   import (
       "github.com/Mentalist-init/assignment01bca"
   )


## Usage

Here's a brief example of how to use the package in your Go code:


package main

import (
    "github.com/Mentalist-init/assignment01bca"
)

func main() {

    // Create a new block
    block := assignment01bca.NewBlock("bob to alice", 123, "previousHash")

    // Display all blocks
    assignment01bca.DisplayBlocks()

    // Change the transaction of a block
    assignment01bca.ChangeBlock(block, "alice to bob")

    // Verify the blockchain
    isVerified := assignment01bca.VerifyChain()

    if isVerified {
        fmt.Println("Blockchain is verified.")
    } else {
        fmt.Println("Blockchain is not verified.")
    }
}

## Functions

The "assignment01bca" package provides the following functions:

1. NewBlock(transaction string, nonce int, previousHash string) *Block: Creates a new block and adds it to the blockchain.
2. DisplayBlocks(): Displays all blocks in the blockchain.
3. ChangeBlock(block *Block, newTransaction string): Changes the transaction of a given block.
4. VerifyChain() bool: Verifies the integrity of the blockchain.
5. CalculateHash(block Block) string: Calculates the hash of a block.

## License

This package is licensed under the MIT License. See the LICENSE file for details.

