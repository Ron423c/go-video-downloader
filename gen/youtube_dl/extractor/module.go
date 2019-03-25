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
 * __init__/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/__init__.py
 */

package extractor

import (
	Ωextractors "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/extractors"
	Ωlib "github.com/tenta-browser/go-video-downloader/lib"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	RegisterExtractor      λ.Object
	ϒextractorDict         λ.Object
	ϒgen_extractor_classes λ.Object
	ϒget_info_extractor    λ.Object
	ϒkey                   λ.Object
	ϒklass                 λ.Object
	τmp0                   λ.Object
	τmp1                   λ.Object
)

func init() {
	λ.InitModule(func() {
		RegisterExtractor = Ωlib.RegisterExtractor
		ϒextractorDict = λ.NewDictWithTable(map[λ.Object]λ.Object{})
		τmp0 = λ.Cal(λ.BuiltinIter, Ωextractors.ϒ__ALL__)
		for {
			if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
				break
			}
			ϒklass = τmp1
			ϒkey = λ.Cal(λ.GetAttr(ϒklass, "ie_key", nil))
			λ.Cal(RegisterExtractor, ϒkey, ϒklass)
			λ.SetItem(ϒextractorDict, ϒkey, ϒklass)
		}
		ϒgen_extractor_classes = λ.NewFunction("gen_extractor_classes",
			nil,
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				return Ωextractors.ϒ__ALL__
			})
		ϒget_info_extractor = λ.NewFunction("get_info_extractor",
			[]λ.Param{
				{Name: "ie_name"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒie_name = λargs[0]
				)
				return λ.Cal(λ.GetAttr(ϒextractorDict, "get", nil), ϒie_name)
			})
	})
}
