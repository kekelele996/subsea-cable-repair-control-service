package service

import "context"

func RunMobilizeSession(ctx context.Context, spans []string) []string {
	child, cancel := context.WithCancel(ctx)
	results := make(chan string, len(spans))
	for _, span := range spans {
		span := span
		go func() {
			select {
			case results <- span:
			case <-child.Done():
			}
		}()
	}
	cancel()
	out := []string{}
	for i := 0; i < len(spans); i++ {
		select {
		case span := <-results:
			out = append(out, span)
		default:
			return out
		}
	}
	return out
}
