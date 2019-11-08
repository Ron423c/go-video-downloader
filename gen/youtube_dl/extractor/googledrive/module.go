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
 * googledrive/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/googledrive.py
 */

package googledrive

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError    λ.Object
	GoogleDriveIE     λ.Object
	InfoExtractor     λ.Object
	ϒdetermine_ext    λ.Object
	ϒint_or_none      λ.Object
	ϒlowercase_escape λ.Object
	ϒupdate_url_query λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒlowercase_escape = Ωutils.ϒlowercase_escape
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		GoogleDriveIE = λ.Cal(λ.TypeType, λ.NewStr("GoogleDriveIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GoogleDriveIE__FORMATS_EXT  λ.Object
				GoogleDriveIE__VALID_URL    λ.Object
				GoogleDriveIE__real_extract λ.Object
			)
			GoogleDriveIE__VALID_URL = λ.NewStr("(?x)\n                        https?://\n                            (?:\n                                (?:docs|drive)\\.google\\.com/\n                                (?:\n                                    (?:uc|open)\\?.*?id=|\n                                    file/d/\n                                )|\n                                video\\.google\\.com/get_player\\?.*?docid=\n                            )\n                            (?P<id>[a-zA-Z0-9_-]{28,})\n                    ")
			GoogleDriveIE__FORMATS_EXT = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("5"):  λ.NewStr("flv"),
				λ.NewStr("6"):  λ.NewStr("flv"),
				λ.NewStr("13"): λ.NewStr("3gp"),
				λ.NewStr("17"): λ.NewStr("3gp"),
				λ.NewStr("18"): λ.NewStr("mp4"),
				λ.NewStr("22"): λ.NewStr("mp4"),
				λ.NewStr("34"): λ.NewStr("flv"),
				λ.NewStr("35"): λ.NewStr("flv"),
				λ.NewStr("36"): λ.NewStr("3gp"),
				λ.NewStr("37"): λ.NewStr("mp4"),
				λ.NewStr("38"): λ.NewStr("mp4"),
				λ.NewStr("43"): λ.NewStr("webm"),
				λ.NewStr("44"): λ.NewStr("webm"),
				λ.NewStr("45"): λ.NewStr("webm"),
				λ.NewStr("46"): λ.NewStr("webm"),
				λ.NewStr("59"): λ.NewStr("mp4"),
			})
			GoogleDriveIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒadd_source_format    λ.Object
						ϒconfirm              λ.Object
						ϒconfirmation_webpage λ.Object
						ϒduration             λ.Object
						ϒf                    λ.Object
						ϒfmt                  λ.Object
						ϒfmt_list             λ.Object
						ϒfmt_stream           λ.Object
						ϒfmt_stream_map       λ.Object
						ϒfmt_stream_split     λ.Object
						ϒformat_id            λ.Object
						ϒformat_url           λ.Object
						ϒformats              λ.Object
						ϒhl                   λ.Object
						ϒmobj                 λ.Object
						ϒreason               λ.Object
						ϒresolution           λ.Object
						ϒresolutions          λ.Object
						ϒself                 = λargs[0]
						ϒsource_url           λ.Object
						ϒsubtitles_id         λ.Object
						ϒtitle                λ.Object
						ϒttsurl               λ.Object
						ϒurl                  = λargs[1]
						ϒurlh                 λ.Object
						ϒvideo_id             λ.Object
						ϒwebpage              λ.Object
						τmp0                  λ.Object
						τmp1                  λ.Object
						τmp2                  λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://docs.google.com/file/d/%s"), ϒvideo_id), ϒvideo_id)
					ϒtitle = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("\"title\"\\s*,\\s*\"([^\"]+)"),
							ϒwebpage,
							λ.NewStr("title"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒduration = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("\"length_seconds\"\\s*,\\s*\"([^\"]+)"),
						ϒwebpage,
						λ.NewStr("length seconds"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒformats = λ.NewList()
					ϒfmt_stream_map = λ.Cal(λ.GetAttr(λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("\"fmt_stream_map\"\\s*,\\s*\"([^\"]+)"),
						ϒwebpage,
						λ.NewStr("fmt stream map"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("")},
					}), "split", nil), λ.NewStr(","))
					ϒfmt_list = λ.Cal(λ.GetAttr(λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("\"fmt_list\"\\s*,\\s*\"([^\"]+)"),
						ϒwebpage,
						λ.NewStr("fmt_list"),
					), λ.KWArgs{
						{Name: "default", Value: λ.NewStr("")},
					}), "split", nil), λ.NewStr(","))
					if λ.IsTrue(func() λ.Object {
						if λv := ϒfmt_stream_map; !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒfmt_list
						}
					}()) {
						ϒresolutions = λ.NewDictWithTable(map[λ.Object]λ.Object{})
						τmp0 = λ.Cal(λ.BuiltinIter, ϒfmt_list)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfmt = τmp1
							ϒmobj = λ.Cal(Ωre.ϒsearch, λ.NewStr("^(?P<format_id>\\d+)/(?P<width>\\d+)[xX](?P<height>\\d+)"), ϒfmt)
							if λ.IsTrue(ϒmobj) {
								λ.SetItem(ϒresolutions, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("format_id")), λ.NewTuple(
									λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("width"))),
									λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("height"))),
								))
							}
						}
						τmp0 = λ.Cal(λ.BuiltinIter, ϒfmt_stream_map)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfmt_stream = τmp1
							ϒfmt_stream_split = λ.Cal(λ.GetAttr(ϒfmt_stream, "split", nil), λ.NewStr("|"))
							if λ.IsTrue(λ.Lt(λ.Cal(λ.BuiltinLen, ϒfmt_stream_split), λ.NewInt(2))) {
								continue
							}
							τmp2 = λ.GetItem(ϒfmt_stream_split, λ.NewSlice(λ.None, λ.NewInt(2), λ.None))
							ϒformat_id = λ.GetItem(τmp2, λ.NewInt(0))
							ϒformat_url = λ.GetItem(τmp2, λ.NewInt(1))
							ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       λ.Cal(ϒlowercase_escape, ϒformat_url),
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("ext"):       λ.GetItem(λ.GetAttr(ϒself, "_FORMATS_EXT", nil), ϒformat_id),
							})
							ϒresolution = λ.Cal(λ.GetAttr(ϒresolutions, "get", nil), ϒformat_id)
							if λ.IsTrue(ϒresolution) {
								λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("width"):  λ.GetItem(ϒresolution, λ.NewInt(0)),
									λ.NewStr("height"): λ.GetItem(ϒresolution, λ.NewInt(1)),
								}))
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
						}
					}
					ϒsource_url = λ.Cal(ϒupdate_url_query, λ.NewStr("https://drive.google.com/uc"), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):     ϒvideo_id,
						λ.NewStr("export"): λ.NewStr("download"),
					}))
					ϒurlh = λ.Call(λ.GetAttr(ϒself, "_request_webpage", nil), λ.NewArgs(
						ϒsource_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: λ.NewStr("Requesting source file")},
						{Name: "errnote", Value: λ.NewStr("Unable to request source file")},
						{Name: "fatal", Value: λ.False},
					})
					if λ.IsTrue(ϒurlh) {
						ϒadd_source_format = λ.NewFunction("add_source_format",
							[]λ.Param{
								{Name: "src_url"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒsrc_url = λargs[0]
								)
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("url"):       ϒsrc_url,
									λ.NewStr("ext"):       λ.Cal(λ.GetAttr(λ.Cal(ϒdetermine_ext, ϒtitle, λ.NewStr("mp4")), "lower", nil)),
									λ.NewStr("format_id"): λ.NewStr("source"),
									λ.NewStr("quality"):   λ.NewInt(1),
								}))
								return λ.None
							})
						if λ.IsTrue(λ.Cal(λ.GetAttr(λ.GetAttr(ϒurlh, "headers", nil), "get", nil), λ.NewStr("Content-Disposition"))) {
							λ.Cal(ϒadd_source_format, ϒsource_url)
						} else {
							ϒconfirmation_webpage = λ.Call(λ.GetAttr(ϒself, "_webpage_read_content", nil), λ.NewArgs(
								ϒurlh,
								ϒurl,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "note", Value: λ.NewStr("Downloading confirmation page")},
								{Name: "errnote", Value: λ.NewStr("Unable to confirm download")},
								{Name: "fatal", Value: λ.False},
							})
							if λ.IsTrue(ϒconfirmation_webpage) {
								ϒconfirm = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.NewStr("confirm=([^&\"\\']+)"),
									ϒconfirmation_webpage,
									λ.NewStr("confirmation code"),
								), λ.KWArgs{
									{Name: "fatal", Value: λ.False},
								})
								if λ.IsTrue(ϒconfirm) {
									λ.Cal(ϒadd_source_format, λ.Cal(ϒupdate_url_query, ϒsource_url, λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("confirm"): ϒconfirm,
									})))
								}
							}
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformats))) {
						ϒreason = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("\"reason\"\\s*,\\s*\"([^\"]+)"),
							ϒwebpage,
							λ.NewStr("reason"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒreason) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(ϒreason), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒhl = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("\"hl\"\\s*,\\s*\"([^\"]+)"),
						ϒwebpage,
						λ.NewStr("hl"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒsubtitles_id = λ.None
					ϒttsurl = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("\"ttsurl\"\\s*,\\s*\"([^\"]+)"),
						ϒwebpage,
						λ.NewStr("ttsurl"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒttsurl) {
						ϒsubtitles_id = λ.GetItem(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒttsurl, "encode", nil), λ.NewStr("utf-8")), "decode", nil), λ.NewStr("unicode_escape")), "split", nil), λ.NewStr("=")), λ.Neg(λ.NewInt(1)))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    ϒvideo_id,
						λ.NewStr("title"): ϒtitle,
						λ.NewStr("thumbnail"): λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}),
						λ.NewStr("duration"):           ϒduration,
						λ.NewStr("formats"):            ϒformats,
						λ.NewStr("subtitles"):          λ.Cal(λ.GetAttr(ϒself, "extract_subtitles", nil), ϒvideo_id, ϒsubtitles_id, ϒhl),
						λ.NewStr("automatic_captions"): λ.Cal(λ.GetAttr(ϒself, "extract_automatic_captions", nil), ϒvideo_id, ϒsubtitles_id, ϒhl),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_FORMATS_EXT"):  GoogleDriveIE__FORMATS_EXT,
				λ.NewStr("_VALID_URL"):    GoogleDriveIE__VALID_URL,
				λ.NewStr("_real_extract"): GoogleDriveIE__real_extract,
			})
		}())
	})
}
