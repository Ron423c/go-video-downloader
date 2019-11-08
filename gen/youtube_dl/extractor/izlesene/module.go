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
 * izlesene/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/izlesene.py
 */

package izlesene

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                λ.Object
	IzleseneIE                   λ.Object
	ϒcompat_str                  λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
	ϒdetermine_ext               λ.Object
	ϒfloat_or_none               λ.Object
	ϒget_element_by_id           λ.Object
	ϒint_or_none                 λ.Object
	ϒparse_iso8601               λ.Object
	ϒstr_to_int                  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒstr_to_int = Ωutils.ϒstr_to_int
		IzleseneIE = λ.Cal(λ.TypeType, λ.NewStr("IzleseneIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				IzleseneIE__VALID_URL    λ.Object
				IzleseneIE__real_extract λ.Object
			)
			IzleseneIE__VALID_URL = λ.NewStr("(?x)\n        https?://(?:(?:www|m)\\.)?izlesene\\.com/\n        (?:video|embedplayer)/(?:[^/]+/)?(?P<id>[0-9]+)\n        ")
			IzleseneIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcomment_count λ.Object
						ϒdescription   λ.Object
						ϒduration      λ.Object
						ϒext           λ.Object
						ϒformats       λ.Object
						ϒheight        λ.Object
						ϒquality       λ.Object
						ϒself          = λargs[0]
						ϒsource_url    λ.Object
						ϒstream        λ.Object
						ϒthumbnail     λ.Object
						ϒtimestamp     λ.Object
						ϒtitle         λ.Object
						ϒuploader      λ.Object
						ϒurl           = λargs[1]
						ϒvideo         λ.Object
						ϒvideo_id      λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://www.izlesene.com/video/%s"), ϒvideo_id), ϒvideo_id)
					ϒvideo = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("videoObj\\s*=\\s*({.+?})\\s*;\\s*\\n"), ϒwebpage, λ.NewStr("streams")), ϒvideo_id)
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("videoTitle")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(λ.GetItem(ϒvideo, λ.NewStr("media")), λ.NewStr("level")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒstream = τmp1
						ϒsource_url = λ.Cal(λ.GetAttr(ϒstream, "get", nil), λ.NewStr("source"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒsource_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒsource_url, ϒcompat_str)))
							}
						}()) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒurl, λ.NewStr("mp4"))
						ϒquality = λ.Cal(λ.GetAttr(ϒstream, "get", nil), λ.NewStr("value"))
						ϒheight = λ.Cal(ϒint_or_none, ϒquality)
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format_id"): func() λ.Object {
								if λ.IsTrue(ϒquality) {
									return λ.Mod(λ.NewStr("%sp"), ϒquality)
								} else {
									return λ.NewStr("sd")
								}
							}(),
							λ.NewStr("url"):    λ.Cal(ϒcompat_urllib_parse_unquote, ϒsource_url),
							λ.NewStr("ext"):    ϒext,
							λ.NewStr("height"): ϒheight,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒthumbnail = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("posterURL")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_proto_relative_url", nil), λ.NewArgs(λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)), λ.KWArgs{
								{Name: "scheme", Value: λ.NewStr("http:")},
							})
						}
					}()
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("adduserUsername\\s*=\\s*'([^']+)';"),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("uploadDate"), ϒwebpage, λ.NewStr("upload date")))
					ϒduration = λ.Call(ϒfloat_or_none, λ.NewArgs(func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("duration")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.NewStr("videoduration[\"\\']?\\s*=\\s*([\"\\'])(?P<value>(?:(?!\\1).)+)\\1"),
								ϒwebpage,
								λ.NewStr("duration"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
								{Name: "group", Value: λ.NewStr("value")},
							})
						}
					}()), λ.KWArgs{
						{Name: "scale", Value: λ.NewInt(1000)},
					})
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Cal(ϒget_element_by_id, λ.NewStr("videoViewCount"), ϒwebpage))
					ϒcomment_count = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("comment_count\\s*=\\s*\\'([^\\']+)\\';"),
						ϒwebpage,
						λ.NewStr("comment_count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("description"):   ϒdescription,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("uploader_id"):   ϒuploader,
						λ.NewStr("timestamp"):     ϒtimestamp,
						λ.NewStr("duration"):      ϒduration,
						λ.NewStr("view_count"):    λ.Cal(ϒint_or_none, ϒview_count),
						λ.NewStr("comment_count"): λ.Cal(ϒint_or_none, ϒcomment_count),
						λ.NewStr("age_limit"):     λ.Cal(λ.GetAttr(ϒself, "_family_friendly_search", nil), ϒwebpage),
						λ.NewStr("formats"):       ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    IzleseneIE__VALID_URL,
				λ.NewStr("_real_extract"): IzleseneIE__real_extract,
			})
		}())
	})
}
