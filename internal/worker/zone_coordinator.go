package worker

import "context"

func CoordinateSpans(ctx context.Context, spans []string) []string {
	results := make(chan string, len(spans))
	for _, span := range spans {
		span := span
		go func() {
			select {
			case results <- span:
			case <-ctx.Done():
			}
		}()
	}
	select {
	case first := <-results:
		return []string{first}
	case <-ctx.Done():
		return nil
	}
}
