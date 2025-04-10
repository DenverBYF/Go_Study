package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"example.com/m/v2/lua"
	"github.com/mozillazg/go-pinyin"
)

const (
	mockPriKey = `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMB1JhhkzETa4GcK
13MWhPQdq61vpWAyPFF8iFW9SIIKUg0slTr/hzpEaEfS08G87QDRD7DdMnnGjCpJ
tlwm2gO86swMaTcDlLGHJyMqe3QLWft07wWJkXRlzOLCToik5AJ2f8gCz6OX19Z0
x739ZiCvswIBnjMZyMJ8x/9lkMytAgMBAAECgYA1wULXs2e+n6foAAY3dgXCrXrj
IXj+immAlRj7Ybgb4kgSt+4ioXai1rKNxRwyU9Oa08nfK5lKnCNQOHCAArYHssrr
//2I9CC3ABgQIIJyo//NtCFzCrZJmYc16Nw++ojJtxDqphQZ0SYFqYwqTh9pYYm3
+nTNytziR5/MNXR8vQJBAPHtDDjsuRNYLQUNcoeTZ/2onkt8vMt5vkrYkNWoKqpE
XcYHsB8YbpIa2YYMoskgOlFa1OX2dpaM0ORd1i9LvVsCQQDLp2Ea8YV0Bie3/6V3
1nrBJTU4j/VcOB/yv9rv2a5lsqx7LS2mj8i+wFjdFXsjlCBWjaJ0sDjRdUuVVjv9
mxSXAkEAzk22QYjlHrDv20Ina6PxZyd2rarWmLHd65eYkwqQL5iTWv9NWocMK33I
B38ZBmh8MspBiUVOxX2Z2VSIBNcSxwJAd/t131zZ5iBWODY6c17+VVqpf1h5EsrS
L79Oqq6R68KXkb5tPctKVu+VqzMjqDN11eh+BMdpwiWb0TDMCT7bowJAHvsAJhun
bUrjsZou/LJ/NwSxKXwepquy2FjtlPHsVPT55aw4zOQeL2JtyWfjXygnXDSaj4hg
wRe5gJ2GGI/eFA==
-----END PRIVATE KEY-----`
	mockPubKey = `-----BEGIN CERTIFICATE-----
MIICJzCCAZACCQDxJI6TbIUHIDANBgkqhkiG9w0BAQsFADBYMQswCQYDVQQGEwJh
ZDELMAkGA1UECAwCZGQxDDAKBgNVBAcMA2FzZDEMMAoGA1UECgwDc2FkMQwwCgYD
VQQLDANhc2QxEjAQBgkqhkiG9w0BCQEWA2FzZDAeFw0yMTA3MDUwNDEzMjVaFw0y
MjA3MDUwNDEzMjVaMFgxCzAJBgNVBAYTAmFkMQswCQYDVQQIDAJkZDEMMAoGA1UE
BwwDYXNkMQwwCgYDVQQKDANzYWQxDDAKBgNVBAsMA2FzZDESMBAGCSqGSIb3DQEJ
ARYDYXNkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDAdSYYZMxE2uBnCtdz
FoT0Hautb6VgMjxRfIhVvUiCClINLJU6/4c6RGhH0tPBvO0A0Q+w3TJ5xowqSbZc
JtoDvOrMDGk3A5SxhycjKnt0C1n7dO8FiZF0Zcziwk6IpOQCdn/IAs+jl9fWdMe9
/WYgr7MCAZ4zGcjCfMf/ZZDMrQIDAQABMA0GCSqGSIb3DQEBCwUAA4GBALG2tanW
rvdxoUcsD3grD1JhpOidTkpLEJhzLyX/PDCvbiHtBBGJZjeZXORldHK3Sy0TqnX2
HQ5Q16UXeaHtlchHUoECp7fRUMTRbbitk1zms4RL1znazAsyhB+rkaDMCoABXhPQ
ZDYK/B4LP8zxC6aXk2rkSyZSkrtpkf57ZIYe
-----END CERTIFICATE-----`
)

func main() {
	lua.LuaTest()
	/*mockPk, err := ncrypto.DecodePrivateKey([]byte(mockPriKey)).PKCS8().RSAPrivateKey()
	fmt.Println(mockPk, err)
	a := "asdasdasdasdasd"

	sign, err := ncrypto.RSASignPKCS1v15([]byte(a), mockPk, crypto.SHA256)
	if err != nil {
		fmt.Println("Error signing:", err)
	}

	h := base64.StdEncoding.EncodeToString(sign)
	fmt.Println()

	cert, err := ncrypto.DecodeCertificate([]byte(mockPubKey))
	if err != nil {
		fmt.Println("Error decoding certificate:", err)
	}
	pub, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Invalid public key type")
	}
	fmt.Println(pub)

	fmt.Println(ncrypto.RSAVerifyPKCS1v15([]byte(a), []byte(h), pub, crypto.SHA256))
	return*/
	/*iconCsv, err := os.Open("icon_url.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer iconCsv.Close()
	iconMap := make(map[string]string)
	reader := csv.NewReader(iconCsv)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}
	for _, record := range records {
		iconMap[record[0]] = record[1]
	}

	f, err := os.Open("cardbin.xlsx")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	ef, err := excelize.OpenReader(f)
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return
	}
	rows, err := ef.GetRows("整理版")

	resFile, err := os.Create("bankcard.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer resFile.Close()
	writer := csv.NewWriter(resFile)
	defer writer.Flush()
	for _, row := range rows {
		newRow := make([]string, 0)
		bankName := strings.Trim(strings.Split(row[0], "(")[0], "\n")
		py := toPinYin(bankName) + ".png"
		newRow = append(newRow, py, bankName)
		for _, val := range row[1:] {
			newRow = append(newRow, strings.Trim(val, "\n"))
		}
		icon := iconMap[py]
		newRow = append(newRow, icon)
		err = writer.Write(newRow)
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}*/

}

func renameFilesInDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}
		originalFileName := file.Name()
		newFileName := generateNewFileName(originalFileName)
		if newFileName != "" {
			oldPath := filepath.Join(dir, originalFileName)
			newPath := filepath.Join(dir, newFileName)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
			fmt.Printf("Renamed \"%s\" to \"%s\"\n", originalFileName, newFileName)
		}
	}

	return nil
}

func toPinYin(s string) string {
	a := pinyin.NewArgs()
	list := pinyin.Pinyin(s, a)

	res := ""
	for _, l := range list {
		for _, ll := range l {
			res = res + ll
		}
		res += "_"
	}
	return strings.Replace(res, "yin_xing_", "yin_hang", -1)
}

func generateNewFileName(originalFileName string) string {
	/*a := pinyin.NewArgs()
	list := pinyin.Pinyin(originalFileName, a)

	res := ""
	for _, l := range list {
		for _, ll := range l {
			res = res + ll
		}
		res += "_"
	}
	return res*/
	//return strings.Replace(originalFileName, "yin_han", "yin_hang", -1)
	return originalFileName
}
