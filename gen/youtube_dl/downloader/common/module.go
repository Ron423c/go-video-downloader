// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * common/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/downloader/common.py
 */

package common

import (
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FileDownloader       λ.Object
	ϒerror_to_compat_str λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒerror_to_compat_str = Ωutils.ϒerror_to_compat_str
		FileDownloader = λ.Cal(λ.TypeType, λ.NewStr("FileDownloader"), λ.NewTuple(λ.ObjectType), func() λ.Dict {
			λ.NewStr("File Downloader class.\n\n    File downloader objects are the ones responsible of downloading the\n    actual video file and writing it to disk.\n\n    File downloaders accept a lot of parameters. In order not to saturate\n    the object constructor with arguments, it receives a dictionary of\n    options instead.\n\n    Available options:\n\n    verbose:            Print additional info to stdout.\n    quiet:              Do not print messages to stdout.\n    ratelimit:          Download speed limit, in bytes/sec.\n    retries:            Number of times to retry for HTTP error 5xx\n    buffersize:         Size of download buffer in bytes.\n    noresizebuffer:     Do not automatically resize the download buffer.\n    continuedl:         Try to continue downloads if possible.\n    noprogress:         Do not print the progress bar.\n    logtostderr:        Log messages to stderr instead of stdout.\n    consoletitle:       Display progress in console window's titlebar.\n    nopart:             Do not use temporary .part files.\n    updatetime:         Use the Last-modified header to set output file timestamps.\n    test:               Download only first bytes to test the downloader.\n    min_filesize:       Skip files smaller than this size\n    max_filesize:       Skip files larger than this size\n    xattr_set_filesize: Set ytdl.filesize user xattribute with expected size.\n    external_downloader_args:  A list of additional command-line arguments for the\n                        external downloader.\n    hls_use_mpegts:     Use the mpegts container for HLS videos.\n    http_chunk_size:    Size of a chunk for chunk-based HTTP downloading. May be\n                        useful for bypassing bandwidth throttling imposed by\n                        a webserver (experimental)\n\n    Subclasses of this one must re-define the real_download method.\n    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
	})
}
