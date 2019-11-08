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
 * myspass/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/myspass.py
 */

package myspass

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor   λ.Object
	MySpassIE       λ.Object
	ϒcompat_str     λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
	ϒxpath_text     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒxpath_text = Ωutils.ϒxpath_text
		MySpassIE = λ.Cal(λ.TypeType, λ.NewStr("MySpassIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				MySpassIE__TEST         λ.Object
				MySpassIE__VALID_URL    λ.Object
				MySpassIE__real_extract λ.Object
			)
			MySpassIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?myspass\\.de/([^/]+/)*(?P<id>\\d+)")
			MySpassIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.myspass.de/myspass/shows/tvshows/absolute-mehrheit/Absolute-Mehrheit-vom-17022013-Die-Highlights-Teil-2--/11741/"),
				λ.NewStr("md5"): λ.NewStr("0b49f4844a068f8b33f4b7c88405862b"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("11741"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("description"): λ.NewStr("Wer kann in die Fußstapfen von Wolfgang Kubicki treten und die Mehrheit der Zuschauer hinter sich versammeln? Wird vielleicht sogar die Absolute Mehrheit geknackt und der Jackpot von 200.000 Euro mit nach Hause genommen?"),
					λ.NewStr("title"):       λ.NewStr("17.02.2013 - Die Highlights, Teil 2"),
				}),
			})
			MySpassIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒgroup        λ.Object
						ϒgroup_int    λ.Object
						ϒmetadata     λ.Object
						ϒself         = λargs[0]
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒvideo_id_int λ.Object
						ϒvideo_url    λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒmetadata = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Add(λ.NewStr("http://www.myspass.de/myspass/includes/apps/video/getvideometadataxml.php?id="), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒmetadata,
						λ.NewStr("title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒvideo_url = λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("url_flv"), λ.NewStr("download url"), λ.True)
					ϒvideo_id_int = λ.Cal(λ.IntType, ϒvideo_id)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒsearch, λ.NewStr("/myspass2009/\\d+/(\\d+)/(\\d+)/(\\d+)/"), ϒvideo_url), "groups", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒgroup = τmp1
						ϒgroup_int = λ.Cal(λ.IntType, ϒgroup)
						if λ.IsTrue(λ.Gt(ϒgroup_int, ϒvideo_id_int)) {
							ϒvideo_url = λ.Cal(λ.GetAttr(ϒvideo_url, "replace", nil), ϒgroup, λ.Cal(ϒcompat_str, λ.FloorDiv(ϒgroup_int, ϒvideo_id_int)))
						}
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):             ϒvideo_id,
						λ.NewStr("url"):            ϒvideo_url,
						λ.NewStr("title"):          ϒtitle,
						λ.NewStr("thumbnail"):      λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("imagePreview")),
						λ.NewStr("description"):    λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("description")),
						λ.NewStr("duration"):       λ.Cal(ϒparse_duration, λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("duration"))),
						λ.NewStr("series"):         λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("format")),
						λ.NewStr("season_number"):  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("season"))),
						λ.NewStr("season_id"):      λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("season_id")),
						λ.NewStr("episode"):        ϒtitle,
						λ.NewStr("episode_number"): λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒmetadata, λ.NewStr("episode"))),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         MySpassIE__TEST,
				λ.NewStr("_VALID_URL"):    MySpassIE__VALID_URL,
				λ.NewStr("_real_extract"): MySpassIE__real_extract,
			})
		}())
	})
}
