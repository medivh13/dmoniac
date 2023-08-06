package dmoniac

import models "dmoniac/src/infra/models/dmoniac"

func ToRiwayat(datas []*models.GetRiwayatModel) []*GetRiwayatRespDTO {
	var resp []*GetRiwayatRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnRiwayat(m))
	}
	return resp
}

func ToReturnRiwayat(d *models.GetRiwayatModel) *GetRiwayatRespDTO {
	return &GetRiwayatRespDTO{
		ID:        d.ID,
		UserID:    d.UserID,
		UserName: d.UserName,
		Stadium:     d.Stadium,
		CreatedAt: d.CreatedAt,
	}
}
