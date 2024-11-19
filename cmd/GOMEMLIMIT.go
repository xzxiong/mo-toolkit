/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"math"
	"os"

	"github.com/matrixorigin/matrixone/pkg/pb/query"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/xzxiong/mo-toolkit/pkg/setup"
)

// GOMEMLimitCmd represents the GOMEMLimit command
var GOMEMLimitCmd = &cobra.Command{
	Use:   "GOMEMLimit",
	Short: "helps to call MO QueryService/GOMEMLimit api",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goMemLimit()
	},
}

func goMemLimit() {
	ctx := context.Background()
	client, err := setup.GetQueryClient(nil)
	if err != nil {
		os.Exit(1)
	}
	defer client.Close()
	logger := setup.GetLogger()

	addr := queryService.GetAddress()

	if err = CheckQueryServiceVersion(ctx, logger, client, GOMEMLimitVersion); err != nil {
		os.Exit(1)
	}

	req := client.NewRequest(query.CmdMethod_GOMEMLIMIT)
	req.GoMemLimitRequest.MemLimitBytes = *goMemLimitConfig.Bytes
	deadlineCtx, dcCancel := context.WithTimeout(ctx, *queryService.Timeout)
	resp, err := client.SendMessage(deadlineCtx, addr, req)
	if err != nil {
		logger.Error("failed to request QueryService/GOMEMLIMIT", zap.Error(err))
		dcCancel()
		os.Exit(1)
	}
	logger.Info("GOMEMLIMIT cmd",
		zap.String("addr", addr),
		zap.Int64("req", req.GoMemLimitRequest.MemLimitBytes),
		zap.Int64("resp", resp.GoMemLimitResponse.MemLimitBytes),
	)
	client.Release(resp)
	dcCancel()
}

type GOMEMLimitConfig struct {
	Bytes *int64
}

var goMemLimitConfig GOMEMLimitConfig

func init() {
	queryServiceCmd.AddCommand(GOMEMLimitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GOMEMLimitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	goMemLimitConfig.Bytes = GOMEMLimitCmd.Flags().Int64("bytes", math.MaxInt64, "call GOMEMLIMIT")
}
