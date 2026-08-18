package domain

import "time"

type RepairState string

const (
	RepairDraft     RepairState = "draft"
	RepairScheduled RepairState = "queued"
	RepairExecuting RepairState = "running"
	RepairAborted   RepairState = "rejected"
	RepairCompleted RepairState = "completed"
)

type RepairPlan struct {
	ID        string
	CableID   string
	Spans     []string
	State     RepairState
	Revision  int
	CreatedAt time.Time
}

type RepairJob struct {
	ID      string
	PlanID  string
	Spans   []string
	Attempt int
}
type Lease struct {
	CableID   string
	Owner     string
	ExpiresAt time.Time
}
type Hazard struct {
	Code     string
	Severity string
	Message  string
}
type CompletionDecision struct {
	PlanID  string
	Allowed bool
	Reason  string
}

type SafetyProvider interface {
	Evaluate(RepairPlan) ([]Hazard, error)
}
