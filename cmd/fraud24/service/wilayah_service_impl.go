package service

import (
	"encoding/json"
	"fmt"
	"github.com/fathoor/fraud24/cmd/fraud24/entity"
	"github.com/fathoor/fraud24/internal/exception"
	"github.com/fathoor/fraud24/internal/util"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

type wilayahServiceImpl struct {
}

func (w *wilayahServiceImpl) InitWilayah() {
	//GetProvinsi()
	//GetKota()
	GetKecamatan()
	GetKelurahan()
	GetTPS()
}

func GetProvinsi() {
	agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/0.json", util.BaseURL))

	var provinsi []entity.Wilayah
	agent.Struct(&provinsi)

	result := make([]entity.Wilayah, len(provinsi))
	for i, p := range provinsi {
		p.Nama = strings.ToUpper(p.Nama)
		result[i] = p
	}

	data, err := json.Marshal(result)
	exception.PanicIfNeeded(err)

	file, err := os.OpenFile("../../assets/PROVINSI.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	_, err = file.Write(data)
	exception.PanicIfNeeded(err)
}

func GetKota() {
	file, err := os.OpenFile("../../assets/PROVINSI.json", os.O_RDWR, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var provinsi []entity.Wilayah
	err = json.NewDecoder(file).Decode(&provinsi)
	exception.PanicIfNeeded(err)

	result := make([]entity.Wilayah, 0)
	for _, p := range provinsi {
		agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s.json", util.BaseURL, p.Kode))

		var kota []entity.Wilayah
		agent.Struct(&kota)

		for _, k := range kota {
			k.Nama = strings.ToUpper(k.Nama)
			result = append(result, k)
		}
	}

	data, err := json.Marshal(result)
	exception.PanicIfNeeded(err)

	f, err := os.OpenFile("../../assets/KOTA.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer f.Close()
	_, err = f.Write(data)
	exception.PanicIfNeeded(err)
}

func GetKecamatan() {
	//file, err := os.OpenFile("../../assets/KOTA.json", os.O_RDWR, 0644)
	//exception.PanicIfNeeded(err)
	//
	//defer file.Close()
	//var kota []entity.Wilayah
	//err = json.NewDecoder(file).Decode(&kota)
	//exception.PanicIfNeeded(err)

	kota := entity.Wilayah{
		Kode: "3578",
		Nama: "KOTA SURABAYA",
	}

	result := make([]entity.Wilayah, 0)
	//for _, k := range kota {
	//	agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s.json", util.BaseURL, k.Kode[0:2], k.Kode))
	//
	//	var kecamatan []entity.Wilayah
	//	agent.Struct(&kecamatan)
	//
	//	for _, kc := range kecamatan {
	//		kc.Nama = strings.ToUpper(kc.Nama)
	//		result = append(result, kc)
	//	}
	//}

	agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s.json", util.BaseURL, kota.Kode[0:2], kota.Kode))

	var kecamatan []entity.Wilayah
	agent.Struct(&kecamatan)

	for _, kc := range kecamatan {
		kc.Nama = strings.ToUpper(kc.Nama)
		result = append(result, kc)
	}

	data, err := json.Marshal(result)
	exception.PanicIfNeeded(err)

	f, err := os.OpenFile("../../assets/KECAMATAN.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer f.Close()
	_, err = f.Write(data)
	exception.PanicIfNeeded(err)
}

func GetKelurahan() {
	file, err := os.OpenFile("../../assets/KECAMATAN.json", os.O_RDWR, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var kecamatan []entity.Wilayah
	err = json.NewDecoder(file).Decode(&kecamatan)
	exception.PanicIfNeeded(err)

	result := make([]entity.Wilayah, 0)
	for _, kc := range kecamatan {
		agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s/%s.json", util.BaseURL, kc.Kode[0:2], kc.Kode[0:4], kc.Kode))

		var kelurahan []entity.Wilayah
		agent.Struct(&kelurahan)

		for _, kl := range kelurahan {
			kl.Nama = strings.ToUpper(kl.Nama)
			result = append(result, kl)
		}
	}

	data, err := json.Marshal(result)
	exception.PanicIfNeeded(err)

	f, err := os.OpenFile("../../assets/KELURAHAN.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer f.Close()
	_, err = f.Write(data)
	exception.PanicIfNeeded(err)
}

func GetTPS() {
	file, err := os.OpenFile("../../assets/KELURAHAN.json", os.O_RDWR, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var kelurahan []entity.Wilayah
	err = json.NewDecoder(file).Decode(&kelurahan)
	exception.PanicIfNeeded(err)

	result := make([]entity.Wilayah, 0)
	for _, kl := range kelurahan {
		agent := fiber.Get(fmt.Sprintf("%swilayah/pemilu/ppwp/%s/%s/%s/%s.json", util.BaseURL, kl.Kode[0:2], kl.Kode[0:4], kl.Kode[0:6], kl.Kode))

		var tps []entity.Wilayah
		agent.Struct(&tps)

		for _, t := range tps {
			t.Nama = strings.ToUpper(t.Nama)
			result = append(result, t)
		}
	}

	data, err := json.Marshal(result)
	exception.PanicIfNeeded(err)

	f, err := os.OpenFile("../../assets/TPS.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer f.Close()
	_, err = f.Write(data)
	exception.PanicIfNeeded(err)
}

func ProvideWilayahService() WilayahService {
	return &wilayahServiceImpl{}
}
