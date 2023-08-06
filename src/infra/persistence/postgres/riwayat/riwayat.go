package dmoniac

import (
	dto "dmoniac/src/app/dto/dmoniac"
	models "dmoniac/src/infra/models/dmoniac"
	"log"

	"github.com/jmoiron/sqlx"
)

type RiwayatRepository interface {
	Create(data *dto.DmoniacReqDTO, stadium int64) error
	GetList(data *dto.GetRiwayatReqDTO) ([]*models.GetRiwayatModel, error)
}

const (
	CreateRiwayat = `insert into riwayat (user_id, stadium, ppb) values ($1, $2, $3)`
	GetRiwayat    = `select r.id, r.user_id, u.name, r.created_at, r.stadium
	from users u
	JOIN riwayat r ON
	u.id = r.user_id
	where u.id = $1`
)

var statement PreparedStatement

type PreparedStatement struct {
	createRiwayat *sqlx.Stmt
	getRiwayat    *sqlx.Stmt
}

type riwayatRepo struct {
	Connection *sqlx.DB
}

func NewRiwayatRepository(db *sqlx.DB) RiwayatRepository {
	repo := &riwayatRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *riwayatRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *riwayatRepo) {
	statement = PreparedStatement{
		createRiwayat: m.Preparex(CreateRiwayat),
		getRiwayat:    m.Preparex(GetRiwayat),
	}
}

func (p *riwayatRepo) Create(data *dto.DmoniacReqDTO, stadium int64) error {

	_, err := statement.createRiwayat.Exec(data.UserId, stadium, data.Ppb)

	if err != nil {
		log.Println("Failed Query Create riwayat : ", err.Error())
		return err
	}

	return nil
}

func (p *riwayatRepo) GetList(data *dto.GetRiwayatReqDTO) ([]*models.GetRiwayatModel, error) {
	var resultData []*models.GetRiwayatModel
	var err error

	err = statement.getRiwayat.Select(&resultData, data.UserId)

	if err != nil {
		return nil, err
	}

	return resultData, nil
}
