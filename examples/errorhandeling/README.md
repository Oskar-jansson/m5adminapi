# Error Handling


## sdkError
The SDK communicates errors through its message interface. Always check for errors when making calls:

```go
result, sdkErr := api.Version.Get(ctx)
if sdkErr != nil {
    // Handle SDK error
    println("Error:", sdkErr)
}

if result != nil {
    // Process result
    println(result)
}
```

`SdkError` consistce of two parts:
- Err `error`
- ErrorMessage `api provided error`

Example error:
```go
{
    // SDK-level error (not from the API)
    Err: "something went wrong",

    // API response error
    ErrorMessage: nil
}
```

Sdk also provides following methods for sdkError:
- `Unwrap()`
- `AsError()`

Example:
```go
result, sdkErr := api.Version.Get(ctx)

var err error

if err = sdkErr.AsError(); err != nil {

	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("request timed out")
		return
	}

	if errors.Is(err, context.Canceled) {
		fmt.Println("request was canceled")
		return
	}

	fmt.Println("request failed:", err)
	return
}

if result != nil {
	fmt.Println("version:", result)
}
```


## Patterns
Handling Different Error Scenarios

1. General API errors:
Use the client’s exposed hooks to handle these. Hooks are triggered for every request and can implement retry logic automatically.
This method should only apply to general errors as hooks are applied to all request from same `AdminApiConnection`.

2. Request-specific or unique errors:
These occur closer to the actual request execution and should be handled manually in your code as retry will most likly not solve problem.