package main

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	// Inisialisasi konfigurasi pengiriman email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "youremail@example.com")
	mailer.SetHeader("To", "pemulihan7cf@gmail.com")
	mailer.SetHeader("Subject", "Email Masuk")
	mailer.SetBody("text/plain", "Selamat datang di email mahasiswa")

	// Konfigurasi koneksi SMTP
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "youremail@example.com", "yourpassword")

	// Mengirim email
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully.")
}
