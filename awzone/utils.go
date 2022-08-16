package awzone

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import "strings"

func isBoringMimeType(mime string) bool {
	if strings.HasPrefix(mime, "text/plain") {
		return true
	}

	if strings.HasPrefix(mime, "application/octet-stream") {
		return true
	}

	return false
}
