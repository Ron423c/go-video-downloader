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
 * indavideo/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/indavideo.py
 */

package indavideo

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	IndavideoEmbedIE  λ.Object
	InfoExtractor     λ.Object
	ϒcompat_str       λ.Object
	ϒint_or_none      λ.Object
	ϒparse_age_limit  λ.Object
	ϒparse_iso8601    λ.Object
	ϒupdate_url_query λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_age_limit = Ωutils.ϒparse_age_limit
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		IndavideoEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("IndavideoEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				IndavideoEmbedIE__VALID_URL λ.Object
			)
			IndavideoEmbedIE__VALID_URL = λ.NewStr("https?://(?:(?:embed\\.)?indavideo\\.hu/player/video/|assets\\.indavideo\\.hu/swf/player\\.swf\\?.*\\b(?:v(?:ID|id))=)(?P<id>[\\da-f]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): IndavideoEmbedIE__VALID_URL,
			})
		}())
	})
}
