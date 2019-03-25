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
 * bpb/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bpb.py
 */

package bpb

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BpbIE          λ.Object
	InfoExtractor  λ.Object
	ϒdetermine_ext λ.Object
	ϒjs_to_json    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		BpbIE = λ.Cal(λ.TypeType, λ.NewStr("BpbIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BpbIE__TEST         λ.Object
				BpbIE__VALID_URL    λ.Object
				BpbIE__real_extract λ.Object
			)
			BpbIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?bpb\\.de/mediathek/(?P<id>[0-9]+)/")
			BpbIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.bpb.de/mediathek/297/joachim-gauck-zu-1989-und-die-erinnerung-an-die-ddr"),
				λ.NewStr("md5"): λ.NewStr("c4f84c8a8044ca9ff68bb8441d300b3f"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("297"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Joachim Gauck zu 1989 und die Erinnerung an die DDR"),
					λ.NewStr("description"): λ.NewStr("Joachim Gauck, erster Beauftragter für die Stasi-Unterlagen, spricht auf dem Geschichtsforum über die friedliche Revolution 1989 und eine \"gewisse Traurigkeit\" im Umgang mit der DDR-Vergangenheit."),
				}),
			})
			BpbIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒformats          λ.Object
						ϒquality          λ.Object
						ϒself             = λargs[0]
						ϒtitle            λ.Object
						ϒurl              = λargs[1]
						ϒvideo_id         λ.Object
						ϒvideo_info       λ.Object
						ϒvideo_info_dicts λ.Object
						ϒvideo_url        λ.Object
						ϒwebpage          λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<h2 class=\"white\">(.*?)</h2>"), ϒwebpage, λ.NewStr("title"))
					ϒvideo_info_dicts = λ.Cal(Ωre.ϒfindall, λ.NewStr("({\\s*src\\s*:\\s*'https?://film\\.bpb\\.de/[^}]+})"), ϒwebpage)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_info_dicts)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_info = τmp1
						ϒvideo_info = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							ϒvideo_info,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒjs_to_json},
							{Name: "fatal", Value: λ.False},
						})
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_info))) {
							continue
						}
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("src"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
							continue
						}
						ϒquality = func() λ.Object {
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒvideo_url, λ.NewStr("_high")))) {
								return λ.NewStr("high")
							} else {
								return λ.NewStr("low")
							}
						}()
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒvideo_url,
							λ.NewStr("preference"): func() λ.Object {
								if λ.IsTrue(λ.Eq(ϒquality, λ.NewStr("high"))) {
									return λ.NewInt(10)
								} else {
									return λ.NewInt(0)
								}
							}(),
							λ.NewStr("format_note"): ϒquality,
							λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
								ϒquality,
								λ.Cal(ϒdetermine_ext, ϒvideo_url),
							)),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         BpbIE__TEST,
				λ.NewStr("_VALID_URL"):    BpbIE__VALID_URL,
				λ.NewStr("_real_extract"): BpbIE__real_extract,
			})
		}())
	})
}
