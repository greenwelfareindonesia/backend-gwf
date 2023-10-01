// package email

// import (
// 	"gopkg.in/gomail.v2"
// )

// func InitSMTPDialer() *gomail.Dialer {
//     // Ganti dengan konfigurasi sesuai SMTP server Anda.
//     return gomail.NewDialer("smtp.gmail.com", 587, "raihanalfarisi2@gmail.com", "obmhkhrsrckkcxni")
// }

// // func InitSMTPDialer(smtpUsername, smtpPassword string) *gomail.Dialer {
// //     // Ganti dengan konfigurasi sesuai SMTP server Anda.
// //     return gomail.NewDialer("smtp.gmail.com", 587, smtpUsername, smtpPassword)
// // }

// // func SendEmail(from, to, subject, body, smtpUsername, smtpPassword string) error {
// //     // Buat pesan email baru.
// //     email := gomail.NewMessage()
// //     email.SetHeader("From", from)
// //     email.SetHeader("To", to)
// //     email.SetHeader("Subject", subject)
// //     email.SetBody("text/plain", body)

// //     // Ambil Dialer SMTP yang telah dikonfigurasi dengan alamat email pengirim.
// //     d := InitSMTPDialer(smtpUsername, smtpPassword)

// //     // Kirim email.
// //     if err := d.DialAndSend(email); err != nil {
// //         return err
// //     }
// //     return nil
// // }

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

package email

import (
	"gopkg.in/gomail.v2"
)

func InitSMTPDialer() *gomail.Dialer {
    // Ganti dengan konfigurasi sesuai SMTP server Anda.
    return gomail.NewDialer("smtp.gmail.com", 587, "raihanalfarisi2@gmail.com", "obmhkhrsrckkcxni")
}

func SendEmail(to, subject, name string, from string, body string) error {
    // Buat pesan email baru.
    email := gomail.NewMessage()
    email.SetHeader("From", "noreply@example.com")
    email.SetHeader("To", "raihanalfarisi2@gmail.com")
    email.SetHeader("Subject", subject)

    // email.SetBody("text/plain", from)
    email.SetBody("text/plain", "Email        : " + from + "\n" + "nama        : " + name + "\n" + "Message  : " + body)

    // Ambil Dialer SMTP yang telah dikonfigurasi.
    d := InitSMTPDialer()

    // Kirim email.
    if err := d.DialAndSend(email); err != nil {
        return err
    }
    return nil
}