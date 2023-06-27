package main

import (
	"context"

	pb "github.com/GoogleCloudPlatform/microservices-demo/src/frontend/genproto"
)

func (fe *apiServer) getProducts(ctx context.Context) ([]*pb.Product, error) {
	resp, err := pb.NewProductCatalogServiceClient(fe.productCatalogSvcConn).
		ListProducts(ctx, &pb.Empty{})
	return resp.GetProducts(), err
}
