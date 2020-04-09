# Go Logger

## Links
- https://dev-journal.in/2019/05/27/adding-uber-go-zap-logger-to-golang-project
- https://www.scalyr.com/blog/getting-started-quickly-with-go-logging
- https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce

```go
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")
```
