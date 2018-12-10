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
 * viqeo/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/viqeo.py
 */

package viqeo

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	ViqeoIE       λ.Object
	ϒint_or_none  λ.Object
	ϒstr_or_none  λ.Object
	ϒurl_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ViqeoIE = λ.Cal(λ.TypeType, λ.NewStr("ViqeoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ViqeoIE__TESTS        λ.Object
				ViqeoIE__VALID_URL    λ.Object
				ViqeoIE__real_extract λ.Object
			)
			ViqeoIE__VALID_URL = λ.NewStr("(?x)\n                        (?:\n                            viqeo:|\n                            https?://cdn\\.viqeo\\.tv/embed/*\\?.*?\\bvid=|\n                            https?://api\\.viqeo\\.tv/v\\d+/data/startup?.*?\\bvideo(?:%5B%5D|\\[\\])=\n                        )\n                        (?P<id>[\\da-f]+)\n                    ")
			ViqeoIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://cdn.viqeo.tv/embed/?vid=cde96f09d25f39bee837"),
					λ.NewStr("md5"): λ.NewStr("a169dd1a6426b350dca4296226f21e76"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("cde96f09d25f39bee837"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("cde96f09d25f39bee837"),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):  λ.NewInt(76),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("viqeo:cde96f09d25f39bee837"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://api.viqeo.tv/v1/data/startup?video%5B%5D=71bbec412ade45c3216c&profile=112"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			ViqeoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata       λ.Object
						ϒduration   λ.Object
						ϒf          λ.Object
						ϒformat_id  λ.Object
						ϒformats    λ.Object
						ϒis_audio   λ.Object
						ϒmedia_file λ.Object
						ϒmedia_kind λ.Object
						ϒmedia_type λ.Object
						ϒmedia_url  λ.Object
						ϒself       = λargs[0]
						ϒthumbnails λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("https://cdn.viqeo.tv/embed/?vid=%s"), ϒvideo_id), ϒvideo_id)
					ϒdata = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("SLOT_DATA\\s*=\\s*({.+?})\\s*;"), ϒwebpage, λ.NewStr("slot data")), ϒvideo_id)
					ϒformats = λ.NewList()
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒdata, λ.NewStr("mediaFiles")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmedia_file = τmp1
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_file, λ.DictType)))) {
							continue
						}
						ϒmedia_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒmedia_file, "get", nil), λ.NewStr("url")))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒmedia_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒmedia_url, "startswith", nil), λ.NewTuple(
									λ.NewStr("http"),
									λ.NewStr("//"),
								))))
							}
						}()) {
							continue
						}
						ϒmedia_type = λ.Cal(ϒstr_or_none, λ.Cal(λ.GetAttr(ϒmedia_file, "get", nil), λ.NewStr("type")))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒmedia_type))) {
							continue
						}
						ϒmedia_kind = λ.Cal(λ.GetAttr(λ.GetItem(λ.Cal(λ.GetAttr(ϒmedia_type, "split", nil), λ.NewStr("/")), λ.NewInt(0)), "lower", nil))
						ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):    ϒmedia_url,
							λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmedia_file, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmedia_file, "get", nil), λ.NewStr("height"))),
						})
						ϒformat_id = λ.Cal(ϒstr_or_none, λ.Cal(λ.GetAttr(ϒmedia_file, "get", nil), λ.NewStr("quality")))
						if λ.IsTrue(λ.Eq(ϒmedia_kind, λ.NewStr("image"))) {
							λ.SetItem(ϒf, λ.NewStr("id"), ϒformat_id)
							λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), ϒf)
						} else {
							if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
								λ.NewStr("video"),
								λ.NewStr("audio"),
							), ϒmedia_kind))) {
								ϒis_audio = λ.Eq(ϒmedia_kind, λ.NewStr("audio"))
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("format_id"): func() λ.Object {
										if λ.IsTrue(ϒis_audio) {
											return λ.NewStr("audio")
										} else {
											return ϒformat_id
										}
									}(),
									λ.NewStr("fps"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmedia_file, "get", nil), λ.NewStr("fps"))),
									λ.NewStr("vcodec"): func() λ.Object {
										if λ.IsTrue(ϒis_audio) {
											return λ.NewStr("none")
										} else {
											return λ.None
										}
									}(),
								}))
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("duration")))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("title"):      ϒvideo_id,
						λ.NewStr("duration"):   ϒduration,
						λ.NewStr("thumbnails"): ϒthumbnails,
						λ.NewStr("formats"):    ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        ViqeoIE__TESTS,
				λ.NewStr("_VALID_URL"):    ViqeoIE__VALID_URL,
				λ.NewStr("_real_extract"): ViqeoIE__real_extract,
			})
		}())
	})
}
