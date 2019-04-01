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
 * cctv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/cctv.py
 */

package cctv

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CCTVIE             λ.Object
	InfoExtractor      λ.Object
	ϒcompat_str        λ.Object
	ϒfloat_or_none     λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		CCTVIE = λ.Cal(λ.TypeType, λ.NewStr("CCTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CCTVIE__TESTS        λ.Object
				CCTVIE__VALID_URL    λ.Object
				CCTVIE__real_extract λ.Object
			)
			CCTVIE__VALID_URL = λ.NewStr("https?://(?:(?:[^/]+)\\.(?:cntv|cctv)\\.(?:com|cn)|(?:www\\.)?ncpa-classic\\.com)/(?:[^/]+/)*?(?P<id>[^/?#&]+?)(?:/index)?(?:\\.s?html|[?#&]|$)")
			CCTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://sports.cntv.cn/2016/02/12/ARTIaBRxv4rTT1yWf1frW2wi160212.shtml"),
					λ.NewStr("md5"): λ.NewStr("d61ec00a493e09da810bf406a078f691"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("5ecdbeab623f4973b40ff25f18b174e8"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("[NBA]二少联手砍下46分 雷霆主场击败鹈鹕（快讯）"),
						λ.NewStr("description"): λ.NewStr("md5:7e14a5328dc5eb3d1cd6afbbe0574e95"),
						λ.NewStr("duration"):    λ.NewInt(98),
						λ.NewStr("uploader"):    λ.NewStr("songjunjie"),
						λ.NewStr("timestamp"):   λ.NewInt(1455279956),
						λ.NewStr("upload_date"): λ.NewStr("20160212"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://tv.cctv.com/2016/02/05/VIDEUS7apq3lKrHG9Dncm03B160205.shtml"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("efc5d49e5b3b4ab2b34f3a502b73d3ae"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("[赛车]“车王”舒马赫恢复情况成谜（快讯）"),
						λ.NewStr("description"): λ.NewStr("2月4日，蒙特泽莫罗透露了关于“车王”舒马赫恢复情况，但情况是否属实遭到了质疑。"),
						λ.NewStr("duration"):    λ.NewInt(37),
						λ.NewStr("uploader"):    λ.NewStr("shujun"),
						λ.NewStr("timestamp"):   λ.NewInt(1454677291),
						λ.NewStr("upload_date"): λ.NewStr("20160205"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://english.cntv.cn/special/four_comprehensives/index.shtml"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("4bb9bb4db7a6471ba85fdeda5af0381e"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("NHnews008 ANNUAL POLITICAL SEASON"),
						λ.NewStr("description"): λ.NewStr("Four Comprehensives"),
						λ.NewStr("duration"):    λ.NewInt(60),
						λ.NewStr("uploader"):    λ.NewStr("zhangyunlei"),
						λ.NewStr("timestamp"):   λ.NewInt(1425385521),
						λ.NewStr("upload_date"): λ.NewStr("20150303"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://cctv.cntv.cn/lm/tvseries_russian/yilugesanghua/index.shtml"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("b15f009ff45c43968b9af583fc2e04b2"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Путь，усыпанный космеями Серия 1"),
						λ.NewStr("description"): λ.NewStr("Путь, усыпанный космеями"),
						λ.NewStr("duration"):    λ.NewInt(2645),
						λ.NewStr("uploader"):    λ.NewStr("renxue"),
						λ.NewStr("timestamp"):   λ.NewInt(1477479241),
						λ.NewStr("upload_date"): λ.NewStr("20161026"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ncpa-classic.com/2013/05/22/VIDE1369219508996867.shtml"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("a194cfa7f18c426b823d876668325946"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("小泽征尔音乐塾 音乐梦想无国界"),
						λ.NewStr("duration"):    λ.NewInt(2173),
						λ.NewStr("timestamp"):   λ.NewInt(1369248264),
						λ.NewStr("upload_date"): λ.NewStr("20130522"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ncpa-classic.com/clt/more/416/index.shtml"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("a8606119a4884588a79d81c02abecc16"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("来自维也纳的新年贺礼"),
						λ.NewStr("description"): λ.NewStr("md5:f13764ae8dd484e84dd4b39d5bcba2a7"),
						λ.NewStr("duration"):    λ.NewInt(1578),
						λ.NewStr("uploader"):    λ.NewStr("djy"),
						λ.NewStr("timestamp"):   λ.NewInt(1482942419),
						λ.NewStr("upload_date"): λ.NewStr("20161228"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("expected_warnings"): λ.NewList(λ.NewStr("Failed to download m3u8 information")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://ent.cntv.cn/2016/01/18/ARTIjprSSJH8DryTVr5Bx8Wb160118.shtml"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://tv.cntv.cn/video/C39296/e0210d949f113ddfb38d31f00a4e5c44"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://english.cntv.cn/2016/09/03/VIDEhnkB5y9AgHyIEVphCEz1160903.shtml"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://tv.cctv.com/2016/09/07/VIDE5C1FnlX5bUywlrjhxXOV160907.shtml"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://tv.cntv.cn/video/C39296/95cfac44cabd3ddc4a9438780a4e5c44"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			CCTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchapters_key λ.Object
						ϒdata         λ.Object
						ϒdescription  λ.Object
						ϒduration     λ.Object
						ϒformats      λ.Object
						ϒhls_url      λ.Object
						ϒquality      λ.Object
						ϒself         = λargs[0]
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒuploader     λ.Object
						ϒurl          = λargs[1]
						ϒvideo        λ.Object
						ϒvideo_id     λ.Object
						ϒvideo_url    λ.Object
						ϒwebpage      λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewList(
						λ.NewStr("var\\s+guid\\s*=\\s*[\"\\']([\\da-fA-F]+)"),
						λ.NewStr("videoCenterId[\"\\']\\s*,\\s*[\"\\']([\\da-fA-F]+)"),
						λ.NewStr("changePlayer\\s*\\(\\s*[\"\\']([\\da-fA-F]+)"),
						λ.NewStr("load[Vv]ideo\\s*\\(\\s*[\"\\']([\\da-fA-F]+)"),
						λ.NewStr("var\\s+initMyAray\\s*=\\s*[\"\\']([\\da-fA-F]+)"),
						λ.NewStr("var\\s+ids\\s*=\\s*\\[[\"\\']([\\da-fA-F]+)"),
					), ϒwebpage, λ.NewStr("video id"))
					ϒdata = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("http://vdn.apps.cntv.cn/api/getHttpVideoInfo.do"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("pid"):      ϒvideo_id,
							λ.NewStr("url"):      ϒurl,
							λ.NewStr("idl"):      λ.NewInt(32),
							λ.NewStr("idlr"):     λ.NewInt(32),
							λ.NewStr("modifyed"): λ.NewStr("false"),
						})},
					})
					ϒtitle = λ.GetItem(ϒdata, λ.NewStr("title"))
					ϒformats = λ.NewList()
					ϒvideo = λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("video"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒvideo, λ.DictType)) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.NewTuple(
							λ.NewStr("lowChapters"),
							λ.NewStr("chapters"),
						)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒquality = λ.GetItem(τmp2, λ.NewInt(0))
							ϒchapters_key = λ.GetItem(τmp2, λ.NewInt(1))
							ϒvideo_url = λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, ϒchapters_key), λ.NewInt(0)), λ.NewStr("url"))
								}), ϒcompat_str)
							if λ.IsTrue(ϒvideo_url) {
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):        ϒvideo_url,
									λ.NewStr("format_id"):  λ.NewStr("http"),
									λ.NewStr("quality"):    ϒquality,
									λ.NewStr("preference"): λ.Neg(λ.NewInt(1)),
								}))
							}
						}
					}
					ϒhls_url = λ.Cal(ϒtry_get, ϒdata, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.NewStr("hls_url"))
						}), ϒcompat_str)
					if λ.IsTrue(ϒhls_url) {
						ϒhls_url = λ.Cal(Ωre.ϒsub, λ.NewStr("maxbr=\\d+&?"), λ.NewStr(""), ϒhls_url)
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒhls_url,
							ϒvideo_id,
							λ.NewStr("mp4"),
						), λ.KWArgs{
							{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒuploader = λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("editer_name"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("description"),
						ϒwebpage,
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("f_pgmtime")))
					ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.NewStr("totalLength"))
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("uploader"):    ϒuploader,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        CCTVIE__TESTS,
				λ.NewStr("_VALID_URL"):    CCTVIE__VALID_URL,
				λ.NewStr("_real_extract"): CCTVIE__real_extract,
			})
		}())
	})
}