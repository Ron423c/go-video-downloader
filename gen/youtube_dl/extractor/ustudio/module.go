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
 * ustudio/module.go: transpiled from https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/ustudio.py
 */

package ustudio

import (
	Ωre "github.com/tenta-browser/go-video-downloader/gen/re"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	UstudioEmbedIE   λ.Object
	UstudioIE        λ.Object
	ϒint_or_none     λ.Object
	ϒunescapeHTML    λ.Object
	ϒunified_strdate λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒunified_strdate = Ωutils.ϒunified_strdate
		ϒunescapeHTML = Ωutils.ϒunescapeHTML
		UstudioIE = λ.Cal(λ.TypeType, λ.NewStr("UstudioIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UstudioIE__VALID_URL λ.Object
			)
			UstudioIE__VALID_URL = λ.NewStr("https?://(?:(?:www|v1)\\.)?ustudio\\.com/video/(?P<id>[^/]+)/(?P<display_id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): UstudioIE__VALID_URL,
			})
		}())
		UstudioEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("UstudioEmbedIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				UstudioEmbedIE_IE_NAME       λ.Object
				UstudioEmbedIE__TEST         λ.Object
				UstudioEmbedIE__VALID_URL    λ.Object
				UstudioEmbedIE__real_extract λ.Object
			)
			UstudioEmbedIE_IE_NAME = λ.NewStr("ustudio:embed")
			UstudioEmbedIE__VALID_URL = λ.NewStr("https?://(?:(?:app|embed)\\.)?ustudio\\.com/embed/(?P<uid>[^/]+)/(?P<id>[^/]+)")
			UstudioEmbedIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("http://app.ustudio.com/embed/DeN7VdYRDKhP/Uw7G1kMCe65T"),
				λ.NewStr("md5"): λ.NewStr("47c0be52a09b23a7f40de9469cec58f4"),
				λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("id"):          λ.NewStr("Uw7G1kMCe65T"),
					λ.NewStr("ext"):         λ.NewStr("mp4"),
					λ.NewStr("title"):       λ.NewStr("5 Things IT Should Know About Video"),
					λ.NewStr("description"): λ.NewStr("md5:93d32650884b500115e158c5677d25ad"),
					λ.NewStr("uploader_id"): λ.NewStr("DeN7VdYRDKhP"),
				}),
			})
			UstudioEmbedIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒext         λ.Object
						ϒformats     λ.Object
						ϒheight      λ.Object
						ϒimage       λ.Object
						ϒimage_url   λ.Object
						ϒqualities   λ.Object
						ϒquality     λ.Object
						ϒquality_url λ.Object
						ϒself        = λargs[0]
						ϒthumbnails  λ.Object
						ϒtitle       λ.Object
						ϒuploader_id λ.Object
						ϒurl         = λargs[1]
						ϒvideo_data  λ.Object
						ϒvideo_id    λ.Object
						τmp0         λ.Object
						τmp1         λ.Object
						τmp2         λ.Object
						τmp3         λ.Object
					)
					τmp0 = λ.Cal(λ.GetAttr(λ.Cal(Ωre.ϒmatch, λ.GetAttr(ϒself, "_VALID_URL", nil), ϒurl), "groups", nil))
					ϒuploader_id = λ.GetItem(τmp0, λ.NewInt(0))
					ϒvideo_id = λ.GetItem(τmp0, λ.NewInt(1))
					ϒvideo_data = λ.GetItem(λ.GetItem(λ.Cal(λ.GetAttr(ϒself, "_download_json", nil), λ.Mod(λ.NewStr("http://app.ustudio.com/embed/%s/%s/config.json"), λ.NewTuple(
						ϒuploader_id,
						ϒvideo_id,
					)), ϒvideo_id), λ.NewStr("videos")), λ.NewInt(0))
					ϒtitle = λ.GetItem(ϒvideo_data, λ.NewStr("name"))
					ϒformats = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("transcodes"), λ.NewDictWithTable(map[λ.Object]λ.Object{})), "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒext = λ.GetItem(τmp2, λ.NewInt(0))
						ϒqualities = λ.GetItem(τmp2, λ.NewInt(1))
						τmp2 = λ.Cal(λ.BuiltinIter, ϒqualities)
						for {
							if τmp3 = λ.NextDefault(τmp2, λ.AfterLast); τmp3 == λ.AfterLast {
								break
							}
							ϒquality = τmp3
							ϒquality_url = λ.Cal(λ.GetAttr(ϒquality, "get", nil), λ.NewStr("url"))
							if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒquality_url))) {
								continue
							}
							ϒheight = λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒquality, "get", nil), λ.NewStr("height")))
							λ.Cal(λ.GetAttr(ϒformats, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("format_id"): func() λ.Object {
									if λ.IsTrue(ϒheight) {
										return λ.Mod(λ.NewStr("%s-%dp"), λ.NewTuple(
											ϒext,
											ϒheight,
										))
									} else {
										return ϒext
									}
								}(),
								λ.NewStr("url"):    ϒquality_url,
								λ.NewStr("width"):  λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒquality, "get", nil), λ.NewStr("width"))),
								λ.NewStr("height"): ϒheight,
							}))
						}
					}
					λ.Cal(λ.GetAttr(ϒself, "_sort_formats", nil), ϒformats)
					ϒthumbnails = λ.NewList()
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("images"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒimage = τmp1
						ϒimage_url = λ.Cal(λ.GetAttr(ϒimage, "get", nil), λ.NewStr("url"))
						if λ.IsTrue(λ.NewBool(!λ.IsTrue(ϒimage_url))) {
							continue
						}
						λ.Cal(λ.GetAttr(ϒthumbnails, "append", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("url"): ϒimage_url,
						}))
					}
					return λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("description")),
						λ.NewStr("duration"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("duration"))),
						λ.NewStr("uploader_id"): ϒuploader_id,
						λ.NewStr("tags"):        λ.Cal(λ.GetAttr(ϒvideo_data, "get", nil), λ.NewStr("keywords")),
						λ.NewStr("thumbnails"):  ϒthumbnails,
						λ.NewStr("formats"):     ϒformats,
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       UstudioEmbedIE_IE_NAME,
				λ.NewStr("_TEST"):         UstudioEmbedIE__TEST,
				λ.NewStr("_VALID_URL"):    UstudioEmbedIE__VALID_URL,
				λ.NewStr("_real_extract"): UstudioEmbedIE__real_extract,
			})
		}())
	})
}
