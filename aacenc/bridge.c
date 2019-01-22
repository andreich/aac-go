// #include <stdint.h>
// #include <stdio.h>
#include "voAAC.h"

// function hooks "panic: runtime error: cgo argument has Go pointer to Go pointer"
VO_U32 encode(VO_HANDLE hCodec,void *pcm, int pcmlen, void *aac, int aacBufferSize, VO_U32 *err){
	VO_CODECBUFFER input,output;
    VO_AUDIO_OUTPUTINFO outInfo;
	input.Buffer=pcm;
    input.Length=pcmlen;
    VO_U32 ret=voAACEncSetInputData(hCodec, &input);
    if (ret!=VO_ERR_NONE){
        *err=ret;
        return 0;
    }
	output.Buffer=aac;
    output.Length=aacBufferSize;
    ret=voAACEncGetOutputData(hCodec,&output,&outInfo);
    if (ret!=VO_ERR_NONE){
        *err=ret;
        return 0;
    }
    return output.Length;
}
