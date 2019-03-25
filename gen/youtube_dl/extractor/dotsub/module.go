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
 * dotsub/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/dotsub.py
 */

package dotsub

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DotsubIE       λ.Object
	InfoExtractor  λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		DotsubIE = λ.Cal(λ.TypeType, λ.NewStr("DotsubIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DotsubIE__TESTS        λ.Object
				DotsubIE__VALID_URL    λ.Object
				DotsubIE__real_extract λ.Object
			)
			DotsubIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?dotsub\\.com/view/(?P<id>[^/]+)")
			DotsubIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://dotsub.com/view/9c63db2a-fa95-4838-8e6e-13deafe47f09"),
					λ.NewStr("md5"): λ.NewStr("21c7ff600f545358134fea762a6d42b6"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("9c63db2a-fa95-4838-8e6e-13deafe47f09"),
						λ.NewStr("ext"):         λ.NewStr("flv"),
						λ.NewStr("title"):       λ.NewStr("MOTIVATION - \"It's Possible\" Best Inspirational Video Ever"),
						λ.NewStr("description"): λ.NewStr("md5:41af1e273edbbdfe4e216a78b9d34ac6"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://dotsub.com/media/9c63db2a-fa95-4838-8e6e-13deafe47f09/p"),
						λ.NewStr("duration"):    λ.NewInt(198),
						λ.NewStr("uploader"):    λ.NewStr("liuxt"),
						λ.NewStr("timestamp"):   λ.NewFloat(1385778501.104),
						λ.NewStr("upload_date"): λ.NewStr("20131130"),
						λ.NewStr("view_count"):  λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://dotsub.com/view/747bcf58-bd59-45b7-8c8c-ac312d084ee6"),
					λ.NewStr("md5"): λ.NewStr("2bb4a83896434d5c26be868c609429a3"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("168006778"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Apartments and flats in Raipur the white symphony"),
						λ.NewStr("description"): λ.NewStr("md5:784d0639e6b7d1bc29530878508e38fe"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://dotsub.com/media/747bcf58-bd59-45b7-8c8c-ac312d084ee6/p"),
						λ.NewStr("duration"):    λ.NewInt(290),
						λ.NewStr("timestamp"):   λ.NewFloat(1476767794.281),
						λ.NewStr("upload_date"): λ.NewStr("20161018"),
						λ.NewStr("uploader"):    λ.NewStr("parthivi001"),
						λ.NewStr("uploader_id"): λ.NewStr("user52596202"),
						λ.NewStr("view_count"):  λ.IntType,
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Vimeo")),
				}),
			)
			DotsubIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo       λ.Object
						ϒinfo_dict  λ.Object
						ϒself       = λargs[0]
						ϒsetup_data λ.Object
						ϒurl        = λargs[1]
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
						ϒwebpage    λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://dotsub.com/api/media/%s/metadata"), ϒvideo_id), ϒvideo_id)
					ϒvideo_url = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("mediaURI"))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
						ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
						ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewList(
								λ.NewStr("<source[^>]+src=\"([^\"]+)\""),
								λ.NewStr("\"file\"\\s*:\\s*\\'([^\\']+)"),
							),
							ϒwebpage,
							λ.NewStr("video url"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						ϒinfo_dict = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):  ϒvideo_id,
							λ.NewStr("url"): ϒvideo_url,
							λ.NewStr("ext"): λ.NewStr("flv"),
						})
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
						ϒsetup_data = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("(?s)data-setup=([\\'\"])(?P<content>(?!\\1).+?)\\1"),
							ϒwebpage,
							λ.NewStr("setup data"),
						), λ.KWArgs{
							{Name: "group", Value: λ.NewStr("content")},
						}), ϒvideo_id)
						ϒinfo_dict = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("_type"): λ.NewStr("url_transparent"),
							λ.NewStr("url"):   λ.GetItem(ϒsetup_data, λ.NewStr("src")),
						})
					}
					λ.Cal(λ.GetAttr(ϒinfo_dict, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"):       λ.GetItem(ϒinfo, λ.NewStr("title")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("description")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("screenshotURI")),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("duration")), λ.NewInt(1000)),
						λ.NewStr("uploader"):    λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("user")),
						λ.NewStr("timestamp"):   λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("dateCreated")), λ.NewInt(1000)),
						λ.NewStr("view_count"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("numberOfViews"))),
					}))
					return ϒinfo_dict
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        DotsubIE__TESTS,
				λ.NewStr("_VALID_URL"):    DotsubIE__VALID_URL,
				λ.NewStr("_real_extract"): DotsubIE__real_extract,
			})
		}())
	})
}
