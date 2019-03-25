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
 * packtpub/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/packtpub.py
 */

package packtpub

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError     λ.Object
	InfoExtractor      λ.Object
	PacktPubBaseIE     λ.Object
	PacktPubCourseIE   λ.Object
	PacktPubIE         λ.Object
	ϒclean_html        λ.Object
	ϒcompat_HTTPError  λ.Object
	ϒcompat_str        λ.Object
	ϒremove_end        λ.Object
	ϒstrip_or_none     λ.Object
	ϒunified_timestamp λ.Object
	ϒurljoin           λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_HTTPError = Ωcompat.ϒcompat_HTTPError
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒremove_end = Ωutils.ϒremove_end
		ϒstrip_or_none = Ωutils.ϒstrip_or_none
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒurljoin = Ωutils.ϒurljoin
		PacktPubBaseIE = λ.Cal(λ.TypeType, λ.NewStr("PacktPubBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PacktPubBaseIE__MAPT_REST  λ.Object
				PacktPubBaseIE__PACKT_BASE λ.Object
			)
			PacktPubBaseIE__PACKT_BASE = λ.NewStr("https://www.packtpub.com")
			PacktPubBaseIE__MAPT_REST = λ.Mod(λ.NewStr("%s/mapt-rest"), PacktPubBaseIE__PACKT_BASE)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_MAPT_REST"):  PacktPubBaseIE__MAPT_REST,
				λ.NewStr("_PACKT_BASE"): PacktPubBaseIE__PACKT_BASE,
			})
		}())
		PacktPubIE = λ.Cal(λ.TypeType, λ.NewStr("PacktPubIE"), λ.NewTuple(PacktPubBaseIE), func() λ.Dict {
			var (
				PacktPubIE__NETRC_MACHINE   λ.Object
				PacktPubIE__TESTS           λ.Object
				PacktPubIE__TOKEN           λ.Object
				PacktPubIE__VALID_URL       λ.Object
				PacktPubIE__download_json   λ.Object
				PacktPubIE__handle_error    λ.Object
				PacktPubIE__real_extract    λ.Object
				PacktPubIE__real_initialize λ.Object
			)
			PacktPubIE__VALID_URL = λ.NewStr("https?://(?:(?:www\\.)?packtpub\\.com/mapt|subscription\\.packtpub\\.com)/video/[^/]+/(?P<course_id>\\d+)/(?P<chapter_id>\\d+)/(?P<id>\\d+)")
			PacktPubIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.packtpub.com/mapt/video/web-development/9781787122215/20528/20530/Project+Intro"),
					λ.NewStr("md5"): λ.NewStr("1e74bd6cfd45d7d07666f4684ef58f70"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("20530"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Project Intro"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:(?i)^https?://.*\\.jpg"),
						λ.NewStr("timestamp"):   λ.NewInt(1490918400),
						λ.NewStr("upload_date"): λ.NewStr("20170331"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://subscription.packtpub.com/video/web_development/9781787122215/20528/20530/project-intro"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			PacktPubIE__NETRC_MACHINE = λ.NewStr("packtpub")
			PacktPubIE__TOKEN = λ.None
			PacktPubIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmessage  λ.Object
						ϒpassword λ.Object
						ϒself     = λargs[0]
						ϒusername λ.Object
						τmp0      λ.Object
						τmp1      λ.Object
					)
					_ = τmp0
					_ = τmp1
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_login_info", nil))
					ϒusername = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpassword = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(ϒusername == λ.None)) {
						return λ.None
					}
					τmp0, τmp1 = func() (λexit λ.Object, λret λ.Object) {
						defer λ.CatchMulti(
							nil,
							&λ.Catcher{ExtractorError, func(λex λ.BaseException) {
								var ϒe λ.Object = λex
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Cal(λ.BuiltinIsInstance, λ.GetAttr(ϒe, "cause", nil), ϒcompat_HTTPError); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(λ.NewTuple(
											λ.NewInt(400),
											λ.NewInt(401),
											λ.NewInt(404),
										), λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "code", nil)))
									}
								}()) {
									ϒmessage = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.GetAttr(ϒe, "cause", nil), "read", nil)), "decode", nil)), λ.None), λ.NewStr("message"))
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒmessage), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
								panic(λ.Raise(λex))
							}},
						)
						λ.SetAttr(ϒself, "_TOKEN", λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
							λ.Add(λ.GetAttr(ϒself, "_MAPT_REST", nil), λ.NewStr("/users/tokens")),
							λ.None,
							λ.NewStr("Downloading Authorization Token"),
						), λ.KWArgs{
							{Name: "data", Value: λ.Cal(λ.GetAttr(λ.Cal(Ωjson.ϒdumps, λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("email"):    ϒusername,
								λ.NewStr("password"): ϒpassword,
							})), "encode", nil))},
						}), λ.NewStr("data")), λ.NewStr("access")))
						return λ.BlockExitNormally, nil
					}()
					return λ.None
				})
			PacktPubIE__handle_error = λ.NewFunction("_handle_error",
				[]λ.Param{
					{Name: "self"},
					{Name: "response"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒresponse = λargs[1]
						ϒself     = λargs[0]
					)
					if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒresponse, "get", nil), λ.NewStr("status")), λ.NewStr("success"))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("% said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							λ.GetItem(ϒresponse, λ.NewStr("message")),
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			PacktPubIE__download_json = λ.NewFunction("_download_json",
				[]λ.Param{
					{Name: "self"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs     = λargs[1]
						ϒkwargs   = λargs[2]
						ϒresponse λ.Object
						ϒself     = λargs[0]
					)
					ϒresponse = λ.Call(λ.GetAttr(λ.Cal(λ.SuperType, PacktPubIE, ϒself), "_download_json", nil), λ.NewArgs(λ.Unpack(λ.AsStarred(ϒargs))...), λ.KWArgs{
						{Name: "", Value: ϒkwargs},
					})
					λ.Cal(λ.GetAttr(ϒself, "_handle_error", nil), ϒresponse)
					return ϒresponse
				})
			PacktPubIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchapter_id   λ.Object
						ϒcontent      λ.Object
						ϒcourse_id    λ.Object
						ϒcourse_title λ.Object
						ϒheaders      λ.Object
						ϒmetadata     λ.Object
						ϒmobj         λ.Object
						ϒself         = λargs[0]
						ϒthumbnail    λ.Object
						ϒtimestamp    λ.Object
						ϒtitle        λ.Object
						ϒurl          = λargs[1]
						ϒvideo        λ.Object
						ϒvideo_id     λ.Object
						ϒvideo_url    λ.Object
						τmp0          λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					τmp0 = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("course_id"), λ.NewStr("chapter_id"), λ.NewStr("id"))
					ϒcourse_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒchapter_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(2))
					ϒheaders = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					if λ.IsTrue(λ.GetAttr(ϒself, "_TOKEN", nil)) {
						λ.SetItem(ϒheaders, λ.NewStr("Authorization"), λ.Add(λ.NewStr("Bearer "), λ.GetAttr(ϒself, "_TOKEN", nil)))
					}
					ϒvideo = λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Mod(λ.NewStr("%s/users/me/products/%s/chapters/%s/sections/%s"), λ.NewTuple(
							λ.GetAttr(ϒself, "_MAPT_REST", nil),
							ϒcourse_id,
							ϒchapter_id,
							ϒvideo_id,
						)),
						ϒvideo_id,
						λ.NewStr("Downloading JSON video"),
					), λ.KWArgs{
						{Name: "headers", Value: ϒheaders},
					}), λ.NewStr("data"))
					ϒcontent = λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("content"))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒcontent))) {
						λ.Cal(λ.GetAttr(ϒself, "raise_login_required", nil), λ.NewStr("This video is locked"))
					}
					ϒvideo_url = λ.GetItem(ϒcontent, λ.NewStr("file"))
					ϒmetadata = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("%s/products/%s/chapters/%s/sections/%s/metadata"), λ.NewTuple(
						λ.GetAttr(ϒself, "_MAPT_REST", nil),
						ϒcourse_id,
						ϒchapter_id,
						ϒvideo_id,
					)), ϒvideo_id), λ.NewStr("data"))
					ϒtitle = λ.GetItem(ϒmetadata, λ.NewStr("pageTitle"))
					ϒcourse_title = λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("title"))
					if λ.IsTrue(ϒcourse_title) {
						ϒtitle = λ.Cal(ϒremove_end, ϒtitle, λ.Mod(λ.NewStr(" - %s"), ϒcourse_title))
					}
					ϒtimestamp = λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("publicationDate")))
					ϒthumbnail = λ.Cal(ϒurljoin, λ.GetAttr(ϒself, "_PACKT_BASE", nil), λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("filepath")))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("url"):       ϒvideo_url,
						λ.NewStr("title"):     ϒtitle,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("timestamp"): ϒtimestamp,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_NETRC_MACHINE"):   PacktPubIE__NETRC_MACHINE,
				λ.NewStr("_TESTS"):           PacktPubIE__TESTS,
				λ.NewStr("_TOKEN"):           PacktPubIE__TOKEN,
				λ.NewStr("_VALID_URL"):       PacktPubIE__VALID_URL,
				λ.NewStr("_download_json"):   PacktPubIE__download_json,
				λ.NewStr("_handle_error"):    PacktPubIE__handle_error,
				λ.NewStr("_real_extract"):    PacktPubIE__real_extract,
				λ.NewStr("_real_initialize"): PacktPubIE__real_initialize,
			})
		}())
		PacktPubCourseIE = λ.Cal(λ.TypeType, λ.NewStr("PacktPubCourseIE"), λ.NewTuple(PacktPubBaseIE), func() λ.Dict {
			var (
				PacktPubCourseIE__VALID_URL λ.Object
				PacktPubCourseIE_suitable   λ.Object
			)
			PacktPubCourseIE__VALID_URL = λ.NewStr("(?P<url>https?://(?:(?:www\\.)?packtpub\\.com/mapt|subscription\\.packtpub\\.com)/video/[^/]+/(?P<id>\\d+))")
			PacktPubCourseIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Cal(λ.GetAttr(PacktPubIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, PacktPubCourseIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			PacktPubCourseIE_suitable = λ.Cal(λ.ClassMethodType, PacktPubCourseIE_suitable)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): PacktPubCourseIE__VALID_URL,
				λ.NewStr("suitable"):   PacktPubCourseIE_suitable,
			})
		}())
	})
}
