# Dominion app-channel

## Requirements

1. Install [ganache-cli](https://github.com/trufflesuite/ganache-cli)
2. Install GoLang
3. Run following to install solidity using asdf
    ```sh
    asdf plugin add solidity
    asdf install solidity 0.7.6
    ```

## Execution

1. Generate interface for DominionApp on Blockchain
   ```sh
    cd ./contracts && ./generate.sh && cd ..
    ```
2. Start a local blockchain with required acc
   ```sh
   ./ganache-cli.sh
    ```
3. Use a new terminal to start the application
   ```sh
   go run .
   ```