# Time

## Find elapsed time till now

## Convert the timestamp to the specified timezone

```go
  loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(fmt.Sprintf("Failed to load timezone: %v", err))
	}

		localTime := payment.Timestamp.In(loc)

		// Extract the date part (YYYY-MM-DD)
		date := localTime.Format("2006-01-02")
```
