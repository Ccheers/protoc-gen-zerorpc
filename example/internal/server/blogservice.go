package server

import (
	context "context"
	v1 "github.com/Ccheers/protoc-gen-zerorpc/example/api/product/app/v1"
	logic "github.com/Ccheers/protoc-gen-zerorpc/example/internal/logic"
	svc "github.com/Ccheers/protoc-gen-zerorpc/example/internal/svc"
)

// This is a compile-time assertion to ensure that this generated file
type BlogServiceServerImpl struct {
	svcCtx *svc.ServiceContext
	v1.UnimplementedBlogServiceServer
}

func NewBlogServiceServer(svcCtx *svc.ServiceContext) v1.BlogServiceServer {
	return &BlogServiceServerImpl{
		svcCtx: svcCtx,
	}
}
func (x *BlogServiceServerImpl) GetArticles(ctx context.Context, in *v1.GetArticlesReq) (*v1.GetArticlesResp, error) {
	handler := logic.NewGetArticlesLogic(ctx, x.svcCtx)
	return handler.GetArticles(in)
}
func (x *BlogServiceServerImpl) CreateArticle(ctx context.Context, in *v1.Article) (*v1.Article, error) {
	handler := logic.NewCreateArticleLogic(ctx, x.svcCtx)
	return handler.CreateArticle(in)
}
func (x *BlogServiceServerImpl) SvrStream(stream v1.BlogService_SvrStreamServer) error {
	ctx := stream.Context()
	handler := logic.NewSvrStreamLogic(ctx, x.svcCtx)
	return handler.SvrStream(stream)
}
func (x *BlogServiceServerImpl) CliStream(in *v1.SvrStreamRequest, stream v1.BlogService_CliStreamServer) error {
	ctx := stream.Context()
	handler := logic.NewCliStreamLogic(ctx, x.svcCtx)
	return handler.CliStream(in, stream)
}
func (x *BlogServiceServerImpl) BothStream(stream v1.BlogService_BothStreamServer) error {
	ctx := stream.Context()
	handler := logic.NewBothStreamLogic(ctx, x.svcCtx)
	return handler.BothStream(stream)
}
