// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ResultInsert interface {
	IsResultInsert()
}

type Karyawan struct {
	ID         int    `json:"id"`
	IDKaryawan string `json:"idKaryawan"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	NoHp       string `json:"noHp"`
}

type ResultJenisBarang struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func (ResultJenisBarang) IsResultInsert() {}
