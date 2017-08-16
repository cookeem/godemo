package main

import (
	"context"
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/codec"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/plugin"
	"godemo/rpcx/gencode"
	"time"
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *gencode.Args, reply *gencode.Reply) error {
	begin := time.Now()
	reply.C = args.A * args.B
	log.Infof("%d*%d=%d, took %v", args.A, args.B, reply.C, time.Since(begin))
	return nil
}

func (t *Arith) Error(args *gencode.Args, reply *gencode.Reply) error {
	panic("ERROR")
}

var addr = flag.String("s", "127.0.0.1:8972", "service address")
var e = flag.String("e", "http://127.0.0.1:2379", "etcd URL")
var n = flag.String("n", "Arith", "Service Name")

func main() {
	flag.Parse()

	log.Info("server start")
	server := rpcx.NewServer()
	server.ServerCodecFunc = codec.NewGencodeServerCodec
	rplugin := &plugin.EtcdV3RegisterPlugin{
		ServiceAddress:      "tcp@" + *addr,
		EtcdServers:         []string{*e},
		BasePath:            "/rpcx",
		Metrics:             metrics.NewRegistry(),
		UpdateIntervalInSec: 20,
	}
	rplugin.Start()
	server.PluginContainer.Add(rplugin)
	server.PluginContainer.Add(plugin.NewMetricsPlugin())
	server.RegisterName(*n, new(Arith), "weight=1&m=devops")
	server.Serve("tcp", *addr)
}
