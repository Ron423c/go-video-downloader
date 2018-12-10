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
 * vuclip/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/vuclip.py
 */

package vuclip

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	VuClipIE                      λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒparse_duration               λ.Object
	ϒremove_end                   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ExtractorError = Ωutils.ExtractorError
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒremove_end = Ωutils.ϒremove_end
		VuClipIE = λ.Cal(λ.TypeType, λ.NewStr("VuClipIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				VuClipIE__TEST         λ.Object
				VuClipIE__VALID_URL    λ.Object
				VuClipIE__real_extract λ.Object
			)
			VuClipIE__VALID_URL = λ.NewStr("https?://(?:m\\.)?vuclip\\.com/w\\?.*?cid=(?P<id>[0-9]+)")
			VuClipIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://m.vuclip.com/w?cid=1129900602&bu=8589892792&frm=w&z=34801&op=0&oc=843169247&section=recommend"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):       λ.NewStr("1129900602"),
					λ.NewStr("ext"):      λ.NewStr("3gp"),
					λ.NewStr("title"):    λ.NewStr("Top 10 TV Convicts"),
					λ.NewStr("duration"): λ.NewInt(733),
				}),
			})
			VuClipIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒad_m       λ.Object
						ϒadfree_url λ.Object
						ϒduration   λ.Object
						ϒerror_msg  λ.Object
						ϒformats    λ.Object
						ϒself       = λargs[0]
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						ϒurlr       λ.Object
						ϒvideo_id   λ.Object
						ϒvideo_url  λ.Object
						ϒwebpage    λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒad_m = λ.Cal(Ωre.ϒsearch, λ.NewStr("value=\"No.*?\" onClick=\"location.href='([^\"']+)'\""), ϒwebpage)
					if λ.IsTrue(ϒad_m) {
						ϒurlr = λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl)
						ϒadfree_url = λ.Add(λ.Add(λ.Add(λ.GetAttr(ϒurlr, "scheme", nil), λ.NewStr("://")), λ.GetAttr(ϒurlr, "netloc", nil)), λ.Cal(λ.GetAttr(ϒad_m, "group", nil), λ.NewInt(1)))
						ϒwebpage = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
							ϒadfree_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "note", Value: λ.NewStr("Download post-ad page")},
						})
					}
					ϒerror_msg = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<p class=\"message\">(.*?)</p>"),
						ϒwebpage,
						λ.NewStr("error message"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒerror_msg) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒerror_msg,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<a[^>]+href=\"([^\"]+)\"[^>]*><img[^>]+src=\"[^\"]*/play\\.gif"),
						ϒwebpage,
						λ.NewStr("video URL"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvideo_url) {
						ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒvideo_url,
						}))
					} else {
						ϒformats = λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_html5_media_entries", nil), ϒurl, ϒwebpage, ϒvideo_id), λ.NewInt(0)), λ.NewStr("formats"))
					}
					ϒtitle = λ.Cal(ϒremove_end, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<title>(.*?)-\\s*Vuclip</title>"), ϒwebpage, λ.NewStr("title")), "strip", nil)), λ.NewStr(" - Video"))
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("[(>]([0-9]+:[0-9]+)(?:<span|\\))"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):       ϒvideo_id,
						λ.NewStr("formats"):  ϒformats,
						λ.NewStr("title"):    ϒtitle,
						λ.NewStr("duration"): ϒduration,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         VuClipIE__TEST,
				λ.NewStr("_VALID_URL"):    VuClipIE__VALID_URL,
				λ.NewStr("_real_extract"): VuClipIE__real_extract,
			})
		}())
	})
}
