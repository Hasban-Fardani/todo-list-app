package models

import "time"

type Todo struct {
	Id        int       `json:"id" db:"id"`
	IdUser    int       `json:"idUser" db:"idUser"`
	IdGroup   int       `json:"idGroup" db:"idGroup"`
	Nama      string    `json:"nama" db:"nama"`
	Deskripsi string    `json:"deskripsi" db:"deskripsi"`
	StartAt   time.Time `json:"startAt" db:"startAt"`
	EndAt     time.Time `json:"endAt" db:"endAt"`
}

type TodoGroup struct {
	Id             int    `json:"id" db:"id"`
	IdUser         int    `json:"idUser" db:"idUser"`
	Nama           string `json:"nama" db:"nama"`
	SkalaPrioritas int    `json:"skalaPrioritas" db:"skalaPrioritas"`
	Warna          string `json:"warna" db:"warna"`
}
