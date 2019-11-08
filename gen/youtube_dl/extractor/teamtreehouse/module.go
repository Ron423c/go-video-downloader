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
 * teamtreehouse/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/teamtreehouse.py
 */

package teamtreehouse

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError        λ.Object
	InfoExtractor         λ.Object
	TeamTreeHouseIE       λ.Object
	ϒclean_html           λ.Object
	ϒdetermine_ext        λ.Object
	ϒfloat_or_none        λ.Object
	ϒget_element_by_class λ.Object
	ϒget_element_by_id    λ.Object
	ϒparse_duration       λ.Object
	ϒremove_end           λ.Object
	ϒurlencode_postdata   λ.Object
	ϒurljoin              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒget_element_by_class = Ωutils.ϒget_element_by_class
		ϒget_element_by_id = Ωutils.ϒget_element_by_id
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒremove_end = Ωutils.ϒremove_end
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		ϒurljoin = Ωutils.ϒurljoin
		TeamTreeHouseIE = λ.Cal(λ.TypeType, λ.NewStr("TeamTreeHouseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TeamTreeHouseIE__NETRC_MACHINE   λ.Object
				TeamTreeHouseIE__VALID_URL       λ.Object
				TeamTreeHouseIE__real_extract    λ.Object
				TeamTreeHouseIE__real_initialize λ.Object
			)
			TeamTreeHouseIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?teamtreehouse\\.com/library/(?P<id>[^/]+)")
			TeamTreeHouseIE__NETRC_MACHINE = λ.NewStr("teamtreehouse")
			TeamTreeHouseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata          λ.Object
						ϒemail         λ.Object
						ϒerror_message λ.Object
						ϒpassword      λ.Object
						ϒself          = λargs[0]
						ϒsignin_page   λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_get_login_info", nil))
					ϒemail = λ.GetItem(τmp0, λ.NewInt(0))
					ϒpassword = λ.GetItem(τmp0, λ.NewInt(1))
					if λ.IsTrue(λ.NewBool(ϒemail == λ.None)) {
						return λ.None
					}
					ϒsignin_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewStr("https://teamtreehouse.com/signin"), λ.None, λ.NewStr("Downloading signin page"))
					ϒdata = λ.Cal(λ.GetAttr(ϒself, "_form_hidden_inputs", nil), λ.NewStr("new_user_session"), ϒsignin_page)
					λ.Cal(λ.GetAttr(ϒdata, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("user_session[email]"):    ϒemail,
						λ.NewStr("user_session[password]"): ϒpassword,
					}))
					ϒerror_message = λ.Cal(ϒget_element_by_class, λ.NewStr("error-message"), λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(
						λ.NewStr("https://teamtreehouse.com/person_session"),
						λ.None,
						λ.NewStr("Logging in"),
					), λ.KWArgs{
						{Name: "data", Value: λ.Cal(ϒurlencode_postdata, ϒdata)},
					}))
					if λ.IsTrue(ϒerror_message) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Cal(ϒclean_html, ϒerror_message)), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			TeamTreeHouseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒchapter         λ.Object
						ϒchapter_number  λ.Object
						ϒdescription     λ.Object
						ϒdisplay_id      λ.Object
						ϒduration        λ.Object
						ϒentries         λ.Object
						ϒextract_urls    λ.Object
						ϒinfo            λ.Object
						ϒis_preview      λ.Object
						ϒself            = λargs[0]
						ϒstages_page     λ.Object
						ϒstages_path     λ.Object
						ϒsteps_list      λ.Object
						ϒsubtitle        λ.Object
						ϒsubtitles       λ.Object
						ϒtitle           λ.Object
						ϒurl             = λargs[1]
						ϒwebpage         λ.Object
						ϒworkshop_videos λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
						τmp3             λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewList(
						λ.NewStr("og:title"),
						λ.NewStr("twitter:title"),
					), ϒwebpage)
					ϒdescription = λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewList(
						λ.NewStr("description"),
						λ.NewStr("og:description"),
						λ.NewStr("twitter:description"),
					), ϒwebpage)
					ϒentries = λ.Cal(λ.GetAttr(ϒself, "_parse_html5_media_entries", nil), ϒurl, ϒwebpage, ϒdisplay_id)
					if λ.IsTrue(ϒentries) {
						ϒinfo = λ.GetItem(ϒentries, λ.NewInt(0))
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("subtitles"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "values", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒsubtitles = τmp1
							τmp2 = λ.Cal(λ.BuiltinIter, ϒsubtitles)
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								ϒsubtitle = τmp3
								λ.SetItem(ϒsubtitle, λ.NewStr("ext"), λ.Cal(ϒdetermine_ext, λ.GetItem(ϒsubtitle, λ.NewStr("url")), λ.NewStr("srt")))
							}
						}
						ϒis_preview = λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr("data-preview=\"true\"")))
						if λ.IsTrue(ϒis_preview) {
							λ.Cal(λ.GetAttr(ϒself, "report_warning", nil), λ.NewStr("This is just a preview. You need to be signed in with a Basic account to download the entire video."), ϒdisplay_id)
							ϒduration = λ.NewInt(30)
						} else {
							ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("data-duration=\"(\\d+)\""), ϒwebpage, λ.NewStr("duration")), λ.NewInt(1000))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒduration))) {
								ϒduration = λ.Cal(ϒparse_duration, λ.Cal(ϒget_element_by_id, λ.NewStr("video-duration"), ϒwebpage))
							}
						}
						λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):          ϒdisplay_id,
							λ.NewStr("title"):       ϒtitle,
							λ.NewStr("description"): ϒdescription,
							λ.NewStr("duration"):    ϒduration,
						}))
						return ϒinfo
					} else {
						ϒextract_urls = λ.NewFunction("extract_urls",
							[]λ.Param{
								{Name: "html"},
								{Name: "extract_info", Def: λ.None},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒentry        λ.Object
									ϒextract_info = λargs[1]
									ϒhtml         = λargs[0]
									ϒpage_url     λ.Object
									ϒpath         λ.Object
									τmp0          λ.Object
									τmp1          λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfindall, λ.NewStr("<a[^>]+href=\"([^\"]+)\""), ϒhtml))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒpath = τmp1
									ϒpage_url = λ.Cal(ϒurljoin, ϒurl, ϒpath)
									ϒentry = λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("_type"):  λ.NewStr("url_transparent"),
										λ.NewStr("id"):     λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒpage_url),
										λ.NewStr("url"):    ϒpage_url,
										λ.NewStr("id_key"): λ.Cal(λ.GetAttr(ϒself, "ie_key", nil)),
									})
									if λ.IsTrue(ϒextract_info) {
										λ.Cal(λ.GetAttr(ϒentry, "update", nil), ϒextract_info)
									}
									λ.Cal(λ.GetAttr(ϒentries, "append", nil), ϒentry)
								}
								return λ.None
							})
						ϒworkshop_videos = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("(?s)<ul[^>]+id=\"workshop-videos\"[^>]*>(.+?)</ul>"),
							ϒwebpage,
							λ.NewStr("workshop videos"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
						if λ.IsTrue(ϒworkshop_videos) {
							λ.Cal(ϒextract_urls, ϒworkshop_videos)
						} else {
							ϒstages_path = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("(?s)<div[^>]+id=\"syllabus-stages\"[^>]+data-url=\"([^\"]+)\""), ϒwebpage, λ.NewStr("stages path"))
							if λ.IsTrue(ϒstages_path) {
								ϒstages_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Cal(ϒurljoin, ϒurl, ϒstages_path), ϒdisplay_id, λ.NewStr("Downloading stages page"))
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, λ.Cal(Ωre.ϒfindall, λ.NewStr("(?s)<h2[^>]*>\\s*(.+?)\\s*</h2>.+?<ul[^>]*>(.+?)</ul>"), ϒstages_page), λ.NewInt(1)))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									τmp2 = τmp1
									ϒchapter_number = λ.GetItem(τmp2, λ.NewInt(0))
									τmp3 = λ.GetItem(τmp2, λ.NewInt(1))
									ϒchapter = λ.GetItem(τmp3, λ.NewInt(0))
									ϒsteps_list = λ.GetItem(τmp3, λ.NewInt(1))
									λ.Cal(ϒextract_urls, ϒsteps_list, λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("chapter"):        ϒchapter,
										λ.NewStr("chapter_number"): ϒchapter_number,
									}))
								}
								ϒtitle = λ.Cal(ϒremove_end, ϒtitle, λ.NewStr(" Course"))
							}
						}
						return λ.Cal(λ.GetAttr(ϒself, "playlist_result", nil), ϒentries, ϒdisplay_id, ϒtitle, ϒdescription)
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_NETRC_MACHINE"):   TeamTreeHouseIE__NETRC_MACHINE,
				λ.NewStr("_VALID_URL"):       TeamTreeHouseIE__VALID_URL,
				λ.NewStr("_real_extract"):    TeamTreeHouseIE__real_extract,
				λ.NewStr("_real_initialize"): TeamTreeHouseIE__real_initialize,
			})
		}())
	})
}
