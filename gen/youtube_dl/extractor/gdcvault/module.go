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
 * gdcvault/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/gdcvault.py
 */

package gdcvault

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	GDCVaultIE          λ.Object
	HEADRequest         λ.Object
	InfoExtractor       λ.Object
	ϒsanitized_Request  λ.Object
	ϒurlencode_postdata λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		HEADRequest = Ωutils.HEADRequest
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒurlencode_postdata = Ωutils.ϒurlencode_postdata
		GDCVaultIE = λ.Cal(λ.TypeType, λ.NewStr("GDCVaultIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GDCVaultIE__TESTS        λ.Object
				GDCVaultIE__VALID_URL    λ.Object
				GDCVaultIE__real_extract λ.Object
			)
			GDCVaultIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?gdcvault\\.com/play/(?P<id>\\d+)/(?P<name>(\\w|-)+)?")
			GDCVaultIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gdcvault.com/play/1019721/Doki-Doki-Universe-Sweet-Simple"),
					λ.NewStr("md5"): λ.NewStr("7ce8388f544c88b7ac11c7ab1b593704"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("1019721"),
						λ.NewStr("display_id"): λ.NewStr("Doki-Doki-Universe-Sweet-Simple"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("title"):      λ.NewStr("Doki-Doki Universe: Sweet, Simple and Genuine (GDC Next 10)"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gdcvault.com/play/1015683/Embracing-the-Dark-Art-of"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("1015683"),
						λ.NewStr("display_id"): λ.NewStr("Embracing-the-Dark-Art-of"),
						λ.NewStr("ext"):        λ.NewStr("flv"),
						λ.NewStr("title"):      λ.NewStr("Embracing the Dark Art of Mathematical Modeling in AI"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gdcvault.com/play/1015301/Thexder-Meets-Windows-95-or"),
					λ.NewStr("md5"): λ.NewStr("a5eb77996ef82118afbbe8e48731b98e"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("1015301"),
						λ.NewStr("display_id"): λ.NewStr("Thexder-Meets-Windows-95-or"),
						λ.NewStr("ext"):        λ.NewStr("flv"),
						λ.NewStr("title"):      λ.NewStr("Thexder Meets Windows 95, or Writing Great Games in the Windows 95 Environment"),
					}),
					λ.NewStr("skip"): λ.NewStr("Requires login"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://gdcvault.com/play/1020791/"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://gdcvault.com/play/1023460/Tenacious-Design-and-The-Interface"),
					λ.NewStr("md5"): λ.NewStr("a8efb6c31ed06ca8739294960b2dbabd"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("1023460"),
						λ.NewStr("ext"):        λ.NewStr("mp4"),
						λ.NewStr("display_id"): λ.NewStr("Tenacious-Design-and-The-Interface"),
						λ.NewStr("title"):      λ.NewStr("Tenacious Design and The Interface of 'Destiny'"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gdcvault.com/play/1014631/Classic-Game-Postmortem-PAC"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):    λ.NewStr("1014631"),
						λ.NewStr("ext"):   λ.NewStr("flv"),
						λ.NewStr("title"): λ.NewStr("How to Create a Good Game - From My Experience of Designing Pac-Man"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
						λ.NewStr("format"):        λ.NewStr("jp"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.gdcvault.com/play/1435/An-American-engine-in-Tokyo"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         λ.NewStr("1435"),
						λ.NewStr("display_id"): λ.NewStr("An-American-engine-in-Tokyo"),
						λ.NewStr("ext"):        λ.NewStr("flv"),
						λ.NewStr("title"):      λ.NewStr("An American Engine in Tokyo:/nThe collaboration of Epic Games and Square Enix/nFor THE LAST REMINANT"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
			)
			GDCVaultIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						PLAYER_REGEX λ.Object
						ϒdirect_url  λ.Object
						ϒdisplay_id  λ.Object
						ϒhead        λ.Object
						ϒlogin_res   λ.Object
						ϒmobj        λ.Object
						ϒself        = λargs[0]
						ϒstart_page  λ.Object
						ϒtitle       λ.Object
						ϒurl         = λargs[1]
						ϒvideo_id    λ.Object
						ϒvideo_url   λ.Object
						ϒwebpage_url λ.Object
						ϒxml_name    λ.Object
						ϒxml_root    λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒdisplay_id = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("name")); λ.IsTrue(λv) {
							return λv
						} else {
							return ϒvideo_id
						}
					}()
					ϒwebpage_url = λ.Add(λ.NewStr("http://www.gdcvault.com/play/"), ϒvideo_id)
					ϒstart_page = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒwebpage_url, ϒdisplay_id)
					ϒdirect_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
						λ.NewStr("s1\\.addVariable\\(\"file\",\\s*encodeURIComponent\\(\"(/[^\"]+)\"\\)\\);"),
						ϒstart_page,
						λ.NewStr("url"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(ϒdirect_url) {
						ϒtitle = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<td><strong>Session Name</strong></td>\\s*<td>(.*?)</td>"), ϒstart_page, λ.NewStr("title"))
						ϒvideo_url = λ.Add(λ.NewStr("http://www.gdcvault.com"), ϒdirect_url)
						ϒhead = λ.Cal(λ.GetAttr(ϒself, "_request_webpage", nil), λ.Cal(HEADRequest, ϒvideo_url), ϒvideo_id)
						ϒvideo_url = λ.Cal(λ.GetAttr(ϒhead, "geturl", nil))
						return λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("id"):         ϒvideo_id,
							λ.NewStr("display_id"): ϒdisplay_id,
							λ.NewStr("url"):        ϒvideo_url,
							λ.NewStr("title"):      ϒtitle,
						})
					}
					PLAYER_REGEX = λ.NewStr("<iframe src=\"(?P<xml_root>.+?)/(?:gdc-)?player.*?\\.html.*?\".*?</iframe>")
					ϒxml_root = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						PLAYER_REGEX,
						ϒstart_page,
						λ.NewStr("xml root"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒxml_root == λ.None)) {
						ϒlogin_res = λ.Cal(λ.GetAttr(ϒself, "_login", nil), ϒwebpage_url, ϒdisplay_id)
						if λ.IsTrue(λ.NewBool(ϒlogin_res == λ.None)) {
							λ.Cal(λ.GetAttr(ϒself, "report_warning", nil), λ.NewStr("Could not login."))
						} else {
							ϒstart_page = ϒlogin_res
							ϒxml_root = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), PLAYER_REGEX, ϒstart_page, λ.NewStr("xml root"))
						}
					}
					ϒxml_name = λ.Call(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewArgs(
						λ.NewStr("<iframe src=\".*?\\?xml=(.+?\\.xml).*?\".*?</iframe>"),
						ϒstart_page,
						λ.NewStr("xml filename"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(λ.NewBool(ϒxml_name == λ.None)) {
						ϒxml_name = λ.Cal(λ.GetAttr(ϒself, "_html_search_regex", nil), λ.NewStr("<iframe src=\".*?\\?xmlURL=xml/(?P<xml_file>.+?\\.xml).*?\".*?</iframe>"), ϒstart_page, λ.NewStr("xml filename"))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("_type"):      λ.NewStr("url_transparent"),
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("display_id"): ϒdisplay_id,
						λ.NewStr("url"): λ.Mod(λ.NewStr("%s/xml/%s"), λ.NewTuple(
							ϒxml_root,
							ϒxml_name,
						)),
						λ.NewStr("ie_key"): λ.NewStr("DigitallySpeaking"),
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        GDCVaultIE__TESTS,
				λ.NewStr("_VALID_URL"):    GDCVaultIE__VALID_URL,
				λ.NewStr("_real_extract"): GDCVaultIE__real_extract,
			})
		}())
	})
}
