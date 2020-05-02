//go:generate mockgen -source store.go -destination mock/store_mock.go -package mock

package reportstore

type Store interface {
	CreateReport(r CreateReportRequest) error
	GetUserReport(r GetUserReportRequest) (GetUserReportResponse, error)
	UserHasAccess(r UserHasAccessRequest) (UserHasAccessResponse, error)
	AppendMessage(r AppendMessageRequest) error
}

type CreateReportRequest struct {
	ReportID string
	UserID   string
	Status   string
	Title    string
}

type GetUserReportRequest struct {
	UserID   string
	ReportID string
}

type GetUserReportResponse struct {
	Report Report
}

type UserHasAccessRequest struct {
	UserID   string
	ReportID string
}

type UserHasAccessResponse struct {
	HasAcess bool
}

type AppendMessageRequest struct {
	ReportID  string
	MessageID string
	Body      string
}
