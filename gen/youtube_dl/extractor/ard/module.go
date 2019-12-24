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
 * ard/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ard.py
 */

package ard

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ARDBetaMediathekIE       λ.Object
	ARDIE                    λ.Object
	ARDMediathekIE           λ.Object
	ExtractorError           λ.Object
	InfoExtractor            λ.Object
	ϒcompat_etree_fromstring λ.Object
	ϒdetermine_ext           λ.Object
	ϒint_or_none             λ.Object
	ϒparse_duration          λ.Object
	ϒqualities               λ.Object
	ϒstr_or_none             λ.Object
	ϒtry_get                 λ.Object
	ϒunified_strdate         λ.Object
	ϒunified_timestamp       λ.Object
	ϒupdate_url_query        λ.Object
	ϒurl_or_none             λ.Object
	ϒxpath_text              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒqualities = Ωutils.ϒqualities
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒxpath_text = Ωutils.ϒxpath_text
		ϒcompat_etree_fromstring = Ωcompat.ϒcompat_etree_fromstring
		ARDMediathekIE = λ.Cal(λ.TypeType, λ.NewStr("ARDMediathekIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ARDMediathekIE_IE_NAME             λ.Object
				ARDMediathekIE__VALID_URL          λ.Object
				ARDMediathekIE__extract_formats    λ.Object
				ARDMediathekIE__extract_media_info λ.Object
				ARDMediathekIE__real_extract       λ.Object
				ARDMediathekIE_suitable            λ.Object
			)
			ARDMediathekIE_IE_NAME = λ.NewStr("ARD:mediathek")
			ARDMediathekIE__VALID_URL = λ.NewStr("^https?://(?:(?:(?:www|classic)\\.)?ardmediathek\\.de|mediathek\\.(?:daserste|rbb-online)\\.de|one\\.ard\\.de)/(?:.*/)(?P<video_id>[0-9]+|[^0-9][^/\\?]+)[^/\\?]*(?:\\?.*)?")
			ARDMediathekIE_suitable = λ.NewFunction("suitable",
				[]λ.Param{
					{Name: "cls"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcls = λargs[0]
						ϒurl = λargs[1]
					)
					return func() λ.Object {
						if λ.IsTrue(λ.Cal(λ.GetAttr(ARDBetaMediathekIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, ARDMediathekIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			ARDMediathekIE_suitable = λ.Cal(λ.ClassMethodType, ARDMediathekIE_suitable)
			ARDMediathekIE__extract_media_info = λ.NewFunction("_extract_media_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "media_info_url"},
					{Name: "webpage"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒduration       λ.Object
						ϒformats        λ.Object
						ϒis_live        λ.Object
						ϒmedia_info     λ.Object
						ϒmedia_info_url = λargs[1]
						ϒself           = λargs[0]
						ϒsubtitle_url   λ.Object
						ϒsubtitles      λ.Object
						ϒthumbnail      λ.Object
						ϒvideo_id       = λargs[3]
						ϒwebpage        = λargs[2]
					)
					ϒmedia_info = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒmedia_info_url, ϒvideo_id, λ.NewStr("Downloading media JSON"))
					ϒformats = λ.Cal(λ.GetAttr(ϒself, "_extract_formats", nil), ϒmedia_info, ϒvideo_id)
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformats))) {
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, λ.NewStr("\"fsk\"")))) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("This video is only available after 20:00")), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						} else {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_geoblocked"))) {
								panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("This video is not available due to geo restriction")), λ.KWArgs{
									{Name: "expected", Value: λ.True},
								})))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒduration = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_duration")))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_previewImage"))
					ϒis_live = λ.NewBool(λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_isLive")) == λ.True)
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒsubtitle_url = λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_subtitleUrl"))
					if λ.IsTrue(ϒsubtitle_url) {
						λ.SetItem(ϒsubtitles, λ.NewStr("de"), λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("ext"): λ.NewStr("ttml"),
							λ.NewStr("url"): ϒsubtitle_url,
						})))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("duration"):  ϒduration,
						λ.NewStr("thumbnail"): ϒthumbnail,
						λ.NewStr("is_live"):   ϒis_live,
						λ.NewStr("formats"):   ϒformats,
						λ.NewStr("subtitles"): ϒsubtitles,
					})
				})
			ARDMediathekIE__extract_formats = λ.NewFunction("_extract_formats",
				[]λ.Param{
					{Name: "self"},
					{Name: "media_info"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext         λ.Object
						ϒf           λ.Object
						ϒformats     λ.Object
						ϒm           λ.Object
						ϒmedia       λ.Object
						ϒmedia_array λ.Object
						ϒmedia_info  = λargs[1]
						ϒnum         λ.Object
						ϒquality     λ.Object
						ϒself        = λargs[0]
						ϒserver      λ.Object
						ϒstream      λ.Object
						ϒstream_url  λ.Object
						ϒstream_urls λ.Object
						ϒtype_       λ.Object
						ϒvideo_id    = λargs[2]
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
						τmp4         λ.Object
						τmp5         λ.Object
					)
					ϒtype_ = λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_type"))
					ϒmedia_array = λ.Cal(λ.GetAttr(ϒmedia_info, "get", nil), λ.NewStr("_mediaArray"), λ.NewList())
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.EnumerateIteratorType, ϒmedia_array))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒnum = λ.GetItem(τmp2, λ.NewInt(0))
						ϒmedia = λ.GetItem(τmp2, λ.NewInt(1))
						τmp2 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒmedia, "get", nil), λ.NewStr("_mediaStreamArray"), λ.NewList()))
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒstream = τmp3
							ϒstream_urls = λ.Cal(λ.GetAttr(ϒstream, "get", nil), λ.NewStr("_stream"))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒstream_urls))) {
								continue
							}
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒstream_urls, λ.ListType)))) {
								ϒstream_urls = λ.NewList(ϒstream_urls)
							}
							ϒquality = λ.Cal(λ.GetAttr(ϒstream, "get", nil), λ.NewStr("_quality"))
							ϒserver = λ.Cal(λ.GetAttr(ϒstream, "get", nil), λ.NewStr("_server"))
							τmp4 = λ.Cal(λ.BuiltinIter, ϒstream_urls)
							for {
								if τmp5 = λ.NextDefault(τmp4, λ.AfterLast); τmp5 == λ.AfterLast {
									break
								}
								ϒstream_url = τmp5
								if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(ϒurl_or_none, ϒstream_url)))) {
									continue
								}
								ϒext = λ.Cal(ϒdetermine_ext, ϒstream_url)
								if λ.IsTrue(func() λ.Object {
									if λv := λ.Ne(ϒquality, λ.NewStr("auto")); !λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(λ.Contains(λ.NewTuple(
											λ.NewStr("f4m"),
											λ.NewStr("m3u8"),
										), ϒext))
									}
								}()) {
									continue
								}
								if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("f4m"))) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
										λ.Cal(ϒupdate_url_query, ϒstream_url, λ.NewDictWithTable(map[λ.Object]λ.Object{
											λ.NewStr("hdcore"): λ.NewStr("3.1.1"),
											λ.NewStr("plugin"): λ.NewStr("aasp-3.1.1.69.124"),
										})),
										ϒvideo_id,
									), λ.KWArgs{
										{Name: "f4m_id", Value: λ.NewStr("hds")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
										λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
											ϒstream_url,
											ϒvideo_id,
											λ.NewStr("mp4"),
										), λ.KWArgs{
											{Name: "m3u8_id", Value: λ.NewStr("hls")},
											{Name: "fatal", Value: λ.False},
										}))
									} else {
										if λ.IsTrue(func() λ.Object {
											if λv := ϒserver; !λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Cal(λ.GetAttr(ϒserver, "startswith", nil), λ.NewStr("rtmp"))
											}
										}()) {
											ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("url"):       ϒserver,
												λ.NewStr("play_path"): ϒstream_url,
												λ.NewStr("format_id"): λ.Mod(λ.NewStr("a%s-rtmp-%s"), λ.NewTuple(
													ϒnum,
													ϒquality,
												)),
											})
										} else {
											ϒf = λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("url"): ϒstream_url,
												λ.NewStr("format_id"): λ.Mod(λ.NewStr("a%s-%s-%s"), λ.NewTuple(
													ϒnum,
													ϒext,
													ϒquality,
												)),
											})
										}
										ϒm = λ.Cal(Ωre.ϒsearch, λ.NewStr("_(?P<width>\\d+)x(?P<height>\\d+)\\.mp4$"), ϒstream_url)
										if λ.IsTrue(ϒm) {
											λ.Cal(λ.GetAttr(ϒf, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
												λ.NewStr("width"):  λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("width"))),
												λ.NewStr("height"): λ.Cal(λ.IntType, λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("height"))),
											}))
										}
										if λ.IsTrue(λ.Eq(ϒtype_, λ.NewStr("audio"))) {
											λ.SetItem(ϒf, λ.NewStr("vcodec"), λ.NewStr("none"))
										}
										λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒf)
									}
								}
							}
						}
					}
					return ϒformats
				})
			ARDMediathekIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ERRORS         λ.Object
						QUALITIES      λ.Object
						ϒdescription   λ.Object
						ϒdoc           λ.Object
						ϒdocument_id   λ.Object
						ϒfid           λ.Object
						ϒfid_m         λ.Object
						ϒformats       λ.Object
						ϒfurl          λ.Object
						ϒinfo          λ.Object
						ϒm             λ.Object
						ϒmedia_streams λ.Object
						ϒmessage       λ.Object
						ϒnumid         λ.Object
						ϒpattern       λ.Object
						ϒself          = λargs[0]
						ϒthumbnail     λ.Object
						ϒtitle         λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						ϒwebpage       λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
						τmp2           λ.Object
					)
					ϒm = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdocument_id = λ.None
					ϒnumid = λ.Cal(Ωre.ϒsearch, λ.NewStr("documentId=([0-9]+)"), ϒurl)
					if λ.IsTrue(ϒnumid) {
						τmp0 = λ.Cal(λ.GetAttr(ϒnumid, "group", nil), λ.NewInt(1))
						ϒdocument_id = τmp0
						ϒvideo_id = τmp0
					} else {
						ϒvideo_id = λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("video_id"))
					}
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ERRORS = λ.NewTuple(
						λ.NewTuple(
							λ.NewStr(">Leider liegt eine Störung vor."),
							λ.NewStr("Video %s is unavailable"),
						),
						λ.NewTuple(
							λ.NewStr(">Der gewünschte Beitrag ist nicht mehr verfügbar.<"),
							λ.NewStr("Video %s is no longer available"),
						),
					)
					τmp0 = λ.Cal(λ.BuiltinIter, ERRORS)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒpattern = λ.GetItem(τmp2, λ.NewInt(0))
						ϒmessage = λ.GetItem(τmp2, λ.NewInt(1))
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwebpage, ϒpattern))) {
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(ϒmessage, ϒvideo_id)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
						}
					}
					if λ.IsTrue(λ.Cal(Ωre.ϒsearch, λ.NewStr("[\\?&]rss($|[=&])"), ϒurl)) {
						ϒdoc = λ.Cal(ϒcompat_etree_fromstring, λ.Cal(λ.GetAttr(ϒwebpage, "encode", nil), λ.NewStr("utf-8")))
						if λ.IsTrue(λ.Eq(λ.GetAttr(ϒdoc, "tag", nil), λ.NewStr("rss"))) {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.None), "_extract_rss", nil), ϒurl, ϒvideo_id, ϒdoc)
						}
					}
					ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewList(
						λ.NewStr("<h1(?:\\s+class=\"boxTopHeadline\")?>(.*?)</h1>"),
						λ.NewStr("<meta name=\"dcterms\\.title\" content=\"(.*?)\"/>"),
						λ.NewStr("<h4 class=\"headline\">(.*?)</h4>"),
						λ.NewStr("<title[^>]*>(.*?)</title>"),
					), ϒwebpage, λ.NewStr("title"))
					ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
						λ.NewStr("dcterms.abstract"),
						ϒwebpage,
						λ.NewStr("description"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒdescription == λ.None)) {
						ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
							λ.NewStr("description"),
							ϒwebpage,
							λ.NewStr("meta description"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					if λ.IsTrue(λ.NewBool(ϒdescription == λ.None)) {
						ϒdescription = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
							λ.NewStr("<p\\s+class=\"teasertext\">(.+?)</p>"),
							ϒwebpage,
							λ.NewStr("teaser text"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						})
					}
					ϒthumbnail = λ.Call(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒmedia_streams = λ.Cal(Ωre.ϒfindall, λ.NewStr("(?x)\n            mediaCollection\\.addMediaStream\\([0-9]+,\\s*[0-9]+,\\s*\"[^\"]*\",\\s*\n            \"([^\"]+)\""), ϒwebpage)
					if λ.IsTrue(ϒmedia_streams) {
						QUALITIES = λ.Cal(ϒqualities, λ.NewList(
							λ.NewStr("lo"),
							λ.NewStr("hi"),
							λ.NewStr("hq"),
						))
						ϒformats = λ.NewList()
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.SetType, ϒmedia_streams))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒfurl = τmp1
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒfurl, "endswith", nil), λ.NewStr(".f4m"))) {
								ϒfid = λ.NewStr("f4m")
							} else {
								ϒfid_m = λ.Cal(Ωre.ϒmatch, λ.NewStr(".*\\.([^.]+)\\.[^.]+$"), ϒfurl)
								ϒfid = func() λ.Object {
									if λ.IsTrue(ϒfid_m) {
										return λ.Cal(λ.GetAttr(ϒfid_m, "group", nil), λ.NewInt(1))
									} else {
										return λ.None
									}
								}()
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("quality"):   λ.Cal(QUALITIES, ϒfid),
								λ.NewStr("format_id"): ϒfid,
								λ.NewStr("url"):       ϒfurl,
							}))
						}
						λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
						ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("formats"): ϒformats,
						})
					} else {
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒdocument_id))) {
							ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("/play/(?:config|media)/(\\d+)"), ϒwebpage, λ.NewStr("media id"))
						}
						ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_extract_media_info", nil), λ.Mod(λ.NewStr("http://www.ardmediathek.de/play/media/%s"), ϒvideo_id), ϒwebpage, ϒvideo_id)
					}
					λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"): ϒvideo_id,
						λ.NewStr("title"): func() λ.Object {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("is_live"))) {
								return λ.Cal(λ.GetAttr(ϒself, "_live_title", nil), ϒtitle)
							} else {
								return ϒtitle
							}
						}(),
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"):   ϒthumbnail,
					}))
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):             ARDMediathekIE_IE_NAME,
				λ.NewStr("_VALID_URL"):          ARDMediathekIE__VALID_URL,
				λ.NewStr("_extract_formats"):    ARDMediathekIE__extract_formats,
				λ.NewStr("_extract_media_info"): ARDMediathekIE__extract_media_info,
				λ.NewStr("_real_extract"):       ARDMediathekIE__real_extract,
				λ.NewStr("suitable"):            ARDMediathekIE_suitable,
			})
		}())
		ARDIE = λ.Cal(λ.TypeType, λ.NewStr("ARDIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ARDIE__VALID_URL λ.Object
			)
			ARDIE__VALID_URL = λ.NewStr("(?P<mainurl>https?://(www\\.)?daserste\\.de/[^?#]+/videos/(?P<display_id>[^/?#]+)-(?P<id>[0-9]+))\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): ARDIE__VALID_URL,
			})
		}())
		ARDBetaMediathekIE = λ.Cal(λ.TypeType, λ.NewStr("ARDBetaMediathekIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ARDBetaMediathekIE__VALID_URL    λ.Object
				ARDBetaMediathekIE__real_extract λ.Object
			)
			ARDBetaMediathekIE__VALID_URL = λ.NewStr("https://(?:beta|www)\\.ardmediathek\\.de/[^/]+/(?:player|live)/(?P<video_id>[a-zA-Z0-9]+)(?:/(?P<display_id>[^/?#]+))?")
			ARDBetaMediathekIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdata         λ.Object
						ϒdata_json    λ.Object
						ϒdisplay_id   λ.Object
						ϒext          λ.Object
						ϒformat_url   λ.Object
						ϒformats      λ.Object
						ϒgeoblocked   λ.Object
						ϒmobj         λ.Object
						ϒquality      λ.Object
						ϒres          λ.Object
						ϒself         = λargs[0]
						ϒsubtitle_url λ.Object
						ϒsubtitles    λ.Object
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
						ϒwebpage      λ.Object
						ϒwidget       λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("video_id"))
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("display_id")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒdata_json = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("window\\.__APOLLO_STATE__\\s*=\\s*(\\{.*);\\n"), ϒwebpage, λ.NewStr("json"))
					ϒdata = λ.Cal(λ.GetAttr(ϒself, "_parse_json", nil), ϒdata_json, ϒdisplay_id)
					ϒres = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
					})
					ϒformats = λ.NewList()
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒgeoblocked = λ.False
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒdata, "values", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒwidget = τmp1
						if λ.IsTrue(λ.NewBool(λ.Cal(λ.GetAttr(ϒwidget, "get", nil), λ.NewStr("_geoblocked")) == λ.True)) {
							ϒgeoblocked = λ.True
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwidget, λ.NewStr("_duration")))) {
							λ.SetItem(ϒres, λ.NewStr("duration"), λ.Cal(ϒint_or_none, λ.GetItem(ϒwidget, λ.NewStr("_duration"))))
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwidget, λ.NewStr("clipTitle")))) {
							λ.SetItem(ϒres, λ.NewStr("title"), λ.GetItem(ϒwidget, λ.NewStr("clipTitle")))
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwidget, λ.NewStr("_previewImage")))) {
							λ.SetItem(ϒres, λ.NewStr("thumbnail"), λ.GetItem(ϒwidget, λ.NewStr("_previewImage")))
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwidget, λ.NewStr("broadcastedOn")))) {
							λ.SetItem(ϒres, λ.NewStr("timestamp"), λ.Cal(ϒunified_timestamp, λ.GetItem(ϒwidget, λ.NewStr("broadcastedOn"))))
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwidget, λ.NewStr("synopsis")))) {
							λ.SetItem(ϒres, λ.NewStr("description"), λ.GetItem(ϒwidget, λ.NewStr("synopsis")))
						}
						ϒsubtitle_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒwidget, "get", nil), λ.NewStr("_subtitleUrl")))
						if λ.IsTrue(ϒsubtitle_url) {
							λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), λ.NewStr("de"), λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("ext"): λ.NewStr("ttml"),
								λ.NewStr("url"): ϒsubtitle_url,
							}))
						}
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒwidget, λ.NewStr("_quality")))) {
							ϒformat_url = λ.Cal(ϒurl_or_none, λ.Cal(ϒtry_get, ϒwidget, λ.NewFunction("<lambda>",
								[]λ.Param{
									{Name: "x"},
								},
								0, false, false,
								func(λargs []λ.Object) λ.Object {
									var (
										ϒx = λargs[0]
									)
									return λ.GetItem(λ.GetItem(λ.GetItem(ϒx, λ.NewStr("_stream")), λ.NewStr("json")), λ.NewInt(0))
								})))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒformat_url))) {
								continue
							}
							ϒext = λ.Cal(ϒdetermine_ext, ϒformat_url)
							if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("f4m"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
									λ.Add(ϒformat_url, λ.NewStr("?hdcore=3.11.0")),
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "f4m_id", Value: λ.NewStr("hds")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								if λ.IsTrue(λ.Eq(ϒext, λ.NewStr("m3u8"))) {
									λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
										ϒformat_url,
										ϒvideo_id,
										λ.NewStr("mp4"),
									), λ.KWArgs{
										{Name: "m3u8_id", Value: λ.NewStr("hls")},
										{Name: "fatal", Value: λ.False},
									}))
								} else {
									if λ.IsTrue(ϒgeoblocked) {
										continue
									}
									ϒquality = λ.Cal(ϒstr_or_none, λ.Cal(λ.GetAttr(ϒwidget, "get", nil), λ.NewStr("_quality")))
									λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("format_id"): func() λ.Object {
											if λ.IsTrue(ϒquality) {
												return λ.Add(λ.NewStr("http-"), ϒquality)
											} else {
												return λ.NewStr("http")
											}
										}(),
										λ.NewStr("url"):        ϒformat_url,
										λ.NewStr("preference"): λ.NewInt(10),
									}))
								}
							}
						}
					}
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒformats)); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒgeoblocked
						}
					}()) {
						λ.Call(λ.GetAttr(ϒself, "raise_geo_restricted", nil), nil, λ.KWArgs{
							{Name: "msg", Value: λ.NewStr("This video is not available due to geoblocking")},
							{Name: "countries", Value: λ.NewList(λ.NewStr("DE"))},
						})
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					λ.Cal(λ.GetAttr(ϒres, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("subtitles"): ϒsubtitles,
						λ.NewStr("formats"):   ϒformats,
					}))
					return ϒres
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    ARDBetaMediathekIE__VALID_URL,
				λ.NewStr("_real_extract"): ARDBetaMediathekIE__real_extract,
			})
		}())
	})
}
