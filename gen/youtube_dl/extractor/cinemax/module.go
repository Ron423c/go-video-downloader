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
 * cinemax/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/cinemax.py
 */

package cinemax

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωhbo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/hbo"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CinemaxIE λ.Object
	HBOBaseIE λ.Object
)

func init() {
	λ.InitModule(func() {
		HBOBaseIE = Ωhbo.HBOBaseIE
		CinemaxIE = λ.Cal(λ.TypeType, λ.StrLiteral("CinemaxIE"), λ.NewTuple(HBOBaseIE), func() λ.Dict {
			var (
				CinemaxIE__VALID_URL    λ.Object
				CinemaxIE__real_extract λ.Object
			)
			CinemaxIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?cinemax\\.com/(?P<path>[^/]+/video/[0-9a-z-]+-(?P<id>\\d+))")
			CinemaxIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo     λ.Object
						ϒpath     λ.Object
						ϒself     = λargs[0]
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						τmp0      λ.Object
					)
					τmp0 = λ.Calm(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups")
					ϒpath = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.IntLiteral(1))
					ϒinfo = λ.Calm(ϒself, "_extract_info", λ.Mod(λ.StrLiteral("https://www.cinemax.com/%s.xml"), ϒpath), ϒvideo_id)
					λ.SetItem(ϒinfo, λ.StrLiteral("id"), ϒvideo_id)
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    CinemaxIE__VALID_URL,
				"_real_extract": CinemaxIE__real_extract,
			})
		}())
	})
}
