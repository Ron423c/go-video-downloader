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
 * lynda/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/lynda.py
 */

package lynda

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError      λ.Object
	InfoExtractor       λ.Object
	LyndaBaseIE         λ.Object
	LyndaCourseIE       λ.Object
	LyndaIE             λ.Object
	ϒcompat_str         λ.Object
	ϒint_or_none        λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		LyndaBaseIE = λ.Cal(λ.TypeType, λ.NewStr("LyndaBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				LyndaBaseIE__NETRC_MACHINE   λ.Object
				LyndaBaseIE__login           λ.Object
				LyndaBaseIE__real_initialize λ.Object
			)
			LyndaBaseIE__NETRC_MACHINE = λ.NewStr("lynda")
			LyndaBaseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Cal(λ.GetAttr(ϒself, "_login", nil))
					return λ.None
				})
			LyndaBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒpassword      λ.Object
						ϒpassword_form λ.Object
						ϒself          = λargs[0]
						ϒsignin_form   λ.Object
						ϒsignin_page   λ.Object
						ϒsignin_url    λ.Object
						ϒusername      λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_login_info", nil))
					ϒusername = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpassword = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(ϒusername == λ.None)) {
						return λ.None
					}
					ϒsignin_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.GetAttr(ϒself, "_SIGNIN_URL", nil), λ.None, λ.NewStr("Downloading signin page"))
					if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒp   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
									λ.NewStr("isLoggedIn\\s*:\\s*true"),
									λ.NewStr("logout\\.aspx"),
									λ.NewStr(">Log out<"),
								))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒp = τmp1
									λgen.Yield(λ.Cal(Ωre.ϒsearch, ϒp, ϒsignin_page))
								}
								return λ.None
							})
						})))) {
						return λ.None
					}
					ϒsignin_form = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)(<form[^>]+data-form-name=[\"\\']signin[\"\\'][^>]*>.+?</form>)"), ϒsignin_page, λ.NewStr("signin form"))
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_login_step", nil), ϒsignin_form, λ.GetAttr(ϒself, "_PASSWORD_URL", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("email"): ϒusername,
					}), λ.NewStr("Submitting email"), λ.GetAttr(ϒself, "_SIGNIN_URL", nil))
					ϒsignin_page = λ.GetItem(τmp0, λ.NewInt(0))
					ϒsignin_url = λ.GetItem(τmp0, λ.NewInt(1))
					ϒpassword_form = λ.GetItem(ϒsignin_page, λ.NewStr("body"))
					λ.Cal(λ.GetAttr(ϒself, "_login_step", nil), ϒpassword_form, λ.GetAttr(ϒself, "_USER_URL", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("email"):    ϒusername,
						λ.NewStr("password"): ϒpassword,
					}), λ.NewStr("Submitting password"), ϒsignin_url)
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_NETRC_MACHINE"):   LyndaBaseIE__NETRC_MACHINE,
				λ.NewStr("_login"):           LyndaBaseIE__login,
				λ.NewStr("_real_initialize"): LyndaBaseIE__real_initialize,
			})
		}())
		LyndaIE = λ.Cal(λ.TypeType, λ.NewStr("LyndaIE"), λ.NewTuple(LyndaBaseIE), func() λ.Dict {
			var (
				LyndaIE_IE_NAME       λ.Object
				LyndaIE__TESTS        λ.Object
				LyndaIE__VALID_URL    λ.Object
				LyndaIE__real_extract λ.Object
			)
			LyndaIE_IE_NAME = λ.NewStr("lynda")
			LyndaIE__VALID_URL = λ.NewStr("(?x)\n                    https?://\n                        (?:www\\.)?(?:lynda\\.com|educourse\\.ga)/\n                        (?:\n                            (?:[^/]+/){2,3}(?P<course_id>\\d+)|\n                            player/embed\n                        )/\n                        (?P<id>\\d+)\n                    ")
			LyndaIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.lynda.com/Bootstrap-tutorials/Using-exercise-files/110885/114408-4.html"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):       λ.NewStr("114408"),
						λ.NewStr("ext"):      λ.NewStr("mp4"),
						λ.NewStr("title"):    λ.NewStr("Using the exercise files"),
						λ.NewStr("duration"): λ.NewInt(68),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.lynda.com/player/embed/133770?tr=foo=1;bar=g;fizz=rt&fs=0"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://educourse.ga/Bootstrap-tutorials/Using-exercise-files/110885/114408-4.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.lynda.com/de/Graphic-Design-tutorials/Willkommen-Grundlagen-guten-Gestaltung/393570/393572-4.html"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			LyndaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcdn                   λ.Object
						ϒconviva               λ.Object
						ϒcourse_id             λ.Object
						ϒduration              λ.Object
						ϒfmts                  λ.Object
						ϒformat_id             λ.Object
						ϒformat_url            λ.Object
						ϒformats               λ.Object
						ϒformats_dict          λ.Object
						ϒmobj                  λ.Object
						ϒplay                  λ.Object
						ϒprioritized_stream    λ.Object
						ϒprioritized_stream_id λ.Object
						ϒprioritized_streams   λ.Object
						ϒquery                 λ.Object
						ϒself                  = λargs[0]
						ϒsubtitles             λ.Object
						ϒtitle                 λ.Object
						ϒurl                   = λargs[1]
						ϒurls                  λ.Object
						ϒvideo                 λ.Object
						ϒvideo_id              λ.Object
						τmp0                   λ.Object
						τmp1                   λ.Object
						τmp2                   λ.Object
						τmp3                   λ.Object
						τmp4                   λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒcourse_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course_id"))
					ϒquery = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("videoId"): ϒvideo_id,
						λ.NewStr("type"):    λ.NewStr("video"),
					})
					ϒvideo = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.NewStr("https://www.lynda.com/ajax/player"),
						ϒvideo_id,
						λ.NewStr("Downloading video JSON"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
						{Name: "query", Value: ϒquery},
					})
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo))) {
						λ.SetItem(ϒquery, λ.NewStr("courseId"), ϒcourse_id)
						ϒplay = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://www.lynda.com/ajax/course/%s/%s/play"), λ.NewTuple(
							ϒcourse_id,
							ϒvideo_id,
						)), ϒvideo_id, λ.NewStr("Downloading play JSON"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒplay))) {
							λ.Cal(λ.GetAttr(ϒself, "_raise_unavailable", nil), ϒvideo_id)
						}
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, ϒplay)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒformats_dict = τmp1
							ϒurls = λ.Cal(λ.GetAttr(ϒformats_dict, "get", nil), λ.NewStr("urls"))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒurls, λ.DictType)))) {
								continue
							}
							ϒcdn = λ.Cal(λ.GetAttr(ϒformats_dict, "get", nil), λ.NewStr("name"))
							τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒurls, "items", nil)))
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								τmp4 = τmp3
								ϒformat_id = λ.GetItem(τmp4, λ.NewInt(0))
								ϒformat_url = λ.GetItem(τmp4, λ.NewInt(1))
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformat_url))) {
									continue
								}
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"): ϒformat_url,
									λ.NewStr("format_id"): func() λ.Object {
										if λ.IsTrue(ϒcdn) {
											return λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
												ϒcdn,
												ϒformat_id,
											))
										} else {
											return ϒformat_id
										}
									}(),
									λ.NewStr("height"): λ.Cal(ϒint_or_none, ϒformat_id),
								}))
							}
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
						ϒconviva = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.NewStr("https://www.lynda.com/ajax/player/conviva"),
							ϒvideo_id,
							λ.NewStr("Downloading conviva JSON"),
						), λ.KWArgs{
							{Name: "query", Value: ϒquery},
						})
						return λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):           ϒvideo_id,
							λ.NewStr("title"):        λ.GetItem(ϒconviva, λ.NewStr("VideoTitle")),
							λ.NewStr("description"):  λ.Cal(λ.GetAttr(ϒconviva, "get", nil), λ.NewStr("VideoDescription")),
							λ.NewStr("release_year"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒconviva, "get", nil), λ.NewStr("ReleaseYear"))),
							λ.NewStr("duration"):     λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒconviva, "get", nil), λ.NewStr("Duration"))),
							λ.NewStr("creator"):      λ.Cal(λ.GetAttr(ϒconviva, "get", nil), λ.NewStr("Author")),
							λ.NewStr("formats"):      ϒformats,
						})
					}
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒvideo, λ.NewStr("Status")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("lynda returned error: %s"), λ.GetItem(ϒvideo, λ.NewStr("Message")))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					if λ.IsTrue(λ.NewBool(λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("HasAccess")) == λ.False)) {
						λ.Cal(λ.GetAttr(ϒself, "_raise_unavailable", nil), ϒvideo_id)
					}
					ϒvideo_id = λ.Cal(ϒcompat_str, func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("ID")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}())
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("DurationInSeconds")))
					ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("Title"))
					ϒformats = λ.NewList()
					ϒfmts = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("Formats"))
					if λ.IsTrue(ϒfmts) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒf   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒfmts)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒf = τmp1
										if λ.IsTrue(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("Url"))) {
											λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("url"):      λ.GetItem(ϒf, λ.NewStr("Url")),
												λ.NewStr("ext"):      λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("Extension")),
												λ.NewStr("width"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("Width"))),
												λ.NewStr("height"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("Height"))),
												λ.NewStr("filesize"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("FileSize"))),
												λ.NewStr("format_id"): func() λ.Object {
													if λ.IsTrue(λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("Resolution"))) {
														return λ.Cal(ϒcompat_str, λ.Cal(λ.GetAttr(ϒf, "get", nil), λ.NewStr("Resolution")))
													} else {
														return λ.None
													}
												}(),
											}))
										}
									}
									return λ.None
								})
							}))))
					}
					ϒprioritized_streams = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("PrioritizedStreams"))
					if λ.IsTrue(ϒprioritized_streams) {
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒprioritized_streams, "items", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒprioritized_stream_id = λ.GetItem(τmp2, λ.NewInt(0))
							ϒprioritized_stream = λ.GetItem(τmp2, λ.NewInt(1))
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒformat_id λ.Object
											ϒvideo_url λ.Object
											τmp0       λ.Object
											τmp1       λ.Object
											τmp2       λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒprioritized_stream, "items", nil)))
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											τmp2 = τmp1
											ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
											ϒvideo_url = λ.GetItem(τmp2, λ.NewInt(1))
											λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("url"):    ϒvideo_url,
												λ.NewStr("height"): λ.Cal(ϒint_or_none, ϒformat_id),
												λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s-%s"), λ.NewTuple(
													ϒprioritized_stream_id,
													ϒformat_id,
												)),
											}))
										}
										return λ.None
									})
								}))))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_check_formats", nil), ϒformats, ϒvideo_id)
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒsubtitles = λ.Cal(λ.GetAttr(ϒself, "extract_subtitles", nil), ϒvideo_id)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("duration"):  ϒduration,
						λ.NewStr("subtitles"): ϒsubtitles,
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       LyndaIE_IE_NAME,
				λ.NewStr("_TESTS"):        LyndaIE__TESTS,
				λ.NewStr("_VALID_URL"):    LyndaIE__VALID_URL,
				λ.NewStr("_real_extract"): LyndaIE__real_extract,
			})
		}())
		LyndaCourseIE = λ.Cal(λ.TypeType, λ.NewStr("LyndaCourseIE"), λ.NewTuple(LyndaBaseIE), func() λ.Dict {
			var (
				LyndaCourseIE__VALID_URL λ.Object
			)
			LyndaCourseIE__VALID_URL = λ.NewStr("https?://(?:www|m)\\.(?:lynda\\.com|educourse\\.ga)/(?P<coursepath>(?:[^/]+/){2,3}(?P<courseid>\\d+))-2\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): LyndaCourseIE__VALID_URL,
			})
		}())
	})
}
