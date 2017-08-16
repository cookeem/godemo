package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/clientselector"
	"github.com/smallnest/rpcx/codec"
	"github.com/smallnest/rpcx/log"
	"godemo/rpcx/gencode"
	"math/rand"
	"time"
)

var e = flag.String("e", "http://127.0.0.1:2379", "etcd URL")
var n = flag.String("n", "Arith", "Service Name")

func main() {
	flag.Parse()

	//basePath = "/rpcx/" + serviceName
	s := clientselector.NewEtcdV3ClientSelector([]string{*e}, "/rpcx/"+*n, time.Minute, rpcx.RoundRobin, time.Minute)
	client := rpcx.NewClient(s)
	client.FailMode = rpcx.Failover
	client.ClientCodecFunc = codec.NewGencodeClientCodec

	log.Info("client start")

	for i := 0; i < 10; i++ {
		begin := time.Now()
		a := int64(rand.Intn(100))
		b := int64(rand.Intn(100))
		args := &gencode.Args{a, b}
		var reply gencode.Reply
		err := client.Call(context.Background(), *n+".Mul", args, &reply)
		if err != nil {
			log.Infof("error for "+*n+": %d*%d, %v, took %v", args.A, args.B, err, time.Since(begin))
		} else {
			log.Infof(*n+": %d*%d=%d, took %v", args.A, args.B, reply.C, time.Since(begin))
		}
	}

	client.Close()
}
