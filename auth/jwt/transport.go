package jwt

//
//import (
//	"context"
//	"fmt"
//	"strings"
//
//	"google.golang.org/grpc/metadata"
//)
//
//const (
//	bearer       string = "bearer"
//	bearerFormat string = "Bearer %s"
//)
//
//// GRPCToContext moves a JWT from grpc metadata to context. Particularly
//// userful for servers.
//func GRPCToContext() grpc.ServerRequestFunc {
//	return func(ctx context.Context, md metadata.MD) context.Context {
//		// capital "Key" is illegal in HTTP/2.
//		authHeader, ok := md["authorization"]
//		if !ok {
//			return ctx
//		}
//
//		token, ok := extractTokenFromAuthHeader(authHeader[0])
//		if ok {
//			ctx = context.WithValue(ctx, JWTTokenContextKey, token)
//		}
//
//		return ctx
//	}
//}
//
//// ContextToGRPC moves a JWT from context to grpc metadata. Particularly
//// useful for clients.
//func ContextToGRPC() grpc.ClientRequestFunc {
//	return func(ctx context.Context, md *metadata.MD) context.Context {
//		token, ok := ctx.Value(JWTTokenContextKey).(string)
//		if ok {
//			// capital "Key" is illegal in HTTP/2.
//			(*md)["authorization"] = []string{generateAuthHeaderFromToken(token)}
//		}
//
//		return ctx
//	}
//}
//
//func extractTokenFromAuthHeader(val string) (token string, ok bool) {
//	authHeaderParts := strings.Split(val, " ")
//	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
//		return "", false
//	}
//
//	return authHeaderParts[1], true
//}
//
//func generateAuthHeaderFromToken(token string) string {
//	return fmt.Sprintf(bearerFormat, token)
//}
