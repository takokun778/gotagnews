package port

import "github.com/takokun778/gotagnews/internal/usecase"

type GotagNoticeInput struct {
	usecase.Input
}

type GotagNoticeOutput struct {
	usecase.Output
}

type GotagNoticeUsecase interface {
	usecase.Usecase[GotagNoticeInput, GotagNoticeOutput]
}
