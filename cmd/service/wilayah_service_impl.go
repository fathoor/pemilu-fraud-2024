package service

import (
	"encoding/json"
	"fmt"
	"github.com/fathoor/pemilu-fraud-2024/cmd/entity"
	"github.com/fathoor/pemilu-fraud-2024/internal/exception"
	"github.com/fathoor/pemilu-fraud-2024/internal/util"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

type wilayahServiceImpl struct {
}

func (w *wilayahServiceImpl) GetWilayah() {
	GetProvinsi()
	GetKota()
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

	data, err := json.MarshalIndent(result, "", "  ")
	exception.PanicIfNeeded(err)

	file, err := os.OpenFile("/app/assets/PROVINSI.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	_, err = file.Write(data)
	exception.PanicIfNeeded(err)
}

func GetKota() {
	file, err := os.OpenFile("/app/assets/PROVINSI.json", os.O_RDWR, 0644)
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

	data, err := json.MarshalIndent(result, "", "  ")
	exception.PanicIfNeeded(err)

	f, err := os.OpenFile("/app/assets/KOTA.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	exception.PanicIfNeeded(err)

	defer f.Close()
	_, err = f.Write(data)
	exception.PanicIfNeeded(err)
}

func ProvideWilayahService() WilayahService {
	return &wilayahServiceImpl{}
}
