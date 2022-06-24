package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
)

func main() {
	var (
		fileName      = "demo/qrcode/qr.png"
		qrCodeContent = "https://github.com"
	)

	// 生成二维码
	if err := generateQrCode(qrCodeContent, fileName); err != nil {
		fmt.Println(err)
		return
	}

	// 获取二维码内容
	qrMatrix, err := getQrCodeFromFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qrMatrix.Content)
}

func getQrCodeFromFile(fileName string) (*qrcode.Matrix, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}(file)
	qrMatrix, err := qrcode.Decode(file)
	if err != nil {
		return nil, err
	}
	return qrMatrix, nil
}

func generateQrCode(qrCodeContent string, fileName string) error {
	// 生成二维码
	qrCode, err := qr.Encode(qrCodeContent, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	// 调整二维码高度
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return err
	}

	// 创建二维码文件
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}(file)

	// 写入二维码到文件
	if err := png.Encode(file, qrCode); err != nil {
		return err
	}
	return nil
}
