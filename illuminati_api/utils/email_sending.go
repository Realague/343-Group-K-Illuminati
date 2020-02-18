package utils

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"errors"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func SendConfirmationEmail(c *gin.Context, user db.User) error {
	token, err := GenerateToken(c, jwt.MapClaims{
		"id": user.Id.Hex(),
	})

	if err {
		return errors.New("internal server error")
	}

	url := config.Config.EmailConfirmationEndpoint + token

	return SendEmail("Confirm Email", createTemplate(url, "Valider mon adresse email", "Cliquez sur le bouton ci-dessous pour valider votre adresse mail"), user.Email)
}

func SendRecoverPasswordEmail(c *gin.Context, user db.User) error {
	token, err := GenerateToken(c, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(config.Config.RecoverPasswordTokenValidityTime * time.Minute).Unix(),
	})

	if err {
		return errors.New("internal server error")
	}

	url := config.Config.RecoverPasswordUrl + token

	return SendEmail("Recover password", createTemplate(url, "Récupérer mon mot de passe", "Cliquez sur le bouton ci-dessous pour récupérer votre mot de passe"), user.Email)
}

func SendEmail(subject string, body string, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Config.Email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer("smtp.gmail.com", 465, config.Config.Email, config.Config.EmailPassword)
	return d.DialAndSend(m)
}

func createTemplate(url string, button string, title string) string {
	return `<table style="width:100%;max-width:600px" width="100%" cellSpacing="0" cellPadding="0" border="0" align="center">
            <tbody>
            <tr>
                <td role="modules-container" style="padding:0px 15px 30px 15px;color:#45494c;text-align:left"
                    width="100%" bgcolor="#F5F8FA" align="left">
                    <table role="module" style="table-layout:fixed"
                           width="100%" cellSpacing="0" cellPadding="0" border="0">
                        <tbody>
                        <tr>
                            <td style="font-size:6px;line-height:10px;background-color:#f5f8fa;padding:30px 0px 30px 0px"
                                valign="top" align="center">
                                <a href="https://layana.eu" target="_blank"
                                   data-saferedirecturl="https://www.google.com/url?q=&amp;source=gmail&amp;ust=1563300891949000&amp;usg=AFQjCNGTlB0c_F6WHQqzTQDCXnHEUuxjRQ"><img
                                    src="https://image.noelshack.com/fichiers/2019/29/2/1563310054-logofondtransparent.png"
                                    alt="Layana"
                                    style="display:block;color:#000;text-decoration:none;font-family:Helvetica,arial,sans-serif;font-size:16px;max-width:40%!important;width:40%;height:auto!important"
                                    width="228" border="0"/> </a>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                    <table role="module"
                           style="padding:015px 30px 0px 30px;background-color:#ffffff;box-sizing:border-box"
                           width="100%" cellSpacing="0" cellPadding="0" border="0" bgcolor="#ffffff" align="center">
                        <tbody>
                        <tr role="module-content">
                            <td valign="top" height="100%">
                                <table
                                    style="width:510.000px;border-spacing:0;border-collapse:collapse;margin:0px 0px 0px 0px"
                                    width="510.000" cellSpacing="0" cellPadding="0" border="0" bgcolor="#ffffff"
                                    align="left">
                                    <tbody>
                                    <tr>
                                        <td style="padding:0px;margin:0px;border-spacing:0">
                                            <table role="module"
                                                   style="table-layout:fixed" width="100%" cellSpacing="0"
                                                   cellPadding="0" border="0">
                                                <tbody>
                                                <tr>
                                                    <td style="padding:0px 0px 0px 0px;line-height:22px;text-align:inherit"
                                                        valign="top" height="100%" bgcolor="">
                                                        <div style="text-align:center"><span style="color:#000000"><span
                                                            style="font-size:20px">` + title + `</span></span><span
                                                            style="color:rgb(69,73,76);font-family:arial;font-size:14px;font-style:normal;font-variant-ligatures:normal;font-variant-caps:normal;font-weight:400">&nbsp;
                                                            <img goomoji="1f447" data-goomoji="1f447"
                                                                 style="margin:0 0.2ex;vertical-align:middle;max-height:24px"
                                                                 alt="👇" src="https://mail.google.com/mail/e/1f447"
                                                                 data-image-whitelisted=""/></span>
                                                        </div>
                                                    </td>
                                                </tr>
                                                </tbody>
                                            </table>
                                            <table role="module" style="table-layout:fixed" width="100%" cellSpacing="0"
                                                   cellPadding="0" border="0">
                                                <tbody>
                                                <tr>
                                                    <td style="padding:30px 030px 30px 030px" align="center">
                                                        <table
                                                            className="m_-2381348781736920438button-css__deep-table___2OZyb m_-2381348781736920438wrapper-mobile"
                                                            style="text-align:center" cellSpacing="0" cellPadding="0"
                                                            border="0">
                                                            <tbody>
                                                            <tr>
                                                                <td className="m_-2381348781736920438inner-td"
                                                                    style="border-radius:6px;font-size:16px;text-align:center;background-color:inherit"
                                                                    bgcolor="#207A38" align="center"><a
                                                                    style="background-color:#207A38;height:px;width:px;font-size:16px;line-height:px;font-family:arial,helvetica,sans-serif;color:#ffffff;padding:12px 18px 12px 18px;text-decoration:none;border-radius:6px;border:1px solid #5a74b7;display:inline-block;border-width:0px;font-weight:700"
                                                                    href="` + url + `"
																	data-saferedirecturl="https://www.google.com/url?q=` + url + `&amp;source=gmail&amp;ust=1563300891949000&amp;usg=AFQjCNGTlB0c_F6WHQqzTQDCXnHEUuxjRQ"
                                                                    target="_blank">` + button + `</a></td>
                                                            </tr>
                                                            </tbody>
                                                        </table>
                                                    </td>
                                                </tr>
                                                </tbody>
                                            </table>
                                        </td>
                                    </tr>
                                    </tbody>
                                </table>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                </td>
            </tr>
            </tbody>
        </table>`
}
