# Logger

A simple logging wrapper


Basic usage:

```
    logger.Log("Foobar")
    logger.Error("Foobar")
    logger.Fatal("Foobar")
    logger.Debug(someVar)
    logger.Debug(fmt.Sprintf("Value is: %v", someVal))
```

Enable debug lines with `logger.SetDebug()`

