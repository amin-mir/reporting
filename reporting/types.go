package reporting

import "github.com/amin-mir/reporting/reportstore"

type CreateReportRequest struct {
	UserID string
	Title  string
}

type CreateReportResponse struct {
	ReportID string
}

type AppendMessageRequest struct {
	UserID   string
	ReportID string
	Body     string
}

type AppendMessageResponse struct {
	MessageID string
}

type UpdateReportStatusRequest struct {
	UserID   string
	ReportID string
	Status   string
}

type UpdateReportStatusResponse struct{}

type GetUserReportsRequest struct {
	UserID string
}

type GetUserReportsResponse struct {
	Reports []reportstore.Report
}
