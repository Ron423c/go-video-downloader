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
 * cinchcast/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/cinchcast.py
 */

package cinchcast

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CinchcastIE      λ.Object
	InfoExtractor    λ.Object
	ϒunified_strdate λ.Object
	ϒxpath_text      λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒxpath_text = Ωutils.ϒxpath_text
		CinchcastIE = λ.Cal(λ.TypeType, λ.NewStr("CinchcastIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				CinchcastIE__TESTS        λ.Object
				CinchcastIE__VALID_URL    λ.Object
				CinchcastIE__real_extract λ.Object
			)
			CinchcastIE__VALID_URL = λ.NewStr("https?://player\\.cinchcast\\.com/.*?(?:assetId|show_id)=(?P<id>[0-9]+)")
			CinchcastIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.cinchcast.com/?show_id=5258197&platformId=1&assetType=single"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("5258197"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Train Your Brain to Up Your Game with Coach Mandy"),
						λ.NewStr("upload_date"): λ.NewStr("20130816"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://player.cinchcast.com/?platformId=1&#038;assetType=single&#038;assetId=7141703"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			CinchcastIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒbackup_url  λ.Object
						ϒdate_str    λ.Object
						ϒdoc         λ.Object
						ϒformats     λ.Object
						ϒitem        λ.Object
						ϒself        = λargs[0]
						ϒtitle       λ.Object
						ϒupload_date λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒdoc = λ.Cal(λ.GetAttr(ϒself, "_download_xml", nil), λ.Mod(λ.NewStr("http://www.blogtalkradio.com/playerasset/mrss?assetType=single&assetId=%s"), ϒvideo_id), ϒvideo_id)
					ϒitem = λ.Cal(λ.GetAttr(ϒdoc, "find", nil), λ.NewStr(".//item"))
					ϒtitle = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒitem,
						λ.NewStr("./title"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.True},
					})
					ϒdate_str = λ.Cal(ϒxpath_text, ϒitem, λ.NewStr("./{http://developer.longtailvideo.com/trac/}date"))
					ϒupload_date = λ.Call(ϒunified_strdate, λ.NewArgs(ϒdate_str), λ.KWArgs{
						{Name: "day_first", Value: λ.False},
					})
					ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("format_id"): λ.NewStr("main"),
						λ.NewStr("url"):       λ.GetItem(λ.GetAttr(λ.Cal(λ.GetAttr(ϒitem, "find", nil), λ.NewStr("./{http://search.yahoo.com/mrss/}content")), "attrib", nil), λ.NewStr("url")),
					}))
					ϒbackup_url = λ.Cal(ϒxpath_text, ϒitem, λ.NewStr("./{http://developer.longtailvideo.com/trac/}backupContent"))
					if λ.IsTrue(ϒbackup_url) {
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("preference"): λ.NewInt(2),
							λ.NewStr("format_id"):  λ.NewStr("backup"),
							λ.NewStr("url"):        ϒbackup_url,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("upload_date"): ϒupload_date,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        CinchcastIE__TESTS,
				λ.NewStr("_VALID_URL"):    CinchcastIE__VALID_URL,
				λ.NewStr("_real_extract"): CinchcastIE__real_extract,
			})
		}())
	})
}