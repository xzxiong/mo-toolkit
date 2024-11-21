# mo-toolkit


## Base cmd
```
$ ./build/mo-toolkit -h

Usage:
  mo-toolkit [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  pullMetrics  helps to pull metrics data from metric-given-server, through http://.../metrics url.
  queryService toolkit to call mo/query-service
  version      A brief description of your command

Flags:
  -h, --help     help for mo-toolkit
  -t, --toggle   Help message for toggle

Use "mo-toolkit [command] --help" for more information about a command.
```


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
