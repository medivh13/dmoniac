package dmoniac

/*
 * Author      : Jody (github.com/medivh13)
 * Modifier    :
 * Domain      : dmoniac
 */
import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type DmoniacDTOInterface interface {
	Validate() error
}

type DmoniacReqDTO struct {
	Ppb    int64 `json:"ppb"`
	UserId int64  `json:"user_id"`
}

func (dto *DmoniacReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Ppb, validation.Required),
		validation.Field(&dto.UserId, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetRiwayatReqDTO struct {
	UserId int64 `json:"user_id"`
}

type GetRiwayatRespDTO struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	UserName  string     `json:"user_name"`
	Stadium   int64     `json:"stadium"`
	CreatedAt time.Time `json:"created_at"`
}
