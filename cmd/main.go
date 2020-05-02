package main

import (
	"fmt"
	"github.com/alessio-perugini/bao"
	"github.com/spf13/cobra"
)

var (
	version = "version"
)

func main() {
	config := new(bao.Config)

	var cmd = &cobra.Command{
		Use:   "bao",
		Short: "ip log address parser",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			bao.GetIpFromLog()
		},
		Version: fmt.Sprintf("%s", "version"),
	}

	cmd.Flags().StringVarP(&config.GeoIpDb, "geoip-db", "g", "", "geoipdb")
	cmd.Flags().StringVarP(&config.LinuxLog, "flog", "i", "", "log file to analyze")
	cmd.Flags().StringVarP(&config.OnlyIpFile, "fplain-ip", "p", "", "file to save blacklisted plain ip")
	cmd.Flags().StringVarP(&config.DetailedIpFile, "fdetailed-ip", "d", "", "file to save detail about blacklisted ip")
	cmd.Flags().StringVarP(&config.NationFilter, "filter", "f", "en", "type the country to whitelist separate with | ")
	bao.NewConfig(config)

	cmd.Execute()
}
