package server

import (
	"context"
	"fmt"
	"github.com/2336260845/hippo_recall/config"
	"github.com/2336260845/hippo_recall/gen-go/recall"
	"github.com/apache/thrift/lib/go/thrift"
	"strconv"
)

type RecallServer struct{}

func (qas *RecallServer) Recall(ctx context.Context, req *recall.RecallParam) (r []*recall.Doc, err error) {
	if req.Query == "" {
		return nil, fmt.Errorf("query is empty")
	}

	if config.GetConfig().Debug {
		for i := 0; i < 10; i++ {
			r = append(r, &recall.Doc{Title: "测试文章:"+strconv.Itoa(i+1), Summary: "test summary"})
		}
	}

	//TODO 实现召回函数
	return r, nil
}

func ThriftInit(conf *config.Config) {
	transport, err := thrift.NewTServerSocket(conf.ThriftAddress)
	if err != nil {
		panic(fmt.Sprintf("ThriftInit NewTServerSocket error, err=%+v", err.Error()))
	}

	handler := &RecallServer{}
	processor := recall.NewRecallServiceProcessor(handler)
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(fmt.Sprintf("ThriftInit thrift Serve error, err=%+v", err.Error()))
	}
}
