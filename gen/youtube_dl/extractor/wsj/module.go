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
 * wsj/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/wsj.py
 */

package wsj

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	WSJArticleIE     λ.Object
	WSJIE            λ.Object
	ϒfloat_or_none   λ.Object
	ϒint_or_none     λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒunified_strdate = Ωutils.ϒunified_strdate
		WSJIE = λ.Cal(λ.TypeType, λ.StrLiteral("WSJIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WSJIE__VALID_URL    λ.Object
				WSJIE__real_extract λ.Object
			)
			WSJIE__VALID_URL = λ.StrLiteral("(?x)\n                        (?:\n                            https?://video-api\\.wsj\\.com/api-video/player/iframe\\.html\\?.*?\\bguid=|\n                            https?://(?:www\\.)?(?:wsj|barrons)\\.com/video/(?:[^/]+/)+|\n                            wsj:\n                        )\n                        (?P<id>[a-fA-F0-9-]{36})\n                    ")
			WSJIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒf4m_url  λ.Object
						ϒformats  λ.Object
						ϒinfo     λ.Object
						ϒm3u8_url λ.Object
						ϒmp4_url  λ.Object
						ϒself     = λargs[0]
						ϒtbr      λ.Object
						ϒtitle    λ.Object
						ϒurl      = λargs[1]
						ϒv        λ.Object
						ϒvideo_id λ.Object
						τmp0      λ.Object
						τmp1      λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒinfo = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("http://video-api.wsj.com/api-video/find_all_videos.asp"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"type":  λ.StrLiteral("guid"),
							"count": λ.IntLiteral(1),
							"query": ϒvideo_id,
							"fields": λ.Calm(λ.StrLiteral(","), "join", λ.NewTuple(
								λ.StrLiteral("type"),
								λ.StrLiteral("hls"),
								λ.StrLiteral("videoMP4List"),
								λ.StrLiteral("thumbnailList"),
								λ.StrLiteral("author"),
								λ.StrLiteral("description"),
								λ.StrLiteral("name"),
								λ.StrLiteral("duration"),
								λ.StrLiteral("videoURL"),
								λ.StrLiteral("titletag"),
								λ.StrLiteral("formattedCreationDate"),
								λ.StrLiteral("keywords"),
								λ.StrLiteral("editor"),
							)),
						})},
					}), λ.StrLiteral("items")), λ.IntLiteral(0))
					ϒtitle = λ.Calm(ϒinfo, "get", λ.StrLiteral("name"), λ.Calm(ϒinfo, "get", λ.StrLiteral("titletag")))
					ϒformats = λ.NewList()
					ϒf4m_url = λ.Calm(ϒinfo, "get", λ.StrLiteral("videoURL"))
					if λ.IsTrue(ϒf4m_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							ϒf4m_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.StrLiteral("hds")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒm3u8_url = λ.Calm(ϒinfo, "get", λ.StrLiteral("hls"))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							λ.GetItem(ϒinfo, λ.StrLiteral("hls")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "ext", Value: λ.StrLiteral("mp4")},
							{Name: "entry_protocol", Value: λ.StrLiteral("m3u8_native")},
							{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒinfo, "get", λ.StrLiteral("videoMP4List"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒv = τmp1
						ϒmp4_url = λ.Calm(ϒv, "get", λ.StrLiteral("url"))
						if !λ.IsTrue(ϒmp4_url) {
							continue
						}
						ϒtbr = λ.Cal(ϒint_or_none, λ.Calm(ϒv, "get", λ.StrLiteral("bitrate")))
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url": ϒmp4_url,
							"format_id": λ.Add(λ.StrLiteral("http"), func() λ.Object {
								if λ.IsTrue(ϒtbr) {
									return λ.Mod(λ.StrLiteral("-%d"), ϒtbr)
								} else {
									return λ.StrLiteral("")
								}
							}()),
							"tbr":    ϒtbr,
							"width":  λ.Cal(ϒint_or_none, λ.Calm(ϒv, "get", λ.StrLiteral("width"))),
							"height": λ.Cal(ϒint_or_none, λ.Calm(ϒv, "get", λ.StrLiteral("height"))),
							"fps":    λ.Cal(ϒfloat_or_none, λ.Calm(ϒv, "get", λ.StrLiteral("fps"))),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒvideo_id,
						"formats":     ϒformats,
						"thumbnails":  λ.Calm(ϒinfo, "get", λ.StrLiteral("thumbnailList")),
						"creator":     λ.Calm(ϒinfo, "get", λ.StrLiteral("author")),
						"uploader_id": λ.Calm(ϒinfo, "get", λ.StrLiteral("editor")),
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒinfo, "get", λ.StrLiteral("duration"))),
						"upload_date": λ.Call(ϒunified_strdate, λ.NewArgs(λ.Calm(ϒinfo, "get", λ.StrLiteral("formattedCreationDate"))), λ.KWArgs{
							{Name: "day_first", Value: λ.False},
						}),
						"title":      ϒtitle,
						"categories": λ.Calm(ϒinfo, "get", λ.StrLiteral("keywords")),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    WSJIE__VALID_URL,
				"_real_extract": WSJIE__real_extract,
			})
		}())
		WSJArticleIE = λ.Cal(λ.TypeType, λ.StrLiteral("WSJArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				WSJArticleIE__VALID_URL    λ.Object
				WSJArticleIE__real_extract λ.Object
			)
			WSJArticleIE__VALID_URL = λ.StrLiteral("(?i)https?://(?:www\\.)?wsj\\.com/articles/(?P<id>[^/?#&]+)")
			WSJArticleIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒarticle_id λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
					)
					ϒarticle_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒwebpage = λ.Calm(ϒself, "_download_webpage", ϒurl, ϒarticle_id)
					ϒvideo_id = λ.Calm(ϒself, "_search_regex", λ.StrLiteral("data-src=[\"\\']([a-fA-F0-9-]{36})"), ϒwebpage, λ.StrLiteral("video id"))
					return λ.Calm(ϒself, "url_result", λ.Mod(λ.StrLiteral("wsj:%s"), ϒvideo_id), λ.Calm(WSJIE, "ie_key"), ϒvideo_id)
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    WSJArticleIE__VALID_URL,
				"_real_extract": WSJArticleIE__real_extract,
			})
		}())
	})
}
