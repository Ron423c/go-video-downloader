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
 * photobucket/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/photobucket.py
 */

package photobucket

import (
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor                λ.Object
	PhotobucketIE                λ.Object
	ϒcompat_urllib_parse_unquote λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_unquote = Ωcompat.ϒcompat_urllib_parse_unquote
		PhotobucketIE = λ.Cal(λ.TypeType, λ.NewStr("PhotobucketIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PhotobucketIE__TEST         λ.Object
				PhotobucketIE__VALID_URL    λ.Object
				PhotobucketIE__real_extract λ.Object
			)
			PhotobucketIE__VALID_URL = λ.NewStr("https?://(?:[a-z0-9]+\\.)?photobucket\\.com/.*(([\\?\\&]current=)|_)(?P<id>.*)\\.(?P<ext>(flv)|(mp4))")
			PhotobucketIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://media.photobucket.com/user/rachaneronas/media/TiredofLinkBuildingTryBacklinkMyDomaincom_zpsc0c3b9fa.mp4.html?filters[term]=search&filters[primary]=videos&filters[secondary]=images&sort=1&o=0"),
				λ.NewStr("md5"): λ.NewStr("7dabfb92b0a31f6c16cebc0f8e60ff99"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("zpsc0c3b9fa"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("timestamp"):   λ.NewInt(1367669341),
					λ.NewStr("upload_date"): λ.NewStr("20130504"),
					λ.NewStr("uploader"):    λ.NewStr("rachaneronas"),
					λ.NewStr("title"):       λ.NewStr("Tired of Link Building? Try BacklinkMyDomain.com!"),
				}),
			})
			PhotobucketIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo            λ.Object
						ϒinfo_json       λ.Object
						ϒmobj            λ.Object
						ϒself            = λargs[0]
						ϒurl             = λargs[1]
						ϒvideo_extension λ.Object
						ϒvideo_id        λ.Object
						ϒwebpage         λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒvideo_extension = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("ext"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					λ.Cal(λ.GetAttr(ϒself, "report_extraction", nil), ϒvideo_id)
					ϒinfo_json = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("Pb\\.Data\\.Shared\\.put\\(Pb\\.Data\\.Shared\\.MEDIA, (.*?)\\);"), ϒwebpage, λ.NewStr("info json"))
					ϒinfo = λ.Cal(Ωjson.ϒloads, ϒinfo_json)
					ϒurl = λ.Cal(ϒcompat_urllib_parse_unquote, λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("file=(.+\\.mp4)"), λ.GetItem(λ.GetItem(ϒinfo, λ.NewStr("linkcodes")), λ.NewStr("html")), λ.NewStr("url")))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("url"):       ϒurl,
						λ.NewStr("uploader"):  λ.GetItem(ϒinfo, λ.NewStr("username")),
						λ.NewStr("timestamp"): λ.GetItem(ϒinfo, λ.NewStr("creationDate")),
						λ.NewStr("title"):     λ.GetItem(ϒinfo, λ.NewStr("title")),
						λ.NewStr("ext"):       ϒvideo_extension,
						λ.NewStr("thumbnail"): λ.GetItem(ϒinfo, λ.NewStr("thumbUrl")),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TEST"):         PhotobucketIE__TEST,
				λ.NewStr("_VALID_URL"):    PhotobucketIE__VALID_URL,
				λ.NewStr("_real_extract"): PhotobucketIE__real_extract,
			})
		}())
	})
}
