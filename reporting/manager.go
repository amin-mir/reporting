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
	var resp reportstore.UserHasAccessResponse
	resp, err = m.store.UserHasAccess(reportstore.UserHasAccessRequest{
		UserID:   request.UserID,
		ReportID: request.ReportID,
	})
	if err != nil {
		return
	}
	if !resp.HasAccess {
		err = ErrNotHavePermission
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

func (m *ReportManager) UpdateReportStatus(request UpdateReportStatusRequest) (response UpdateReportStatusResponse, err error) {
	reportStatus := reportstore.ParseReportStatus(request.Status)
	if reportStatus == reportstore.ReportStatusUnknown {
		err = ErrUnknownReportStatus
		return
	}

	var resp reportstore.UserHasAccessResponse
	resp, err = m.store.UserHasAccess(reportstore.UserHasAccessRequest{
		UserID:   request.UserID,
		ReportID: request.ReportID,
	})
	if err != nil {
		return
	}
	if !resp.HasAccess {
		err = ErrNotHavePermission
		return
	}

	err = m.store.UpdateReportStatus(reportstore.UpdateReportStatusRequest{
		ReportID: request.ReportID,
		Status:   reportStatus,
	})
	return
}

func (m *ReportManager) GetUserReports(request GetUserReportsRequest) (response GetUserReportsResponse, err error) {
	var resp reportstore.GetUserReportsResponse
	resp, err = m.store.GetUserReports(reportstore.GetUserReportsRequest{
		UserID: request.UserID,
	})
	response.Reports = resp.Reports
	return
}
