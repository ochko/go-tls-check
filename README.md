### SSL/TLS certificate validator

#### Use binary
```sh
go install github.com/ochko/go-tls-check/cmd/tls-check

tls-check example.com your.site.com
tls-check -w 24h -t 5s example.com your.site.com
```

#### Options
```sh
tls-cert-check [options] hostname1 hostname2 ...
  options:
  -t string
    	Connection timeout. (default "10s")
  -w string
    	Allowd time before certificate expiration. (default "72h")
```

### Importing the package

```golang
import "github.com/ochko/go-tls-check/validator"

...

expirationDays, err := validator.Check("example.com", time.Hour*24, time.Second*3)

...

```

### Using the docker image

```sh
docker build . -t tls-check
docker run -ti --rm tls-check example.com
```

### License

MIT
