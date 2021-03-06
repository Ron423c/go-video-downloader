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
 * sverigesradio/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/sverigesradio.py
 */

package sverigesradio

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor              λ.Object
	SverigesRadioBaseIE        λ.Object
	SverigesRadioEpisodeIE     λ.Object
	SverigesRadioPublicationIE λ.Object
	ϒdetermine_ext             λ.Object
	ϒint_or_none               λ.Object
	ϒstr_or_none               λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒstr_or_none = Ωutils.ϒstr_or_none
		SverigesRadioBaseIE = λ.Cal(λ.TypeType, λ.StrLiteral("SverigesRadioBaseIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				SverigesRadioBaseIE__BASE_URL                 λ.Object
				SverigesRadioBaseIE__CODING_FORMAT_TO_ABR_MAP λ.Object
				SverigesRadioBaseIE__EXT_TO_CODEC_MAP         λ.Object
				SverigesRadioBaseIE__QUALITIES                λ.Object
				SverigesRadioBaseIE__real_extract             λ.Object
			)
			SverigesRadioBaseIE__BASE_URL = λ.StrLiteral("https://sverigesradio.se/sida/playerajax/")
			SverigesRadioBaseIE__QUALITIES = λ.NewList(
				λ.StrLiteral("low"),
				λ.StrLiteral("medium"),
				λ.StrLiteral("high"),
			)
			SverigesRadioBaseIE__EXT_TO_CODEC_MAP = λ.DictLiteral(map[string]string{
				"mp3": "mp3",
				"m4a": "aac",
			})
			SverigesRadioBaseIE__CODING_FORMAT_TO_ABR_MAP = λ.DictLiteral(map[int]int{
				5:  128,
				11: 192,
				12: 32,
				13: 96,
			})
			SverigesRadioBaseIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒabr            λ.Object
						ϒaudio_id       λ.Object
						ϒaudio_url      λ.Object
						ϒaudio_url_data λ.Object
						ϒcoding_format  λ.Object
						ϒext            λ.Object
						ϒformats        λ.Object
						ϒitem           λ.Object
						ϒquality        λ.Object
						ϒquery          λ.Object
						ϒself           = λargs[0]
						ϒtitle          λ.Object
						ϒurl            = λargs[1]
						ϒurls           λ.Object
						τmp0            λ.Object
						τmp1            λ.Object
					)
					ϒaudio_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒquery = λ.DictLiteral(map[string]λ.Object{
						"id":   ϒaudio_id,
						"type": λ.GetAttr(ϒself, "_AUDIO_TYPE", nil),
					})
					ϒitem = λ.GetItem(λ.GetItem(λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Add(λ.GetAttr(ϒself, "_BASE_URL", nil), λ.StrLiteral("audiometadata")),
						ϒaudio_id,
						λ.StrLiteral("Downloading audio JSON metadata"),
					), λ.KWArgs{
						{Name: "query", Value: ϒquery},
					}), λ.StrLiteral("items")), λ.IntLiteral(0))
					ϒtitle = λ.GetItem(ϒitem, λ.StrLiteral("subtitle"))
					λ.SetItem(ϒquery, λ.StrLiteral("format"), λ.StrLiteral("iis"))
					ϒurls = λ.NewList()
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetAttr(ϒself, "_QUALITIES", nil))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒquality = τmp1
						λ.SetItem(ϒquery, λ.StrLiteral("quality"), ϒquality)
						ϒaudio_url_data = func() λ.Object {
							if λv := λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
								λ.Add(λ.GetAttr(ϒself, "_BASE_URL", nil), λ.StrLiteral("getaudiourl")),
								ϒaudio_id,
								λ.Mod(λ.StrLiteral("Downloading %s format JSON metadata"), ϒquality),
							), λ.KWArgs{
								{Name: "fatal", Value: λ.False},
								{Name: "query", Value: ϒquery},
							}); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.DictLiteral(map[λ.Object]λ.Object{})
							}
						}()
						ϒaudio_url = λ.Calm(ϒaudio_url_data, "get", λ.StrLiteral("audioUrl"))
						if λ.IsTrue(func() λ.Object {
							if λv := λ.NewBool(!λ.IsTrue(ϒaudio_url)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.NewBool(λ.Contains(ϒurls, ϒaudio_url))
							}
						}()) {
							continue
						}
						λ.Calm(ϒurls, "append", ϒaudio_url)
						ϒext = λ.Cal(ϒdetermine_ext, ϒaudio_url)
						ϒcoding_format = λ.Calm(ϒaudio_url_data, "get", λ.StrLiteral("codingFormat"))
						ϒabr = func() λ.Object {
							if λv := λ.Cal(ϒint_or_none, λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.StrLiteral("_a(\\d+)\\.m4a"),
								ϒaudio_url,
								λ.StrLiteral("audio bitrate"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Calm(λ.GetAttr(ϒself, "_CODING_FORMAT_TO_ABR_MAP", nil), "get", ϒcoding_format)
							}
						}()
						λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
							"abr":       ϒabr,
							"acodec":    λ.Calm(λ.GetAttr(ϒself, "_EXT_TO_CODEC_MAP", nil), "get", ϒext),
							"ext":       ϒext,
							"format_id": λ.Cal(ϒstr_or_none, ϒcoding_format),
							"vcodec":    λ.StrLiteral("none"),
							"url":       ϒaudio_url,
						}))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id":          ϒaudio_id,
						"title":       ϒtitle,
						"formats":     ϒformats,
						"series":      λ.Calm(ϒitem, "get", λ.StrLiteral("title")),
						"duration":    λ.Cal(ϒint_or_none, λ.Calm(ϒitem, "get", λ.StrLiteral("duration"))),
						"thumbnail":   λ.Calm(ϒitem, "get", λ.StrLiteral("displayimageurl")),
						"description": λ.Calm(ϒitem, "get", λ.StrLiteral("description")),
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_BASE_URL":                 SverigesRadioBaseIE__BASE_URL,
				"_CODING_FORMAT_TO_ABR_MAP": SverigesRadioBaseIE__CODING_FORMAT_TO_ABR_MAP,
				"_EXT_TO_CODEC_MAP":         SverigesRadioBaseIE__EXT_TO_CODEC_MAP,
				"_QUALITIES":                SverigesRadioBaseIE__QUALITIES,
				"_real_extract":             SverigesRadioBaseIE__real_extract,
			})
		}())
		SverigesRadioPublicationIE = λ.Cal(λ.TypeType, λ.StrLiteral("SverigesRadioPublicationIE"), λ.NewTuple(SverigesRadioBaseIE), func() λ.Dict {
			var (
				SverigesRadioPublicationIE_IE_NAME     λ.Object
				SverigesRadioPublicationIE__AUDIO_TYPE λ.Object
				SverigesRadioPublicationIE__VALID_URL  λ.Object
			)
			SverigesRadioPublicationIE_IE_NAME = λ.StrLiteral("sverigesradio:publication")
			SverigesRadioPublicationIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?sverigesradio\\.se/sida/(?:artikel|gruppsida)\\.aspx\\?.*?\\bartikel=(?P<id>[0-9]+)")
			SverigesRadioPublicationIE__AUDIO_TYPE = λ.StrLiteral("publication")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":     SverigesRadioPublicationIE_IE_NAME,
				"_AUDIO_TYPE": SverigesRadioPublicationIE__AUDIO_TYPE,
				"_VALID_URL":  SverigesRadioPublicationIE__VALID_URL,
			})
		}())
		SverigesRadioEpisodeIE = λ.Cal(λ.TypeType, λ.StrLiteral("SverigesRadioEpisodeIE"), λ.NewTuple(SverigesRadioBaseIE), func() λ.Dict {
			var (
				SverigesRadioEpisodeIE_IE_NAME     λ.Object
				SverigesRadioEpisodeIE__AUDIO_TYPE λ.Object
				SverigesRadioEpisodeIE__VALID_URL  λ.Object
			)
			SverigesRadioEpisodeIE_IE_NAME = λ.StrLiteral("sverigesradio:episode")
			SverigesRadioEpisodeIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?sverigesradio\\.se/(?:sida/)?avsnitt/(?P<id>[0-9]+)")
			SverigesRadioEpisodeIE__AUDIO_TYPE = λ.StrLiteral("episode")
			return λ.ClassDictLiteral(map[string]λ.Object{
				"IE_NAME":     SverigesRadioEpisodeIE_IE_NAME,
				"_AUDIO_TYPE": SverigesRadioEpisodeIE__AUDIO_TYPE,
				"_VALID_URL":  SverigesRadioEpisodeIE__VALID_URL,
			})
		}())
	})
}
