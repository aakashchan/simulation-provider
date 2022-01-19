package main

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	cli "github.com/virtual-kubelet/node-cli"
	logruslogger "github.com/virtual-kubelet/virtual-kubelet/log/logrus"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"github.com/virtual-kubelet/virtual-kubelet/trace/opencensus"
)

var (
	buildVersion    = "N/A"
	buildTime       = "N/A"
	k8sVersion      = "v1.19.10" // This should follow the version of k8s.io/client-go we are importing
	numberOfWorkers = 50
)
func main() {
	ctx := cli.ContextWithCancelOnSignal(context.Background())

	logger := logrus.StandardLogger()
	log.L = logruslogger.FromLogrus(logrus.NewEntry(logger))
	logConfig := &logruscli.Config{LogLevel: "info"}

	trace.T = opencensus.Adapter{}
	traceConfig := opencensuscli.Config{
		AvailableExporters: map[string]opencensuscli.ExporterInitFunc{
			"ocagent": initOCAgent,
		},
	}


}
