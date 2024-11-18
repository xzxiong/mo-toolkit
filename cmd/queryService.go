/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/matrixorigin/matrixone/pkg/defines"
	"github.com/spf13/cobra"
)

// queryServiceCmd represents the queryService command
var queryServiceCmd = &cobra.Command{
	Use:   "queryService",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("queryService called")
	},
}

type QueryServiceConfig struct {
	Port    *int
	Host    *string
	Timeout *time.Duration
}

func (c QueryServiceConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", *c.Host, *c.Port)
}

var queryService QueryServiceConfig

const (
	GOMaxProcsVersion = defines.MORPCMinVersion
)

func init() {
	rootCmd.AddCommand(queryServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryServiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryServiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	queryService.Port = queryServiceCmd.Flags().Int("query-service-port", 18002, "mo query-service port, example: 6004 (in cloud)")
	queryService.Host = queryServiceCmd.Flags().String("query-service-host", "127.0.0.1", "mo query-service host.")
	queryService.Timeout = queryServiceCmd.Flags().Duration("query-service-timeout", 3*time.Second, "timeout")
}
