package util

import (
	"os"
	"os/signal"
	"syscall"

	"eos2git.cec.lab.emc.com/ECS/baremetal-csi-plugin.git/pkg/base/rpc"
)

//SetupSignalHandler set up channel for SIGTERM signal, when SIGTERM is caught function try to shutdown service
func SetupSignalHandler(server *rpc.ServerRunner) {
	sigint := make(chan os.Signal, 1)

	signal.Notify(sigint, syscall.SIGTERM)

	//Wait SIGTERM handler
	<-sigint

	// We received an interrupt signal, shut down.
	server.StopServer()
}