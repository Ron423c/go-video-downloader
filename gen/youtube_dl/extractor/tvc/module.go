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
 * tvc/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/tvc.py
 */

package tvc

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	TVCArticleIE  λ.Object
	TVCIE         λ.Object
	ϒclean_html   λ.Object
	ϒint_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		TVCIE = λ.Cal(λ.TypeType, λ.NewStr("TVCIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVCIE__VALID_URL    λ.Object
				TVCIE__real_extract λ.Object
			)
			TVCIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tvc\\.ru/video/iframe/id/(?P<id>\\d+)")
			TVCIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformat_id λ.Object
						ϒformats   λ.Object
						ϒinfo      λ.Object
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvideo     λ.Object
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						τmp0       λ.Object
						τmp1       λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒvideo = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://www.tvc.ru/video/json/id/%s"), ϒvideo_id), ϒvideo_id)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("path"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("quality"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒinfo = τmp1
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
							continue
						}
						ϒformat_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("cdnvideo/([^/]+?)(?:-[^/]+?)?/"),
							ϒvideo_url,
							λ.NewStr("format id"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒvideo_url,
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("height"))),
							λ.NewStr("tbr"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("bitrate"))),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     λ.GetItem(ϒvideo, λ.NewStr("title")),
						λ.NewStr("thumbnail"): λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("picture")),
						λ.NewStr("duration"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    TVCIE__VALID_URL,
				λ.NewStr("_real_extract"): TVCIE__real_extract,
			})
		}())
		TVCArticleIE = λ.Cal(λ.TypeType, λ.NewStr("TVCArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVCArticleIE__VALID_URL    λ.Object
				TVCArticleIE__real_extract λ.Object
			)
			TVCArticleIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tvc\\.ru/(?!video/iframe/id/)(?P<id>[^?#]+)")
			TVCArticleIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself    = λargs[0]
						ϒurl     = λargs[1]
						ϒwebpage λ.Object
					)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):       λ.NewStr("url_transparent"),
						λ.NewStr("ie_key"):      λ.NewStr("TVC"),
						λ.NewStr("url"):         λ.Cal(λ.GetAttr(ϒself, "_og_search_video_url", nil), ϒwebpage),
						λ.NewStr("title"):       λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)),
						λ.NewStr("description"): λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    TVCArticleIE__VALID_URL,
				λ.NewStr("_real_extract"): TVCArticleIE__real_extract,
			})
		}())
	})
}
