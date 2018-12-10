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
 * vevo/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/vevo.py
 */

package vevo

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError    λ.Object
	InfoExtractor     λ.Object
	VevoBaseIE        λ.Object
	VevoIE            λ.Object
	VevoPlaylistIE    λ.Object
	ϒcompat_HTTPError λ.Object
	ϒcompat_str       λ.Object
	ϒint_or_none      λ.Object
	ϒparse_iso8601    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		VevoBaseIE = λ.Cal(λ.TypeType, λ.NewStr("VevoBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VevoBaseIE__extract_json λ.Object
			)
			VevoBaseIE__extract_json = λ.NewFunction("_extract_json",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
						ϒwebpage  = λargs[1]
					)
					return λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("window\\.__INITIAL_STORE__\\s*=\\s*({.+?});\\s*</script>"), ϒwebpage, λ.NewStr("initial store")), ϒvideo_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_extract_json"): VevoBaseIE__extract_json,
			})
		}())
		VevoIE = λ.Cal(λ.TypeType, λ.NewStr("VevoIE"), λ.NewTuple(VevoBaseIE), func() λ.Dict {
			var (
				VevoIE__TESTS          λ.Object
				VevoIE__VALID_URL      λ.Object
				VevoIE__VERSIONS       λ.Object
				VevoIE__call_api       λ.Object
				VevoIE__initialize_api λ.Object
				VevoIE__real_extract   λ.Object
			)
			λ.NewStr("\n    Accepts urls from vevo.com or in the format 'vevo:{id}'\n    (currently used by MTVIE and MySpaceIE)\n    ")
			VevoIE__VALID_URL = λ.NewStr("(?x)\n        (?:https?://(?:www\\.)?vevo\\.com/watch/(?!playlist|genre)(?:[^/]+/(?:[^/]+/)?)?|\n           https?://cache\\.vevo\\.com/m/html/embed\\.html\\?video=|\n           https?://videoplayer\\.vevo\\.com/embed/embedded\\?videoId=|\n           vevo:)\n        (?P<id>[^&?#]+)")
			VevoIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.vevo.com/watch/hurts/somebody-to-die-for/GB1101300280"),
					λ.NewStr("md5"): λ.NewStr("95ee28ee45e70130e3ab02b0f579ae23"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("GB1101300280"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Hurts - Somebody to Die For"),
						λ.NewStr("timestamp"):   λ.NewInt(1372057200),
						λ.NewStr("upload_date"): λ.NewStr("20130624"),
						λ.NewStr("uploader"):    λ.NewStr("Hurts"),
						λ.NewStr("track"):       λ.NewStr("Somebody to Die For"),
						λ.NewStr("artist"):      λ.NewStr("Hurts"),
						λ.NewStr("genre"):       λ.NewStr("Pop"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(
						λ.NewStr("Unable to download SMIL file"),
						λ.NewStr("Unable to download info"),
					),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("note"): λ.NewStr("v3 SMIL format"),
					λ.NewStr("url"):  λ.NewStr("http://www.vevo.com/watch/cassadee-pope/i-wish-i-could-break-your-heart/USUV71302923"),
					λ.NewStr("md5"):  λ.NewStr("f6ab09b034f8c22969020b042e5ac7fc"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("USUV71302923"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Cassadee Pope - I Wish I Could Break Your Heart"),
						λ.NewStr("timestamp"):   λ.NewInt(1392796919),
						λ.NewStr("upload_date"): λ.NewStr("20140219"),
						λ.NewStr("uploader"):    λ.NewStr("Cassadee Pope"),
						λ.NewStr("track"):       λ.NewStr("I Wish I Could Break Your Heart"),
						λ.NewStr("artist"):      λ.NewStr("Cassadee Pope"),
						λ.NewStr("genre"):       λ.NewStr("Country"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(
						λ.NewStr("Unable to download SMIL file"),
						λ.NewStr("Unable to download info"),
					),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("note"): λ.NewStr("Age-limited video"),
					λ.NewStr("url"):  λ.NewStr("https://www.vevo.com/watch/justin-timberlake/tunnel-vision-explicit/USRV81300282"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("USRV81300282"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Justin Timberlake - Tunnel Vision (Explicit)"),
						λ.NewStr("age_limit"):   λ.NewInt(18),
						λ.NewStr("timestamp"):   λ.NewInt(1372888800),
						λ.NewStr("upload_date"): λ.NewStr("20130703"),
						λ.NewStr("uploader"):    λ.NewStr("Justin Timberlake"),
						λ.NewStr("track"):       λ.NewStr("Tunnel Vision (Explicit)"),
						λ.NewStr("artist"):      λ.NewStr("Justin Timberlake"),
						λ.NewStr("genre"):       λ.NewStr("Pop"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(
						λ.NewStr("Unable to download SMIL file"),
						λ.NewStr("Unable to download info"),
					),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("note"): λ.NewStr("No video_info"),
					λ.NewStr("url"):  λ.NewStr("http://www.vevo.com/watch/k-camp-1/Till-I-Die/USUV71503000"),
					λ.NewStr("md5"):  λ.NewStr("8b83cc492d72fc9cf74a02acee7dc1b0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("USUV71503000"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("K Camp ft. T.I. - Till I Die"),
						λ.NewStr("age_limit"):   λ.NewInt(18),
						λ.NewStr("timestamp"):   λ.NewInt(1449468000),
						λ.NewStr("upload_date"): λ.NewStr("20151207"),
						λ.NewStr("uploader"):    λ.NewStr("K Camp"),
						λ.NewStr("track"):       λ.NewStr("Till I Die"),
						λ.NewStr("artist"):      λ.NewStr("K Camp"),
						λ.NewStr("genre"):       λ.NewStr("Hip-Hop"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(
						λ.NewStr("Unable to download SMIL file"),
						λ.NewStr("Unable to download info"),
					),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("note"): λ.NewStr("Featured test"),
					λ.NewStr("url"):  λ.NewStr("https://www.vevo.com/watch/lemaitre/Wait/USUV71402190"),
					λ.NewStr("md5"):  λ.NewStr("d28675e5e8805035d949dc5cf161071d"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("USUV71402190"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Lemaitre ft. LoLo - Wait"),
						λ.NewStr("age_limit"):   λ.NewInt(0),
						λ.NewStr("timestamp"):   λ.NewInt(1413432000),
						λ.NewStr("upload_date"): λ.NewStr("20141016"),
						λ.NewStr("uploader"):    λ.NewStr("Lemaitre"),
						λ.NewStr("track"):       λ.NewStr("Wait"),
						λ.NewStr("artist"):      λ.NewStr("Lemaitre"),
						λ.NewStr("genre"):       λ.NewStr("Electronic"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(
						λ.NewStr("Unable to download SMIL file"),
						λ.NewStr("Unable to download info"),
					),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("note"): λ.NewStr("Only available via webpage"),
					λ.NewStr("url"):  λ.NewStr("http://www.vevo.com/watch/GBUV71600656"),
					λ.NewStr("md5"):  λ.NewStr("67e79210613865b66a47c33baa5e37fe"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("GBUV71600656"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("ABC - Viva Love"),
						λ.NewStr("age_limit"):   λ.NewInt(0),
						λ.NewStr("timestamp"):   λ.NewInt(1461830400),
						λ.NewStr("upload_date"): λ.NewStr("20160428"),
						λ.NewStr("uploader"):    λ.NewStr("ABC"),
						λ.NewStr("track"):       λ.NewStr("Viva Love"),
						λ.NewStr("artist"):      λ.NewStr("ABC"),
						λ.NewStr("genre"):       λ.NewStr("Pop"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(λ.NewStr("Failed to download video versions info")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.vevo.com/watch/INS171400764"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.vevo.com/watch/boostee/pop-corn-clip-officiel/FR1A91600909"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			VevoIE__VERSIONS = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewInt(0): λ.NewStr("youtube"),
				λ.NewInt(1): λ.NewStr("level3"),
				λ.NewInt(2): λ.NewStr("akamai"),
				λ.NewInt(3): λ.NewStr("level3"),
				λ.NewInt(4): λ.NewStr("amazon"),
			})
			VevoIE__initialize_api = λ.NewFunction("_initialize_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒauth_info λ.Object
						ϒself      = λargs[0]
						ϒvideo_id  = λargs[1]
						ϒwebpage   λ.Object
					)
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.NewStr("https://accounts.vevo.com/token"),
						λ.None,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Retrieving oauth token")},
						{Name: "errnote", Value: λ.NewStr("Unable to retrieve oauth token")},
						{Name: "data", Value: λ.Cal(λ.GetAttr(λ.Cal(Ωjson.ϒdumps, λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("client_id"):  λ.NewStr("SPupX1tvqFEopQ1YS6SS"),
							λ.NewStr("grant_type"): λ.NewStr("urn:vevo:params:oauth:grant-type:anonymous"),
						})), "encode", nil), λ.NewStr("utf-8"))},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Content-Type"): λ.NewStr("application/json"),
						})},
					})
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.NewStr("(?i)THIS PAGE IS CURRENTLY UNAVAILABLE IN YOUR REGION"), ϒwebpage)) {
						λ.Cal(λ.GetAttr(ϒself, "raise_geo_restricted", nil), λ.Mod(λ.NewStr("%s said: This page is currently unavailable in your region"), λ.GetAttr(ϒself, "IE_NAME", nil)))
					}
					ϒauth_info = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), ϒwebpage, ϒvideo_id)
					λ.SetAttr(ϒself, "_api_url_template", λ.Add(λ.Add(λ.Cal(λ.GetAttr(ϒself, "http_scheme", nil)), λ.NewStr("//apiv2.vevo.com/%s?token=")), λ.GetItem(ϒauth_info, λ.NewStr("legacy_token"))))
					return λ.None
				})
			VevoIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs          = λargs[2]
						ϒdata          λ.Object
						ϒerror_message λ.Object
						ϒerrors        λ.Object
						ϒkwargs        = λargs[3]
						ϒpath          = λargs[1]
						ϒself          = λargs[0]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					_ = τmp0
					_ = τmp1
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								ϒe := λex
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), ϒcompat_HTTPError)) {
									ϒerrors = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "read", nil)), "decode", nil)), λ.None), λ.NewStr("errors"))
									ϒerror_message = λ.Cal(λ.GetAttr(λ.NewStr(", "), "join", nil), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
										nil,
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
												var (
													ϒerror λ.Object
													τmp0   λ.Object
													τmp1   λ.Object
												)
												τmp0 = λ.Cal(λ.BuiltinIter, ϒerrors)
												for {
													if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
														break
													}
													ϒerror = τmp1
													λgen.Yield(λ.GetItem(ϒerror, λ.NewStr("message")))
												}
												return λ.None
											})
										}))))
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
										λ.GetAttr(ϒself, "IE_NAME", nil),
										ϒerror_message,
									))), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
								panic(λ.Raise(λex))
							}},
						)
						ϒdata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(λ.Unpack(
							λ.Mod(λ.GetAttr(ϒself, "_api_url_template", nil), ϒpath),
							λ.AsStarred(ϒargs),
						)...), λ.KWArgs{
							{Name: "", Value: ϒkwargs},
						})
						return λ.BlockExitNormally, nil
					}()
					return ϒdata
				})
			VevoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit       λ.Object
						ϒartist          λ.Object
						ϒartists         λ.Object
						ϒcurr_artist     λ.Object
						ϒfeatured_artist λ.Object
						ϒformats         λ.Object
						ϒgenre           λ.Object
						ϒgenres          λ.Object
						ϒis_explicit     λ.Object
						ϒjson_data       λ.Object
						ϒm               λ.Object
						ϒself            = λargs[0]
						ϒtitle           λ.Object
						ϒtrack           λ.Object
						ϒuploader        λ.Object
						ϒurl             = λargs[1]
						ϒversion         λ.Object
						ϒversion_url     λ.Object
						ϒvideo_id        λ.Object
						ϒvideo_info      λ.Object
						ϒvideo_version   λ.Object
						ϒvideo_versions  λ.Object
						ϒwebpage         λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					λ.Cal(λ.GetAttr(ϒself, "_initialize_api", nil), ϒvideo_id)
					ϒvideo_info = λ.Cal(λ.GetAttr(ϒself, "_call_api", nil), λ.Mod(λ.NewStr("video/%s"), ϒvideo_id), ϒvideo_id, λ.NewStr("Downloading api video info"), λ.NewStr("Failed to download video info"))
					ϒvideo_versions = λ.Call(λ.GetAttr(ϒself, "_call_api", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("video/%s/streams"), ϒvideo_id),
						ϒvideo_id,
						λ.NewStr("Downloading video versions info"),
						λ.NewStr("Failed to download video versions info"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_versions))) {
						ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
						ϒjson_data = λ.Cal(λ.GetAttr(ϒself, "_extract_json", nil), ϒwebpage, ϒvideo_id)
						if λ.IsTrue(λ.NewBool(λ.Contains(λ.Cal(λ.GetAttr(ϒjson_data, "get", nil), λ.NewStr("default"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), λ.NewStr("streams")))) {
							ϒvideo_versions = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒjson_data, λ.NewStr("default")), λ.NewStr("streams")), ϒvideo_id), λ.NewInt(0))
						} else {
							ϒvideo_versions = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒkey   λ.Object
											ϒvalue λ.Object
											τmp0   λ.Object
											τmp1   λ.Object
											τmp2   λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(λ.GetItem(ϒjson_data, λ.NewStr("apollo")), λ.NewStr("data")), "items", nil)))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											τmp2 = τmp1
											ϒkey = λ.GetItem(τmp2, λ.NewInt(0))
											ϒvalue = λ.GetItem(τmp2, λ.NewInt(1))
											if λ.IsTrue(λ.Cal(λ.GetAttr(ϒkey, "startswith", nil), λ.Mod(λ.NewStr("%s.streams"), ϒvideo_id))) {
												λgen.Yield(ϒvalue)
											}
										}
										return λ.None
									})
								})))
						}
					}
					ϒuploader = λ.None
					ϒartist = λ.None
					ϒfeatured_artist = λ.None
					ϒartists = λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("artists"))
					τmp0 = λ.Cal(λ.BuiltinIter, ϒartists)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒcurr_artist = τmp1
						if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒcurr_artist, "get", nil), λ.NewStr("role")), λ.NewStr("Featured"))) {
							ϒfeatured_artist = λ.GetItem(ϒcurr_artist, λ.NewStr("name"))
						} else {
							τmp2 = λ.GetItem(ϒcurr_artist, λ.NewStr("name"))
							ϒartist = τmp2
							ϒuploader = τmp2
						}
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_versions)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_version = τmp1
						ϒversion = λ.Cal(λ.GetAttr(λ.GetAttr(ϒself, "_VERSIONS", nil), "get", nil), λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("version")), λ.NewStr("generic"))
						ϒversion_url = λ.Cal(λ.GetAttr(ϒvideo_version, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒversion_url))) {
							continue
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒversion_url, λ.NewStr(".ism")))) {
							continue
						} else {
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒversion_url, λ.NewStr(".mpd")))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
									ϒversion_url,
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "mpd_id", Value: λ.Mod(λ.NewStr("dash-%s"), ϒversion)},
									{Name: "note", Value: λ.Mod(λ.NewStr("Downloading %s MPD information"), ϒversion)},
									{Name: "errnote", Value: λ.Mod(λ.NewStr("Failed to download %s MPD information"), ϒversion)},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒversion_url, λ.NewStr(".m3u8")))) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒversion_url,
										ϒvideo_id,
										λ.NewStr("mp4"),
										λ.NewStr("m3u8_native"),
									), λ.KWArgs{
										{Name: "m3u8_id", Value: λ.Mod(λ.NewStr("hls-%s"), ϒversion)},
										{Name: "note", Value: λ.Mod(λ.NewStr("Downloading %s m3u8 information"), ϒversion)},
										{Name: "errnote", Value: λ.Mod(λ.NewStr("Failed to download %s m3u8 information"), ϒversion)},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒm = λ.Cal(Ωre.ϒsearch, λ.NewStr("(?xi)\n                    _(?P<width>[0-9]+)x(?P<height>[0-9]+)\n                    _(?P<vcodec>[a-z0-9]+)\n                    _(?P<vbr>[0-9]+)\n                    _(?P<acodec>[a-z0-9]+)\n                    _(?P<abr>[0-9]+)\n                    \\.(?P<ext>[a-z0-9]+)"), ϒversion_url)
									if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒm))) {
										continue
									}
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): ϒversion_url,
										λ.NewStr("format_id"): λ.Mod(λ.NewStr("http-%s-%s"), λ.NewTuple(
											ϒversion,
											λ.GetItem(ϒvideo_version, λ.NewStr("quality")),
										)),
										λ.NewStr("vcodec"): λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("vcodec")),
										λ.NewStr("acodec"): λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("acodec")),
										λ.NewStr("vbr"):    λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("vbr"))),
										λ.NewStr("abr"):    λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("abr"))),
										λ.NewStr("ext"):    λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("ext")),
										λ.NewStr("width"):  λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("width"))),
										λ.NewStr("height"): λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("height"))),
									}))
								}
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtrack = λ.GetItem(ϒvideo_info, λ.NewStr("title"))
					if λ.IsTrue(ϒfeatured_artist) {
						ϒartist = λ.Mod(λ.NewStr("%s ft. %s"), λ.NewTuple(
							ϒartist,
							ϒfeatured_artist,
						))
					}
					ϒtitle = func() λ.Object {
						if λ.IsTrue(ϒartist) {
							return λ.Mod(λ.NewStr("%s - %s"), λ.NewTuple(
								ϒartist,
								ϒtrack,
							))
						} else {
							return ϒtrack
						}
					}()
					ϒgenres = λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("genres"))
					ϒgenre = func() λ.Object {
						if λ.IsTrue(func() λ.Object {
							if λv := ϒgenres; !λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Cal(λ.BuiltinIsInstance, ϒgenres, λ.ListType); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, λ.GetItem(ϒgenres, λ.NewInt(0)), ϒcompat_str)
							}
						}()) {
							return λ.GetItem(ϒgenres, λ.NewInt(0))
						} else {
							return λ.None
						}
					}()
					ϒis_explicit = λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("isExplicit"))
					if λ.IsTrue(λ.NewBool(ϒis_explicit == λ.True)) {
						ϒage_limit = λ.NewInt(18)
					} else {
						if λ.IsTrue(λ.NewBool(ϒis_explicit == λ.False)) {
							ϒage_limit = λ.NewInt(0)
						} else {
							ϒage_limit = λ.None
						}
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):      ϒvideo_id,
						λ.NewStr("title"):   ϒtitle,
						λ.NewStr("formats"): ϒformats,
						λ.NewStr("thumbnail"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("imageUrl")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("thumbnailUrl"))
							}
						}(),
						λ.NewStr("timestamp"):  λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("releaseDate"))),
						λ.NewStr("uploader"):   ϒuploader,
						λ.NewStr("duration"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("view_count"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_info, "get", nil), λ.NewStr("views"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("total"))),
						λ.NewStr("age_limit"):  ϒage_limit,
						λ.NewStr("track"):      ϒtrack,
						λ.NewStr("artist"):     ϒuploader,
						λ.NewStr("genre"):      ϒgenre,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):          VevoIE__TESTS,
				λ.NewStr("_VALID_URL"):      VevoIE__VALID_URL,
				λ.NewStr("_VERSIONS"):       VevoIE__VERSIONS,
				λ.NewStr("_call_api"):       VevoIE__call_api,
				λ.NewStr("_initialize_api"): VevoIE__initialize_api,
				λ.NewStr("_real_extract"):   VevoIE__real_extract,
			})
		}())
		VevoPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("VevoPlaylistIE"), λ.NewTuple(VevoBaseIE), func() λ.Dict {
			var (
				VevoPlaylistIE__VALID_URL λ.Object
			)
			VevoPlaylistIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?vevo\\.com/watch/(?P<kind>playlist|genre)/(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): VevoPlaylistIE__VALID_URL,
			})
		}())
	})
}
