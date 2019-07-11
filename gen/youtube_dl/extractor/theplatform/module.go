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
 * theplatform/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/theplatform.py
 */

package theplatform

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωadobepass "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/adobepass"
	Ωonce "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/once"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	AdobePassIE                   λ.Object
	ExtractorError                λ.Object
	OnceIE                        λ.Object
	ThePlatformBaseIE             λ.Object
	ThePlatformFeedIE             λ.Object
	ThePlatformIE                 λ.Object
	ϒ_x                           λ.Object
	ϒcompat_parse_qs              λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdefault_ns                   λ.Object
	ϒdetermine_ext                λ.Object
	ϒfind_xpath_attr              λ.Object
	ϒfloat_or_none                λ.Object
	ϒint_or_none                  λ.Object
	ϒmimetype2ext                 λ.Object
	ϒsanitized_Request            λ.Object
	ϒunsmuggle_url                λ.Object
	ϒupdate_url_query             λ.Object
	ϒxpath_with_ns                λ.Object
)

func init() {
	λ.InitModule(func() {
		OnceIE = Ωonce.OnceIE
		AdobePassIE = Ωadobepass.AdobePassIE
		ϒcompat_parse_qs = Ωcompat.ϒcompat_parse_qs
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒfloat_or_none = Ωutils.ϒfloat_or_none
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒsanitized_Request = Ωutils.ϒsanitized_Request
		ϒunsmuggle_url = Ωutils.ϒunsmuggle_url
		ϒupdate_url_query = Ωutils.ϒupdate_url_query
		ϒxpath_with_ns = Ωutils.ϒxpath_with_ns
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒfind_xpath_attr = Ωutils.ϒfind_xpath_attr
		ϒdefault_ns = λ.NewStr("http://www.w3.org/2005/SMIL21/Language")
		ϒ_x = λ.NewFunction("<lambda>",
			[]λ.Param{
				{Name: "p"},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒp = λargs[0]
				)
				return λ.Cal(ϒxpath_with_ns, ϒp, λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("smil"): ϒdefault_ns,
				}))
			})
		ThePlatformBaseIE = λ.Cal(λ.TypeType, λ.NewStr("ThePlatformBaseIE"), λ.NewTuple(OnceIE), func() λ.Dict {
			var (
				ThePlatformBaseIE__TP_TLD                        λ.Object
				ThePlatformBaseIE__download_theplatform_metadata λ.Object
				ThePlatformBaseIE__extract_theplatform_metadata  λ.Object
				ThePlatformBaseIE__extract_theplatform_smil      λ.Object
				ThePlatformBaseIE__parse_theplatform_metadata    λ.Object
			)
			ThePlatformBaseIE__TP_TLD = λ.NewStr("com")
			ThePlatformBaseIE__extract_theplatform_smil = λ.NewFunction("_extract_theplatform_smil",
				[]λ.Param{
					{Name: "self"},
					{Name: "smil_url"},
					{Name: "video_id"},
					{Name: "note", Def: λ.NewStr("Downloading SMIL data")},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_format       λ.Object
						ϒerror_element λ.Object
						ϒexception     λ.Object
						ϒformats       λ.Object
						ϒhdnea2        λ.Object
						ϒmedia_url     λ.Object
						ϒmeta          λ.Object
						ϒnote          = λargs[3]
						ϒself          = λargs[0]
						ϒsmil_formats  λ.Object
						ϒsmil_url      = λargs[1]
						ϒsubtitles     λ.Object
						ϒvideo_id      = λargs[2]
						τmp0           λ.Object
						τmp1           λ.Object
					)
					ϒmeta = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						ϒsmil_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "note", Value: ϒnote},
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("format"): λ.NewStr("SMIL"),
						})},
						{Name: "headers", Value: λ.Cal(λ.GetAttr(ϒself, "geo_verification_headers", nil))},
					})
					ϒerror_element = λ.Cal(ϒfind_xpath_attr, ϒmeta, λ.Cal(ϒ_x, λ.NewStr(".//smil:ref")), λ.NewStr("src"))
					if λ.IsTrue(λ.NewBool(ϒerror_element != λ.None)) {
						ϒexception = λ.Cal(ϒfind_xpath_attr, ϒerror_element, λ.Cal(ϒ_x, λ.NewStr(".//smil:param")), λ.NewStr("name"), λ.NewStr("exception"))
						if λ.IsTrue(λ.NewBool(ϒexception != λ.None)) {
							if λ.IsTrue(λ.Eq(λ.Cal(λ.GetAttr(ϒexception, "get", nil), λ.NewStr("value")), λ.NewStr("GeoLocationBlocked"))) {
								λ.Cal(λ.GetAttr(ϒself, "raise_geo_restricted", nil), λ.GetItem(λ.GetAttr(ϒerror_element, "attrib", nil), λ.NewStr("abstract")))
							} else {
								if λ.IsTrue(λ.Cal(λ.GetAttr(λ.GetItem(λ.GetAttr(ϒerror_element, "attrib", nil), λ.NewStr("src")), "startswith", nil), λ.Mod(λ.NewStr("http://link.theplatform.%s/s/errorFiles/Unavailable."), λ.GetAttr(ϒself, "_TP_TLD", nil)))) {
									panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.GetItem(λ.GetAttr(ϒerror_element, "attrib", nil), λ.NewStr("abstract"))), λ.KWArgs{
										{Name: "expected", Value: λ.True},
									})))
								}
							}
						}
					}
					ϒsmil_formats = λ.Call(λ.GetAttr(ϒself, "_parse_smil_formats", nil), λ.NewArgs(
						ϒmeta,
						ϒsmil_url,
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "namespace", Value: ϒdefault_ns},
						{Name: "f4m_params", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("g"):      λ.NewStr("UXWGVKRWHFSP"),
							λ.NewStr("hdcore"): λ.NewStr("3.0.3"),
						})},
						{Name: "transform_rtmp_url", Value: λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "streamer"},
								{Name: "src"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒsrc      = λargs[1]
									ϒstreamer = λargs[0]
								)
								return λ.NewTuple(
									ϒstreamer,
									λ.Add(λ.NewStr("mp4:"), ϒsrc),
								)
							})},
					})
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, ϒsmil_formats)
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒ_format = τmp1
						if λ.IsTrue(λ.Cal(λ.GetAttr(OnceIE, "suitable", nil), λ.GetItem(ϒ_format, λ.NewStr("url")))) {
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), λ.Cal(λ.GetAttr(ϒself, "_extract_once_formats", nil), λ.GetItem(ϒ_format, λ.NewStr("url"))))
						} else {
							ϒmedia_url = λ.GetItem(ϒ_format, λ.NewStr("url"))
							if λ.IsTrue(λ.Eq(λ.Cal(ϒdetermine_ext, ϒmedia_url), λ.NewStr("m3u8"))) {
								ϒhdnea2 = λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒself, "_get_cookies", nil), ϒmedia_url), "get", nil), λ.NewStr("hdnea2"))
								if λ.IsTrue(ϒhdnea2) {
									λ.SetItem(ϒ_format, λ.NewStr("url"), λ.Cal(ϒupdate_url_query, ϒmedia_url, λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("hdnea3"): λ.GetAttr(ϒhdnea2, "value", nil),
									})))
								}
							}
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), ϒ_format)
						}
					}
					ϒsubtitles = λ.Cal(λ.GetAttr(ϒself, "_parse_smil_subtitles", nil), ϒmeta, ϒdefault_ns)
					return λ.NewTuple(
						ϒformats,
						ϒsubtitles,
					)
				})
			ThePlatformBaseIE__download_theplatform_metadata = λ.NewFunction("_download_theplatform_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo_url λ.Object
						ϒpath     = λargs[1]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					ϒinfo_url = λ.Mod(λ.NewStr("http://link.theplatform.%s/s/%s?format=preview"), λ.NewTuple(
						λ.GetAttr(ϒself, "_TP_TLD", nil),
						ϒpath,
					))
					return λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒinfo_url, ϒvideo_id)
				})
			ThePlatformBaseIE__parse_theplatform_metadata = λ.NewFunction("_parse_theplatform_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "info"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒ_add_chapter λ.Object
						ϒcaption      λ.Object
						ϒcaptions     λ.Object
						ϒchapter      λ.Object
						ϒchapters     λ.Object
						ϒduration     λ.Object
						ϒinfo         = λargs[1]
						ϒlang         λ.Object
						ϒmime         λ.Object
						ϒself         = λargs[0]
						ϒsrc          λ.Object
						ϒsubtitles    λ.Object
						ϒtp_chapters  λ.Object
						τmp0          λ.Object
						τmp1          λ.Object
						τmp2          λ.Object
					)
					_ = ϒself
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒcaptions = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("captions"))
					if λ.IsTrue(λ.Cal(λ.BuiltinIsInstance, ϒcaptions, λ.ListType)) {
						τmp0 = λ.Cal(λ.BuiltinIter, ϒcaptions)
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒcaption = τmp1
							τmp2 = λ.NewTuple(
								λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("lang"), λ.NewStr("en")),
								λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("src")),
								λ.Cal(λ.GetAttr(ϒcaption, "get", nil), λ.NewStr("type")),
							)
							ϒlang = λ.GetItem(τmp2, λ.NewInt(0))
							ϒsrc = λ.GetItem(τmp2, λ.NewInt(1))
							ϒmime = λ.GetItem(τmp2, λ.NewInt(2))
							λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒsubtitles, "setdefault", nil), ϒlang, λ.NewList()), "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("ext"): λ.Cal(ϒmimetype2ext, ϒmime),
								λ.NewStr("url"): ϒsrc,
							}))
						}
					}
					ϒduration = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("duration"))
					ϒtp_chapters = λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("chapters"), λ.NewList())
					ϒchapters = λ.NewList()
					if λ.IsTrue(ϒtp_chapters) {
						ϒ_add_chapter = λ.NewFunction("_add_chapter",
							[]λ.Param{
								{Name: "start_time"},
								{Name: "end_time"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒend_time   = λargs[1]
									ϒstart_time = λargs[0]
								)
								ϒstart_time = λ.Cal(ϒfloat_or_none, ϒstart_time, λ.NewInt(1000))
								ϒend_time = λ.Cal(ϒfloat_or_none, ϒend_time, λ.NewInt(1000))
								if λ.IsTrue(func() λ.Object {
									if λv := λ.NewBool(ϒstart_time == λ.None); λ.IsTrue(λv) {
										return λv
									} else {
										return λ.NewBool(ϒend_time == λ.None)
									}
								}()) {
									return λ.None
								}
								λ.Cal(λ.GetAttr(ϒchapters, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("start_time"): ϒstart_time,
									λ.NewStr("end_time"):   ϒend_time,
								}))
								return λ.None
							})
						τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒtp_chapters, λ.NewSlice(λ.None, λ.Neg(λ.NewInt(1)), λ.None)))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒchapter = τmp1
							λ.Cal(ϒ_add_chapter, λ.Cal(λ.GetAttr(ϒchapter, "get", nil), λ.NewStr("startTime")), λ.Cal(λ.GetAttr(ϒchapter, "get", nil), λ.NewStr("endTime")))
						}
						λ.Cal(ϒ_add_chapter, λ.Cal(λ.GetAttr(λ.GetItem(ϒtp_chapters, λ.Neg(λ.NewInt(1))), "get", nil), λ.NewStr("startTime")), func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(λ.GetItem(ϒtp_chapters, λ.Neg(λ.NewInt(1))), "get", nil), λ.NewStr("endTime")); λ.IsTrue(λv) {
								return λv
							} else {
								return ϒduration
							}
						}())
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("title"):       λ.GetItem(ϒinfo, λ.NewStr("title")),
						λ.NewStr("subtitles"):   ϒsubtitles,
						λ.NewStr("description"): λ.GetItem(ϒinfo, λ.NewStr("description")),
						λ.NewStr("thumbnail"):   λ.GetItem(ϒinfo, λ.NewStr("defaultThumbnailUrl")),
						λ.NewStr("duration"):    λ.Cal(ϒfloat_or_none, ϒduration, λ.NewInt(1000)),
						λ.NewStr("timestamp"): func() λ.Object {
							if λv := λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("pubDate")), λ.NewInt(1000)); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.None
							}
						}(),
						λ.NewStr("uploader"): λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("billingCode")),
						λ.NewStr("chapters"): ϒchapters,
					})
				})
			ThePlatformBaseIE__extract_theplatform_metadata = λ.NewFunction("_extract_theplatform_metadata",
				[]λ.Param{
					{Name: "self"},
					{Name: "path"},
					{Name: "video_id"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinfo     λ.Object
						ϒpath     = λargs[1]
						ϒself     = λargs[0]
						ϒvideo_id = λargs[2]
					)
					ϒinfo = λ.Cal(λ.GetAttr(ϒself, "_download_theplatform_metadata", nil), ϒpath, ϒvideo_id)
					return λ.Cal(λ.GetAttr(ϒself, "_parse_theplatform_metadata", nil), ϒinfo)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TP_TLD"):                        ThePlatformBaseIE__TP_TLD,
				λ.NewStr("_download_theplatform_metadata"): ThePlatformBaseIE__download_theplatform_metadata,
				λ.NewStr("_extract_theplatform_metadata"):  ThePlatformBaseIE__extract_theplatform_metadata,
				λ.NewStr("_extract_theplatform_smil"):      ThePlatformBaseIE__extract_theplatform_smil,
				λ.NewStr("_parse_theplatform_metadata"):    ThePlatformBaseIE__parse_theplatform_metadata,
			})
		}())
		ThePlatformIE = λ.Cal(λ.TypeType, λ.NewStr("ThePlatformIE"), λ.NewTuple(
			ThePlatformBaseIE,
			AdobePassIE,
		), func() λ.Dict {
			var (
				ThePlatformIE__TESTS        λ.Object
				ThePlatformIE__VALID_URL    λ.Object
				ThePlatformIE__real_extract λ.Object
			)
			ThePlatformIE__VALID_URL = λ.NewStr("(?x)\n        (?:https?://(?:link|player)\\.theplatform\\.com/[sp]/(?P<provider_id>[^/]+)/\n           (?:(?:(?:[^/]+/)+select/)?(?P<media>media/(?:guid/\\d+/)?)?|(?P<config>(?:[^/\\?]+/(?:swf|config)|onsite)/select/))?\n         |theplatform:)(?P<id>[^/\\?&]+)")
			ThePlatformIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://link.theplatform.com/s/dJ5BDC/e9I_cZgTgIPd/meta.smil?format=smil&Tracking=true&mbr=true"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("e9I_cZgTgIPd"),
						λ.NewStr("ext"):         λ.NewStr("flv"),
						λ.NewStr("title"):       λ.NewStr("Blackberry's big, bold Z30"),
						λ.NewStr("description"): λ.NewStr("The Z30 is Blackberry's biggest, baddest mobile messaging device yet."),
						λ.NewStr("duration"):    λ.NewInt(247),
						λ.NewStr("timestamp"):   λ.NewInt(1383239700),
						λ.NewStr("upload_date"): λ.NewStr("20131031"),
						λ.NewStr("uploader"):    λ.NewStr("CBSI-NEW"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("skip"): λ.NewStr("404 Not Found"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://link.theplatform.com/s/kYEXFC/22d_qsQ6MIRT"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("22d_qsQ6MIRT"),
						λ.NewStr("ext"):         λ.NewStr("flv"),
						λ.NewStr("description"): λ.NewStr("md5:ac330c9258c04f9d7512cf26b9595409"),
						λ.NewStr("title"):       λ.NewStr("Tesla Model S: A second step towards a cleaner motoring future"),
						λ.NewStr("timestamp"):   λ.NewInt(1426176191),
						λ.NewStr("upload_date"): λ.NewStr("20150312"),
						λ.NewStr("uploader"):    λ.NewStr("CBSI-NEW"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://player.theplatform.com/p/D6x-PC/pulse_preview/embed/select/media/yMBg9E8KFxZD"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("yMBg9E8KFxZD"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("description"): λ.NewStr("md5:644ad9188d655b742f942bf2e06b002d"),
						λ.NewStr("title"):       λ.NewStr("HIGHLIGHTS: USA bag first ever series Cup win"),
						λ.NewStr("uploader"):    λ.NewStr("EGSM"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://player.theplatform.com/p/NnzsPC/widget/select/media/4Y0TlYUr_ZT7"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://player.theplatform.com/p/2E2eJC/nbcNewsOffsite?guid=tdy_or_siri_150701"),
					λ.NewStr("md5"): λ.NewStr("fb96bb3d85118930a5b055783a3bd992"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("tdy_or_siri_150701"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("iPhone Siri’s sassy response to a math question has people talking"),
						λ.NewStr("description"): λ.NewStr("md5:a565d1deadd5086f3331d57298ec6333"),
						λ.NewStr("duration"):    λ.NewFloat(83.0),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*\\.jpg$"),
						λ.NewStr("timestamp"):   λ.NewInt(1435752600),
						λ.NewStr("upload_date"): λ.NewStr("20150701"),
						λ.NewStr("uploader"):    λ.NewStr("NBCU-NEWS"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://player.theplatform.com/p/NnzsPC/onsite_universal/select/media/guid/2410887629/2928790?fwsitesection=nbc_the_blacklist_video_library&autoPlay=true&carouselID=137781"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			ThePlatformIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcombined_subtitles λ.Object
						ϒconfig             λ.Object
						ϒconfig_url         λ.Object
						ϒfeed_id            λ.Object
						ϒfeed_script        λ.Object
						ϒformats            λ.Object
						ϒheaders            λ.Object
						ϒmobj               λ.Object
						ϒpath               λ.Object
						ϒprovider_id        λ.Object
						ϒqs_dict            λ.Object
						ϒrelease_url        λ.Object
						ϒrequest            λ.Object
						ϒret                λ.Object
						ϒscript             λ.Object
						ϒscripts            λ.Object
						ϒself               = λargs[0]
						ϒsig                λ.Object
						ϒsmil_url           λ.Object
						ϒsmuggled_data      λ.Object
						ϒsource_url         λ.Object
						ϒsubtitles          λ.Object
						ϒurl                = λargs[1]
						ϒvideo_id           λ.Object
						ϒwebpage            λ.Object
						τmp0                λ.Object
						τmp1                λ.Object
					)
					τmp0 = λ.Cal(ϒunsmuggle_url, ϒurl, λ.NewDictWithTable(map[λ.Object]λ.Object{}))
					ϒurl = λ.GetItem(τmp0, λ.NewInt(0))
					ϒsmuggled_data = λ.GetItem(τmp0, λ.NewInt(1))
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒprovider_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("provider_id"))
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒprovider_id))) {
						ϒprovider_id = λ.NewStr("dJ5BDC")
					}
					ϒpath = λ.Add(ϒprovider_id, λ.NewStr("/"))
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("media"))) {
						τmp0 = λ.IAdd(ϒpath, λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("media")))
						ϒpath = τmp0
					}
					τmp0 = λ.IAdd(ϒpath, ϒvideo_id)
					ϒpath = τmp0
					ϒqs_dict = λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒurl), "query", nil))
					if λ.IsTrue(λ.NewBool(λ.Contains(ϒqs_dict, λ.NewStr("guid")))) {
						ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒvideo_id)
						ϒscripts = λ.Cal(Ωre.ϒfindall, λ.NewStr("<script[^>]+src=\"([^\"]+)\""), ϒwebpage)
						ϒfeed_id = λ.None
						τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.ReversedIteratorType, ϒscripts))
						for {
							if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
								break
							}
							ϒscript = τmp1
							ϒfeed_script = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), λ.Cal(λ.GetAttr(ϒself, "_proto_relative_url", nil), ϒscript, λ.NewStr("http:")), ϒvideo_id, λ.NewStr("Downloading feed script"))
							ϒfeed_id = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("defaultFeedId\\s*:\\s*\"([^\"]+)\""),
								ϒfeed_script,
								λ.NewStr("default feed id"),
							), λ.KWArgs{
								{Name: "default", Value: λ.None},
							})
							if λ.IsTrue(λ.NewBool(ϒfeed_id != λ.None)) {
								break
							}
						}
						if λ.IsTrue(λ.NewBool(ϒfeed_id == λ.None)) {
							panic(λ.Raise(λ.Cal(ExtractorError, λ.NewStr("Unable to find feed id"))))
						}
						return λ.Cal(λ.GetAttr(ϒself, "url_result", nil), λ.Mod(λ.NewStr("http://feed.theplatform.com/f/%s/%s?byGuid=%s"), λ.NewTuple(
							ϒprovider_id,
							ϒfeed_id,
							λ.GetItem(λ.GetItem(ϒqs_dict, λ.NewStr("guid")), λ.NewInt(0)),
						)))
					}
					if λ.IsTrue(λ.Cal(λ.GetAttr(ϒsmuggled_data, "get", nil), λ.NewStr("force_smil_url"), λ.False)) {
						ϒsmil_url = ϒurl
					} else {
						if λ.IsTrue(λ.NewBool(λ.Contains(ϒurl, λ.NewStr("/guid/")))) {
							ϒheaders = λ.NewDictWithTable(map[λ.Object]λ.Object{})
							ϒsource_url = λ.Cal(λ.GetAttr(ϒsmuggled_data, "get", nil), λ.NewStr("source_url"))
							if λ.IsTrue(ϒsource_url) {
								λ.SetItem(ϒheaders, λ.NewStr("Referer"), ϒsource_url)
							}
							ϒrequest = λ.Call(ϒsanitized_Request, λ.NewArgs(ϒurl), λ.KWArgs{
								{Name: "headers", Value: ϒheaders},
							})
							ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒrequest, ϒvideo_id)
							ϒsmil_url = λ.Call(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewArgs(
								λ.NewStr("<link[^>]+href=([\"\\'])(?P<url>.+?)\\1[^>]+type=[\"\\']application/smil\\+xml"),
								ϒwebpage,
								λ.NewStr("smil url"),
							), λ.KWArgs{
								{Name: "group", Value: λ.NewStr("url")},
							})
							ϒpath = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("link\\.theplatform\\.com/s/((?:[^/?#&]+/)+[^/?#&]+)"), ϒsmil_url, λ.NewStr("path"))
							τmp0 = λ.IAdd(ϒsmil_url, func() λ.Object {
								if λ.IsTrue(λ.NewBool(!λ.Contains(ϒsmil_url, λ.NewStr("?")))) {
									return λ.NewStr("?")
								} else {
									return λ.Add(λ.NewStr("&"), λ.NewStr("formats=m3u,mpeg4"))
								}
							}())
							ϒsmil_url = τmp0
						} else {
							if λ.IsTrue(λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("config"))) {
								ϒconfig_url = λ.Add(ϒurl, λ.NewStr("&form=json"))
								ϒconfig_url = λ.Cal(λ.GetAttr(ϒconfig_url, "replace", nil), λ.NewStr("swf/"), λ.NewStr("config/"))
								ϒconfig_url = λ.Cal(λ.GetAttr(ϒconfig_url, "replace", nil), λ.NewStr("onsite/"), λ.NewStr("onsite/config/"))
								ϒconfig = λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒconfig_url, ϒvideo_id, λ.NewStr("Downloading config"))
								if λ.IsTrue(λ.NewBool(λ.Contains(ϒconfig, λ.NewStr("releaseUrl")))) {
									ϒrelease_url = λ.GetItem(ϒconfig, λ.NewStr("releaseUrl"))
								} else {
									ϒrelease_url = λ.Mod(λ.NewStr("http://link.theplatform.com/s/%s?mbr=true"), ϒpath)
								}
								ϒsmil_url = λ.Add(ϒrelease_url, λ.NewStr("&formats=MPEG4&manifest=f4m"))
							} else {
								ϒsmil_url = λ.Mod(λ.NewStr("http://link.theplatform.com/s/%s?mbr=true"), ϒpath)
							}
						}
					}
					ϒsig = λ.Cal(λ.GetAttr(ϒsmuggled_data, "get", nil), λ.NewStr("sig"))
					if λ.IsTrue(ϒsig) {
						ϒsmil_url = λ.Cal(λ.GetAttr(ϒself, "_sign_url", nil), ϒsmil_url, λ.GetItem(ϒsig, λ.NewStr("key")), λ.GetItem(ϒsig, λ.NewStr("secret")))
					}
					τmp0 = λ.Cal(λ.GetAttr(ϒself, "_extract_theplatform_smil", nil), ϒsmil_url, ϒvideo_id)
					ϒformats = λ.GetItem(τmp0, λ.NewInt(0))
					ϒsubtitles = λ.GetItem(τmp0, λ.NewInt(1))
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒret = λ.Cal(λ.GetAttr(ϒself, "_extract_theplatform_metadata", nil), ϒpath, ϒvideo_id)
					ϒcombined_subtitles = λ.Cal(λ.GetAttr(ϒself, "_merge_subtitles", nil), λ.Cal(λ.GetAttr(ϒret, "get", nil), λ.NewStr("subtitles"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), ϒsubtitles)
					λ.Cal(λ.GetAttr(ϒret, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        ϒvideo_id,
						λ.NewStr("formats"):   ϒformats,
						λ.NewStr("subtitles"): ϒcombined_subtitles,
					}))
					return ϒret
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        ThePlatformIE__TESTS,
				λ.NewStr("_VALID_URL"):    ThePlatformIE__VALID_URL,
				λ.NewStr("_real_extract"): ThePlatformIE__real_extract,
			})
		}())
		ThePlatformFeedIE = λ.Cal(λ.TypeType, λ.NewStr("ThePlatformFeedIE"), λ.NewTuple(ThePlatformBaseIE), func() λ.Dict {
			var (
				ThePlatformFeedIE__URL_TEMPLATE      λ.Object
				ThePlatformFeedIE__VALID_URL         λ.Object
				ThePlatformFeedIE__extract_feed_info λ.Object
				ThePlatformFeedIE__real_extract      λ.Object
			)
			ThePlatformFeedIE__URL_TEMPLATE = λ.NewStr("%s//feed.theplatform.com/f/%s/%s?form=json&%s")
			ThePlatformFeedIE__VALID_URL = λ.NewStr("https?://feed\\.theplatform\\.com/f/(?P<provider_id>[^/]+)/(?P<feed_id>[^?/]+)\\?(?:[^&]+&)*(?P<filter>by(?:Gui|I)d=(?P<id>[^&]+))")
			ThePlatformFeedIE__extract_feed_info = λ.NewFunction("_extract_feed_info",
				[]λ.Param{
					{Name: "self"},
					{Name: "provider_id"},
					{Name: "feed_id"},
					{Name: "filter_query"},
					{Name: "video_id"},
					{Name: "custom_fields", Def: λ.None},
					{Name: "asset_types_query", Def: λ.NewDictWithTable(map[λ.Object]λ.Object{})},
					{Name: "account_id", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒaccount_id        = λargs[7]
						ϒasset_type        λ.Object
						ϒasset_types       λ.Object
						ϒasset_types_query = λargs[6]
						ϒcategories        λ.Object
						ϒcur_formats       λ.Object
						ϒcur_subtitles     λ.Object
						ϒcur_video_id      λ.Object
						ϒcustom_fields     = λargs[5]
						ϒduration          λ.Object
						ϒentry             λ.Object
						ϒfeed_id           = λargs[2]
						ϒfile_asset_types  λ.Object
						ϒfilter_query      = λargs[3]
						ϒfirst_video_id    λ.Object
						ϒformats           λ.Object
						ϒitem              λ.Object
						ϒmain_smil_url     λ.Object
						ϒprovider_id       = λargs[1]
						ϒquery             λ.Object
						ϒreal_url          λ.Object
						ϒret               λ.Object
						ϒself              = λargs[0]
						ϒsmil_url          λ.Object
						ϒsubtitles         λ.Object
						ϒthumbnails        λ.Object
						ϒtimestamp         λ.Object
						ϒvideo_id          = λargs[4]
						τmp0               λ.Object
						τmp1               λ.Object
						τmp2               λ.Object
						τmp3               λ.Object
						τmp4               λ.Object
					)
					ϒreal_url = λ.Mod(λ.GetAttr(ϒself, "_URL_TEMPLATE", nil), λ.NewTuple(
						λ.Cal(λ.GetAttr(ϒself, "http_scheme", nil)),
						ϒprovider_id,
						ϒfeed_id,
						ϒfilter_query,
					))
					ϒentry = λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), ϒreal_url, ϒvideo_id), λ.NewStr("entries")), λ.NewInt(0))
					ϒmain_smil_url = func() λ.Object {
						if λ.IsTrue(ϒaccount_id) {
							return λ.Mod(λ.NewStr("http://link.theplatform.com/s/%s/media/guid/%d/%s"), λ.NewTuple(
								ϒprovider_id,
								ϒaccount_id,
								λ.GetItem(ϒentry, λ.NewStr("guid")),
							))
						} else {
							return λ.Cal(λ.GetAttr(ϒentry, "get", nil), λ.NewStr("plmedia$publicUrl"))
						}
					}()
					ϒformats = λ.NewList()
					ϒsubtitles = λ.NewDictWithTable(map[λ.Object]λ.Object{})
					ϒfirst_video_id = λ.None
					ϒduration = λ.None
					ϒasset_types = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(ϒentry, λ.NewStr("media$content")))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒitem = τmp1
						ϒsmil_url = λ.GetItem(ϒitem, λ.NewStr("plfile$url"))
						ϒcur_video_id = λ.Cal(λ.GetAttr(ThePlatformIE, "_match_id", nil), ϒsmil_url)
						if λ.IsTrue(λ.NewBool(ϒfirst_video_id == λ.None)) {
							ϒfirst_video_id = ϒcur_video_id
							ϒduration = λ.Cal(ϒfloat_or_none, λ.Cal(λ.GetAttr(ϒitem, "get", nil), λ.NewStr("plfile$duration")))
						}
						ϒfile_asset_types = func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒitem, "get", nil), λ.NewStr("plfile$assetTypes")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.GetItem(λ.Cal(ϒcompat_parse_qs, λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒsmil_url), "query", nil)), λ.NewStr("assetTypes"))
							}
						}()
						τmp2 = λ.Cal(λ.BuiltinIter, ϒfile_asset_types)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒasset_type = τmp3
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒasset_types, ϒasset_type))) {
								continue
							}
							λ.Cal(λ.GetAttr(ϒasset_types, "append", nil), ϒasset_type)
							ϒquery = λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("mbr"):        λ.NewStr("true"),
								λ.NewStr("formats"):    λ.GetItem(ϒitem, λ.NewStr("plfile$format")),
								λ.NewStr("assetTypes"): ϒasset_type,
							})
							if λ.IsTrue(λ.NewBool(λ.Contains(ϒasset_types_query, ϒasset_type))) {
								λ.Cal(λ.GetAttr(ϒquery, "update", nil), λ.GetItem(ϒasset_types_query, ϒasset_type))
							}
							τmp4 = λ.Cal(λ.GetAttr(ϒself, "_extract_theplatform_smil", nil), λ.Cal(ϒupdate_url_query, func() λ.Object {
								if λv := ϒmain_smil_url; λ.IsTrue(λv) {
									return λv
								} else {
									return ϒsmil_url
								}
							}(), ϒquery), ϒvideo_id, λ.Mod(λ.NewStr("Downloading SMIL data for %s"), ϒasset_type))
							ϒcur_formats = λ.GetItem(τmp4, λ.NewInt(0))
							ϒcur_subtitles = λ.GetItem(τmp4, λ.NewInt(1))
							λ.Cal(λ.GetAttr(ϒformats, "extend", nil), ϒcur_formats)
							ϒsubtitles = λ.Cal(λ.GetAttr(ϒself, "_merge_subtitles", nil), ϒsubtitles, ϒcur_subtitles)
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnails = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒthumbnail λ.Object
									τmp0       λ.Object
									τmp1       λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒentry, "get", nil), λ.NewStr("media$thumbnails"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒthumbnail = τmp1
									λgy.Yield(λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("url"):    λ.GetItem(ϒthumbnail, λ.NewStr("plfile$url")),
										λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒthumbnail, "get", nil), λ.NewStr("plfile$width"))),
										λ.NewStr("height"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒthumbnail, "get", nil), λ.NewStr("plfile$height"))),
									}))
								}
								return λ.None
							})
						})))
					ϒtimestamp = λ.Call(ϒint_or_none, λ.NewArgs(λ.Cal(λ.GetAttr(ϒentry, "get", nil), λ.NewStr("media$availableDate"))), λ.KWArgs{
						{Name: "scale", Value: λ.NewInt(1000)},
					})
					ϒcategories = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
						nil,
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
								var (
									ϒitem λ.Object
									τmp0  λ.Object
									τmp1  λ.Object
								)
								τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒentry, "get", nil), λ.NewStr("media$categories"), λ.NewList()))
								for {
									if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
										break
									}
									ϒitem = τmp1
									λgy.Yield(λ.GetItem(ϒitem, λ.NewStr("media$name")))
								}
								return λ.None
							})
						})))
					ϒret = λ.Cal(λ.GetAttr(ϒself, "_extract_theplatform_metadata", nil), λ.Mod(λ.NewStr("%s/%s"), λ.NewTuple(
						ϒprovider_id,
						ϒfirst_video_id,
					)), ϒvideo_id)
					ϒsubtitles = λ.Cal(λ.GetAttr(ϒself, "_merge_subtitles", nil), ϒsubtitles, λ.GetItem(ϒret, λ.NewStr("subtitles")))
					λ.Cal(λ.GetAttr(ϒret, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):         ϒvideo_id,
						λ.NewStr("formats"):    ϒformats,
						λ.NewStr("subtitles"):  ϒsubtitles,
						λ.NewStr("thumbnails"): ϒthumbnails,
						λ.NewStr("duration"):   ϒduration,
						λ.NewStr("timestamp"):  ϒtimestamp,
						λ.NewStr("categories"): ϒcategories,
					}))
					if λ.IsTrue(ϒcustom_fields) {
						λ.Cal(λ.GetAttr(ϒret, "update", nil), λ.Cal(ϒcustom_fields, ϒentry))
					}
					return ϒret
				})
			ThePlatformFeedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒfeed_id      λ.Object
						ϒfilter_query λ.Object
						ϒmobj         λ.Object
						ϒprovider_id  λ.Object
						ϒself         = λargs[0]
						ϒurl          = λargs[1]
						ϒvideo_id     λ.Object
					)
					ϒmobj = λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("id"))
					ϒprovider_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("provider_id"))
					ϒfeed_id = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("feed_id"))
					ϒfilter_query = λ.Cal(λ.GetAttr(ϒmobj, "group", nil), λ.NewStr("filter"))
					return λ.Cal(λ.GetAttr(ϒself, "_extract_feed_info", nil), ϒprovider_id, ϒfeed_id, ϒfilter_query, ϒvideo_id)
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_URL_TEMPLATE"):      ThePlatformFeedIE__URL_TEMPLATE,
				λ.NewStr("_VALID_URL"):         ThePlatformFeedIE__VALID_URL,
				λ.NewStr("_extract_feed_info"): ThePlatformFeedIE__extract_feed_info,
				λ.NewStr("_real_extract"):      ThePlatformFeedIE__real_extract,
			})
		}())
	})
}
