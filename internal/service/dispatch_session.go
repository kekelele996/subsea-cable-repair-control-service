package service

import (
	"context"
	"sync"
)

func RunMobilizeSession(ctx context.Context, spans []string) []string {
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan string, len(spans))
	var workers sync.WaitGroup
	workers.Add(len(spans))
	for _, span := range spans {
		span := span
		go func() {
			defer workers.Done()
			select {
			case results <- span:
			case <-child.Done():
			}
		}()
	}
	go func() { workers.Wait(); close(results) }()
	out := make([]string, 0, len(spans))
	for span := range results {
		out = append(out, span)
	}
	return out
}
