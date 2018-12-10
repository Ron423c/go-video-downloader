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
 * digg/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/digg.py
 */

package digg

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DiggIE        λ.Object
	InfoExtractor λ.Object
	ϒjs_to_json   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒjs_to_json = Ωutils.ϒjs_to_json
		DiggIE = λ.Cal(λ.TypeType, λ.NewStr("DiggIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DiggIE__TESTS        λ.Object
				DiggIE__VALID_URL    λ.Object
				DiggIE__real_extract λ.Object
			)
			DiggIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?digg\\.com/video/(?P<id>[^/?#&]+)")
			DiggIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://digg.com/video/sci-fi-short-jonah-daniel-kaluuya-get-out"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("LcqvmS0b"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("'Get Out' Star Daniel Kaluuya Goes On 'Moby Dick'-Like Journey In Sci-Fi Short 'Jonah'"),
						λ.NewStr("description"): λ.NewStr("md5:541bb847648b6ee3d6514bc84b82efda"),
						λ.NewStr("upload_date"): λ.NewStr("20180109"),
						λ.NewStr("timestamp"):   λ.NewInt(1515530551),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://digg.com/video/dog-boat-seal-play"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://digg.com/video/dream-girl-short-film"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			DiggIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒinfo       λ.Object
						ϒprovider   λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒwebpage    λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒinfo = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("(?s)video_info\\s*=\\s*({.+?});\\n"),
							ϒwebpage,
							λ.NewStr("video info"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("{}")},
						}),
						ϒdisplay_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
						{Name: "fatal", Value: λ.False},
					})
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("video_id"))
					if λ.IsTrue(ϒvideo_id) {
						ϒprovider = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("provider_name"))
						if λ.IsTrue(λ.Eq(ϒprovider, λ.NewStr("youtube"))) {
							return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒvideo_id), λ.KWArgs{
								{Name: "ie", Value: λ.NewStr("Youtube")},
								{Name: "video_id", Value: ϒvideo_id},
							})
						} else {
							if λ.IsTrue(λ.Eq(ϒprovider, λ.NewStr("jwplayer"))) {
								return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Mod(λ.NewStr("jwplatform:%s"), ϒvideo_id)), λ.KWArgs{
									{Name: "ie", Value: λ.NewStr("JWPlatform")},
									{Name: "video_id", Value: ϒvideo_id},
								})
							}
						}
					}
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒurl, λ.NewStr("Generic"))
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        DiggIE__TESTS,
				λ.NewStr("_VALID_URL"):    DiggIE__VALID_URL,
				λ.NewStr("_real_extract"): DiggIE__real_extract,
			})
		}())
	})
}
