package models

import (
	"net/http"
)

type KedatanganKapal struct {
	Id                 int    `json:"id"`
	Id_Kapal           int    `json:"id_kapal`
	Id_Pelabuhan       int    `json:"id_pelabuhan`
	Id_Jenis_Muatan    int    `json:"id_jenis_muatan"`
	Daerah_Penangkapan string `json:"daerah_penangkapan`
	Tgl_Keberangkatan  string `json:"tgl_keberangkatan`
	Tgl_Kedatangan     string `json:"tgl_kedatangan"`
	Created_At         string `json:"created_at`
	Updated_At         string `json:"updated_at`
}

func GetAllKedatanganKapal() (Response, error) {
	var obj KedatanganKapal
	var arrobj []KedatanganKapal
	var res Response

	con := db.DBConnect()

	sqlStatement := "SELECT *FROM kedatangan_kapal.kedatangan_kapal"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Id_Kapal, &obj.Id_Pelabuhan, &obj.Id_Jenis_Muatan, &obj.Daerah_Penangkapan, obj.Tgl_Keberangkatan, obj.Tgl_Kedatangan, obj.Created_At, obj.Updated_At)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Massage = "succes"
	res.Data = arrobj

	return res, nil
}
