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
 * onionstudios/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/onionstudios.py
 */

package onionstudios

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	OnionStudiosIE λ.Object
	ϒcompat_str    λ.Object
	ϒjs_to_json    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒjs_to_json = Ωutils.ϒjs_to_json
		OnionStudiosIE = λ.Cal(λ.TypeType, λ.NewStr("OnionStudiosIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				OnionStudiosIE__TESTS        λ.Object
				OnionStudiosIE__VALID_URL    λ.Object
				OnionStudiosIE__real_extract λ.Object
			)
			OnionStudiosIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?onionstudios\\.com/(?:video(?:s/[^/]+-|/)|embed\\?.*\\bid=)(?P<id>\\d+)(?!-)")
			OnionStudiosIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.onionstudios.com/videos/hannibal-charges-forward-stops-for-a-cocktail-2937"),
					λ.NewStr("md5"): λ.NewStr("5a118d466d62b5cd03647cf2c593977f"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("3459881"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Hannibal charges forward, stops for a cocktail"),
						λ.NewStr("description"): λ.NewStr("md5:545299bda6abf87e5ec666548c6a9448"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("uploader"):    λ.NewStr("a.v. club"),
						λ.NewStr("upload_date"): λ.NewStr("20150619"),
						λ.NewStr("timestamp"):   λ.NewInt(1434728546),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.onionstudios.com/embed?id=2855&autoplay=true"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.onionstudios.com/video/6139.json"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			OnionStudiosIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmcp_id   λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewStr("http://onionstudios.com/embed/dc94dc2899fe644c0e7241fa04c1b732.js"), ϒvideo_id)
					ϒmcp_id = λ.Cal(ϒcompat_str, λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("window\\.mcpMapping\\s*=\\s*({.+?});"), ϒwebpage, λ.NewStr("MCP Mapping")), ϒvideo_id, ϒjs_to_json), ϒvideo_id), λ.NewStr("mcp_id")))
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Add(λ.NewStr("http://kinja.com/ajax/inset/iframe?id=mcp-"), ϒmcp_id), λ.NewStr("KinjaEmbed"), ϒmcp_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        OnionStudiosIE__TESTS,
				λ.NewStr("_VALID_URL"):    OnionStudiosIE__VALID_URL,
				λ.NewStr("_real_extract"): OnionStudiosIE__real_extract,
			})
		}())
	})
}
