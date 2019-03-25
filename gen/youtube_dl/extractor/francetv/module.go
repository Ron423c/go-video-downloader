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
 * francetv/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/francetv.py
 */

package francetv

import (
	Ωcompat "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/compat"
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωdailymotion "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/dailymotion"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	CultureboxIE              λ.Object
	DailymotionIE             λ.Object
	ExtractorError            λ.Object
	FranceTVBaseInfoExtractor λ.Object
	FranceTVEmbedIE           λ.Object
	FranceTVIE                λ.Object
	FranceTVInfoIE            λ.Object
	FranceTVInfoSportIE       λ.Object
	FranceTVJeunesseIE        λ.Object
	FranceTVSiteIE            λ.Object
	GenerationWhatIE          λ.Object
	InfoExtractor             λ.Object
	ϒclean_html               λ.Object
	ϒcompat_str               λ.Object
	ϒdetermine_ext            λ.Object
	ϒint_or_none              λ.Object
	ϒparse_duration           λ.Object
	ϒtry_get                  λ.Object
	ϒurl_or_none              λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒcompat_str = Ωcompat.ϒcompat_str
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ExtractorError = Ωutils.ExtractorError
		ϒint_or_none = Ωutils.ϒint_or_none
		ϒparse_duration = Ωutils.ϒparse_duration
		ϒtry_get = Ωutils.ϒtry_get
		ϒurl_or_none = Ωutils.ϒurl_or_none
		DailymotionIE = Ωdailymotion.DailymotionIE
		FranceTVBaseInfoExtractor = λ.Cal(λ.TypeType, λ.NewStr("FranceTVBaseInfoExtractor"), λ.NewTuple(InfoExtractor), func() λ.Dict {

			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
		FranceTVIE = λ.Cal(λ.TypeType, λ.NewStr("FranceTVIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				FranceTVIE__VALID_URL λ.Object
			)
			FranceTVIE__VALID_URL = λ.NewStr("(?x)\n                    (?:\n                        https?://\n                            sivideo\\.webservices\\.francetelevisions\\.fr/tools/getInfosOeuvre/v2/\\?\n                            .*?\\bidDiffusion=[^&]+|\n                        (?:\n                            https?://videos\\.francetv\\.fr/video/|\n                            francetv:\n                        )\n                        (?P<id>[^@]+)(?:@(?P<catalog>.+))?\n                    )\n                    ")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FranceTVIE__VALID_URL,
			})
		}())
		FranceTVSiteIE = λ.Cal(λ.TypeType, λ.NewStr("FranceTVSiteIE"), λ.NewTuple(FranceTVBaseInfoExtractor), func() λ.Dict {
			var (
				FranceTVSiteIE__VALID_URL λ.Object
			)
			FranceTVSiteIE__VALID_URL = λ.NewStr("https?://(?:(?:www\\.)?france\\.tv|mobile\\.france\\.tv)/(?:[^/]+/)*(?P<id>[^/]+)\\.html")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FranceTVSiteIE__VALID_URL,
			})
		}())
		FranceTVEmbedIE = λ.Cal(λ.TypeType, λ.NewStr("FranceTVEmbedIE"), λ.NewTuple(FranceTVBaseInfoExtractor), func() λ.Dict {
			var (
				FranceTVEmbedIE__VALID_URL λ.Object
			)
			FranceTVEmbedIE__VALID_URL = λ.NewStr("https?://embed\\.francetv\\.fr/*\\?.*?\\bue=(?P<id>[^&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FranceTVEmbedIE__VALID_URL,
			})
		}())
		FranceTVInfoIE = λ.Cal(λ.TypeType, λ.NewStr("FranceTVInfoIE"), λ.NewTuple(FranceTVBaseInfoExtractor), func() λ.Dict {
			var (
				FranceTVInfoIE__VALID_URL λ.Object
			)
			FranceTVInfoIE__VALID_URL = λ.NewStr("https?://(?:www|mobile|france3-regions)\\.francetvinfo\\.fr/(?:[^/]+/)*(?P<id>[^/?#&.]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FranceTVInfoIE__VALID_URL,
			})
		}())
		FranceTVInfoSportIE = λ.Cal(λ.TypeType, λ.NewStr("FranceTVInfoSportIE"), λ.NewTuple(FranceTVBaseInfoExtractor), func() λ.Dict {
			var (
				FranceTVInfoSportIE__VALID_URL λ.Object
			)
			FranceTVInfoSportIE__VALID_URL = λ.NewStr("https?://sport\\.francetvinfo\\.fr/(?:[^/]+/)*(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FranceTVInfoSportIE__VALID_URL,
			})
		}())
		GenerationWhatIE = λ.Cal(λ.TypeType, λ.NewStr("GenerationWhatIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				GenerationWhatIE_IE_NAME       λ.Object
				GenerationWhatIE__TESTS        λ.Object
				GenerationWhatIE__VALID_URL    λ.Object
				GenerationWhatIE__real_extract λ.Object
			)
			GenerationWhatIE_IE_NAME = λ.NewStr("france2.fr:generation-what")
			GenerationWhatIE__VALID_URL = λ.NewStr("https?://generation-what\\.francetv\\.fr/[^/]+/video/(?P<id>[^/?#&]+)")
			GenerationWhatIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://generation-what.francetv.fr/portrait/video/present-arms"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("wtvKYUG45iw"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("Generation What - Garde à vous - FRA"),
						λ.NewStr("uploader"):    λ.NewStr("Generation What"),
						λ.NewStr("uploader_id"): λ.NewStr("UCHH9p1eetWCgt4kXBYCb3_w"),
						λ.NewStr("upload_date"): λ.NewStr("20160411"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("skip_download"): λ.True,
					}),
					λ.NewStr("add_ie"): λ.NewList(λ.NewStr("Youtube")),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("http://generation-what.francetv.fr/europe/video/present-arms"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			GenerationWhatIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒdisplay_id λ.Object
						ϒself       = λargs[0]
						ϒurl        = λargs[1]
						ϒwebpage    λ.Object
						ϒyoutube_id λ.Object
					)
					ϒdisplay_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒwebpage = λ.Cal(λ.GetAttr(ϒself, "_download_webpage", nil), ϒurl, ϒdisplay_id)
					ϒyoutube_id = λ.Cal(λ.GetAttr(ϒself, "_search_regex", nil), λ.NewStr("window\\.videoURL\\s*=\\s*'([0-9A-Za-z_-]{11})';"), ϒwebpage, λ.NewStr("youtube id"))
					return λ.Call(λ.GetAttr(ϒself, "url_result", nil), λ.NewArgs(ϒyoutube_id), λ.KWArgs{
						{Name: "ie", Value: λ.NewStr("Youtube")},
						{Name: "video_id", Value: ϒyoutube_id},
					})
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("IE_NAME"):       GenerationWhatIE_IE_NAME,
				λ.NewStr("_TESTS"):        GenerationWhatIE__TESTS,
				λ.NewStr("_VALID_URL"):    GenerationWhatIE__VALID_URL,
				λ.NewStr("_real_extract"): GenerationWhatIE__real_extract,
			})
		}())
		CultureboxIE = λ.Cal(λ.TypeType, λ.NewStr("CultureboxIE"), λ.NewTuple(FranceTVBaseInfoExtractor), func() λ.Dict {
			var (
				CultureboxIE__VALID_URL λ.Object
			)
			CultureboxIE__VALID_URL = λ.NewStr("https?://(?:m\\.)?culturebox\\.francetvinfo\\.fr/(?:[^/]+/)*(?P<id>[^/?#&]+)")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): CultureboxIE__VALID_URL,
			})
		}())
		FranceTVJeunesseIE = λ.Cal(λ.TypeType, λ.NewStr("FranceTVJeunesseIE"), λ.NewTuple(FranceTVBaseInfoExtractor), func() λ.Dict {
			var (
				FranceTVJeunesseIE__VALID_URL λ.Object
			)
			FranceTVJeunesseIE__VALID_URL = λ.NewStr("(?P<url>https?://(?:www\\.)?(?:zouzous|ludo)\\.fr/heros/(?P<id>[^/?#&]+))")
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_VALID_URL"): FranceTVJeunesseIE__VALID_URL,
			})
		}())
	})
}
