package authInterface

import "zeus/credentials"

type ClientAuthInterface interface {
	ClientInit() *credentials.Authdetails
}
