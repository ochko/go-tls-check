### SSL/TLS certificate validator

#### Use binary
```sh
go install github.com/ochko/go-tls-check/cmd/tls-check

tls-check example.com
tls-check -w 24h -t 5s example.com
```

Exit code is `0` when validation was successful, `1` if there is any issue.
It also prints some information in json format, so that you can collect expiration days of your deployed certificates:
```json
{
  "status": "ok",
  "certificateCheckHost": "example.com",
  "expirationDays": 305,
  "msg":"valid certificate"
}
```
When there is an issue:
```json
{
  "status": "ng",
  "certificateCheckHost":"unknown.com",
  "expirationDays":0,
  "msg":"dial tcp 23.253.58.227:443: i/o timeout"
}
```

#### Options
```
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
docker run tls-check example.com
```

### License

MIT
