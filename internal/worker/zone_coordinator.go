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
	out := make([]string, 0, len(spans))
	for len(out) < len(spans) {
		select {
		case span := <-results:
			out = append(out, span)
		case <-ctx.Done():
			return out
		}
	}
	return out
}
