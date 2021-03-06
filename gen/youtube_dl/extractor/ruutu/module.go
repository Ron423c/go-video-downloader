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
 * ruutu/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/ruutu.py
 */

package ruutu

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError                λ.Object
	InfoExtractor                 λ.Object
	RuutuIE                       λ.Object
	ϒcompat_urllib_parse_urlparse λ.Object
	ϒdetermine_ext                λ.Object
	ϒint_or_none                  λ.Object
	ϒxpath_attr                   λ.Object
	ϒxpath_text                   λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_urllib_parse_urlparse = Ωcompat.ϒcompat_urllib_parse_urlparse
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒxpath_attr = Ωutils.ϒxpath_attr
		ϒxpath_text = Ωutils.ϒxpath_text
		RuutuIE = λ.Cal(λ.TypeType, λ.StrLiteral("RuutuIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				RuutuIE__VALID_URL    λ.Object
				RuutuIE__real_extract λ.Object
			)
			RuutuIE__VALID_URL = λ.StrLiteral("https?://(?:www\\.)?(?:ruutu|supla)\\.fi/(?:video|supla)/(?P<id>\\d+)")
			RuutuIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdrm             λ.Object
						ϒextract_formats λ.Object
						ϒformats         λ.Object
						ϒprocessed_urls  λ.Object
						ϒself            = λargs[0]
						ϒurl             = λargs[1]
						ϒvideo_id        λ.Object
						ϒvideo_xml       λ.Object
					)
					ϒvideo_id = λ.Calm(ϒself, "_match_id", ϒurl)
					ϒvideo_xml = λ.Call(λ.GetAttr(ϒself, "_download_xml", nil), λ.NewArgs(
						λ.StrLiteral("https://gatling.nelonenmedia.fi/media-xml-cache"),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.DictLiteral(map[string]λ.Object{
							"id": ϒvideo_id,
						})},
					})
					ϒformats = λ.NewList()
					ϒprocessed_urls = λ.NewList()
					ϒextract_formats = λ.NewFunction("extract_formats",
						[]λ.Param{
							{Name: "node"},
						},
						0, false, false,
						func(λargs []λ.Object) λ.Object {
							var (
								ϒchild      λ.Object
								ϒext        λ.Object
								ϒformat_id  λ.Object
								ϒheight     λ.Object
								ϒlabel      λ.Object
								ϒnode       = λargs[0]
								ϒpreference λ.Object
								ϒproto      λ.Object
								ϒtbr        λ.Object
								ϒvideo_url  λ.Object
								ϒwidth      λ.Object
								τmp0        λ.Object
								τmp1        λ.Object
								τmp2        λ.Object
							)
							τmp0 = λ.Cal(λ.BuiltinIter, ϒnode)
							for {
								if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
									break
								}
								ϒchild = τmp1
								if λ.IsTrue(λ.Calm(λ.GetAttr(ϒchild, "tag", nil), "endswith", λ.StrLiteral("Files"))) {
									λ.Cal(ϒextract_formats, ϒchild)
								} else {
									if λ.IsTrue(λ.Calm(λ.GetAttr(ϒchild, "tag", nil), "endswith", λ.StrLiteral("File"))) {
										ϒvideo_url = λ.GetAttr(ϒchild, "text", nil)
										if λ.IsTrue(func() λ.Object {
											if λv := λ.NewBool(!λ.IsTrue(ϒvideo_url)); λ.IsTrue(λv) {
												return λv
											} else if λv := λ.NewBool(λ.Contains(ϒprocessed_urls, ϒvideo_url)); λ.IsTrue(λv) {
												return λv
											} else {
												return λ.Cal(λ.BuiltinAny, λ.Cal(λ.NewFunction("<generator>",
													nil,
													0, false, false,
													func(λargs []λ.Object) λ.Object {
														return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
															var (
																ϒp   λ.Object
																τmp0 λ.Object
																τmp1 λ.Object
															)
															τmp0 = λ.Cal(λ.BuiltinIter, λ.NewTuple(
																λ.StrLiteral("NOT_USED"),
																λ.StrLiteral("NOT-USED"),
															))
															for {
																if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
																	break
																}
																ϒp = τmp1
																λgy.Yield(λ.NewBool(λ.Contains(ϒvideo_url, ϒp)))
															}
															return λ.None
														})
													})))
											}
										}()) {
											continue
										}
										λ.Calm(ϒprocessed_urls, "append", ϒvideo_url)
										ϒext = λ.Cal(ϒdetermine_ext, ϒvideo_url)
										if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("m3u8"))) {
											λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_m3u8_formats", nil), λ.NewArgs(
												ϒvideo_url,
												ϒvideo_id,
												λ.StrLiteral("mp4"),
											), λ.KWArgs{
												{Name: "m3u8_id", Value: λ.StrLiteral("hls")},
												{Name: "fatal", Value: λ.False},
											}))
										} else {
											if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("f4m"))) {
												λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_f4m_formats", nil), λ.NewArgs(
													ϒvideo_url,
													ϒvideo_id,
												), λ.KWArgs{
													{Name: "f4m_id", Value: λ.StrLiteral("hds")},
													{Name: "fatal", Value: λ.False},
												}))
											} else {
												if λ.IsTrue(λ.Eq(ϒext, λ.StrLiteral("mpd"))) {
													continue
													λ.Calm(ϒformats, "extend", λ.Call(λ.GetAttr(ϒself, "_extract_mpd_formats", nil), λ.NewArgs(
														ϒvideo_url,
														ϒvideo_id,
													), λ.KWArgs{
														{Name: "mpd_id", Value: λ.StrLiteral("dash")},
														{Name: "fatal", Value: λ.False},
													}))
												} else {
													if λ.IsTrue(func() λ.Object {
														if λv := λ.Eq(ϒext, λ.StrLiteral("mp3")); λ.IsTrue(λv) {
															return λv
														} else {
															return λ.Eq(λ.GetAttr(ϒchild, "tag", nil), λ.StrLiteral("AudioMediaFile"))
														}
													}()) {
														λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
															"format_id": λ.StrLiteral("audio"),
															"url":       ϒvideo_url,
															"vcodec":    λ.StrLiteral("none"),
														}))
													} else {
														ϒproto = λ.GetAttr(λ.Cal(ϒcompat_urllib_parse_urlparse, ϒvideo_url), "scheme", nil)
														if λ.IsTrue(func() λ.Object {
															if λv := λ.NewBool(!λ.IsTrue(λ.Calm(λ.GetAttr(ϒchild, "tag", nil), "startswith", λ.StrLiteral("HTTP")))); !λ.IsTrue(λv) {
																return λv
															} else {
																return λ.Ne(ϒproto, λ.StrLiteral("rtmp"))
															}
														}()) {
															continue
														}
														ϒpreference = func() λ.Object {
															if λ.IsTrue(λ.Eq(ϒproto, λ.StrLiteral("rtmp"))) {
																return λ.Neg(λ.IntLiteral(1))
															} else {
																return λ.IntLiteral(1)
															}
														}()
														ϒlabel = λ.Calm(ϒchild, "get", λ.StrLiteral("label"))
														ϒtbr = λ.Cal(ϒint_or_none, λ.Calm(ϒchild, "get", λ.StrLiteral("bitrate")))
														ϒformat_id = func() λ.Object {
															if λ.IsTrue(func() λ.Object {
																if λv := ϒlabel; λ.IsTrue(λv) {
																	return λv
																} else {
																	return ϒtbr
																}
															}()) {
																return λ.Mod(λ.StrLiteral("%s-%s"), λ.NewTuple(
																	ϒproto,
																	func() λ.Object {
																		if λ.IsTrue(ϒlabel) {
																			return ϒlabel
																		} else {
																			return ϒtbr
																		}
																	}(),
																))
															} else {
																return ϒproto
															}
														}()
														if !λ.IsTrue(λ.Calm(ϒself, "_is_valid_url", ϒvideo_url, ϒvideo_id, ϒformat_id)) {
															continue
														}
														τmp2 = λ.Cal(λ.ListType, λ.Cal(λ.NewFunction("<generator>",
															nil,
															0, false, false,
															func(λargs []λ.Object) λ.Object {
																return λ.NewGenerator(func(λgy λ.Yielder) λ.Object {
																	var (
																		ϒx   λ.Object
																		τmp0 λ.Object
																		τmp1 λ.Object
																	)
																	τmp0 = λ.Cal(λ.BuiltinIter, λ.GetItem(λ.Calm(λ.Calm(ϒchild, "get", λ.StrLiteral("resolution"), λ.StrLiteral("x")), "split", λ.StrLiteral("x")), λ.NewSlice(λ.None, λ.IntLiteral(2), λ.None)))
																	for {
																		if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
																			break
																		}
																		ϒx = τmp1
																		λgy.Yield(λ.Cal(ϒint_or_none, ϒx))
																	}
																	return λ.None
																})
															})))
														ϒwidth = λ.GetItem(τmp2, λ.IntLiteral(0))
														ϒheight = λ.GetItem(τmp2, λ.IntLiteral(1))
														λ.Calm(ϒformats, "append", λ.DictLiteral(map[string]λ.Object{
															"format_id":  ϒformat_id,
															"url":        ϒvideo_url,
															"width":      ϒwidth,
															"height":     ϒheight,
															"tbr":        ϒtbr,
															"preference": ϒpreference,
														}))
													}
												}
											}
										}
									}
								}
							}
							return λ.None
						})
					λ.Cal(ϒextract_formats, λ.Calm(ϒvideo_xml, "find", λ.StrLiteral("./Clip")))
					ϒdrm = λ.Call(ϒxpath_text, λ.NewArgs(
						ϒvideo_xml,
						λ.StrLiteral("./Clip/DRM"),
					), λ.KWArgs{
						{Name: "default", Value: λ.None},
					})
					if λ.IsTrue(func() λ.Object {
						if λv := λ.NewBool(!λ.IsTrue(ϒformats)); !λ.IsTrue(λv) {
							return λv
						} else {
							return ϒdrm
						}
					}()) {
						panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.StrLiteral("This video is DRM protected.")), λ.KWArgs{
							{Name: "expected", Value: λ.True},
						})))
					}
					λ.Calm(ϒself, "_sort_formats", ϒformats)
					return λ.DictLiteral(map[string]λ.Object{
						"id": ϒvideo_id,
						"title": λ.Call(ϒxpath_attr, λ.NewArgs(
							ϒvideo_xml,
							λ.StrLiteral(".//Behavior/Program"),
							λ.StrLiteral("program_name"),
							λ.StrLiteral("title"),
						), λ.KWArgs{
							{Name: "fatal", Value: λ.True},
						}),
						"description": λ.Cal(ϒxpath_attr, ϒvideo_xml, λ.StrLiteral(".//Behavior/Program"), λ.StrLiteral("description"), λ.StrLiteral("description")),
						"thumbnail":   λ.Cal(ϒxpath_attr, ϒvideo_xml, λ.StrLiteral(".//Behavior/Startpicture"), λ.StrLiteral("href"), λ.StrLiteral("thumbnail")),
						"duration":    λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo_xml, λ.StrLiteral(".//Runtime"), λ.StrLiteral("duration"))),
						"age_limit":   λ.Cal(ϒint_or_none, λ.Cal(ϒxpath_text, ϒvideo_xml, λ.StrLiteral(".//AgeLimit"), λ.StrLiteral("age limit"))),
						"formats":     ϒformats,
					})
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"_VALID_URL":    RuutuIE__VALID_URL,
				"_real_extract": RuutuIE__real_extract,
			})
		}())
	})
}
