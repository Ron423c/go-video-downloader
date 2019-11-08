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
 * ccma/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ccma.py
 */

package ccma

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CCMAIE            λ.Object
	InfoExtractor     λ.Object
	ϒclean_html       λ.Object
	ϒint_or_none      λ.Object
	ϒparse_duration   λ.Object
	ϒparse_iso8601    λ.Object
	ϒparse_resolution λ.Object
	ϒurl_or_none      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒparse_resolution = Ωutils.ϒparse_resolution
		ϒurl_or_none = Ωutils.ϒurl_or_none
		CCMAIE = λ.Cal(λ.TypeType, λ.NewStr("CCMAIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CCMAIE__VALID_URL    λ.Object
				CCMAIE__real_extract λ.Object
			)
			CCMAIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?ccma\\.cat/(?:[^/]+/)*?(?P<type>video|audio)/(?P<id>\\d+)")
			CCMAIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdurada        λ.Object
						ϒduration      λ.Object
						ϒf             λ.Object
						ϒformat_       λ.Object
						ϒformat_url    λ.Object
						ϒformats       λ.Object
						ϒimatges       λ.Object
						ϒinformacio    λ.Object
						ϒlabel         λ.Object
						ϒmedia         λ.Object
						ϒmedia_id      λ.Object
						ϒmedia_type    λ.Object
						ϒmedia_url     λ.Object
						ϒself          = λargs[0]
						ϒsub_url       λ.Object
						ϒsubtitles     λ.Object
						ϒsubtitols     λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒtimestamp     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒmedia_type = λ.GetItem(τmp0, λ.NewInt(0))
					ϒmedia_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒmedia = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("http://dinamics.ccma.cat/pvideo/media.jsp"),
						ϒmedia_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("media"): ϒmedia_type,
							λ.NewStr("idint"): ϒmedia_id,
						})},
					})
					ϒformats = λ.NewList()
					ϒmedia_url = λ.GetItem(λ.GetItem(ϒmedia, λ.NewStr("media")), λ.NewStr("url"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒmedia_url, λ.ListType)) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒmedia_url)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒformat_ = τmp1
							ϒformat_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒformat_, "get", nil), λ.NewStr("file")))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformat_url))) {
								continue
							}
							ϒlabel = λ.Cal(λ.GetAttr(ϒformat_, "get", nil), λ.NewStr("label"))
							ϒf = λ.Cal(ϒparse_resolution, ϒlabel)
							λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒformat_url,
								λ.NewStr("format_id"): ϒlabel,
							}))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
					} else {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒmedia_url,
							λ.NewStr("vcodec"): func() λ.Object {
								if λ.IsTrue(λ.Eq(ϒmedia_type, λ.NewStr("audio"))) {
									return λ.NewStr("none")
								} else {
									return λ.None
								}
							}(),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒinformacio = λ.GetItem(ϒmedia, λ.NewStr("informacio"))
					ϒtitle = λ.GetItem(ϒinformacio, λ.NewStr("titol"))
					ϒdurada = λ.Cal(λ.GetAttr(ϒinformacio, "get", nil), λ.NewStr("durada"), λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					ϒduration = func() λ.Object {
						if λv := λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒdurada, "get", nil), λ.NewStr("milisegons")), λ.NewInt(1000)); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒdurada, "get", nil), λ.NewStr("text")))
						}
					}()
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒinformacio, "get", nil), λ.NewStr("data_emissio"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("utc")))
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒsubtitols = λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("subtitols"), λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					if λ.IsTrue(ϒsubtitols) {
						ϒsub_url = λ.Cal(λ.GetAttr(ϒsubtitols, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(ϒsub_url) {
							λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒsubtitols, "get", nil), λ.NewStr("iso")); λ.IsTrue(λv) {
									return λv
								} else if λv := λ.Cal(λ.GetAttr(ϒsubtitols, "get", nil), λ.NewStr("text")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewStr("ca")
								}
							}(), λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒsub_url,
							}))
						}
					}
					ϒthumbnails = λ.NewList()
					ϒimatges = λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("imatges"), λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					if λ.IsTrue(ϒimatges) {
						ϒthumbnail_url = λ.Cal(λ.GetAttr(ϒimatges, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(ϒthumbnail_url) {
							ϒthumbnails = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):    ϒthumbnail_url,
								λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒimatges, "get", nil), λ.NewStr("amplada"))),
								λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒimatges, "get", nil), λ.NewStr("alcada"))),
							}))
						}
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒmedia_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒinformacio, "get", nil), λ.NewStr("descripcio"))),
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("thumbnails"):  ϒthumbnails,
						λ.NewStr("subtitles"):   ϒsubtitles,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    CCMAIE__VALID_URL,
				λ.NewStr("_real_extract"): CCMAIE__real_extract,
			})
		}())
	})
}
