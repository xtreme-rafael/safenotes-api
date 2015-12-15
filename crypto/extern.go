package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

var RandRead = rand.Read
var Base64EncodeToString = base64.URLEncoding.EncodeToString
