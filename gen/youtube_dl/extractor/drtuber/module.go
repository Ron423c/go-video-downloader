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
 * drtuber/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/drtuber.py
 */

package drtuber

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DrTuberIE       λ.Object
	InfoExtractor   λ.Object
	NO_DEFAULT      λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒstr_to_int     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_to_int = Ωutils.ϒstr_to_int
		DrTuberIE = λ.Cal(λ.TypeType, λ.NewStr("DrTuberIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DrTuberIE__VALID_URL    λ.Object
				DrTuberIE__real_extract λ.Object
			)
			DrTuberIE__VALID_URL = λ.NewStr("https?://(?:(?:www|m)\\.)?drtuber\\.com/(?:video|embed)/(?P<id>\\d+)(?:/(?P<display_id>[\\w-]+))?")
			DrTuberIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcategories    λ.Object
						ϒcats_str      λ.Object
						ϒcomment_count λ.Object
						ϒdislike_count λ.Object
						ϒdisplay_id    λ.Object
						ϒduration      λ.Object
						ϒextract_count λ.Object
						ϒformat_id     λ.Object
						ϒformats       λ.Object
						ϒlike_count    λ.Object
						ϒmobj          λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_data    λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
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
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://www.drtuber.com/video/%s"), ϒvideo_id), ϒdisplay_id)
					ϒvideo_data = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("http://www.drtuber.com/player_config_json/"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("vid"):       ϒvideo_id,
							λ.NewStr("embed"):     λ.NewInt(0),
							λ.NewStr("aid"):       λ.NewInt(0),
							λ.NewStr("domain_id"): λ.NewInt(0),
						})},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒvideo_data, λ.NewStr("files")), "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒvideo_url = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(ϒvideo_url) {
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("quality"): func() λ.Object {
									if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("hq"))) {
										return λ.NewInt(2)
									} else {
										return λ.NewInt(1)
									}
								}(),
								λ.NewStr("url"): ϒvideo_url,
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒduration = func() λ.Object {
						if λv := λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration"))); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration_format")))
						}
					}()
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewTuple(
						λ.NewStr("<h1[^>]+class=[\"\\']title[^>]+>([^<]+)"),
						λ.NewStr("<title>([^<]+)\\s*@\\s+DrTuber"),
						λ.NewStr("class=\"title_watch\"[^>]*><(?:p|h\\d+)[^>]*>([^<]+)<"),
						λ.NewStr("<p[^>]+class=\"title_substrate\">([^<]+)</p>"),
						λ.NewStr("<title>([^<]+) - \\d+"),
					), ϒwebpage, λ.NewStr("title"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("poster=\"([^\"]+)\""),
						ϒwebpage,
						λ.NewStr("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒextract_count = λ.NewFunction("extract_count",
						[]λ.Param{
							{Name: "id_"},
							{Name: "name"},
							{Name: "default", Def: NO_DEFAULT},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdefault = λargs[2]
								ϒid_     = λargs[0]
								ϒname    = λargs[1]
							)
							return λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
								λ.Mod(λ.NewStr("<span[^>]+(?:class|id)=\"%s\"[^>]*>([\\d,\\.]+)</span>"), ϒid_),
								ϒwebpage,
								λ.Mod(λ.NewStr("%s count"), ϒname),
							), λ.KWArgs{
								{Name: "default", Value: ϒdefault},
								{Name: "fatal", Value: λ.False},
							}))
						})
					ϒlike_count = λ.Cal(ϒextract_count, λ.NewStr("rate_likes"), λ.NewStr("like"))
					ϒdislike_count = λ.Call(ϒextract_count, λ.NewArgs(
						λ.NewStr("rate_dislikes"),
						λ.NewStr("dislike"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒcomment_count = λ.Cal(ϒextract_count, λ.NewStr("comments_count"), λ.NewStr("comment"))
					ϒcats_str = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<div[^>]+class=\"categories_list\">(.+?)</div>"),
						ϒwebpage,
						λ.NewStr("categories"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒcategories = func() λ.Object {
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒcats_str))) {
							return λ.NewList()
						} else {
							return λ.Cal(Ωre.ϒfindall, λ.NewStr("<a title=\"([^\"]+)\""), ϒcats_str)
						}
					}()
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("display_id"):    ϒdisplay_id,
						λ.NewStr("formats"):       ϒformats,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("thumbnail"):     ϒthumbnail,
						λ.NewStr("like_count"):    ϒlike_count,
						λ.NewStr("dislike_count"): ϒdislike_count,
						λ.NewStr("comment_count"): ϒcomment_count,
						λ.NewStr("categories"):    ϒcategories,
						λ.NewStr("age_limit"):     λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage),
						λ.NewStr("duration"):      ϒduration,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    DrTuberIE__VALID_URL,
				λ.NewStr("_real_extract"): DrTuberIE__real_extract,
			})
		}())
	})
}
