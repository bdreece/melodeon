package web

import "mime"

func init() {
	if err := mime.AddExtensionType(".cjs", "text/javascript"); err != nil {
		panic(err)
	}
}
