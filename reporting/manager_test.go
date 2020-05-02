package reporting

import (
	"errors"
	"testing"

	"github.com/amin-mir/reporting/reportstore"
	mockreportstore "github.com/amin-mir/reporting/reportstore/mock"
	mockuuid "github.com/amin-mir/reporting/uuid/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ReportManagerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl              *gomock.Controller
	mockReportStore   *mockreportstore.MockStore
	mockUUIDGenerator *mockuuid.MockGenerator

	manager *ReportManager
}

func TestReportManagerSuite(t *testing.T) {
	suite.Run(t, new(ReportManagerSuite))
}

func (s *ReportManagerSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.mockReportStore = mockreportstore.NewMockStore(s.ctrl)
	s.mockUUIDGenerator = mockuuid.NewMockGenerator(s.ctrl)

	s.manager = NewReportManager(s.mockUUIDGenerator, s.mockReportStore)
}

func (s *ReportManagerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *ReportManagerSuite) TestCreateReport() {
	reportID := "reportid"
	userID := "userid"
	title := "title"

	s.mockUUIDGenerator.EXPECT().Generate().Return(reportID).Times(1)

	s.mockReportStore.EXPECT().CreateReport(gomock.Eq(reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   userID,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    title,
	})).Return(nil).Times(1)

	actualResponse, err := s.manager.CreateReport(CreateReportRequest{
		UserID: userID,
		Title:  title,
	})
	s.NoError(err)

	expectedResponse := CreateReportResponse{
		ReportID: reportID,
	}
	s.Equal(expectedResponse, actualResponse)
}

func (s *ReportManagerSuite) TestCreateReportError() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportid").Times(1)

	createError := errors.New("create report error")
	s.mockReportStore.EXPECT().CreateReport(gomock.Any()).Return(createError).Times(1)

	_, err := s.manager.CreateReport(CreateReportRequest{})
	s.Equal(createError, err)
}

func (s *ReportManagerSuite) TestAppendMessage_UserHasAccessError() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportid").Times(1)

	hasAccessError := errors.New("has access error")
	s.mockReportStore.
		EXPECT().
		UserHasAccess(gomock.Any()).
		Return(reportstore.UserHasAccessResponse{}, hasAccessError).
		Times(1)

	_, err := s.manager.AppendMessage(AppendMessageRequest{})
	s.Equal(hasAccessError, err)
}

func (s *ReportManagerSuite) TestAppendMessage_UserNoAccess() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportid").Times(1)

	s.mockReportStore.
		EXPECT().
		UserHasAccess(gomock.Any()).
		Return(reportstore.UserHasAccessResponse{HasAccess: false}, nil).
		Times(1)

	_, err := s.manager.AppendMessage(AppendMessageRequest{})
	s.Equal(ErrNotHavePermission, err)
}

func (s *ReportManagerSuite) TestAppendMessage_AppendMessageError() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportid").Times(1)

	s.mockReportStore.EXPECT().UserHasAccess(gomock.Any()).Return(reportstore.UserHasAccessResponse{HasAccess: true}, nil).Times(1)

	appendError := errors.New("append message error")
	s.mockReportStore.EXPECT().AppendMessage(gomock.Any()).Return(appendError).Times(1)

	_, err := s.manager.AppendMessage(AppendMessageRequest{})
	s.Equal(appendError, err)
}

func (s *ReportManagerSuite) TestAppendMessage() {
	userID := "userid"
	reportID := "reportid"
	messageID := "messageid"
	body := "body"

	s.mockUUIDGenerator.EXPECT().Generate().Return(messageID).Times(1)

	s.mockReportStore.
		EXPECT().
		UserHasAccess(gomock.Eq(reportstore.UserHasAccessRequest{UserID: userID, ReportID: reportID})).
		Return(reportstore.UserHasAccessResponse{HasAccess: true}, nil).
		Times(1)

	s.mockReportStore.
		EXPECT().
		AppendMessage(gomock.Eq(reportstore.AppendMessageRequest{ReportID: reportID, MessageID: messageID, Body: body})).
		Return(nil).
		Times(1)

	actualResponse, err := s.manager.AppendMessage(AppendMessageRequest{
		UserID:   userID,
		ReportID: reportID,
		Body:     body,
	})
	s.NoError(err)

	expectedResponse := AppendMessageResponse{
		MessageID: "messageid",
	}
	s.Equal(expectedResponse, actualResponse)
}

