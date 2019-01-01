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
 * faz/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/faz.py
 */

package faz

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FazIE                    λ.Object
	InfoExtractor            λ.Object
	ϒcompat_etree_fromstring λ.Object
	ϒint_or_none             λ.Object
	ϒxpath_element           λ.Object
	ϒxpath_text              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ϒxpath_element = Ωutils.ϒxpath_element
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒint_or_none = Ωutils.ϒint_or_none
		FazIE = λ.Cal(λ.TypeType, λ.NewStr("FazIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FazIE_IE_NAME       λ.Object
				FazIE__TESTS        λ.Object
				FazIE__VALID_URL    λ.Object
				FazIE__real_extract λ.Object
			)
			FazIE_IE_NAME = λ.NewStr("faz.net")
			FazIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?faz\\.net/(?:[^/]+/)*.*?-(?P<id>\\d+)\\.html")
			FazIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.faz.net/multimedia/videos/stockholm-chemie-nobelpreis-fuer-drei-amerikanische-forscher-12610585.html"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("12610585"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Stockholm: Chemie-Nobelpreis für drei amerikanische Forscher"),
						λ.NewStr("description"): λ.NewStr("md5:1453fbf9a0d041d985a47306192ea253"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.faz.net/aktuell/politik/berlin-gabriel-besteht-zerreissprobe-ueber-datenspeicherung-13659345.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.faz.net/berlin-gabriel-besteht-zerreissprobe-ueber-datenspeicherung-13659345.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.faz.net/-13659345.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.faz.net/aktuell/politik/-13659345.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.faz.net/foobarblafasel-13659345.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			FazIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcode         λ.Object
						ϒconfig       λ.Object
						ϒdescription  λ.Object
						ϒencoding     λ.Object
						ϒencoding_url λ.Object
						ϒencodings    λ.Object
						ϒf            λ.Object
						ϒformats      λ.Object
						ϒmedia        λ.Object
						ϒmobj         λ.Object
						ϒperform_url  λ.Object
						ϒpref         λ.Object
						ϒself         = λargs[0]
						ϒtbr          λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒdescription = λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)
					ϒmedia = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("data-videojs-media='([^']+)"), ϒwebpage, λ.NewStr("media"))
					if λ.IsTrue(λ.Eq(ϒmedia, λ.NewStr("extern"))) {
						ϒperform_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<iframe[^>]+?src='((?:http:)?//player\\.performgroup\\.com/eplayer/eplayer\\.html#/?[0-9a-f]{26}\\.[0-9a-z]{26})"), ϒwebpage, λ.NewStr("perform url"))
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒperform_url)
					}
					ϒconfig = λ.Cal(ϒcompat_etree_fromstring, ϒmedia)
					ϒencodings = λ.Cal(ϒxpath_element, ϒconfig, λ.NewStr("ENCODINGS"), λ.NewStr("encodings"), λ.True)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewList(
						λ.NewStr("LOW"),
						λ.NewStr("HIGH"),
						λ.NewStr("HQ"),
					)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒpref = λ.GetItem(τmp2, λ.NewInt(0))
						ϒcode = λ.GetItem(τmp2, λ.NewInt(1))
						ϒencoding = λ.Cal(ϒxpath_element, ϒencodings, ϒcode)
						if λ.IsTrue(λ.NewBool(ϒencoding != λ.None)) {
							ϒencoding_url = λ.Cal(ϒxpath_text, ϒencoding, λ.NewStr("FILENAME"))
							if λ.IsTrue(ϒencoding_url) {
								ϒtbr = λ.Cal(ϒxpath_text, ϒencoding, λ.NewStr("AVERAGEBITRATE"), λ.NewInt(1000))
								if λ.IsTrue(ϒtbr) {
									ϒtbr = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒtbr, "replace", nil), λ.NewStr(","), λ.NewStr(".")))
								}
								ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):       ϒencoding_url,
									λ.NewStr("format_id"): λ.Cal(λ.GetAttr(ϒcode, "lower", nil)),
									λ.NewStr("quality"):   ϒpref,
									λ.NewStr("tbr"):       ϒtbr,
									λ.NewStr("vcodec"):    λ.Cal(ϒxpath_text, ϒencoding, λ.NewStr("CODEC")),
								})
								ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("(\\d+)x(\\d+)_(\\d+)\\.mp4"), ϒencoding_url)
								if λ.IsTrue(ϒmobj) {
									λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("width"):  λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(1))),
										λ.NewStr("height"): λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(2))),
										λ.NewStr("tbr"): func() λ.Object {
											if λv := ϒtbr; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(3)))
											}
										}(),
									}))
								}
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):      ϒvideo_id,
						λ.NewStr("title"):   λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage),
						λ.NewStr("formats"): ϒformats,
						λ.NewStr("description"): func() λ.Object {
							if λ.IsTrue(ϒdescription) {
								return λ.Cal(λ.GetAttr(ϒdescription, "strip", nil))
							} else {
								return λ.None
							}
						}(),
						λ.NewStr("thumbnail"): λ.Cal(ϒxpath_text, ϒconfig, λ.NewStr("STILL/STILL_BIG")),
						λ.NewStr("duration"):  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒconfig, λ.NewStr("DURATION"))),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       FazIE_IE_NAME,
				λ.NewStr("_TESTS"):        FazIE__TESTS,
				λ.NewStr("_VALID_URL"):    FazIE__VALID_URL,
				λ.NewStr("_real_extract"): FazIE__real_extract,
			})
		}())
	})
}