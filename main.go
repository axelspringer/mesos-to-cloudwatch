package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/axelspringer/mesos-to-cloudwatch/metrics"
	. "github.com/axelspringer/mesos-to-cloudwatch/services"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/urfave/cli"
)

// Collect metrics about enabled metric
func Collect(metrics []Metric, c CloudWatchService, cli *cli.Context) {
	id := "tortuga-mesos-master"
	for _, metric := range metrics {
		metric.Collect(id, c, cli)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "mesos"
	app.Usage = "Publish Custom Metrics to CloudWatch"
	app.Version = "1.0.0"
	app.Author = "me and you"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "mesos",
			Usage: "Collect mesos metrics",
		},
		cli.StringFlag{
			Name:  "master",
			Usage: "List of Mesos Master Server",
			Value: "http://127.0.0.1:5050",
		},
		cli.StringFlag{
			Name:  "region",
			Usage: "AWS region",
			Value: "us-west-1",
		},
		cli.IntFlag{
			Name:  "interval",
			Usage: "Time interval",
			Value: 5,
		},
		cli.BoolFlag{
			Name:  "once",
			Usage: "Run once (i.e. not on an interval)",
		},
		cli.StringFlag{
			Name:  "namespace",
			Usage: "Namespace for the metric data",
			Value: "CustomMetrics",
		},
	}
	app.Action = func(c *cli.Context) error {
		enabledMetrics := make([]string, 0)
		metrics := make([]Metric, 0)
		if c.Bool("mesos") {
			metrics = append(metrics, Mesos{})
			enabledMetrics = append(enabledMetrics, "mesos")
		}

		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			panic("Unable to load SDK config")
		}

		cfg.Region = c.String("region")
		cloudWatch := CloudWatchService{
			Config: cfg,
		}

		interval := c.Int("interval")

		fmt.Printf("Features enabled: %s\n", strings.Join(enabledMetrics, ", "))

		var collect = func() {
			Collect(metrics, cloudWatch, c)
		}

		if c.Bool("once") {
			collect()
		} else {
			fmt.Printf("Interval: %d Seconds\n", interval)

			duration := time.Duration(interval) * time.Second
			for range time.Tick(duration) {
				collect()
			}
		}

		return nil
	}
	app.Run(os.Args)
}
