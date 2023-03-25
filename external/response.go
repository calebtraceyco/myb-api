package external

import (
	"github.com/calebtraceyco/models/pkg/response"
	"github.com/jackc/pgx/v5/pgconn"
)

type Response struct {
	Details []any            `json:"Details,omitempty"`
	Message response.Message `json:"Message,omitempty"`
}

func (res *Response) SetErrorLog(errs []error, trace, statusCode string) {
	for _, err := range errs {
		res.Message.ErrorLog = append(res.Message.ErrorLog, response.ErrorLog{
			StatusCode: statusCode,
			Trace:      trace,
			RootCause:  err.Error(),
		})
	}
}

type ExecResponse struct {
	Status pgconn.CommandTag `json:"Status"`
}
