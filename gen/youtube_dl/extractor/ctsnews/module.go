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
 * ctsnews/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ctsnews.py
 */

package ctsnews

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CtsNewsIE          λ.Object
	InfoExtractor      λ.Object
	YoutubeIE          λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		YoutubeIE = Ωyoutube.YoutubeIE
		CtsNewsIE = λ.Cal(λ.TypeType, λ.StrLiteral("CtsNewsIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CtsNewsIE__VALID_URL    λ.Object
				CtsNewsIE__real_extract λ.Object
			)
			CtsNewsIE__VALID_URL = λ.StrLiteral("https?://news\\.cts\\.com\\.tw/[a-z]+/[a-z]+/\\d+/(?P<id>\\d+)\\.html")
			CtsNewsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdatetime_str λ.Object
						ϒdescription  λ.Object
						ϒmp4_feed     λ.Object
						ϒnews_id      λ.Object
						ϒpage         λ.Object
						ϒself         = λargs[0]
						ϒthumbnail    λ.Object
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo_url    λ.Object
						ϒyoutube_url  λ.Object
					)
					ϒnews_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒnews_id)
					ϒnews_id = λ.Calm(λ.Calm(ϒself, "_hidden_inputs", ϒpage), "get", λ.StrLiteral("get_id"))
					if λ.IsTrue(ϒnews_id) {
						ϒmp4_feed = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.StrLiteral("http://news.cts.com.tw/action/test_mp4feed.php"),
							ϒnews_id,
						), λ.KWArgs{
							{Name: "note", Value: λ.StrLiteral("Fetching feed")},
							{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
								"news_id": ϒnews_id,
							})},
						})
						ϒvideo_url = λ.GetItem(ϒmp4_feed, λ.StrLiteral("source_url"))
					} else {
						λ.Calm(ϒself, "to_screen", λ.StrLiteral("Not CTSPlayer video, trying Youtube..."))
						ϒyoutube_url = λ.Calm(YoutubeIE, "_extract_url", ϒpage)
						return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒyoutube_url), λ.KWArgs{
							{Name: "ie", Value: λ.StrLiteral("Youtube")},
						})
					}
					ϒdescription = λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("description"), ϒpage)
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.StrLiteral("title"),
						ϒpage,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒthumbnail = λ.Calm(ϒself, "_html_search_meta", λ.StrLiteral("image"), ϒpage)
					ϒdatetime_str = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.StrLiteral("(\\d{4}/\\d{2}/\\d{2} \\d{2}:\\d{2})"),
						ϒpage,
						λ.StrLiteral("date and time"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒtimestamp = λ.None
					if λ.IsTrue(ϒdatetime_str) {
						ϒtimestamp = λ.Sub(λ.Cal(ϒunified_timestamp, ϒdatetime_str), λ.Mul(λ.IntLiteral(8), λ.IntLiteral(3600)))
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒnews_id,
						"url":         ϒvideo_url,
						"title":       ϒtitle,
						"description": ϒdescription,
						"thumbnail":   ϒthumbnail,
						"timestamp":   ϒtimestamp,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    CtsNewsIE__VALID_URL,
				"_real_extract": CtsNewsIE__real_extract,
			})
		}())
	})
}
