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
 * spankbang/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/spankbang.py
 */

package spankbang

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	SpankBangIE         λ.Object
	SpankBangPlaylistIE λ.Object
	ϒdetermine_ext      λ.Object
	ϒmerge_dicts        λ.Object
	ϒorderedSet         λ.Object
	ϒparse_duration     λ.Object
	ϒparse_resolution   λ.Object
	ϒstr_to_int         λ.Object
	ϒurl_or_none        λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒorderedSet = Ωutils.ϒorderedSet
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_resolution = Ωutils.ϒparse_resolution
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		SpankBangIE = λ.Cal(λ.TypeType, λ.NewStr("SpankBangIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SpankBangIE__VALID_URL    λ.Object
				SpankBangIE__real_extract λ.Object
			)
			SpankBangIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?spankbang\\.com/(?P<id>[\\da-z]+)/(?:video|play|embed)\\b")
			SpankBangIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						STREAM_URL_PREFIX λ.Object
						ϒage_limit        λ.Object
						ϒdescription      λ.Object
						ϒduration         λ.Object
						ϒextract_format   λ.Object
						ϒformat_id        λ.Object
						ϒformat_url       λ.Object
						ϒformats          λ.Object
						ϒinfo             λ.Object
						ϒmobj             λ.Object
						ϒself             = λargs[0]
						ϒstream           λ.Object
						ϒstream_key       λ.Object
						ϒthumbnail        λ.Object
						ϒtitle            λ.Object
						ϒuploader         λ.Object
						ϒurl              = λargs[1]
						ϒvideo_id         λ.Object
						ϒview_count       λ.Object
						ϒwebpage          λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
						τmp2              λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.Cal(λ.GetAttr(ϒurl, "replace", nil), λ.Mod(λ.NewStr("/%s/embed"), ϒvideo_id), λ.Mod(λ.NewStr("/%s/video"), ϒvideo_id)),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Cookie"): λ.NewStr("country=US"),
						})},
					})
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.NewStr("<[^>]+\\b(?:id|class)=[\"\\']video_removed"), ϒwebpage)) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Video %s is not available"), ϒvideo_id)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒformats = λ.NewList()
					ϒextract_format = λ.NewFunction("extract_format",
						[]λ.Param{
							{Name: "format_id"},
							{Name: "format_url"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒext        λ.Object
								ϒf          λ.Object
								ϒf_url      λ.Object
								ϒformat_id  = λargs[0]
								ϒformat_url = λargs[1]
							)
							ϒf_url = λ.Cal(ϒurl_or_none, ϒformat_url)
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒf_url))) {
								return λ.None
							}
							ϒf = λ.Cal(ϒparse_resolution, ϒformat_id)
							ϒext = λ.Cal(ϒdetermine_ext, ϒf_url)
							if λ.IsTrue(func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒformat_id, "startswith", nil), λ.NewStr("m3u8")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒext, λ.NewStr("m3u8"))
								}
							}()) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒf_url,
									ϒvideo_id,
									λ.NewStr("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
									{Name: "m3u8_id", Value: λ.NewStr("hls")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.GetAttr(ϒformat_id, "startswith", nil), λ.NewStr("mpd")); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Eq(ϒext, λ.NewStr("mpd"))
									}
								}()) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
										ϒf_url,
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "mpd_id", Value: λ.NewStr("dash")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									if λ.IsTrue(func() λ.Object {
										if λv := λ.Eq(ϒext, λ.NewStr("mp4")); λ.IsTrue(λv) {
											return λv
										} else if λv := λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("width")); λ.IsTrue(λv) {
											return λv
										} else {
											return λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("height"))
										}
									}()) {
										λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"):       ϒf_url,
											λ.NewStr("format_id"): ϒformat_id,
										}))
										λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
									}
								}
							}
							return λ.None
						})
					STREAM_URL_PREFIX = λ.NewStr("stream_url_")
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.Mod(λ.NewStr("%s(?P<id>[^\\s=]+)\\s*=\\s*([\"\\'])(?P<url>(?:(?!\\2).)+)\\2"), STREAM_URL_PREFIX), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmobj = τmp1
						λ.Cal(ϒextract_format, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"), λ.NewStr("url")))
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformats))) {
						ϒstream_key = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("data-streamkey\\s*=\\s*([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
							ϒwebpage,
							λ.NewStr("stream key"),
						), λ.KWArgs{
							{Name: "group", Value: λ.NewStr("value")},
						})
						ϒstream = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.NewStr("https://spankbang.com/api/videos/stream"),
							ϒvideo_id,
							λ.NewStr("Downloading stream JSON"),
						), λ.KWArgs{
							{Name: "data", Value: λ.Cal(ϒurlencode_postdata, λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):   ϒstream_key,
								λ.NewStr("data"): λ.NewInt(0),
							}))},
							{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("Referer"):          ϒurl,
								λ.NewStr("X-Requested-With"): λ.NewStr("XMLHttpRequest"),
							})},
						})
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒstream, "items", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
							ϒformat_url = λ.GetItem(τmp2, λ.NewInt(1))
							if λ.IsTrue(func() λ.Object {
								if λv := ϒformat_url; !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.BuiltinIsInstance, ϒformat_url, λ.ListType)
								}
							}()) {
								ϒformat_url = λ.GetItem(ϒformat_url, λ.NewInt(0))
							}
							λ.Cal(ϒextract_format, ϒformat_id, ϒformat_url)
						}
					}
					λ.Call(λ.GetAttr(ϒself, "_sort_formats", nil), λ.NewArgs(ϒformats), λ.KWArgs{
						{Name: "field_preference", Value: λ.NewTuple(
							λ.NewStr("preference"),
							λ.NewStr("height"),
							λ.NewStr("width"),
							λ.NewStr("fps"),
							λ.NewStr("tbr"),
							λ.NewStr("format_id"),
						)},
					})
					ϒinfo = λ.Call(λ.GetAttr(ϒself, "_search_json_ld", nil), λ.NewArgs(
						ϒwebpage,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "default", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{})},
					})
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<h1[^>]*>(.+?)</h1>"),
						ϒwebpage,
						λ.NewStr("title"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div[^>]+\\bclass=[\"\\']bottom[^>]+>\\s*<p>[^<]*</p>\\s*<p>([^<]+)"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.NewStr("(?s)<li[^>]+class=[\"\\']profile[^>]+>(.+?)</a>"),
							λ.NewStr("class=\"user\"[^>]*><img[^>]+>([^<]+)"),
						),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div[^>]+\\bclass=[\"\\']right_side[^>]+>\\s*<span>([^<]+)"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("([\\d,.]+)\\s+plays"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒage_limit = λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage)
					return λ.Cal(ϒmerge_dicts, λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): ϒvideo_id,
						λ.NewStr("title"): func() λ.Object {
							if λv := ϒtitle; λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("uploader"):    ϒuploader,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("view_count"):  ϒview_count,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("age_limit"):   ϒage_limit,
					}), ϒinfo)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    SpankBangIE__VALID_URL,
				λ.NewStr("_real_extract"): SpankBangIE__real_extract,
			})
		}())
		SpankBangPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("SpankBangPlaylistIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SpankBangPlaylistIE__VALID_URL λ.Object
			)
			SpankBangPlaylistIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?spankbang\\.com/(?P<id>[\\da-z]+)/playlist/[^/]+")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): SpankBangPlaylistIE__VALID_URL,
			})
		}())
	})
}
