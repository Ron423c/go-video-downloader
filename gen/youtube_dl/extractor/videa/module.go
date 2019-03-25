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
 * videa/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/videa.py
 */

package videa

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	VideaIE        λ.Object
	ϒint_or_none   λ.Object
	ϒmimetype2ext  λ.Object
	ϒparse_codecs  λ.Object
	ϒxpath_element λ.Object
	ϒxpath_text    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_codecs = Ωutils.ϒparse_codecs
		ϒxpath_element = Ωutils.ϒxpath_element
		ϒxpath_text = Ωutils.ϒxpath_text
		VideaIE = λ.Cal(λ.TypeType, λ.NewStr("VideaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VideaIE__TESTS        λ.Object
				VideaIE__VALID_URL    λ.Object
				VideaIE__real_extract λ.Object
			)
			VideaIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        videa(?:kid)?\\.hu/\n                        (?:\n                            videok/(?:[^/]+/)*[^?#&]+-|\n                            player\\?.*?\\bv=|\n                            player/v/\n                        )\n                        (?P<id>[^?#&]+)\n                    ")
			VideaIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://videa.hu/videok/allatok/az-orult-kigyasz-285-kigyot-kigyo-8YfIAjxwWGwT8HVQ"),
					λ.NewStr("md5"): λ.NewStr("97a7af41faeaffd9f1fc864a7c7e7603"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("8YfIAjxwWGwT8HVQ"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("Az őrült kígyász 285 kígyót enged szabadon"),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*"),
						λ.NewStr("duration"):  λ.NewInt(21),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://videa.hu/videok/origo/jarmuvek/supercars-elozes-jAHDWfWSJH5XuFhH"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://videa.hu/player?v=8YfIAjxwWGwT8HVQ"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://videa.hu/player/v/8YfIAjxwWGwT8HVQ?autoplay=1"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://videakid.hu/videok/origo/jarmuvek/supercars-elozes-jAHDWfWSJH5XuFhH"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://videakid.hu/player?v=8YfIAjxwWGwT8HVQ"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://videakid.hu/player/v/8YfIAjxwWGwT8HVQ?autoplay=1"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			VideaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit  λ.Object
						ϒduration   λ.Object
						ϒf          λ.Object
						ϒformats    λ.Object
						ϒinfo       λ.Object
						ϒis_adult   λ.Object
						ϒself       = λargs[0]
						ϒsource     λ.Object
						ϒsource_url λ.Object
						ϒsources    λ.Object
						ϒthumbnail  λ.Object
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒinfo = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						λ.NewStr("http://videa.hu/videaplayer_get_xml.php"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("v"): ϒvideo_id,
						})},
					})
					ϒvideo = λ.Call(ϒxpath_element, λ.NewArgs(
						ϒinfo,
						λ.NewStr(".//video"),
						λ.NewStr("video"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒsources = λ.Call(ϒxpath_element, λ.NewArgs(
						ϒinfo,
						λ.NewStr(".//video_sources"),
						λ.NewStr("sources"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒvideo,
						λ.NewStr("./title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒsources, "findall", nil), λ.NewStr("./video_source")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						ϒsource_url = λ.GetAttr(ϒsource, "text", nil)
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒsource_url))) {
							continue
						}
						ϒf = λ.Cal(ϒparse_codecs, λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("codecs")))
						λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒsource_url,
							λ.NewStr("ext"): func() λ.Object {
								if λv := λ.Cal(ϒmimetype2ext, λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("mimetype"))); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewStr("mp4")
								}
							}(),
							λ.NewStr("format_id"): λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("name")),
							λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("height"))),
						}))
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnail = λ.Cal(ϒxpath_text, ϒvideo, λ.NewStr("./poster_src"))
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo, λ.NewStr("./duration")))
					ϒage_limit = λ.None
					ϒis_adult = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒvideo,
						λ.NewStr("./is_adult_content"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒis_adult) {
						ϒage_limit = func() λ.Object {
							if λ.IsTrue(λ.Eq(ϒis_adult, λ.NewStr("1"))) {
								return λ.NewInt(18)
							} else {
								return λ.NewInt(0)
							}
						}()
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("duration"):  ϒduration,
						λ.NewStr("age_limit"): ϒage_limit,
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        VideaIE__TESTS,
				λ.NewStr("_VALID_URL"):    VideaIE__VALID_URL,
				λ.NewStr("_real_extract"): VideaIE__real_extract,
			})
		}())
	})
}
