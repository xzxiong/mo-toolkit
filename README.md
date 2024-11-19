# mo-toolkit


## CMD queryService

```
$ ./build/mo-toolkit queryService -h

Usage:
  mo-toolkit queryService [flags]
  mo-toolkit queryService [command]

Available Commands:
  GOGCPercent helps to call MO QueryService/GOGCPercent api
  GOMEMLimit  helps to call MO QueryService/GOMEMLimit api
  GOMaxProcs  helps to call MO QueryService/GOMaxProcs api
  MetaCache   helps to call MO QueryService/MetaCache api

Flags:
  -h, --help                             help for queryService
      --query-service-host string        mo query-service host. (default "127.0.0.1")
      --query-service-port int           mo query-service port, example: 6004 (in cloud) (default 18002)
      --query-service-timeout duration   timeout (default 3s)

Use "mo-toolkit queryService [command] --help" for more information about a command.
```
