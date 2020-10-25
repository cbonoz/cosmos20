<!--CosmosHub API Module

### Concept

Cosmos API module (apis) - interact with arbitrary api's from your cosmos app using this prebuilt cosmos hub module.

Going after the build a module challenge for the cosmos hackatom contest.

### Use cases

- Make api calls in response to key events such as block starts or ends, arbitrary messages, or ad-hoc queries.
- Perform programmatic access to other services such as other blockchains, third party api's
- Provide a simple interface to configure a set of API-based commands.

### Usage

1. Add the /x/apis module to your app
2. Configure the module with your external API's and desired parameters.
3. Create subscribers to process emitted events.

-->

### Running the project

`go build -v ./...`

### Running tests

`go test -v ./...`

### Dev notes

- brew install tendermint/tap/starport

### Useful links

- https://cosmos.network/intro
- https://github.com/cosmos/cosmos-sdk
- https://github.com/tendermint/starport/blob/develop/docs/1%20Introduction/2%20Install.md
- https://docs.cosmos.network/master/building-modules/intro.html
- https://docs.google.com/document/d/16XTVG6j2TQw53DWjs_QAEBXNFukxRzQohsmL1K3JYVA/edit?usp=sharing
- https://tutorials.cosmos.network/
- https://five.hackatom.org/resources

**cosmos20** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
