# Logrus logger adapter for Echo

## Overview

Simple adapter for [Logrus](https://github.com/sirupsen/logrus) logger
for [Echo](https://github.com/labstack/echo).

Advantages over other implementations:
- doesn't use unconfigurable singleton logger
- doesn't panic when it shouldn't (and I believe such adapter should **NEVER** panic)

## Usage

```go
e := echo.New()

logger := logrus.New() // or use `logrus.StandardLogger()` for predefined singleton

// Logger adapter is used by `echo.Echo`
adapter := echologrus.LoggerAdapter{logger}
e.Logger = adapter

// Logger middleware
middleware := echologrus.Middleware(logger)
// or
middleware := echologrus.Middleware(logger, echologrus.WithSkipper(...))

e.Use(middleware) 
```
