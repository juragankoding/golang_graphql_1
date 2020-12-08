package models

type Barang struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Description string `json:"description"`
}

type JenisBarang struct {
	ID          int    `json:"id"`
	JenisBarang string `json:"jenis_barnag"`
}
