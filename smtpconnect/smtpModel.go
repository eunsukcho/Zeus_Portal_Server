package smtpconnect

type SmtpInfo struct {
	AdminAddress   string `json:"AdminAddress" binding:"required"`
	SmtpAddress    string `json:"SmtpAddress" binding:"required"`
	Port           string `json:"Port" binding:"required"`
	Password       string `json:"Password" binding:"required"`
	Authentication string `json:"Authentication" binding:"required"`
}
