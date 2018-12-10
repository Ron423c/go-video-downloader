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
 * sportdeutschland/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/sportdeutschland.py
 */

package sportdeutschland

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor      λ.Object
	SportDeutschlandIE λ.Object
	ϒparse_iso8601     λ.Object
	ϒsanitized_Request λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		SportDeutschlandIE = λ.Cal(λ.TypeType, λ.NewStr("SportDeutschlandIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SportDeutschlandIE__TESTS        λ.Object
				SportDeutschlandIE__VALID_URL    λ.Object
				SportDeutschlandIE__real_extract λ.Object
			)
			SportDeutschlandIE__VALID_URL = λ.NewStr("https?://sportdeutschland\\.tv/(?P<sport>[^/?#]+)/(?P<id>[^?#/]+)(?:$|[?#])")
			SportDeutschlandIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://sportdeutschland.tv/badminton/live-li-ning-badminton-weltmeisterschaft-2014-kopenhagen"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("live-li-ning-badminton-weltmeisterschaft-2014-kopenhagen"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("re:Li-Ning Badminton Weltmeisterschaft 2014 Kopenhagen"),
						λ.NewStr("categories"):  λ.NewList(λ.NewStr("Badminton")),
						λ.NewStr("view_count"):  λ.IntType,
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("description"): λ.NewStr("re:Die Badminton-WM 2014 aus Kopenhagen bei Sportdeutschland\\.TV"),
						λ.NewStr("timestamp"):   λ.IntType,
						λ.NewStr("upload_date"): λ.NewStr("re:^201408[23][0-9]$"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.NewStr("Live stream"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://sportdeutschland.tv/li-ning-badminton-wm-2014/lee-li-ning-badminton-weltmeisterschaft-2014-kopenhagen-herren-einzel-wei-vs"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("lee-li-ning-badminton-weltmeisterschaft-2014-kopenhagen-herren-einzel-wei-vs"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("upload_date"): λ.NewStr("20140825"),
						λ.NewStr("description"): λ.NewStr("md5:60a20536b57cee7d9a4ec005e8687504"),
						λ.NewStr("timestamp"):   λ.NewInt(1408976060),
						λ.NewStr("duration"):    λ.NewInt(2732),
						λ.NewStr("title"):       λ.NewStr("Li-Ning Badminton Weltmeisterschaft 2014 Kopenhagen: Herren Einzel, Wei Lee vs. Keun Lee"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("view_count"):  λ.IntType,
						λ.NewStr("categories"):  λ.NewList(λ.NewStr("Li-Ning Badminton WM 2014")),
					}),
				}),
			)
			SportDeutschlandIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_url     λ.Object
						ϒasset       λ.Object
						ϒbase_url    λ.Object
						ϒbase_url_el λ.Object
						ϒcategories  λ.Object
						ϒdata        λ.Object
						ϒformats     λ.Object
						ϒm3u8_url    λ.Object
						ϒmobj        λ.Object
						ϒreq         λ.Object
						ϒself        = λargs[0]
						ϒsmil_doc    λ.Object
						ϒsmil_url    λ.Object
						ϒsport_id    λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒsport_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("sport"))
					ϒapi_url = λ.Mod(λ.NewStr("http://proxy.vidibusdynamic.net/sportdeutschland.tv/api/permalinks/%s/%s?access_token=true"), λ.NewTuple(
						ϒsport_id,
						ϒvideo_id,
					))
					ϒreq = λ.Call(ϒsanitized_Request, λ.NewArgs(ϒapi_url), λ.KWArgs{
						{Name: "headers", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("Accept"):  λ.NewStr("application/vnd.vidibus.v2.html+json"),
							λ.NewStr("Referer"): ϒurl,
						})},
					})
					ϒdata = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒreq, ϒvideo_id)
					ϒasset = λ.GetItem(ϒdata, λ.NewStr("asset"))
					ϒcategories = λ.NewList(λ.GetItem(λ.GetItem(ϒdata, λ.NewStr("section")), λ.NewStr("title")))
					ϒformats = λ.NewList()
					ϒsmil_url = λ.GetItem(ϒasset, λ.NewStr("video"))
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒsmil_url, λ.NewStr(".smil")))) {
						ϒm3u8_url = λ.Cal(λ.GetAttr(ϒsmil_url, "replace", nil), λ.NewStr(".smil"), λ.NewStr(".m3u8"))
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "ext", Value: λ.NewStr("mp4")},
						}))
						ϒsmil_doc = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
							ϒsmil_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "note", Value: λ.NewStr("Downloading SMIL metadata")},
						})
						ϒbase_url_el = λ.Cal(λ.GetAttr(ϒsmil_doc, "find", nil), λ.NewStr("./head/meta"))
						if λ.IsTrue(ϒbase_url_el) {
							ϒbase_url = λ.GetItem(λ.GetAttr(ϒbase_url_el, "attrib", nil), λ.NewStr("base"))
						}
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒn   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒsmil_doc, "findall", nil), λ.NewStr("./body/video")))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒn = τmp1
										λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("format_id"): λ.NewStr("rmtp"),
											λ.NewStr("url"): func() λ.Object {
												if λ.IsTrue(ϒbase_url_el) {
													return ϒbase_url
												} else {
													return λ.GetItem(λ.GetAttr(ϒn, "attrib", nil), λ.NewStr("src"))
												}
											}(),
											λ.NewStr("play_path"):   λ.GetItem(λ.GetAttr(ϒn, "attrib", nil), λ.NewStr("src")),
											λ.NewStr("ext"):         λ.NewStr("flv"),
											λ.NewStr("preference"):  λ.Neg(λ.NewInt(100)),
											λ.NewStr("format_note"): λ.NewStr("Seems to fail at example stream"),
										}))
									}
									return λ.None
								})
							}))))
					} else {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒsmil_url,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("title"):       λ.GetItem(ϒasset, λ.NewStr("title")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("image")),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("teaser")),
						λ.NewStr("duration"):    λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("duration")),
						λ.NewStr("categories"):  ϒcategories,
						λ.NewStr("view_count"):  λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("views")),
						λ.NewStr("rtmp_live"):   λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("live")),
						λ.NewStr("timestamp"):   λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("date"))),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        SportDeutschlandIE__TESTS,
				λ.NewStr("_VALID_URL"):    SportDeutschlandIE__VALID_URL,
				λ.NewStr("_real_extract"): SportDeutschlandIE__real_extract,
			})
		}())
	})
}
