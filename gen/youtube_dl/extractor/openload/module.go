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
 * openload/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/openload.py
 */

package openload

import (
	Ωbrowser "github.com/tenta-browser/go-video-downloader/gen/lib/browser"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrowserWrapper   λ.Object
	ExtractorError   λ.Object
	PhantomJSwrapper λ.Object
	ϒcompat_kwargs   λ.Object
	ϒstd_headers     λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ExtractorError = Ωutils.ExtractorError
		ϒstd_headers = Ωutils.ϒstd_headers
		BrowserWrapper = Ωbrowser.BrowserWrapper
		PhantomJSwrapper = λ.Cal(λ.TypeType, λ.NewStr("PhantomJSwrapper"), λ.NewTuple(BrowserWrapper), func() λ.Dict {

			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
	})
}
