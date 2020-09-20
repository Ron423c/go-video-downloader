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
 * twitch/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/twitch.py
 */

package twitch

import (
	Ωcollections "github.com/tenta-browser/go-video-downloader/gen/collections"
	Ωjson "github.com/tenta-browser/go-video-downloader/gen/json"
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                 λ.Object
	InfoExtractor                  λ.Object
	TwitchBaseIE                   λ.Object
	TwitchClipsIE                  λ.Object
	TwitchCollectionIE             λ.Object
	TwitchGraphQLBaseIE            λ.Object
	TwitchPlaylistBaseIE           λ.Object
	TwitchStreamIE                 λ.Object
	TwitchVideosClipsIE            λ.Object
	TwitchVideosCollectionsIE      λ.Object
	TwitchVideosIE                 λ.Object
	TwitchVodIE                    λ.Object
	ϒclean_html                    λ.Object
	ϒcompat_kwargs                 λ.Object
	ϒcompat_parse_qs               λ.Object
	ϒcompat_str                    λ.Object
	ϒcompat_urllib_parse_urlencode λ.Object
	ϒcompat_urllib_parse_urlparse  λ.Object
	ϒfloat_or_none                 λ.Object
	ϒint_or_none                   λ.Object
	ϒparse_duration                λ.Object
	ϒparse_iso8601                 λ.Object
	ϒqualities                     λ.Object
	ϒtry_get                       λ.Object
	ϒunified_timestamp             λ.Object
	ϒupdate_url_query              λ.Object
	ϒurl_or_none                   λ.Object
	ϒurljoin                       λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_kwargs = Ωcompat.ϒcompat_kwargs
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒcompat_urllib_parse_urlencode = Ωcompat.ϒcompat_urllib_parse_urlencode
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒclean_html = Ωutils.ϒclean_html
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒqualities = Ωutils.ϒqualities
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒurl_or_none = Ωutils.ϒurl_or_none
		ϒurljoin = Ωutils.ϒurljoin
		TwitchBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				TwitchBaseIE__CLIENT_ID       λ.Object
				TwitchBaseIE__NETRC_MACHINE   λ.Object
				TwitchBaseIE__login           λ.Object
				TwitchBaseIE__real_initialize λ.Object
			)
			TwitchBaseIE__CLIENT_ID = λ.StrLiteral("kimne78kx3ncx6brgo4mv6wki5h1ko")
			TwitchBaseIE__NETRC_MACHINE = λ.StrLiteral("twitch")
			TwitchBaseIE__real_initialize = λ.NewFunction("_real_initialize",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					λ.Calm(ϒself, "_login")
					return λ.None
				})
			TwitchBaseIE__login = λ.NewFunction("_login",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfail          λ.Object
						ϒhandle        λ.Object
						ϒlogin_page    λ.Object
						ϒlogin_step    λ.Object
						ϒpassword      λ.Object
						ϒredirect_page λ.Object
						ϒself          = λargs[0]
						ϒtfa_token     λ.Object
						ϒusername      λ.Object
						τmp0           λ.Object
					)
					τmp0 = λ.Calm(ϒself, "_get_login_info")
					ϒusername = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒpassword = λ.GetItem(τmp0, λ.IntLiteral(1))
					if ϒusername == λ.None {
						return λ.None
					}
					ϒfail = λ.NewFunction("fail",
						[]λ.Param{
							{Name: "message"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒmessage = λargs[0]
							)
							panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.Mod(λ.StrLiteral("Unable to login. Twitch said: %s"), ϒmessage)), λ.KWArgs{
								{Name: "expected", Value: λ.True},
							})))
							return λ.None
						})
					ϒlogin_step = λ.NewFunction("login_step",
						[]λ.Param{
							{Name: "page"},
							{Name: "urlh"},
							{Name: "note"},
							{Name: "data"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒdata         = λargs[3]
								ϒerror        λ.Object
								ϒform         λ.Object
								ϒheaders      λ.Object
								ϒnote         = λargs[2]
								ϒpage         = λargs[0]
								ϒpage_url     λ.Object
								ϒpost_url     λ.Object
								ϒredirect_url λ.Object
								ϒresponse     λ.Object
								ϒurlh         = λargs[1]
							)
							ϒform = λ.Calm(ϒself, "_hidden_inputs", ϒpage)
							λ.Calm(ϒform, "update", ϒdata)
							ϒpage_url = λ.Calm(ϒurlh, "geturl")
							ϒpost_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("<form[^>]+action=([\"\\'])(?P<url>.+?)\\1"),
								ϒpage,
								λ.StrLiteral("post url"),
							), λ.KWArgs{
								{Name: "default", Value: λ.GetAttr(ϒself, "_LOGIN_POST_URL", nil)},
								{Name: "group", Value: λ.StrLiteral("url")},
							})
							ϒpost_url = λ.Cal(ϒurljoin, ϒpage_url, ϒpost_url)
							ϒheaders = λ.DictLiteral(map[string]λ.Object{
								"Referer":      ϒpage_url,
								"Origin":       ϒpage_url,
								"Content-Type": λ.StrLiteral("text/plain;charset=UTF-8"),
							})
							ϒresponse = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								ϒpost_url,
								λ.None,
								ϒnote,
							), λ.KWArgs{
								{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, ϒform), "encode")},
								{Name: "headers", Value: ϒheaders},
								{Name: "expected_status", Value: λ.IntLiteral(400)},
							})
							ϒerror = func() λ.Object {
								if λv := λ.Calm(ϒresponse, "get", λ.StrLiteral("error_description")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.Calm(ϒresponse, "get", λ.StrLiteral("error_code"))
								}
							}()
							if λ.IsTrue(ϒerror) {
								λ.Cal(ϒfail, ϒerror)
							}
							if λ.Contains(λ.Calm(ϒresponse, "get", λ.StrLiteral("message"), λ.StrLiteral("")), λ.StrLiteral("Authenticated successfully")) {
								return λ.NewTuple(
									λ.None,
									λ.None,
								)
							}
							ϒredirect_url = λ.Cal(ϒurljoin, ϒpost_url, func() λ.Object {
								if λv := λ.Calm(ϒresponse, "get", λ.StrLiteral("redirect")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.GetItem(ϒresponse, λ.StrLiteral("redirect_path"))
								}
							}())
							return λ.Call(λ.GetAttr(ϒself, "_download_webpage_handle", nil), λ.NewArgs(
								ϒredirect_url,
								λ.None,
								λ.StrLiteral("Downloading login redirect page"),
							), λ.KWArgs{
								{Name: "headers", Value: ϒheaders},
							})
						})
					τmp0 = λ.Calm(ϒself, "_download_webpage_handle", λ.GetAttr(ϒself, "_LOGIN_FORM_URL", nil), λ.None, λ.StrLiteral("Downloading login page"))
					ϒlogin_page = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒhandle = λ.GetItem(τmp0, λ.IntLiteral(1))
					if λ.Contains(ϒlogin_page, λ.StrLiteral("blacklist_message")) {
						λ.Cal(ϒfail, λ.Cal(ϒclean_html, ϒlogin_page))
					}
					τmp0 = λ.Cal(ϒlogin_step, ϒlogin_page, ϒhandle, λ.StrLiteral("Logging in"), λ.DictLiteral(map[string]λ.Object{
						"username":  ϒusername,
						"password":  ϒpassword,
						"client_id": λ.GetAttr(ϒself, "_CLIENT_ID", nil),
					}))
					ϒredirect_page = λ.GetItem(τmp0, λ.IntLiteral(0))
					ϒhandle = λ.GetItem(τmp0, λ.IntLiteral(1))
					if !λ.IsTrue(ϒredirect_page) {
						return λ.None
					}
					if λ.Cal(Ωre.ϒsearch, λ.StrLiteral("(?i)<form[^>]+id=\"two-factor-submit\""), ϒredirect_page) != λ.None {
						ϒtfa_token = λ.Calm(ϒself, "_get_tfa_info", λ.StrLiteral("two-factor authentication token"))
						λ.Cal(ϒlogin_step, ϒredirect_page, ϒhandle, λ.StrLiteral("Submitting TFA token"), λ.DictLiteral(map[string]λ.Object{
							"authy_token":  ϒtfa_token,
							"remember_2fa": λ.StrLiteral("true"),
						}))
					}
					return λ.None
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_CLIENT_ID":       TwitchBaseIE__CLIENT_ID,
				"_NETRC_MACHINE":   TwitchBaseIE__NETRC_MACHINE,
				"_login":           TwitchBaseIE__login,
				"_real_initialize": TwitchBaseIE__real_initialize,
			})
		}())
		TwitchVodIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVodIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchVodIE__VALID_URL λ.Object
			)
			TwitchVodIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/(?:[^/]+/v(?:ideo)?|videos)/|\n                            player\\.twitch\\.tv/\\?.*?\\bvideo=v?\n                        )\n                        (?P<id>\\d+)\n                    ")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchVodIE__VALID_URL,
			})
		}())
		TwitchGraphQLBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchGraphQLBaseIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		TwitchCollectionIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchCollectionIE"), λ.NewTuple(TwitchGraphQLBaseIE), func() λ.Dict {
			var (
				TwitchCollectionIE__VALID_URL λ.Object
			)
			TwitchCollectionIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|go|m)\\.)?twitch\\.tv/collections/(?P<id>[^/]+)")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchCollectionIE__VALID_URL,
			})
		}())
		TwitchPlaylistBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchPlaylistBaseIE"), λ.NewTuple(TwitchGraphQLBaseIE), func() λ.Dict {

			return λ.ClassDictLiteral(map[λ.Object]λ.Object{})
		}())
		TwitchVideosIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVideosIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchVideosIE_Broadcast          λ.Object
				TwitchVideosIE__DEFAULT_BROADCAST λ.Object
				TwitchVideosIE__DEFAULT_SORTED_BY λ.Object
				TwitchVideosIE__VALID_URL         λ.Object
				TwitchVideosIE_suitable           λ.Object
			)
			TwitchVideosIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|go|m)\\.)?twitch\\.tv/(?P<id>[^/]+)/(?:videos|profile)")
			TwitchVideosIE_Broadcast = λ.Cal(Ωcollections.ϒnamedtuple, λ.StrLiteral("Broadcast"), λ.NewList(
				λ.StrLiteral("type"),
				λ.StrLiteral("label"),
			))
			TwitchVideosIE__DEFAULT_BROADCAST = λ.Cal(TwitchVideosIE_Broadcast, λ.None, λ.StrLiteral("All Videos"))
			TwitchVideosIE__DEFAULT_SORTED_BY = λ.StrLiteral("Date")
			TwitchVideosIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒie  λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										TwitchVideosClipsIE,
										TwitchVideosCollectionsIE,
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒie = τmp1
										λgy.Yield(λ.Calm(ϒie, "suitable", ϒurl))
									}
									return λ.None
								})
							})))) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, TwitchVideosIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			TwitchVideosIE_suitable = λ.Cal(λ.ClassMethodType, TwitchVideosIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"Broadcast":          TwitchVideosIE_Broadcast,
				"_DEFAULT_BROADCAST": TwitchVideosIE__DEFAULT_BROADCAST,
				"_DEFAULT_SORTED_BY": TwitchVideosIE__DEFAULT_SORTED_BY,
				"_VALID_URL":         TwitchVideosIE__VALID_URL,
				"suitable":           TwitchVideosIE_suitable,
			})
		}())
		TwitchVideosClipsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVideosClipsIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchVideosClipsIE_Clip          λ.Object
				TwitchVideosClipsIE__DEFAULT_CLIP λ.Object
				TwitchVideosClipsIE__VALID_URL    λ.Object
			)
			TwitchVideosClipsIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|go|m)\\.)?twitch\\.tv/(?P<id>[^/]+)/(?:clips|videos/*?\\?.*?\\bfilter=clips)")
			TwitchVideosClipsIE_Clip = λ.Cal(Ωcollections.ϒnamedtuple, λ.StrLiteral("Clip"), λ.NewList(
				λ.StrLiteral("filter"),
				λ.StrLiteral("label"),
			))
			TwitchVideosClipsIE__DEFAULT_CLIP = λ.Cal(TwitchVideosClipsIE_Clip, λ.StrLiteral("LAST_WEEK"), λ.StrLiteral("Top 7D"))
			return λ.ClassDictLiteral(map[string]λ.Object{
				"Clip":          TwitchVideosClipsIE_Clip,
				"_DEFAULT_CLIP": TwitchVideosClipsIE__DEFAULT_CLIP,
				"_VALID_URL":    TwitchVideosClipsIE__VALID_URL,
			})
		}())
		TwitchVideosCollectionsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchVideosCollectionsIE"), λ.NewTuple(TwitchPlaylistBaseIE), func() λ.Dict {
			var (
				TwitchVideosCollectionsIE__VALID_URL λ.Object
			)
			TwitchVideosCollectionsIE__VALID_URL = λ.StrLiteral("https?://(?:(?:www|go|m)\\.)?twitch\\.tv/(?P<id>[^/]+)/videos/*?\\?.*?\\bfilter=collections")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchVideosCollectionsIE__VALID_URL,
			})
		}())
		TwitchStreamIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchStreamIE"), λ.NewTuple(TwitchGraphQLBaseIE), func() λ.Dict {
			var (
				TwitchStreamIE__VALID_URL λ.Object
				TwitchStreamIE_suitable   λ.Object
			)
			TwitchStreamIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/|\n                            player\\.twitch\\.tv/\\?.*?\\bchannel=\n                        )\n                        (?P<id>[^/#?]+)\n                    ")
			TwitchStreamIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
							nil,
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
									var (
										ϒie  λ.Object
										τmp0 λ.Object
										τmp1 λ.Object
									)
									τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
										TwitchVodIE,
										TwitchCollectionIE,
										TwitchVideosIE,
										TwitchVideosClipsIE,
										TwitchVideosCollectionsIE,
										TwitchClipsIE,
									))
									for {
										if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
											break
										}
										ϒie = τmp1
										λgy.Yield(λ.Calm(ϒie, "suitable", ϒurl))
									}
									return λ.None
								})
							})))) {
							return λ.False
						} else {
							return λ.Calm(λ.Cal(λ.SuperType, TwitchStreamIE, ϒcls), "suitable", ϒurl)
						}
					}()
				})
			TwitchStreamIE_suitable = λ.Cal(λ.ClassMethodType, TwitchStreamIE_suitable)
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL": TwitchStreamIE__VALID_URL,
				"suitable":   TwitchStreamIE_suitable,
			})
		}())
		TwitchClipsIE = λ.Cal(λ.TypeType, λ.StrLiteral("TwitchClipsIE"), λ.NewTuple(TwitchBaseIE), func() λ.Dict {
			var (
				TwitchClipsIE_IE_NAME       λ.Object
				TwitchClipsIE__VALID_URL    λ.Object
				TwitchClipsIE__real_extract λ.Object
			)
			TwitchClipsIE_IE_NAME = λ.StrLiteral("twitch:clips")
			TwitchClipsIE__VALID_URL = λ.StrLiteral("(?x)\n                    https?://\n                        (?:\n                            clips\\.twitch\\.tv/(?:embed\\?.*?\\bclip=|(?:[^/]+/)*)|\n                            (?:(?:www|go|m)\\.)?twitch\\.tv/[^/]+/clip/\n                        )\n                        (?P<id>[^/?#&]+)\n                    ")
			TwitchClipsIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒclip          λ.Object
						ϒformats       λ.Object
						ϒmobj          λ.Object
						ϒoption        λ.Object
						ϒself          = λargs[0]
						ϒsource        λ.Object
						ϒthumb         λ.Object
						ϒthumbnail_id  λ.Object
						ϒthumbnail_url λ.Object
						ϒthumbnails    λ.Object
						ϒurl           = λargs[1]
						ϒvideo_id      λ.Object
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒclip = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.StrLiteral("https://gql.twitch.tv/gql"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "data", Value: λ.Calm(λ.Cal(Ωjson.ϒdumps, λ.DictLiteral(map[string]λ.Object{
							"query": λ.Mod(λ.StrLiteral("{\n  clip(slug: \"%s\") {\n    broadcaster {\n      displayName\n    }\n    createdAt\n    curator {\n      displayName\n      id\n    }\n    durationSeconds\n    id\n    tiny: thumbnailURL(width: 86, height: 45)\n    small: thumbnailURL(width: 260, height: 147)\n    medium: thumbnailURL(width: 480, height: 272)\n    title\n    videoQualities {\n      frameRate\n      quality\n      sourceURL\n    }\n    viewCount\n  }\n}"), ϒvideo_id),
						})), "encode")},
						{Name: "headers", Value: λ.DictLiteral(map[string]λ.Object{
							"Client-ID": λ.GetAttr(ϒself, "_CLIENT_ID", nil),
						})},
					}), λ.StrLiteral("data")), λ.StrLiteral("clip"))
					if !λ.IsTrue(ϒclip) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This clip is no longer available")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Calm(ϒclip, "get", λ.StrLiteral("videoQualities"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒoption = τmp1
						if !λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒoption, λ.DictType)) {
							continue
						}
						ϒsource = λ.Cal(ϒurl_or_none, λ.Calm(ϒoption, "get", λ.StrLiteral("sourceURL")))
						if !λ.IsTrue(ϒsource) {
							continue
						}
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"url":       ϒsource,
							"format_id": λ.Calm(ϒoption, "get", λ.StrLiteral("quality")),
							"height":    λ.Cal(ϒint_or_none, λ.Calm(ϒoption, "get", λ.StrLiteral("quality"))),
							"fps":       λ.Cal(ϒint_or_none, λ.Calm(ϒoption, "get", λ.StrLiteral("frameRate"))),
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
						λ.StrLiteral("tiny"),
						λ.StrLiteral("small"),
						λ.StrLiteral("medium"),
					))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒthumbnail_id = τmp1
						ϒthumbnail_url = λ.Calm(ϒclip, "get", ϒthumbnail_id)
						if !λ.IsTrue(ϒthumbnail_url) {
							continue
						}
						ϒthumb = λ.DictLiteral(map[string]λ.Object{
							"id":  ϒthumbnail_id,
							"url": ϒthumbnail_url,
						})
						ϒmobj = λ.Cal(Ωre.ϒsearch, λ.StrLiteral("-(\\d+)x(\\d+)\\."), ϒthumbnail_url)
						if λ.IsTrue(ϒmobj) {
							λ.Calm(ϒthumb, "update", λ.DictLiteral(map[string]λ.Object{
								"height": λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(2))),
								"width":  λ.Cal(λ.IntType, λ.Calm(ϒmobj, "group", λ.IntLiteral(1))),
							}))
						}
						λ.Calm(ϒthumbnails, "append", ϒthumb)
					}
					return λ.DictLiteral(map[string]λ.Object{
						"id": func() λ.Object {
							if λv := λ.Calm(ϒclip, "get", λ.StrLiteral("id")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"title": func() λ.Object {
							if λv := λ.Calm(ϒclip, "get", λ.StrLiteral("title")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒvideo_id
							}
						}(),
						"formats":    ϒformats,
						"duration":   λ.Cal(ϒint_or_none, λ.Calm(ϒclip, "get", λ.StrLiteral("durationSeconds"))),
						"views":      λ.Cal(ϒint_or_none, λ.Calm(ϒclip, "get", λ.StrLiteral("viewCount"))),
						"timestamp":  λ.Cal(ϒunified_timestamp, λ.Calm(ϒclip, "get", λ.StrLiteral("createdAt"))),
						"thumbnails": ϒthumbnails,
						"creator": λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("broadcaster")), λ.StrLiteral("displayName"))
							}), ϒcompat_str),
						"uploader": λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("curator")), λ.StrLiteral("displayName"))
							}), ϒcompat_str),
						"uploader_id": λ.Cal(ϒtry_get, ϒclip, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.StrLiteral("curator")), λ.StrLiteral("id"))
							}), ϒcompat_str),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":       TwitchClipsIE_IE_NAME,
				"_VALID_URL":    TwitchClipsIE__VALID_URL,
				"_real_extract": TwitchClipsIE__real_extract,
			})
		}())
	})
}
