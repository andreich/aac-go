package aac

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/youpy/go-wav"
)

func TestEncode(t *testing.T) {
	file, err := os.Open(filepath.Join("testdata", "test.wav"))
	if err != nil {
		t.Fatal(err)
	}

	wr := wav.NewReader(file)
	f, err := wr.Format()
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	opts := &Options{}
	opts.SampleRate = int(f.SampleRate)
	opts.NumChannels = int(f.NumChannels)

	enc, err := NewEncoder(buf, opts)
	if err != nil {
		t.Fatal(err)
	}

	err = enc.Encode(wr)
	if err != nil {
		t.Error(err)
	}

	err = enc.Close()
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile(filepath.Join(os.TempDir(), "test.aac"), buf.Bytes(), 0644)
	if err != nil {
		t.Error(err)
	}

	if want, got := 8192, len(buf.Bytes()); want != got {
		t.Errorf("encoded file length %d is different from expected length %d", got, want)
	}
}

type testReader struct {
	bufLen int
	from   io.Reader
	adjust func() int
}

func (t *testReader) Read(in []byte) (int, error) {
	defer func() {
		t.bufLen = t.adjust()
	}()
	if t.bufLen > len(in) {
		t.bufLen = len(in)
	}
	if t.bufLen == 0 {
		// This is not EOF, just unavailable data at this point.
		return 0, nil
	}
	buf := make([]byte, t.bufLen)
	n, err := t.from.Read(buf)
	if n > 0 {
		copy(in, buf[:n])
	}
	return n, err
}

func TestEncodeVariableReadLength(t *testing.T) {
	file, err := os.Open(filepath.Join("testdata", "test.wav"))
	if err != nil {
		t.Fatal(err)
	}

	wr := wav.NewReader(file)
	f, err := wr.Format()
	if err != nil {
		t.Fatal(err)
	}

	tr := &testReader{
		bufLen: 1,
		from:   wr,
	}
	bufLenPrev := 1
	tr.adjust = func() int {
		bufLenPrev += 1
		if bufLenPrev%5 == 0 {
			return 0
		}
		return bufLenPrev
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	opts := &Options{}
	opts.SampleRate = int(f.SampleRate)
	opts.NumChannels = int(f.NumChannels)

	enc, err := NewEncoder(buf, opts)
	if err != nil {
		t.Fatal(err)
	}

	err = enc.Encode(tr)
	if err != nil {
		t.Error(err)
	}

	err = enc.Close()
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile(filepath.Join(os.TempDir(), "test-2.aac"), buf.Bytes(), 0644)
	if err != nil {
		t.Error(err)
	}

	if want, got := 8192, len(buf.Bytes()); want != got {
		t.Errorf("encoded file length %d is different from expected length %d", got, want)
	}
}
