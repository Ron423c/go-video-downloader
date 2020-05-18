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
 * lenta/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/lenta.py
 */

package lenta

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	LentaIE       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		LentaIE = λ.Cal(λ.TypeType, λ.StrLiteral("LentaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LentaIE__VALID_URL    λ.Object
				LentaIE__real_extract λ.Object
			)
			LentaIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?lenta\\.ru/[^/]+/\\d+/\\d+/\\d+/(?P<id>[^/?#&]+)")
			LentaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
					)
					ϒdisplay_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒdisplay_id)
					ϒvideo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("vid\\s*:\\s*[\"\\']?(\\d+)"),
						ϒwebpage,
						λ.StrLiteral("eagleplatform id"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvideo_id) {
						return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Mod(λ.StrLiteral("eagleplatform:lentaru.media.eagleplatform.com:%s"), ϒvideo_id)), λ.KWArgs{
							{Name: "ie", Value: λ.StrLiteral("EaglePlatform")},
							{Name: "video_id", Value: ϒvideo_id},
						})
					}
					return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒurl), λ.KWArgs{
						{Name: "ie", Value: λ.StrLiteral("Generic")},
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    LentaIE__VALID_URL,
				"_real_extract": LentaIE__real_extract,
			})
		}())
	})
}
