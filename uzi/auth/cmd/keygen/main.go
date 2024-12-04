package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определяем флаги командной строки
	publicKeyPath := flag.String("public", "public_key.pem", "Путь для сохранения открытого ключа")
	privateKeyPath := flag.String("private", "private_key.pem", "Путь для сохранения закрытого ключа")

	// Парсинг флагов
	flag.Parse()

	// Генерируем пару ключей RSA
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Ошибка генерации закрытого ключа:", err)
		return
	}

	// Получаем открытый ключ из закрытого
	publicKey := &privateKey.PublicKey

	// Сохраняем закрытый ключ в PEM-формате
	privateKeyFile, err := os.Create(*privateKeyPath)
	if err != nil {
		fmt.Println("Ошибка создания файла закрытого ключа:", err)
		return
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		fmt.Println("Ошибка кодирования закрытого ключа:", err)
		return
	}

	// Сохраняем открытый ключ в PEM-формате
	publicKeyFile, err := os.Create(*publicKeyPath)
	if err != nil {
		fmt.Println("Ошибка создания файла открытого ключа:", err)
		return
	}
	defer publicKeyFile.Close()

	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Ошибка кодирования открытого ключа:", err)
		return
	}

	err = pem.Encode(publicKeyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})
	if err != nil {
		fmt.Println("Ошибка кодирования открытого ключа:", err)
		return
	}

	fmt.Println("Ключи успешно сгенерированы и сохранены в файлы")
	fmt.Printf("- Закрытый ключ: %s\n", *privateKeyPath)
	fmt.Printf("- Открытый ключ: %s\n", *publicKeyPath)
}
