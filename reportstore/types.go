package reportstore

import "time"

type ReportStatus int

const (
	ReportStatusUnknown ReportStatus = iota
	ReportStatusPending
	ReportStatusReviewing
	ReportStatusResolved
)

func (s ReportStatus) String() string {
	switch s {
	case ReportStatusPending:
		return "pending"
	case ReportStatusReviewing:
		return "reviewing"
	case ReportStatusResolved:
		return "resolved"
	default:
		return ""
	}
}

func ParseReportStatus(s string) ReportStatus {
	switch s {
	case "pending":
		return ReportStatusPending
	case "reviewing":
		return ReportStatusReviewing
	case "resolved":
		return ReportStatusResolved
	default:
		return ReportStatusUnknown
	}
}

type Report struct {
	ReportID    string
	UserID      string
	ResolverID  string
	ReviewerIDs []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      ReportStatus
	Title       string
}

type Message struct {
	ReportID  string
	MessageID string
	Body      string
	CreatedAt time.Time
}
