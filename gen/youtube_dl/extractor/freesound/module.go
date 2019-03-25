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
 * freesound/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/freesound.py
 */

package freesound

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	FreesoundIE           λ.Object
	InfoExtractor         λ.Object
	ϒfloat_or_none        λ.Object
	ϒget_element_by_class λ.Object
	ϒget_element_by_id    λ.Object
	ϒunified_strdate      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		ϒunified_strdate = Ωutils.ϒunified_strdate
		FreesoundIE = λ.Cal(λ.TypeType, λ.NewStr("FreesoundIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FreesoundIE__TEST         λ.Object
				FreesoundIE__VALID_URL    λ.Object
				FreesoundIE__real_extract λ.Object
			)
			FreesoundIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?freesound\\.org/people/[^/]+/sounds/(?P<id>[^/]+)")
			FreesoundIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.freesound.org/people/miklovan/sounds/194503/"),
				λ.NewStr("md5"): λ.NewStr("12280ceb42c81f19a515c745eae07650"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("194503"),
					λ.NewStr("ext"):         λ.NewStr("mp3"),
					λ.NewStr("title"):       λ.NewStr("gulls in the city.wav"),
					λ.NewStr("description"): λ.NewStr("the sounds of seagulls in the city"),
					λ.NewStr("duration"):    λ.NewFloat(130.233),
					λ.NewStr("uploader"):    λ.NewStr("miklovan"),
					λ.NewStr("upload_date"): λ.NewStr("20130715"),
					λ.NewStr("tags"):        λ.ListType,
				}),
			})
			FreesoundIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						LQ_FORMAT    λ.Object
						ϒaudio_id    λ.Object
						ϒaudio_url   λ.Object
						ϒaudio_urls  λ.Object
						ϒchannels    λ.Object
						ϒdescription λ.Object
						ϒduration    λ.Object
						ϒformats     λ.Object
						ϒself        = λargs[0]
						ϒtags        λ.Object
						ϒtags_str    λ.Object
						ϒtitle       λ.Object
						ϒupload_date λ.Object
						ϒuploader    λ.Object
						ϒurl         = λargs[1]
						ϒwebpage     λ.Object
					)
					ϒaudio_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒaudio_id)
					ϒaudio_url = λ.Cal(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewStr("audio"), ϒwebpage, λ.NewStr("song url"))
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewStr("audio:title"), ϒwebpage, λ.NewStr("song title"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("(?s)id=[\"\\']sound_description[\"\\'][^>]*>(.+?)</div>"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒduration = λ.Call(ϒfloat_or_none, λ.NewArgs(λ.Cal(ϒget_element_by_class, λ.NewStr("duration"), ϒwebpage)), λ.KWArgs{
						{Name: "scale", Value: λ.NewInt(1000)},
					})
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Cal(ϒget_element_by_id, λ.NewStr("sound_date"), ϒwebpage))
					ϒuploader = λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
						λ.NewStr("audio:artist"),
						ϒwebpage,
						λ.NewStr("uploader"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒchannels = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("Channels</dt><dd>(.+?)</dd>"),
						ϒwebpage,
						λ.NewStr("channels info"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					ϒtags_str = λ.Cal(ϒget_element_by_class, λ.NewStr("tags"), ϒwebpage)
					ϒtags = func() λ.Object {
						if λ.IsTrue(ϒtags_str) {
							return λ.Cal(Ωre.ϒfindall, λ.NewStr("<a[^>]+>([^<]+)"), ϒtags_str)
						} else {
							return λ.None
						}
					}()
					ϒaudio_urls = λ.NewList(ϒaudio_url)
					LQ_FORMAT = λ.NewStr("-lq.mp3")
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒaudio_url, LQ_FORMAT))) {
						λ.Cal(λ.GetAttr(ϒaudio_urls, "append", nil), λ.Cal(λ.GetAttr(ϒaudio_url, "replace", nil), LQ_FORMAT, λ.NewStr("-hq.mp3")))
					}
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒformat_url λ.Object
									ϒquality    λ.Object
									τmp0        λ.Object
									τmp1        λ.Object
									τmp2        λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒaudio_urls))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒquality = λ.GetItem(τmp2, λ.NewInt(0))
									ϒformat_url = λ.GetItem(τmp2, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):         ϒformat_url,
										λ.NewStr("format_note"): ϒchannels,
										λ.NewStr("quality"):     ϒquality,
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒaudio_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("uploader"):    ϒuploader,
						λ.NewStr("upload_date"): ϒupload_date,
						λ.NewStr("tags"):        ϒtags,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         FreesoundIE__TEST,
				λ.NewStr("_VALID_URL"):    FreesoundIE__VALID_URL,
				λ.NewStr("_real_extract"): FreesoundIE__real_extract,
			})
		}())
	})
}
