package dao

import (
	"testing"
)

func TestBtcBannerSetGet(t *testing.T) {
	err := BtcBanner{Volume: 100}.Set()

	if err != nil {
		t.Error(err)
	}

	res, err := BtcBanner{}.Get()
	if err != nil {
		t.Error(err)
	}

	t.Logf("BtcBanner:  %v \n", res)

}
