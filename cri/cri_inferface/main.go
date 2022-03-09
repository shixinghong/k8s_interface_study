package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"k8s.io/klog/v2"
)

func main() {

	//grpc.WithTransportCredentials(insecure.NewCredentials())
	//opts := []grpc.DialOption{
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//}

	addr := "unix:///run/containerd/containerd.sock"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		klog.Exit(err)
	}

	req := v1alpha2.VersionRequest{}
	rsp := v1alpha2.VersionResponse{}
	err = conn.Invoke(context.Background(), "/runtime.v1alpha2.RuntimeService/Version", &req, &rsp)
	if err != nil {
		klog.Exit(err)
	}
	fmt.Println(rsp)
}
