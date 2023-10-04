package contact

// import (
// 	"gopkg.in/gomail.v2"
// )

// func InitSMTPDialer() *gomail.Dialer {
//     // Ganti dengan konfigurasi sesuai SMTP server Anda.
//     return gomail.NewDialer("smtp.gmail.com", 587, "raihanalfarisi2@gmail.com", "cvjpkxpdbhhlxjzc")
// }

// func SendEmail(to, subject, body string) error {
//     // Buat pesan email baru.
//     email := gomail.NewMessage()
//     email.SetHeader("From", "noreply@example.com")
//     email.SetHeader("To", to)
//     email.SetHeader("Subject", subject)
//     email.SetBody("text/plain", body)

//     // Ambil Dialer SMTP yang telah dikonfigurasi.
//     d := InitSMTPDialer()

//     // Kirim email.
//     if err := d.DialAndSend(email); err != nil {
//         return err
//     }
//     return nil
// }
