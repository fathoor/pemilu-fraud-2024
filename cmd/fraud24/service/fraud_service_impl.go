package service

import (
	"encoding/json"
	"fmt"
	"github.com/fathoor/fraud24/cmd/fraud24/entity"
	"github.com/fathoor/fraud24/internal/exception"
	"github.com/fathoor/fraud24/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

type fraudServiceImpl struct {
}

func (f *fraudServiceImpl) FraudCheck() []entity.Fraud {
	file, err := os.OpenFile("/app/assets/KECAMATAN.json", os.O_RDWR, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var kecamatan []entity.Wilayah
	err = json.NewDecoder(file).Decode(&kecamatan)
	exception.PanicIfNeeded(err)

	response := make([]entity.Fraud, 0)
	for _, kecamatan := range kecamatan {
		log.Debug(fmt.Sprintf("Currently processing KECAMATAN %s", kecamatan.Nama))
		agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s/%s.json", util.BaseURL, kecamatan.Kode[0:2], kecamatan.Kode[0:4], kecamatan.Kode))

		var kelurahan []entity.Wilayah
		agent.Struct(&kelurahan)

		for _, kelurahan := range kelurahan {
			log.Debug(fmt.Sprintf("Currently processing KELURAHAN %s", kelurahan.Nama))
			agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s/%s/%s.json", util.BaseURL, kelurahan.Kode[0:2], kelurahan.Kode[0:4], kelurahan.Kode[0:6], kelurahan.Kode))

			var tps []entity.Wilayah
			agent.Struct(&tps)

			for _, tps := range tps {
				log.Debug(fmt.Sprintf("Currently processing %s KECAMATAN %s KELURAHAN %s", tps.Nama, kecamatan.Nama, kelurahan.Nama))
				agent := fiber.Get(fmt.Sprintf("%spemilu/hhcw/ppwp/%s/%s/%s/%s/%s.json", util.BaseURL, tps.Kode[0:2], tps.Kode[0:4], tps.Kode[0:6], tps.Kode[0:10], tps.Kode))

				var fraud entity.TPS
				agent.Struct(&fraud)

				if !fraud.StatusAdm {
					continue
				}

				totalSuara := fraud.Chart[util.Kandidat["1"]] + fraud.Chart[util.Kandidat["2"]] + fraud.Chart[util.Kandidat["3"]]
				if totalSuara > fraud.Administrasi.SuaraSah {
					log.Debug(fmt.Sprintf("Fraud detected at Kecamatan %s, Kelurahan %s, TPS %s", kecamatan.Nama, kelurahan.Nama, tps.Nama))
					response = append(response, entity.Fraud{
						KodeTPS: tps.Kode,
						Suara: struct {
							SuaraSah      int `json:"suara_sah"`
							SuaraTidakSah int `json:"suara_tidak_sah"`
							SuaraTotal    int `json:"suara_total"`
						}{
							SuaraSah:      fraud.Administrasi.SuaraSah,
							SuaraTidakSah: fraud.Administrasi.SuaraTidakSah,
							SuaraTotal:    fraud.Administrasi.SuaraTotal,
						},
						Hasil: struct {
							Anies   int `json:"01"`
							Prabowo int `json:"02"`
							Ganjar  int `json:"03"`
						}{
							Anies:   fraud.Chart[util.Kandidat["1"]],
							Prabowo: fraud.Chart[util.Kandidat["2"]],
							Ganjar:  fraud.Chart[util.Kandidat["3"]],
						},
						Provinsi:  "JAWA TIMUR",
						Kota:      "KOTA SURABAYA",
						Kecamatan: kecamatan.Nama,
						Kelurahan: kelurahan.Nama,
						TPS:       tps.Nama,
						Image:     fraud.Images,
						URL:       fmt.Sprintf("%spilpres/hitung-suara/%s/%s/%s/%s/%s", util.PemiluURL, tps.Kode[0:2], tps.Kode[0:4], tps.Kode[0:6], tps.Kode[0:10], tps.Kode),
						Timestamp: fraud.Ts,
					})
				}
			}
		}
	}

	data, err := json.Marshal(response)
	exception.PanicIfNeeded(err)

	logs, err := os.OpenFile("/app/logs/KOTA_SURABAYA.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer logs.Close()
	_, err = logs.Write(data)
	exception.PanicIfNeeded(err)

	return response
}

func (f *fraudServiceImpl) FraudCheckCache() []entity.Fraud {
	file, err := os.OpenFile("/app/logs/KOTA_SURABAYA.json", os.O_RDONLY, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var fraud []entity.Fraud
	err = json.NewDecoder(file).Decode(&fraud)
	exception.PanicIfNeeded(err)

	return fraud
}

func ProvideFraudService() FraudService {
	return &fraudServiceImpl{}
}
