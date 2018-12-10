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
 * tvanouvelles/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/tvanouvelles.py
 */

package tvanouvelles

import (
	Ωbrightcove "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/brightcove"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BrightcoveNewIE       λ.Object
	InfoExtractor         λ.Object
	TVANouvellesArticleIE λ.Object
	TVANouvellesIE        λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		BrightcoveNewIE = Ωbrightcove.BrightcoveNewIE
		TVANouvellesIE = λ.Cal(λ.TypeType, λ.NewStr("TVANouvellesIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVANouvellesIE_BRIGHTCOVE_URL_TEMPLATE λ.Object
				TVANouvellesIE__TEST                   λ.Object
				TVANouvellesIE__VALID_URL              λ.Object
				TVANouvellesIE__real_extract           λ.Object
			)
			TVANouvellesIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tvanouvelles\\.ca/videos/(?P<id>\\d+)")
			TVANouvellesIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.tvanouvelles.ca/videos/5117035533001"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("5117035533001"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("L’industrie du taxi dénonce l’entente entre Québec et Uber: explications"),
					λ.NewStr("description"): λ.NewStr("md5:479653b7c8cf115747bf5118066bd8b3"),
					λ.NewStr("uploader_id"): λ.NewStr("1741764581"),
					λ.NewStr("timestamp"):   λ.NewInt(1473352030),
					λ.NewStr("upload_date"): λ.NewStr("20160908"),
				}),
				λ.NewStr("add_ie"): λ.NewList(λ.NewStr("BrightcoveNew")),
			})
			TVANouvellesIE_BRIGHTCOVE_URL_TEMPLATE = λ.NewStr("http://players.brightcove.net/1741764581/default_default/index.html?videoId=%s")
			TVANouvellesIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbrightcove_id λ.Object
						ϒself          = λargs[0]
						ϒurl           = λargs[1]
					)
					ϒbrightcove_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.GetAttr(ϒself, "BRIGHTCOVE_URL_TEMPLATE", nil), ϒbrightcove_id), λ.Cal(λ.GetAttr(BrightcoveNewIE, "ie_key", nil)), ϒbrightcove_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("BRIGHTCOVE_URL_TEMPLATE"): TVANouvellesIE_BRIGHTCOVE_URL_TEMPLATE,
				λ.NewStr("_TEST"):                   TVANouvellesIE__TEST,
				λ.NewStr("_VALID_URL"):              TVANouvellesIE__VALID_URL,
				λ.NewStr("_real_extract"):           TVANouvellesIE__real_extract,
			})
		}())
		TVANouvellesArticleIE = λ.Cal(λ.TypeType, λ.NewStr("TVANouvellesArticleIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TVANouvellesArticleIE__VALID_URL λ.Object
				TVANouvellesArticleIE_suitable   λ.Object
			)
			TVANouvellesArticleIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?tvanouvelles\\.ca/(?:[^/]+/)+(?P<id>[^/?#&]+)")
			TVANouvellesArticleIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Cal(λ.GetAttr(TVANouvellesIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, TVANouvellesArticleIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			TVANouvellesArticleIE_suitable = λ.Cal(λ.ClassMethodType, TVANouvellesArticleIE_suitable)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): TVANouvellesArticleIE__VALID_URL,
				λ.NewStr("suitable"):   TVANouvellesArticleIE_suitable,
			})
		}())
	})
}
