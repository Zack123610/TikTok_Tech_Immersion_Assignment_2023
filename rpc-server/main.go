package main

import (
	"log"
	// "context"
	// "encoding/json"

	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB

	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
