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
 * dailymotion/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/dailymotion.py
 */

package dailymotion

import (
	Ωrandom "github.com/tenta-browser/go-video-downloader/gen/random"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωstring "github.com/tenta-browser/go-video-downloader/gen/string"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	DailymotionBaseInfoExtractor λ.Object
	DailymotionIE                λ.Object
	DailymotionPlaylistIE        λ.Object
	DailymotionUserIE            λ.Object
	ExtractorError               λ.Object
	InfoExtractor                λ.Object
	ϒdetermine_ext               λ.Object
	ϒerror_to_compat_str         λ.Object
	ϒint_or_none                 λ.Object
	ϒmimetype2ext                λ.Object
	ϒparse_iso8601               λ.Object
	ϒsanitized_Request           λ.Object
	ϒstr_to_int                  λ.Object
	ϒtry_get                     λ.Object
	ϒunescapeHTML                λ.Object
	ϒupdate_url_query            λ.Object
	ϒurl_or_none                 λ.Object
	ϒurlencode_postdata          λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒerror_to_compat_str = Ωutils.ϒerror_to_compat_str
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒstr_to_int = Ωutils.ϒstr_to_int
		ϒtry_get = Ωutils.ϒtry_get
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		DailymotionBaseInfoExtractor = λ.Cal(λ.TypeType, λ.NewStr("DailymotionBaseInfoExtractor"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				DailymotionBaseInfoExtractor__build_request          λ.Object
				DailymotionBaseInfoExtractor__download_webpage_no_ff λ.Object
			)
			DailymotionBaseInfoExtractor__build_request = λ.NewFunction("_build_request",
				[]λ.Param{
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒrequest λ.Object
						ϒurl     = λargs[0]
					)
					λ.NewStr("Build a request with the family filter disabled")
					ϒrequest = λ.Cal(ϒsanitized_Request, ϒurl)
					λ.Cal(λ.GetAttr(ϒrequest, "add_header", nil), λ.NewStr("Cookie"), λ.NewStr("family_filter=off; ff=off"))
					return ϒrequest
				})
			DailymotionBaseInfoExtractor__build_request = λ.Cal(λ.StaticMethodType, DailymotionBaseInfoExtractor__build_request)
			DailymotionBaseInfoExtractor__download_webpage_no_ff = λ.NewFunction("_download_webpage_no_ff",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, true, true,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒargs    = λargs[2]
						ϒkwargs  = λargs[3]
						ϒrequest λ.Object
						ϒself    = λargs[0]
						ϒurl     = λargs[1]
					)
					ϒrequest = λ.Cal(λ.GetAttr(ϒself, "_build_request", nil), ϒurl)
					return λ.Call(λ.GetAttr(ϒself, "_download_webpage", nil), λ.NewArgs(λ.Unpack(
						ϒrequest,
						λ.AsStarred(ϒargs),
					)...), λ.KWArgs{
						{Name: "", Value: ϒkwargs},
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_build_request"):          DailymotionBaseInfoExtractor__build_request,
				λ.NewStr("_download_webpage_no_ff"): DailymotionBaseInfoExtractor__download_webpage_no_ff,
			})
		}())
		DailymotionIE = λ.Cal(λ.TypeType, λ.NewStr("DailymotionIE"), λ.NewTuple(DailymotionBaseInfoExtractor), func() λ.Dict {
			var (
				DailymotionIE_IE_NAME       λ.Object
				DailymotionIE__TESTS        λ.Object
				DailymotionIE__VALID_URL    λ.Object
				DailymotionIE__check_error  λ.Object
				DailymotionIE__extract_urls λ.Object
				DailymotionIE__real_extract λ.Object
			)
			DailymotionIE__VALID_URL = λ.NewStr("(?ix)\n                    https?://\n                        (?:\n                            (?:(?:www|touch)\\.)?dailymotion\\.[a-z]{2,3}/(?:(?:(?:embed|swf|\\#)/)?video|swf)|\n                            (?:www\\.)?lequipe\\.fr/video\n                        )\n                        /(?P<id>[^/?_]+)\n                    ")
			DailymotionIE_IE_NAME = λ.NewStr("dailymotion")
			DailymotionIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.dailymotion.com/video/x5kesuj_office-christmas-party-review-jason-bateman-olivia-munn-t-j-miller_news"),
					λ.NewStr("md5"): λ.NewStr("074b95bdee76b9e3654137aee9c79dfe"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("x5kesuj"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Office Christmas Party Review –  Jason Bateman, Olivia Munn, T.J. Miller"),
						λ.NewStr("description"): λ.NewStr("Office Christmas Party Review -  Jason Bateman, Olivia Munn, T.J. Miller"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?:.*\\.(?:jpg|png)$"),
						λ.NewStr("duration"):    λ.NewInt(187),
						λ.NewStr("timestamp"):   λ.NewInt(1493651285),
						λ.NewStr("upload_date"): λ.NewStr("20170501"),
						λ.NewStr("uploader"):    λ.NewStr("Deadline"),
						λ.NewStr("uploader_id"): λ.NewStr("x1xm8ri"),
						λ.NewStr("age_limit"):   λ.NewInt(0),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.dailymotion.com/video/x2iuewm_steam-machine-models-pricing-listed-on-steam-store-ign-news_videogames"),
					λ.NewStr("md5"): λ.NewStr("2137c41a8e78554bb09225b8eb322406"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("x2iuewm"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Steam Machine Models, Pricing Listed on Steam Store - IGN News"),
						λ.NewStr("description"): λ.NewStr("Several come bundled with the Steam Controller."),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?:.*\\.(?:jpg|png)$"),
						λ.NewStr("duration"):    λ.NewInt(74),
						λ.NewStr("timestamp"):   λ.NewInt(1425657362),
						λ.NewStr("upload_date"): λ.NewStr("20150306"),
						λ.NewStr("uploader"):    λ.NewStr("IGN"),
						λ.NewStr("uploader_id"): λ.NewStr("xijv66"),
						λ.NewStr("age_limit"):   λ.NewInt(0),
						λ.NewStr("view_count"):  λ.IntType,
					}),
					λ.NewStr("skip"): λ.NewStr("video gone"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.dailymotion.com/video/x149uew_katy-perry-roar-official_musi"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"):       λ.NewStr("Roar (Official)"),
						λ.NewStr("id"):          λ.NewStr("USUV71301934"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("uploader"):    λ.NewStr("Katy Perry"),
						λ.NewStr("upload_date"): λ.NewStr("20130905"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("skip"): λ.NewStr("VEVO is only available in some countries"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.dailymotion.com/video/xyh2zz_leanna-decker-cyber-girl-of-the-year-desires-nude-playboy-plus_redband"),
					λ.NewStr("md5"): λ.NewStr("0d667a7b9cebecc3c89ee93099c4159d"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("xyh2zz"),
						λ.NewStr("ext"):       λ.NewStr("mp4"),
						λ.NewStr("title"):     λ.NewStr("Leanna Decker - Cyber Girl Of The Year Desires Nude [Playboy Plus]"),
						λ.NewStr("uploader"):  λ.NewStr("HotWaves1012"),
						λ.NewStr("age_limit"): λ.NewInt(18),
					}),
					λ.NewStr("skip"): λ.NewStr("video gone"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.dailymotion.com/video/xhza0o"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.dailymotion.com/video/x20su5f_the-power-of-nightmares-1-the-rise-of-the-politics-of-fear-bbc-2004_news"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.dailymotion.com/swf/video/x3n92nf"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.dailymotion.com/swf/x3ss1m_funny-magic-trick-barry-and-stuart_fun"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.lequipe.fr/video/x791mem"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.lequipe.fr/video/k7MtHciueyTcrFtFKA2"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			DailymotionIE__extract_urls = λ.NewFunction("_extract_urls",
				[]λ.Param{
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒmobj    λ.Object
						ϒurls    λ.Object
						ϒwebpage = λargs[0]
						τmp0     λ.Object
						τmp1     λ.Object
					)
					ϒurls = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("<(?:(?:embed|iframe)[^>]+?src=|input[^>]+id=[\\'\"]dmcloudUrlEmissionSelect[\\'\"][^>]+value=)([\"\\'])(?P<url>(?:https?:)?//(?:www\\.)?dailymotion\\.com/(?:embed|swf)/video/.+?)\\1"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmobj = τmp1
						λ.Cal(λ.GetAttr(ϒurls, "append", nil), λ.Cal(ϒunescapeHTML, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("url"))))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(Ωre.ϒfinditer, λ.NewStr("(?s)DM\\.player\\([^,]+,\\s*{.*?video[\\'\"]?\\s*:\\s*[\"\\']?(?P<id>[0-9a-zA-Z]+).+?}\\s*\\);"), ϒwebpage))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒmobj = τmp1
						λ.Cal(λ.GetAttr(ϒurls, "append", nil), λ.Add(λ.NewStr("https://www.dailymotion.com/embed/video/"), λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))))
					}
					return ϒurls
				})
			DailymotionIE__extract_urls = λ.Cal(λ.StaticMethodType, DailymotionIE__extract_urls)
			DailymotionIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒage_limit       λ.Object
						ϒcomment_count   λ.Object
						ϒdescription     λ.Object
						ϒduration        λ.Object
						ϒembed_page      λ.Object
						ϒext             λ.Object
						ϒf               λ.Object
						ϒformat_id       λ.Object
						ϒformats         λ.Object
						ϒheight          λ.Object
						ϒi               λ.Object
						ϒinfo            λ.Object
						ϒkey             λ.Object
						ϒm               λ.Object
						ϒm3u8_formats    λ.Object
						ϒm_size          λ.Object
						ϒmedia           λ.Object
						ϒmedia_list      λ.Object
						ϒmedia_url       λ.Object
						ϒmetadata        λ.Object
						ϒmetadata_url    λ.Object
						ϒn               λ.Object
						ϒpassword        λ.Object
						ϒplayer          λ.Object
						ϒplayer_v5       λ.Object
						ϒquality         λ.Object
						ϒr               λ.Object
						ϒself            = λargs[0]
						ϒsubtitle        λ.Object
						ϒsubtitle_lang   λ.Object
						ϒsubtitles       λ.Object
						ϒsubtitles_data  λ.Object
						ϒt               λ.Object
						ϒthumbnail       λ.Object
						ϒtimestamp       λ.Object
						ϒtitle           λ.Object
						ϒtype_           λ.Object
						ϒuploader        λ.Object
						ϒuploader_id     λ.Object
						ϒurl             = λargs[1]
						ϒus64e           λ.Object
						ϒvevo_id         λ.Object
						ϒvideo_id        λ.Object
						ϒvideo_subtitles λ.Object
						ϒvideo_url       λ.Object
						ϒview_count      λ.Object
						ϒview_count_str  λ.Object
						ϒwebpage         λ.Object
						ϒwidth           λ.Object
						τmp0             λ.Object
						τmp1             λ.Object
						τmp2             λ.Object
						τmp3             λ.Object
						τmp4             λ.Object
						τmp5             λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_no_ff", nil), λ.Mod(λ.NewStr("https://www.dailymotion.com/video/%s"), ϒvideo_id), ϒvideo_id)
					ϒage_limit = λ.Cal(λ.GetAttr(ϒself, "_rta_search", nil), ϒwebpage)
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("description"), ϒwebpage, λ.NewStr("description"))
						}
					}()
					ϒview_count_str = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewTuple(
							λ.NewStr("<meta[^>]+itemprop=\"interactionCount\"[^>]+content=\"UserPlays:([\\s\\d,.]+)\""),
							λ.NewStr("video_views_count[^>]+>\\s+([\\s\\d\\,.]+)"),
						),
						ϒwebpage,
						λ.NewStr("view count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒview_count_str) {
						ϒview_count_str = λ.Cal(Ωre.ϒsub, λ.NewStr("\\s"), λ.NewStr(""), ϒview_count_str)
					}
					ϒview_count = λ.Cal(ϒstr_to_int, ϒview_count_str)
					ϒcomment_count = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<meta[^>]+itemprop=\"interactionCount\"[^>]+content=\"UserComments:(\\d+)\""),
						ϒwebpage,
						λ.NewStr("comment count"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒplayer_v5 = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewList(
							λ.NewStr("buildPlayer\\(({.+?})\\);\\n"),
							λ.NewStr("playerV5\\s*=\\s*dmp\\.create\\([^,]+?,\\s*({.+?})\\);"),
							λ.NewStr("buildPlayer\\(({.+?})\\);"),
							λ.NewStr("var\\s+config\\s*=\\s*({.+?});"),
							λ.NewStr("__PLAYER_CONFIG__\\s*=\\s*({.+?});"),
						),
						ϒwebpage,
						λ.NewStr("player v5"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒplayer_v5) {
						ϒplayer = func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
								ϒplayer_v5,
								ϒvideo_id,
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewDictWithTable(map[λ.Object]λ.Object{})
							}
						}()
						ϒmetadata = λ.Cal(ϒtry_get, ϒplayer, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(ϒx, λ.NewStr("metadata"))
							}), λ.DictType)
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒmetadata))) {
							ϒmetadata_url = λ.Cal(ϒurl_or_none, λ.Cal(ϒtry_get, ϒplayer, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("context")), λ.NewStr("metadata_template_url1"))
								})))
							if λ.IsTrue(ϒmetadata_url) {
								ϒmetadata_url = λ.Cal(λ.GetAttr(ϒmetadata_url, "replace", nil), λ.NewStr(":videoId"), ϒvideo_id)
							} else {
								ϒmetadata_url = λ.Cal(ϒupdate_url_query, λ.Mod(λ.NewStr("https://www.dailymotion.com/player/metadata/video/%s"), ϒvideo_id), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("embedder"):    ϒurl,
									λ.NewStr("integration"): λ.NewStr("inline"),
									λ.NewStr("GK_PV5_NEON"): λ.NewStr("1"),
								}))
							}
							ϒmetadata = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒmetadata_url, ϒvideo_id, λ.NewStr("Downloading metadata JSON"))
						}
						if λ.IsTrue(λ.Eq(λ.Cal(ϒtry_get, ϒmetadata, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("error")), λ.NewStr("type"))
							})), λ.NewStr("password_protected"))) {
							ϒpassword = λ.Cal(λ.GetAttr(λ.GetAttr(λ.GetAttr(ϒself, "_downloader", nil), "params", nil), "get", nil), λ.NewStr("videopassword"))
							if λ.IsTrue(ϒpassword) {
								ϒr = λ.Cal(λ.IntType, λ.GetItem(λ.GetItem(ϒmetadata, λ.NewStr("id")), λ.NewSlice(λ.NewInt(1), λ.None, λ.None)), λ.NewInt(36))
								ϒus64e = λ.NewFunction("<lambda>",
									[]λ.Param{
										{Name: "x"},
									},
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										var (
											ϒx = λargs[0]
										)
										return λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.Cal(λ.None, ϒx), "decode", nil)), "strip", nil), λ.NewStr("="))
									})
								ϒt = λ.Cal(λ.GetAttr(λ.NewStr(""), "join", nil), λ.Cal(λ.NewFunction("<generator>",
									nil,
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
											var (
												ϒi   λ.Object
												τmp0 λ.Object
												τmp1 λ.Object
											)
											_ = ϒi
											τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.RangeType, λ.NewInt(10)))
											for {
												if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
													break
												}
												ϒi = τmp1
												λgy.Yield(λ.Cal(Ωrandom.ϒchoice, Ωstring.ϒascii_letters))
											}
											return λ.None
										})
									})))
								ϒn = λ.Cal(ϒus64e, λ.Cal(λ.None, λ.NewStr("I"), ϒr))
								ϒi = λ.Cal(ϒus64e, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(λ.None, "md5", nil), λ.Cal(λ.GetAttr(λ.Mod(λ.NewStr("%s%d%s"), λ.NewTuple(
									ϒpassword,
									ϒr,
									ϒt,
								)), "encode", nil))), "digest", nil)))
								ϒmetadata = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Add(λ.Add(λ.Add(λ.NewStr("http://www.dailymotion.com/player/metadata/video/p"), ϒi), ϒt), ϒn), ϒvideo_id)
							}
						}
						λ.Cal(λ.GetAttr(ϒself, "_check_error", nil), ϒmetadata)
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.GetItem(ϒmetadata, λ.NewStr("qualities")), "items", nil)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							τmp2 = τmp1
							ϒquality = λ.GetItem(τmp2, λ.NewInt(0))
							ϒmedia_list = λ.GetItem(τmp2, λ.NewInt(1))
							τmp2 = λ.Cal(λ.BuiltinIter, ϒmedia_list)
							for {
								if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
									break
								}
								ϒmedia = τmp3
								ϒmedia_url = λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("url"))
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒmedia_url))) {
									continue
								}
								ϒtype_ = λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("type"))
								if λ.IsTrue(λ.Eq(ϒtype_, λ.NewStr("application/vnd.lumberjack.manifest"))) {
									continue
								}
								ϒext = func() λ.Object {
									if λv := λ.Cal(ϒmimetype2ext, ϒtype_); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.Cal(ϒdetermine_ext, ϒmedia_url)
									}
								}()
								if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
									ϒm3u8_formats = λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒmedia_url,
										ϒvideo_id,
										λ.NewStr("mp4"),
									), λ.KWArgs{
										{Name: "preference", Value: λ.Neg(λ.NewInt(1))},
										{Name: "m3u8_id", Value: λ.NewStr("hls")},
										{Name: "fatal", Value: λ.False},
									})
									τmp4 = λ.Cal(λ.BuiltinIter, ϒm3u8_formats)
									for {
										if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
											break
										}
										ϒf = τmp5
										λ.SetItem(ϒf, λ.NewStr("url"), λ.GetItem(λ.Cal(λ.GetAttr(λ.GetItem(ϒf, λ.NewStr("url")), "split", nil), λ.NewStr("#")), λ.NewInt(0)))
										λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
									}
								} else {
									if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("f4m"))) {
										λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
											ϒmedia_url,
											ϒvideo_id,
										), λ.KWArgs{
											{Name: "preference", Value: λ.Neg(λ.NewInt(1))},
											{Name: "f4m_id", Value: λ.NewStr("hds")},
											{Name: "fatal", Value: λ.False},
										}))
									} else {
										ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("url"):       ϒmedia_url,
											λ.NewStr("format_id"): λ.Mod(λ.NewStr("http-%s"), ϒquality),
											λ.NewStr("ext"):       ϒext,
										})
										ϒm = λ.Cal(Ωre.ϒsearch, λ.NewStr("H264-(?P<width>\\d+)x(?P<height>\\d+)"), ϒmedia_url)
										if λ.IsTrue(ϒm) {
											λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("width"):  λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("width"))),
												λ.NewStr("height"): λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("height"))),
											}))
										}
										λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
									}
								}
							}
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
						ϒtitle = λ.GetItem(ϒmetadata, λ.NewStr("title"))
						ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("duration")))
						ϒtimestamp = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("created_time")))
						ϒthumbnail = λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("poster_url"))
						ϒuploader = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("owner"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("screenname"))
						ϒuploader_id = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("owner"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("id"))
						ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
						ϒsubtitles_data = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("subtitles"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("data"), λ.NewDictWithTable(map[λ.Object]λ.Object{}))
						if λ.IsTrue(func() λ.Object {
							if λv := ϒsubtitles_data; !λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.BuiltinIsInstance, ϒsubtitles_data, λ.DictType)
							}
						}()) {
							τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒsubtitles_data, "items", nil)))
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								τmp2 = τmp1
								ϒsubtitle_lang = λ.GetItem(τmp2, λ.NewInt(0))
								ϒsubtitle = λ.GetItem(τmp2, λ.NewInt(1))
								λ.SetItem(ϒsubtitles, ϒsubtitle_lang, λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
									nil,
									0, false, false,
									func(λargs []λ.Object) λ.Object {
										return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
											var (
												ϒsubtitle_url λ.Object
												τmp0          λ.Object
												τmp1          λ.Object
											)
											τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒsubtitle, "get", nil), λ.NewStr("urls"), λ.NewList()))
											for {
												if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
													break
												}
												ϒsubtitle_url = τmp1
												λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
													λ.NewStr("ext"): λ.Cal(ϒdetermine_ext, ϒsubtitle_url),
													λ.NewStr("url"): ϒsubtitle_url,
												}))
											}
											return λ.None
										})
									}))))
							}
						}
						return λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):            ϒvideo_id,
							λ.NewStr("title"):         ϒtitle,
							λ.NewStr("description"):   ϒdescription,
							λ.NewStr("thumbnail"):     ϒthumbnail,
							λ.NewStr("duration"):      ϒduration,
							λ.NewStr("timestamp"):     ϒtimestamp,
							λ.NewStr("uploader"):      ϒuploader,
							λ.NewStr("uploader_id"):   ϒuploader_id,
							λ.NewStr("age_limit"):     ϒage_limit,
							λ.NewStr("view_count"):    ϒview_count,
							λ.NewStr("comment_count"): ϒcomment_count,
							λ.NewStr("formats"):       ϒformats,
							λ.NewStr("subtitles"):     ϒsubtitles,
						})
					}
					ϒvevo_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<link rel=\"video_src\" href=\"[^\"]*?vevo\\.com[^\"]*?video=(?P<id>[\\w]*)"),
						ϒwebpage,
						λ.NewStr("vevo embed"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒvevo_id) {
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("vevo:%s"), ϒvevo_id), λ.NewStr("Vevo"))
					}
					ϒembed_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage_no_ff", nil), λ.Mod(λ.NewStr("https://www.dailymotion.com/embed/video/%s"), ϒvideo_id), ϒvideo_id, λ.NewStr("Downloading embed page"))
					ϒtimestamp = λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("video:release_date"), ϒwebpage, λ.NewStr("upload date")))
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("var info = ({.*?}),$"),
						ϒembed_page,
						λ.NewStr("video info"),
					), λ.KWArgs{
						{Name: "flags", Value: Ωre.MULTILINE},
					}), ϒvideo_id)
					λ.Cal(λ.GetAttr(ϒself, "_check_error", nil), ϒinfo)
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetAttr(ϒself, "_FORMATS", nil))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒkey = λ.GetItem(τmp2, λ.NewInt(0))
						ϒformat_id = λ.GetItem(τmp2, λ.NewInt(1))
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), ϒkey)
						if λ.IsTrue(λ.NewBool(ϒvideo_url != λ.None)) {
							ϒm_size = λ.Cal(Ωre.ϒsearch, λ.NewStr("H264-(\\d+)x(\\d+)"), ϒvideo_url)
							if λ.IsTrue(λ.NewBool(ϒm_size != λ.None)) {
								τmp2 = λ.Cal(λ.MapIteratorType, ϒint_or_none, λ.NewTuple(
									λ.Cal(λ.GetAttr(ϒm_size, "group", nil), λ.NewInt(1)),
									λ.Cal(λ.GetAttr(ϒm_size, "group", nil), λ.NewInt(2)),
								))
								ϒwidth = λ.GetItem(τmp2, λ.NewInt(0))
								ϒheight = λ.GetItem(τmp2, λ.NewInt(1))
							} else {
								τmp2 = λ.NewTuple(
									λ.None,
									λ.None,
								)
								ϒwidth = λ.GetItem(τmp2, λ.NewInt(0))
								ϒheight = λ.GetItem(τmp2, λ.NewInt(1))
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("url"):       ϒvideo_url,
								λ.NewStr("ext"):       λ.NewStr("mp4"),
								λ.NewStr("format_id"): ϒformat_id,
								λ.NewStr("width"):     ϒwidth,
								λ.NewStr("height"):    ϒheight,
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒvideo_subtitles = λ.Cal(λ.GetAttr(ϒself, "extract_subtitles", nil), ϒvideo_id, ϒwebpage)
					ϒtitle = λ.Call(λ.GetAttr(ϒself, "_og_search_title", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒtitle == λ.None)) {
						ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("(?s)<span\\s+id=\"video_title\"[^>]*>(.*?)</span>"), ϒwebpage, λ.NewStr("title"))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("formats"):     ϒformats,
						λ.NewStr("uploader"):    λ.GetItem(ϒinfo, λ.NewStr("owner.screenname")),
						λ.NewStr("timestamp"):   ϒtimestamp,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("subtitles"):   ϒvideo_subtitles,
						λ.NewStr("thumbnail"):   λ.GetItem(ϒinfo, λ.NewStr("thumbnail_url")),
						λ.NewStr("age_limit"):   ϒage_limit,
						λ.NewStr("view_count"):  ϒview_count,
						λ.NewStr("duration"):    λ.GetItem(ϒinfo, λ.NewStr("duration")),
					})
				})
			DailymotionIE__check_error = λ.NewFunction("_check_error",
				[]λ.Param{
					{Name: "self"},
					{Name: "info"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒerror λ.Object
						ϒinfo  = λargs[1]
						ϒself  = λargs[0]
						ϒtitle λ.Object
					)
					ϒerror = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("error"))
					if λ.IsTrue(ϒerror) {
						ϒtitle = func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒerror, "get", nil), λ.NewStr("title")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetItem(ϒerror, λ.NewStr("message"))
							}
						}()
						if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒerror, "get", nil), λ.NewStr("code")), λ.NewStr("DM007"))) {
							λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), nil, λ.KWArgs{
								{Name: "msg", Value: ϒtitle},
							})
						}
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.NewStr("%s said: %s"), λ.NewTuple(
							λ.GetAttr(ϒself, "IE_NAME", nil),
							ϒtitle,
						))), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       DailymotionIE_IE_NAME,
				λ.NewStr("_TESTS"):        DailymotionIE__TESTS,
				λ.NewStr("_VALID_URL"):    DailymotionIE__VALID_URL,
				λ.NewStr("_check_error"):  DailymotionIE__check_error,
				λ.NewStr("_extract_urls"): DailymotionIE__extract_urls,
				λ.NewStr("_real_extract"): DailymotionIE__real_extract,
			})
		}())
		DailymotionPlaylistIE = λ.Cal(λ.TypeType, λ.NewStr("DailymotionPlaylistIE"), λ.NewTuple(DailymotionBaseInfoExtractor), func() λ.Dict {
			var (
				DailymotionPlaylistIE__VALID_URL λ.Object
			)
			DailymotionPlaylistIE__VALID_URL = λ.NewStr("(?:https?://)?(?:www\\.)?dailymotion\\.[a-z]{2,3}/playlist/(?P<id>x[0-9a-z]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): DailymotionPlaylistIE__VALID_URL,
			})
		}())
		DailymotionUserIE = λ.Cal(λ.TypeType, λ.NewStr("DailymotionUserIE"), λ.NewTuple(DailymotionBaseInfoExtractor), func() λ.Dict {
			var (
				DailymotionUserIE__VALID_URL λ.Object
			)
			DailymotionUserIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?dailymotion\\.[a-z]{2,3}/(?!(?:embed|swf|#|video|playlist)/)(?:(?:old/)?user/)?(?P<user>[^/]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): DailymotionUserIE__VALID_URL,
			})
		}())
	})
}
