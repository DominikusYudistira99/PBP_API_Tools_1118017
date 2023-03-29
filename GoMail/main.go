// package main

// Masalah Less secure apps
// import (
// 	"fmt"
// 	"log"

// 	mail "github.com/xhit/go-simple-mail"
// )

// func main() {
// 	// Create a new email message
// 	email := mail.NewMSG()

// 	// Set the subject and body of the message
// 	email.SetSubject("Email konfirmasi Pesan Masuk ")
// 	email.SetBody(mail.TextHTML, "<h1>Haloo Yudis,</h1><p>Selamat datang di Email Mahasiswa ITHB.</p>")

// 	// Set the sender and recipient of the message
// 	email.SetFrom("emailmahasiswa@gmail.com").
// 		AddTo("yudistiraprimapradana@gmail.com")

// 	// Set up the SMTP client to send the message
// 	server := mail.NewSMTPClient()
// 	server.Host = "smtp.gmail.com"
// 	server.Port = 587
// 	server.Username = "emailmahasiswa@gmail.com"
// 	server.Password = "12345678"
// 	server.Encryption = mail.EncryptionTLS

// 	// Send the email
// 	err := email.SendWith(server)
// 	if err != nil {
// 		log.Fatalf("failed to send email: %v", err)
// 	}

// 	fmt.Println("Email sent successfully!")
