package entity

type Fraud struct {
	Kode      string `json:"kode"`
	Timestamp string `json:"timestamp"`
	Suara     struct {
		SuaraSah      int `json:"suara_sah"`
		SuaraTidakSah int `json:"suara_tidak_sah"`
		SuaraTotal    int `json:"suara_total"`
	} `json:"suara"`
	Hasil struct {
		Anies   int `json:"01"`
		Prabowo int `json:"02"`
		Ganjar  int `json:"03"`
	} `json:"hasil"`
	Provinsi  string   `json:"provinsi"`
	Kota      string   `json:"kota"`
	Kecamatan string   `json:"kecamatan"`
	Kelurahan string   `json:"kelurahan"`
	TPS       string   `json:"tps"`
	Images    []string `json:"images"`
	URL       string   `json:"url"`
}
