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

// MetaCacheCmd represents the MetaCache command
var MetaCacheCmd = &cobra.Command{
	Use:   "MetaCache",
	Short: "helps to call MO QueryService/MetaCache api",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		metaCache()
	},
}

func metaCache() {
	ctx := context.Background()
	client, err := setup.GetQueryClient(nil)
	if err != nil {
		os.Exit(1)
	}
	defer client.Close()
	logger := setup.GetLogger()

	addr := queryService.GetAddress()

	if err = CheckQueryServiceVersion(ctx, logger, client, MetadataCacheVersion); err != nil {
		os.Exit(1)
	}

	req := client.NewRequest(query.CmdMethod_MetadataCache)
	req.MetadataCacheRequest.CacheSize = *metaCacheConfig.Bytes
	deadlineCtx, dcCancel := context.WithTimeout(ctx, *queryService.Timeout)
	resp, err := client.SendMessage(deadlineCtx, addr, req)
	if err != nil {
		logger.Error("failed to request QueryService/MetadataCache", zap.Error(err))
		dcCancel()
		os.Exit(1)
	}
	logger.Info("MetadataCache cmd",
		zap.String("addr", addr),
		zap.Int64("req", req.MetadataCacheRequest.CacheSize),
		zap.Int64("resp", resp.MetadataCacheResponse.CacheCapacity),
	)
	client.Release(resp)
	dcCancel()
}

type MetaCacheConfig struct {
	Bytes *int64
}

var metaCacheConfig MetaCacheConfig

func init() {
	queryServiceCmd.AddCommand(MetaCacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MetaCacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	metaCacheConfig.Bytes = MetaCacheCmd.Flags().Int64("bytes", 2147483648, "calling queryService/MetaCache to change")
}
