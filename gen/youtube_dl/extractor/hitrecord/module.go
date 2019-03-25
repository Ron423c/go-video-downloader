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
 * hitrecord/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/hitrecord.py
 */

package hitrecord

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	HitRecordIE    λ.Object
	InfoExtractor  λ.Object
	ϒclean_html    λ.Object
	ϒcompat_str    λ.Object
	ϒfloat_or_none λ.Object
	ϒint_or_none   λ.Object
	ϒtry_get       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		HitRecordIE = λ.Cal(λ.TypeType, λ.NewStr("HitRecordIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				HitRecordIE__TEST         λ.Object
				HitRecordIE__VALID_URL    λ.Object
				HitRecordIE__real_extract λ.Object
			)
			HitRecordIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?hitrecord\\.org/records/(?P<id>\\d+)")
			HitRecordIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://hitrecord.org/records/2954362"),
				λ.NewStr("md5"): λ.NewStr("fe1cdc2023bce0bbb95c39c57426aa71"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):            λ.NewStr("2954362"),
					λ.NewStr("ext"):           λ.NewStr("mp4"),
					λ.NewStr("title"):         λ.NewStr("A Very Different World (HITRECORD x ACLU)"),
					λ.NewStr("description"):   λ.NewStr("md5:e62defaffab5075a5277736bead95a3d"),
					λ.NewStr("duration"):      λ.NewFloat(139.327),
					λ.NewStr("timestamp"):     λ.NewInt(1471557582),
					λ.NewStr("upload_date"):   λ.NewStr("20160818"),
					λ.NewStr("uploader"):      λ.NewStr("Zuzi.C12"),
					λ.NewStr("uploader_id"):   λ.NewStr("362811"),
					λ.NewStr("view_count"):    λ.IntType,
					λ.NewStr("like_count"):    λ.IntType,
					λ.NewStr("comment_count"): λ.IntType,
					λ.NewStr("tags"):          λ.ListType,
				}),
			})
			HitRecordIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself      = λargs[0]
						ϒtags      λ.Object
						ϒtags_list λ.Object
						ϒtitle     λ.Object
						ϒurl       = λargs[1]
						ϒvideo     λ.Object
						ϒvideo_id  λ.Object
						ϒvideo_url λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒvideo = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://hitrecord.org/api/web/records/%s"), ϒvideo_id), ϒvideo_id)
					ϒtitle = λ.GetItem(ϒvideo, λ.NewStr("title"))
					ϒvideo_url = λ.GetItem(λ.GetItem(ϒvideo, λ.NewStr("source_url")), λ.NewStr("mp4_url"))
					ϒtags = λ.None
					ϒtags_list = λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
						[]λ.Param{
							{Name: "x"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒx = λargs[0]
							)
							return λ.GetItem(ϒx, λ.NewStr("tags"))
						}), λ.ListType)
					if λ.IsTrue(ϒtags_list) {
						ϒtags = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒt   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒtags_list)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒt = τmp1
										if λ.IsTrue(func() λ.Object {
											if λv := λ.Cal(λ.BuiltinIsInstance, ϒt, λ.DictType); !λ.IsTrue(λv) {
												return λv
											} else if λv := λ.Cal(λ.GetAttr(ϒt, "get", nil), λ.NewStr("text")); !λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Cal(λ.BuiltinIsInstance, λ.GetItem(ϒt, λ.NewStr("text")), ϒcompat_str)
											}
										}()) {
											λgen.Yield(λ.GetItem(ϒt, λ.NewStr("text")))
										}
									}
									return λ.None
								})
							})))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("body"))),
						λ.NewStr("duration"):    λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("duration")), λ.NewInt(1000)),
						λ.NewStr("timestamp"):   λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("created_at_i"))),
						λ.NewStr("uploader"): λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("user")), λ.NewStr("username"))
							}), ϒcompat_str),
						λ.NewStr("uploader_id"): λ.Cal(ϒtry_get, ϒvideo, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.Cal(ϒcompat_str, λ.GetItem(λ.GetItem(ϒx, λ.NewStr("user")), λ.NewStr("id")))
							})),
						λ.NewStr("view_count"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("total_views_count"))),
						λ.NewStr("like_count"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("hearts_count"))),
						λ.NewStr("comment_count"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("comments_count"))),
						λ.NewStr("tags"):          ϒtags,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         HitRecordIE__TEST,
				λ.NewStr("_VALID_URL"):    HitRecordIE__VALID_URL,
				λ.NewStr("_real_extract"): HitRecordIE__real_extract,
			})
		}())
	})
}
