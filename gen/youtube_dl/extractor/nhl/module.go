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
 * nhl/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/nhl.py
 */

package nhl

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	NHLBaseIE       λ.Object
	NHLIE           λ.Object
	ϒcompat_str     λ.Object
	ϒdetermine_ext  λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒparse_iso8601  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒparse_duration = Ωutils.ϒparse_duration
		NHLBaseIE = λ.Cal(λ.TypeType, λ.NewStr("NHLBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NHLBaseIE__real_extract λ.Object
			)
			NHLBaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcuts           λ.Object
						ϒext            λ.Object
						ϒformats        λ.Object
						ϒheight         λ.Object
						ϒm3u8_formats   λ.Object
						ϒplayback       λ.Object
						ϒplayback_url   λ.Object
						ϒself           = λargs[0]
						ϒsite           λ.Object
						ϒthumbnail_data λ.Object
						ϒthumbnail_url  λ.Object
						ϒthumbnails     λ.Object
						ϒtitle          λ.Object
						ϒtmp_id         λ.Object
						ϒurl            = λargs[1]
						ϒvideo          λ.Object
						ϒvideo_data     λ.Object
						ϒvideo_id       λ.Object
						ϒvideos         λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒsite = λ.GetItem(τmp0, λ.NewInt(0))
					ϒtmp_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://%s/%s/%sid/v1/%s/details/web-v1.json"), λ.NewTuple(
						λ.GetAttr(ϒself, "_CONTENT_DOMAIN", nil),
						λ.GetItem(ϒsite, λ.NewSlice(λ.None, λ.NewInt(3), λ.None)),
						func() λ.Object {
							if λ.IsTrue(λ.Eq(ϒsite, λ.NewStr("mlb"))) {
								return λ.NewStr("item/")
							} else {
								return λ.NewStr("")
							}
						}(),
						ϒtmp_id,
					)), ϒtmp_id)
					if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("type")), λ.NewStr("video"))) {
						ϒvideo_data = λ.GetItem(ϒvideo_data, λ.NewStr("media"))
						ϒvideo = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("video"))
						if λ.IsTrue(ϒvideo) {
							ϒvideo_data = ϒvideo
						} else {
							ϒvideos = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("videos"))
							if λ.IsTrue(ϒvideos) {
								ϒvideo_data = λ.GetItem(ϒvideos, λ.NewInt(0))
							}
						}
					}
					ϒvideo_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒvideo_data, λ.NewStr("id")))
					ϒtitle = λ.GetItem(ϒvideo_data, λ.NewStr("title"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("playbacks"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒplayback = τmp1
						ϒplayback_url = λ.Cal(λ.GetAttr(ϒplayback, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒplayback_url))) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒplayback_url)
						if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
							ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒplayback_url,
								ϒvideo_id,
								λ.NewStr("mp4"),
								λ.NewStr("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.Cal(λ.GetAttr(ϒplayback, "get", nil), λ.NewStr("name"), λ.NewStr("hls"))},
								{Name: "fatal", Value: λ.False},
							})
							λ.Cal(λ.GetAttr(ϒself, "_check_formats", nil), ϒm3u8_formats, ϒvideo_id)
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), ϒm3u8_formats)
						} else {
							ϒheight = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒplayback, "get", nil), λ.NewStr("height")))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): λ.Cal(λ.GetAttr(ϒplayback, "get", nil), λ.NewStr("name"), λ.Add(λ.NewStr("http"), func() λ.Object {
									if λ.IsTrue(ϒheight) {
										return λ.Mod(λ.NewStr("-%dp"), ϒheight)
									} else {
										return λ.NewStr("")
									}
								}())),
								λ.NewStr("url"):    ϒplayback_url,
								λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒplayback, "get", nil), λ.NewStr("width"))),
								λ.NewStr("height"): ϒheight,
								λ.NewStr("tbr"): λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.NewStr("_(\\d+)[kK]"),
									ϒplayback_url,
									λ.NewStr("bitrate"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})),
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnails = λ.NewList()
					ϒcuts = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("image"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("cuts")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}()
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒcuts, λ.DictType)) {
						ϒcuts = λ.Cal(λ.GetAttr(ϒcuts, "values", nil))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, ϒcuts)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumbnail_data = τmp1
						ϒthumbnail_url = λ.Cal(λ.GetAttr(ϒthumbnail_data, "get", nil), λ.NewStr("src"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒthumbnail_url))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):    ϒthumbnail_url,
							λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒthumbnail_data, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒthumbnail_data, "get", nil), λ.NewStr("height"))),
						}))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("description")),
						λ.NewStr("timestamp"):   λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("date"))),
						λ.NewStr("duration"):    λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("thumbnails"):  ϒthumbnails,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_real_extract"): NHLBaseIE__real_extract,
			})
		}())
		NHLIE = λ.Cal(λ.TypeType, λ.NewStr("NHLIE"), λ.NewTuple(NHLBaseIE), func() λ.Dict {
			var (
				NHLIE_IE_NAME         λ.Object
				NHLIE__CONTENT_DOMAIN λ.Object
				NHLIE__TESTS          λ.Object
				NHLIE__VALID_URL      λ.Object
			)
			NHLIE_IE_NAME = λ.NewStr("nhl.com")
			NHLIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?(?P<site>nhl|wch2016)\\.com/(?:[^/]+/)*c-(?P<id>\\d+)")
			NHLIE__CONTENT_DOMAIN = λ.NewStr("nhl.bamcontent.com")
			NHLIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.nhl.com/video/anisimov-cleans-up-mess/t-277752844/c-43663503"),
					λ.NewStr("md5"): λ.NewStr("0f7b9a8f986fb4b4eeeece9a56416eaf"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("43663503"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Anisimov cleans up mess"),
						λ.NewStr("description"): λ.NewStr("md5:a02354acdfe900e940ce40706939ca63"),
						λ.NewStr("timestamp"):   λ.NewInt(1461288600),
						λ.NewStr("upload_date"): λ.NewStr("20160422"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.nhl.com/news/dennis-wideman-suspended/c-278258934"),
					λ.NewStr("md5"): λ.NewStr("1f39f4ea74c1394dea110699a25b366c"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("40784403"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Wideman suspended by NHL"),
						λ.NewStr("description"): λ.NewStr("Flames defenseman Dennis Wideman was banned 20 games for violation of Rule 40 (Physical Abuse of Officials)"),
						λ.NewStr("upload_date"): λ.NewStr("20160204"),
						λ.NewStr("timestamp"):   λ.NewInt(1454544904),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.nhl.com/predators/video/poile-laviolette-on-subban-trade/t-277437416/c-44315003"),
					λ.NewStr("md5"): λ.NewStr("50b2bb47f405121484dda3ccbea25459"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("44315003"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Poile, Laviolette on Subban trade"),
						λ.NewStr("description"): λ.NewStr("General manager David Poile and head coach Peter Laviolette share their thoughts on acquiring P.K. Subban from Montreal (06/29/16)"),
						λ.NewStr("timestamp"):   λ.NewInt(1467242866),
						λ.NewStr("upload_date"): λ.NewStr("20160629"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.wch2016.com/video/caneur-best-of-game-2-micd-up/t-281230378/c-44983703"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.wch2016.com/news/3-stars-team-europe-vs-team-canada/c-282195068"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):         NHLIE_IE_NAME,
				λ.NewStr("_CONTENT_DOMAIN"): NHLIE__CONTENT_DOMAIN,
				λ.NewStr("_TESTS"):          NHLIE__TESTS,
				λ.NewStr("_VALID_URL"):      NHLIE__VALID_URL,
			})
		}())
	})
}
