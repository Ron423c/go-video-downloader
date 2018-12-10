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
 * canalc2/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/canalc2.py
 */

package canalc2

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	Canalc2IE       λ.Object
	InfoExtractor   λ.Object
	ϒparse_duration λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_duration = Ωutils.ϒparse_duration
		Canalc2IE = λ.Cal(λ.TypeType, λ.NewStr("Canalc2IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Canalc2IE_IE_NAME       λ.Object
				Canalc2IE__TESTS        λ.Object
				Canalc2IE__VALID_URL    λ.Object
				Canalc2IE__real_extract λ.Object
			)
			Canalc2IE_IE_NAME = λ.NewStr("canalc2.tv")
			Canalc2IE__VALID_URL = λ.NewStr("https?://(?:(?:www\\.)?canalc2\\.tv/video/|archives-canalc2\\.u-strasbg\\.fr/video\\.asp\\?.*\\bidVideo=)(?P<id>\\d+)")
			Canalc2IE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.canalc2.tv/video/12163"),
					λ.NewStr("md5"): λ.NewStr("060158428b650f896c542dfbb3d6487f"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):       λ.NewStr("12163"),
						λ.NewStr("ext"):      λ.NewStr("mp4"),
						λ.NewStr("title"):    λ.NewStr("Terrasses du Numérique"),
						λ.NewStr("duration"): λ.NewInt(122),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://archives-canalc2.u-strasbg.fr/video.asp?idVideo=11427&voir=oui"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			Canalc2IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats   λ.Object
						ϒinfo      λ.Object
						ϒrtmp      λ.Object
						ϒself      = λargs[0]
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						ϒwebpage   λ.Object
						τmp0       λ.Object
						τmp1       λ.Object
						τmp2       λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://www.canalc2.tv/video/%s"), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("(?s)class=\"[^\"]*col_description[^\"]*\">.*?<h3>(.+?)</h3>"), ϒwebpage, λ.NewStr("title"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("file\\s*=\\s*([\"\\'])(.+?)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						_ = λ.GetItem(τmp2, λ.NewInt(0))
						ϒvideo_url = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo_url, "startswith", nil), λ.NewStr("rtmp://"))) {
							ϒrtmp = λ.Cal(Ωre.ϒsearch, λ.NewStr("^(?P<url>rtmp://[^/]+/(?P<app>.+/))(?P<play_path>mp4:.+)$"), ϒvideo_url)
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       λ.Cal(λ.GetAttr(ϒrtmp, "group", nil), λ.NewStr("url")),
								λ.NewStr("format_id"): λ.NewStr("rtmp"),
								λ.NewStr("ext"):       λ.NewStr("flv"),
								λ.NewStr("app"):       λ.Cal(λ.GetAttr(ϒrtmp, "group", nil), λ.NewStr("app")),
								λ.NewStr("play_path"): λ.Cal(λ.GetAttr(ϒrtmp, "group", nil), λ.NewStr("play_path")),
								λ.NewStr("page_url"):  ϒurl,
							}))
						} else {
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒvideo_url,
								λ.NewStr("format_id"): λ.NewStr("http"),
							}))
						}
					}
					if λ.IsTrue(ϒformats) {
						ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("formats"): ϒformats,
						})
					} else {
						ϒinfo = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_html5_media_entries", nil), ϒurl, ϒwebpage, ϒurl), λ.NewInt(0))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), λ.GetItem(ϒinfo, λ.NewStr("formats")))
					λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("title"): ϒtitle,
						λ.NewStr("duration"): λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("id=[\"\\']video_duree[\"\\'][^>]*>([^<]+)"),
							ϒwebpage,
							λ.NewStr("duration"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.False},
						})),
					}))
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       Canalc2IE_IE_NAME,
				λ.NewStr("_TESTS"):        Canalc2IE__TESTS,
				λ.NewStr("_VALID_URL"):    Canalc2IE__VALID_URL,
				λ.NewStr("_real_extract"): Canalc2IE__real_extract,
			})
		}())
	})
}
