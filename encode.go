// Package aac provides AAC codec encoder based on [VisualOn AAC encoder](https://github.com/mstorsjo/vo-aacenc) library.
package aac

// #include <stdlib.h>
import "C"

import (
	"io"
	"unsafe"

	"github.com/aam335/aac-go/aacenc"
)

// Options represent encoding options.
type Options struct {
	// Audio file sample rate
	SampleRate int
	// Encoder bit rate in bits/sec
	BitRate int
	// Number of channels on input (1,2)
	NumChannels int
}

// Encoder type.
type Encoder struct {
	w      io.Writer
	aacEnc *aacenc.Encoder
	insize int
	inbuf  []byte
	outbuf []byte
}

// NewEncoder returns new AAC encoder.
func NewEncoder(w io.Writer, opts *Options) (e *Encoder, err error) {
	e = &Encoder{}
	e.w = w
	e.aacEnc = aacenc.New()

	if opts.BitRate == 0 {
		opts.BitRate = 64000
	}

	ret := e.aacEnc.Init(aacenc.VoAudioCodingAac)
	err = aacenc.ErrorFromResult(ret)
	if err != nil {
		return
	}

	var params aacenc.AacencParam
	params.SampleRate = int32(opts.SampleRate)
	params.BitRate = int32(opts.BitRate)
	params.NChannels = int16(opts.NumChannels)
	params.AdtsUsed = 1

	ret = e.aacEnc.SetParam(aacenc.VoPidAacEncparam, unsafe.Pointer(&params))
	err = aacenc.ErrorFromResult(ret)
	if err != nil {
		return
	}

	e.insize = int(opts.NumChannels) * 2 * 1024

	e.inbuf = make([]byte, e.insize)
	e.outbuf = make([]byte, 20480)

	return
}

// Encode encodes data from reader.
func (e *Encoder) Encode(r io.Reader) (err error) {
	var outinfo aacenc.VoAudioOutputinfo
	var input, output aacenc.VoCodecBuffer
	var n int
	var prevRead int
loop:
	for {
		n, err = r.Read(e.inbuf[prevRead:])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		prevRead += n
		if prevRead < e.insize {
			continue
		}
		n = prevRead
		prevRead = 0

		input.Buffer = C.CBytes(e.inbuf)
		input.Length = uint64(n)

		ret := e.aacEnc.SetInputData(&input)
		err = aacenc.ErrorFromResult(ret)
		if err != nil {
			break loop
		}

		output.Buffer = C.CBytes(e.outbuf)
		output.Length = uint64(len(e.outbuf))

		ret = e.aacEnc.GetOutputData(&output, &outinfo)
		err = aacenc.ErrorFromResult(ret)
		if err != nil {
			break loop
		}

		_, err = e.w.Write(C.GoBytes(output.Buffer, C.int(output.Length)))
		if err != nil {
			break loop
		}
		C.free(input.Buffer)
		input.Buffer = nil
		C.free(output.Buffer)
		output.Buffer = nil
	}

	if input.Buffer != nil {
		C.free(input.Buffer)
	}
	if output.Buffer != nil {
		C.free(output.Buffer)
	}

	return nil
}

// Close closes encoder.
func (e *Encoder) Close() error {
	ret := e.aacEnc.Uninit()
	return aacenc.ErrorFromResult(ret)
}
