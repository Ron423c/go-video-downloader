// Code generated by transpiler. DO NOT EDIT.

/**
 * Go Video Downloader
 *
 *    Copyright 2019 Tenta, LLC
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
 * __init__/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/postprocessor/__init__.py
 */

package postprocessor

import (
	Ωffmpeg "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/postprocessor/ffmpeg"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FFmpegPostProcessor λ.Object
)

func init() {
	λ.InitModule(func() {
		FFmpegPostProcessor = Ωffmpeg.FFmpegPostProcessor
	})
}
