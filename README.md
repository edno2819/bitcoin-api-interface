# Interface to Bitcoin Node

This project provides an interface to interact with a Bitcoin node through the Bitcoin RPC API. It serves as a scaffold to create additional functions and wrappers for native Bitcoin RPC methods.

## Features

The program connects to a Bitcoin node and includes wrappers for several commonly used commands:

- **Blockchain Information:**
  - **Get Block Info:** Retrieve detailed information about specific blocks.
  - **Estimate Raw Fee:** Estimate the approximate fee per kilobyte needed for a transaction to begin confirmation within a certain number of blocks.

- **Wallet Management:**
  - **Create Wallet:** Initialize a new wallet on the Bitcoin node.
  - **Dump Wallet:** Export all wallet keys to a file for backup or analysis.
  - **Get Wallet Info:** Obtain detailed information about the current wallet, including balance and transaction data.
  - **Get Wallet Balance:** Check the total balance of the wallet.
  - **List Descriptors:** List all the script descriptors in the wallet, which define how keys and addresses are derived.

## Getting Started

### Prerequisites

- **Bitcoin Core:** Ensure you have Bitcoin Core installed on your system. You can download it from the [official website](https://bitcoin.org/en/download).
- **Bitcoin Node:** A running Bitcoin node is required to use the RPC interface.

### Starting the Bitcoin Node

To start the Bitcoin node in daemon mode, run the following command:

```bash
bitcoind --daemon
```

This will start the node in the background, allowing you to interact with it via RPC calls.

### Configuration

Ensure your `bitcoin.conf` file is properly configured to allow RPC connections. At a minimum, include the following settings:

```
server=1
rpcuser=your_rpc_username
rpcpassword=your_rpc_password
```

Replace `your_rpc_username` and `your_rpc_password` with your desired credentials.


## Usage

The interface provides functions to interact with the Bitcoin node. Below are examples of how to use some of the provided functions.

### Get Block Information

Retrieve detailed information about a specific block:

```go
rpcInterface.GetBlockInfo("blockhash")
```

### Estimate Raw Fee

Estimate the fee needed for transaction confirmation:

```go
rpcInterface.EstimateRawFee(confTarget)
```

- `confTarget`: The confirmation target in blocks (e.g., 6 for approximately one hour).

### Create a New Wallet

Create a new wallet on the node:

```go
rpcInterface.CreateWallet("wallet_name")
```

### Dump Wallet

Export all wallet keys to a file:

```go
rpcInterface.DumpWallet("/path/to/dumpfile.txt")
```

**Security Note:** The dump file contains sensitive information. Ensure it is stored securely.

### Get Wallet Information

Obtain detailed information about the current wallet:

```go
rpcInterface.GetWalletInfo()
```

### Get Wallet Balance

Check the total balance of the wallet:

```go
rpcInterface.GetWalletBalance()
```

### List Descriptors

List all script descriptors in the wallet:

```go
rpcInterface.ListDescriptors(includePrivate)
```

- `includePrivate`: Set to `true` to include private keys in the output.

## Additional Bitcoin CLI Commands

For more Bitcoin CLI commands and detailed documentation, visit the official Bitcoin Core documentation:

- [Bitcoin Core RPC Documentation](https://developer.bitcoin.org/reference/rpc/index.html)


---

**Disclaimer:** Interacting with a Bitcoin node and managing wallets involves significant security considerations. Ensure you understand the implications of the commands you execute and protect sensitive data accordingly.