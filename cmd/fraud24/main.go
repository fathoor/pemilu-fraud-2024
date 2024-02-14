package main

import (
	"flag"
	"fmt"
	"github.com/fathoor/fraud24/cmd/fraud24/service"
	"os"
)

func main() {
	kode := flag.String("k", "", "Kode TPS (13 digit)")
	flag.Parse()

	if *kode == "" {
		fmt.Println("Error: Membutuhkan kode TPS sebagai argumen (-k)")
		os.Exit(1)
	}

	fraud := service.ProvideFraudService()
	result := fraud.FraudCheck(*kode)

	if result != "" {
		fmt.Println(result)
	} else {
		fmt.Println("Tidak terdeteksi kecurangan")
	}

	os.Exit(0)
}
