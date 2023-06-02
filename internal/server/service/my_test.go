package service

import (
	"testing"
	"vediomeeting/internal/helper"
)

func TestName(t *testing.T) {
	println(helper.GetMd5("123456"))
}
