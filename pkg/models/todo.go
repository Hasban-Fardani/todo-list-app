package models

import "time"

type Todo struct {
	Id           int
	IdUser       int
	IdGroup      int
	NamaKegiatan string
	Deskripsi    string
	StartAt      time.Time
	EndAt        time.Time
}

type TodoGroup struct {
	Id             int
	IdUser         int
	Nama           string
	SkalaPrioritas int
	Warna          string
}
