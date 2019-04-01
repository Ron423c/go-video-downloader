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
 * cbssports/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/cbssports.py
 */

package cbssports

import (
	Ωcbs "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/cbs"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CBSBaseIE   λ.Object
	CBSSportsIE λ.Object
)

func init() {
	λ.InitModule(func() {
		CBSBaseIE = Ωcbs.CBSBaseIE
		CBSSportsIE = λ.Cal(λ.TypeType, λ.NewStr("CBSSportsIE"), λ.NewTuple(CBSBaseIE), func() λ.Dict {
			var (
				CBSSportsIE__TESTS              λ.Object
				CBSSportsIE__VALID_URL          λ.Object
				CBSSportsIE__extract_video_info λ.Object
				CBSSportsIE__real_extract       λ.Object
			)
			CBSSportsIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?cbssports\\.com/[^/]+/(?:video|news)/(?P<id>[^/?#&]+)")
			CBSSportsIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.cbssports.com/nba/video/donovan-mitchell-flashes-star-potential-in-game-2-victory-over-thunder/"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1214315075735"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Donovan Mitchell flashes star potential in Game 2 victory over Thunder"),
						λ.NewStr("description"): λ.NewStr("md5:df6f48622612c2d6bd2e295ddef58def"),
						λ.NewStr("timestamp"):   λ.NewInt(1524111457),
						λ.NewStr("upload_date"): λ.NewStr("20180419"),
						λ.NewStr("uploader"):    λ.NewStr("CBSI-NEW"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.cbssports.com/nba/news/nba-playoffs-2018-watch-76ers-vs-heat-game-3-series-schedule-tv-channel-online-stream/"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			CBSSportsIE__extract_video_info = λ.NewFunction("_extract_video_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "filter_query"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfilter_query = λargs[1]
						ϒself         = λargs[0]
						ϒvideo_id     = λargs[2]
					)
					return λ.Cal(λ.GetAttr(ϒself, "_extract_feed_info", nil), λ.NewStr("dJ5BDC"), λ.NewStr("VxxJg8Ymh8sE"), ϒfilter_query, ϒvideo_id)
				})
			CBSSportsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewList(
						λ.NewStr("(?:=|%26)pcid%3D(\\d+)"),
						λ.NewStr("embedVideo(?:Container)?_(\\d+)"),
					), ϒwebpage, λ.NewStr("video id"))
					return λ.Cal(λ.GetAttr(ϒself, "_extract_video_info", nil), λ.Mod(λ.NewStr("byId=%s"), ϒvideo_id), ϒvideo_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):              CBSSportsIE__TESTS,
				λ.NewStr("_VALID_URL"):          CBSSportsIE__VALID_URL,
				λ.NewStr("_extract_video_info"): CBSSportsIE__extract_video_info,
				λ.NewStr("_real_extract"):       CBSSportsIE__real_extract,
			})
		}())
	})
}