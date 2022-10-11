package logger

import (
	"strconv"
	"strings"
)

// Severity is the severity of the event described in a log entry. These
// guideline severity levels are ordered, with numerically smaller levels
// treated as less severe than numerically larger levels.
type Severity int

const (
	// SeverityDefault means the log entry has no assigned severity level.
	SeverityDefault = Severity(0)
	// Debug means debug or trace information.
	SeverityDebug = Severity(100)
	// Info means routine information, such as ongoing status or performance.
	SeverityInfo = Severity(200)
	// Notice means normal but significant events, such as start up, shut down, or configuration.
	SeverityNotice = Severity(300)
	// Warning means events that might cause problems.
	SeverityWarning = Severity(400)
	// Error means events that are likely to cause problems.
	SeverityError = Severity(500)
	// Critical means events that cause more severe problems or brief outages.
	SeverityCritical = Severity(600)
	// Alert means a person must take an action immediately.
	SeverityAlert = Severity(700)
	// Emergency means one or more systems are unusable.
	SeverityEmergency = Severity(800)
)

var severityName = map[Severity]string{
	SeverityDefault:   "SeverityDefault",
	SeverityDebug:     "Debug",
	SeverityInfo:      "Info",
	SeverityNotice:    "Notice",
	SeverityWarning:   "Warning",
	SeverityError:     "Error",
	SeverityCritical:  "Critical",
	SeverityAlert:     "Alert",
	SeverityEmergency: "Emergency",
}

// String converts a severity level to a string.
func (v Severity) String() string {
	// same as proto.EnumName
	s, ok := severityName[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// ParseSeverity returns the Severity whose name equals s, ignoring case. It
// returns SeverityDefault if no Severity matches.
func ParseSeverity(s string) Severity {
	sl := strings.ToLower(s)
	for sev, name := range severityName {
		if strings.ToLower(name) == sl {
			return sev
		}
	}
	return SeverityDefault
}
