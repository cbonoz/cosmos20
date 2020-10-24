<!--
order: 6
-->

# End Block

At the end of each block, the current price is calculated as the median of all raw prices for each request. The logic is as follows:

```go
// EndBlocker updates the current apis
func EndBlocker(ctx sdk.Context, k Keeper) {
	// Update the current price of each asset.
	for _, request := range k.GetRequests(ctx) {
		if request.Active {
			err := k.SetCurrentPrices(ctx, request.RequestID)
			if err != nil {
				// In the event of failure, emit an event.
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						EventTypeNoValidPrices,
						sdk.NewAttribute(AttributeRequestID, fmt.Sprintf("%s", request.RequestID)),
					),
				)
				continue
			}
		}
	}
	return
}
```
