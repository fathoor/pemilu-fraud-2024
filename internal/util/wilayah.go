package util

import (
	"encoding/json"
	"github.com/fathoor/pemilu-fraud-2024/cmd/entity"
	"github.com/fathoor/pemilu-fraud-2024/internal/exception"
	"os"
)

func FindProvinsi(p string) entity.Wilayah {
	file, err := os.OpenFile("/app/assets/PROVINSI.json", os.O_RDWR, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var provinsi []entity.Wilayah
	err = json.NewDecoder(file).Decode(&provinsi)
	exception.PanicIfNeeded(err)

	for _, provinsi := range provinsi {
		if provinsi.Kode == p {
			return provinsi
		}
	}

	return entity.Wilayah{}
}

func FindKota(k string) entity.Wilayah {
	file, err := os.OpenFile("/app/assets/KOTA.json", os.O_RDWR, 0644)
	exception.PanicIfNeeded(err)

	defer file.Close()
	var kota []entity.Wilayah
	err = json.NewDecoder(file).Decode(&kota)
	exception.PanicIfNeeded(err)

	for _, kota := range kota {
		if kota.Nama == k {
			return kota
		}
	}

	return entity.Wilayah{}
}
