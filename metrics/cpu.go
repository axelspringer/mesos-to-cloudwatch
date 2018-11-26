package metrics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	. "github.com/axelspringer/mesos-to-cloudwatch/services"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/urfave/cli"
)

// Collect Swap usage
func (d Mesos) Collect(instanceID string, c CloudWatchService, cli *cli.Context) {
	servers := strings.Split(cli.String("master"), ",")
	var h = &http.Client{Timeout: 10 * time.Second}

	for i := range servers {

		url := servers[i] + "/metrics/snapshot"

		r, err := h.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)

		var data Mesos
		err = decoder.Decode(&data)

		if err != nil {
			fmt.Println("%T\n%s\n%#v\n", err, err, err)
		}

		if data.MasterElected != 0 {
			fmt.Println("Server is Master", servers[i])
			mesosMetrics := (data.MasterCpusUsed * 100) / float64(data.MasterCpusTotal)

			dimensionKey := "mesos"
			dimensions := []cloudwatch.Dimension{
				cloudwatch.Dimension{
					Name:  &dimensionKey,
					Value: &instanceID,
				},
			}

			mesosCpusUsedData := constructMetricDatum("CpusUsed", float64(mesosMetrics), cloudwatch.StandardUnitPercent, dimensions)
			c.Publish(mesosCpusUsedData, cli.String("namespace"))

			log.Printf("CPUs - Utilization:%v%% \n", mesosMetrics)
			break
		}

		fmt.Println("Server is no Master: ", servers[i])
		continue
	}

}
