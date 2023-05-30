package main

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

// 長度
const uuidCount = 6

func Uuid() string {
	id := uuid.New()
	str := base64.StdEncoding.EncodeToString(id[:])
	str = strings.ReplaceAll(str, "/", "_")
	return str[:uuidCount]
}
