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
 * netzkino/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/netzkino.py
 */

package netzkino

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor  λ.Object
	NetzkinoIE     λ.Object
	ϒclean_html    λ.Object
	ϒint_or_none   λ.Object
	ϒjs_to_json    λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		NetzkinoIE = λ.Cal(λ.TypeType, λ.NewStr("NetzkinoIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NetzkinoIE__TEST         λ.Object
				NetzkinoIE__VALID_URL    λ.Object
				NetzkinoIE__real_extract λ.Object
			)
			NetzkinoIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?netzkino\\.de/\\#!/(?P<category>[^/]+)/(?P<id>[^/]+)")
			NetzkinoIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://www.netzkino.de/#!/scifikino/rakete-zum-mond"),
				λ.NewStr("md5"): λ.NewStr("92a3f8b76f8d7220acce5377ea5d4873"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("rakete-zum-mond"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("Rakete zum Mond (Endstation Mond, Destination Moon)"),
					λ.NewStr("comments"):    λ.NewStr("mincount:3"),
					λ.NewStr("description"): λ.NewStr("md5:1eddeacc7e62d5a25a2d1a7290c64a28"),
					λ.NewStr("upload_date"): λ.NewStr("20120813"),
					λ.NewStr("thumbnail"):   λ.NewStr("re:https?://.*\\.jpg$"),
					λ.NewStr("timestamp"):   λ.NewInt(1344858571),
					λ.NewStr("age_limit"):   λ.NewInt(12),
				}),
				λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("skip_download"): λ.NewStr("Download only works from Germany"),
				}),
			})
			NetzkinoIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_info      λ.Object
						ϒapi_url       λ.Object
						ϒavo_js        λ.Object
						ϒcategory_id   λ.Object
						ϒcomments      λ.Object
						ϒcustom_fields λ.Object
						ϒfilm_fn       λ.Object
						ϒformats       λ.Object
						ϒinfo          λ.Object
						ϒmobj          λ.Object
						ϒproduction_js λ.Object
						ϒself          = λargs[0]
						ϒsuffix        λ.Object
						ϒtemplates     λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒcategory_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("category"))
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒapi_url = λ.Mod(λ.NewStr("http://api.netzkino.de.simplecache.net/capi-2.0a/categories/%s.json?d=www"), ϒcategory_id)
					ϒapi_info = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒapi_url, ϒvideo_id)
					ϒinfo = λ.Cal(λ.BuiltinNext, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒp   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒapi_info, λ.NewStr("posts")))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒp = τmp1
									if λ.IsTrue(λ.Eq(λ.GetItem(ϒp, λ.NewStr("slug")), ϒvideo_id)) {
										λgen.Yield(ϒp)
									}
								}
								return λ.None
							})
						})))
					ϒcustom_fields = λ.GetItem(ϒinfo, λ.NewStr("custom_fields"))
					ϒproduction_js = λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.NewStr("http://www.netzkino.de/beta/dist/production.min.js"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Downloading player code")},
					})
					ϒavo_js = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("var urlTemplate=(\\{.*?\"\\})"), ϒproduction_js, λ.NewStr("URL templates"))
					ϒtemplates = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
						ϒavo_js,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "transform_source", Value: ϒjs_to_json},
					})
					ϒsuffix = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("hds"): λ.NewStr(".mp4/manifest.f4m"),
						λ.NewStr("hls"): λ.NewStr(".mp4/master.m3u8"),
						λ.NewStr("pmd"): λ.NewStr(".mp4"),
					})
					ϒfilm_fn = λ.GetItem(λ.GetItem(ϒcustom_fields, λ.NewStr("Streaming")), λ.NewInt(0))
					ϒformats = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒkey λ.Object
									ϒtpl λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
									τmp2 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒtemplates, "items", nil)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒkey = λ.GetItem(τmp2, λ.NewInt(0))
									ϒtpl = λ.GetItem(τmp2, λ.NewInt(1))
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("format_id"): ϒkey,
										λ.NewStr("ext"):       λ.NewStr("mp4"),
										λ.NewStr("url"):       λ.Add(λ.Cal(λ.GetAttr(ϒtpl, "replace", nil), λ.NewStr("{}"), ϒfilm_fn), λ.GetItem(ϒsuffix, ϒkey)),
									}))
								}
								return λ.None
							})
						})))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒcomments = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒc   λ.Object
									τmp0 λ.Object
									τmp1 λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("comments"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒc = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("timestamp"): λ.Call(ϒparse_iso8601, λ.NewArgs(λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("date"))), λ.KWArgs{
											{Name: "delimiter", Value: λ.NewStr(" ")},
										}),
										λ.NewStr("id"):     λ.GetItem(ϒc, λ.NewStr("id")),
										λ.NewStr("author"): λ.GetItem(ϒc, λ.NewStr("name")),
										λ.NewStr("html"):   λ.GetItem(ϒc, λ.NewStr("content")),
										λ.NewStr("parent"): func() λ.Object {
											if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒc, "get", nil), λ.NewStr("parent"), λ.NewInt(0)), λ.NewInt(0))) {
												return λ.NewStr("root")
											} else {
												return λ.GetItem(ϒc, λ.NewStr("parent"))
											}
										}(),
									}))
								}
								return λ.None
							})
						})))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("formats"):   ϒformats,
						λ.NewStr("comments"):  ϒcomments,
						λ.NewStr("title"):     λ.GetItem(ϒinfo, λ.NewStr("title")),
						λ.NewStr("age_limit"): λ.Cal(ϒint_or_none, λ.GetItem(λ.Cal(λ.GetAttr(ϒcustom_fields, "get", nil), λ.NewStr("FSK")), λ.NewInt(0))),
						λ.NewStr("timestamp"): λ.Call(ϒparse_iso8601, λ.NewArgs(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("date"))), λ.KWArgs{
							{Name: "delimiter", Value: λ.NewStr(" ")},
						}),
						λ.NewStr("description"):    λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("content"))),
						λ.NewStr("thumbnail"):      λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("thumbnail")),
						λ.NewStr("playlist_title"): λ.Cal(λ.GetAttr(ϒapi_info, "get", nil), λ.NewStr("title")),
						λ.NewStr("playlist_id"):    ϒcategory_id,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         NetzkinoIE__TEST,
				λ.NewStr("_VALID_URL"):    NetzkinoIE__VALID_URL,
				λ.NewStr("_real_extract"): NetzkinoIE__real_extract,
			})
		}())
	})
}
