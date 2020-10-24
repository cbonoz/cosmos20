<!--
order: 5
-->

# Parameters

The apis module has the following parameters:

| Key      | Type            | Example       | Description                                  |
| -------- | --------------- | ------------- | -------------------------------------------- |
| Requests | array (Request) | [{see below}] | array of params for each request in the apis |

Each `Request` has the following parameters

| Key        | Type               | Example                  | Description                                                      |
| ---------- | ------------------ | ------------------------ | ---------------------------------------------------------------- |
| RequestID  | string             | "bnb:usd"                | identifier for the request -- **must** be unique across requests |
| BaseAsset  | string             | "bnb"                    | the base asset for the request pair                              |
| QuoteAsset | string             | "usd"                    | the quote asset for the request pair                             |
| Oracles    | array (AccAddress) | ["kava1...", "kava1..."] | addresses which can post prices for the request                  |
| Active     | bool               | true                     | flag to disable oracle interactions with the module              |
