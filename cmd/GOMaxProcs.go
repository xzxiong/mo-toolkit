/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/matrixorigin/matrixone/pkg/pb/query"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/xzxiong/mo-toolkit/pkg/setup"
)

// GOMaxProcsCmd represents the GOMaxProcs command
var GOMaxProcsCmd = &cobra.Command{
	Use:   "GOMaxProcs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goMaxProcs()
	},
}

func goMaxProcs() {
	ctx := context.Background()
	client, err := setup.GetQueryClient(nil)
	if err != nil {
		os.Exit(1)
	}
	defer client.Close()
	logger := setup.GetLogger()

	addr := queryService.GetAddress()

	// Get Version
	req := client.NewRequest(query.CmdMethod_GetProtocolVersion)
	req.GetProtocolVersion = &query.GetProtocolVersionRequest{}
	deadlineCtx, dcCancel := context.WithTimeout(ctx, *queryService.Timeout)
	resp, err := client.SendMessage(deadlineCtx, addr, req)
	if err != nil {
		logger.Error("failed to request QueryService", zap.Error(err))
		dcCancel()
		os.Exit(1)
	}
	if resp.GetProtocolVersion.Version < GOMaxProcsVersion {
		logger.Error("target mo query service is too old", zap.String("addr", addr),
			zap.Int64("version", resp.GetProtocolVersion.Version),
			zap.Int64("target", GOMaxProcsVersion),
		)
		os.Exit(1)
	}
	logger.Info("GetProtocolVersion",
		zap.String("addr", addr), zap.Int64("version", resp.GetProtocolVersion.Version))
	client.Release(resp)
	dcCancel()

	// query GOMAXPROCS
	req = client.NewRequest(query.CmdMethod_GOMAXPROCS)
	req.GoMaxProcsRequest.MaxProcs = *goMaxProcsConfig.Value
	deadlineCtx, dcCancel = context.WithTimeout(ctx, *queryService.Timeout)
	resp, err = client.SendMessage(deadlineCtx, addr, req)
	if err != nil {
		logger.Error("failed to request QueryService/GOMAXPROCS", zap.Error(err))
		dcCancel()
		os.Exit(1)
	}
	logger.Info("GOMAXPROCS query",
		zap.String("addr", addr),
		zap.Int32("MaxProcs", resp.GoMaxProcsResponse.MaxProcs),
	)
	client.Release(resp)
	dcCancel()
}

type GoMaxProcsCmdConfig struct {
	Value *int32
}

var goMaxProcsConfig GoMaxProcsCmdConfig

func init() {
	queryServiceCmd.AddCommand(GOMaxProcsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GOMaxProcsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GOMaxProcsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	goMaxProcsConfig.Value = GOMaxProcsCmd.Flags().Int32P("value", "v", 0, "call mo query-service/GOMaxProcs with spec ")
}
