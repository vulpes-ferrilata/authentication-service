package errors

type AppError interface {
	error
	WebError
	GrpcError
}
