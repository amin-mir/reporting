//go:generate mockgen -source store.go -destination mock/store_mock.go -package mock

package reportstore

type Store interface {
	CreateReport(r CreateReportRequest) error
	UpdateReportStatus(r UpdateReportStatusRequest) error
	GetUserReport(r GetUserReportRequest) (GetUserReportResponse, error)
	GetUserReports(r GetUserReportsRequest) (GetUserReportsResponse, error)
	UserHasAccess(r UserHasAccessRequest) (UserHasAccessResponse, error)
	AppendMessage(r AppendMessageRequest) error
}

type CreateReportRequest struct {
	ReportID string
	UserID   string
	Status   string
	Title    string
}

type UpdateReportStatusRequest struct {
	ReportID string
	Status   ReportStatus
}

type GetUserReportRequest struct {
	UserID   string
	ReportID string
}

type GetUserReportResponse struct {
	Report Report
}

type GetUserReportsRequest struct {
	UserID string
}

type GetUserReportsResponse struct {
	Reports []Report
}

type UserHasAccessRequest struct {
	UserID   string
	ReportID string
}

type UserHasAccessResponse struct {
	HasAccess bool
}

type AppendMessageRequest struct {
	ReportID  string
	MessageID string
	Body      string
}
