package main

/*
#cgo CFLAGS: -IC:/ffmpeg/include
#cgo LDFLAGS: -LC:/ffmpeg/lib/ -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavutil/opt.h>
#include <libavdevice/avdevice.h>

static const AVStream *go_av_streams_get(const AVStream **streams,unsigned int n)
{
    return streams[n];
}
*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println(C.avdevice_version())
}
