<!--
order: 4
-->

# Events

The `x/apis` module emits the following events:

## MsgPostPrice

| Type                 | Attribute Key | Attribute Value    |
| -------------------- | ------------- | ------------------ |
| oracle_updated_price | request_id    | `{request ID}`     |
| oracle_updated_price | oracle        | `{oracle}`         |
| oracle_updated_price | request_price | `{price}`          |
| oracle_updated_price | expiry        | `{expiry}`         |
| message              | module        | apis               |
| message              | sender        | `{sender address}` |

## BeginBlock

| Type                  | Attribute Key | Attribute Value |
| --------------------- | ------------- | --------------- |
| request_price_updated | request_id    | `{request ID}`  |
| request_price_updated | request_price | `{price}`       |
| no_valid_prices       | request_id    | `{request ID}`  |
