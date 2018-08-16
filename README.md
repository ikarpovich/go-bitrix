Bitrix REST API client in Go
============================

This library implements Bitrix and Bitrix24 [REST API](https://training.bitrix24.com/rest_help/)

Features
-------
- Generic REST client, responses and errors handling
- Auto collection of methods (full) and entities (partial) from PHP code
- WebHook Auth


Installation
------------
```bash
go get -u github.com/ikarpovich/go-bitrix
```

Usage
-----
```go
See [examples](https://github.com/ikarpovich/go-bitrix/blob/master/examples/common_methods/main.go)
```

Roadmap
-------

- CRUD operations
- Event handling
- OAuth
- Secure method calls (JWS)