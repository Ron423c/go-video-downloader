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
 * xhamster/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/xhamster.py
 */

package xhamster

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	XHamsterEmbedIE     λ.Object
	XHamsterIE          λ.Object
	XHamsterUserIE      λ.Object
	ϒclean_html         λ.Object
	ϒcompat_str         λ.Object
	ϒdetermine_ext      λ.Object
	ϒdict_get           λ.Object
	ϒextract_attributes λ.Object
	ϒint_or_none        λ.Object
	ϒparse_duration     λ.Object
	ϒtry_get            λ.Object
	ϒunified_strdate    λ.Object
	ϒurl_or_none        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒdict_get = Ωutils.ϒdict_get
		ϒextract_attributes = Ωutils.ϒextract_attributes
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurl_or_none = Ωutils.ϒurl_or_none
		XHamsterIE = λ.Cal(λ.TypeType, λ.StrLiteral("XHamsterIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XHamsterIE__DOMAINS      λ.Object
				XHamsterIE__VALID_URL    λ.Object
				XHamsterIE__real_extract λ.Object
			)
			XHamsterIE__DOMAINS = λ.StrLiteral("(?:xhamster\\.(?:com|one|desi)|xhms\\.pro|xhamster\\d+\\.com)")
			XHamsterIE__VALID_URL = λ.Mod(λ.StrLiteral("(?x)\n                    https?://\n                        (?:.+?\\.)?%s/\n                        (?:\n                            movies/(?P<id>[\\dA-Za-z]+)/(?P<display_id>[^/]*)\\.html|\n                            videos/(?P<display_id_2>[^/]*)-(?P<id_2>[\\dA-Za-z]+)\n                        )\n                    "), XHamsterIE__DOMAINS)
			XHamsterIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit       λ.Object
						ϒc               λ.Object
						ϒc_name          λ.Object
						ϒcategories      λ.Object
						ϒcategories_html λ.Object
						ϒcategories_list λ.Object
						ϒcomment_count   λ.Object
						ϒdescription     λ.Object
						ϒdesktop_url     λ.Object
						ϒdislike_count   λ.Object
						ϒdisplay_id      λ.Object
						ϒduration        λ.Object
						ϒerror           λ.Object
						ϒfilesize        λ.Object
						ϒformat_id       λ.Object
						ϒformat_item     λ.Object
						ϒformat_url      λ.Object
						ϒformat_urls     λ.Object
						ϒformats         λ.Object
						ϒformats_dict    λ.Object
						ϒget_height      λ.Object
						ϒinitials        λ.Object
						ϒlike_count      λ.Object
						ϒmobj            λ.Object
						ϒquality         λ.Object
						ϒself            = λargs[0]
						ϒsources         λ.Object
						ϒthumbnail       λ.Object
						ϒtitle           λ.Object
						ϒupload_date     λ.Object
						ϒuploader        λ.Object
						ϒurl             = λargs[1]
						ϒurlh            λ.Object
						ϒvideo           λ.Object
						ϒvideo_id        λ.Object
						ϒvideo_url       λ.Object
						ϒview_count      λ.Object
						ϒwebpage         λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
						τmp3             λ.Object
						τmp4             λ.Object
					)
					ϒurl = λ.Cal(Ωre.ϒsub, λ.StrLiteral("^(https?://(?:.+?\\.)?)m\\."), λ.StrLiteral("\\1"), ϒurl)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = func() λ.Object {
						if λv := λ.Calm(ϒmobj, "group", λ.StrLiteral("id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒmobj, "group", λ.StrLiteral("id_2"))
						}
					}()
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id_2"))
						}
					}()
					ϒdesktop_url = λ.Cal(Ωre.ϒsub, λ.StrLiteral("^(https?://(?:.+?\\.)?)m\\."), λ.StrLiteral("\\1"), ϒurl)
					τmp0 = λ.Calm(ϒself, "_download_webpage_handle", ϒdesktop_url, ϒvideo_id)
					ϒwebpage = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒurlh = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒerror = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<div[^>]+id=[\"\\']videoClosed[\"\\'][^>]*>(.+?)</div>"),
						ϒwebpage,
						λ.StrLiteral("error"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒerror) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒerror), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒage_limit = λ.Calm(ϒself, "_rta_search", ϒwebpage)
					ϒget_height = λ.NewFunction("get_height",
						[]λ.Param{
							{Name: "s"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒs = λargs[0]
							)
							return λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("^(\\d+)[pP]"),
								ϒs,
								λ.StrLiteral("height"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
						})
					ϒinitials = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewTuple(
								λ.StrLiteral("window\\.initials\\s*=\\s*({.+?})\\s*;\\s*</script>"),
								λ.StrLiteral("window\\.initials\\s*=\\s*({.+?})\\s*;"),
							),
							ϒwebpage,
							λ.StrLiteral("initials"),
						), λ.KWArgs{
							{Name: "default", Value: λ.StrLiteral("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒinitials) {
						ϒvideo = λ.GetItem(ϒinitials, λ.StrLiteral("videoModel"))
						ϒtitle = λ.GetItem(ϒvideo, λ.StrLiteral("title"))
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(λ.GetItem(ϒvideo, λ.StrLiteral("sources")), "items"))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
							ϒformats_dict = λ.GetItem(τmp2, λ.IntLiteral(1))
							if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformats_dict, λ.DictType)) {
								continue
							}
							τmp2 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒformats_dict, "items"))
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								τmp4 = τmp3
								ϒquality = λ.GetItem(τmp4, λ.IntLiteral(0))
								ϒformat_item = λ.GetItem(τmp4, λ.IntLiteral(1))
								if λ.IsTrue(λ.Eq(ϒformat_id, λ.StrLiteral("download"))) {
									continue
									if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_item, λ.DictType)) {
										continue
									}
									ϒformat_url = λ.Calm(ϒformat_item, "get", λ.StrLiteral("link"))
									ϒfilesize = λ.Call(ϒint_or_none, λ.NewArgs(λ.Calm(ϒformat_item, "get", λ.StrLiteral("size"))), λ.KWArgs{
										{Name: "invscale", Value: λ.IntLiteral(1000000)},
									})
								} else {
									ϒformat_url = ϒformat_item
									ϒfilesize = λ.None
								}
								ϒformat_url = λ.Cal(ϒurl_or_none, ϒformat_url)
								if !λ.IsTrue(ϒformat_url) {
									continue
								}
								λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
									"format_id": λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
										ϒformat_id,
										ϒquality,
									)),
									"url":      ϒformat_url,
									"ext":      λ.Cal(ϒdetermine_ext, ϒformat_url, λ.StrLiteral("mp4")),
									"height":   λ.Cal(ϒget_height, ϒquality),
									"filesize": ϒfilesize,
									"http_headers": λ.DictLiteral(map[string]λ.Object{
										"Referer": λ.Calm(ϒurlh, "geturl"),
									}),
								}))
							}
						}
						λ.Calm(ϒself, "_sort_formats", ϒformats)
						ϒcategories_list = λ.Calm(ϒvideo, "get", λ.StrLiteral("categories"))
						if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒcategories_list, λ.ListType)) {
							ϒcategories = λ.NewList()
							τmp0 = λ.Cal(λ.BuiltinIter, ϒcategories_list)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒc = τmp1
								if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒc, λ.DictType)) {
									continue
								}
								ϒc_name = λ.Calm(ϒc, "get", λ.StrLiteral("name"))
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒc_name, ϒcompat_str)) {
									λ.Calm(ϒcategories, "append", ϒc_name)
								}
							}
						} else {
							ϒcategories = λ.None
						}
						return λ.DictLiteral(map[string]λ.Object{
							"id":          ϒvideo_id,
							"display_id":  ϒdisplay_id,
							"title":       ϒtitle,
							"description": λ.Calm(ϒvideo, "get", λ.StrLiteral("description")),
							"timestamp":   λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("created"))),
							"uploader": λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("author")), λ.StrLiteral("name"))
								}), ϒcompat_str),
							"thumbnail":  λ.Calm(ϒvideo, "get", λ.StrLiteral("thumbURL")),
							"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("duration"))),
							"view_count": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("views"))),
							"like_count": λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("rating")), λ.StrLiteral("likes"))
								}), λ.IntType)),
							"dislike_count": λ.Cal(ϒint_or_none, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("rating")), λ.StrLiteral("dislikes"))
								}), λ.IntType)),
							"comment_count": λ.Cal(ϒint_or_none, λ.Calm(ϒvideo, "get", λ.StrLiteral("views"))),
							"age_limit":     ϒage_limit,
							"categories":    ϒcategories,
							"formats":       ϒformats,
						})
					}
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.NewList(
						λ.StrLiteral("<h1[^>]*>([^<]+)</h1>"),
						λ.StrLiteral("<meta[^>]+itemprop=\".*?caption.*?\"[^>]+content=\"(.+?)\""),
						λ.StrLiteral("<title[^>]*>(.+?)(?:,\\s*[^,]*?\\s*Porn\\s*[^,]*?:\\s*xHamster[^<]*| - xHamster\\.com)</title>"),
					), ϒwebpage, λ.StrLiteral("title"))
					ϒformats = λ.NewList()
					ϒformat_urls = λ.Cal(λ.SetType)
					ϒsources = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.StrLiteral("sources\\s*:\\s*({.+?})\\s*,?\\s*\\n"),
							ϒwebpage,
							λ.StrLiteral("sources"),
						), λ.KWArgs{
							{Name: "default", Value: λ.StrLiteral("{}")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒsources, "items"))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.IntLiteral(0))
						ϒformat_url = λ.GetItem(τmp2, λ.IntLiteral(1))
						ϒformat_url = λ.Cal(ϒurl_or_none, ϒformat_url)
						if !λ.IsTrue(ϒformat_url) {
							continue
						}
						if λ.Contains(ϒformat_urls, ϒformat_url) {
							continue
						}
						λ.Calm(ϒformat_urls, "add", ϒformat_url)
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"format_id": ϒformat_id,
							"url":       ϒformat_url,
							"height":    λ.Cal(ϒget_height, ϒformat_id),
						}))
					}
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.StrLiteral("file\\s*:\\s*(?P<q>[\"'])(?P<mp4>.+?)(?P=q)"),
							λ.StrLiteral("<a\\s+href=(?P<q>[\"'])(?P<mp4>.+?)(?P=q)\\s+class=[\"']mp4Thumb"),
							λ.StrLiteral("<video[^>]+file=(?P<q>[\"'])(?P<mp4>.+?)(?P=q)[^>]*>"),
						),
						ϒwebpage,
						λ.StrLiteral("video url"),
					), λ.KWArgs{
						{Name: "group", Value: λ.StrLiteral("mp4")},
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(func() λ.Object {
						if λv := ϒvideo_url; !λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewBool(!λ.Contains(ϒformat_urls, ϒvideo_url))
						}
					}()) {
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒvideo_url,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("<span>Description: </span>([^<]+)"), ϒwebpage)
					ϒdescription = func() λ.Object {
						if λ.IsTrue(ϒmobj) {
							return λ.Calm(ϒmobj, "group", λ.IntLiteral(1))
						} else {
							return λ.None
						}
					}()
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("hint=[\"\\'](\\d{4}-\\d{2}-\\d{2}) \\d{2}:\\d{2}:\\d{2} [A-Z]{3,4}"),
						ϒwebpage,
						λ.StrLiteral("upload date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("<span[^>]+itemprop=[\"\\']author[^>]+><a[^>]+><span[^>]+>([^<]+)"),
						ϒwebpage,
						λ.StrLiteral("uploader"),
					), λ.KWArgs{
						{Name: "default", Value: λ.StrLiteral("anonymous")},
					})
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.StrLiteral("[\"']thumbUrl[\"']\\s*:\\s*(?P<q>[\"'])(?P<thumbnail>.+?)(?P=q)"),
							λ.StrLiteral("<video[^>]+\"poster\"=(?P<q>[\"'])(?P<thumbnail>.+?)(?P=q)[^>]*>"),
						),
						ϒwebpage,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "group", Value: λ.StrLiteral("thumbnail")},
					})
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.StrLiteral("<[^<]+\\bitemprop=[\"\\']duration[\"\\'][^<]+\\bcontent=[\"\\'](.+?)[\"\\']"),
							λ.StrLiteral("Runtime:\\s*</span>\\s*([\\d:]+)"),
						),
						ϒwebpage,
						λ.StrLiteral("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("content=[\"\\']User(?:View|Play)s:(\\d+)"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("hint=[\\'\"](?P<likecount>\\d+) Likes / (?P<dislikecount>\\d+) Dislikes"), ϒwebpage)
					τmp0 = func() λ.Object {
						if λ.IsTrue(ϒmobj) {
							return λ.NewTuple(
								λ.Calm(ϒmobj, "group", λ.StrLiteral("likecount")),
								λ.Calm(ϒmobj, "group", λ.StrLiteral("dislikecount")),
							)
						} else {
							return λ.NewTuple(
								λ.None,
								λ.None,
							)
						}
					}()
					ϒlike_count = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒdislike_count = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("</label>Comments \\((?P<commentcount>\\d+)\\)</div>"), ϒwebpage)
					ϒcomment_count = func() λ.Object {
						if λ.IsTrue(ϒmobj) {
							return λ.Calm(ϒmobj, "group", λ.StrLiteral("commentcount"))
						} else {
							return λ.IntLiteral(0)
						}
					}()
					ϒcategories_html = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(?s)<table.+?(<span>Categories:.+?)</table>"),
						ϒwebpage,
						λ.StrLiteral("categories"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒcategories = func() λ.Object {
						if λ.IsTrue(ϒcategories_html) {
							return λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
										var (
											ϒcategory λ.Object
											τmp0      λ.Object
											τmp1      λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<a[^>]+>(.+?)</a>"), ϒcategories_html))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒcategory = τmp1
											λgy.Yield(λ.Cal(ϒclean_html, ϒcategory))
										}
										return λ.None
									})
								})))
						} else {
							return λ.None
						}
					}()
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"display_id":    ϒdisplay_id,
						"title":         ϒtitle,
						"description":   ϒdescription,
						"upload_date":   ϒupload_date,
						"uploader":      ϒuploader,
						"thumbnail":     ϒthumbnail,
						"duration":      ϒduration,
						"view_count":    ϒview_count,
						"like_count":    λ.Cal(ϒint_or_none, ϒlike_count),
						"dislike_count": λ.Cal(ϒint_or_none, ϒdislike_count),
						"comment_count": λ.Cal(ϒint_or_none, ϒcomment_count),
						"age_limit":     ϒage_limit,
						"categories":    ϒcategories,
						"formats":       ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_DOMAINS":      XHamsterIE__DOMAINS,
				"_VALID_URL":    XHamsterIE__VALID_URL,
				"_real_extract": XHamsterIE__real_extract,
			})
		}())
		XHamsterEmbedIE = λ.Cal(λ.TypeType, λ.StrLiteral("XHamsterEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XHamsterEmbedIE__VALID_URL    λ.Object
				XHamsterEmbedIE__real_extract λ.Object
			)
			XHamsterEmbedIE__VALID_URL = λ.Mod(λ.StrLiteral("https?://(?:.+?\\.)?%s/xembed\\.php\\?video=(?P<id>\\d+)"), λ.GetAttr(XHamsterIE, "_DOMAINS", nil))
			XHamsterEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself      = λargs[0]
						ϒurl       = λargs[1]
						ϒvars      λ.Object
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒvideo_id)
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.Calm(λ.StrLiteral("href=\"(https?://xhamster\\.com/(?:movies/{0}/[^\"]*\\.html|videos/[^/]*-{0})[^\"]*)\""), "format", ϒvideo_id),
						ϒwebpage,
						λ.StrLiteral("xhamster url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if !λ.IsTrue(ϒvideo_url) {
						ϒvars = λ.Calm(ϒself, "_parse_json", λ.Calm(ϒself, "_search_regex", λ.StrLiteral("vars\\s*:\\s*({.+?})\\s*,\\s*\\n"), ϒwebpage, λ.StrLiteral("vars")), ϒvideo_id)
						ϒvideo_url = λ.Cal(ϒdict_get, ϒvars, λ.NewTuple(
							λ.StrLiteral("downloadLink"),
							λ.StrLiteral("homepageLink"),
							λ.StrLiteral("commentsLink"),
							λ.StrLiteral("shareUrl"),
						))
					}
					return λ.Calm(ϒself, "url_result", ϒvideo_url, λ.StrLiteral("XHamster"))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    XHamsterEmbedIE__VALID_URL,
				"_real_extract": XHamsterEmbedIE__real_extract,
			})
		}())
		XHamsterUserIE = λ.Cal(λ.TypeType, λ.StrLiteral("XHamsterUserIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				XHamsterUserIE__VALID_URL λ.Object
			)
			XHamsterUserIE__VALID_URL = λ.Mod(λ.StrLiteral("https?://(?:.+?\\.)?%s/users/(?P<id>[^/?#&]+)"), λ.GetAttr(XHamsterIE, "_DOMAINS", nil))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": XHamsterUserIE__VALID_URL,
			})
		}())
	})
}
