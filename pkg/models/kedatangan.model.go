package models

import (
	"fmt"
	"log"
	"mini_project/pkg/database"
	"time"
)

type KedatanganKapal struct {
	Id                 int       `gorm:"primaryKey" json:"id"`
	Id_Kapal           int       `gorm:"foreignKey" json:"id_kapal`
	Id_Pelabuhan       int       `gorm:"foreignKey" json:"id_pelabuhan`
	Id_Jenis_Muatan    int       `gorm:"foreignKey" json:"id_jenis_muatan"`
	Daerah_Penangkapan string    `json:"daerah_penangkapan`
	Tgl_Keberangkatan  time.Time `json:"tgl_keberangkatan`
	Tgl_Kedatangan     time.Time `json:"tgl_kedatangan"`
	Created_At         time.Time `json:"created_at`
	Updated_At         time.Time `json:"updated_at`
}

func CreateKedatanganKapal() (Response, error) {
	db := database.DB

	sqlStatement := `INSERT * FROM kedatangan_kapal(daerah_penangkapan, tgl_keberangkatan, tgl_kedatangan) VALUES ($1, $2, $3) RETURNING id`

	var res Response
	var id int
	err := db.Query(sqlStatement, kedatangan_kapal.Daerah_Penangkapan, kedatangan_kapal.Tgl_Keberangkatan, kedatangan_kapal.Tgl_Kedatangan).Scan(&id)

	if err != nil {
		log.Fatalf("Cannot Execution The Query. %v", err)
	}
	fmt.Printf("Record Insert Data Single %v", id)

	return res, err
}
