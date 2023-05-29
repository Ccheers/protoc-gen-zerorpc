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
