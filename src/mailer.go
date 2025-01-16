package main

import (
	gomail "gopkg.in/mail.v2"
)

func mailService() {
	m := gomail.NewMessage()

	m.SetHeader("From", "yanisching@gmail.com")
	m.SetHeader("To", "yanisching@gmail.com", "davidjonesgan@gmail.com")
	m.SetAddressHeader("Cc", "yanisching@gmail.com", "yanisching")
	m.SetHeader("Subject", "Set or reset your password")

	m.SetBody("text/html", "Hello <b>Username</b>!<h3>We have received a request to set or reset your password.</h3>http://www.resetpassword.com<br><p>If you did not request to set a new password or change your password, you can disregard this email and your password will remain the same.</p><p>Should you experience any issues accessing your Discovery App account, please contact us immediately.</p>Sincerely,<br>Discovery App")
	m.Attach("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTy3tp8KjQ8wP9i1hyZ0n2WLSbVkxHRwKldfw&s")

	d := gomail.NewDialer("smtp.gmail.com", 587, "yanisching@gmail.com", "yjpyfqdkwkczydef")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

//Method with limitation to send email
// func mailService() {
// 	auth := smtp.PlainAuth("", "yanisching@gmail.com", "yjpyfqdkwkczydef", "smtp.gmail.com")

// 	to := []string{"yanisching@gmail.com"}

// 	msg := []byte("To:yanisching@gmail.com\r\n" + "Subject: Why aren't you using Mailtrap yet?\r\n" + "\r\n" + "Here's the space for our great sales pitch\r\n")

// 	err := smtp.SendMail("smtp.gmail.com:587", auth, "yanisching@gmail.com", to, msg)
// 	if err != nil {
// 		slog.Warn("Fail to send email from server.", "error", err)
// 	}
// }
