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
 * sunporno/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sunporno.py
 */

package sunporno

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	SunPornoIE      λ.Object
	ϒdetermine_ext  λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒqualities      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		SunPornoIE = λ.Cal(λ.TypeType, λ.StrLiteral("SunPornoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SunPornoIE__VALID_URL    λ.Object
				SunPornoIE__real_extract λ.Object
			)
			SunPornoIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www\\.)?sunporno\\.com/videos|embeds\\.sunporno\\.com/embed)/(?P<id>\\d+)")
			SunPornoIE__real_extract = λ.NewFunction("_real_extract",
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
						ϒformats       λ.Object
						ϒquality       λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_ext     λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", λ.Mod(λ.StrLiteral("http://www.sunporno.com/videos/%s"), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.Calm(ϒself, "_html_search_regex", λ.StrLiteral("<title>([^<]+)</title>"), ϒwebpage, λ.StrLiteral("title"))
					ϒdescription = λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("description"), ϒwebpage, λ.StrLiteral("description"))
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("poster=\"([^\"]+)\""),
						ϒwebpage,
						λ.StrLiteral("thumbnail"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.StrLiteral("itemprop=\"duration\"[^>]*>\\s*(\\d+:\\d+)\\s*<"),
							λ.StrLiteral(">Duration:\\s*<span[^>]+>\\s*(\\d+:\\d+)\\s*<"),
						),
						ϒwebpage,
						λ.StrLiteral("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒview_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("class=\"views\">(?:<noscript>)?\\s*(\\d+)\\s*<"),
						ϒwebpage,
						λ.StrLiteral("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒcomment_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(\\d+)</b> Comments?"),
						ϒwebpage,
						λ.StrLiteral("comment count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "default", Value: λ.None},
					}))
					ϒformats = λ.NewList()
					ϒquality = λ.Cal(ϒqualities, λ.NewList(
						λ.StrLiteral("mp4"),
						λ.StrLiteral("flv"),
					))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.StrLiteral("<(?:source|video) src=\"([^\"]+)\""), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_url = τmp1
						ϒvideo_ext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒvideo_url,
							"format_id": ϒvideo_ext,
							"quality":   λ.Cal(ϒquality, ϒvideo_ext),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":            ϒvideo_id,
						"title":         ϒtitle,
						"description":   ϒdescription,
						"thumbnail":     ϒthumbnail,
						"duration":      ϒduration,
						"view_count":    ϒview_count,
						"comment_count": ϒcomment_count,
						"formats":       ϒformats,
						"age_limit":     λ.IntLiteral(18),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    SunPornoIE__VALID_URL,
				"_real_extract": SunPornoIE__real_extract,
			})
		}())
	})
}