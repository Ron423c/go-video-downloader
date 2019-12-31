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
 * porn91/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/porn91.py
 */

package porn91

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError  λ.Object
	InfoExtractor   λ.Object
	Porn91IE        λ.Object
	ϒint_or_none    λ.Object
	ϒparse_duration λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒint_or_none = Ωutils.ϒint_or_none
		ExtractorError = Ωutils.ExtractorError
		Porn91IE = λ.Cal(λ.TypeType, λ.NewStr("Porn91IE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				Porn91IE_IE_NAME       λ.Object
				Porn91IE__VALID_URL    λ.Object
				Porn91IE__real_extract λ.Object
			)
			Porn91IE_IE_NAME = λ.NewStr("91porn")
			Porn91IE__VALID_URL = λ.NewStr("(?:https?://)(?:www\\.|)91porn\\.com/.+?\\?viewkey=(?P<id>[\\w\\d]+)")
			Porn91IE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcomment_count  λ.Object
						ϒduration       λ.Object
						ϒinfo_dict      λ.Object
						ϒself           = λargs[0]
						ϒtitle          λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒvideo_link_url λ.Object
						ϒvideopage      λ.Object
						ϒwebpage        λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					λ.Cal(λ.GetAttr(ϒself, "_set_cookie", nil), λ.NewStr("91porn.com"), λ.NewStr("language"), λ.NewStr("cn_CN"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Mod(λ.NewStr("http://91porn.com/view_video.php?viewkey=%s"), ϒvideo_id), ϒvideo_id)
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr("作为游客，你每天只可观看10个视频")))) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("91 Porn says: Daily limit 10 videos exceeded")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<div id=\"viewvideo-title\">([^<]+)</div>"), ϒwebpage, λ.NewStr("title"))
					ϒtitle = λ.Cal(λ.GetAttr(ϒtitle, "replace", nil), λ.NewStr("\n"), λ.NewStr(""))
					ϒvideo_link_url = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<textarea[^>]+id=[\"\\']fm-video_link[^>]+>([^<]+)</textarea>"), ϒwebpage, λ.NewStr("video link"))
					ϒvideopage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒvideo_link_url, ϒvideo_id)
					ϒinfo_dict = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_parse_html5_media_entries", nil), ϒurl, ϒvideopage, ϒvideo_id), λ.NewInt(0))
					ϒduration = λ.Cal(ϒparse_duration, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("时长:\\s*</span>\\s*(\\d+:\\d+)"),
						ϒwebpage,
						λ.NewStr("duration"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					ϒcomment_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("留言:\\s*</span>\\s*(\\d+)"),
						ϒwebpage,
						λ.NewStr("comment count"),
					), λ.KWArgs{
						{Name: "fatal", Value: λ.False},
					}))
					λ.Cal(λ.GetAttr(ϒinfo_dict, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):            ϒvideo_id,
						λ.NewStr("title"):         ϒtitle,
						λ.NewStr("duration"):      ϒduration,
						λ.NewStr("comment_count"): ϒcomment_count,
						λ.NewStr("age_limit"):     λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage),
					}))
					return ϒinfo_dict
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       Porn91IE_IE_NAME,
				λ.NewStr("_VALID_URL"):    Porn91IE__VALID_URL,
				λ.NewStr("_real_extract"): Porn91IE__real_extract,
			})
		}())
	})
}
