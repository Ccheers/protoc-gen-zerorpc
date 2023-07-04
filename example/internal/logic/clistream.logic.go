package logic

import (
	context "context"

	v1 "github.com/Ccheers/protoc-gen-zerorpc/example/api/product/app/v1"
	svc "github.com/Ccheers/protoc-gen-zerorpc/example/internal/svc"
	logx "github.com/zeromicro/go-zero/core/logx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
type CliStreamLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCliStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CliStreamLogic {
	return &CliStreamLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *CliStreamLogic) CliStream(req *v1.SvrStreamRequest, stream v1.BlogService_CliStreamServer) error {
	panic("implement me")
}
