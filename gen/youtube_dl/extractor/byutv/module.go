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
 * byutv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/byutv.py
 */

package byutv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BYUtvIE       λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		BYUtvIE = λ.Cal(λ.TypeType, λ.NewStr("BYUtvIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BYUtvIE__TESTS        λ.Object
				BYUtvIE__VALID_URL    λ.Object
				BYUtvIE__real_extract λ.Object
			)
			BYUtvIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?byutv\\.org/(?:watch|player)/(?!event/)(?P<id>[0-9a-f-]+)(?:/(?P<display_id>[^/?#&]+))?")
			BYUtvIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.byutv.org/watch/6587b9a3-89d2-42a6-a7f7-fd2f81840a7d/studio-c-season-5-episode-5"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("ZvanRocTpW-G5_yZFeltTAMv6jxOU9KH"),
						λ.NewStr("display_id"):  λ.NewStr("studio-c-season-5-episode-5"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Season 5 Episode 5"),
						λ.NewStr("description"): λ.NewStr("md5:1d31dc18ef4f075b28f6a65937d22c65"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*"),
						λ.NewStr("duration"):    λ.NewFloat(1486.486),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Ooyala")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.byutv.org/watch/6587b9a3-89d2-42a6-a7f7-fd2f81840a7d"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.byutv.org/player/27741493-dc83-40b0-8420-e7ae38a2ae98/byu-football-toledo-vs-byu-93016?listid=4fe0fee5-0d3c-4a29-b725-e4948627f472&listindex=0&q=toledo"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			BYUtvIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒep         λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
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
					ϒep = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("https://api.byutv.org/api3/catalog/getvideosforcontent"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("contentid"):       ϒvideo_id,
							λ.NewStr("channel"):         λ.NewStr("byutv"),
							λ.NewStr("x-byutv-context"): λ.NewStr("web$US"),
						})},
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("x-byutv-context"):     λ.NewStr("web$US"),
							λ.NewStr("x-byutv-platformkey"): λ.NewStr("xsaaw9c7y5"),
						})},
					}), λ.NewStr("ooyalaVOD"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):       λ.NewStr("url_transparent"),
						λ.NewStr("ie_key"):      λ.NewStr("Ooyala"),
						λ.NewStr("url"):         λ.Mod(λ.NewStr("ooyala:%s"), λ.GetItem(ϒep, λ.NewStr("providerId"))),
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       λ.Cal(λ.GetAttr(ϒep, "get", nil), λ.NewStr("title")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒep, "get", nil), λ.NewStr("description")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒep, "get", nil), λ.NewStr("imageThumbnail")),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        BYUtvIE__TESTS,
				λ.NewStr("_VALID_URL"):    BYUtvIE__VALID_URL,
				λ.NewStr("_real_extract"): BYUtvIE__real_extract,
			})
		}())
	})
}
