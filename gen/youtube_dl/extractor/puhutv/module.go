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
 * puhutv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/puhutv.py
 */

package puhutv

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError     λ.Object
	InfoExtractor      λ.Object
	PuhuTVIE           λ.Object
	PuhuTVSerieIE      λ.Object
	ϒcompat_HTTPError  λ.Object
	ϒcompat_str        λ.Object
	ϒfloat_or_none     λ.Object
	ϒint_or_none       λ.Object
	ϒparse_resolution  λ.Object
	ϒstr_or_none       λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
	ϒurl_or_none       λ.Object
	ϒurljoin           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒparse_resolution = Ωutils.ϒparse_resolution
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurljoin = Ωutils.ϒurljoin
		PuhuTVIE = λ.Cal(λ.TypeType, λ.NewStr("PuhuTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PuhuTVIE_IE_NAME       λ.Object
				PuhuTVIE__TESTS        λ.Object
				PuhuTVIE__VALID_URL    λ.Object
				PuhuTVIE__real_extract λ.Object
			)
			PuhuTVIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?puhutv\\.com/(?P<id>[^/?#&]+)-izle")
			PuhuTVIE_IE_NAME = λ.NewStr("puhutv")
			PuhuTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://puhutv.com/sut-kardesler-izle"),
					λ.NewStr("md5"): λ.NewStr("a347470371d56e1585d1b2c8dab01c96"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):           λ.NewStr("5085"),
						λ.NewStr("display_id"):   λ.NewStr("sut-kardesler"),
						λ.NewStr("ext"):          λ.NewStr("mp4"),
						λ.NewStr("title"):        λ.NewStr("Süt Kardeşler"),
						λ.NewStr("description"):  λ.NewStr("md5:ca09da25b7e57cbb5a9280d6e48d17aa"),
						λ.NewStr("thumbnail"):    λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("duration"):     λ.NewFloat(4832.44),
						λ.NewStr("creator"):      λ.NewStr("Arzu Film"),
						λ.NewStr("timestamp"):    λ.NewInt(1561062602),
						λ.NewStr("upload_date"):  λ.NewStr("20190620"),
						λ.NewStr("release_year"): λ.NewInt(1976),
						λ.NewStr("view_count"):   λ.IntType,
						λ.NewStr("tags"):         λ.ListType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://puhutv.com/jet-sosyete-1-bolum-izle"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://puhutv.com/dip-1-bolum-izle"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			PuhuTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒadd_http_from_hls λ.Object
						ϒcontent           λ.Object
						ϒcreator           λ.Object
						ϒdisplay_id        λ.Object
						ϒf                 λ.Object
						ϒformat_id         λ.Object
						ϒformats           λ.Object
						ϒgenre             λ.Object
						ϒgenre_name        λ.Object
						ϒimage_id          λ.Object
						ϒimage_url         λ.Object
						ϒimages            λ.Object
						ϒinfo              λ.Object
						ϒis_hls            λ.Object
						ϒlang              λ.Object
						ϒm3u8_f            λ.Object
						ϒm3u8_formats      λ.Object
						ϒmedia_url         λ.Object
						ϒplaylist          λ.Object
						ϒquality           λ.Object
						ϒself              = λargs[0]
						ϒshow              λ.Object
						ϒsub_url           λ.Object
						ϒsubtitle          λ.Object
						ϒsubtitles         λ.Object
						ϒt                 λ.Object
						ϒtags              λ.Object
						ϒthumbnails        λ.Object
						ϒtitle             λ.Object
						ϒurl               = λargs[1]
						ϒurls              λ.Object
						ϒvideo             λ.Object
						ϒvideo_format      λ.Object
						ϒvideo_id          λ.Object
						ϒvideos            λ.Object
						τmp0               λ.Object
						τmp1               λ.Object
						τmp2               λ.Object
						τmp3               λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒinfo = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Cal(ϒurljoin, ϒurl, λ.Mod(λ.NewStr("/api/slug/%s-izle"), ϒdisplay_id)), ϒdisplay_id), λ.NewStr("data"))
					ϒvideo_id = λ.Cal(ϒcompat_str, λ.GetItem(ϒinfo, λ.NewStr("id")))
					ϒshow = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("name")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒshow, λ.NewStr("name"))
						}
					}()
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("display_name"))) {
						ϒtitle = λ.Mod(λ.NewStr("%s %s"), λ.NewTuple(
							ϒtitle,
							λ.GetItem(ϒinfo, λ.NewStr("display_name")),
						))
					}
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), ϒcompat_HTTPError); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil), λ.NewInt(403))
									}
								}()) {
									λ.Cal(λ.GetAttr(ϒself, "raise_geo_restricted", nil))
								}
								panic(λ.Raise(λex))
							}},
						)
						ϒvideos = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Mod(λ.NewStr("https://puhutv.com/api/assets/%s/videos"), ϒvideo_id),
							ϒdisplay_id,
							λ.NewStr("Downloading video JSON"),
						), λ.KWArgs{
							{Name: "headers", Value: λ.Cal(λ.GetAttr(ϒself, "geo_verification_headers", nil))},
						})
						return λ.BlockExitNormally, nil
					}()
					ϒurls = λ.NewList()
					ϒformats = λ.NewList()
					ϒadd_http_from_hls = λ.NewFunction("add_http_from_hls",
						[]λ.Param{
							{Name: "m3u8_f"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒf        λ.Object
								ϒhttp_url λ.Object
								ϒm3u8_f   = λargs[0]
							)
							ϒhttp_url = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.GetItem(ϒm3u8_f, λ.NewStr("url")), "replace", nil), λ.NewStr("/hls/"), λ.NewStr("/mp4/")), "replace", nil), λ.NewStr("/chunklist.m3u8"), λ.NewStr(".mp4"))
							if λ.IsTrue(λ.Ne(ϒhttp_url, λ.GetItem(ϒm3u8_f, λ.NewStr("url")))) {
								ϒf = λ.Cal(λ.GetAttr(ϒm3u8_f, "copy", nil))
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("format_id"): λ.Cal(λ.GetAttr(λ.GetItem(ϒf, λ.NewStr("format_id")), "replace", nil), λ.NewStr("hls"), λ.NewStr("http")),
									λ.NewStr("protocol"):  λ.NewStr("http"),
									λ.NewStr("url"):       ϒhttp_url,
								}))
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
							}
							return λ.None
						})
					τmp1 = λ.Cal(λ.BuiltinIter, λ.GetItem(λ.GetItem(ϒvideos, λ.NewStr("data")), λ.NewStr("videos")))
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒvideo = τmp0
						ϒmedia_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("url")))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒmedia_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurls, ϒmedia_url))
							}
						}()) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒurls, "append", nil), ϒmedia_url)
						ϒplaylist = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("is_playlist"))
						if λ.IsTrue(func() λ.Object {
							if λv := func() λ.Object {
								if λv := λ.Eq(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("stream_type")), λ.NewStr("hls")); !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(ϒplaylist == λ.True)
								}
							}(); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒmedia_url, λ.NewStr("playlist.m3u8")))
							}
						}()) {
							ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒmedia_url,
								ϒvideo_id,
								λ.NewStr("mp4"),
							), λ.KWArgs{
								{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							})
							τmp2 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								ϒm3u8_f = τmp3
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒm3u8_f)
								λ.Cal(ϒadd_http_from_hls, ϒm3u8_f)
							}
							continue
						}
						ϒquality = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("quality")))
						ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):    ϒmedia_url,
							λ.NewStr("ext"):    λ.NewStr("mp4"),
							λ.NewStr("height"): ϒquality,
						})
						ϒvideo_format = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("video_format"))
						ϒis_hls = func() λ.Object {
							if λv := func() λ.Object {
								if λv := λ.Eq(ϒvideo_format, λ.NewStr("hls")); λ.IsTrue(λv) {
									return λv
								} else if λv := λ.NewBool(λ.Contains(ϒmedia_url, λ.NewStr("/hls/"))); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(λ.Contains(ϒmedia_url, λ.NewStr("/chunklist.m3u8")))
								}
							}(); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(ϒplaylist == λ.False)
							}
						}()
						if λ.IsTrue(ϒis_hls) {
							ϒformat_id = λ.NewStr("hls")
							λ.SetItem(ϒf, λ.NewStr("protocol"), λ.NewStr("m3u8_native"))
						} else {
							if λ.IsTrue(λ.Eq(ϒvideo_format, λ.NewStr("mp4"))) {
								ϒformat_id = λ.NewStr("http")
							} else {
								continue
							}
						}
						if λ.IsTrue(ϒquality) {
							τmp2 = λ.IAdd(ϒformat_id, λ.Mod(λ.NewStr("-%sp"), ϒquality))
							ϒformat_id = τmp2
						}
						λ.SetItem(ϒf, λ.NewStr("format_id"), ϒformat_id)
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						if λ.IsTrue(ϒis_hls) {
							λ.Cal(ϒadd_http_from_hls, ϒf)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒcreator = λ.Cal(ϒtry_get, ϒshow, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("producer")), λ.NewStr("name"))
						}), ϒcompat_str)
					ϒcontent = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("content")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒimages = func() λ.Object {
						if λv := λ.Cal(ϒtry_get, ϒcontent, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("images")), λ.NewStr("wide"))
							}), λ.DictType); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒthumbnails = λ.NewList()
					τmp1 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒimages, "items", nil)))
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						τmp2 = τmp0
						ϒimage_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒimage_url = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒimage_url, ϒcompat_str)))) {
							continue
						}
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒimage_url, "startswith", nil), λ.NewTuple(
							λ.NewStr("http"),
							λ.NewStr("//"),
						))))) {
							ϒimage_url = λ.Mod(λ.NewStr("https://%s"), ϒimage_url)
						}
						ϒt = λ.Cal(ϒparse_resolution, ϒimage_id)
						λ.Cal(λ.GetAttr(ϒt, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):  ϒimage_id,
							λ.NewStr("url"): ϒimage_url,
						}))
						λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), ϒt)
					}
					ϒtags = λ.NewList()
					τmp1 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒshow, "get", nil), λ.NewStr("genres")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒgenre = τmp0
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒgenre, λ.DictType)))) {
							continue
						}
						ϒgenre_name = λ.Cal(λ.GetAttr(ϒgenre, "get", nil), λ.NewStr("name"))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒgenre_name; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, ϒgenre_name, ϒcompat_str)
							}
						}()) {
							λ.Cal(λ.GetAttr(ϒtags, "append", nil), ϒgenre_name)
						}
					}
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					τmp1 = λ.Cal(λ.BuiltinIter, func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒcontent, "get", nil), λ.NewStr("subtitles")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewList()
						}
					}())
					for {
						if τmp0 = λ.NextDefault(τmp1, λ.AfterLast); τmp0 == λ.AfterLast {
							break
						}
						ϒsubtitle = τmp0
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒsubtitle, λ.DictType)))) {
							continue
						}
						ϒlang = λ.Cal(λ.GetAttr(ϒsubtitle, "get", nil), λ.NewStr("language"))
						ϒsub_url = λ.Cal(ϒurl_or_none, func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒsubtitle, "get", nil), λ.NewStr("url")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒsubtitle, "get", nil), λ.NewStr("file"))
							}
						}())
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒlang)); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒlang, ϒcompat_str))); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(ϒsub_url))
							}
						}()) {
							continue
						}
						λ.SetItem(ϒsubtitles, λ.Cal(λ.GetAttr(λ.GetAttr(ϒself, "_SUBTITLE_LANGS", nil), "get", nil), ϒlang, ϒlang), λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒsub_url,
						})))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
						λ.NewStr("title"):      ϒtitle,
						λ.NewStr("description"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("description")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒshow, "get", nil), λ.NewStr("description"))
							}
						}(),
						λ.NewStr("season_id"):      λ.Cal(ϒstr_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("season_id"))),
						λ.NewStr("season_number"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("season_number"))),
						λ.NewStr("episode_number"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("episode_number"))),
						λ.NewStr("release_year"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒshow, "get", nil), λ.NewStr("released_at"))),
						λ.NewStr("timestamp"):      λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("created_at"))),
						λ.NewStr("creator"):        ϒcreator,
						λ.NewStr("view_count"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒcontent, "get", nil), λ.NewStr("watch_count"))),
						λ.NewStr("duration"):       λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒcontent, "get", nil), λ.NewStr("duration_in_ms")), λ.NewInt(1000)),
						λ.NewStr("tags"):           ϒtags,
						λ.NewStr("subtitles"):      ϒsubtitles,
						λ.NewStr("thumbnails"):     ϒthumbnails,
						λ.NewStr("formats"):        ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       PuhuTVIE_IE_NAME,
				λ.NewStr("_TESTS"):        PuhuTVIE__TESTS,
				λ.NewStr("_VALID_URL"):    PuhuTVIE__VALID_URL,
				λ.NewStr("_real_extract"): PuhuTVIE__real_extract,
			})
		}())
		PuhuTVSerieIE = λ.Cal(λ.TypeType, λ.NewStr("PuhuTVSerieIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PuhuTVSerieIE__VALID_URL λ.Object
			)
			PuhuTVSerieIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?puhutv\\.com/(?P<id>[^/?#&]+)-detay")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): PuhuTVSerieIE__VALID_URL,
			})
		}())
	})
}
