package md

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"

	"gim/pkg/gerrors"
)

const (
	CtxUserID    = "user_id"
	CtxDeviceID  = "device_id"
	CtxToken     = "token"
	CtxRequestID = "request_id"
)

func ContextWithRequestId(ctx context.Context, requestId int64) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(CtxRequestID, strconv.FormatInt(requestId, 10)))
}

func Get(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values, ok := md[key]
	if !ok || len(values) == 0 {
		return ""
	}
	return values[0]
}

// GetCtxRequestId 获取ctx的app_id
func GetCtxRequestId(ctx context.Context) int64 {
	requestIdStr := Get(ctx, CtxRequestID)
	requestId, err := strconv.ParseInt(requestIdStr, 10, 64)
	if err != nil {
		return 0
	}
	return requestId
}

// GetCtxData 获取ctx的用户数据，依次返回user_id,device_id
func GetCtxData(ctx context.Context) (uint64, uint64, error) {
	var (
		userId   uint64
		deviceId uint64
		err      error
	)

	userIdStr := Get(ctx, CtxUserID)
	userId, err = strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return 0, 0, gerrors.ErrUnauthorized
	}

	deviceIdStr := Get(ctx, CtxDeviceID)
	deviceId, err = strconv.ParseUint(deviceIdStr, 10, 64)
	if err != nil {
		return 0, 0, gerrors.ErrUnauthorized
	}
	return userId, deviceId, nil
}

// GetCtxToken 获取ctx的token
func GetCtxToken(ctx context.Context) string {
	return Get(ctx, CtxToken)
}

// NewAndCopyRequestId 创建一个context,并且复制RequestId
func NewAndCopyRequestId(ctx context.Context) context.Context {
	newCtx := context.TODO()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return newCtx
	}

	requestIds, ok := md[CtxRequestID]
	if !ok && len(requestIds) == 0 {
		return newCtx
	}
	return metadata.NewOutgoingContext(newCtx, metadata.Pairs(CtxRequestID, requestIds[0]))
}
