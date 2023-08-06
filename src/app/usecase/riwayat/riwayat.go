package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : article
 */

import (
	dto "dmoniac/src/app/dto/dmoniac"
	"log"
	
	riwayatRepo "dmoniac/src/infra/persistence/postgres/riwayat"
)

type RiwayatUCInterface interface {
	Create(data *dto.DmoniacReqDTO) error
	GetList(req *dto.GetRiwayatReqDTO) ([]*dto.GetRiwayatRespDTO, error)
}

type riwayatUseCase struct {
	RiwayatRepo riwayatRepo.RiwayatRepository
}

func NewRiwayatUseCase(articleRepo riwayatRepo.RiwayatRepository) RiwayatUCInterface {
	return &riwayatUseCase{
		RiwayatRepo: articleRepo,
	}
}

func (uc *riwayatUseCase) Create(data *dto.DmoniacReqDTO) error {

	stadium := scoringStadium(data.Ppb)

	err := uc.RiwayatRepo.Create(data, stadium)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *riwayatUseCase) GetList(req *dto.GetRiwayatReqDTO) ([]*dto.GetRiwayatRespDTO, error) {
	
	data, err := uc.RiwayatRepo.GetList(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToRiwayat(data), nil
}

func scoringStadium (ppb int64) int64 {
	var stadium int64
	if ppb <= 636{
		stadium = 0
	}
	if ppb >= 636 && ppb <= 1020{
		stadium = 1
	}
	if ppb >= 1020 && ppb <= 1943{
		stadium = 2
	}
	if ppb >= 1943 && ppb <= 4421{
		stadium = 3
	}
	if ppb >= 4421 && ppb <= 4421{
		stadium = 4
	}
	if ppb >= 12781{
		stadium = 5
	}

	return stadium
}
