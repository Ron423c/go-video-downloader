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
 * ign/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ign.py
 */

package ign

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	IGNIE          λ.Object
	InfoExtractor  λ.Object
	OneUPIE        λ.Object
	PCMagIE        λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		IGNIE = λ.Cal(λ.TypeType, λ.NewStr("IGNIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				IGNIE_IE_NAME           λ.Object
				IGNIE__API_URL_TEMPLATE λ.Object
				IGNIE__EMBED_RE         λ.Object
				IGNIE__TESTS            λ.Object
				IGNIE__VALID_URL        λ.Object
				IGNIE__find_video_id    λ.Object
				IGNIE__get_video_info   λ.Object
				IGNIE__real_extract     λ.Object
			)
			λ.NewStr("\n    Extractor for some of the IGN sites, like www.ign.com, es.ign.com de.ign.com.\n    Some videos of it.ign.com are also supported\n    ")
			IGNIE__VALID_URL = λ.NewStr("https?://.+?\\.ign\\.com/(?:[^/]+/)?(?P<type>videos|show_videos|articles|feature|(?:[^/]+/\\d+/video))(/.+)?/(?P<name_or_id>.+)")
			IGNIE_IE_NAME = λ.NewStr("ign.com")
			IGNIE__API_URL_TEMPLATE = λ.NewStr("http://apis.ign.com/video/v3/videos/%s")
			IGNIE__EMBED_RE = λ.NewStr("<iframe[^>]+?[\"\\']((?:https?:)?//.+?\\.ign\\.com.+?/embed.+?)[\"\\']")
			IGNIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ign.com/videos/2013/06/05/the-last-of-us-review"),
					λ.NewStr("md5"): λ.NewStr("febda82c4bafecd2d44b6e1a18a595f8"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("8f862beef863986b2785559b9e1aa599"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("The Last of Us Review"),
						λ.NewStr("description"): λ.NewStr("md5:c8946d4260a4d43a00d5ae8ed998870c"),
						λ.NewStr("timestamp"):   λ.NewInt(1370440800),
						λ.NewStr("upload_date"): λ.NewStr("20130605"),
						λ.NewStr("uploader_id"): λ.NewStr("cberidon@ign.com"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://me.ign.com/en/feature/15775/100-little-things-in-gta-5-that-will-blow-your-mind"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): λ.NewStr("100-little-things-in-gta-5-that-will-blow-your-mind"),
					}),
					λ.NewStr("playlist"): λ.NewList(
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("5ebbd138523268b93c9141af17bec937"),
								λ.NewStr("ext"):         λ.NewStr("mp4"),
								λ.NewStr("title"):       λ.NewStr("GTA 5 Video Review"),
								λ.NewStr("description"): λ.NewStr("Rockstar drops the mic on this generation of games. Watch our review of the masterly Grand Theft Auto V."),
								λ.NewStr("timestamp"):   λ.NewInt(1379339880),
								λ.NewStr("upload_date"): λ.NewStr("20130916"),
								λ.NewStr("uploader_id"): λ.NewStr("danieljkrupa@gmail.com"),
							}),
						}),
						λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("id"):          λ.NewStr("638672ee848ae4ff108df2a296418ee2"),
								λ.NewStr("ext"):         λ.NewStr("mp4"),
								λ.NewStr("title"):       λ.NewStr("26 Twisted Moments from GTA 5 in Slow Motion"),
								λ.NewStr("description"): λ.NewStr("The twisted beauty of GTA 5 in stunning slow motion."),
								λ.NewStr("timestamp"):   λ.NewInt(1386878820),
								λ.NewStr("upload_date"): λ.NewStr("20131212"),
								λ.NewStr("uploader_id"): λ.NewStr("togilvie@ign.com"),
							}),
						}),
					),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.ign.com/articles/2014/08/15/rewind-theater-wild-trailer-gamescom-2014?watch"),
					λ.NewStr("md5"): λ.NewStr("618fedb9c901fd086f6f093564ef8558"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("078fdd005f6d3c02f63d795faa1b984f"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Rewind Theater - Wild Trailer Gamescom 2014"),
						λ.NewStr("description"): λ.NewStr("Brian and Jared explore Michel Ancel's captivating new preview."),
						λ.NewStr("timestamp"):   λ.NewInt(1408047180),
						λ.NewStr("upload_date"): λ.NewStr("20140814"),
						λ.NewStr("uploader_id"): λ.NewStr("jamesduggan1990@gmail.com"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://me.ign.com/en/videos/112203/video/how-hitman-aims-to-be-different-than-every-other-s"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://me.ign.com/ar/angry-birds-2/106533/video/lrd-ldyy-lwl-lfylm-angry-birds"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://www.ign.com/articles/2017/06/08/new-ducktales-short-donalds-birthday-doesnt-go-as-planned"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			IGNIE__find_video_id = λ.NewFunction("_find_video_id",
				[]λ.Param{
					{Name: "self"},
					{Name: "webpage"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒres_id  λ.Object
						ϒself    = λargs[0]
						ϒwebpage = λargs[1]
					)
					ϒres_id = λ.NewList(
						λ.NewStr("\"video_id\"\\s*:\\s*\"(.*?)\""),
						λ.NewStr("class=\"hero-poster[^\"]*?\"[^>]*id=\"(.+?)\""),
						λ.NewStr("data-video-id=\"(.+?)\""),
						λ.NewStr("<object id=\"vid_(.+?)\""),
						λ.NewStr("<meta name=\"og:image\" content=\".*/(.+?)-(.+?)/.+.jpg\""),
						λ.NewStr("videoId&quot;\\s*:\\s*&quot;(.+?)&quot;"),
						λ.NewStr("videoId[\"\\']\\s*:\\s*[\"\\']([^\"\\']+?)[\"\\']"),
					)
					return λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						ϒres_id,
						ϒwebpage,
						λ.NewStr("video id"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
				})
			IGNIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒentries       λ.Object
						ϒmobj          λ.Object
						ϒmultiple_urls λ.Object
						ϒname_or_id    λ.Object
						ϒpage_type     λ.Object
						ϒself          = λargs[0]
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒwebpage       λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒname_or_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("name_or_id"))
					ϒpage_type = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("type"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒname_or_id)
					if λ.IsTrue(λ.Ne(ϒpage_type, λ.NewStr("video"))) {
						ϒmultiple_urls = λ.Cal(Ωre.ϒfindall, λ.NewStr("<param name=\"flashvars\"[^>]*value=\"[^\"]*?url=(https?://www\\.ign\\.com/videos/.*?)[\"&]"), ϒwebpage)
						if λ.IsTrue(ϒmultiple_urls) {
							ϒentries = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
								nil,
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
										var (
											ϒu   λ.Object
											τmp0 λ.Object
											τmp1 λ.Object
										)
										τmp0 = λ.Cal(λ.BuiltinIter, ϒmultiple_urls)
										for {
											if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
												break
											}
											ϒu = τmp1
											λgen.Yield(λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒu), λ.KWArgs{
												{Name: "ie", Value: λ.NewStr("IGN")},
											}))
										}
										return λ.None
									})
								})))
							return λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("_type"):   λ.NewStr("playlist"),
								λ.NewStr("id"):      ϒname_or_id,
								λ.NewStr("entries"): ϒentries,
							})
						}
					}
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_find_video_id", nil), ϒwebpage)
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒvideo_id))) {
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.GetAttr(ϒself, "_EMBED_RE", nil), ϒwebpage, λ.NewStr("embed url")))
					}
					return λ.Cal(λ.GetAttr(ϒself, "_get_video_info", nil), ϒvideo_id)
				})
			IGNIE__get_video_info = λ.NewFunction("_get_video_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒapi_data   λ.Object
						ϒasset      λ.Object
						ϒf4m_url    λ.Object
						ϒformats    λ.Object
						ϒm3u8_url   λ.Object
						ϒmetadata   λ.Object
						ϒself       = λargs[0]
						ϒthumbnails λ.Object
						ϒvideo_id   = λargs[1]
						τmp0        λ.Object
						τmp1        λ.Object
					)
					ϒapi_data = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.GetAttr(ϒself, "_API_URL_TEMPLATE", nil), ϒvideo_id), ϒvideo_id)
					ϒformats = λ.NewList()
					ϒm3u8_url = λ.Cal(λ.GetAttr(λ.GetItem(ϒapi_data, λ.NewStr("refs")), "get", nil), λ.NewStr("m3uUrl"))
					if λ.IsTrue(ϒm3u8_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
							ϒm3u8_url,
							ϒvideo_id,
							λ.NewStr("mp4"),
							λ.NewStr("m3u8_native"),
						), λ.KWArgs{
							{Name: "m3u8_id", Value: λ.NewStr("hls")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					ϒf4m_url = λ.Cal(λ.GetAttr(λ.GetItem(ϒapi_data, λ.NewStr("refs")), "get", nil), λ.NewStr("f4mUrl"))
					if λ.IsTrue(ϒf4m_url) {
						λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
							ϒf4m_url,
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "f4m_id", Value: λ.NewStr("hds")},
							{Name: "fatal", Value: λ.False},
						}))
					}
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒapi_data, λ.NewStr("assets")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒasset = τmp1
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):    λ.GetItem(ϒasset, λ.NewStr("url")),
							λ.NewStr("tbr"):    λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("actual_bitrate_kbps")),
							λ.NewStr("fps"):    λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("frame_rate")),
							λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("height"))),
							λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒasset, "get", nil), λ.NewStr("width"))),
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgen λ.Generator) λ.Object {
								var (
									ϒthumbnail λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒapi_data, "get", nil), λ.NewStr("thumbnails"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒthumbnail = τmp1
									λgen.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): λ.GetItem(ϒthumbnail, λ.NewStr("url")),
									}))
								}
								return λ.None
							})
						})))
					ϒmetadata = λ.GetItem(ϒapi_data, λ.NewStr("metadata"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒapi_data, "get", nil), λ.NewStr("videoId")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						λ.NewStr("title"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("longTitle")); λ.IsTrue(λv) {
								return λv
							} else if λv := λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("name")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetItem(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("title"))
							}
						}(),
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("description")),
						λ.NewStr("timestamp"):   λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("publishDate"))),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("display_id"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("slug")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						λ.NewStr("uploader_id"): λ.Cal(λ.GetAttr(ϒmetadata, "get", nil), λ.NewStr("creator")),
						λ.NewStr("thumbnails"):  ϒthumbnails,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):           IGNIE_IE_NAME,
				λ.NewStr("_API_URL_TEMPLATE"): IGNIE__API_URL_TEMPLATE,
				λ.NewStr("_EMBED_RE"):         IGNIE__EMBED_RE,
				λ.NewStr("_TESTS"):            IGNIE__TESTS,
				λ.NewStr("_VALID_URL"):        IGNIE__VALID_URL,
				λ.NewStr("_find_video_id"):    IGNIE__find_video_id,
				λ.NewStr("_get_video_info"):   IGNIE__get_video_info,
				λ.NewStr("_real_extract"):     IGNIE__real_extract,
			})
		}())
		OneUPIE = λ.Cal(λ.TypeType, λ.NewStr("OneUPIE"), λ.NewTuple(IGNIE), func() λ.Dict {
			var (
				OneUPIE__VALID_URL λ.Object
			)
			OneUPIE__VALID_URL = λ.NewStr("https?://gamevideos\\.1up\\.com/(?P<type>video)/id/(?P<name_or_id>.+)\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): OneUPIE__VALID_URL,
			})
		}())
		PCMagIE = λ.Cal(λ.TypeType, λ.NewStr("PCMagIE"), λ.NewTuple(IGNIE), func() λ.Dict {
			var (
				PCMagIE__VALID_URL λ.Object
			)
			PCMagIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?pcmag\\.com/(?P<type>videos|article2)(/.+)?/(?P<name_or_id>.+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): PCMagIE__VALID_URL,
			})
		}())
	})
}