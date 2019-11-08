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
 * unistra/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/unistra.py
 */

package unistra

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	UnistraIE     λ.Object
	ϒqualities    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒqualities = Ωutils.ϒqualities
		UnistraIE = λ.Cal(λ.TypeType, λ.NewStr("UnistraIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UnistraIE__VALID_URL    λ.Object
				UnistraIE__real_extract λ.Object
			)
			UnistraIE__VALID_URL = λ.NewStr("https?://utv\\.unistra\\.fr/(?:index|video)\\.php\\?id_video\\=(?P<id>\\d+)")
			UnistraIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒfile_path   λ.Object
						ϒfiles       λ.Object
						ϒformat_id   λ.Object
						ϒformats     λ.Object
						ϒmobj        λ.Object
						ϒquality     λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒfiles = λ.Cal(λ.SetType, λ.Cal(Ωre.ϒfindall, λ.NewStr("file\\s*:\\s*\"(/[^\"]+)\""), ϒwebpage))
					ϒquality = λ.Cal(ϒqualities, λ.NewList(
						λ.NewStr("SD"),
						λ.NewStr("HD"),
					))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒfiles)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒfile_path = τmp1
						ϒformat_id = func() λ.Object {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒfile_path, "endswith", nil), λ.NewStr("-HD.mp4"))) {
								return λ.NewStr("HD")
							} else {
								return λ.NewStr("SD")
							}
						}()
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       λ.Mod(λ.NewStr("http://vod-flash.u-strasbg.fr:8080%s"), ϒfile_path),
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("quality"):   λ.Cal(ϒquality, ϒformat_id),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<title>UTV - (.*?)</"), ϒwebpage, λ.NewStr("title"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<meta name=\"Description\" content=\"(.*?)\""),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "flags", Value: Ωre.DOTALL},
					})
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("image: \"(.*?)\""), ϒwebpage, λ.NewStr("thumbnail"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    UnistraIE__VALID_URL,
				λ.NewStr("_real_extract"): UnistraIE__real_extract,
			})
		}())
	})
}
