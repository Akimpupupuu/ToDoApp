package users_transport_http

import (
	"net/http"

	"github.com/Akimpupupuu/ToDoApp/internal/core/domain"
	core_logger "github.com/Akimpupupuu/ToDoApp/internal/core/logger"
	core_http_request "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/request"
	core_http_response "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/response"
	core_http_types "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/types"
	core_http_utils "github.com/Akimpupupuu/ToDoApp/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetIntPathValues(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId path value")
		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userPatch := userPatchFromRequst(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequst(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
