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
 * eporner/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/eporner.py
 */

package eporner

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	EpornerIE       λ.Object
	ExtractorError  λ.Object
	InfoExtractor   λ.Object
	ϒcompat_str     λ.Object
	ϒencode_base_n  λ.Object
	ϒint_or_none    λ.Object
	ϒmerge_dicts    λ.Object
	ϒparse_duration λ.Object
	ϒstr_to_int     λ.Object
	ϒurl_or_none    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒencode_base_n = Ωutils.ϒencode_base_n
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmerge_dicts = Ωutils.ϒmerge_dicts
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒurl_or_none = Ωutils.ϒurl_or_none
		EpornerIE = λ.Cal(λ.TypeType, λ.NewStr("EpornerIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				EpornerIE__VALID_URL    λ.Object
				EpornerIE__real_extract λ.Object
			)
			EpornerIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?eporner\\.com/(?:hd-porn|embed)/(?P<id>\\w+)(?:/(?P<display_id>[\\w-]+))?")
			EpornerIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcalc_hash    λ.Object
						ϒdisplay_id   λ.Object
						ϒduration     λ.Object
						ϒformat_dict  λ.Object
						ϒformat_id    λ.Object
						ϒformats      λ.Object
						ϒformats_dict λ.Object
						ϒfps          λ.Object
						ϒhash         λ.Object
						ϒheight       λ.Object
						ϒjson_ld      λ.Object
						ϒkind         λ.Object
						ϒmobj         λ.Object
						ϒself         = λargs[0]
						ϒsources      λ.Object
						ϒsrc          λ.Object
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒurlh         λ.Object
						ϒvideo        λ.Object
						ϒvideo_id     λ.Object
						ϒview_count   λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
						τmp3          λ.Object
						τmp4          λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), ϒurl, ϒdisplay_id)
					ϒwebpage = λ.GetItem(τmp0, λ.NewInt(0))
					ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), λ.Cal(ϒcompat_str, λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil))))
					ϒhash = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("hash\\s*:\\s*[\"\\']([\\da-f]{32})"), ϒwebpage, λ.NewStr("hash"))
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<title>(.+?) - EPORNER"), ϒwebpage, λ.NewStr("title"))
						}
					}()
					ϒcalc_hash = λ.NewFunction("calc_hash",
						[]λ.Param{
							{Name: "s"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒs = λargs[0]
							)
							return λ.Cal(λ.GetAttr(λ.NewStr(""), "join", nil), λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒlb  λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.RangeType, λ.NewInt(0), λ.NewInt(32), λ.NewInt(8)))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒlb = τmp1
											λgy.Yield(λ.Cal(ϒencode_base_n, λ.Cal(λ.IntType, λ.GetItem(ϒs, λ.NewSlice(ϒlb, λ.Add(ϒlb, λ.NewInt(8)), λ.None)), λ.NewInt(16)), λ.NewInt(36)))
										}
										return λ.None
									})
								})))
						})
					ϒvideo = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("http://www.eporner.com/xhr/video/%s"), ϒvideo_id),
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Downloading video JSON")},
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("hash"):     λ.Cal(ϒcalc_hash, ϒhash),
							λ.NewStr("device"):   λ.NewStr("generic"),
							λ.NewStr("domain"):   λ.NewStr("www.eporner.com"),
							λ.NewStr("fallback"): λ.NewStr("false"),
						})},
					})
					if λ.IsTrue(λ.NewBool(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("available")) == λ.False)) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒvideo, λ.NewStr("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒsources = λ.GetItem(ϒvideo, λ.NewStr("sources"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒsources, "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒkind = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformats_dict = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformats_dict, λ.DictType)))) {
							continue
						}
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒformats_dict, "items", nil)))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							τmp4 = τmp3
							ϒformat_id = λ.GetItem(τmp4, λ.NewInt(0))
							ϒformat_dict = λ.GetItem(τmp4, λ.NewInt(1))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_dict, λ.DictType)))) {
								continue
							}
							ϒsrc = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒformat_dict, "get", nil), λ.NewStr("src")))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒsrc)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒsrc, "startswith", nil), λ.NewStr("http"))))
								}
							}()) {
								continue
							}
							if λ.IsTrue(λ.Eq(ϒkind, λ.NewStr("hls"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒsrc,
									ϒdisplay_id,
									λ.NewStr("mp4"),
								), λ.KWArgs{
									{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
									{Name: "m3u8_id", Value: ϒkind},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.NewStr("(\\d+)[pP]"),
									ϒformat_id,
									λ.NewStr("height"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}))
								ϒfps = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.NewStr("(\\d+)fps"),
									ϒformat_id,
									λ.NewStr("fps"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								}))
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):       ϒsrc,
									λ.NewStr("format_id"): ϒformat_id,
									λ.NewStr("height"):    ϒheight,
									λ.NewStr("fps"):       ϒfps,
								}))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒjson_ld = λ.Call(λ.GetAttr(ϒself, "_search_json_ld", nil), λ.NewArgs(
						ϒwebpage,
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "default", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{})},
					})
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("duration"),
						ϒwebpage,
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("id=\"cinemaviews\">\\s*([0-9,]+)\\s*<small>views"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.Cal(ϒmerge_dicts, ϒjson_ld, λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
						λ.NewStr("title"):      ϒtitle,
						λ.NewStr("duration"):   ϒduration,
						λ.NewStr("view_count"): ϒview_count,
						λ.NewStr("formats"):    ϒformats,
						λ.NewStr("age_limit"):  λ.NewInt(18),
					}))
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    EpornerIE__VALID_URL,
				λ.NewStr("_real_extract"): EpornerIE__real_extract,
			})
		}())
	})
}
