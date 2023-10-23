package email

import (
	"fmt"
	"strings"
)

var websiteName = "cashr"

func verificationEmail(email, verificationLink string) string {
	template := fmt.Sprintf(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html lang="en">

  <head data-id="__react-email-head">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  </head>
  <div id="__react-email-preview" style="display:none;overflow:hidden;line-height:1px;opacity:0;max-height:0;max-width:0">Verify your %s account</div>

  <body data-id="__react-email-body" style="background-color:rgb(255,255,255);margin-top:auto;margin-bottom:auto;margin-left:auto;margin-right:auto;font-family:ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji">
    <table align="center" width="100%%" data-id="__react-email-container" role="presentation" cellSpacing="0" cellPadding="0" border="0" style="max-width:37.5em;border-width:1px;border-style:solid;border-color:rgb(234,234,234);border-radius:0.25rem;margin-top:40px;margin-bottom:40px;margin-left:auto;margin-right:auto;padding:20px;width:420px">
      <tbody>
        <tr style="width:100%%">
          <td>
            <h1 data-id="react-email-heading" style="color:rgb(41,37,36);font-size:24px;font-weight:400;text-align:center;padding:0px;margin-top:30px;margin-left:0px;margin-right:0px"><strong>%s</strong></h1>
            <h1 data-id="react-email-heading" style="font-size:0.875rem;line-height:1.25rem;font-weight:400;text-align:center;margin-bottom:24px">One step closer to full financial control.</h1>
            <p data-id="react-email-text" style="font-size:14px;line-height:24px;margin:16px 0;color:rgb(41,37,36)">Hello,</p>
            <p data-id="react-email-text" style="font-size:14px;line-height:24px;margin:16px 0;color:rgb(41,37,36)">Thanks for joining <strong>%s</strong>, the easiest and most convenient way to manage your day-to-day expenses and income.</p>
            <table align="center" width="100%%" data-id="react-email-section" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="text-align:center;margin-top:32px;margin-bottom:32px">
              <tbody>
                <tr>
                  <td><a href="%s" data-id="react-email-button" target="_blank" style="line-height:100%%;text-decoration:none;display:inline-block;max-width:100%%;padding:12px 20px;background-color:rgb(41,37,36);border-radius:0.25rem;color:rgb(255,255,255);font-size:12px;font-weight:600;text-decoration-line:none;text-align:center"><span></span><span style="max-width:100%%;display:inline-block;line-height:120%%;mso-padding-alt:0px;mso-text-raise:9px">Verify account</span><span></span></a></td>
                </tr>
              </tbody>
            </table>
            <p data-id="react-email-text" style="font-size:14px;line-height:24px;margin:16px 0;color:rgb(41,37,36);overflow-wrap:normal;word-break:normal;max-width:380px">or copy and paste this URL into your browser: <a href="%s" data-id="react-email-link" target="_blank" style="color:rgb(37,99,235);text-decoration:none;text-decoration-line:none;overflow-wrap:break-word">%s</a></p>
            <hr data-id="react-email-hr" style="width:100%%;border:none;border-top:1px solid #eaeaea;border-width:1px;border-style:solid;border-color:rgb(234,234,234);margin-top:26px;margin-bottom:26px;margin-left:0px;margin-right:0px" />
            <p data-id="react-email-text" style="font-size:12px;line-height:24px;margin:16px 0;color:rgb(102,102,102)">This invitation was intended for <span style="color:rgb(41,37,36)">%s</span>. If you were not expecting this invitation, you can ignore this email. If you are concerned about your account&#x27;s safety, please reply to this email to get in touch with us.</p>
          </td>
        </tr>
      </tbody>
    </table>
  </body>

</html>
`, websiteName, websiteName, websiteName, verificationLink, verificationLink, verificationLink, email)
	return template
}

func verificationEmailText(email, verificationLink string) string {
	template := fmt.Sprintf(`
%s
ONE STEP CLOSER TO FULL FINANCIAL CONTROL.

Hello,

Thanks for joining %s, the easiest and most convenient way to manage your
day-to-day expenses and income.

Verify account [%s]

or copy and paste this URL into your browser:
%s [%s]

--------------------------------------------------------------------------------

This invitation was intended for %s. If you were not expecting
this invitation, you can ignore this email. If you are concerned about your
account's safety, please reply to this email to get in touch with us.`, strings.ToUpper(websiteName), websiteName, verificationLink, verificationLink, verificationLink, email)
	return template
}

func resetPasswordEmail(email, resetPasswordLink string) string {
	template := fmt.Sprintf(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html lang="en">
  <head data-id="__react-email-head">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  </head>
  <div
    id="__react-email-preview"
    style="
      display: none;
      overflow: hidden;
      line-height: 1px;
      opacity: 0;
      max-height: 0;
      max-width: 0;
    "
  >
    Verify your %s account
  </div>

  <body
    data-id="__react-email-body"
    style="
      background-color: rgb(255, 255, 255);
      margin-top: auto;
      margin-bottom: auto;
      margin-left: auto;
      margin-right: auto;
      font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont,
        Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans, sans-serif,
        Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol, Noto Color Emoji;
    "
  >
    <table
      align="center"
      width="100%%"
      data-id="__react-email-container"
      role="presentation"
      cellspacing="0"
      cellpadding="0"
      border="0"
      style="
        max-width: 37.5em;
        border-width: 1px;
        border-style: solid;
        border-color: rgb(234, 234, 234);
        border-radius: 0.25rem;
        margin-top: 40px;
        margin-bottom: 40px;
        margin-left: auto;
        margin-right: auto;
        padding: 20px;
        width: 420px;
      "
    >
      <tbody>
        <tr style="width: 100%%">
          <td>
            <h1
              data-id="react-email-heading"
              style="
                color: rgb(41, 37, 36);
                font-size: 24px;
                font-weight: 400;
                text-align: center;
                padding: 0px;
                margin-top: 30px;
                margin-left: 0px;
                margin-right: 0px;
              "
            >
              <strong>cashr</strong>
            </h1>
            <p
              data-id="react-email-text"
              style="
                font-size: 14px;
                line-height: 24px;
                margin: 16px 0;
                color: rgb(41, 37, 36);
              "
            >
              Hello,
            </p>
            <p
              data-id="react-email-text"
              style="
                font-size: 14px;
                line-height: 24px;
                margin: 16px 0;
                color: rgb(41, 37, 36);
              "
            >
              We see that you want to reset your password. Please click the
              button below to reset your password.
            </p>
            <table
              align="center"
              width="100%%"
              data-id="react-email-section"
              border="0"
              cellpadding="0"
              cellspacing="0"
              role="presentation"
              style="text-align: center; margin-top: 32px; margin-bottom: 32px"
            >
              <tbody>
                <tr>
                  <td>
                    <a
                      href="%s"
                      data-id="react-email-button"
                      target="_blank"
                      style="
                        line-height: 100%%;
                        text-decoration: none;
                        display: inline-block;
                        max-width: 100%%;
                        padding: 12px 20px;
                        background-color: rgb(41, 37, 36);
                        border-radius: 0.25rem;
                        color: rgb(255, 255, 255);
                        font-size: 12px;
                        font-weight: 600;
                        text-decoration-line: none;
                        text-align: center;
                      "
                      ><span></span
                      ><span
                        style="
                          max-width: 100%%;
                          display: inline-block;
                          line-height: 120%%;
                          mso-padding-alt: 0px;
                          mso-text-raise: 9px;
                        "
                        >Reset password</span
                      ><span></span
                    ></a>
                  </td>
                </tr>
              </tbody>
            </table>
            <p
              data-id="react-email-text"
              style="
                font-size: 14px;
                line-height: 24px;
                margin: 16px 0;
                color: rgb(41, 37, 36);
                overflow-wrap: normal;
                word-break: normal;
                max-width: 380px;
              "
            >
              or copy and paste this URL into your browser:
              <a
                href="%s"
                data-id="react-email-link"
                target="_blank"
                style="
                  color: rgb(37, 99, 235);
                  text-decoration: none;
                  text-decoration-line: none;
                  overflow-wrap: break-word;
                "
                >%s</a
              >
            </p>
            <hr
              data-id="react-email-hr"
              style="
                width: 100%%;
                border: none;
                border-top: 1px solid #eaeaea;
                border-width: 1px;
                border-style: solid;
                border-color: rgb(234, 234, 234);
                margin-top: 26px;
                margin-bottom: 26px;
                margin-left: 0px;
                margin-right: 0px;
              "
            />
            <p
              data-id="react-email-text"
              style="
                font-size: 12px;
                line-height: 24px;
                margin: 16px 0;
                color: rgb(102, 102, 102);
              "
            >
              This invitation was intended for
              <span style="color: rgb(41, 37, 36)">%s</span>. If
              you were not expecting this invitation, you can ignore this email.
              If you are concerned about your account&#x27;s safety, please
              reply to this email to get in touch with us.
            </p>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`, websiteName, resetPasswordLink, resetPasswordLink, resetPasswordLink, email)

	return template
}

func resetPasswordEmailText(email, resetPasswordLink string) string {
	template := fmt.Sprintf(`
%s

Hello,

We see that you want to reset your password. Please click the button below to
reset your password.

Reset password [%s]

or copy and paste this URL into your browser:
%s
[%s]

--------------------------------------------------------------------------------

This invitation was intended for %s. If you were not expecting
this invitation, you can ignore this email. If you are concerned about your
account's safety, please reply to this email to get in touch with us.
`, strings.ToUpper(websiteName), resetPasswordLink, resetPasswordLink, resetPasswordLink, email)
	return template
}
