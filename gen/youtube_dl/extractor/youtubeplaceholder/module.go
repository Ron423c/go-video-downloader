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
 * youtubeplaceholder/module.go: transpiled from https://github.com/ytdl-org/youtube-dl/blob/master/youtube_dl/extractor/youtubeplaceholder.py
 */

package youtubeplaceholder

import (
	Ωcommon "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/extractor/common"
	Ωutils "github.com/tenta-browser/go-video-downloader/gen/youtube_dl/utils"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ExtractorError       λ.Object
	InfoExtractor        λ.Object
	YoutubePlaceholderIE λ.Object
)

func init() {
	λ.InitModule(func() {
		InfoExtractor = Ωcommon.InfoExtractor
		ExtractorError = Ωutils.ExtractorError
		YoutubePlaceholderIE = λ.Cal(λ.TypeType, λ.NewStr("YoutubePlaceholderIE"), λ.NewTuple(InfoExtractor), func() λ.Dict {
			var (
				YoutubePlaceholderIE__PLAYLIST_ID_RE λ.Object
				YoutubePlaceholderIE__TEST           λ.Object
				YoutubePlaceholderIE__VALID_URL      λ.Object
				YoutubePlaceholderIE__real_extract   λ.Object
			)
			YoutubePlaceholderIE__PLAYLIST_ID_RE = λ.NewStr("(?:PL|LL|EC|UU|FL|RD|UL|TL)[0-9A-Za-z-_]{10,}")
			YoutubePlaceholderIE__VALID_URL = λ.Mod(λ.NewStr("(?x)^\n                     (\n                         (?:https?://|//)                                    # http(s):// or protocol-independent URL\n                         (?:(?:(?:(?:\\w+\\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie)?\\.com/|\n                            (?:www\\.)?deturl\\.com/www\\.youtube\\.com/|\n                            (?:www\\.)?pwnyoutube\\.com/|\n                            (?:www\\.)?hooktube\\.com/|\n                            (?:www\\.)?yourepeat\\.com/|\n                            tube\\.majestyc\\.net/|\n                            (?:www\\.)?invidio\\.us/|\n                            youtube\\.googleapis\\.com/)                        # the various hostnames, with wildcard subdomains\n                         (?:.*?\\#/)?                                          # handle anchor (#/) redirect urls\n                         (?:                                                  # the various things that can precede the ID:\n                             (?:(?:v|embed|e)/(?!videoseries))                # v/ or embed/ or e/\n                             |(?:                                             # or the v= param in all its forms\n                                 (?:(?:watch|movie)(?:_popup)?(?:\\.php)?/?)?  # preceding watch(_popup|.php) or nothing (like /?v=xxxx)\n                                 (?:\\?|\\#!?)                                  # the params delimiter ? or # or #!\n                                 (?:.*?[&;])??                                # any other preceding param (like /?s=tuff&v=xxxx or ?s=tuff&amp;v=V36LpHqtcDY)\n                                 v=\n                             )\n                         ))\n                         |(?:\n                            youtu\\.be|                                        # just youtu.be/xxxx\n                            vid\\.plus|                                        # or vid.plus/xxxx\n                            zwearz\\.com/watch|                                # or zwearz.com/watch/xxxx\n                         )/\n                         |(?:www\\.)?cleanvideosearch\\.com/media/action/yt/watch\\?videoId=\n                         )\n                     )?                                                       # all until now is optional -> you can pass the naked ID\n                     ([0-9A-Za-z_-]{11})                                      # here is it! the YouTube video ID\n                     (?!.*?\\blist=\n                        (?:\n                            %(playlist_id)s|                                  # combined list/video URLs are handled by the playlist IE\n                            WL                                                # WL are handled by the watch later IE\n                        )\n                     )\n                     #(?(1).+)?                                                # if we found the ID, everything can follow\n                     $"), λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("playlist_id"): YoutubePlaceholderIE__PLAYLIST_ID_RE,
			}))
			YoutubePlaceholderIE__TEST = λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("url"): λ.NewStr("https://www.youtube.com/watch?v=MuAGGZNfUkU"),
			})
			YoutubePlaceholderIE__real_extract = λ.NewFunction("_real_extract",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
						ϒurl  = λargs[1]
					)
					_ = ϒself
					_ = ϒurl
					panic(λ.Raise(λ.Call(ExtractorError, λ.NewArgs(λ.NewStr("YOUTUBE_PLACEHOLDER")), λ.KWArgs{
						{Name: "expected", Value: λ.True},
					})))
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("_PLAYLIST_ID_RE"): YoutubePlaceholderIE__PLAYLIST_ID_RE,
				λ.NewStr("_TEST"):           YoutubePlaceholderIE__TEST,
				λ.NewStr("_VALID_URL"):      YoutubePlaceholderIE__VALID_URL,
				λ.NewStr("_real_extract"):   YoutubePlaceholderIE__real_extract,
			})
		}())
	})
}
