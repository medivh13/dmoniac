package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "dmoniac/src/app/dto/dmoniac"
	usecases "dmoniac/src/app/usecase/riwayat"
	common_error "dmoniac/src/infra/errors"
	"dmoniac/src/interface/rest/response"
)

type RiwayatHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
}

type riwayatHandler struct {
	response response.IResponseClient
	usecase  usecases.RiwayatUCInterface
}

func NewRiwayatHandler(r response.IResponseClient, h usecases.RiwayatUCInterface) RiwayatHandlerInterface {
	return &riwayatHandler{
		response: r,
		usecase:  h,
	}
}

func (h *riwayatHandler) Create(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.DmoniacReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Create(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Adding New Riwayat",
		nil,
		nil,
	)
}

func (h *riwayatHandler) GetList(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetRiwayatReqDTO{}

	if r.URL.Query().Get("user_id") != "" {
		var err error
		getDTO.UserId, err = strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)

		if err != nil {
			h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		}
	}



	data, err := h.usecase.GetList(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Article",
		data,
		nil,
	)
}
