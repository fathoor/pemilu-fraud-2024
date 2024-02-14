package service

import (
	"encoding/json"
	"fmt"
	"github.com/fathoor/fraud24/cmd/fraud24/entity"
	"github.com/fathoor/fraud24/cmd/fraud24/util"
	"io"
	"net/http"
	"os"
)

type fraudServiceImpl struct {
}

func (service *fraudServiceImpl) FraudCheck(k string) string {
	kode := k[0:2] + "/" + k[0:4] + "/" + k[0:6] + "/" + k[0:10] + "/" + k
	tpsURL := fmt.Sprintf("%spemilu/hhcw/ppwp/%s.json", util.BaseURL, kode)

	response, err := http.Get(tpsURL)
	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Println("Error: TPS tidak ditemukan")
		os.Exit(1)
	}

	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var tps entity.TPS
	err = json.Unmarshal(body, &tps)
	if err != nil {
		fmt.Println("Error: Tidak dapat mem-parse response")
		os.Exit(1)
	}

	totalSuara := tps.Chart[util.Kandidat["1"]] + tps.Chart[util.Kandidat["2"]] + tps.Chart[util.Kandidat["3"]]
	if totalSuara > tps.Administrasi.SuaraTotal {
		return fmt.Sprintf("Terdeteksi kecurangan! Total suara tidak sesuai dengan jumlah pemilih (%d/%d) pada TPS", totalSuara, tps.Administrasi.SuaraTotal)
	}

	return ""
}

func ProvideFraudService() FraudService {
	return &fraudServiceImpl{}
}