func (s *ReportManagerSuite) TestUpdateReportStatus_UnknownReportStatus() {
	_, err := s.manager.UpdateReportStatus(UpdateReportStatusRequest{
		Status: "some unknonw status",
	})
	s.Equal(ErrUnknownReportStatus, err)
}

func (s *ReportManagerSuite) TestUpdateReportStatus_UserHasAccessError() {
	hasAccessError := errors.New("has access error")
	s.mockReportStore.
		EXPECT().
		UserHasAccess(gomock.Any()).
		Return(reportstore.UserHasAccessResponse{}, hasAccessError).
		Times(1)

	_, err := s.manager.UpdateReportStatus(UpdateReportStatusRequest{Status: reportstore.ReportStatusPending.String()})
	s.Equal(hasAccessError, err)
}

func (s *ReportManagerSuite) TestUpdateReportStatus_UserNoAccess() {
	s.mockReportStore.
		EXPECT().
		UserHasAccess(gomock.Any()).
		Return(reportstore.UserHasAccessResponse{HasAccess: false}, nil).
		Times(1)

	_, err := s.manager.UpdateReportStatus(UpdateReportStatusRequest{Status: reportstore.ReportStatusPending.String()})
	s.Equal(ErrNotHavePermission, err)
}

func (s *ReportManagerSuite) TestUpdateReportStatus() {
	userID := "userid"
	reportID := "reportid"
	status := reportstore.ReportStatusPending

	s.mockReportStore.
		EXPECT().
		UserHasAccess(gomock.Eq(reportstore.UserHasAccessRequest{UserID: userID, ReportID: reportID})).
		Return(reportstore.UserHasAccessResponse{HasAccess: true}, nil).
		Times(1)

	s.mockReportStore.
		EXPECT().
		UpdateReportStatus(gomock.Eq(reportstore.UpdateReportStatusRequest{ReportID: reportID, Status: status})).
		Return(nil).
		Times(1)

	actualResponse, err := s.manager.UpdateReportStatus(UpdateReportStatusRequest{UserID: userID, ReportID: reportID, Status: status.String()})
	s.NoError(err)

	expectedResponse := UpdateReportStatusResponse{}
	s.Equal(expectedResponse, actualResponse)
}

func (s *ReportManagerSuite) TestGetUserReportsError() {
	getError := errors.New("get error")
	s.mockReportStore.
		EXPECT().
		GetUserReports(gomock.Any()).
		Return(reportstore.GetUserReportsResponse{}, getError).
		Times(1)

	_, err := s.manager.GetUserReports(GetUserReportsRequest{})
	s.Equal(getError, err)
}

func (s *ReportManagerSuite) TestGetUserReports() {
	userID := "userid"
	reports := []reportstore.Report{
		{
			ReportID: "reportid1",
			UserID:   userID,
		},
		{
			ReportID: "reportid2",
			UserID:   userID,
		},
	}

	s.mockReportStore.
		EXPECT().
		GetUserReports(gomock.Eq(reportstore.GetUserReportsRequest{UserID: userID})).
		Return(reportstore.GetUserReportsResponse{Reports: reports}, nil).
		Times(1)

	actualResponse, err := s.manager.GetUserReports(GetUserReportsRequest{
		UserID: userID,
	})
	s.NoError(err)

	expectedResponse := GetUserReportsResponse{
		Reports: reports,
	}
	s.Equal(expectedResponse, actualResponse)
}
