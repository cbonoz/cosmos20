<p align='center'>
    <img src="./img/apikit.png"/>
</p>

## CosmosHub API Module

### Concept

Cosmos API module (apis) - interact with arbitrary api's from your cosmos app using this prebuilt cosmos hub module.

Going after the build a module challenge for the cosmos hackatom contest for increasing Cosmos utility and go to market.

Note this is a prototype/hackathon concept module.

### Value adds

This module is for you if:

- You want to easily make api calls in response to key events such as block starts or ends, arbitrary messages, or ad-hoc queries.
- Perform programmatic access to other services such as other blockchains, third party api's (increase interoperability).
- Provide a simple interface to configure a set of API-based commands even without detailed knowledge of go http or cosmos. Simply provide the urls and payloads and the module can take care of the rest.
- Centralize all your API-related work in your cosmos application.

### Integrating with your project

1. Add the /x/apis module to your app
2. Configure the module with your external API's and desired parameters.
3. Create subscribers to process emitted events.
4. Edit or fork the module however you wish.

### Example usage

1. Adjust `DefaultParams` in params.go of the apis module (currently configured in the cosmos20 app to poll for weather and emit an event).
2. `starport serve -v`

You should see these requests going out in response to the EndBlock event.

<img src="./img/debug.png" width=600>

3. Subscribe to the events via the tendermint websocket.

<img src="./img/subscribe.png" width=600>

Can see evidence of the events emitted here.

<img src="./img/event.png" width=600>

### Running the project

`go build -v ./...`

### Running tests

`go test -v ./...`

### Dev notes

- brew install tendermint/tap/starport
- starport serve

### Useful links

- https://www.youtube.com/watch?v=PCSzgJCgwdE
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
