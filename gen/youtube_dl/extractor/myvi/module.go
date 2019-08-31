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
 * myvi/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/myvi.py
 */

package myvi

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωvimple "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/vimple"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	MyviEmbedIE   λ.Object
	MyviIE        λ.Object
	SprutoBaseIE  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		SprutoBaseIE = Ωvimple.SprutoBaseIE
		MyviIE = λ.Cal(λ.TypeType, λ.NewStr("MyviIE"), λ.NewTuple(SprutoBaseIE), func() λ.Dict {
			var (
				MyviIE__VALID_URL λ.Object
			)
			MyviIE__VALID_URL = λ.NewStr("(?x)\n                        (?:\n                            https?://\n                                (?:www\\.)?\n                                myvi\\.\n                                (?:\n                                    (?:ru/player|tv)/\n                                    (?:\n                                        (?:\n                                            embed/html|\n                                            flash|\n                                            api/Video/Get\n                                        )/|\n                                        content/preloader\\.swf\\?.*\\bid=\n                                    )|\n                                    ru/watch/\n                                )|\n                            myvi:\n                        )\n                        (?P<id>[\\da-zA-Z_-]+)\n                    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): MyviIE__VALID_URL,
			})
		}())
		MyviEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("MyviEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MyviEmbedIE__VALID_URL λ.Object
				MyviEmbedIE_suitable   λ.Object
			)
			MyviEmbedIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?myvi\\.tv/(?:[^?]+\\?.*?\\bv=|embed/)(?P<id>[\\da-z]+)")
			MyviEmbedIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.GetAttr(MyviIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, MyviEmbedIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			MyviEmbedIE_suitable = λ.Cal(λ.ClassMethodType, MyviEmbedIE_suitable)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): MyviEmbedIE__VALID_URL,
				λ.NewStr("suitable"):   MyviEmbedIE_suitable,
			})
		}())
	})
}
