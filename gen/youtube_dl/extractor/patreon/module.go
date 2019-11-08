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
 * patreon/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/patreon.py
 */

package patreon

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	InfoExtractor    λ.Object
	KNOWN_EXTENSIONS λ.Object
	PatreonIE        λ.Object
	ϒclean_html      λ.Object
	ϒdetermine_ext   λ.Object
	ϒint_or_none     λ.Object
	ϒmimetype2ext    λ.Object
	ϒparse_iso8601   λ.Object
	ϒstr_or_none     λ.Object
	ϒtry_get         λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ϒclean_html = Ωutils.ϒclean_html
		ϒdetermine_ext = Ωutils.ϒdetermine_ext
		ϒint_or_none = Ωutils.ϒint_or_none
		KNOWN_EXTENSIONS = Ωutils.KNOWN_EXTENSIONS
		ϒmimetype2ext = Ωutils.ϒmimetype2ext
		ϒparse_iso8601 = Ωutils.ϒparse_iso8601
		ϒstr_or_none = Ωutils.ϒstr_or_none
		ϒtry_get = Ωutils.ϒtry_get
		PatreonIE = λ.Cal(λ.TypeType, λ.NewStr("PatreonIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				PatreonIE__TESTS        λ.Object
				PatreonIE__VALID_URL    λ.Object
				PatreonIE__real_extract λ.Object
			)
			PatreonIE__VALID_URL = λ.NewStr("https?://(?:www\\.)?patreon\\.com/(?:creation\\?hid=|posts/(?:[\\w-]+-)?)(?P<id>\\d+)")
			PatreonIE__TESTS = λ.NewList(
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.patreon.com/creation?hid=743933"),
					λ.NewStr("md5"): λ.NewStr("e25505eec1053a6e6813b8ed369875cc"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("743933"),
						λ.NewStr("ext"):         λ.NewStr("mp3"),
						λ.NewStr("title"):       λ.NewStr("Episode 166: David Smalley of Dogma Debate"),
						λ.NewStr("description"): λ.NewStr("md5:713b08b772cd6271b9f3906683cfacdf"),
						λ.NewStr("uploader"):    λ.NewStr("Cognitive Dissonance Podcast"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*$"),
						λ.NewStr("timestamp"):   λ.NewInt(1406473987),
						λ.NewStr("upload_date"): λ.NewStr("20140727"),
						λ.NewStr("uploader_id"): λ.NewStr("87145"),
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("http://www.patreon.com/creation?hid=754133"),
					λ.NewStr("md5"): λ.NewStr("3eb09345bf44bf60451b8b0b81759d0a"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):        λ.NewStr("754133"),
						λ.NewStr("ext"):       λ.NewStr("mp3"),
						λ.NewStr("title"):     λ.NewStr("CD 167 Extra"),
						λ.NewStr("uploader"):  λ.NewStr("Cognitive Dissonance Podcast"),
						λ.NewStr("thumbnail"): λ.NewStr("re:^https?://.*$"),
					}),
					λ.NewStr("skip"): λ.NewStr("Patron-only content"),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"): λ.NewStr("https://www.patreon.com/creation?hid=1682498"),
					λ.NewStr("info_dict"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          λ.NewStr("SU4fj_aEMVw"),
						λ.NewStr("ext"):         λ.NewStr("mp4"),
						λ.NewStr("title"):       λ.NewStr("I'm on Patreon!"),
						λ.NewStr("uploader"):    λ.NewStr("TraciJHines"),
						λ.NewStr("thumbnail"):   λ.NewStr("re:^https?://.*$"),
						λ.NewStr("upload_date"): λ.NewStr("20150211"),
						λ.NewStr("description"): λ.NewStr("md5:c5a706b1f687817a3de09db1eb93acd4"),
						λ.NewStr("uploader_id"): λ.NewStr("TraciJHines"),
					}),
					λ.NewStr("params"): λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("noplaylist"):    λ.True,
						λ.NewStr("skip_download"): λ.True,
					}),
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.patreon.com/posts/episode-166-of-743933"),
					λ.NewStr("only_matching"): λ.True,
				}),
				λ.NewDictWithTable(map[λ.Object]λ.Object{
					λ.NewStr("url"):           λ.NewStr("https://www.patreon.com/posts/743933"),
					λ.NewStr("only_matching"): λ.True,
				}),
			)
			λ.NewStr("\n    def _login(self):\n        username, password = self._get_login_info()\n        if username is None:\n            return\n\n        login_form = {\n            'redirectUrl': 'http://www.patreon.com/',\n            'email': username,\n            'password': password,\n        }\n\n        request = sanitized_Request(\n            'https://www.patreon.com/processLogin',\n            compat_urllib_parse_urlencode(login_form).encode('utf-8')\n        )\n        login_page = self._download_webpage(request, None, note='Logging in')\n\n        if re.search(r'onLoginFailed', login_page):\n            raise ExtractorError('Unable to login, incorrect username and/or password', expected=True)\n\n    def _real_initialize(self):\n        self._login()\n    ")
			PatreonIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒattributes       λ.Object
						ϒdownload_url     λ.Object
						ϒembed_url        λ.Object
						ϒext              λ.Object
						ϒi                λ.Object
						ϒi_type           λ.Object
						ϒimage            λ.Object
						ϒinfo             λ.Object
						ϒmedia_attributes λ.Object
						ϒpost             λ.Object
						ϒpost_file        λ.Object
						ϒself             = λargs[0]
						ϒtitle            λ.Object
						ϒurl              = λargs[1]
						ϒuser_attributes  λ.Object
						ϒvideo_id         λ.Object
						τmp0              λ.Object
						τmp1              λ.Object
					)
					ϒvideo_id = λ.Cal(λ.GetAttr(ϒself, "_match_id", nil), ϒurl)
					ϒpost = λ.Call(λ.GetAttr(ϒself, "_download_json", nil), λ.NewArgs(
						λ.Add(λ.NewStr("https://www.patreon.com/api/posts/"), ϒvideo_id),
						ϒvideo_id,
					), λ.KWArgs{
						{Name: "query", Value: λ.NewDictWithTable(map[λ.Object]λ.Object{
							λ.NewStr("fields[media]"):                 λ.NewStr("download_url,mimetype,size_bytes"),
							λ.NewStr("fields[post]"):                  λ.NewStr("comment_count,content,embed,image,like_count,post_file,published_at,title"),
							λ.NewStr("fields[user]"):                  λ.NewStr("full_name,url"),
							λ.NewStr("json-api-use-default-includes"): λ.NewStr("false"),
							λ.NewStr("include"):                       λ.NewStr("media,user"),
						})},
					})
					ϒattributes = λ.GetItem(λ.GetItem(ϒpost, λ.NewStr("data")), λ.NewStr("attributes"))
					ϒtitle = λ.Cal(λ.GetAttr(λ.GetItem(ϒattributes, λ.NewStr("title")), "strip", nil))
					ϒimage = func() λ.Object {
						if λv := λ.Cal(λ.GetAttr(ϒattributes, "get", nil), λ.NewStr("image")); λ.IsTrue(λv) {
							return λv
						} else {
							return λ.NewDictWithTable(map[λ.Object]λ.Object{})
						}
					}()
					ϒinfo = λ.NewDictWithTable(map[λ.Object]λ.Object{
						λ.NewStr("id"):          ϒvideo_id,
						λ.NewStr("title"):       ϒtitle,
						λ.NewStr("description"): λ.Cal(ϒclean_html, λ.Cal(λ.GetAttr(ϒattributes, "get", nil), λ.NewStr("content"))),
						λ.NewStr("thumbnail"): func() λ.Object {
							if λv := λ.Cal(λ.GetAttr(ϒimage, "get", nil), λ.NewStr("large_url")); λ.IsTrue(λv) {
								return λv
							} else {
								return λ.Cal(λ.GetAttr(ϒimage, "get", nil), λ.NewStr("url"))
							}
						}(),
						λ.NewStr("timestamp"):     λ.Cal(ϒparse_iso8601, λ.Cal(λ.GetAttr(ϒattributes, "get", nil), λ.NewStr("published_at"))),
						λ.NewStr("like_count"):    λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒattributes, "get", nil), λ.NewStr("like_count"))),
						λ.NewStr("comment_count"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒattributes, "get", nil), λ.NewStr("comment_count"))),
					})
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(ϒpost, "get", nil), λ.NewStr("included"), λ.NewList()))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						ϒi = τmp1
						ϒi_type = λ.Cal(λ.GetAttr(ϒi, "get", nil), λ.NewStr("type"))
						if λ.IsTrue(λ.Eq(ϒi_type, λ.NewStr("media"))) {
							ϒmedia_attributes = func() λ.Object {
								if λv := λ.Cal(λ.GetAttr(ϒi, "get", nil), λ.NewStr("attributes")); λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewDictWithTable(map[λ.Object]λ.Object{})
								}
							}()
							ϒdownload_url = λ.Cal(λ.GetAttr(ϒmedia_attributes, "get", nil), λ.NewStr("download_url"))
							ϒext = λ.Cal(ϒmimetype2ext, λ.Cal(λ.GetAttr(ϒmedia_attributes, "get", nil), λ.NewStr("mimetype")))
							if λ.IsTrue(func() λ.Object {
								if λv := ϒdownload_url; !λ.IsTrue(λv) {
									return λv
								} else {
									return λ.NewBool(λ.Contains(KNOWN_EXTENSIONS, ϒext))
								}
							}()) {
								λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
									λ.NewStr("ext"):      ϒext,
									λ.NewStr("filesize"): λ.Cal(ϒint_or_none, λ.Cal(λ.GetAttr(ϒmedia_attributes, "get", nil), λ.NewStr("size_bytes"))),
									λ.NewStr("url"):      ϒdownload_url,
								}))
							}
						} else {
							if λ.IsTrue(λ.Eq(ϒi_type, λ.NewStr("user"))) {
								ϒuser_attributes = λ.Cal(λ.GetAttr(ϒi, "get", nil), λ.NewStr("attributes"))
								if λ.IsTrue(ϒuser_attributes) {
									λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
										λ.NewStr("uploader"):     λ.Cal(λ.GetAttr(ϒuser_attributes, "get", nil), λ.NewStr("full_name")),
										λ.NewStr("uploader_id"):  λ.Cal(ϒstr_or_none, λ.Cal(λ.GetAttr(ϒi, "get", nil), λ.NewStr("id"))),
										λ.NewStr("uploader_url"): λ.Cal(λ.GetAttr(ϒuser_attributes, "get", nil), λ.NewStr("url")),
									}))
								}
							}
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("url"))))) {
						ϒembed_url = λ.Cal(ϒtry_get, ϒattributes, λ.NewFunction("<lambda>",
							[]λ.Param{
								{Name: "x"},
							},
							0, false, false,
							func(λargs []λ.Object) λ.Object {
								var (
									ϒx = λargs[0]
								)
								return λ.GetItem(λ.GetItem(ϒx, λ.NewStr("embed")), λ.NewStr("url"))
							}))
						if λ.IsTrue(ϒembed_url) {
							λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("_type"): λ.NewStr("url"),
								λ.NewStr("url"):   ϒembed_url,
							}))
						}
					}
					if λ.IsTrue(λ.NewBool(!λ.IsTrue(λ.Cal(λ.GetAttr(ϒinfo, "get", nil), λ.NewStr("url"))))) {
						ϒpost_file = λ.GetItem(ϒattributes, λ.NewStr("post_file"))
						ϒext = λ.Cal(ϒdetermine_ext, λ.Cal(λ.GetAttr(ϒpost_file, "get", nil), λ.NewStr("name")))
						if λ.IsTrue(λ.NewBool(λ.Contains(KNOWN_EXTENSIONS, ϒext))) {
							λ.Cal(λ.GetAttr(ϒinfo, "update", nil), λ.NewDictWithTable(map[λ.Object]λ.Object{
								λ.NewStr("ext"): ϒext,
								λ.NewStr("url"): λ.GetItem(ϒpost_file, λ.NewStr("url")),
							}))
						}
					}
					return ϒinfo
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_TESTS"):        PatreonIE__TESTS,
				λ.NewStr("_VALID_URL"):    PatreonIE__VALID_URL,
				λ.NewStr("_real_extract"): PatreonIE__real_extract,
			})
		}())
	})
}
