package reporting

import "errors"

var (
	ErrNotHavePermission   = errors.New("Request is not permitted.")
	ErrUnknownReportStatus = errors.New("Report status is unknown.")
)
