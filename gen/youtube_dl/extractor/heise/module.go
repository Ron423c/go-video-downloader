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
 * heise/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/heise.py
 */

package heise

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωkaltura "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/kaltura"
	Ωyoutube "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/youtube"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	HeiseIE        λ.Object
	InfoExtractor  λ.Object
	KalturaIE      λ.Object
	NO_DEFAULT     λ.Object
	YoutubeIE      λ.Object
	ϒdetermine_ext λ.Object
	ϒint_or_none   λ.Object
	ϒparse_iso8601 λ.Object
	ϒsmuggle_url   λ.Object
	ϒxpath_text    λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		KalturaIE = Ωkaltura.KalturaIE
		YoutubeIE = Ωyoutube.YoutubeIE
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		NO_DEFAULT = Ωutils.NO_DEFAULT
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒsmuggle_url = Ωutils.ϒsmuggle_url
		ϒxpath_text = Ωutils.ϒxpath_text
		HeiseIE = λ.Cal(λ.TypeType, λ.NewStr("HeiseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				HeiseIE__VALID_URL    λ.Object
				HeiseIE__real_extract λ.Object
			)
			HeiseIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?heise\\.de/(?:[^/]+/)+[^/]+-(?P<id>[0-9]+)\\.html")
			HeiseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_make_kaltura_result λ.Object
						ϒcontainer_id         λ.Object
						ϒdescription          λ.Object
						ϒdoc                  λ.Object
						ϒext                  λ.Object
						ϒextract_title        λ.Object
						ϒformats              λ.Object
						ϒheight               λ.Object
						ϒkaltura_id           λ.Object
						ϒkaltura_url          λ.Object
						ϒlabel                λ.Object
						ϒself                 = λargs[0]
						ϒsequenz_id           λ.Object
						ϒsource_node          λ.Object
						ϒtitle                λ.Object
						ϒurl                  = λargs[1]
						ϒvideo_id             λ.Object
						ϒvideo_url            λ.Object
						ϒwebpage              λ.Object
						ϒyt_urls              λ.Object
						τmp0                  λ.Object
						τmp1                  λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
					ϒextract_title = λ.NewFunction("extract_title",
						[]λ.Param{
							{Name: "default", Def: NO_DEFAULT},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdefault = λargs[0]
								ϒtitle   λ.Object
							)
							ϒtitle = λ.Call(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewArgs(
								λ.NewTuple(
									λ.NewStr("fulltitle"),
									λ.NewStr("title"),
								),
								ϒwebpage,
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(func() λ.Object {
								if λv := λ.NewBool(!λ.IsTrue(ϒtitle)); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Eq(ϒtitle, λ.NewStr("c't"))
								}
							}()) {
								ϒtitle = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
									λ.NewStr("<div[^>]+class=\"videoplayerjw\"[^>]+data-title=\"([^\"]+)\""),
									ϒwebpage,
									λ.NewStr("title"),
								), λ.KWArgs{
									{Name: "default", Value: λ.None},
								})
							}
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒtitle))) {
								ϒtitle = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
									λ.NewStr("<h1[^>]+\\bclass=[\"\\']article_page_title[^>]+>(.+?)<"),
									ϒwebpage,
									λ.NewStr("title"),
								), λ.KWArgs{
									{Name: "default", Value: ϒdefault},
								})
							}
							return ϒtitle
						})
					ϒtitle = λ.Call(ϒextract_title, nil, λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒdescription = func() λ.Object {
						if λv := λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("description"), ϒwebpage)
						}
					}()
					ϒ_make_kaltura_result = λ.NewFunction("_make_kaltura_result",
						[]λ.Param{
							{Name: "kaltura_url"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒkaltura_url = λargs[0]
							)
							return λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("_type"): λ.NewStr("url_transparent"),
								λ.NewStr("url"): λ.Cal(ϒsmuggle_url, ϒkaltura_url, λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("source_url"): ϒurl,
								})),
								λ.NewStr("ie_key"):      λ.Cal(λ.GetAttr(KalturaIE, "ie_key", nil)),
								λ.NewStr("title"):       ϒtitle,
								λ.NewStr("description"): ϒdescription,
							})
						})
					ϒkaltura_url = λ.Cal(λ.GetAttr(KalturaIE, "_extract_url", nil), ϒwebpage)
					if λ.IsTrue(ϒkaltura_url) {
						return λ.Cal(ϒ_make_kaltura_result, ϒkaltura_url)
					}
					ϒkaltura_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("entry-id=([\"\\'])(?P<id>(?:(?!\\1).)+)\\1"),
						ϒwebpage,
						λ.NewStr("kaltura id"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
						{Name: "group", Value: λ.NewStr("id")},
					})
					if λ.IsTrue(ϒkaltura_id) {
						return λ.Cal(ϒ_make_kaltura_result, λ.Mod(λ.NewStr("kaltura:2238431:%s"), ϒkaltura_id))
					}
					ϒyt_urls = λ.Cal(λ.GetAttr(YoutubeIE, "_extract_urls", nil), ϒwebpage)
					if λ.IsTrue(ϒyt_urls) {
						return λ.Call(λ.GetAttr(ϒself, "playlist_from_matches", nil), λ.NewArgs(
							ϒyt_urls,
							ϒvideo_id,
							ϒtitle,
						), λ.KWArgs{
							{Name: "ie", Value: λ.Cal(λ.GetAttr(YoutubeIE, "ie_key", nil))},
						})
					}
					ϒtitle = λ.Cal(ϒextract_title)
					ϒcontainer_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<div class=\"videoplayerjw\"[^>]+data-container=\"([0-9]+)\""), ϒwebpage, λ.NewStr("container ID"))
					ϒsequenz_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("<div class=\"videoplayerjw\"[^>]+data-sequenz=\"([0-9]+)\""), ϒwebpage, λ.NewStr("sequenz ID"))
					ϒdoc = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						λ.NewStr("http://www.heise.de/videout/feed"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("container"): ϒcontainer_id,
							λ.NewStr("sequenz"):   ϒsequenz_id,
						})},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒdoc, "findall", nil), λ.NewStr(".//{http://rss.jwpcdn.com/}source")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒsource_node = τmp1
						ϒlabel = λ.GetItem(λ.GetAttr(ϒsource_node, "attrib", nil), λ.NewStr("label"))
						ϒheight = λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("^(.*?_)?([0-9]+)p$"),
							ϒlabel,
							λ.NewStr("height"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}))
						ϒvideo_url = λ.GetItem(λ.GetAttr(ϒsource_node, "attrib", nil), λ.NewStr("file"))
						ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url, λ.NewStr(""))
						λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):         ϒvideo_url,
							λ.NewStr("format_note"): ϒlabel,
							λ.NewStr("format_id"): λ.Mod(λ.NewStr("%s_%s"), λ.NewTuple(
								ϒext,
								ϒlabel,
							)),
							λ.NewStr("height"): ϒheight,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("thumbnail"): func() λ.Object {
							if λv := λ.Cal(ϒxpath_text, ϒdoc, λ.NewStr(".//{http://rss.jwpcdn.com/}image")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒself, "_og_search_thumbnail", nil), ϒwebpage)
							}
						}(),
						λ.NewStr("timestamp"): λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒself, "_html_search_meta", nil), λ.NewStr("date"), ϒwebpage)),
						λ.NewStr("formats"):   ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    HeiseIE__VALID_URL,
				λ.NewStr("_real_extract"): HeiseIE__real_extract,
			})
		}())
	})
}
