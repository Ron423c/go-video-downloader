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
 * businessinsider/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/businessinsider.py
 */

package businessinsider

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωjwplatform "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/jwplatform"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BusinessInsiderIE λ.Object
	InfoExtractor     λ.Object
	JWPlatformIE      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		JWPlatformIE = Ωjwplatform.JWPlatformIE
		BusinessInsiderIE = λ.Cal(λ.TypeType, λ.NewStr("BusinessInsiderIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BusinessInsiderIE__TESTS        λ.Object
				BusinessInsiderIE__VALID_URL    λ.Object
				BusinessInsiderIE__real_extract λ.Object
			)
			BusinessInsiderIE__VALID_URL = λ.NewStr("https?://(?:[^/]+\\.)?businessinsider\\.(?:com|nl)/(?:[^/]+/)*(?P<id>[^/?#&]+)")
			BusinessInsiderIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://uk.businessinsider.com/how-much-radiation-youre-exposed-to-in-everyday-life-2016-6"),
					λ.NewStr("md5"): λ.NewStr("ca237a53a8eb20b6dc5bd60564d4ab3e"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("hZRllCfw"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Here's how much radiation you're exposed to in everyday life"),
						λ.NewStr("description"): λ.NewStr("md5:9a0d6e2c279948aadaa5e84d6d9b99bd"),
						λ.NewStr("upload_date"): λ.NewStr("20170709"),
						λ.NewStr("timestamp"):   λ.NewInt(1499606400),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.businessinsider.nl/5-scientifically-proven-things-make-you-less-attractive-2017-7/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.businessinsider.com/excel-index-match-vlookup-video-how-to-2015-2?IR=T"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			BusinessInsiderIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒjwplatform_id λ.Object
						ϒself          = λargs[0]
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒwebpage       λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒjwplatform_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewTuple(
						λ.NewStr("data-media-id=[\"\\']([a-zA-Z0-9]{8})"),
						λ.NewStr("id=[\"\\']jwplayer_([a-zA-Z0-9]{8})"),
						λ.NewStr("id[\"\\']?\\s*:\\s*[\"\\']?([a-zA-Z0-9]{8})"),
					), ϒwebpage, λ.NewStr("jwplatform id"))
					return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Mod(λ.NewStr("jwplatform:%s"), ϒjwplatform_id)), λ.KWArgs{
						{Name: "ie", Value: λ.Cal(λ.GetAttr(JWPlatformIE, "ie_key", nil))},
						{Name: "video_id", Value: ϒvideo_id},
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        BusinessInsiderIE__TESTS,
				λ.NewStr("_VALID_URL"):    BusinessInsiderIE__VALID_URL,
				λ.NewStr("_real_extract"): BusinessInsiderIE__real_extract,
			})
		}())
	})
}
