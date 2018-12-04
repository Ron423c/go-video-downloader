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
 * ustream/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ustream.py
 */

package ustream

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError   λ.Object
	InfoExtractor    λ.Object
	UstreamChannelIE λ.Object
	UstreamIE        λ.Object
	ϒcompat_str      λ.Object
	ϒfloat_or_none   λ.Object
	ϒint_or_none     λ.Object
	ϒmimetype2ext    λ.Object
	ϒstr_or_none     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒstr_or_none = Ωutils.ϒstr_or_none
		UstreamIE = λ.Cal(λ.TypeType, λ.NewStr("UstreamIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UstreamIE_IE_NAME       λ.Object
				UstreamIE__VALID_URL    λ.Object
				UstreamIE__extract_url  λ.Object
				UstreamIE__real_extract λ.Object
			)
			UstreamIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ustream\\.tv/(?P<type>recorded|embed|embed/recorded)/(?P<id>\\d+)")
			UstreamIE_IE_NAME = λ.NewStr("ustream")
			UstreamIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmobj    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("<iframe[^>]+?src=([\"\\'])(?P<url>http://www\\.ustream\\.tv/embed/.+?)\\1"), ϒwebpage)
					if λ.IsTrue(λ.NewBool(ϒmobj != λ.None)) {
						return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
					}
					return λ.None
				})
			UstreamIE__extract_url = λ.Cal(λ.StaticMethodType, UstreamIE__extract_url)
			UstreamIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcontent_video_ids λ.Object
						ϒdescription       λ.Object
						ϒdesktop_url       λ.Object
						ϒduration          λ.Object
						ϒerror             λ.Object
						ϒfilesize          λ.Object
						ϒformats           λ.Object
						ϒhls_streams       λ.Object
						ϒm                 λ.Object
						ϒparams            λ.Object
						ϒself              = λargs[0]
						ϒthumbnails        λ.Object
						ϒtimestamp         λ.Object
						ϒtitle             λ.Object
						ϒuploader          λ.Object
						ϒuploader_id       λ.Object
						ϒurl               = λargs[1]
						ϒvideo             λ.Object
						ϒvideo_id          λ.Object
						ϒview_count        λ.Object
						ϒwebpage           λ.Object
					)
					ϒm = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("id"))
					if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("type")), λ.NewStr("embed/recorded"))) {
						ϒvideo_id = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("id"))
						ϒdesktop_url = λ.Add(λ.NewStr("http://www.ustream.tv/recorded/"), ϒvideo_id)
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒdesktop_url, λ.NewStr("Ustream"))
					}
					if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("type")), λ.NewStr("embed"))) {
						ϒvideo_id = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("id"))
						ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
						ϒcontent_video_ids = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("ustream\\.vars\\.offAirContentVideoIds=([^;]+);"), ϒwebpage, λ.NewStr("content video IDs")), ϒvideo_id)
						return λ.Cal(λ.GetAttr(ϒself, "playlist_result", nil), λ.Cal(λ.MapIteratorType, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "u"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒu = λargs[0]
								)
								return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Add(λ.NewStr("http://www.ustream.tv/recorded/"), ϒu), λ.NewStr("Ustream"))
							}), ϒcontent_video_ids), ϒvideo_id)
					}
					ϒparams = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://api.ustream.tv/videos/%s.json"), ϒvideo_id), ϒvideo_id)
					ϒerror = λ.Cal(λ.GetAttr(ϒparams, "get", nil), λ.NewStr("error"))
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s returned error: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo = λ.GetItem(ϒparams, λ.NewStr("video"))
					ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("title"))
					ϒfilesize = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("file_size")))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒformat_id λ.Object
									ϒvideo_url λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
									τmp2       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo, λ.NewStr("media_urls")), "items", nil)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
									ϒvideo_url = λ.GetItem(τmp2, λ.NewInt(1))
									if λ.IsTrue(ϒvideo_url) {
										λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("id"):       ϒvideo_id,
											λ.NewStr("url"):      ϒvideo_url,
											λ.NewStr("ext"):      ϒformat_id,
											λ.NewStr("filesize"): ϒfilesize,
										}))
									}
								}
								return λ.None
							})
						})))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformats))) {
						ϒhls_streams = λ.Call(λ.GetAttr(ϒself, "_get_streams", nil), λ.NewArgs(
							ϒurl,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "app_id_ver", Value: λ.NewTuple(
								λ.NewInt(11),
								λ.NewInt(2),
							)},
						})
						if λ.IsTrue(ϒhls_streams) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								λ.GetItem(λ.GetItem(ϒhls_streams, λ.NewInt(0)), λ.NewStr("url")),
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "ext", Value: λ.NewStr("mp4")},
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
							}))
						}
						λ.NewStr("\n            # DASH streams handling is incomplete as 'url' is missing\n            dash_streams = self._get_streams(url, video_id, app_id_ver=(3, 1))\n            if dash_streams:\n                formats.extend(self._parse_segmented_mp4(dash_streams))\n            ")
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒdescription = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("description"))
					ϒtimestamp = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("created_at")))
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("length")))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("views")))
					ϒuploader = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("owner"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("username"))
					ϒuploader_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("owner"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("id"))
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒthumbnail_id  λ.Object
									ϒthumbnail_url λ.Object
									τmp0           λ.Object
									τmp1           λ.Object
									τmp2           λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("thumbnail"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "items", nil)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒthumbnail_id = λ.GetItem(τmp2, λ.NewInt(0))
									ϒthumbnail_url = λ.GetItem(τmp2, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("id"):  ϒthumbnail_id,
										λ.NewStr("url"): ϒthumbnail_url,
									}))
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnails"):  ϒthumbnails,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("view_count"):  ϒview_count,
						λ.NewStr("uploader"):    ϒuploader,
						λ.NewStr("uploader_id"): ϒuploader_id,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       UstreamIE_IE_NAME,
				λ.NewStr("_VALID_URL"):    UstreamIE__VALID_URL,
				λ.NewStr("_extract_url"):  UstreamIE__extract_url,
				λ.NewStr("_real_extract"): UstreamIE__real_extract,
			})
		}())
		UstreamChannelIE = λ.Cal(λ.TypeType, λ.NewStr("UstreamChannelIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UstreamChannelIE__VALID_URL λ.Object
			)
			UstreamChannelIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ustream\\.tv/channel/(?P<slug>.+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): UstreamChannelIE__VALID_URL,
			})
		}())
	})
}
