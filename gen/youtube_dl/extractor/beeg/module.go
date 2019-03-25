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
 * beeg/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/beeg.py
 */

package beeg

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BeegIE             λ.Object
	InfoExtractor      λ.Object
	ϒcompat_str        λ.Object
	ϒint_or_none       λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		BeegIE = λ.Cal(λ.TypeType, λ.NewStr("BeegIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BeegIE__TESTS        λ.Object
				BeegIE__VALID_URL    λ.Object
				BeegIE__real_extract λ.Object
			)
			BeegIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?beeg\\.(?:com|porn(?:/video)?)/(?P<id>\\d+)")
			BeegIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://beeg.com/5416503"),
					λ.NewStr("md5"): λ.NewStr("a1a1b1a8bc70a89e49ccfd113aed0820"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("5416503"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Sultry Striptease"),
						λ.NewStr("description"): λ.NewStr("md5:d22219c09da287c14bed3d6c37ce4bc2"),
						λ.NewStr("timestamp"):   λ.NewInt(1391813355),
						λ.NewStr("upload_date"): λ.NewStr("20140207"),
						λ.NewStr("duration"):    λ.NewInt(383),
						λ.NewStr("tags"):        λ.ListType,
						λ.NewStr("age_limit"):   λ.NewInt(18),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://beeg.porn/video/5416503"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://beeg.porn/5416503"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			BeegIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_path     λ.Object
						ϒbeeg_version λ.Object
						ϒdescription  λ.Object
						ϒdisplay_id   λ.Object
						ϒduration     λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒheight       λ.Object
						ϒself         = λargs[0]
						ϒseries       λ.Object
						ϒtags         λ.Object
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo        λ.Object
						ϒvideo_id     λ.Object
						ϒvideo_url    λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒbeeg_version = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("beeg_version\\s*=\\s*([\\da-zA-Z_-]+)"),
						ϒwebpage,
						λ.NewStr("beeg version"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("1546225636701")},
					})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.NewStr(""),
						λ.NewStr("api."),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒapi_path = τmp1
						ϒvideo = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.NewStr("https://%sbeeg.com/api/v6/%s/video/%s"), λ.NewTuple(
								ϒapi_path,
								ϒbeeg_version,
								ϒvideo_id,
							)),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "fatal", Value: λ.Eq(ϒapi_path, λ.NewStr("api."))},
						})
						if λ.IsTrue(ϒvideo) {
							break
						}
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo, "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒvideo_url = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
							continue
						}
						ϒheight = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("^(\\d+)[pP]$"),
							ϒformat_id,
							λ.NewStr("height"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒheight))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.Cal(λ.GetAttr(ϒvideo_url, "replace", nil), λ.NewStr("{DATA_MARKERS}"), λ.Mod(λ.NewStr("data=pc_XX__%s_0"), ϒbeeg_version)), λ.NewStr("https:")),
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("height"):    λ.Cal(λ.IntType, ϒheight),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("title"))
					ϒvideo_id = λ.Cal(ϒcompat_str, func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}())
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("code"))
					ϒdescription = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("desc"))
					ϒseries = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("ps_name"))
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("date")))
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("duration")))
					ϒtags = func() λ.Object {
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("tags"))) {
							return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒtag λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo, λ.NewStr("tags")), "split", nil), λ.NewStr(",")))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒtag = τmp1
											λgen.Yield(λ.Cal(λ.GetAttr(ϒtag, "strip", nil)))
										}
										return λ.None
									})
								})))
						} else {
							return λ.None
						}
					}()
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("series"):      ϒseries,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("tags"):        ϒtags,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("age_limit"):   λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        BeegIE__TESTS,
				λ.NewStr("_VALID_URL"):    BeegIE__VALID_URL,
				λ.NewStr("_real_extract"): BeegIE__real_extract,
			})
		}())
	})
}
