package logic

import (
	context "context"

	v1 "github.com/Ccheers/protoc-gen-zerorpc/example/api/product/app/v1"
	svc "github.com/Ccheers/protoc-gen-zerorpc/example/internal/svc"
	logx "github.com/zeromicro/go-zero/core/logx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the Ccheers/protoc-gen-zeroapi package it is being compiled against.
type BothStreamLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBothStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BothStreamLogic {
	return &BothStreamLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *BothStreamLogic) BothStream(stream v1.BlogService_BothStreamServer) error {
	panic("implement me")
}
