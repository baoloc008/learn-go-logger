# Zap Logger
## Why Uber-go zap
- It provides both structured logging and printf style logging
- It is supposedly very fast.

## Usage
- Create a Logger, by calling `zap.NewProduction()` / `zap.NewDevelopment()` or `zap.NewExample()`
- Each of the above will create a logger. The only difference is in the information it will log. e.g the production logger logs the calling function information, date and time, etc by default.
- Call Info/Error, etc on the Logger.
- By default the logs will come to the application console.