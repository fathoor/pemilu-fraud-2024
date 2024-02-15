package entity

type Fraud struct {
	KodeTPS string `json:"kode_tps"`
	Suara   struct {
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
	Image     []string `json:"image"`
	URL       string   `json:"url"`
	Timestamp string   `json:"timestamp"`
}
