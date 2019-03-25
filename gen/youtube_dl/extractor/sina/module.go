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
 * sina/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sina.py
 */

package sina

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError            λ.Object
	HEADRequest               λ.Object
	InfoExtractor             λ.Object
	SinaIE                    λ.Object
	ϒclean_html               λ.Object
	ϒget_element_by_attribute λ.Object
	ϒint_or_none              λ.Object
	ϒqualities                λ.Object
	ϒupdate_url_query         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		HEADRequest = Ωutils.HEADRequest
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒqualities = Ωutils.ϒqualities
		ϒget_element_by_attribute = Ωutils.ϒget_element_by_attribute
		ϒclean_html = Ωutils.ϒclean_html
		SinaIE = λ.Cal(λ.TypeType, λ.NewStr("SinaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SinaIE__TESTS        λ.Object
				SinaIE__VALID_URL    λ.Object
				SinaIE__real_extract λ.Object
			)
			SinaIE__VALID_URL = λ.NewStr("(?x)https?://(?:.*?\\.)?video\\.sina\\.com\\.cn/\n                        (?:\n                            (?:view/|.*\\#)(?P<video_id>\\d+)|\n                            .+?/(?P<pseudo_id>[^/?#]+)(?:\\.s?html)|\n                            # This is used by external sites like Weibo\n                            api/sinawebApi/outplay.php/(?P<token>.+?)\\.swf\n                        )\n                  ")
			SinaIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://video.sina.com.cn/news/spj/topvideoes20160504/?opsubject_id=top1#250576622"),
					λ.NewStr("md5"): λ.NewStr("d38433e2fc886007729735650ae4b3e9"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("250576622"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("现场:克鲁兹宣布退选 特朗普将稳获提名"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://video.sina.com.cn/v/b/101314253-1290078633.html"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("101314253"),
						λ.NewStr("ext"):   λ.NewStr("flv"),
						λ.NewStr("title"): λ.NewStr("军方提高对朝情报监视级别"),
					}),
					λ.NewStr("skip"): λ.NewStr("the page does not exist or has been deleted"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://video.sina.com.cn/view/250587748.html"),
					λ.NewStr("md5"): λ.NewStr("3d1807a25c775092aab3bc157fff49b4"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("250587748"),
						λ.NewStr("ext"):   λ.NewStr("mp4"),
						λ.NewStr("title"): λ.NewStr("瞬间泪目：8年前汶川地震珍贵视频首曝光"),
					}),
				}),
			)
			SinaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription λ.Object
						ϒerror       λ.Object
						ϒfile_api    λ.Object
						ϒfile_id     λ.Object
						ϒformats     λ.Object
						ϒmobj        λ.Object
						ϒpreference  λ.Object
						ϒpseudo_id   λ.Object
						ϒquality     λ.Object
						ϒquality_id  λ.Object
						ϒrequest     λ.Object
						ϒself        = λargs[0]
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒurlh        λ.Object
						ϒvideo_data  λ.Object
						ϒvideo_id    λ.Object
						ϒwebpage     λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("video_id"))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_id))) {
						if λ.IsTrue(λ.NewBool(λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("token")) != λ.None)) {
							λ.Cal(λ.GetAttr(ϒself, "to_screen", nil), λ.NewStr("Getting video id"))
							ϒrequest = λ.Cal(HEADRequest, ϒurl)
							τmp0 = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_handle", nil), ϒrequest, λ.NewStr("NA"), λ.False)
							_ = λ.GetItem(τmp0, λ.NewInt(0))
							ϒurlh = λ.GetItem(τmp0, λ.NewInt(1))
							return λ.Cal(λ.GetAttr(ϒself, "_real_extract", nil), λ.Cal(λ.GetAttr(ϒurlh, "geturl", nil)))
						} else {
							ϒpseudo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("pseudo_id"))
							ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒpseudo_id)
							ϒerror = λ.Cal(ϒget_element_by_attribute, λ.NewStr("class"), λ.NewStr("errtitle"), ϒwebpage)
							if λ.IsTrue(ϒerror) {
								panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
									λ.GetAttr(ϒself, "IE_NAME", nil),
									λ.Cal(ϒclean_html, ϒerror),
								))), λ.KWArgs{
									{Name: "expected", Value: λ.True},
								})))
							}
							ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("video_id\\s*:\\s*'(\\d+)'"), ϒwebpage, λ.NewStr("video id"))
						}
					}
					ϒvideo_data = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("http://s.video.sina.com.cn/video/h5play"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("video_id"): ϒvideo_id,
						})},
					})
					if λ.IsTrue(λ.Ne(λ.GetItem(ϒvideo_data, λ.NewStr("code")), λ.NewInt(1))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒvideo_data, λ.NewStr("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					} else {
						ϒvideo_data = λ.GetItem(ϒvideo_data, λ.NewStr("data"))
						ϒtitle = λ.GetItem(ϒvideo_data, λ.NewStr("title"))
						ϒdescription = λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("description"))
						if λ.IsTrue(ϒdescription) {
							ϒdescription = λ.Cal(λ.GetAttr(ϒdescription, "strip", nil))
						}
						ϒpreference = λ.Cal(ϒqualities, λ.NewList(
							λ.NewStr("cif"),
							λ.NewStr("sd"),
							λ.NewStr("hd"),
							λ.NewStr("fhd"),
							λ.NewStr("ffd"),
						))
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("videos"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("mp4"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "items", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒquality_id = λ.GetItem(τmp2, λ.NewInt(0))
							ϒquality = λ.GetItem(τmp2, λ.NewInt(1))
							ϒfile_api = λ.Cal(λ.GetAttr(ϒquality, "get", nil), λ.NewStr("file_api"))
							ϒfile_id = λ.Cal(λ.GetAttr(ϒquality, "get", nil), λ.NewStr("file_id"))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒfile_api)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(!λ.IsTrue(ϒfile_id))
								}
							}()) {
								continue
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): ϒquality_id,
								λ.NewStr("url"): λ.Cal(ϒupdate_url_query, ϒfile_api, λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("vid"): ϒfile_id,
								})),
								λ.NewStr("preference"): λ.Cal(ϒpreference, ϒquality_id),
								λ.NewStr("ext"):        λ.NewStr("mp4"),
							}))
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
						return λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):          ϒvideo_id,
							λ.NewStr("title"):       ϒtitle,
							λ.NewStr("description"): ϒdescription,
							λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("image")),
							λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("length"))),
							λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("create_time"))),
							λ.NewStr("formats"):     ϒformats,
						})
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        SinaIE__TESTS,
				λ.NewStr("_VALID_URL"):    SinaIE__VALID_URL,
				λ.NewStr("_real_extract"): SinaIE__real_extract,
			})
		}())
	})
}
