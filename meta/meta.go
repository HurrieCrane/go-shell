package meta

import (
	"time"
)

// ParsingMeta contains meta information
// about a single parsing operation
// which is alive for a single parsing
// operation
type ParsingMeta struct {
	StartedParsing time.Time
	StoppedParsing time.Time
}

// StartRecording will record the current UTC time
func (m ParsingMeta) StartRecording() {
	m.StartedParsing = time.Now().UTC()
}

// StopRecording will record the current UTC time
func (m ParsingMeta) StopRecording() {
	m.StoppedParsing = time.Now().UTC()
}

// GetParsingTime returns the total time, in seconds,
// that it took to parse a block
func (m ParsingMeta) GetParsingTime() int {
	return m.StoppedParsing.Second() - m.StartedParsing.Second()
}
