package logic

import (
	context "context"

	v1 "github.com/Ccheers/protoc-gen-zerorpc/example/api/product/app/v1"
	svc "github.com/Ccheers/protoc-gen-zerorpc/example/internal/svc"
	logx "github.com/zeromicro/go-zero/core/logx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
type SvrStreamLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSvrStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SvrStreamLogic {
	return &SvrStreamLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *SvrStreamLogic) SvrStream(stream v1.BlogService_SvrStreamServer) error {
	panic("implement me")
}
