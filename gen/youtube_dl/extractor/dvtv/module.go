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
 * dvtv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/dvtv.py
 */

package dvtv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DVTVIE         λ.Object
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒjs_to_json    λ.Object
	ϒmimetype2ext  λ.Object
	ϒparse_iso8601 λ.Object
	ϒtry_get       λ.Object
	ϒunescapeHTML  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		DVTVIE = λ.Cal(λ.TypeType, λ.StrLiteral("DVTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DVTVIE_IE_NAME               λ.Object
				DVTVIE__VALID_URL            λ.Object
				DVTVIE__parse_video_metadata λ.Object
				DVTVIE__real_extract         λ.Object
			)
			DVTVIE_IE_NAME = λ.StrLiteral("dvtv")
			DVTVIE__VALID_URL = λ.StrLiteral("https?://video\\.aktualne\\.cz/(?:[^/]+/)+r~(?P<id>[0-9a-f]{32})")
			DVTVIE__parse_video_metadata = λ.NewFunction("_parse_video_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "js"},
					{Name: "video_id"},
					{Name: "timestamp"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata         λ.Object
						ϒext          λ.Object
						ϒf            λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒheight       λ.Object
						ϒjs           = λargs[1]
						ϒlabel        λ.Object
						ϒlive_starter λ.Object
						ϒself         = λargs[0]
						ϒtimestamp    = λargs[3]
						ϒtitle        λ.Object
						ϒtracks       λ.Object
						ϒvideo        λ.Object
						ϒvideo_id     = λargs[2]
						ϒvideo_type   λ.Object
						ϒvideo_url    λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
						τmp3          λ.Object
						τmp4          λ.Object
						τmp5          λ.Object
					)
					ϒdata = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						ϒjs,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
					})
					ϒtitle = λ.Cal(ϒunescapeHTML, λ.GetItem(ϒdata, λ.StrLiteral("title")))
					ϒlive_starter = λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("plugins")), λ.StrLiteral("liveStarter"))
						}), λ.DictType)
					if λ.IsTrue(ϒlive_starter) {
						λ.Calm(ϒdata, "update", ϒlive_starter)
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.Calm(ϒdata, "get", λ.StrLiteral("tracks"), λ.DictLiteral(map[λ.Object]λ.Object{})), "values"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒtracks = τmp1
						τmp2 = λ.Cal(λ.BuiltinIter, ϒtracks)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒvideo = τmp3
							ϒvideo_url = λ.Calm(ϒvideo, "get", λ.StrLiteral("src"))
							if !λ.IsTrue(ϒvideo_url) {
								continue
							}
							ϒvideo_type = λ.Calm(ϒvideo, "get", λ.StrLiteral("type"))
							ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url, λ.Cal(ϒmimetype2ext, ϒvideo_type))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Eq(ϒvideo_type, λ.StrLiteral("application/vnd.apple.mpegurl")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒext, λ.StrLiteral("m3u8"))
								}
							}()) {
								λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒvideo_id,
									λ.StrLiteral("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
									{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Eq(ϒvideo_type, λ.StrLiteral("application/dash+xml")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(ϒext, λ.StrLiteral("mpd"))
									}
								}()) {
									λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
										ϒvideo_url,
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "mpd_id", Value: λ.StrLiteral("dash")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒlabel = λ.Calm(ϒvideo, "get", λ.StrLiteral("label"))
									ϒheight = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
										λ.StrLiteral("^(\\d+)[pP]"),
										func() λ.Object {
											if λv := ϒlabel; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.StrLiteral("")
											}
										}(),
										λ.StrLiteral("height"),
									), λ.KWArgs{
										{Name: "default", Value: λ.None},
									})
									ϒformat_id = λ.NewList(λ.StrLiteral("http"))
									τmp4 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										ϒext,
										ϒlabel,
									))
									for {
										if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
											break
										}
										ϒf = τmp5
										if λ.IsTrue(ϒf) {
											λ.Calm(ϒformat_id, "append", ϒf)
										}
									}
									λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
										"url":       ϒvideo_url,
										"format_id": λ.Calm(λ.StrLiteral("-"), "join", ϒformat_id),
										"height":    λ.Cal(ϒint_or_none, ϒheight),
									}))
								}
							}
						}
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id": func() λ.Object {
							if λv := λ.Calm(ϒdata, "get", λ.StrLiteral("mediaid")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"title":       ϒtitle,
						"description": λ.Calm(ϒdata, "get", λ.StrLiteral("description")),
						"thumbnail":   λ.Calm(ϒdata, "get", λ.StrLiteral("image")),
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒdata, "get", λ.StrLiteral("duration"))),
						"timestamp":   λ.Cal(ϒint_or_none, ϒtimestamp),
						"formats":     ϒformats,
					})
				})
			DVTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒitem      λ.Object
						ϒitems     λ.Object
						ϒself      = λargs[0]
						ϒtimestamp λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("article:published_time"),
						ϒwebpage,
						λ.StrLiteral("published time"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒitems = λ.Cal(Ωre.ϒfindall, λ.StrLiteral("(?s)playlist\\.push\\(({.+?})\\);"), ϒwebpage)
					if λ.IsTrue(ϒitems) {
						return λ.Calm(ϒself, "playlist_result", λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒi   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒitems)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒi = τmp1
										λgy.Yield(λ.Calm(ϒself, "_parse_video_metadata", ϒi, ϒvideo_id, ϒtimestamp))
									}
									return λ.None
								})
							}))), ϒvideo_id, λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("twitter:title"), ϒwebpage))
					}
					ϒitem = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)BBXPlayer\\.setup\\((.+?)\\);"),
						ϒwebpage,
						λ.StrLiteral("video"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒitem) {
						ϒitem = λ.Cal(Ωre.ϒsub, λ.StrLiteral("\\w+?\\((.+)\\)"), λ.StrLiteral("\\1"), ϒitem)
						return λ.Calm(ϒself, "_parse_video_metadata", ϒitem, ϒvideo_id, ϒtimestamp)
					}
					panic(λ.Raise(λ.Cal(ExtractorError, λ.StrLiteral("Could not find neither video nor playlist"))))
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":               DVTVIE_IE_NAME,
				"_VALID_URL":            DVTVIE__VALID_URL,
				"_parse_video_metadata": DVTVIE__parse_video_metadata,
				"_real_extract":         DVTVIE__real_extract,
			})
		}())
	})
}
