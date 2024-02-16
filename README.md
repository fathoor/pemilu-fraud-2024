## Pemilu Fraud 2024 API
### Fraud Detection for Pemilu 2024

***Saat ini hanya mampu melakukan pengecekan pada tingkat Kabupaten/Kota***

### Request

```http
curl -X GET "https://pemilu-fraud.fathoor.cloud/?kota=[NAMA KOTA]"
```

```shell
...
[Info] [16/20] CHECKING TPS 016 KECAMATAN ASEM ROWO KELURAHAN TAMBAK SARIOSO
[Info] [17/20] CHECKING TPS 017 KECAMATAN ASEM ROWO KELURAHAN TAMBAK SARIOSO
[Info] [18/20] CHECKING TPS 018 KECAMATAN ASEM ROWO KELURAHAN TAMBAK SARIOSO
[Info] [19/20] CHECKING TPS 019 KECAMATAN ASEM ROWO KELURAHAN TAMBAK SARIOSO
[Info] [20/20] CHECKING TPS 020 KECAMATAN ASEM ROWO KELURAHAN TAMBAK SARIOSO
[Info] [1/63] CHECKING TPS 001 KECAMATAN BENOWO KELURAHAN KANDANGAN
[Info] [2/63] CHECKING TPS 002 KECAMATAN BENOWO KELURAHAN KANDANGAN
[Info] [3/63] CHECKING TPS 003 KECAMATAN BENOWO KELURAHAN KANDANGAN
[Info] [1] ANOMALY FOUND AT KECAMATAN BENOWO KELURAHAN KANDANGAN TPS 003
[Info] [4/63] CHECKING TPS 004 KECAMATAN BENOWO KELURAHAN KANDANGAN
...
[Info] FOUND A TOTAL OF [28] ANOMALIES AT KOTA SURABAYA
```

### Response

```json
{
  "kode": "3204321002073", 
  "timestamp": "2024-02-14 21:46:01", 
  "suara": {
    "suara_sah": 216, 
    "suara_tidak_sah": 6, 
    "suara_total": 224
  }, 
  "hasil": {
    "01": 73, 
    "02": 138, 
    "03": 7
  }, 
  "provinsi": "JAWA BARAT", 
  "kota": "BANDUNG", 
  "kecamatan": "BALEENDAH", 
  "kelurahan": "ANDIR", 
  "tps": "TPS 073", 
  "images": [
    "https://sirekap-obj-formc.kpu.go.id/59ff/pemilu/ppwp/32/04/32/10/02/3204321002073-20240214-193447--b82d0577-ea3a-45b8-ba7f-c2167cb491e9.jpg", 
    "https://sirekap-obj-formc.kpu.go.id/59ff/pemilu/ppwp/32/04/32/10/02/3204321002073-20240214-193712--a808fddd-30f6-4a6c-952d-d0f44f5a00f0.jpg", 
    "https://sirekap-obj-formc.kpu.go.id/59ff/pemilu/ppwp/32/04/32/10/02/3204321002073-20240214-193933--123b8674-c4a9-4815-9385-6e8f1d86e0b0.jpg"
  ], 
  "url": "https://pemilu2024.kpu.go.id/pilpres/hitung-suara/32/3204/320432/3204321002/3204321002073"
}
```

### Endpoint
| Method  | Endpoint                               | Query | Description            |
|:-------:|:---------------------------------------|:-----:|:-----------------------|
|   GET   | https://pemilu-fraud.fathoor.cloud/ | kota  | Get latest cached data |
