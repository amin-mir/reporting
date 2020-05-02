package reporting

import (
	"github.com/amin-mir/reporting/reportstore"
	"github.com/amin-mir/reporting/uuid"
)

type ReportManager struct {
	uuidGenerator uuid.Generator
	store         reportstore.Store
}

func NewReportManager(gen uuid.Generator, store reportstore.Store) *ReportManager {
	return &ReportManager{
		uuidGenerator: gen,
		store:         store,
	}
}

func (m *ReportManager) CreateReport(request CreateReportRequest) (response CreateReportResponse, err error) {
	reportID := m.uuidGenerator.Generate()

	r := reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   request.UserID,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    request.Title,
	}
	err = m.store.CreateReport(r)

	response.ReportID = r.ReportID
	return
}

func (m *ReportManager) AppendMessage(request AppendMessageRequest) (response AppendMessageResponse, err error) {
	messageID := m.uuidGenerator.Generate()

	// check if user has permission
	_, err = m.store.GetUserReport(reportstore.GetUserReportRequest{
		UserID:   request.UserID,
		ReportID: request.ReportID,
	})
	if err != nil {
		return
	}

	// append the message to report
	err = m.store.AppendMessage(reportstore.AppendMessageRequest{
		ReportID:  request.ReportID,
		MessageID: messageID,
		Body:      request.Body,
	})

	response.MessageID = messageID
	return
}
