/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/spf13/cobra"
)

// pullMetricsCmd represents the pullMetrics command
var pullMetricsCmd = &cobra.Command{
	Use:   "pullMetrics",
	Short: "helps to pull metrics data from metric-given-server, through http://.../metrics url.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pullMetricOp()
	},
}

func pullMetricOp() {

	address := fmt.Sprintf("http://%s:%d/metrics", pullMetricsCfg.Host, pullMetricsCfg.Port)

	// 创建一个HTTP客户端
	client := &http.Client{}
	// 创建一个HTTP请求，用于获取指标数据
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	// 使用promhttp包来处理响应并获取指标数据
	//metricFamilies, err := promhttp.ParseMetricFamilies(resp.Body)
	//if err != nil {
	//	fmt.Println("Error parsing metrics:", err)
	//	return
	//}

	parser := &expfmt.TextParser{}
	metricFamilies, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		fmt.Println("Error parsing metrics:", err)
		return
	}
	for name, metricFamily := range metricFamilies {
		fmt.Printf("Metric Name: %s\n", name)
		switch metricFamily.GetType() {
		case dto.MetricType_COUNTER:
			for _, metric := range metricFamily.GetMetric() {
				counter := metric.GetCounter()
				fmt.Printf("  Counter Value: %v, \tLabels: %v\n", counter.GetValue(), metric.GetLabel())
			}
		case dto.MetricType_GAUGE:
			for _, metric := range metricFamily.GetMetric() {
				gauge := metric.GetGauge()
				fmt.Printf("  Gauge Value: %v, \tLabels: %v\n", gauge.GetValue(), metric.GetLabel())
			}
			// 可以根据指标类型（如Histogram、Summary等）添加更多的处理逻辑
		}
	}
}

type pullMetricsConfig struct {
	Port    int
	Host    string
	Timeout time.Duration
}

var pullMetricsCfg pullMetricsConfig

func init() {
	rootCmd.AddCommand(pullMetricsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullMetricsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullMetricsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pullMetricsCmd.Flags().IntVar(&pullMetricsCfg.Port, "port", 7001, "port to access the metrics server")
	pullMetricsCmd.Flags().StringVar(&pullMetricsCfg.Host, "host", "127.0.0.1", "host to access the metrics server")
}
