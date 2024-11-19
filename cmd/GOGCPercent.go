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

// GOGCPercentCmd represents the GOGCPercent command
var GOGCPercentCmd = &cobra.Command{
	Use:   "GOGCPercent",
	Short: "helps to call MO QueryService/GOGCPercent api",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goGCPercent()
	},
}

func goGCPercent() {
	ctx := context.Background()
	client, err := setup.GetQueryClient(nil)
	if err != nil {
		os.Exit(1)
	}
	defer client.Close()
	logger := setup.GetLogger()

	addr := queryService.GetAddress()

	if err = CheckQueryServiceVersion(ctx, logger, client, GOGCPercentVersion); err != nil {
		os.Exit(1)
	}

	req := client.NewRequest(query.CmdMethod_GOGCPercent)
	req.GoGCPercentRequest.Percent = *goGCPercentConfig.Percent
	deadlineCtx, dcCancel := context.WithTimeout(ctx, *queryService.Timeout)
	resp, err := client.SendMessage(deadlineCtx, addr, req)
	if err != nil {
		logger.Error("failed to request QueryService/GoGCPercent", zap.Error(err))
		dcCancel()
		os.Exit(1)
	}
	logger.Info("GOGCPercent cmd",
		zap.String("addr", addr),
		zap.Int32("req", req.GoGCPercentRequest.Percent),
		zap.Int32("resp", resp.GoGCPercentResponse.Percent),
	)
	client.Release(resp)
	dcCancel()
}

type GOGCPercentConfig struct {
	Percent *int32
}

var goGCPercentConfig GOGCPercentConfig

func init() {
	queryServiceCmd.AddCommand(GOGCPercentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GOGCPercentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	goGCPercentConfig.Percent = GOGCPercentCmd.Flags().Int32("percent", 100, "Call GOGCPercent, set the target percent")
}
