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
 * nova/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/nova.py
 */

package nova

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	NovaEmbedIE      λ.Object
	NovaIE           λ.Object
	ϒclean_html      λ.Object
	ϒint_or_none     λ.Object
	ϒjs_to_json      λ.Object
	ϒqualities       λ.Object
	ϒunified_strdate λ.Object
	ϒurl_or_none     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒjs_to_json = Ωutils.ϒjs_to_json
		ϒqualities = Ωutils.ϒqualities
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒurl_or_none = Ωutils.ϒurl_or_none
		NovaEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("NovaEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NovaEmbedIE__VALID_URL λ.Object
			)
			NovaEmbedIE__VALID_URL = λ.NewStr("https?://media\\.cms\\.nova\\.cz/embed/(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): NovaEmbedIE__VALID_URL,
			})
		}())
		NovaIE = λ.Cal(λ.TypeType, λ.NewStr("NovaIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				NovaIE__VALID_URL    λ.Object
				NovaIE__real_extract λ.Object
			)
			NovaIE__VALID_URL = λ.NewStr("https?://(?:[^.]+\\.)?(?P<site>tv(?:noviny)?|tn|novaplus|vymena|fanda|krasna|doma|prask)\\.nova\\.cz/(?:[^/]+/)+(?P<id>[^/]+?)(?:\\.html|/|$)")
			NovaIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						DEFAULT_SITE_ID λ.Object
						SITES           λ.Object
						ϒconfig         λ.Object
						ϒconfig_params  λ.Object
						ϒconfig_url     λ.Object
						ϒdescription    λ.Object
						ϒdisplay_id     λ.Object
						ϒembed_id       λ.Object
						ϒformats        λ.Object
						ϒm              λ.Object
						ϒmediafile      λ.Object
						ϒmobj           λ.Object
						ϒparams         λ.Object
						ϒplayer         λ.Object
						ϒself           = λargs[0]
						ϒsite           λ.Object
						ϒsite_id        λ.Object
						ϒthumbnail      λ.Object
						ϒtitle          λ.Object
						ϒupload_date    λ.Object
						ϒurl            = λargs[1]
						ϒvideo_id       λ.Object
						ϒvideo_url      λ.Object
						ϒwebpage        λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒsite = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("site"))
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒembed_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("<iframe[^>]+\\bsrc=[\"\\'](?:https?:)?//media\\.cms\\.nova\\.cz/embed/([^/?#&]+)"),
						ϒwebpage,
						λ.NewStr("embed url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒembed_id) {
						return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(λ.Mod(λ.NewStr("https://media.cms.nova.cz/embed/%s"), ϒembed_id)), λ.KWArgs{
							{Name: "ie", Value: λ.Cal(λ.GetAttr(NovaEmbedIE, "ie_key", nil))},
							{Name: "video_id", Value: ϒembed_id},
						})
					}
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewList(
						λ.NewStr("(?:media|video_id)\\s*:\\s*'(\\d+)'"),
						λ.NewStr("media=(\\d+)"),
						λ.NewStr("id=\"article_video_(\\d+)\""),
						λ.NewStr("id=\"player_(\\d+)\""),
					), ϒwebpage, λ.NewStr("video id"))
					ϒconfig_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("src=\"(https?://(?:tn|api)\\.nova\\.cz/bin/player/videojs/config\\.php\\?[^\"]+)\""),
						ϒwebpage,
						λ.NewStr("config url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					ϒconfig_params = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒconfig_url))) {
						ϒplayer = λ.Call(λ.GetAttr(ϒself, "_parse_json", nil), λ.NewArgs(
							λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("(?s)Player\\s*\\(.+?\\s*,\\s*({.+?\\bmedia\\b[\"\\']?\\s*:\\s*[\"\\']?\\d+.+?})\\s*\\)"),
								ϒwebpage,
								λ.NewStr("player"),
							), λ.KWArgs{
								{Name: "default", Value: λ.NewStr("{}")},
							}),
							ϒvideo_id,
						), λ.KWArgs{
							{Name: "transform_source", Value: ϒjs_to_json},
							{Name: "fatal", Value: λ.False},
						})
						if λ.IsTrue(ϒplayer) {
							ϒconfig_url = λ.Cal(ϒurl_or_none, λ.Cal(λ.GetAttr(ϒplayer, "get", nil), λ.NewStr("configUrl")))
							ϒparams = λ.Cal(λ.GetAttr(ϒplayer, "get", nil), λ.NewStr("configParams"))
							if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒparams, λ.DictType)) {
								ϒconfig_params = ϒparams
							}
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒconfig_url))) {
						DEFAULT_SITE_ID = λ.NewStr("23000")
						SITES = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("tvnoviny"): DEFAULT_SITE_ID,
							λ.NewStr("novaplus"): DEFAULT_SITE_ID,
							λ.NewStr("vymena"):   DEFAULT_SITE_ID,
							λ.NewStr("krasna"):   DEFAULT_SITE_ID,
							λ.NewStr("fanda"):    λ.NewStr("30"),
							λ.NewStr("tn"):       λ.NewStr("30"),
							λ.NewStr("doma"):     λ.NewStr("30"),
						})
						ϒsite_id = func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("site=(\\d+)"),
								ϒwebpage,
								λ.NewStr("site id"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(SITES, "get", nil), ϒsite, DEFAULT_SITE_ID)
							}
						}()
						ϒconfig_url = λ.NewStr("https://api.nova.cz/bin/player/videojs/config.php")
						ϒconfig_params = λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("site"):    ϒsite_id,
							λ.NewStr("media"):   ϒvideo_id,
							λ.NewStr("quality"): λ.NewInt(3),
							λ.NewStr("version"): λ.NewInt(1),
						})
					}
					ϒconfig = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						ϒconfig_url,
						ϒdisplay_id,
						λ.NewStr("Downloading config JSON"),
					), λ.KWArgs{
						{Name: "query", Value: ϒconfig_params},
						{Name: "transform_source", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "s"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒs = λargs[0]
								)
								return λ.GetItem(ϒs, λ.NewSlice(λ.Cal(λ.GetAttr(ϒs, "index", nil), λ.NewStr("{")), λ.Add(λ.Cal(λ.GetAttr(ϒs, "rindex", nil), λ.NewStr("}")), λ.NewInt(1)), λ.None))
							})},
					})
					ϒmediafile = λ.GetItem(ϒconfig, λ.NewStr("mediafile"))
					ϒvideo_url = λ.GetItem(ϒmediafile, λ.NewStr("src"))
					ϒm = λ.Cal(Ωre.ϒsearch, λ.NewStr("^(?P<url>rtmpe?://[^/]+/(?P<app>[^/]+?))/&*(?P<playpath>.+)$"), ϒvideo_url)
					if λ.IsTrue(ϒm) {
						ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"):         λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("url")),
							λ.NewStr("app"):         λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("app")),
							λ.NewStr("play_path"):   λ.Cal(λ.GetAttr(ϒm, "group", nil), λ.NewStr("playpath")),
							λ.NewStr("player_path"): λ.NewStr("http://tvnoviny.nova.cz/static/shared/app/videojs/video-js.swf"),
							λ.NewStr("ext"):         λ.NewStr("flv"),
						}))
					} else {
						ϒformats = λ.NewList(λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒvideo_url,
						}))
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒtitle = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒmediafile, "get", nil), λ.NewStr("meta"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "get", nil), λ.NewStr("title")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.Cal(λ.GetAttr(ϒself, "_og_search_title", nil), ϒwebpage)
						}
					}()
					ϒdescription = λ.Cal(ϒclean_html, λ.Call(λ.GetAttr(ϒself, "_og_search_description", nil), λ.NewArgs(ϒwebpage), λ.KWArgs{
						{Name: "default", Value: λ.None},
					}))
					ϒthumbnail = λ.Cal(λ.GetAttr(ϒconfig, "get", nil), λ.NewStr("poster"))
					if λ.IsTrue(λ.Eq(ϒsite, λ.NewStr("novaplus"))) {
						ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
							λ.NewStr("(\\d{1,2}-\\d{1,2}-\\d{4})$"),
							ϒdisplay_id,
							λ.NewStr("upload date"),
						), λ.KWArgs{
							{Name: "default", Value: λ.None},
						}))
					} else {
						if λ.IsTrue(λ.Eq(ϒsite, λ.NewStr("fanda"))) {
							ϒupload_date = λ.Cal(ϒunified_strdate, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("<span class=\"date_time\">(\\d{1,2}\\.\\d{1,2}\\.\\d{4})"),
								ϒwebpage,
								λ.NewStr("upload date"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							}))
						} else {
							ϒupload_date = λ.None
						}
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("display_id"):  ϒdisplay_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): ϒdescription,
						λ.NewStr("upload_date"): ϒupload_date,
						λ.NewStr("thumbnail"):   ϒthumbnail,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"):    NovaIE__VALID_URL,
				λ.NewStr("_real_extract"): NovaIE__real_extract,
			})
		}())
	})
}
