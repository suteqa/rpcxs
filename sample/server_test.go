package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/suteqa/rpcxs/server"
	"github.com/suteqa/rpcxs/serverplugin"
	"testing"
	"time"
)

func TestServder1(t *testing.T) {
	s := server.NewServer()

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@127.0.0.1:8973",
		EtcdServers:    []string{"127.0.0.1:2379"},
		BasePath:       "/rpcx_test",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	err := r.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Plugins.Add(r)

	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", "127.0.0.1:8973")

	defer s.Close()

	if len(r.Services) != 1 {
		fmt.Println("failed to register services in etcd")
		return
	}

	if err := r.Stop(); err != nil {
		fmt.Println(err)
		return
	}
}
