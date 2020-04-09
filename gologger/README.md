# Pros and Cons of Go Logger
## Pros
The biggest advantages is that it is very simple to use. We can set any io.Writer to log output and send write logs to it.

## Cons
- Basic Log Levels only
    - Only has one option for `Print`. Does not supports multiple levels like `INFO`/ `DEBUG`.
- For Error logging, it has `Fatal` and `Panic`
    - Fatal Logging ends program by calling `os.Exit(1)`
    - Panic Logging throws a `panic` after writing the log message.
    - It however lacks a `ERROR` log level, that can log errors without throwing a panic or exiting the program.
- Lacks log message formatting capabilities – e.g logging caller function name and line number, formatting the date and time format, etc.
- Does not provides log rotation capabilities.
