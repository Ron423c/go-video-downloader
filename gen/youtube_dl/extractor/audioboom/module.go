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
 * audioboom/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/audioboom.py
 */

package audioboom

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AudioBoomIE    λ.Object
	InfoExtractor  λ.Object
	ϒfloat_or_none λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		AudioBoomIE = λ.Cal(λ.TypeType, λ.NewStr("AudioBoomIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				AudioBoomIE__TESTS        λ.Object
				AudioBoomIE__VALID_URL    λ.Object
				AudioBoomIE__real_extract λ.Object
			)
			AudioBoomIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?audioboom\\.com/(?:boos|posts)/(?P<id>[0-9]+)")
			AudioBoomIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://audioboom.com/boos/4279833-3-09-2016-czaban-hour-3?t=0"),
					λ.NewStr("md5"): λ.NewStr("63a8d73a055c6ed0f1e51921a10a5a76"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):           λ.NewStr("4279833"),
						λ.NewStr("ext"):          λ.NewStr("mp3"),
						λ.NewStr("title"):        λ.NewStr("3/09/2016 Czaban Hour 3"),
						λ.NewStr("description"):  λ.NewStr("Guest:   Nate Davis - NFL free agency,   Guest:   Stan Gans"),
						λ.NewStr("duration"):     λ.NewFloat(2245.72),
						λ.NewStr("uploader"):     λ.NewStr("SB Nation A.M."),
						λ.NewStr("uploader_url"): λ.NewStr("re:https?://(?:www\\.)?audioboom\\.com/channel/steveczabanyahoosportsradio"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://audioboom.com/posts/4279833-3-09-2016-czaban-hour-3?t=0"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			AudioBoomIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaudio_url    λ.Object
						ϒclip         λ.Object
						ϒclip_store   λ.Object
						ϒclips        λ.Object
						ϒdescription  λ.Object
						ϒduration     λ.Object
						ϒfrom_clip    λ.Object
						ϒself         = λargs[0]
						ϒtitle        λ.Object
						ϒuploader     λ.Object
						ϒuploader_url λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒwebpage      λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒclip = λ.None
					ϒclip_store = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.Mod(λ.NewStr("data-new-clip-store=([\"\\'])(?P<json>{.*?\"clipId\"\\s*:\\s*%s.*?})\\1"), ϒvideo_id),
							ϒwebpage,
							λ.NewStr("clip store"),
						), λ.KWArgs{
							{Name: "default", Value: λ.NewStr("{}")},
							{Name: "group", Value: λ.NewStr("json")},
						}),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒclip_store) {
						ϒclips = λ.Cal(λ.GetAttr(ϒclip_store, "get", nil), λ.NewStr("clips"))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒclips; !λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Cal(λ.BuiltinIsInstance, ϒclips, λ.ListType); !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, λ.GetItem(ϒclips, λ.NewInt(0)), λ.DictType)
							}
						}()) {
							ϒclip = λ.GetItem(ϒclips, λ.NewInt(0))
						}
					}
					ϒfrom_clip = λ.NewFunction("from_clip",
						[]λ.Param{
							{Name: "field"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒfield = λargs[0]
							)
							if λ.IsTrue(ϒclip) {
								return λ.Cal(λ.GetAttr(ϒclip, "get", nil), ϒfield)
							}
							return λ.None
						})
					ϒaudio_url = func() λ.Object {
						if λv := λ.Cal(ϒfrom_clip, λ.NewStr("clipURLPriorToLoading")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewStr("audio"), ϒwebpage, λ.NewStr("audio url"))
						}
					}()
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(ϒfrom_clip, λ.NewStr("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒdescription = func() λ.Object {
						if λv := λ.Cal(ϒfrom_clip, λ.NewStr("description")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_description", nil), ϒwebpage)
						}
					}()
					ϒduration = λ.Cal(ϒfloat_or_none, func() λ.Object {
						if λv := λ.Cal(ϒfrom_clip, λ.NewStr("duration")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("weibo:audio:duration"), ϒwebpage)
						}
					}())
					ϒuploader = func() λ.Object {
						if λv := λ.Cal(ϒfrom_clip, λ.NewStr("author")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Call(λ.GetAttr(ϒself, "_og_search_property", nil), λ.NewArgs(
								λ.NewStr("audio:artist"),
								ϒwebpage,
								λ.NewStr("uploader"),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							})
						}
					}()
					ϒuploader_url = func() λ.Object {
						if λv := λ.Cal(ϒfrom_clip, λ.NewStr("author_url")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("audioboo:channel"), ϒwebpage, λ.NewStr("uploader url"))
						}
					}()
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):           ϒvideo_id,
						λ.NewStr("url"):          ϒaudio_url,
						λ.NewStr("title"):        ϒtitle,
						λ.NewStr("description"):  ϒdescription,
						λ.NewStr("duration"):     ϒduration,
						λ.NewStr("uploader"):     ϒuploader,
						λ.NewStr("uploader_url"): ϒuploader_url,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        AudioBoomIE__TESTS,
				λ.NewStr("_VALID_URL"):    AudioBoomIE__VALID_URL,
				λ.NewStr("_real_extract"): AudioBoomIE__real_extract,
			})
		}())
	})
}
