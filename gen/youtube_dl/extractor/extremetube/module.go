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
 * extremetube/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/extremetube.py
 */

package extremetube

import (
	Ωkeezmovies "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/keezmovies"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtremeTubeIE λ.Object
	KeezMoviesIE  λ.Object
	ϒstr_to_int   λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒstr_to_int = Ωutils.ϒstr_to_int
		KeezMoviesIE = Ωkeezmovies.KeezMoviesIE
		ExtremeTubeIE = λ.Cal(λ.TypeType, λ.NewStr("ExtremeTubeIE"), λ.NewTuple(KeezMoviesIE), func() λ.Dict {
			var (
				ExtremeTubeIE__TESTS        λ.Object
				ExtremeTubeIE__VALID_URL    λ.Object
				ExtremeTubeIE__real_extract λ.Object
			)
			ExtremeTubeIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?extremetube\\.com/(?:[^/]+/)?video/(?P<id>[^/#?&]+)")
			ExtremeTubeIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.extremetube.com/video/music-video-14-british-euro-brit-european-cumshots-swallow-652431"),
					λ.NewStr("md5"): λ.NewStr("92feaafa4b58e82f261e5419f39c60cb"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("music-video-14-british-euro-brit-european-cumshots-swallow-652431"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("Music Video 14 british euro brit european cumshots swallow"),
						λ.NewStr("uploader"):   λ.NewStr("anonim"),
						λ.NewStr("view_count"): λ.IntType,
						λ.NewStr("age_limit"):  λ.NewInt(18),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.extremetube.com/gay/video/abcde-1234"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.extremetube.com/video/latina-slut-fucked-by-fat-black-dick"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.extremetube.com/video/652431"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			ExtremeTubeIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo       λ.Object
						ϒself       = λargs[0]
						ϒuploader   λ.Object
						ϒurl        = λargs[1]
						ϒview_count λ.Object
						ϒwebpage    λ.Object
						τmp0        λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_extract_info", nil), ϒurl)
					ϒwebpage = λ.GetItem(τmp0, λ.NewInt(0))
					ϒinfo = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.GetItem(ϒinfo, λ.NewStr("title"))))) {
						λ.SetItem(ϒinfo, λ.NewStr("title"), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<h1[^>]+title=\"([^\"]+)\"[^>]*>"), ϒwebpage, λ.NewStr("title")))
					}
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("Uploaded by:\\s*</[^>]+>\\s*<a[^>]+>(.+?)</a>"),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("Views:\\s*</[^>]+>\\s*<[^>]+>([\\d,\\.]+)</"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("uploader"):   ϒuploader,
						λ.NewStr("view_count"): ϒview_count,
					}))
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        ExtremeTubeIE__TESTS,
				λ.NewStr("_VALID_URL"):    ExtremeTubeIE__VALID_URL,
				λ.NewStr("_real_extract"): ExtremeTubeIE__real_extract,
			})
		}())
	})
}
