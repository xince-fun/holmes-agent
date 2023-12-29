package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/xince-fun/holmes-agent/pkg/logger"
	"github.com/xince-fun/holmes-agent/pkg/machine"
	"golang.org/x/mod/semver"
)

var (
	Version                 = "main"
	minKernelVersionSupport = "4.16"
	kernelVersionRe         = regexp.MustCompile(`^(\d+\.\d+)`)
)

func compareKernelVersion(kernelVersion string) string {
	return kernelVersionRe.FindString(kernelVersion)
}

// main entry point
func main() {
	log := logger.DefaultLogger.Named("holmes")

	log.Infof("Version: %s, staring Holmes Ebpf Agent ... ", Version)
	pprofPort := 6600

	// we can visit /debug/pprof
	go func() {
		log.Infof("starting PProf HTTP listener, port: %d", pprofPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", pprofPort), nil)
		log.Error("PProf HTTP listener stopped working, %w", err)
	}()

	hostName, kernelVersion, err := machine.Uname()
	if err != nil {
		log.Fatal("failed to get uname from machine: ", err)
	}
	machineId := machine.MachineId()
	log.Infof("HostName: %s KernelVersion: %s MachineId: %s", hostName, kernelVersion, machineId)

	compareVersion := compareKernelVersion(kernelVersion)
	if compareVersion == "" {
		log.Fatal("read invalid kernel version: ", compareVersion)
	}
	if semver.Compare("v"+compareVersion, "v"+minKernelVersionSupport) == -1 {
		log.Fatalf("can't support KernelVersion %s, the minimum Kernel Version required is %s or later", minKernelVersionSupport)
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-ctx.Done()
}
