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
 * canalplus/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/canalplus.py
 */

package canalplus

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CanalplusIE      λ.Object
	InfoExtractor    λ.Object
	ϒint_or_none     λ.Object
	ϒqualities       λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒqualities = Ωutils.ϒqualities
		ϒunified_strdate = Ωutils.ϒunified_strdate
		CanalplusIE = λ.Cal(λ.TypeType, λ.NewStr("CanalplusIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CanalplusIE__GEO_COUNTRIES       λ.Object
				CanalplusIE__SITE_ID_MAP         λ.Object
				CanalplusIE__TESTS               λ.Object
				CanalplusIE__VALID_URL           λ.Object
				CanalplusIE__VIDEO_INFO_TEMPLATE λ.Object
				CanalplusIE__real_extract        λ.Object
			)
			CanalplusIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?(?P<site>mycanal|piwiplus)\\.fr/(?:[^/]+/)*(?P<display_id>[^?/]+)(?:\\.html\\?.*\\bvid=|/p/)(?P<id>\\d+)")
			CanalplusIE__VIDEO_INFO_TEMPLATE = λ.NewStr("http://service.canal-plus.com/video/rest/getVideosLiees/%s/%s?format=json")
			CanalplusIE__SITE_ID_MAP = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("mycanal"):  λ.NewStr("cplus"),
				λ.NewStr("piwiplus"): λ.NewStr("teletoon"),
			})
			CanalplusIE__GEO_COUNTRIES = λ.NewList(λ.NewStr("FR"))
			CanalplusIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.mycanal.fr/d17-emissions/lolywood/p/1397061"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1397061"),
						λ.NewStr("display_id"):  λ.NewStr("lolywood"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Euro 2016 : Je préfère te prévenir - Lolywood - Episode 34"),
						λ.NewStr("description"): λ.NewStr("md5:7d97039d455cb29cdba0d652a0efaa5e"),
						λ.NewStr("upload_date"): λ.NewStr("20160602"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.piwiplus.fr/videos-piwi/pid1405-le-labyrinthe-boing-super-ranger.html?vid=1108190"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("1108190"),
						λ.NewStr("display_id"):  λ.NewStr("pid1405-le-labyrinthe-boing-super-ranger"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("BOING SUPER RANGER - Ep : Le labyrinthe"),
						λ.NewStr("description"): λ.NewStr("md5:4cea7a37153be42c1ba2c1d3064376ff"),
						λ.NewStr("upload_date"): λ.NewStr("20140724"),
					}),
					λ.NewStr("expected_warnings"): λ.NewList(λ.NewStr("HTTP Error 403: Forbidden")),
				}),
			)
			CanalplusIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒformat_id  λ.Object
						ϒformat_url λ.Object
						ϒformats    λ.Object
						ϒinfo_url   λ.Object
						ϒinfos      λ.Object
						ϒmedia      λ.Object
						ϒpreference λ.Object
						ϒself       = λargs[0]
						ϒsite       λ.Object
						ϒsite_id    λ.Object
						ϒthumbnails λ.Object
						ϒtitrage    λ.Object
						ϒurl        = λargs[1]
						ϒvideo_data λ.Object
						ϒvideo_id   λ.Object
						τmp0        λ.Object
						τmp1        λ.Object
						τmp2        λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒsite = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(2))
					ϒsite_id = λ.GetItem(λ.GetAttr(ϒself, "_SITE_ID_MAP", nil), ϒsite)
					ϒinfo_url = λ.Mod(λ.GetAttr(ϒself, "_VIDEO_INFO_TEMPLATE", nil), λ.NewTuple(
						ϒsite_id,
						ϒvideo_id,
					))
					ϒvideo_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒinfo_url, ϒvideo_id, λ.NewStr("Downloading video JSON"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒvideo_data, λ.ListType)) {
						ϒvideo_data = λ.GetItem(λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒvideo λ.Object
										τmp0   λ.Object
										τmp1   λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒvideo_data)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒvideo = τmp1
										if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("ID")), ϒvideo_id)) {
											λgen.Yield(ϒvideo)
										}
									}
									return λ.None
								})
							}))), λ.NewInt(0))
					}
					ϒmedia = λ.GetItem(ϒvideo_data, λ.NewStr("MEDIA"))
					ϒinfos = λ.GetItem(ϒvideo_data, λ.NewStr("INFOS"))
					ϒpreference = λ.Cal(ϒqualities, λ.NewList(
						λ.NewStr("MOBILE"),
						λ.NewStr("BAS_DEBIT"),
						λ.NewStr("HAUT_DEBIT"),
						λ.NewStr("HD"),
					))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒmedia, λ.NewStr("VIDEOS")), "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformat_url = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformat_url))) {
							continue
						}
						if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("HLS"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒformat_url,
								ϒvideo_id,
								λ.NewStr("mp4"),
								λ.NewStr("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: ϒformat_id},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒformat_id, λ.NewStr("HDS"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
									λ.Add(ϒformat_url, λ.NewStr("?hdcore=2.11.3")),
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "f4m_id", Value: ϒformat_id},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):        λ.Add(ϒformat_url, λ.NewStr("?secret=pqzerjlsmdkjfoiuerhsdlfknaes")),
									λ.NewStr("format_id"):  ϒformat_id,
									λ.NewStr("preference"): λ.Cal(ϒpreference, ϒformat_id),
								}))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒimage_id  λ.Object
									ϒimage_url λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
									τmp2       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("images"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "items", nil)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒimage_id = λ.GetItem(τmp2, λ.NewInt(0))
									ϒimage_url = λ.GetItem(τmp2, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("id"):  ϒimage_id,
										λ.NewStr("url"): ϒimage_url,
									}))
								}
								return λ.None
							})
						})))
					ϒtitrage = λ.GetItem(ϒinfos, λ.NewStr("TITRAGE"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
						λ.NewStr("title"): λ.Mod(λ.NewStr("%s - %s"), λ.NewTuple(
							λ.GetItem(ϒtitrage, λ.NewStr("TITRE")),
							λ.GetItem(ϒtitrage, λ.NewStr("SOUS_TITRE")),
						)),
						λ.NewStr("upload_date"):   λ.Cal(ϒunified_strdate, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒinfos, "get", nil), λ.NewStr("PUBLICATION"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("DATE"))),
						λ.NewStr("thumbnails"):    ϒthumbnails,
						λ.NewStr("description"):   λ.Cal(λ.GetAttr(ϒinfos, "get", nil), λ.NewStr("DESCRIPTION")),
						λ.NewStr("duration"):      λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfos, "get", nil), λ.NewStr("DURATION"))),
						λ.NewStr("view_count"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfos, "get", nil), λ.NewStr("NB_VUES"))),
						λ.NewStr("like_count"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfos, "get", nil), λ.NewStr("NB_LIKES"))),
						λ.NewStr("comment_count"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfos, "get", nil), λ.NewStr("NB_COMMENTS"))),
						λ.NewStr("formats"):       ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_GEO_COUNTRIES"):       CanalplusIE__GEO_COUNTRIES,
				λ.NewStr("_SITE_ID_MAP"):         CanalplusIE__SITE_ID_MAP,
				λ.NewStr("_TESTS"):               CanalplusIE__TESTS,
				λ.NewStr("_VALID_URL"):           CanalplusIE__VALID_URL,
				λ.NewStr("_VIDEO_INFO_TEMPLATE"): CanalplusIE__VIDEO_INFO_TEMPLATE,
				λ.NewStr("_real_extract"):        CanalplusIE__real_extract,
			})
		}())
	})
}
