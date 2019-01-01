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
 * drtv/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/drtv.py
 */

package drtv

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DRTVIE            λ.Object
	DRTVLiveIE        λ.Object
	ExtractorError    λ.Object
	InfoExtractor     λ.Object
	ϒfloat_or_none    λ.Object
	ϒint_or_none      λ.Object
	ϒmimetype2ext     λ.Object
	ϒparse_iso8601    λ.Object
	ϒremove_end       λ.Object
	ϒupdate_url_query λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒremove_end = Ωutils.ϒremove_end
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		DRTVIE = λ.Cal(λ.TypeType, λ.NewStr("DRTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DRTVIE_IE_NAME        λ.Object
				DRTVIE__GEO_BYPASS    λ.Object
				DRTVIE__GEO_COUNTRIES λ.Object
				DRTVIE__TESTS         λ.Object
				DRTVIE__VALID_URL     λ.Object
				DRTVIE__real_extract  λ.Object
			)
			DRTVIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?dr\\.dk/(?:tv/se|nyheder|radio/ondemand)/(?:[^/]+/)*(?P<id>[\\da-z-]+)(?:[/#?]|$)")
			DRTVIE__GEO_BYPASS = λ.False
			DRTVIE__GEO_COUNTRIES = λ.NewList(λ.NewStr("DK"))
			DRTVIE_IE_NAME = λ.NewStr("drtv")
			DRTVIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.dr.dk/tv/se/boern/ultra/klassen-ultra/klassen-darlig-taber-10"),
					λ.NewStr("md5"): λ.NewStr("7ae17b4e18eb5d29212f424a7511c184"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("klassen-darlig-taber-10"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Klassen - Dårlig taber (10)"),
						λ.NewStr("description"): λ.NewStr("md5:815fe1b7fa656ed80580f31e8b3c79aa"),
						λ.NewStr("timestamp"):   λ.NewInt(1471991907),
						λ.NewStr("upload_date"): λ.NewStr("20160823"),
						λ.NewStr("duration"):    λ.NewFloat(606.84),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.dr.dk/nyheder/indland/live-christianias-rydning-af-pusher-street-er-i-gang"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("christiania-pusher-street-ryddes-drdkrjpo"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("LIVE Christianias rydning af Pusher Street er i gang"),
						λ.NewStr("description"): λ.NewStr("md5:2a71898b15057e9b97334f61d04e6eb5"),
						λ.NewStr("timestamp"):   λ.NewInt(1472800279),
						λ.NewStr("upload_date"): λ.NewStr("20160902"),
						λ.NewStr("duration"):    λ.NewFloat(131.4),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.dr.dk/tv/se/historien-om-danmark/-/historien-om-danmark-stenalder"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("historien-om-danmark-stenalder"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Historien om Danmark: Stenalder (1)"),
						λ.NewStr("description"): λ.NewStr("md5:8c66dcbc1669bbc6f873879880f37f2a"),
						λ.NewStr("timestamp"):   λ.NewInt(1490401996),
						λ.NewStr("upload_date"): λ.NewStr("20170325"),
						λ.NewStr("duration"):    λ.NewFloat(3502.04),
						λ.NewStr("formats"):     λ.NewStr("mincount:20"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			DRTVIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						LANGS                  λ.Object
						ϒasset                 λ.Object
						ϒasset_target          λ.Object
						ϒbitrate               λ.Object
						ϒdata                  λ.Object
						ϒdescription           λ.Object
						ϒduration              λ.Object
						ϒf                     λ.Object
						ϒf4m_formats           λ.Object
						ϒformat_id             λ.Object
						ϒformats               λ.Object
						ϒkind                  λ.Object
						ϒlang                  λ.Object
						ϒlink                  λ.Object
						ϒpreference            λ.Object
						ϒprogramcard           λ.Object
						ϒrestricted_to_denmark λ.Object
						ϒself                  = λargs[0]
						ϒsubs                  λ.Object
						ϒsubtitles             λ.Object
						ϒsubtitles_list        λ.Object
						ϒtarget                λ.Object
						ϒthumbnail             λ.Object
						ϒtimestamp             λ.Object
						ϒtitle                 λ.Object
						ϒuri                   λ.Object
						ϒurl                   = λargs[1]
						ϒvideo_id              λ.Object
						ϒwebpage               λ.Object
						τmp0                   λ.Object
						τmp1                   λ.Object
						τmp2                   λ.Object
						τmp3                   λ.Object
						τmp4                   λ.Object
						τmp5                   λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr(">Programmet er ikke længere tilgængeligt")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("Video %s is not available"), ϒvideo_id)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewTuple(
						λ.NewStr("data-(?:material-identifier|episode-slug)=\"([^\"]+)\""),
						λ.NewStr("data-resource=\"[^>\"]+mu/programcard/expanded/([^\"]+)\""),
					), ϒwebpage, λ.NewStr("video id"))
					ϒprogramcard = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://www.dr.dk/mu/programcard/expanded/%s"), ϒvideo_id), ϒvideo_id, λ.NewStr("Downloading video JSON"))
					ϒdata = λ.GetItem(λ.GetItem(ϒprogramcard, λ.NewStr("Data")), λ.NewInt(0))
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(ϒremove_end, λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}), λ.NewStr(" | TV | DR")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.GetItem(ϒdata, λ.NewStr("Title"))
						}
					}()
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("Description"))
						}
					}()
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒdata, "get", nil), λ.NewStr("CreatedTime")))
					ϒthumbnail = λ.None
					ϒduration = λ.None
					ϒrestricted_to_denmark = λ.False
					ϒformats = λ.NewList()
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒdata, λ.NewStr("Assets")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒasset = τmp1
						ϒkind = λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("Kind"))
						if λ.IsTrue(λ.Eq(ϒkind, λ.NewStr("Image"))) {
							ϒthumbnail = λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("Uri"))
						} else {
							if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
								λ.NewStr("VideoResource"),
								λ.NewStr("AudioResource"),
							), ϒkind))) {
								ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("DurationInMilliseconds")), λ.NewInt(1000))
								ϒrestricted_to_denmark = λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("RestrictedToDenmark"))
								ϒasset_target = λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("Target"))
								τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("Links"), λ.NewList()))
								for {
									if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
										break
									}
									ϒlink = τmp3
									ϒuri = λ.Cal(λ.GetAttr(ϒlink, "get", nil), λ.NewStr("Uri"))
									if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒuri))) {
										continue
									}
									ϒtarget = λ.Cal(λ.GetAttr(ϒlink, "get", nil), λ.NewStr("Target"))
									ϒformat_id = func() λ.Object {
										if λv := ϒtarget; λ.IsTrue(λv) {
											return λv
										} else {
											return λ.NewStr("")
										}
									}()
									ϒpreference = λ.None
									if λ.IsTrue(λ.NewBool(λ.Contains(λ.NewTuple(
										λ.NewStr("SpokenSubtitles"),
										λ.NewStr("SignLanguage"),
									), ϒasset_target))) {
										ϒpreference = λ.Neg(λ.NewInt(1))
										τmp4 = λ.IAdd(ϒformat_id, λ.Mod(λ.NewStr("-%s"), ϒasset_target))
										ϒformat_id = τmp4
									}
									if λ.IsTrue(λ.Eq(ϒtarget, λ.NewStr("HDS"))) {
										ϒf4m_formats = λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
											λ.Add(ϒuri, λ.NewStr("?hdcore=3.3.0&plugin=aasp-3.3.0.99.43")),
											ϒvideo_id,
											ϒpreference,
										), λ.KWArgs{
											{Name: "f4m_id", Value: ϒformat_id},
											{Name: "fatal", Value: λ.False},
										})
										if λ.IsTrue(λ.Eq(ϒkind, λ.NewStr("AudioResource"))) {
											τmp4 = λ.Cal(λ.BuiltinIter, ϒf4m_formats)
											for {
												if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
													break
												}
												ϒf = τmp5
												λ.SetItem(ϒf, λ.NewStr("vcodec"), λ.NewStr("none"))
											}
										}
										λ.Cal(λ.GetAttr(ϒformats, "extend", nil), ϒf4m_formats)
									} else {
										if λ.IsTrue(λ.Eq(ϒtarget, λ.NewStr("HLS"))) {
											λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
												ϒuri,
												ϒvideo_id,
												λ.NewStr("mp4"),
											), λ.KWArgs{
												{Name: "entry_protocol", Value: λ.NewStr("m3u8_native")},
												{Name: "preference", Value: ϒpreference},
												{Name: "m3u8_id", Value: ϒformat_id},
												{Name: "fatal", Value: λ.False},
											}))
										} else {
											ϒbitrate = λ.Cal(λ.GetAttr(ϒlink, "get", nil), λ.NewStr("Bitrate"))
											if λ.IsTrue(ϒbitrate) {
												τmp4 = λ.IAdd(ϒformat_id, λ.Mod(λ.NewStr("-%s"), ϒbitrate))
												ϒformat_id = τmp4
											}
											λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("url"):       ϒuri,
												λ.NewStr("format_id"): ϒformat_id,
												λ.NewStr("tbr"):       λ.Cal(ϒint_or_none, ϒbitrate),
												λ.NewStr("ext"):       λ.Cal(λ.GetAttr(ϒlink, "get", nil), λ.NewStr("FileFormat")),
												λ.NewStr("vcodec"): func() λ.Object {
													if λ.IsTrue(λ.Eq(ϒkind, λ.NewStr("AudioResource"))) {
														return λ.NewStr("none")
													} else {
														return λ.None
													}
												}(),
												λ.NewStr("preference"): ϒpreference,
											}))
										}
									}
								}
								ϒsubtitles_list = λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("SubtitlesList"))
								if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒsubtitles_list, λ.ListType)) {
									LANGS = λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("Danish"): λ.NewStr("da"),
									})
									τmp2 = λ.Cal(λ.BuiltinIter, ϒsubtitles_list)
									for {
										if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
											break
										}
										ϒsubs = τmp3
										if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒsubs, "get", nil), λ.NewStr("Uri"))))) {
											continue
										}
										ϒlang = func() λ.Object {
											if λv := λ.Cal(λ.GetAttr(ϒsubs, "get", nil), λ.NewStr("Language")); λ.IsTrue(λv) {
												return λv
											} else {
												return λ.NewStr("da")
											}
										}()
										λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), λ.Cal(λ.GetAttr(LANGS, "get", nil), ϒlang, ϒlang), λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"): λ.GetItem(ϒsubs, λ.NewStr("Uri")),
											λ.NewStr("ext"): func() λ.Object {
												if λv := λ.Cal(ϒmimetype2ext, λ.Cal(λ.GetAttr(ϒsubs, "get", nil), λ.NewStr("MimeType"))); λ.IsTrue(λv) {
													return λv
												} else {
													return λ.NewStr("vtt")
												}
											}(),
										}))
									}
								}
							}
						}
					}
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒformats)); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒrestricted_to_denmark
						}
					}()) {
						λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), λ.NewArgs(λ.NewStr("Unfortunately, DR is not allowed to show this program outside Denmark.")), λ.KWArgs{
							{Name: "countries", Value: λ.GetAttr(ϒself, "_GEO_COUNTRIES", nil)},
						})
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("duration"):    ϒduration,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("subtitles"):   ϒsubtitles,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):        DRTVIE_IE_NAME,
				λ.NewStr("_GEO_BYPASS"):    DRTVIE__GEO_BYPASS,
				λ.NewStr("_GEO_COUNTRIES"): DRTVIE__GEO_COUNTRIES,
				λ.NewStr("_TESTS"):         DRTVIE__TESTS,
				λ.NewStr("_VALID_URL"):     DRTVIE__VALID_URL,
				λ.NewStr("_real_extract"):  DRTVIE__real_extract,
			})
		}())
		DRTVLiveIE = λ.Cal(λ.TypeType, λ.NewStr("DRTVLiveIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DRTVLiveIE__VALID_URL λ.Object
			)
			DRTVLiveIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?dr\\.dk/(?:tv|TV)/live/(?P<id>[\\da-z-]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): DRTVLiveIE__VALID_URL,
			})
		}())
	})
}