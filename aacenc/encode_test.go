package aacenc

import (
	"log"
	"testing"
)

func TestBridge(t *testing.T) {
	channels := 2
	properLen := 1024 * channels
	m := make([]int16, properLen, properLen)
	aacEnc := New()
	aacEnc.Init(VoAudioCodingAac)
	ret := aacEnc.SetParamAac(44100, channels)
	if ret != VoErrNone {
		t.Error("SetParamAac failed")
	}

	b, errn := aacEnc.EncodePcmBlock(m)
	if errn != VoErrNone {
		t.Error("EncodePcmBlock failed on proper len", ErrorFromResult(errn))
	}
	log.Print("Encoded: ", len(b))
	b, errn = aacEnc.EncodePcmBlock(m[1:])
	if errn == VoErrNone {
		t.Error("EncodePcmBlock failed on wrong len", ErrorFromResult(errn))
	}

}
