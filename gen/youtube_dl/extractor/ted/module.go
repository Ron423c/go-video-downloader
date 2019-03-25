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
 * ted/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ted.py
 */

package ted

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	TEDIE          λ.Object
	ϒcompat_str    λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒtry_get       λ.Object
	ϒurl_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒurl_or_none = Ωutils.ϒurl_or_none
		TEDIE = λ.Cal(λ.TypeType, λ.NewStr("TEDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TEDIE_IE_NAME               λ.Object
				TEDIE__NATIVE_FORMATS       λ.Object
				TEDIE__TESTS                λ.Object
				TEDIE__VALID_URL            λ.Object
				TEDIE__extract_info         λ.Object
				TEDIE__get_subtitles        λ.Object
				TEDIE__playlist_videos_info λ.Object
				TEDIE__real_extract         λ.Object
				TEDIE__talk_info            λ.Object
			)
			TEDIE_IE_NAME = λ.NewStr("ted")
			TEDIE__VALID_URL = λ.NewStr("(?x)\n        (?P<proto>https?://)\n        (?P<type>www|embed(?:-ssl)?)(?P<urlmain>\\.ted\\.com/\n        (\n            (?P<type_playlist>playlists(?:/\\d+)?) # We have a playlist\n            |\n            ((?P<type_talk>talks)) # We have a simple talk\n            |\n            (?P<type_watch>watch)/[^/]+/[^/]+\n        )\n        (/lang/(.*?))? # The url may contain the language\n        /(?P<name>[\\w-]+) # Here goes the name and then \".html\"\n        .*)$\n        ")
			TEDIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ted.com/talks/dan_dennett_on_our_consciousness.html"),
					λ.NewStr("md5"): λ.NewStr("b0ce2b05ca215042124fbc9e3886493a"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("102"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("The illusion of consciousness"),
						λ.NewStr("description"):   λ.NewStr("Philosopher Dan Dennett makes a compelling argument that not only don't we understand our own consciousness, but that half the time our brains are actively fooling us."),
						λ.NewStr("uploader"):      λ.NewStr("Dan Dennett"),
						λ.NewStr("width"):         λ.NewInt(853),
						λ.NewStr("duration"):      λ.NewInt(1308),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("tags"):          λ.ListType,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.ted.com/talks/vishal_sikka_the_beauty_and_power_of_algorithms"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("6069"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("The beauty and power of algorithms"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.+\\.jpg"),
						λ.NewStr("description"): λ.NewStr("md5:734e352710fb00d840ab87ae31aaf688"),
						λ.NewStr("uploader"):    λ.NewStr("Vishal Sikka"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ted.com/talks/gabby_giffords_and_mark_kelly_be_passionate_be_courageous_be_your_best"),
					λ.NewStr("md5"): λ.NewStr("e6b9617c01a7970ceac8bb2c92c346c0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1972"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Be passionate. Be courageous. Be your best."),
						λ.NewStr("uploader"):    λ.NewStr("Gabby Giffords and Mark Kelly"),
						λ.NewStr("description"): λ.NewStr("md5:5174aed4d0f16021b704120360f72b92"),
						λ.NewStr("duration"):    λ.NewInt(1128),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ted.com/playlists/who_are_the_hackers"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("10"),
						λ.NewStr("title"): λ.NewStr("Who are the hackers?"),
					}),
					λ.NewStr("playlist_mincount"): λ.NewInt(6),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):    λ.NewStr("https://www.ted.com/talks/douglas_adams_parrots_the_universe_and_everything"),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Youtube")),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("_ZG8HBuDjgc"),
						λ.NewStr("ext"):         λ.NewStr("webm"),
						λ.NewStr("title"):       λ.NewStr("Douglas Adams: Parrots the Universe and Everything"),
						λ.NewStr("description"): λ.NewStr("md5:01ad1e199c49ac640cb1196c0e9016af"),
						λ.NewStr("uploader"):    λ.NewStr("University of California Television (UCTV)"),
						λ.NewStr("uploader_id"): λ.NewStr("UCtelevision"),
						λ.NewStr("upload_date"): λ.NewStr("20080522"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.ted.com/talks/tom_thum_the_orchestra_in_my_mouth"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            λ.NewStr("1792"),
						λ.NewStr("ext"):           λ.NewStr("mp4"),
						λ.NewStr("title"):         λ.NewStr("The orchestra in my mouth"),
						λ.NewStr("description"):   λ.NewStr("md5:5d1d78650e2f8dfcbb8ebee2951ac29a"),
						λ.NewStr("uploader"):      λ.NewStr("Tom Thum"),
						λ.NewStr("view_count"):    λ.IntType,
						λ.NewStr("comment_count"): λ.IntType,
						λ.NewStr("tags"):          λ.ListType,
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			TEDIE__NATIVE_FORMATS = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("low"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("width"):  λ.NewInt(320),
					λ.NewStr("height"): λ.NewInt(180),
				}),
				λ.NewStr("medium"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("width"):  λ.NewInt(512),
					λ.NewStr("height"): λ.NewInt(288),
				}),
				λ.NewStr("high"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("width"):  λ.NewInt(854),
					λ.NewStr("height"): λ.NewInt(480),
				}),
			})
			TEDIE__extract_info = λ.NewFunction("_extract_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo_json λ.Object
						ϒself      = λargs[0]
						ϒwebpage   = λargs[1]
					)
					ϒinfo_json = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)q\\(\\s*\"\\w+.init\"\\s*,\\s*({.+})\\)\\s*</script>"), ϒwebpage, λ.NewStr("info json"))
					return λ.Cal(Ωjson.ϒloads, ϒinfo_json)
				})
			TEDIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdesktop_url λ.Object
						ϒm           λ.Object
						ϒname        λ.Object
						ϒself        = λargs[0]
						ϒurl         = λargs[1]
					)
					ϒm = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl, Ωre.VERBOSE)
					if λ.IsTrue(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("type")), "startswith", nil), λ.NewStr("embed"))) {
						ϒdesktop_url = λ.Add(λ.Add(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("proto")), λ.NewStr("www")), λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("urlmain")))
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒdesktop_url, λ.NewStr("TED"))
					}
					ϒname = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("name"))
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("type_talk"))) {
						return λ.Cal(λ.GetAttr(ϒself, "_talk_info", nil), ϒurl, ϒname)
					} else {
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("type_watch"))) {
							return λ.Cal(λ.GetAttr(ϒself, "_watch_info", nil), ϒurl, ϒname)
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_playlist_videos_info", nil), ϒurl, ϒname)
						}
					}
					return λ.None
				})
			TEDIE__playlist_videos_info = λ.NewFunction("_playlist_videos_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "name"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo             λ.Object
						ϒname             = λargs[2]
						ϒplaylist_entries λ.Object
						ϒplaylist_info    λ.Object
						ϒself             = λargs[0]
						ϒurl              = λargs[1]
						ϒwebpage          λ.Object
					)
					λ.NewStr("Returns the videos of the playlist")
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒname, λ.NewStr("Downloading playlist webpage"))
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_extract_info", nil), ϒwebpage)
					ϒplaylist_info = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒinfo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("__INITIAL_DATA__")), λ.NewStr("playlist"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒinfo, λ.NewStr("playlist"))
						}
					}()
					ϒplaylist_entries = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒtalk λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, func() λ.Object {
									if λv := λ.Cal(ϒtry_get, ϒinfo, λ.NewFunction("<lambda>",
										[]λ.Param{
											{Name: "x"},
										},
										0, false, false,
										func(λargs []λ.Object) λ.Object {
											var (
												ϒx = λargs[0]
											)
											return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("__INITIAL_DATA__")), λ.NewStr("talks"))
										}), λ.DictType); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.GetItem(ϒinfo, λ.NewStr("talks"))
									}
								}())
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒtalk = τmp1
									λgen.Yield(λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Add(λ.NewStr("http://www.ted.com/talks/"), λ.GetItem(ϒtalk, λ.NewStr("slug"))), λ.Cal(λ.GetAttr(ϒself, "ie_key", nil))))
								}
								return λ.None
							})
						})))
					return λ.Call(λ.GetAttr(ϒself, "playlist_result", nil), λ.NewArgs(ϒplaylist_entries), λ.KWArgs{
						{Name: "playlist_id", Value: λ.Cal(ϒcompat_str, λ.GetItem(ϒplaylist_info, λ.NewStr("id")))},
						{Name: "playlist_title", Value: λ.GetItem(ϒplaylist_info, λ.NewStr("title"))},
					})
				})
			TEDIE__talk_info = λ.NewFunction("_talk_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "video_name"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaudio_download   λ.Object
						ϒbitrate          λ.Object
						ϒbitrate_url      λ.Object
						ϒdata             λ.Object
						ϒext_url          λ.Object
						ϒexternal         λ.Object
						ϒf                λ.Object
						ϒfinfo            λ.Object
						ϒformat_id        λ.Object
						ϒformats          λ.Object
						ϒh264_url         λ.Object
						ϒhttp_url         λ.Object
						ϒinfo             λ.Object
						ϒm3u8_format      λ.Object
						ϒm3u8_formats     λ.Object
						ϒnative_downloads λ.Object
						ϒplayer_talk      λ.Object
						ϒresource         λ.Object
						ϒresources        λ.Object
						ϒresources_       λ.Object
						ϒself             = λargs[0]
						ϒservice          λ.Object
						ϒstream_url       λ.Object
						ϒstreamer         λ.Object
						ϒtalk_info        λ.Object
						ϒtitle            λ.Object
						ϒurl              = λargs[1]
						ϒvideo_id         λ.Object
						ϒvideo_name       = λargs[2]
						ϒwebpage          λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
						τmp2              λ.Object
						τmp3              λ.Object
					)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_name)
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_extract_info", nil), ϒwebpage)
					ϒdata = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒinfo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.NewStr("__INITIAL_DATA__"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒinfo
						}
					}()
					ϒtalk_info = λ.GetItem(λ.GetItem(ϒdata, λ.NewStr("talks")), λ.NewInt(0))
					ϒtitle = λ.Cal(λ.GetAttr(λ.GetItem(ϒtalk_info, λ.NewStr("title")), "strip", nil))
					ϒnative_downloads = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒtalk_info, λ.NewTuple(
							λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("downloads")), λ.NewStr("nativeDownloads"))
								}),
							λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(ϒx, λ.NewStr("nativeDownloads"))
								}),
						), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒformat_id  λ.Object
									ϒformat_url λ.Object
									τmp0        λ.Object
									τmp1        λ.Object
									τmp2        λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒnative_downloads, "items", nil)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
									ϒformat_url = λ.GetItem(τmp2, λ.NewInt(1))
									if λ.IsTrue(λ.NewBool(ϒformat_url != λ.None)) {
										λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"):       ϒformat_url,
											λ.NewStr("format_id"): ϒformat_id,
											λ.NewStr("format"):    ϒformat_id,
										}))
									}
								}
								return λ.None
							})
						})))
					if λ.IsTrue(ϒformats) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒf = τmp1
							ϒfinfo = λ.Cal(λ.GetAttr(λ.GetAttr(ϒself, "_NATIVE_FORMATS", nil), "get", nil), λ.GetItem(ϒf, λ.NewStr("format_id")))
							if λ.IsTrue(ϒfinfo) {
								λ.Cal(λ.GetAttr(ϒf, "update", nil), ϒfinfo)
							}
						}
					}
					ϒplayer_talk = λ.GetItem(λ.GetItem(ϒtalk_info, λ.NewStr("player_talks")), λ.NewInt(0))
					ϒexternal = λ.Cal(λ.GetAttr(ϒplayer_talk, "get", nil), λ.NewStr("external"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒexternal, λ.DictType)) {
						ϒservice = λ.Cal(λ.GetAttr(ϒexternal, "get", nil), λ.NewStr("service"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒservice, ϒcompat_str)) {
							ϒext_url = λ.None
							if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒservice, "lower", nil)), λ.NewStr("youtube"))) {
								ϒext_url = λ.Cal(λ.GetAttr(ϒexternal, "get", nil), λ.NewStr("code"))
							}
							return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), func() λ.Object {
								if λv := ϒext_url; λ.IsTrue(λv) {
									return λv
								} else {
									return λ.GetItem(ϒexternal, λ.NewStr("uri"))
								}
							}())
						}
					}
					ϒresources_ = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒplayer_talk, "get", nil), λ.NewStr("resources")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒtalk_info, "get", nil), λ.NewStr("resources"))
						}
					}()
					ϒhttp_url = λ.None
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒresources_, "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒresources = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("h264"))) {
							τmp2 = λ.Cal(λ.BuiltinIter, ϒresources)
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								ϒresource = τmp3
								ϒh264_url = λ.Cal(λ.GetAttr(ϒresource, "get", nil), λ.NewStr("file"))
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒh264_url))) {
									continue
								}
								ϒbitrate = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒresource, "get", nil), λ.NewStr("bitrate")))
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"): ϒh264_url,
									λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-%sk"), λ.NewTuple(
										ϒformat_id,
										ϒbitrate,
									)),
									λ.NewStr("tbr"): ϒbitrate,
								}))
								if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.NewStr("\\d+k"), ϒh264_url)) {
									ϒhttp_url = ϒh264_url
								}
							}
						} else {
							if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("rtmp"))) {
								ϒstreamer = λ.Cal(λ.GetAttr(ϒtalk_info, "get", nil), λ.NewStr("streamer"))
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒstreamer))) {
									continue
								}
								τmp2 = λ.Cal(λ.BuiltinIter, ϒresources)
								for {
									if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
										break
									}
									ϒresource = τmp3
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
											ϒformat_id,
											λ.Cal(λ.GetAttr(ϒresource, "get", nil), λ.NewStr("name")),
										)),
										λ.NewStr("url"):       ϒstreamer,
										λ.NewStr("play_path"): λ.GetItem(ϒresource, λ.NewStr("file")),
										λ.NewStr("ext"):       λ.NewStr("flv"),
										λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒresource, "get", nil), λ.NewStr("width"))),
										λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒresource, "get", nil), λ.NewStr("height"))),
										λ.NewStr("tbr"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒresource, "get", nil), λ.NewStr("bitrate"))),
									}))
								}
							} else {
								if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("hls"))) {
									if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒresources, λ.DictType)))) {
										continue
									}
									ϒstream_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒresources, "get", nil), λ.NewStr("stream")))
									if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒstream_url))) {
										continue
									}
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒstream_url,
										ϒvideo_name,
										λ.NewStr("mp4"),
									), λ.KWArgs{
										{Name: "m3u8_id", Value: ϒformat_id},
										{Name: "fatal", Value: λ.False},
									}))
								}
							}
						}
					}
					ϒm3u8_formats = λ.Cal(λ.ListType, λ.Cal(λ.FilterIteratorType, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "f"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒf = λargs[0]
							)
							return func() λ.Object {
								if λv := λ.Eq(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("protocol")), λ.NewStr("m3u8")); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Ne(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("vcodec")), λ.NewStr("none"))
								}
							}()
						}), ϒformats))
					if λ.IsTrue(ϒhttp_url) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒm3u8_format = τmp1
							ϒbitrate = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("(\\d+k)"),
								λ.GetItem(ϒm3u8_format, λ.NewStr("url")),
								λ.NewStr("bitrate"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒbitrate))) {
								continue
							}
							ϒbitrate_url = λ.Cal(Ωre.ϒsub, λ.NewStr("\\d+k"), ϒbitrate, ϒhttp_url)
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒself, "_is_valid_url", nil), ϒbitrate_url, ϒvideo_name, λ.Mod(λ.NewStr("%s bitrate"), ϒbitrate))))) {
								continue
							}
							ϒf = λ.Cal(λ.GetAttr(ϒm3u8_format, "copy", nil))
							λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒbitrate_url,
								λ.NewStr("format_id"): λ.Cal(λ.GetAttr(λ.GetItem(ϒm3u8_format, λ.NewStr("format_id")), "replace", nil), λ.NewStr("hls"), λ.NewStr("http")),
								λ.NewStr("protocol"):  λ.NewStr("http"),
							}))
							if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("acodec")), λ.NewStr("none"))) {
								λ.DelItem(ϒf, λ.NewStr("acodec"))
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
					}
					ϒaudio_download = λ.Cal(λ.GetAttr(ϒtalk_info, "get", nil), λ.NewStr("audioDownload"))
					if λ.IsTrue(ϒaudio_download) {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒaudio_download,
							λ.NewStr("format_id"): λ.NewStr("audio"),
							λ.NewStr("vcodec"):    λ.NewStr("none"),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒvideo_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒtalk_info, λ.NewStr("id")))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("title"): ϒtitle,
						λ.NewStr("uploader"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒplayer_talk, "get", nil), λ.NewStr("speaker")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒtalk_info, "get", nil), λ.NewStr("speaker"))
							}
						}(),
						λ.NewStr("thumbnail"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒplayer_talk, "get", nil), λ.NewStr("thumb")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒtalk_info, "get", nil), λ.NewStr("thumb"))
							}
						}(),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage),
						λ.NewStr("subtitles"):   λ.Cal(λ.GetAttr(ϒself, "_get_subtitles", nil), ϒvideo_id, ϒtalk_info),
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("duration"):    λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒtalk_info, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("view_count"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("viewed_count"))),
						λ.NewStr("comment_count"): λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("comments")), λ.NewStr("count"))
							}))),
						λ.NewStr("tags"): λ.Cal(ϒtry_get, ϒtalk_info, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.NewStr("tags"))
							}), λ.ListType),
					})
				})
			TEDIE__get_subtitles = λ.NewFunction("_get_subtitles",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
					{Name: "talk_info"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒlang_code     λ.Object
						ϒlanguage      λ.Object
						ϒself          = λargs[0]
						ϒsub_lang_list λ.Object
						ϒtalk_info     = λargs[2]
						ϒvideo_id      = λargs[1]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					_ = ϒself
					ϒsub_lang_list = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(ϒtry_get, ϒtalk_info, λ.NewTuple(
						λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("downloads")), λ.NewStr("languages"))
							}),
						λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.NewStr("languages"))
							}),
					), λ.ListType))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒlanguage = τmp1
						ϒlang_code = func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒlanguage, "get", nil), λ.NewStr("languageCode")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒlanguage, "get", nil), λ.NewStr("ianaCode"))
							}
						}()
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒlang_code))) {
							continue
						}
						λ.SetItem(ϒsub_lang_list, ϒlang_code, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒext λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewList(
										λ.NewStr("ted"),
										λ.NewStr("srt"),
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒext = τmp1
										λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"): λ.Mod(λ.NewStr("http://www.ted.com/talks/subtitles/id/%s/lang/%s/format/%s"), λ.NewTuple(
												ϒvideo_id,
												ϒlang_code,
												ϒext,
											)),
											λ.NewStr("ext"): ϒext,
										}))
									}
									return λ.None
								})
							}))))
					}
					return ϒsub_lang_list
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):               TEDIE_IE_NAME,
				λ.NewStr("_NATIVE_FORMATS"):       TEDIE__NATIVE_FORMATS,
				λ.NewStr("_TESTS"):                TEDIE__TESTS,
				λ.NewStr("_VALID_URL"):            TEDIE__VALID_URL,
				λ.NewStr("_extract_info"):         TEDIE__extract_info,
				λ.NewStr("_get_subtitles"):        TEDIE__get_subtitles,
				λ.NewStr("_playlist_videos_info"): TEDIE__playlist_videos_info,
				λ.NewStr("_real_extract"):         TEDIE__real_extract,
				λ.NewStr("_talk_info"):            TEDIE__talk_info,
			})
		}())
	})
}
