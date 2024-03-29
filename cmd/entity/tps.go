package entity

type TPS struct {
	Chart        map[string]int `json:"chart"`
	Images       []string       `json:"images"`
	Administrasi struct {
		SuaraSah      int `json:"suara_sah"`
		SuaraTidakSah int `json:"suara_tidak_sah"`
		SuaraTotal    int `json:"suara_total"`
	} `json:"administrasi"`
	Ts          string `json:"ts"`
	StatusSuara bool   `json:"status_suara"`
	StatusAdm   bool   `json:"status_adm"`
}
