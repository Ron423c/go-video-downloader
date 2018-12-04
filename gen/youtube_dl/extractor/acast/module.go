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
 * acast/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/acast.py
 */

package acast

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ACastChannelIE     λ.Object
	ACastIE            λ.Object
	InfoExtractor      λ.Object
	ϒcompat_str        λ.Object
	ϒfloat_or_none     λ.Object
	ϒint_or_none       λ.Object
	ϒtry_get           λ.Object
	ϒunified_timestamp λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒtry_get = Ωutils.ϒtry_get
		ϒunified_timestamp = Ωutils.ϒunified_timestamp
		ACastIE = λ.Cal(λ.TypeType, λ.NewStr("ACastIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ACastIE_IE_NAME       λ.Object
				ACastIE__TESTS        λ.Object
				ACastIE__VALID_URL    λ.Object
				ACastIE__real_extract λ.Object
			)
			ACastIE_IE_NAME = λ.NewStr("acast")
			ACastIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?acast\\.com/(?P<channel>[^/]+)/(?P<id>[^/#?]+)")
			ACastIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.acast.com/condenasttraveler/-where-are-you-taipei-101-taiwan"),
					λ.NewStr("md5"): λ.NewStr("ada3de5a1e3a2a381327d749854788bb"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("57de3baa-4bb0-487e-9418-2692c1277a34"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("\"Where Are You?\": Taipei 101, Taiwan"),
						λ.NewStr("description"): λ.NewStr("md5:a0b4ef3634e63866b542e5b1199a1a0e"),
						λ.NewStr("timestamp"):   λ.NewInt(1196172000),
						λ.NewStr("upload_date"): λ.NewStr("20071127"),
						λ.NewStr("duration"):    λ.NewInt(211),
						λ.NewStr("creator"):     λ.NewStr("Concierge"),
						λ.NewStr("series"):      λ.NewStr("Condé Nast Traveler Podcast"),
						λ.NewStr("episode"):     λ.NewStr("\"Where Are You?\": Taipei 101, Taiwan"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.acast.com/sparpodcast/2.raggarmordet-rosterurdetforflutna"),
					λ.NewStr("md5"): λ.NewStr("a02393c74f3bdb1801c3ec2695577ce0"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("2a92b283-1a75-4ad8-8396-499c641de0d9"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("2. Raggarmordet - Röster ur det förflutna"),
						λ.NewStr("description"): λ.NewStr("md5:4f81f6d8cf2e12ee21a321d8bca32db4"),
						λ.NewStr("timestamp"):   λ.NewInt(1477346700),
						λ.NewStr("upload_date"): λ.NewStr("20161024"),
						λ.NewStr("duration"):    λ.NewFloat(2766.602563),
						λ.NewStr("creator"):     λ.NewStr("Anton Berg & Martin Johnson"),
						λ.NewStr("series"):      λ.NewStr("Spår"),
						λ.NewStr("episode"):     λ.NewStr("2. Raggarmordet - Röster ur det förflutna"),
					}),
				}),
			)
			ACastIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcast_data  λ.Object
						ϒchannel    λ.Object
						ϒdisplay_id λ.Object
						ϒe          λ.Object
						ϒmedia_url  λ.Object
						ϒs          λ.Object
						ϒself       = λargs[0]
						ϒtitle      λ.Object
						ϒurl        = λargs[1]
						τmp0        λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒchannel = λ.GetItem(τmp0, λ.NewInt(0))
					ϒdisplay_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒs = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://play-api.acast.com/stitch/%s/%s"), λ.NewTuple(
						ϒchannel,
						ϒdisplay_id,
					)), ϒdisplay_id), λ.NewStr("result"))
					ϒmedia_url = λ.GetItem(ϒs, λ.NewStr("url"))
					ϒcast_data = λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("https://play-api.acast.com/splash/%s/%s"), λ.NewTuple(
						ϒchannel,
						ϒdisplay_id,
					)), ϒdisplay_id), λ.NewStr("result"))
					ϒe = λ.GetItem(ϒcast_data, λ.NewStr("episode"))
					ϒtitle = λ.GetItem(ϒe, λ.NewStr("name"))
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.Cal(ϒcompat_str, λ.GetItem(ϒe, λ.NewStr("id"))),
						λ.NewStr("display_id"): ϒdisplay_id,
						λ.NewStr("url"):        ϒmedia_url,
						λ.NewStr("title"):      ϒtitle,
						λ.NewStr("description"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("description")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("summary"))
							}
						}(),
						λ.NewStr("thumbnail"): λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("image")),
						λ.NewStr("timestamp"): λ.Cal(ϒunified_timestamp, λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("publishingDate"))),
						λ.NewStr("duration"): λ.Cal(ϒfloat_or_none, func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒs, "get", nil), λ.NewStr("duration")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("duration"))
							}
						}()),
						λ.NewStr("filesize"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("contentLength"))),
						λ.NewStr("creator"): λ.Cal(ϒtry_get, ϒcast_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("show")), λ.NewStr("author"))
							}), ϒcompat_str),
						λ.NewStr("series"): λ.Cal(ϒtry_get, ϒcast_data, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("show")), λ.NewStr("name"))
							}), ϒcompat_str),
						λ.NewStr("season_number"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("seasonNumber"))),
						λ.NewStr("episode"):        ϒtitle,
						λ.NewStr("episode_number"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒe, "get", nil), λ.NewStr("episodeNumber"))),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       ACastIE_IE_NAME,
				λ.NewStr("_TESTS"):        ACastIE__TESTS,
				λ.NewStr("_VALID_URL"):    ACastIE__VALID_URL,
				λ.NewStr("_real_extract"): ACastIE__real_extract,
			})
		}())
		ACastChannelIE = λ.Cal(λ.TypeType, λ.NewStr("ACastChannelIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				ACastChannelIE__VALID_URL λ.Object
				ACastChannelIE_suitable   λ.Object
			)
			ACastChannelIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?acast\\.com/(?P<id>[^/#?]+)")
			ACastChannelIE_suitable = λ.NewFunction("suitable",
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
						if λ.IsTrue(λ.Cal(λ.GetAttr(ACastIE, "suitable", nil), ϒurl)) {
							return λ.False
						} else {
							return λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, ACastChannelIE, ϒcls), "suitable", nil), ϒurl)
						}
					}()
				})
			ACastChannelIE_suitable = λ.Cal(λ.ClassMethodType, ACastChannelIE_suitable)
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): ACastChannelIE__VALID_URL,
				λ.NewStr("suitable"):   ACastChannelIE_suitable,
			})
		}())
	})
}
