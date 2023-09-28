# cloud-config-loader

```shell
go get github.com/tarantinoz/cloud-config-loader
```

## Usage
```go
import (
    "log"
    "github.com/tarantinoz/cloud-config-loader"
)

type applicationConfig struct {
    Key1 Key1Type `json:"key1" yaml:"key1"`
    // ...
}

func main() {
    client, err := cloudconfigloader.NewClient(
        "localhost:8080/config",
        "APP_NAME",
        "production",
    )
    if err != nil {
        panic(err)
    }

    // get a reader to configuration
    rpd, err := client.Raw()
    if err != nil {
        panic(err)
    }
    
    if rpd == nil {
        panic("get cloud config error")
    }
	
    // or, decode configuration directly into a struct
    var appConfig applicationConfig
    err = client.Decode(&appConfig)
    if err != nil {
        panic(err)
    }
}
```

### Basic Auth

```go
client, err := cloudconfigloader.NewClient(
    server, 
    application, 
    profile, 
    cloudconfig.WithBasicAuth("username", "password"),
)
```

### Reading config as YAML instead of (default) JSON

```go
client, err := cloudconfigloader.NewClient(
    server,
    application,
    profile, 
    cloudconfig.WithFormat(cloudconfig.YAMLFormat)),
)
```