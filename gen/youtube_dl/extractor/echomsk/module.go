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
 * echomsk/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/echomsk.py
 */

package echomsk

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	EchoMskIE     λ.Object
	InfoExtractor λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		EchoMskIE = λ.Cal(λ.TypeType, λ.NewStr("EchoMskIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				EchoMskIE__TEST         λ.Object
				EchoMskIE__VALID_URL    λ.Object
				EchoMskIE__real_extract λ.Object
			)
			EchoMskIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?echo\\.msk\\.ru/sounds/(?P<id>\\d+)")
			EchoMskIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.echo.msk.ru/sounds/1464134.html"),
				λ.NewStr("md5"): λ.NewStr("2e44b3b78daff5b458e4dbc37f191f7c"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):    λ.NewStr("1464134"),
					λ.NewStr("ext"):   λ.NewStr("mp3"),
					λ.NewStr("title"): λ.NewStr("Особое мнение - 29 декабря 2014, 19:08"),
				}),
			})
			EchoMskIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒair_date  λ.Object
						ϒaudio_url λ.Object
						ϒself      = λargs[0]
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo_id  λ.Object
						ϒwebpage   λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒaudio_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<a rel=\"mp3\" href=\"([^\"]+)\">"), ϒwebpage, λ.NewStr("audio URL"))
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<a href=\"/programs/[^\"]+\" target=\"_blank\">([^<]+)</a>"), ϒwebpage, λ.NewStr("title"))
					ϒair_date = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)<div class=\"date\">(.+?)</div>"),
						ϒwebpage,
						λ.NewStr("date"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒair_date) {
						ϒair_date = λ.Cal(Ωre.ϒsub, λ.NewStr("(\\s)\\1+"), λ.NewStr("\\1"), ϒair_date)
						if λ.IsTrue(ϒair_date) {
							ϒtitle = λ.Mod(λ.NewStr("%s - %s"), λ.NewTuple(
								ϒtitle,
								ϒair_date,
							))
						}
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("url"):   ϒaudio_url,
						λ.NewStr("title"): ϒtitle,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         EchoMskIE__TEST,
				λ.NewStr("_VALID_URL"):    EchoMskIE__VALID_URL,
				λ.NewStr("_real_extract"): EchoMskIE__real_extract,
			})
		}())
	})
}
