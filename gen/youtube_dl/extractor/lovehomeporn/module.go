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
 * lovehomeporn/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/lovehomeporn.py
 */

package lovehomeporn

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωnuevo "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/nuevo"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	LoveHomePornIE λ.Object
	NuevoBaseIE    λ.Object
)

func init() {
	λ.InitModule(func() {
		NuevoBaseIE = Ωnuevo.NuevoBaseIE
		LoveHomePornIE = λ.Cal(λ.TypeType, λ.StrLiteral("LoveHomePornIE"), λ.NewTuple(NuevoBaseIE), func() λ.Dict {
			var (
				LoveHomePornIE__VALID_URL    λ.Object
				LoveHomePornIE__real_extract λ.Object
			)
			LoveHomePornIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?lovehomeporn\\.com/video/(?P<id>\\d+)(?:/(?P<display_id>[^/?#&]+))?")
			LoveHomePornIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒinfo       λ.Object
						ϒmobj       λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("id"))
					ϒdisplay_id = λ.Calm(ϒmobj, "group", λ.StrLiteral("display_id"))
					ϒinfo = λ.Calm(ϒself, "_extract_nuevo", λ.Mod(λ.StrLiteral("http://lovehomeporn.com/media/nuevo/config.php?key=%s"), ϒvideo_id), ϒvideo_id)
					λ.Calm(ϒinfo, "update", λ.DictLiteral(map[string]λ.Object{
						"display_id": ϒdisplay_id,
						"age_limit":  λ.IntLiteral(18),
					}))
					return ϒinfo
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    LoveHomePornIE__VALID_URL,
				"_real_extract": LoveHomePornIE__real_extract,
			})
		}())
	})
}
