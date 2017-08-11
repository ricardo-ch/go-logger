# GO-LOGGER

go-logger provides a simple way to log messages through stderr and stdout with a minimal set of configuration.

## Quick start

```
// initialization (optional)
logger.InitLogger(true)

// use of logger
logger.Error("My error message")

logger.Info("My info message")

logger.Debug("My debug message") // only if verbose is true
```

*Error() publish to stderr, Info() and Debug() publish to stdout*

## Features

 - Simple logger to use
 - Publish to stderr and stdout messages

## License
go-logger is licensed under the MIT license. (http://opensource.org/licenses/MIT)

## Contributing
Pull requests are the way to help us here. We will be really gratefull.
