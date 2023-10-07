package gomock

import (
	"github.com/auvitly/assistant/behaviour/internal/gomock/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock.go -package mock

type Service struct {
	Setter Setter
	Getter Getter
}

type Setter interface {
	Set(args ...any) error
}

type Getter interface {
	Get() ([]any, error)
}

func Test(t *testing.T) {
	var ctrl = gomock.NewController(t)
	setter := mock.NewMockSetter(ctrl)
	setter.EXPECT()
}
