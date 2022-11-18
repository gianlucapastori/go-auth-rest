package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func New() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("could not create logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	return sugar, nil
}
