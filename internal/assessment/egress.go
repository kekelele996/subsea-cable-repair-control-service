package assessment

import (
	"sort"
	"strings"
)

type EgressAssessment struct {
	CableID string
	Window  string
	Spans   []string
	Signals map[string]float64
	Notes   []string
}

type EgressResult struct {
	Score      int
	Band       string
	Review     bool
	Reasons    []string
	Normalized []string
}

func NewEgressAssessment(cable, window string, spans []string, signals map[string]float64) EgressAssessment {
	copied := make([]string, len(spans))
	copy(copied, spans)
	signalCopy := map[string]float64{}
	for key, value := range signals {
		signalCopy[strings.ToLower(strings.TrimSpace(key))] = value
	}
	return EgressAssessment{CableID: strings.TrimSpace(cable), Window: strings.TrimSpace(window), Spans: copied, Signals: signalCopy}
}

func (a EgressAssessment) NormalizeSpans() []string {
	out := make([]string, 0, len(a.Spans))
	seen := map[string]bool{}
	for _, span := range a.Spans {
		value := strings.ToLower(strings.TrimSpace(span))
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}

func (a EgressAssessment) Signal(name string) float64 {
	return a.Signals[strings.ToLower(strings.TrimSpace(name))]
}
func (a *EgressAssessment) AddNote(note string) {
	if strings.TrimSpace(note) != "" {
		a.Notes = append(a.Notes, strings.TrimSpace(note))
	}
}

func (a EgressAssessment) Score() int {
	score := 0
	for _, span := range a.NormalizeSpans() {
		score += len(span) % 4
		if strings.Contains(span, "restricted") {
			score += 5
		}
		if strings.Contains(span, "critical") {
			score += 8
		}
	}
	for _, value := range a.Signals {
		if value < 0 {
			score += 10
		} else if value > 90 {
			score += 6
		} else if value > 70 {
			score += 3
		}
	}
	if a.Window == "" {
		score += 4
	}
	if a.CableID == "" {
		score += 7
	}
	return score
}

func (a EgressAssessment) Result() EgressResult {
	score := a.Score()
	band := "normal"
	if score >= 18 {
		band = "elevated"
	}
	if score >= 30 {
		band = "high"
	}
	reasons := make([]string, 0, 3)
	for _, span := range a.NormalizeSpans() {
		if strings.Contains(span, "critical") {
			reasons = append(reasons, "critical span")
		}
	}
	if a.Window == "" {
		reasons = append(reasons, "missing weather window")
	}
	if len(a.NormalizeSpans()) == 0 {
		reasons = append(reasons, "no repair spans")
	}
	return EgressResult{Score: score, Band: band, Review: score >= 18, Reasons: reasons, Normalized: a.NormalizeSpans()}
}

func (a EgressAssessment) Merge(other EgressAssessment) EgressAssessment {
	out := NewEgressAssessment(a.CableID, a.Window, a.Spans, a.Signals)
	if out.CableID == "" {
		out.CableID = other.CableID
	}
	if out.Window == "" {
		out.Window = other.Window
	}
	out.Spans = append(out.Spans, other.Spans...)
	for key, value := range other.Signals {
		out.Signals[key] = value
	}
	out.Notes = append(out.Notes, other.Notes...)
	return out
}

func (a EgressAssessment) Summary() string {
	r := a.Result()
	return strings.Join([]string{a.CableID, a.Window, r.Band, strings.Join(r.Reasons, "|")}, "/")
}
func (a EgressAssessment) RequiresSupervisor() bool {
	r := a.Result()
	return r.Review || r.Band == "high"
}
func (a EgressAssessment) StableKey() string {
	r := a.Result()
	return strings.Join([]string{a.CableID, a.Window, strings.Join(r.Normalized, ",")}, "#")
}
func (a EgressAssessment) WithSignal(name string, value float64) EgressAssessment {
	out := NewEgressAssessment(a.CableID, a.Window, a.Spans, a.Signals)
	out.Signals[strings.ToLower(strings.TrimSpace(name))] = value
	out.Notes = append(out.Notes, a.Notes...)
	return out
}
