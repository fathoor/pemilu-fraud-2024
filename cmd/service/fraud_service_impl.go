package service

import (
	"encoding/json"
	"fmt"
	"github.com/fathoor/pemilu-fraud-2024/cmd/entity"
	"github.com/fathoor/pemilu-fraud-2024/cmd/validation"
	"github.com/fathoor/pemilu-fraud-2024/internal/exception"
	"github.com/fathoor/pemilu-fraud-2024/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

type fraudServiceImpl struct {
}

func (f *fraudServiceImpl) FraudCheck(k string) []entity.Fraud {
	var (
		kota     = validation.ValidateKota(k)
		provinsi = validation.ValidateProvinsi(kota.Kode[0:2])
		response = make([]entity.Fraud, 0)
		counter  = 0
	)

	agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s.json", util.BaseURL, kota.Kode[0:2], kota.Kode))

	var kecamatan []entity.Wilayah
	agent.Struct(&kecamatan)

	for _, kecamatan := range kecamatan {
		agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s/%s.json", util.BaseURL, kecamatan.Kode[0:2], kecamatan.Kode[0:4], kecamatan.Kode))

		var kelurahan []entity.Wilayah
		agent.Struct(&kelurahan)

		for _, kelurahan := range kelurahan {
			agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s/%s/%s.json", util.BaseURL, kelurahan.Kode[0:2], kelurahan.Kode[0:4], kelurahan.Kode[0:6], kelurahan.Kode))

			var tps []entity.Wilayah
			agent.Struct(&tps)

			total := len(tps)

			for current, tps := range tps {
				log.Info(fmt.Sprintf("[%d/%d] CHECKING %s KECAMATAN %s KELURAHAN %s", current+1, total, tps.Nama, kecamatan.Nama, kelurahan.Nama))
				agent := fiber.Get(fmt.Sprintf("%spemilu/hhcw/ppwp/%s/%s/%s/%s/%s.json", util.BaseURL, tps.Kode[0:2], tps.Kode[0:4], tps.Kode[0:6], tps.Kode[0:10], tps.Kode))

				var fraud entity.TPS
				agent.Struct(&fraud)

				hasVote := fraud.Administrasi.SuaraSah != 0 && fraud.Administrasi.SuaraTidakSah != 0 && fraud.Administrasi.SuaraTotal != 0
				hasCandidateVote := fraud.Chart[util.Kandidat["1"]] != 0 && fraud.Chart[util.Kandidat["2"]] != 0 && fraud.Chart[util.Kandidat["3"]] != 0

				if !fraud.StatusAdm && !fraud.StatusSuara {
					continue
				}

				if fraud.Chart == nil || !hasVote || !hasCandidateVote {
					continue
				}

				totalSuara := fraud.Chart[util.Kandidat["1"]] + fraud.Chart[util.Kandidat["2"]] + fraud.Chart[util.Kandidat["3"]]
				if totalSuara != fraud.Administrasi.SuaraSah {
					counter++
					log.Info(fmt.Sprintf("[%d] ANOMALY FOUND AT KECAMATAN %s KELURAHAN %s %s", counter, kecamatan.Nama, kelurahan.Nama, tps.Nama))

					response = append(response, entity.Fraud{
						Kode:      tps.Kode,
						Timestamp: fraud.Ts,
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
						Provinsi:  provinsi.Nama,
						Kota:      kota.Nama,
						Kecamatan: kecamatan.Nama,
						Kelurahan: kelurahan.Nama,
						TPS:       tps.Nama,
						Images:    fraud.Images,
						URL:       fmt.Sprintf("%spilpres/hitung-suara/%s/%s/%s/%s/%s", util.PemiluURL, tps.Kode[0:2], tps.Kode[0:4], tps.Kode[0:6], tps.Kode[0:10], tps.Kode),
					})

					logs, err := os.OpenFile(fmt.Sprintf("/app/logs/%s.json", kota.Kode), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
					exception.PanicIfNeeded(err)

					result, err := json.MarshalIndent(response, "", "  ")
					exception.PanicIfNeeded(err)

					_, err = logs.Write(result)
					exception.PanicIfNeeded(err)
					logs.Close()
				}
			}
		}
	}

	log.Info(fmt.Sprintf("FOUND A TOTAL OF [%d] ANOMALIES AT %s", counter, kota.Nama))

	return response
}

func (f *fraudServiceImpl) FraudCache(k string) []entity.Fraud {
	kota := validation.ValidateKota(k)

	file, err := os.OpenFile(fmt.Sprintf("/app/logs/%s.json", kota.Kode), os.O_RDONLY, 0644)
	if err != nil {
		f.FraudCheck(k)

		panic(exception.NotFoundError{
			Message: "Data belum tersedia",
		})
	}

	defer file.Close()
	var fraud []entity.Fraud
	err = json.NewDecoder(file).Decode(&fraud)
	exception.PanicIfNeeded(err)

	return fraud
}

func ProvideFraudService() FraudService {
	return &fraudServiceImpl{}
}
