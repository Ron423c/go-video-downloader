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
 * radiojavan/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/radiojavan.py
 */

package radiojavan

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	RadioJavanIE        λ.Object
	ϒparse_resolution   λ.Object
	ϒstr_to_int         λ.Object
	ϒunified_strdate    λ.Object
	ϒurlencode_postdata λ.Object
	ϒurljoin            λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_resolution = Ωutils.ϒparse_resolution
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		ϒurljoin = Ωutils.ϒurljoin
		RadioJavanIE = λ.Cal(λ.TypeType, λ.NewStr("RadioJavanIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RadioJavanIE__VALID_URL    λ.Object
				RadioJavanIE__real_extract λ.Object
			)
			RadioJavanIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?radiojavan\\.com/videos/video/(?P<id>[^/]+)/?")
			RadioJavanIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdislike_count λ.Object
						ϒdownload_host λ.Object
						ϒf             λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒlike_count    λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒvideo_path    λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒdownload_host = λ.Cal(λ.GetAttr(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("https://www.radiojavan.com/videos/video_host"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"): ϒvideo_id,
						}))},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Content-Type"): λ.NewStr("application/x-www-form-urlencoded"),
							λ.NewStr("Referer"):      ϒurl,
						})},
					}), "get", nil), λ.NewStr("host"), λ.NewStr("https://host1.rjmusicmedia.com"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("RJ\\.video(?P<format_id>\\d+[pPkK])\\s*=\\s*([\"\\'])(?P<url>(?:(?!\\2).)+)\\2"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						_ = λ.GetItem(τmp2, λ.NewInt(1))
						ϒvideo_path = λ.GetItem(τmp2, λ.NewInt(2))
						ϒf = λ.Cal(ϒparse_resolution, ϒformat_id)
						λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       λ.Cal(ϒurljoin, ϒdownload_host, ϒvideo_path),
							λ.NewStr("format_id"): ϒformat_id,
						}))
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("class=\"date_added\">Date added: ([^<]+)<"),
						ϒwebpage,
						λ.NewStr("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("class=\"views\">Plays: ([\\d,]+)"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒlike_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("class=\"rating\">([\\d,]+) likes"),
						ϒwebpage,
						λ.NewStr("like count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒdislike_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("class=\"rating\">([\\d,]+) dislikes"),
						ϒwebpage,
						λ.NewStr("dislike count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("upload_date"):   ϒupload_date,
						λ.NewStr("view_count"):    ϒview_count,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("dislike_count"): ϒdislike_count,
						λ.NewStr("formats"):       ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    RadioJavanIE__VALID_URL,
				λ.NewStr("_real_extract"): RadioJavanIE__real_extract,
			})
		}())
	})
}
