### SSL/TLS certificate validator

#### Install binary
```sh
go install github.com/ochko/go-tls-check/cmd/tls-check
```

#### Usage
```sh
tls-cert-check [options] hostname1 hostname2 ...
  options:
  -t string
    	Connection timeout. (default "10s")
  -w string
    	Allowd time before certificate expiration. (default "72h")
```

#### Examples
```sh
tls-check example.com your.site.com
tls-check -w 24h -t 5s example.com your.site.com
```

### Using in your code

```golang
import "github.com/ochko/go-tls-check/validator"

...

expirationDays, err := validator.Check("example.com",time.Hour*24, time.Second*3)

...

```

### License

MIT
