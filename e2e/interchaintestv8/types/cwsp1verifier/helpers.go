package cwsp1verifier

import "encoding/base64"

func ToBinary(b []byte) Binary {
	return Binary(base64.StdEncoding.EncodeToString(b))
}
