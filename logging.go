package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(s Service) *LoggingService {
	return &LoggingService{next: s}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("GetCatFact took %v\n Err=%s\n Fact=%v\n", time.Since(start), err, fact.Fact)
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
