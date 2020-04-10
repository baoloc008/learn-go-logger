```go
log.Debug("Useful debugging information.")

log.Info("Something noteworthy happened!")

log.Warn("You should probably take a look at this.")

log.Error("Something failed but I'm not quitting.")

log.Fatal("Bye.") // Calls os.Exit(1) after logging

log.Panic("I'm bailing.") // Calls panic() after logging
```
