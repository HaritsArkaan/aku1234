package models

import "time"

type (
	TugasPendahuluan struct {
		ID        uint      `gorm:"primary_key;autoIncrement:true"`
		Judul     string    `gorm:"judul"`
		SubJudul  string    `gorm:"sub_judul"`
		Kategori  string    `gorm:"kategori"`
		CreatedAt time.Time `gorm:"created_at"`
		Deadline  string    `gorm:"deadline"`
		Deskripsi string    `gorm:"deskripsi"`
	}
)
