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
 * turner/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/turner.py
 */

package turner

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE       λ.Object
	ExtractorError    λ.Object
	TurnerBaseIE      λ.Object
	ϒcompat_str       λ.Object
	ϒdetermine_ext    λ.Object
	ϒfloat_or_none    λ.Object
	ϒint_or_none      λ.Object
	ϒparse_duration   λ.Object
	ϒstrip_or_none    λ.Object
	ϒupdate_url_query λ.Object
	ϒurl_or_none      λ.Object
	ϒxpath_attr       λ.Object
	ϒxpath_text       λ.Object
)

func init() {
	λ.InitModule(func() {
		AdobePassIE = Ωadobepass.AdobePassIE
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒxpath_attr = Ωutils.ϒxpath_attr
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ExtractorError = Ωutils.ExtractorError
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒurl_or_none = Ωutils.ϒurl_or_none
		TurnerBaseIE = λ.Cal(λ.TypeType, λ.NewStr("TurnerBaseIE"), λ.NewTuple(AdobePassIE), func() λ.Dict {
			var (
				TurnerBaseIE__extract_cvp_info  λ.Object
				TurnerBaseIE__extract_timestamp λ.Object
			)
			TurnerBaseIE__extract_timestamp = λ.NewFunction("_extract_timestamp",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_data"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself       = λargs[0]
						ϒvideo_data = λargs[1]
					)
					_ = ϒself
					return λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_attr, ϒvideo_data, λ.NewStr("dateCreated"), λ.NewStr("uts")))
				})
			TurnerBaseIE__extract_cvp_info = λ.NewFunction("_extract_cvp_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "data_src"},
					{Name: "video_id"},
					{Name: "path_data", Def: λ.NewDictWithTable(map[λ.Object]λ.Object{})},
					{Name: "ap_data", Def: λ.NewDictWithTable(map[λ.Object]λ.Object{})},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒap_data          = λargs[4]
						ϒbase_path_data   λ.Object
						ϒcontent_id       λ.Object
						ϒdata_src         = λargs[1]
						ϒext              λ.Object
						ϒf                λ.Object
						ϒformat_id        λ.Object
						ϒformats          λ.Object
						ϒis_live          λ.Object
						ϒlang             λ.Object
						ϒm3u8_formats     λ.Object
						ϒmedia_src        λ.Object
						ϒmobj             λ.Object
						ϒpath_data        = λargs[3]
						ϒrex              λ.Object
						ϒsecure_path_data λ.Object
						ϒself             = λargs[0]
						ϒsource           λ.Object
						ϒsubtitles        λ.Object
						ϒthumbnails       λ.Object
						ϒtitle            λ.Object
						ϒtrack            λ.Object
						ϒtrack_url        λ.Object
						ϒurls             λ.Object
						ϒvideo_data       λ.Object
						ϒvideo_file       λ.Object
						ϒvideo_id         = λargs[2]
						ϒvideo_url        λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
						τmp2              λ.Object
						τmp3              λ.Object
					)
					ϒvideo_data = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), ϒdata_src, ϒvideo_id)
					ϒvideo_id = λ.GetItem(λ.GetAttr(ϒvideo_data, "attrib", nil), λ.NewStr("id"))
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒvideo_data,
						λ.NewStr("headline"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒcontent_id = func() λ.Object {
						if λv := λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("contentId")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒurls = λ.NewList()
					ϒformats = λ.NewList()
					ϒrex = λ.Cal(Ωre.ϒcompile, λ.NewStr("(?P<width>[0-9]+)x(?P<height>[0-9]+)(?:_(?P<bitrate>[0-9]+))?"))
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "findall", nil), λ.NewStr(".//file")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒvideo_file = τmp1
						ϒvideo_url = λ.Cal(λ.GetAttr(λ.GetAttr(ϒvideo_file, "text", nil), "strip", nil))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_url))) {
							continue
						}
						ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo_url, "startswith", nil), λ.NewStr("/mp4:protected/"))) {
							continue
						} else {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo_url, "startswith", nil), λ.NewStr("/secure/"))) {
								ϒsecure_path_data = λ.Cal(λ.GetAttr(ϒpath_data, "get", nil), λ.NewStr("secure"))
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒsecure_path_data))) {
									continue
								}
								ϒvideo_url = λ.Cal(λ.GetAttr(ϒself, "_add_akamai_spe_token", nil), λ.GetItem(ϒsecure_path_data, λ.NewStr("tokenizer_src")), λ.Add(λ.GetItem(ϒsecure_path_data, λ.NewStr("media_src")), ϒvideo_url), ϒcontent_id, ϒap_data)
							} else {
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(Ωre.ϒmatch, λ.NewStr("https?://"), ϒvideo_url)))) {
									ϒbase_path_data = λ.Cal(λ.GetAttr(ϒpath_data, "get", nil), ϒext, λ.Cal(λ.GetAttr(ϒpath_data, "get", nil), λ.NewStr("default"), λ.NewDictWithTable(map[λ.Object]λ.Object{})))
									ϒmedia_src = λ.Cal(λ.GetAttr(ϒbase_path_data, "get", nil), λ.NewStr("media_src"))
									if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒmedia_src))) {
										continue
									}
									ϒvideo_url = λ.Add(ϒmedia_src, ϒvideo_url)
								}
							}
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒurls, ϒvideo_url))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒurls, "append", nil), ϒvideo_url)
						ϒformat_id = λ.Cal(λ.GetAttr(ϒvideo_file, "get", nil), λ.NewStr("bitrate"))
						if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("smil"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_smil_formats", nil), λ.NewArgs(
								ϒvideo_url,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
								ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
									ϒvideo_url,
									ϒvideo_id,
									λ.NewStr("mp4"),
								), λ.KWArgs{
									{Name: "m3u8_id", Value: func() λ.Object {
										if λv := ϒformat_id; λ.IsTrue(λv) {
											return λv
										} else {
											return λ.NewStr("hls")
										}
									}()},
									{Name: "fatal", Value: λ.False},
								})
								if λ.IsTrue(func() λ.Object {
									if λv := λ.NewBool(λ.Contains(ϒvideo_url, λ.NewStr("/secure/"))); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(ϒvideo_url, λ.NewStr("?hdnea=")))
									}
								}()) {
									τmp2 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
									for {
										if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
											break
										}
										ϒf = τmp3
										λ.SetItem(ϒf, λ.NewStr("_seekable"), λ.False)
									}
								}
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), ϒm3u8_formats)
							} else {
								if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("f4m"))) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
										λ.Cal(ϒupdate_url_query, ϒvideo_url, λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("hdcore"): λ.NewStr("3.7.0"),
										})),
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "f4m_id", Value: func() λ.Object {
											if λv := ϒformat_id; λ.IsTrue(λv) {
												return λv
											} else {
												return λ.NewStr("hds")
											}
										}()},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("format_id"): ϒformat_id,
										λ.NewStr("url"):       ϒvideo_url,
										λ.NewStr("ext"):       ϒext,
									})
									ϒmobj = λ.Cal(λ.GetAttr(ϒrex, "search", nil), λ.Add(ϒformat_id, ϒvideo_url))
									if λ.IsTrue(ϒmobj) {
										λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("width"):  λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("width"))),
											λ.NewStr("height"): λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("height"))),
											λ.NewStr("tbr"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("bitrate"))),
										}))
									} else {
										if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒformat_id, ϒcompat_str)) {
											if λ.IsTrue(λ.Cal(λ.GetAttr(ϒformat_id, "isdigit", nil))) {
												λ.SetItem(ϒf, λ.NewStr("tbr"), λ.Cal(λ.IntType, ϒformat_id))
											} else {
												ϒmobj = λ.Cal(Ωre.ϒmatch, λ.NewStr("ios_(audio|[0-9]+)$"), ϒformat_id)
												if λ.IsTrue(ϒmobj) {
													if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(1)), λ.NewStr("audio"))) {
														λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
															λ.NewStr("vcodec"): λ.NewStr("none"),
															λ.NewStr("ext"):    λ.NewStr("m4a"),
														}))
													} else {
														λ.SetItem(ϒf, λ.NewStr("tbr"), λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewInt(1))))
													}
												}
											}
										}
									}
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
								}
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "findall", nil), λ.NewStr("closedCaptions/source")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource = τmp1
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒsource, "findall", nil), λ.NewStr("track")))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒtrack = τmp3
							ϒtrack_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒtrack, "get", nil), λ.NewStr("url")))
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒtrack_url)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Cal(λ.GetAttr(ϒtrack_url, "endswith", nil), λ.NewStr("/big"))
								}
							}()) {
								continue
							}
							ϒlang = func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒtrack, "get", nil), λ.NewStr("lang")); λ.IsTrue(λv) {
									return λv
								} else if λv := λ.Cal(λ.GetAttr(ϒtrack, "get", nil), λ.NewStr("label")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewStr("en")
								}
							}()
							λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), ϒlang, λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): ϒtrack_url,
								λ.NewStr("ext"): λ.Cal(λ.GetAttr(λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("scc"):     λ.NewStr("scc"),
									λ.NewStr("webvtt"):  λ.NewStr("vtt"),
									λ.NewStr("smptett"): λ.NewStr("tt"),
								}), "get", nil), λ.Cal(λ.GetAttr(ϒsource, "get", nil), λ.NewStr("format"))),
							}))
						}
					}
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒimage λ.Object
									τmp0   λ.Object
									τmp1   λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "findall", nil), λ.NewStr("images/image")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒimage = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("id"):     λ.Cal(λ.GetAttr(ϒimage, "get", nil), λ.NewStr("cut")),
										λ.NewStr("url"):    λ.GetAttr(ϒimage, "text", nil),
										λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒimage, "get", nil), λ.NewStr("width"))),
										λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒimage, "get", nil), λ.NewStr("height"))),
									}))
								}
								return λ.None
							})
						})))
					ϒis_live = λ.Eq(λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("isLive")), λ.NewStr("true"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): ϒvideo_id,
						λ.NewStr("title"): func() λ.Object {
							if λ.IsTrue(ϒis_live) {
								return λ.Cal(λ.GetAttr(ϒself, "_live_title", nil), ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("subtitles"):   ϒsubtitles,
						λ.NewStr("thumbnails"):  ϒthumbnails,
						λ.NewStr("thumbnail"):   λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("poster")),
						λ.NewStr("description"): λ.Cal(ϒstrip_or_none, λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("description"))),
						λ.NewStr("duration"): λ.Cal(ϒparse_duration, func() λ.Object {
							if λv := λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("length")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("trt"))
							}
						}()),
						λ.NewStr("timestamp"):      λ.Cal(λ.GetAttr(ϒself, "_extract_timestamp", nil), ϒvideo_data),
						λ.NewStr("upload_date"):    λ.Cal(ϒxpath_attr, ϒvideo_data, λ.NewStr("metas"), λ.NewStr("version")),
						λ.NewStr("series"):         λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("showTitle")),
						λ.NewStr("season_number"):  λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("seasonNumber"))),
						λ.NewStr("episode_number"): λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo_data, λ.NewStr("episodeNumber"))),
						λ.NewStr("is_live"):        ϒis_live,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_extract_cvp_info"):  TurnerBaseIE__extract_cvp_info,
				λ.NewStr("_extract_timestamp"): TurnerBaseIE__extract_timestamp,
			})
		}())
	})
}
