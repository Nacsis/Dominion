# Dominion app-channel

## Requirements

1. Install [ganache-cli](https://github.com/trufflesuite/ganache-cli)
2. Install GoLang
3. Run following to install solidity using asdf
    ```sh
    asdf plugin add solidity
    asdf install solidity 0.7.6
    ```

## Cobra/Viper Setup
1. install [cobra](https://github.com/spf13/cobra) and [cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md)
   ```sh
   go get -u github.com/spf13/cobra@latest
   go install github.com/spf13/cobra-cli@latest
   ```
2. init cmd structure
   ```sh
   cobra-cli init --viper
   ```
   CAUTION: ```cobra-cli init``` overrides files (e.g. ```main.go```, ```cmd/root.go```)

## Execution

1. Generate interface for DominionApp on Blockchain
   ```sh
    cd ./contracts && ./generate.sh && cd ..
    ```
2. Start a local blockchain with required acc
   ```sh
   ./ganache-cli.sh
    ```
3. Use a new terminal to start the demo application
   ```sh
   go run main.go demo
   ```

### interactive playing
...not yet implemented, but will work like this:
1. Start a local blockchain with required acc
   ```sh
   ./ganache-cli.sh
    ```
2. Use a new terminal to start a cli as alice
   ```sh
   go run main.go play --config alice.yml
   ```
3. Use a new terminal to start a cli as bob
   ```sh
   go run main.go play --config bob.yml
   ```