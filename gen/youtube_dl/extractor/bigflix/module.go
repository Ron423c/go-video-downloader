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
 * bigflix/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/bigflix.py
 */

package bigflix

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BigflixIE                    λ.Object
	InfoExtractor                λ.Object
	ϒcompat_b64decode            λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_b64decode = Ωcompat.ϒcompat_b64decode
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		BigflixIE = λ.Cal(λ.TypeType, λ.NewStr("BigflixIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				BigflixIE__TESTS        λ.Object
				BigflixIE__VALID_URL    λ.Object
				BigflixIE__real_extract λ.Object
			)
			BigflixIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?bigflix\\.com/.+/(?P<id>[0-9]+)")
			BigflixIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.bigflix.com/Tamil-movies/Drama-movies/Madarasapatinam/16070"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("16070"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Madarasapatinam"),
						λ.NewStr("description"): λ.NewStr("md5:9f0470b26a4ba8e824c823b5d95c2f6b"),
						λ.NewStr("formats"):     λ.NewStr("mincount:2"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.bigflix.com/Malayalam-movies/Drama-movies/Indian-Rupee/15967"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			BigflixIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdecode_url  λ.Object
						ϒdescription λ.Object
						ϒencoded_url λ.Object
						ϒf           λ.Object
						ϒfile_url    λ.Object
						ϒformats     λ.Object
						ϒheight      λ.Object
						ϒself        = λargs[0]
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage     λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<div[^>]+class=[\"\\']pagetitle[\"\\'][^>]*>(.+?)</div>"), ϒwebpage, λ.NewStr("title"))
					ϒdecode_url = λ.NewFunction("decode_url",
						[]λ.Param{
							{Name: "quoted_b64_url"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒquoted_b64_url = λargs[0]
							)
							return λ.Cal(λ.GetAttr(λ.Cal(ϒcompat_b64decode, λ.Cal(ϒcompat_urllib_parse_unquote, ϒquoted_b64_url)), "decode", nil), λ.NewStr("utf-8"))
						})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("ContentURL_(\\d{3,4})[pP][^=]+=([^&]+)"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒheight = λ.GetItem(τmp2, λ.NewInt(0))
						ϒencoded_url = λ.GetItem(τmp2, λ.NewInt(1))
						ϒvideo_url = λ.Cal(ϒdecode_url, ϒencoded_url)
						ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):       ϒvideo_url,
							λ.NewStr("format_id"): λ.Mod(λ.NewStr("%sp"), ϒheight),
							λ.NewStr("height"):    λ.Cal(λ.IntType, ϒheight),
						})
						if λ.IsTrue(λ.Cal(λ.GetAttr(ϒvideo_url, "startswith", nil), λ.NewStr("rtmp"))) {
							λ.SetItem(ϒf, λ.NewStr("ext"), λ.NewStr("flv"))
						}
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
					}
					ϒfile_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("file=([^&]+)"),
						ϒwebpage,
						λ.NewStr("video url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒfile_url) {
						ϒvideo_url = λ.Cal(ϒdecode_url, ϒfile_url)
						if λ.IsTrue(λ.Cal(λ.BuiltinAll, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
									var (
										ϒf   λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, ϒformats)
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒf = τmp1
										λgen.Yield(λ.Ne(λ.GetItem(ϒf, λ.NewStr("url")), ϒvideo_url))
									}
									return λ.None
								})
							})))) {
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"): λ.Cal(ϒdecode_url, ϒfile_url),
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒdescription = λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("description"), ϒwebpage)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        BigflixIE__TESTS,
				λ.NewStr("_VALID_URL"):    BigflixIE__VALID_URL,
				λ.NewStr("_real_extract"): BigflixIE__real_extract,
			})
		}())
	})
}
