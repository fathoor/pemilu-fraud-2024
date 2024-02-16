package validation

import (
	"github.com/fathoor/pemilu-fraud-2024/cmd/entity"
	"github.com/fathoor/pemilu-fraud-2024/internal/exception"
	"github.com/fathoor/pemilu-fraud-2024/internal/util"
	"net/url"
	"strings"
)

func ValidateKota(k string) entity.Wilayah {
	if k == "" {
		panic(exception.BadRequestError{
			Message: "Kota tidak boleh kosong",
		})
	}

	k, err := url.PathUnescape(k)
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Kota tidak valid",
		})
	}

	k = strings.ToUpper(k)
	kota := util.FindKota(k)
	if kota.Kode == "" {
		panic(exception.NotFoundError{
			Message: "Kota tidak ditemukan",
		})
	}

	return kota
}

func ValidateProvinsi(p string) entity.Wilayah {
	if p == "" {
		panic(exception.BadRequestError{
			Message: "Provinsi tidak boleh kosong",
		})
	}

	provinsi := util.FindProvinsi(p)
	if provinsi.Nama == "" {
		panic(exception.NotFoundError{
			Message: "Provinsi tidak ditemukan",
		})
	}

	return provinsi
}
