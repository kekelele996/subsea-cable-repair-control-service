package service

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/store"

type CompleteCoordinator struct{ Events []string }

func (c *CompleteCoordinator) Complete(repo *store.CompleteRepository, expectedRevision int) error {
	c.Events = append(c.Events, "completed:"+repo.Plan.ID)
	return repo.Commit(expectedRevision)
}
