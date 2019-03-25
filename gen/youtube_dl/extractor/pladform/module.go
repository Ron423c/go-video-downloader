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
 * pladform/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/pladform.py
 */

package pladform

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError λ.Object
	InfoExtractor  λ.Object
	PladformIE     λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒqualities     λ.Object
	ϒxpath_text    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒqualities = Ωutils.ϒqualities
		PladformIE = λ.Cal(λ.TypeType, λ.NewStr("PladformIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PladformIE__VALID_URL   λ.Object
				PladformIE__extract_url λ.Object
			)
			PladformIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:\n                            (?:\n                                out\\.pladform\\.ru/player|\n                                static\\.pladform\\.ru/player\\.swf\n                            )\n                            \\?.*\\bvideoid=|\n                            video\\.pladform\\.ru/catalog/video/videoid/\n                        )\n                        (?P<id>\\d+)\n                    ")
			PladformIE__extract_url = λ.NewFunction("_extract_url",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmobj    λ.Object
						ϒwebpage = λargs[0]
					)
					ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("<iframe[^>]+src=([\"\\'])(?P<url>(?:https?:)?//out\\.pladform\\.ru/player\\?.+?)\\1"), ϒwebpage)
					if λ.IsTrue(ϒmobj) {
						return λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))
					}
					return λ.None
				})
			PladformIE__extract_url = λ.Cal(λ.StaticMethodType, PladformIE__extract_url)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):   PladformIE__VALID_URL,
				λ.NewStr("_extract_url"): PladformIE__extract_url,
			})
		}())
	})
}
