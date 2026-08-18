package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/domain"

func MobilizeManifest(snapshot domain.ManifestSnapshot) <-chan domain.ManifestSnapshot {
	out := make(chan domain.ManifestSnapshot, 1)
	go func() {
		out <- domain.CopyManifestSnapshot(snapshot)
		close(out)
	}()
	return out
}
