package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pursuit/gateway/internal/proto/out/api/portal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	UserClient portal_proto.UserClient
}

func convertGrpcError(err error) (string, int) {
	st := status.Convert(err)
	httpCode := 503
	switch st.Code() {
	case codes.Unauthenticated:
		httpCode = 401
	case codes.PermissionDenied:
		httpCode = 403
	case codes.InvalidArgument, codes.AlreadyExists:
		httpCode = 422
	case codes.NotFound:
		httpCode = 404
	}

	msg := st.Err().Error()
	if httpCode == 503 {
		msg = "Please try again in a few moment"
	}
	return msg, httpCode
}

type tempCreateUser struct {
	Username string          `json:"username"`
	Password json.RawMessage `json:"password"`
}

func (this Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(422)
		return
	}

	var jsonPayload tempCreateUser
	if err := json.Unmarshal(body, &jsonPayload); err != nil {
		w.WriteHeader(422)
		return
	}

	if len(jsonPayload.Password) < 3 {
		w.WriteHeader(422)
		return
	}

	payload := portal_proto.CreateUserPayload{
		Username: jsonPayload.Username,
		Password: jsonPayload.Password[1 : len(jsonPayload.Password)-1],
	}

	_, err = this.UserClient.Create(r.Context(), &payload)
	if err != nil {
		sErr, httpStatus := convertGrpcError(err)
		w.WriteHeader(httpStatus)
		w.Write([]byte(sErr))
		return
	}

	w.WriteHeader(201)
}
