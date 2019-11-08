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
 * performgroup/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/performgroup.py
 */

package performgroup

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	PerformGroupIE λ.Object
	ϒint_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		PerformGroupIE = λ.Cal(λ.TypeType, λ.NewStr("PerformGroupIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PerformGroupIE__VALID_URL    λ.Object
				PerformGroupIE__call_api     λ.Object
				PerformGroupIE__real_extract λ.Object
			)
			PerformGroupIE__VALID_URL = λ.NewStr("https?://player\\.performgroup\\.com/eplayer(?:/eplayer\\.html|\\.js)#/?(?P<id>[0-9a-f]{26})\\.(?P<auth_token>[0-9a-z]{26})")
			PerformGroupIE__call_api = λ.NewFunction("_call_api",
				[]λ.Param{
					{Name: "self"},
					{Name: "service"},
					{Name: "auth_token"},
					{Name: "content_id"},
					{Name: "referer_url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒauth_token  = λargs[2]
						ϒcontent_id  = λargs[3]
						ϒreferer_url = λargs[4]
						ϒself        = λargs[0]
						ϒservice     = λargs[1]
					)
					return λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("http://ep3.performfeeds.com/ep%s/%s/%s/"), λ.NewTuple(
							ϒservice,
							ϒauth_token,
							ϒcontent_id,
						)),
						ϒcontent_id,
					), λ.KWArgs{
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Referer"): ϒreferer_url,
							λ.NewStr("Origin"):  λ.NewStr("http://player.performgroup.com"),
						})},
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("_fmt"): λ.NewStr("json"),
						})},
					})
				})
			PerformGroupIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒauth_token λ.Object
						ϒbootstrap  λ.Object
						ϒc          λ.Object
						ϒc_url      λ.Object
						ϒformat_id  λ.Object
						ϒformats    λ.Object
						ϒhds_url    λ.Object
						ϒhls_url    λ.Object
						ϒmedia      λ.Object
						ϒplayer_id  λ.Object
						ϒself       = λargs[0]
						ϒtbr        λ.Object
						ϒurl        = λargs[1]
						ϒvideo      λ.Object
						ϒvideo_id   λ.Object
						ϒvod        λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒsearch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒplayer_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒauth_token = λ.GetItem(τmp0, λ.NewInt(1))
					ϒbootstrap = λ.Cal(λ.GetAttr(ϒself, "_call_api", nil), λ.NewStr("bootstrap"), ϒauth_token, ϒplayer_id, ϒurl)
					ϒvideo = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒbootstrap, λ.NewStr("config")), λ.NewStr("dataSource")), λ.NewStr("sourceItems")), λ.NewInt(0)), λ.NewStr("videos")), λ.NewInt(0))
					ϒvideo_id = λ.GetItem(ϒvideo, λ.NewStr("uuid"))
					ϒvod = λ.Cal(λ.GetAttr(ϒself, "_call_api", nil), λ.NewStr("vod"), ϒauth_token, ϒvideo_id, ϒurl)
					ϒmedia = λ.GetItem(λ.GetItem(λ.GetItem(λ.GetItem(ϒvod, λ.NewStr("videos")), λ.NewStr("video")), λ.NewInt(0)), λ.NewStr("media"))
					ϒformats = λ.NewList()
					ϒhls_url = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("hls"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("url"))
					if λ.IsTrue(ϒhls_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒhls_url,
							ϒvideo_id,
							λ.NewStr("mp4"),
							λ.NewStr("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒhds_url = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("hds"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("url"))
					if λ.IsTrue(ϒhds_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							λ.Add(ϒhds_url, λ.NewStr("?hdcore")),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.NewStr("hds")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("content"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒc = τmp1
						ϒc_url = λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒc_url))) {
							continue
						}
						ϒtbr = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("bitrate")), λ.NewInt(1000))
						ϒformat_id = λ.NewStr("http")
						if λ.IsTrue(ϒtbr) {
							τmp2 = λ.IAdd(ϒformat_id, λ.Mod(λ.NewStr("-%d"), ϒtbr))
							ϒformat_id = τmp2
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format_id"): ϒformat_id,
							λ.NewStr("url"):       ϒc_url,
							λ.NewStr("tbr"):       ϒtbr,
							λ.NewStr("width"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("width"))),
							λ.NewStr("height"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("height"))),
							λ.NewStr("filesize"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("fileSize"))),
							λ.NewStr("vcodec"):    λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("type")),
							λ.NewStr("fps"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("videoFrameRate"))),
							λ.NewStr("vbr"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("videoRate")), λ.NewInt(1000)),
							λ.NewStr("abr"):       λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("audioRate")), λ.NewInt(1000)),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       λ.GetItem(ϒvideo, λ.NewStr("title")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("description")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("poster")),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("publishedTime")), λ.NewInt(1000)),
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    PerformGroupIE__VALID_URL,
				λ.NewStr("_call_api"):     PerformGroupIE__call_api,
				λ.NewStr("_real_extract"): PerformGroupIE__real_extract,
			})
		}())
	})
}
