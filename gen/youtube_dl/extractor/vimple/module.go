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
 * vimple/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/vimple.py
 */

package vimple

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor λ.Object
	SprutoBaseIE  λ.Object
	VimpleIE      λ.Object
	ϒint_or_none  λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		SprutoBaseIE = λ.Cal(λ.TypeType, λ.NewStr("SprutoBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SprutoBaseIE__extract_spruto λ.Object
			)
			SprutoBaseIE__extract_spruto = λ.NewFunction("_extract_spruto",
				[]λ.Param{
					{Name: "self"},
					{Name: "spruto"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒduration  λ.Object
						ϒformats   λ.Object
						ϒplaylist  λ.Object
						ϒself      = λargs[0]
						ϒspruto    = λargs[1]
						ϒthumbnail λ.Object
						ϒtitle     λ.Object
						ϒvideo_id  = λargs[2]
					)
					ϒplaylist = λ.GetItem(λ.GetItem(ϒspruto, λ.NewStr("playlist")), λ.NewInt(0))
					ϒtitle = λ.GetItem(ϒplaylist, λ.NewStr("title"))
					ϒvideo_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒplaylist, "get", nil), λ.NewStr("videoId")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒthumbnail = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒplaylist, "get", nil), λ.NewStr("posterUrl")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒplaylist, "get", nil), λ.NewStr("thumbnailUrl"))
						}
					}()
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒplaylist, "get", nil), λ.NewStr("duration")))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒf   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒplaylist, λ.NewStr("video")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒf = τmp1
									λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): λ.GetItem(ϒf, λ.NewStr("url")),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("duration"):  ϒduration,
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_extract_spruto"): SprutoBaseIE__extract_spruto,
			})
		}())
		VimpleIE = λ.Cal(λ.TypeType, λ.NewStr("VimpleIE"), λ.NewTuple(SprutoBaseIE), func() λ.Dict {
			var (
				VimpleIE__VALID_URL    λ.Object
				VimpleIE__real_extract λ.Object
			)
			VimpleIE__VALID_URL = λ.NewStr("https?://(?:player\\.vimple\\.(?:ru|co)/iframe|vimple\\.(?:ru|co))/(?P<id>[\\da-f-]{32,36})")
			VimpleIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself     = λargs[0]
						ϒspruto   λ.Object
						ϒurl      = λargs[1]
						ϒvideo_id λ.Object
						ϒwebpage  λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://player.vimple.ru/iframe/%s"), ϒvideo_id), ϒvideo_id)
					ϒspruto = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("sprutoData\\s*:\\s*({.+?}),\\r\\n"), ϒwebpage, λ.NewStr("spruto data")), ϒvideo_id)
					return λ.Cal(λ.GetAttr(ϒself, "_extract_spruto", nil), ϒspruto, ϒvideo_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    VimpleIE__VALID_URL,
				λ.NewStr("_real_extract"): VimpleIE__real_extract,
			})
		}())
	})
}
