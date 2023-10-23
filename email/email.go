package email

import (
	"github.com/resendlabs/resend-go"
	"log"
	"os"
)

func SendEmailVerificationEmail(email, verificationLink string) {
	var apiKey = os.Getenv("RESEND_API_KEY")
	var client = resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "cashr Team <noreply@cashr.klukode.dev>",
		To:      []string{email},
		Subject: "cashr Verification Email",
		Html:    verificationEmail(email, verificationLink),
		Text:    verificationEmailText(email, verificationLink),
	}

	_, err := client.Emails.Send(params)

	if err != nil {
		log.Fatal(err)
	}
}

func SendResetPasswordEmail(email, resetPasswordLink string) {
	var apiKey = os.Getenv("RESEND_API_KEY")
	var client = resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "cashr Team <noreply@cashr.klukode.dev>",
		To:      []string{email},
		Subject: "Reset your cashr password",
		Html:    resetPasswordEmail(email, resetPasswordLink),
		Text:    resetPasswordEmailText(email, resetPasswordLink),
	}

	_, err := client.Emails.Send(params)

	if err != nil {
		log.Fatal(err)
	}
}
