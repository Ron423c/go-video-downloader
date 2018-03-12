/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
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
 * masterregexp.go: Automatically generated regexp by combining URL validation expressions of all extractors
 */

package downloader

func init() {
	masterRegexp = `(?<extr_ABCOTVSClips>https?://clips\.abcotvs\.com/(?:[^/]+/)*video/(?:\d+))|(?<extr_ABCOTVS>https?://(?:abc(?:7(?:news|ny|chicago)?|11|13|30)|6abc)\.com(?:/[^/]+/(?:[^/]+))?/(?:\d+))|(?<extr_AdobeTV>https?://tv\.adobe\.com/(?:(?:fr|de|es|jp)/)?watch/(?:[^/]+)/(?:[^/]+))|(?<extr_AnySex>https?://(?:www\.)?anysex\.com/(?:\d+))|(?<extr_ATTTechChannel>https?://techchannel\.att\.com/play-video\.cfm/([^/]+/)*(?:.+))|(?<extr_AudioBoom>https?://(?:www\.)?audioboom\.com/(?:boos|posts)/(?:[0-9]+))|(?<extr_Bang>https?://(?:www\.)?bang\.com/video/(?:[0-9a-zA-Z_-]+))|(?<extr_Bigflix>https?://(?:www\.)?bigflix\.com/.+/(?:[0-9]+))|(?<extr_Bild>https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?<extr_Camdemy>https?://(?:www\.)?camdemy\.com/media/(?:\d+))|(?<extr_Canalc2>https?://(?:(?:www\.)?canalc2\.tv/video/|archives-canalc2\.u-strasbg\.fr/video\.asp\?.*\bidVideo=)(?:\d+))|(?<extr_Chirbit>https?://(?:www\.)?chirb\.it/(?:(?:wp|pl)/|fb_chirbit_player\.swf\?key=)?(?:[\da-zA-Z]+))|(?<extr_Clubic>https?://(?:www\.)?clubic\.com/video/(?:[^/]+/)*video.*-(?:[0-9]+)\.html)|(?<extr_Clyp>https?://(?:www\.)?clyp\.it/(?:[a-z0-9]+))|(?<extr_Criterion>https?://(?:www\.)?criterion\.com/films/(?:[0-9]+)-.+)|(?<extr_CrooksAndLiars>https?://embed\.crooksandliars\.com/(?:embed|v)/(?:[A-Za-z0-9]+))|(?<extr_Digiteka>https?://(?:www\.)?(?:digiteka\.net|ultimedia\.com)/(?:deliver/(?:generic|musique)(?:/[^/]+)*/(?:src|article)|default/index/video(?:generic|music)/id)/(?:[\d+a-z]+))|(?<extr_Dotsub>https?://(?:www\.)?dotsub\.com/view/(?:[^/]+))|(?<extr_Dumpert>(?:https?)://(?:www\.)?dumpert\.nl/(?:mediabase|embed)/(?:[0-9]+/[0-9a-zA-Z]+))|(?<extr_EbaumsWorld>https?://(?:www\.)?ebaumsworld\.com/videos/[^/]+/(?:\d+))|(?<extr_EchoMsk>https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?<extr_Fczenit>https?://(?:www\.)?fc-zenit\.ru/video/(?:[0-9]+))|(?<extr_FiveTV>http://(?:www\.)?5-tv\.ru/(?:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?<extr_GodTube>https?://(?:www\.)?godtube\.com/watch/\?v=(?:[\da-zA-Z]+))|(?<extr_Hark>https?://(?:www\.)?hark\.com/clips/(?:.+?)-.+)|(?<extr_HentaiStigma>^https?://hentai\.animestigma\.com/(?:[^/]+))|(?<extr_HistoricFilms>https?://(?:www\.)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?<extr_Hitbox>https?://(?:www\.)?(?:hitbox|smashcast)\.tv/(?:[^/]+/)*videos?/(?:[0-9]+))|(?<extr_HitRecord>https?://(?:www\.)?hitrecord\.org/records/(?:\d+))|(?<extr_Huajiao>https?://(?:www\.)?huajiao\.com/l/(?:[0-9]+))|(?<extr_Ir90Tv>https?://(?:www\.)?90tv\.ir/video/(?:[0-9]+)/.*)|(?<extr_Joj>(?:joj:|https?://media\.joj\.sk/embed/)(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}))|(?<extr_Lemonde>https?://(?:.+?\.)?lemonde\.fr/(?:[^/]+/)*(?:[^/]+)\.html)|(?<extr_LoveHomePorn>https?://(?:www\.)?lovehomeporn\.com/video/(?:\d+)(?:/(?:[^/?#&]+))?)|(?<extr_LRT>https?://(?:www\.)?lrt\.lt/mediateka/irasas/(?:[0-9]+))|(?<extr_MacGameStore>https?://(?:www\.)?macgamestore\.com/mediaviewer\.php\?trailer=(?:\d+))|(?<extr_MailRu>https?://(?:(?:www|m)\.)?my\.mail\.ru/(?:video/.*\#video=/?(?:(?:[^/]+/){3}\d+)|(?:(?:(?:[^/]+/){2})video/(?:[^/]+/\d+))\.html|(?:video/embed|\+/video/meta)/(?:\d+)))|(?<extr_Megaphone>https://player\.megaphone\.fm/(?:[A-Z0-9]+))|(?<extr_Mojvideo>https?://(?:www\.)?mojvideo\.com/video-(?:[^/]+)/(?:[a-f0-9]+))|(?<extr_Morningstar>https?://(?:(?:www|news)\.)morningstar\.com/[cC]over/video[cC]enter\.aspx\?id=(?:[0-9]+))|(?<extr_MovingImage>https?://movingimage\.nls\.uk/film/(?:\d+))|(?<extr_NDTV>https?://(?:[^/]+\.)?ndtv\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?:\d+))|(?<extr_NextMediaActionNews>https?://hk\.dv\.nextmedia\.com/actionnews/[^/]+/(?:\d+)/(?:\d+)/\d+)|(?<extr_NextMedia>https?://hk\.apple\.nextmedia\.com/[^/]+/[^/]+/(?:\d+)/(?:\d+))|(?<extr_NTVRu>https?://(?:www\.)?ntv\.ru/(?:[^/]+/)*(?:[^/?#&]+))|(?<extr_Piksel>https?://player\.piksel\.com/v/(?:[a-z0-9]+))|(?<extr_Pinkbike>https?://(?:(?:www\.)?pinkbike\.com/video/|es\.pinkbike\.org/i/kvid/kvid-y5\.swf\?id=)(?:[0-9]+))|(?<extr_PlayFM>https?://(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+))/?(?:$|[?#]))|(?<extr_Podomatic>(?:https?)://(?:(?:[^.]+)\.podomatic\.com/entry|(?:www\.)?podomatic\.com/podcasts/(?:[^/]+)/episodes)/(?:[^/?#&]+))|(?<extr_PornHd>https?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos/(?:\d+)(?:/(?:.+))?)|(?<extr_PornHub>https?://(?:(?:[a-z]+\.)?pornhub\.com/(?:(?:view_video\.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(?:[\da-z]+))|(?<extr_RadioJavan>https?://(?:www\.)?radiojavan\.com/videos/video/(?:[^/]+)/?)|(?<extr_RedTube>https?://(?:(?:www\.)?redtube\.com/|embed\.redtube\.com/\?.*?\bid=)(?:[0-9]+))|(?<extr_RockstarGames>https?://(?:www\.)?rockstargames\.com/videos(?:/video/|#?/?\?.*\bvideo=)(?:\d+))|(?<extr_Rozhlas>https?://(?:www\.)?prehravac\.rozhlas\.cz/audio/(?:[0-9]+))|(?<extr_RTBF>https?://(?:www\.)?rtbf\.be/(?:video/[^?]+\?.*\bid=|ouftivi/(?:[^/]+/)*[^?]+\?.*\bvideoId=|auvio/[^/]+\?.*id=)(?:\d+))|(?<extr_Sapo>https?://(?:(?:v2|www)\.)?videos\.sapo\.(?:pt|cv|ao|mz|tl)/(?:[\da-zA-Z]{20}))|(?<extr_Slutload>^https?://(?:\w+\.)?slutload\.com/video/[^/]+/(?:[^/]+)/?$)|(?<extr_Soundgasm>https?://(?:www\.)?soundgasm\.net/u/(?:[0-9a-zA-Z_-]+)/(?:[0-9a-zA-Z_-]+))|(?<extr_SpankBang>https?://(?:(?:www|m|[a-z]{2})\.)?spankbang\.com/(?:[\da-z]+)/video)|(?<extr_Spankwire>https?://(?:www\.)?(?:spankwire\.com/[^/]*/video(?:[0-9]+)/?))|(?<extr_Stitcher>https?://(?:www\.)?stitcher\.com/podcast/(?:[^/]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+)(?:[/#?&]|$))|(?<extr_Streamable>https?://streamable\.com/(?:[es]/)?(?:\w+))|(?<extr_SunPorno>https?://(?:(?:www\.)?sunporno\.com/videos|embeds\.sunporno\.com/embed)/(?:\d+))|(?<extr_Tass>https?://(?:tass\.ru|itar-tass\.com)/[^/]+/(?:\d+))|(?<extr_TeacherTube>https?://(?:www\.)?teachertube\.com/(viewVideo\.php\?video_id=|music\.php\?music_id=|video/(?:[\da-z-]+-)?|audio/)(?:\d+))|(?<extr_TeleMB>https?://(?:www\.)?telemb\.be/(?:.+?)_d_(?:\d+)\.html)|(?<extr_ThisAmericanLife>https?://(?:www\.)?thisamericanlife\.org/(?:radio-archives/episode/|play_full\.php\?play=)(?:\d+))|(?<extr_TinyPic>https?://(?:.+?\.)?tinypic\.com/player\.php\?v=(?:[^&]+)&s=\d+)|(?<extr_Tumblr>https?://(?:[^/?#&]+)\.tumblr\.com/(?:post|video)/(?:[0-9]+)(?:$|[/?#]))|(?<extr_TVCArticle>https?://(?:www\.)?tvc\.ru/(?!video/iframe/id/)(?:[^?#]+))|(?<extr_TVC>https?://(?:www\.)?tvc\.ru/video/iframe/id/(?:\d+))|(?<extr_TwentyMinuten>https?://(?:www\.)?20min\.ch/(?:videotv/*\?.*?\bvid=|videoplayer/videoplayer\.html\?.*?\bvideoId@)(?:\d+))|(?<extr_TwitchClips>https?://clips\.twitch\.tv/(?:[^/]+/)*(?:[^/?#&]+))|(?<extr_Unistra>https?://utv\.unistra\.fr/(?:index|video)\.php\?id_video\=(?:\d+))|(?<extr_UstudioEmbed>https?://(?:(?:app|embed)\.)?ustudio\.com/embed/(?:[^/]+)/(?:[^/]+))|(?<extr_VideosZ>https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-z]+/(?:[0-9]+)_)|(?<extr_Vidio>https?://(?:www\.)?vidio\.com/watch/(?:\d+)-(?:[^/?#&]+))|(?<extr_Vporn>https?://(?:www\.)?vporn\.com/[^/]+/(?:[^/]+)/(?:\d+))|(?<extr_WDRMobile>https?://mobile-ondemand\.wdr\.de/.*?/fsk(?:[0-9]+)/[0-9]+/[0-9]+/(?:[0-9]+)_(?:[0-9]+))|(?<extr_WeiboMobile>https?://m\.weibo\.cn/status/(?:[0-9]+)(\?.+)?)|(?<extr_XHamsterEmbed>https?://(?:.+?\.)?xhamster\.com/xembed\.php\?video=(?:\d+))|(?<extr_XHamster>https?://(?:.+?\.)?xhamster\.com/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?<extr_XNXX>https?://(?:video|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)|(?<extr_XTube>(?:xtube:|https?://(?:www\.)?xtube\.com/(?:watch\.php\?.*\bv=|video-watch/(?:embedded/)?(?:[^/]+)-))(?:[^/?&#]+))|(?<extr_XVideos>https?://(?:(?:www\.)?xvideos\.com/video|flashservice\.xvideos\.com/embedframe/|static-hw\.xvideos\.com/swf/xv-player\.swf\?.*?\bid_video=)(?:[0-9]+))|(?<extr_XXXYMovies>https?://(?:www\.)?xxxymovies\.com/videos/(?:\d+)/(?:[^/]+))|(?<extr_YinYueTai>https?://v\.yinyuetai\.com/video(?:/h5)?/(?:[0-9]+))|(?<extr_YouPorn>https?://(?:www\.)?youporn\.com/watch/(?:\d+)/(?:[^/?#&]+))|(?<extr_YourUpload>https?://(?:www\.)?(?:yourupload\.com/(?:watch|embed)|embed\.yourupload\.com)/(?:[A-Za-z0-9]+))|(?<extr_YoutubePlaceholder>^((?:https?://|//)(?:(?:(?:(?:\w+\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie)?\.com/|(?:www\.)?deturl\.com/www\.youtube\.com/|(?:www\.)?pwnyoutube\.com/|(?:www\.)?hooktube\.com/|(?:www\.)?yourepeat\.com/|tube\.majestyc\.net/|youtube\.googleapis\.com/)(?:.*?\#/)?(?:(?:(?:v|embed|e)/(?!videoseries))|(?:(?:(?:watch|movie)(?:_popup)?(?:\.php)?/?)?(?:\?|\#!?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.plus|zwearz\.com/watch|)/|(?:www\.)?cleanvideosearch\.com/media/action/yt/watch\?videoId=))?([0-9A-Za-z_-]{11})(?!.*?\blist=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL)[0-9A-Za-z-_]{10,}|WL))$)`
}
