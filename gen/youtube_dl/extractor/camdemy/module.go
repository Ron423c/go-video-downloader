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
 * camdemy/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/camdemy.py
 */

package camdemy

import (
	Ωparse "github.com/tenta-browser/go-video-downloader/gen/urllib/parse"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CamdemyFolderIE                λ.Object
	CamdemyIE                      λ.Object
	InfoExtractor                  λ.Object
	ϒclean_html                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒparse_duration                λ.Object
	ϒstr_to_int                    λ.Object
	ϒunified_strdate               λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒclean_html = Ωutils.ϒclean_html
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒunified_strdate = Ωutils.ϒunified_strdate
		CamdemyIE = λ.Cal(λ.TypeType, λ.NewStr("CamdemyIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CamdemyIE__TESTS        λ.Object
				CamdemyIE__VALID_URL    λ.Object
				CamdemyIE__real_extract λ.Object
			)
			CamdemyIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?camdemy\\.com/media/(?P<id>\\d+)")
			CamdemyIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.camdemy.com/media/5181/"),
					λ.NewStr("md5"): λ.NewStr("5a5562b6a98b37873119102e052e311b"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("5181"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Ch1-1 Introduction, Signals (02-23-2012)"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("creator"):     λ.NewStr("ss11spring"),
						λ.NewStr("duration"):    λ.NewInt(1591),
						λ.NewStr("upload_date"): λ.NewStr("20130114"),
						λ.NewStr("view_count"):  λ.IntType,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.camdemy.com/media/13885"),
					λ.NewStr("md5"): λ.NewStr("4576a3bb2581f86c61044822adbd1249"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("13885"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("EverCam + Camdemy QuickStart"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("description"): λ.NewStr("md5:2a9f989c2b153a2342acee579c6e7db6"),
						λ.NewStr("creator"):     λ.NewStr("evercam"),
						λ.NewStr("duration"):    λ.NewInt(318),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.camdemy.com/media/14842"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("2vsYQzNIsJo"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Excel 2013 Tutorial - How to add Password Protection"),
						λ.NewStr("description"): λ.NewStr("Excel 2013 Tutorial for Beginners - How to add Password Protection"),
						λ.NewStr("upload_date"): λ.NewStr("20130211"),
						λ.NewStr("uploader"):    λ.NewStr("Hun Kim"),
						λ.NewStr("uploader_id"): λ.NewStr("hunkimtutorials"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			CamdemyIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdescription   λ.Object
						ϒfile_list_doc λ.Object
						ϒfile_name     λ.Object
						ϒoembed_obj    λ.Object
						ϒself          = λargs[0]
						ϒsrc_from      λ.Object
						ϒthumb_url     λ.Object
						ϒtitle         λ.Object
						ϒupload_date   λ.Object
						ϒurl           = λargs[1]
						ϒvideo_folder  λ.Object
						ϒvideo_id      λ.Object
						ϒvideo_url     λ.Object
						ϒview_count    λ.Object
						ϒwebpage       λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒsrc_from = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("class=['\\\"]srcFrom['\\\"][^>]*>Sources?(?:\\s+from)?\\s*:\\s*<a[^>]+(?:href|title)=(['\\\"])(?P<url>(?:(?!\\1).)+)\\1"),
						ϒwebpage,
						λ.NewStr("external source"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.NewStr("url")},
					})
					if λ.IsTrue(ϒsrc_from) {
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), ϒsrc_from)
					}
					ϒoembed_obj = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Add(λ.NewStr("http://www.camdemy.com/oembed/?format=json&url="), ϒurl), ϒvideo_id)
					ϒtitle = λ.GetItem(ϒoembed_obj, λ.NewStr("title"))
					ϒthumb_url = λ.GetItem(ϒoembed_obj, λ.NewStr("thumbnail_url"))
					ϒvideo_folder = λ.Cal(Ωparse.ϒurljoin, ϒthumb_url, λ.NewStr("video/"))
					ϒfile_list_doc = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Cal(Ωparse.ϒurljoin, ϒvideo_folder, λ.NewStr("fileList.xml")), ϒvideo_id, λ.NewStr("Downloading filelist XML"))
					ϒfile_name = λ.GetAttr(λ.Cal(λ.GetAttr(ϒfile_list_doc, "find", nil), λ.NewStr("./video/item/fileName")), "text", nil)
					ϒvideo_url = λ.Cal(Ωparse.ϒurljoin, ϒvideo_folder, ϒfile_name)
					ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr(">published on ([^<]+)<"),
						ϒwebpage,
						λ.NewStr("upload date"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒview_count = λ.Cal(ϒstr_to_int, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("role=[\"\\']viewCnt[\"\\'][^>]*>([\\d,.]+) views"),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("description"),
							ϒwebpage,
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒoembed_obj, "get", nil), λ.NewStr("description")))
						}
					}()
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("url"):         ϒvideo_url,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("thumbnail"):   ϒthumb_url,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("creator"):     λ.Cal(λ.GetAttr(ϒoembed_obj, "get", nil), λ.NewStr("author_name")),
						λ.NewStr("duration"):    λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒoembed_obj, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("upload_date"): ϒupload_date,
						λ.NewStr("view_count"):  ϒview_count,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        CamdemyIE__TESTS,
				λ.NewStr("_VALID_URL"):    CamdemyIE__VALID_URL,
				λ.NewStr("_real_extract"): CamdemyIE__real_extract,
			})
		}())
		CamdemyFolderIE = λ.Cal(λ.TypeType, λ.NewStr("CamdemyFolderIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CamdemyFolderIE__VALID_URL λ.Object
			)
			CamdemyFolderIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?camdemy\\.com/folder/(?P<id>\\d+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): CamdemyFolderIE__VALID_URL,
			})
		}())
	})
}
