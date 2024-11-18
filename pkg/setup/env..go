package setup

import (
	"github.com/google/uuid"
	"github.com/matrixorigin/matrixone/pkg/common/morpc"
	"github.com/matrixorigin/matrixone/pkg/common/runtime"
	"github.com/matrixorigin/matrixone/pkg/defines"
	"github.com/matrixorigin/matrixone/pkg/logutil"
	qclient "github.com/matrixorigin/matrixone/pkg/queryservice/client"
	"go.uber.org/zap"
)

var serviceId uuid.UUID

func Init() {
	serviceId = uuid.Must(uuid.NewRandom())
	// init for rpc env
	setupMORuntime(serviceId.String())
}

func GetDefaultRpcConfig() *morpc.Config {
	var cfg morpc.Config
	cfg.Adjust()
	return &cfg
}

func GetQueryClient(rpcCfg *morpc.Config) (qclient.QueryClient, error) {
	if rpcCfg == nil {
		rpcCfg = GetDefaultRpcConfig()
	}
	client, err := qclient.NewQueryClient(serviceId.String(), *rpcCfg)
	if err != nil {
		logger.Error("failed to init QueryService client", zap.Error(err))
		return nil, err
	}
	return client, nil
}

var logger *zap.Logger

func GetLogger() *zap.Logger {
	return logger
}

func setupMORuntime(serviceId string) {
	setupMOLogutil(true)
	logger = logutil.GetGlobalLogger()
	setupMOLogutil(false)
	rt := runtime.DefaultRuntimeWithLevel(zap.InfoLevel)
	rt.SetGlobalVariables(runtime.MOProtocolVersion, defines.MORPCLatestVersion)
	// ## for mo 1.2.*
	//runtime.SetupProcessLevelRuntime(rt)
	// ##for mo 1.3.*
	runtime.SetupServiceBasedRuntime(serviceId, rt)
}

func setupMOLogutil(logs bool) {
	// close mo/pkg log.
	logutil.SetupMOLogger(&logutil.LogConfig{
		Level:           "info",
		Format:          "console",
		Filename:        "",
		MaxSize:         0,
		MaxDays:         0,
		MaxBackups:      0,
		DisableStore:    true,
		DisableLog:      !logs,
		StacktraceLevel: "panic",
	})
}
