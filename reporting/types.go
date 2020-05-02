package reporting

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
