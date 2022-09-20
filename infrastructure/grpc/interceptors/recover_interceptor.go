package interceptors

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRecoverInterceptor() *RecoverInterceptor {
	return &RecoverInterceptor{}
}

type RecoverInterceptor struct{}

func (r RecoverInterceptor) ServerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				detail := fmt.Sprintf("%v", r)
				stacktrace := strings.Split(string(debug.Stack()), "\n")

				stt := status.New(codes.Internal, detail)

				debugInfo := &errdetails.DebugInfo{}
				debugInfo.Detail = detail
				debugInfo.StackEntries = append(debugInfo.StackEntries, stacktrace...)

				stt, _ = stt.WithDetails(debugInfo)

				err = stt.Err()
			}
		}()

		resp, err := handler(ctx, req)

		return resp, err
	}
}
