Bitrix REST API client in Go
============================

This library implements Bitrix and Bitrix24 [REST API](https://training.bitrix24.com/rest_help/)

Installation
------------
```bash
go get -u github.com/ikarpovich/go-bitrix
```

Usage
-----
```go
import "github.com/ikarpovich/go-bitrix"
```

Roadmap
-------
- Basic client with OAuth support
- Abstraction layer for CRUD operations
- Auto collection of methods and entities from PHP code
- WebHook support
- Event handling
- Secure method calls (JWS)