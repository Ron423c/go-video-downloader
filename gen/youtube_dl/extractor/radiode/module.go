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
 * radiode/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/radiode.py
 */

package radiode

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	RadioDeIE     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		RadioDeIE = λ.Cal(λ.TypeType, λ.NewStr("RadioDeIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RadioDeIE_IE_NAME       λ.Object
				RadioDeIE__TEST         λ.Object
				RadioDeIE__VALID_URL    λ.Object
				RadioDeIE__real_extract λ.Object
			)
			RadioDeIE_IE_NAME = λ.NewStr("radio.de")
			RadioDeIE__VALID_URL = λ.NewStr("https?://(?P<id>.+?)\\.(?:radio\\.(?:de|at|fr|pt|es|pl|it)|rad\\.io)")
			RadioDeIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://ndr2.radio.de/"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("ndr2"),
					λ.NewStr("ext"):         λ.NewStr("mp3"),
					λ.NewStr("title"):       λ.NewStr("re:^NDR 2 [0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}$"),
					λ.NewStr("description"): λ.NewStr("md5:591c49c702db1a33751625ebfb67f273"),
					λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.png"),
					λ.NewStr("is_live"):     λ.True,
				}),
				λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("skip_download"): λ.True,
				}),
			})
			RadioDeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbroadcast   λ.Object
						ϒdescription λ.Object
						ϒformats     λ.Object
						ϒjscode      λ.Object
						ϒradio_id    λ.Object
						ϒself        = λargs[0]
						ϒthumbnail   λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒwebpage     λ.Object
					)
					ϒradio_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒradio_id)
					ϒjscode = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("'components/station/stationService':\\s*\\{\\s*'?station'?:\\s*(\\{.*?\\s*\\}),\\n"), ϒwebpage, λ.NewStr("broadcast"))
					ϒbroadcast = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), ϒjscode, ϒradio_id)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_live_title", nil), λ.GetItem(ϒbroadcast, λ.NewStr("name")))
					ϒdescription = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒbroadcast, "get", nil), λ.NewStr("description")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒbroadcast, "get", nil), λ.NewStr("shortDescription"))
						}
					}()
					ϒthumbnail = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒbroadcast, "get", nil), λ.NewStr("picture4Url")); λ.IsTrue(λv) {
							return λv
						} else if λv := λ.Cal(λ.GetAttr(ϒbroadcast, "get", nil), λ.NewStr("picture4TransUrl")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒbroadcast, "get", nil), λ.NewStr("logo100x100"))
						}
					}()
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒstream λ.Object
									τmp0    λ.Object
									τmp1    λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒbroadcast, λ.NewStr("streamUrls")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒstream = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):    λ.GetItem(ϒstream, λ.NewStr("streamUrl")),
										λ.NewStr("ext"):    λ.Cal(λ.GetAttr(λ.GetItem(ϒstream, λ.NewStr("streamContentFormat")), "lower", nil)),
										λ.NewStr("acodec"): λ.GetItem(ϒstream, λ.NewStr("streamContentFormat")),
										λ.NewStr("abr"):    λ.GetItem(ϒstream, λ.NewStr("bitRate")),
										λ.NewStr("asr"):    λ.GetItem(ϒstream, λ.NewStr("sampleRate")),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒradio_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("is_live"):     λ.True,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       RadioDeIE_IE_NAME,
				λ.NewStr("_TEST"):         RadioDeIE__TEST,
				λ.NewStr("_VALID_URL"):    RadioDeIE__VALID_URL,
				λ.NewStr("_real_extract"): RadioDeIE__real_extract,
			})
		}())
	})
}
