package client

import (
	context "context"
	v1 "github.com/Ccheers/protoc-gen-zerorpc/example/api/product/app/v1"
	zrpc "github.com/zeromicro/go-zero/zrpc"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
type BlogServiceClientImpl struct {
	cli zrpc.Client
}

func NewBlogServiceClient(cli zrpc.Client) v1.BlogServiceClient {
	return &BlogServiceClientImpl{
		cli: cli,
	}
}
func (x *BlogServiceClientImpl) GetArticles(ctx context.Context, in *v1.GetArticlesReq, opts ...grpc.CallOption) (*v1.GetArticlesResp, error) {
	client := v1.NewBlogServiceClient(x.cli.Conn())
	return client.GetArticles(ctx, in, opts...)
}
func (x *BlogServiceClientImpl) CreateArticle(ctx context.Context, in *v1.Article, opts ...grpc.CallOption) (*v1.Article, error) {
	client := v1.NewBlogServiceClient(x.cli.Conn())
	return client.CreateArticle(ctx, in, opts...)
}
func (x *BlogServiceClientImpl) SvrStream(ctx context.Context, opts ...grpc.CallOption) (v1.BlogService_SvrStreamClient, error) {
	client := v1.NewBlogServiceClient(x.cli.Conn())
	return client.SvrStream(ctx, opts...)
}
func (x *BlogServiceClientImpl) CliStream(ctx context.Context, in *v1.SvrStreamRequest, opts ...grpc.CallOption) (v1.BlogService_CliStreamClient, error) {
	client := v1.NewBlogServiceClient(x.cli.Conn())
	return client.CliStream(ctx, in, opts...)
}
func (x *BlogServiceClientImpl) BothStream(ctx context.Context, opts ...grpc.CallOption) (v1.BlogService_BothStreamClient, error) {
	client := v1.NewBlogServiceClient(x.cli.Conn())
	return client.BothStream(ctx, opts...)
}
