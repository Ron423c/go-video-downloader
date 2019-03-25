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
 * swrmediathek/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/swrmediathek.py
 */

package swrmediathek

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor       λ.Object
	SWRMediathekIE      λ.Object
	ϒdetermine_protocol λ.Object
	ϒint_or_none        λ.Object
	ϒparse_duration     λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒdetermine_protocol = Ωutils.ϒdetermine_protocol
		SWRMediathekIE = λ.Cal(λ.TypeType, λ.NewStr("SWRMediathekIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SWRMediathekIE__TESTS        λ.Object
				SWRMediathekIE__VALID_URL    λ.Object
				SWRMediathekIE__real_extract λ.Object
			)
			SWRMediathekIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?swrmediathek\\.de/(?:content/)?player\\.htm\\?show=(?P<id>[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12})")
			SWRMediathekIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://swrmediathek.de/player.htm?show=849790d0-dab8-11e3-a953-0026b975f2e6"),
					λ.NewStr("md5"): λ.NewStr("8c5f6f0172753368547ca8413a7768ac"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("849790d0-dab8-11e3-a953-0026b975f2e6"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("SWR odysso"),
						λ.NewStr("description"): λ.NewStr("md5:2012e31baad36162e97ce9eb3f157b8a"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^http:.*\\.jpg$"),
						λ.NewStr("duration"):    λ.NewInt(2602),
						λ.NewStr("upload_date"): λ.NewStr("20140515"),
						λ.NewStr("uploader"):    λ.NewStr("SWR Fernsehen"),
						λ.NewStr("uploader_id"): λ.NewStr("990030"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://swrmediathek.de/player.htm?show=0e1a8510-ddf2-11e3-9be3-0026b975f2e6"),
					λ.NewStr("md5"): λ.NewStr("b10ab854f912eecc5a6b55cd6fc1f545"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("0e1a8510-ddf2-11e3-9be3-0026b975f2e6"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Nachtcafé - Alltagsdroge Alkohol - zwischen Sektempfang und Komasaufen"),
						λ.NewStr("description"): λ.NewStr("md5:e0a3adc17e47db2c23aab9ebc36dbee2"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
						λ.NewStr("duration"):    λ.NewInt(5305),
						λ.NewStr("upload_date"): λ.NewStr("20140516"),
						λ.NewStr("uploader"):    λ.NewStr("SWR Fernsehen"),
						λ.NewStr("uploader_id"): λ.NewStr("990030"),
					}),
					λ.NewStr("skip"): λ.NewStr("redirect to http://swrmediathek.de/index.htm?hinweis=swrlink"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://swrmediathek.de/player.htm?show=bba23e10-cb93-11e3-bf7f-0026b975f2e6"),
					λ.NewStr("md5"): λ.NewStr("4382e4ef2c9d7ce6852535fa867a0dd3"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("bba23e10-cb93-11e3-bf7f-0026b975f2e6"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Saša Stanišic: Vor dem Fest"),
						λ.NewStr("description"): λ.NewStr("md5:5b792387dc3fbb171eb709060654e8c9"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:http://.*\\.jpg"),
						λ.NewStr("duration"):    λ.NewInt(3366),
						λ.NewStr("upload_date"): λ.NewStr("20140520"),
						λ.NewStr("uploader"):    λ.NewStr("SWR 2"),
						λ.NewStr("uploader_id"): λ.NewStr("284670"),
					}),
					λ.NewStr("skip"): λ.NewStr("redirect to http://swrmediathek.de/index.htm?hinweis=swrlink"),
				}),
			)
			SWRMediathekIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒattr         λ.Object
						ϒcodec        λ.Object
						ϒentry        λ.Object
						ϒentry_attr   λ.Object
						ϒentry_pdatet λ.Object
						ϒf_url        λ.Object
						ϒformats      λ.Object
						ϒmedia_type   λ.Object
						ϒself         = λargs[0]
						ϒtitle        λ.Object
						ϒupload_date  λ.Object
						ϒurl          = λargs[1]
						ϒvideo        λ.Object
						ϒvideo_id     λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒvideo = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://swrmediathek.de/AjaxEntry?ekey=%s"), ϒvideo_id), ϒvideo_id, λ.NewStr("Downloading video JSON"))
					ϒattr = λ.GetItem(ϒvideo, λ.NewStr("attr"))
					ϒtitle = λ.GetItem(ϒattr, λ.NewStr("entry_title"))
					ϒmedia_type = λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("entry_etype"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo, "get", nil), λ.NewStr("sub"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒentry = τmp1
						if λ.IsTrue(λ.Ne(λ.Cal(λ.GetAttr(ϒentry, "get", nil), λ.NewStr("name")), λ.NewStr("entry_media"))) {
							continue
						}
						ϒentry_attr = λ.Cal(λ.GetAttr(ϒentry, "get", nil), λ.NewStr("attr"), λ.NewDictWithTable(map[λ.Object]λ.Object{}))
						ϒf_url = λ.Cal(λ.GetAttr(ϒentry_attr, "get", nil), λ.NewStr("val2"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒf_url))) {
							continue
						}
						ϒcodec = λ.Cal(λ.GetAttr(ϒentry_attr, "get", nil), λ.NewStr("val0"))
						if λ.IsTrue(λ.Eq(ϒcodec, λ.NewStr("m3u8"))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
								ϒf_url,
								ϒvideo_id,
								λ.NewStr("mp4"),
								λ.NewStr("m3u8_native"),
							), λ.KWArgs{
								{Name: "m3u8_id", Value: λ.NewStr("hls")},
								{Name: "fatal", Value: λ.False},
							}))
						} else {
							if λ.IsTrue(λ.Eq(ϒcodec, λ.NewStr("f4m"))) {
								λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
									λ.Add(ϒf_url, λ.NewStr("?hdcore=3.7.0")),
									ϒvideo_id,
								), λ.KWArgs{
									{Name: "f4m_id", Value: λ.NewStr("hds")},
									{Name: "fatal", Value: λ.False},
								}))
							} else {
								λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("format_id"): λ.Cal(ϒdetermine_protocol, λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"): ϒf_url,
									})),
									λ.NewStr("url"):     ϒf_url,
									λ.NewStr("quality"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒentry_attr, "get", nil), λ.NewStr("val1"))),
									λ.NewStr("vcodec"): func() λ.Object {
										if λ.IsTrue(λ.Eq(ϒmedia_type, λ.NewStr("Video"))) {
											return ϒcodec
										} else {
											return λ.NewStr("none")
										}
									}(),
									λ.NewStr("acodec"): func() λ.Object {
										if λ.IsTrue(λ.Eq(ϒmedia_type, λ.NewStr("Audio"))) {
											return ϒcodec
										} else {
											return λ.None
										}
									}(),
								}))
							}
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒupload_date = λ.None
					ϒentry_pdatet = λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("entry_pdatet"))
					if λ.IsTrue(ϒentry_pdatet) {
						ϒupload_date = λ.GetItem(ϒentry_pdatet, λ.NewSlice(λ.None, λ.Neg(λ.NewInt(4)), λ.None))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("entry_descl")),
						λ.NewStr("thumbnail"):   λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("entry_image_16_9")),
						λ.NewStr("duration"):    λ.Cal(ϒparse_duration, λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("entry_durat"))),
						λ.NewStr("upload_date"): ϒupload_date,
						λ.NewStr("uploader"):    λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("channel_title")),
						λ.NewStr("uploader_id"): λ.Cal(λ.GetAttr(ϒattr, "get", nil), λ.NewStr("channel_idkey")),
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        SWRMediathekIE__TESTS,
				λ.NewStr("_VALID_URL"):    SWRMediathekIE__VALID_URL,
				λ.NewStr("_real_extract"): SWRMediathekIE__real_extract,
			})
		}())
	})
}
