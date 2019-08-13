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
 * masterregexp.go: Automatically generated regexp by combining URL validation expressions of all extractors
 */

// Code generated by transpiler. DO NOT EDIT.

package downloader

func init() {
	masterRegexps = []string{
		`(?:https?://(?:abcnews\.go\.com/(?:[^/]+/video/(?:[0-9a-z-]+)-|video/embed\?.*?\` +
			`bid=)|fivethirtyeight\.abcnews\.go\.com/video/embed/\d+/)(?:\d+))|(?:https?://(?` +
			`:(?:(?:embed|www)\.)?acast\.com/|play\.acast\.com/s/)(?:[^/]+)/(?:[^/#?]+))|(?:h` +
			`ttps?://video\.tv\.adobe\.com/v/(?:\d+))|(?:(?:aol-video:|https?://(?:www\.)?aol` +
			`\.(?:com|ca|co\.uk|de|jp)/video/(?:[^/]+/)*)(?:[0-9a-f]+))|(?:https?://(?:www\.)` +
			`?allocine\.fr/(?:article|video|film)/(?:fichearticle_gen_carticle=|player_gen_cm` +
			`edia=|fichefilm_gen_cfilm=|video-)(?:[0-9]+)(?:\.html)?)|(?:https?://(?:www\.)?a` +
			`parat\.com/(?:v/|video/video/embed/videohash/)(?:[a-zA-Z0-9]+))|(?:^https?://(?:` +
			`(?:(?:www|classic)\.)?ardmediathek\.de|mediathek\.(?:daserste|rbb-online)\.de|on` +
			`e\.ard\.de)/(?:.*/)(?:[0-9]+|[^0-9][^/\?]+)[^/\?]*(?:\?.*)?)|(?:https?://(?:www\` +
			`.)?arte\.tv/(?:fr|de|en|es|it|pl)/videos/(?:\d{6}-\d{3}-[AF]))|(?:https?://(?:ww` +
			`w\.)?(?:(?:(?:asiancrush|yuyutv|midnightpulp)\.com|cocoro\.tv))/video/(?:[^/]+/)` +
			`?0+(?:\d+)v\b)|(?:https?://(?:www\.)?audiomack\.com/song/(?:[\w/-]+))|(?:https?:` +
			`//(?:www\.)?(?:telezueri\.ch|telebaern\.tv|telem1\.ch)/[^/]+/(?:[^/]+-(?:\d+))(?` +
			`:\#video=(?:[_0-9a-z]+))?)|(?:https?://[^/]+\.bandcamp\.com/track/(?:[^/?#&]+))|` +
			`(?:https?://(?:www\.)?bandcamp\.com/?\?(?:.*?&)?show=(?:\d+))|(?:https?://(?:www` +
			`\.)?beeg\.(?:com|porn(?:/video)?)/(?:\d+))|(?:https?://(?:www\.|pro\.)?beatport\` +
			`.com/track/(?:[^/]+)/(?:[0-9]+))|(?:https?://(?:www\.)?bigflix\.com/.+/(?:[0-9]+` +
			`))|(?:https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.b` +
			`ild\.html)|(?:https?://(?:www\.)?bitchute\.com/(?:video|embed|torrent/[^/]+)/(?:` +
			`[^/?#&]+))|(?:https?://(?:www\.)?biqle\.(?:com|org|ru)/watch/(?:-?\d+_\d+))|(?:h` +
			`ttps?://(?:www\.)?bleacherreport\.com/articles/(?:\d+))|(?:https?://(?:www\.)?bl` +
			`eacherreport\.com/video_embed\?id=(?:[0-9a-f-]{36}|\d{5}))|(?:https?://(?:www\.)` +
			`?bpb\.de/mediathek/(?:[0-9]+)/)|(?:(?:https?://(?:www\.)?br(?:-klassik)?\.de)/(?` +
			`:[a-z0-9\-_]+/)+(?:[a-z0-9\-_]+)\.html)|(?:https?://(?:www\.)?bravotv\.com/(?:[^` +
			`/]+/)+(?:[^/?#]+))|(?:(?:https?://.*brightcove\.com/(services|viewer).*?\?|brigh` +
			`tcove:)(?:.*))|(?:https?://(?:www\.)?camdemy\.com/media/(?:\d+))|(?:https?://(?:` +
			`www\.)?(?:mycanal|piwiplus)\.fr/(?:[^/]+/)*(?:[^?/]+)(?:\.html\?.*\bvid=|/p/)(?:` +
			`\d+))|(?:https?://(?:(?:www\.)?canalc2\.tv/video/|archives-canalc2\.u-strasbg\.f` +
			`r/video\.asp\?.*\bidVideo=)(?:\d+))|(?:(?:cbcplayer:|https?://(?:www\.)?cbc\.ca/` +
			`(?:player/play/|i/caffeine/syndicate/\?mediaId=))(?:\d+))|(?:https?://(?:www\.)?` +
			`media\.ccc\.de/v/(?:[^/?#&]+))|(?:https?://(?:www\.)?ccma\.cat/(?:[^/]+/)*?(?:vi` +
			`deo|audio)/(?:\d+))|(?:https?://(?:(?:[^/]+)\.(?:cntv|cctv)\.(?:com|cn)|(?:www\.` +
			`)?ncpa-classic\.com)/(?:[^/]+/)*?(?:[^/?#&]+?)(?:/index)?(?:\.s?html|[?#&]|$))|(` +
			`?:https?://(?:www\.)?charlierose\.com/(?:video|episode)(?:s|/player)/(?:\d+))|(?` +
			`:https?://(?:www\.)?chilloutzone\.net/video/(?:[\w|-]+)\.html)|(?:https?://(?:ww` +
			`w\.)?chirb\.it/(?:(?:wp|pl)/|fb_chirbit_player\.swf\?key=)?(?:[\da-zA-Z]+))|(?:h` +
			`ttps?://(?:www\.)?cinemax\.com/(?:[^/]+/video/[0-9a-z-]+-(?:\d+)))|(?:https?://(` +
			`?:www\.)?clippituser\.tv/c/(?:[a-z]+))|(?:https?://(?:www\.)?clip\.rs/(?:[^/]+)/` +
			`\d+)|(?:https?://(?:chic|www)\.clipsyndicate\.com/video/play(list/\d+)?/(?:\d+))` +
			`|(?:https?://(?:www\.)?closertotruth\.com/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(` +
			`?:www\.)?clyp\.it/(?:[a-z0-9]+))|(?:https?://(?:(?:edition|www|money)\.)?cnn\.co` +
			`m/(?:video/(?:data/.+?|\?)/)?videos?/(?:.+?/(?:[^/]+?)(?:\.(?:[a-z\-]+)|(?=&))))` +
			`|(?:https?://[^\.]+\.blogs\.cnn\.com/.+)|(?:https?://(?:(?:edition|www)\.)?cnn\.` +
			`com/(?!videos?/))|(?:(?:coub:|https?://(?:coub\.com/(?:view|embed|coubs)/|c-cdn\` +
			`.coub\.com/fb-player\.swf\?.*\bcoub(?:ID|id)=))(?:[\da-z]+))|(?:https?://(?:vide` +
			`o|www|player(?:-backend)?)\.(?:allure|architecturaldigest|arstechnica|bonappetit` +
			`|brides|cnevids|cntraveler|details|epicurious|glamour|golfdigest|gq|newyorker|se` +
			`lf|teenvogue|vanityfair|vogue|wired|wmagazine)\.com/(?:(?:embed(?:js)?|(?:script` +
			`|inline)/video)/(?:[0-9a-f]{24})(?:/(?:[0-9a-f]{24}))?(?:.+?\btarget=(?:[^&]+))?` +
			`|(?:watch|series|video)/(?:[^/?#]+)))|(?:https?://(?:www\.)?cracked\.com/video_(` +
			`?:\d+)_[\da-z-]+\.html)|(?:https?://news\.cts\.com\.tw/[a-z]+/[a-z]+/\d+/(?:\d+)` +
			`\.html)|(?:https?://(?:app\.)?curiositystream\.com/video/(?:\d+))|(?:https?://(?` +
			`:www\.)?dailymail\.co\.uk/(?:video/[^/]+/video-|embed/video/)(?:[0-9]+))|(?:(?i)` +
			`https?://(?:(www|touch)\.)?dailymotion\.[a-z]{2,3}/(?:(?:(?:embed|swf|#)/)?video` +
			`|swf)/(?:[^/?_]+))|(?:https?://(?:www\.)?dagbladet\.no/video/(?:(?:embed|(?:[^/]` +
			`+))/)?(?:[0-9A-Za-z_-]{11}|[a-zA-Z0-9]{8}))|(?:https?://(?:www\.)?digg\.com/vide` +
			`o/(?:[^/?#&]+))|(?:https?://(?:www\.)?dotsub\.com/view/(?:[^/]+))|(?:https?://(?` +
			`:(?:www|m)\.)?drtuber\.com/(?:video|embed)/(?:\d+)(?:/(?:[\w-]+))?)|(?:https?://` +
			`video\.aktualne\.cz/(?:[^/]+/)+r~(?:[0-9a-f]{32}))|(?:https?://(?:www\.)?(?:disc` +
			`overy|tlc|animalplanet|dmax)\.de/(?:.*\#(?:\d+)|(?:[^/]+/)*videos/(?:[^/?#]+)|pr` +
			`ogramme/(?:[^/]+)/video/(?:[^/]+)))|(?:https?://(?:(?:[^/]+\.)?(?:disney\.[a-z]{` +
			`2,3}(?:\.[a-z]{2})?|disney(?:(?:me|latino)\.com|turkiye\.com\.tr|channel\.de)|(?` +
			`:starwars|marvelkids)\.com))/(?:(?:embed/|(?:[^/]+/)+[\w-]+-)(?:[a-z0-9]{24})|(?` +
			`:[^/]+/)?(?:[^/?#]+)))|(?:https?://(?:s?evt\.dispeak|events\.digitallyspeaking)\` +
			`.com/(?:[^/]+/)+xml/(?:[^.]+)\.xml)|(?:https?://(?:www\.)?dw\.com/(?:[^/]+/)+(?:` +
			`av|e)-(?:\d+))|(?:(?:eagleplatform:(?:[^/]+):|https?://(?:.+?\.media\.eagleplatf` +
			`orm\.com)/index/player\?.*\brecord_id=)(?:\d+))|(?:https?://(?:www\.)?ebaumsworl` +
			`d\.com/videos/[^/]+/(?:\d+))|(?:https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))` +
			`|(?:https?://(?:[^.]+\.)?elpais\.com/.*/(?:[^/#?]+)\.html(?:$|[?#]))|(?:https?:/` +
			`/(?:www\.)?engadget\.com/video/(?:[^/?#]+))|(?:https?://(?:www\.)?eporner\.com/(` +
			`?:hd-porn|embed)/(?:\w+)(?:/(?:[\w-]+))?)|(?:https?://?(?:(?:www|v1)\.)?escapist` +
			`magazine\.com/videos/view/[^/]+/(?:[0-9]+))|(?:https?://(?:www\.)?extremetube\.c` +
			`om/(?:[^/]+/)?video/(?:[^/#?&]+))|(?:(?:https?://(?:[\w-]+\.)?(?:facebook\.com|f` +
			`acebookcorewwwi\.onion)/(?:[^#]*?\#!/)?(?:(?:video/video\.php|photo\.php|video\.` +
			`php|video/embed|story\.php)\?(?:.*?)(?:v|video_id|story_fbid)=|[^/]+/videos/(?:[` +
			`^/]+/)?|[^/]+/posts/|groups/[^/]+/permalink/)|facebook:)(?:[0-9]+))|(?:https?://` +
			`(?:[\w-]+\.)?facebook\.com/plugins/video\.php\?.*?\bhref=(?:https.+))|(?:https?:` +
			`//(?:www\.)?faz\.net/(?:[^/]+/)*.*?-(?:\d+)\.html)|(?:https?://(?:www\.)?fc-zeni` +
			`t\.ru/video/(?:[0-9]+))|(?:https?://(?:www\.)?filmweb\.no/(?:trailere|filmnytt)/` +
			`article(?:\d+)\.ece)|(?:(?:5min:|https?://(?:[^/]*?5min\.com/|delivery\.vidible\` +
			`.tv/aol)(?:(?:Scripts/PlayerSeed\.js|playerseed/?)?\?.*?playList=)?)(?:\d+))|(?:` +
			`https?://(?:www\.)?5-tv\.ru/(?:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?:htt` +
			`ps?://(?:www\.|secure\.)?flickr\.com/photos/[\w\-_@]+/(?:\d+))|(?:https?://(?:ww` +
			`w\.)?formula1\.com/(?:content/fom-website/)?en/video/\d{4}/\d{1,2}/(?:.+?)\.html` +
			`)|(?:https?://(?:video\.(?:insider\.)?fox(?:news|business)\.com)/v/(?:video-embe` +
			`d\.html\?video_id=)?(?:\d+))|(?:https?://(?:www\.)?(?:insider\.)?foxnews\.com/(?` +
			`!v)([^/]+/)+(?:[a-z-]+))|(?:https?://(?:www\.)?franceculture\.fr/emissions/(?:[^` +
			`/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?franceinter\.fr/emissions/(?:[^?#]+))|` +
			`(?:https?://generation-what\.francetv\.fr/[^/]+/video/(?:[^/?#&]+))|(?:https?://` +
			`(?:www\.)?freesound\.org/people/[^/]+/sounds/(?:[^/]+))|(?:https?://(?:www\.)?fu` +
			`nk\.net/(?:channel|playlist)/[^/]+/(?:[0-9a-z-]+)-(?:\d+))|(?:https?://(?:www\.)` +
			`?(?:fxnetworks|simpsonsworld)\.com/video/(?:\d+))|(?:https?://(?:www\.)?gameinfo` +
			`rmer\.com/(?:[^/]+/)*(?:[^.?&#]+))|(?:https?://(?:www\.)?gamespot\.com/(?:video|` +
			`article|review)s/(?:[^/]+/\d+-|embed/)(?:\d+))|(?:https?://(?:www\.)?gaskrank\.t` +
			`v/tv/(?:[^/]+)/(?:[^/]+)\.htm)|(?:(?:https?://(?:www\.)?gazeta\.ru/(?:[^/]+/)?vi` +
			`deo/(?:main/)*(?:\d{4}/\d{2}/\d{2}/)?(?:[A-Za-z0-9-_.]+)\.s?html))|(?:https?://(` +
			`?:www\.)?gdcvault\.com/play/(?:\d+)(?:/(?:[\w-]+))?)|(?:https?://(?:www\.)?gfyca` +
			`t\.com/(?:ru/|ifr/|gifs/detail/)?(?:[^-/?#]+))|(?:https?://(?:www\.)?godtube\.co` +
			`m/watch/\?v=(?:[\da-zA-Z]+))|(?:^https?://video\.golem\.de/.+?/(?:.+?)/)|(?:http` +
			`s?://(?:(?:docs|drive)\.google\.com/(?:(?:uc|open)\?.*?id=|file/d/)|video\.googl` +
			`e\.com/get_player\?.*?docid=)(?:[a-zA-Z0-9_-]{28,}))|(?:https?://on-demand\.gput` +
			`echconf\.com/gtc/2015/video/S(?:\d+)\.html)|(?:https?://(?:www\.)?hbo\.com/(?:vi` +
			`deo|embed)(?:/[^/]+)*/(?:[^/?#]+))|(?:https?://(?:www\.)?historicfilms\.com/(?:t` +
			`apes/|play)(?:\d+))|(?:https?://(?:www\.)?hitrecord\.org/records/(?:\d+))|(?:htt` +
			`p?://(?:www\.)?hornbunny\.com/videos/(?:[a-z-]+)-(?:\d+)\.html)|(?:https?://[\da` +
			`-z-]+\.(?:howstuffworks|stuff(?:(?:youshould|theydontwantyouto)know|toblowyourmi` +
			`nd|momnevertoldyou)|(?:brain|car)stuffshow|fwthinking|geniusstuff)\.com/(?:[^/]+` +
			`/)*(?:\d+-)?(?:.+?)-video\.htm)|(?:https?://(?:www\.)?hypem\.com/track/(?:[0-9a-` +
			`z]{5}))|(?:https?://.+?\.ign\.com/(?:[^/]+/)?(?:videos|show_videos|articles|feat` +
			`ure|(?:[^/]+/\d+/video))(/.+)?/(?:.+))|(?:https?://(?:www|m)\.imdb\.com/(?:video` +
			`|title|list).+?[/-]vi(?:\d+))|(?:https?://(?:i\.)?imgur\.com/(?!(?:a|gallery|(?:` +
			`t(?:opic)?|r)/[^/]+)/)(?:[a-zA-Z0-9]+))|(?:https?://(?:i\.)?imgur\.com/(?:galler` +
			`y|(?:t(?:opic)?|r)/[^/]+)/(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?ina\.fr/(?:vid` +
			`eo|audio)/(?:[A-Z0-9_]+))|(?:https?://(?:www\.)?infoq\.com/(?:[^/]+/)+(?:[^/]+))` +
			`|(?:(?:https?://(?:www\.)?instagram\.com/p/(?:[^/?#&]+)))|(?:https?://(?:www\.)?` +
			`90tv\.ir/video/(?:[0-9]+)/.*)|(?:https?://(?:www\.)?ivideon\.com/tv/(?:[^/]+/)*c` +
			`amera/(?:\d+-[\da-f]+)/(?:\d+))|(?:https?://(?:(?:www|m)\.)?izlesene\.com/(?:vid` +
			`eo|embedplayer)/(?:[^/]+/)?(?:[0-9]+))|(?:https?://(?:licensing\.jamendo\.com/[^` +
			`/]+|(?:www\.)?jamendo\.com)/track/(?:[0-9]+)/(?:[^/?#&]+))|(?:https?://.*?\.jeux` +
			`video\.com/.*/(.*?)\.htm)|(?:(?:joj:|https?://media\.joj\.sk/embed/)(?:[^/?#^]+)` +
			`)|(?:(?:https?://(?:content\.jwplatform|cdn\.jwplayer)\.com/(?:(?:feed|player|th` +
			`umb|preview|video)s|jw6|v2/media)/|jwplatform:)(?:[a-zA-Z0-9]{8}))|(?:https?://t` +
			`v\.kakao\.com/channel/(?:\d+)/cliplink/(?:\d+))|(?:(?:kaltura:(?:\d+):(?:[0-9a-z` +
			`_]+)|https?://(:?(?:www|cdnapi(?:sec)?)\.)?kaltura\.com(?::\d+)?/(?:(?:index\.ph` +
			`p/(?:kwidget|extwidget/preview)|html5/html5lib/[^/]+/mwEmbedFrame\.php))(?:/(?:[` +
			`^?]+))?(?:\?(?:.*))?))|(?:https?://(?:www\.)?keezmovies\.com/video/(?:(?:[^/]+)-` +
			`)?(?:\d+))|(?:^https?://(?:(?:www|api)\.)?khanacademy\.org/(?:[^/]+)/(?:[^/]+/){` +
			`,2}(?:[^?#/]+)(?:$|[?#]))|(?:https?://(?:www\.)?kickstarter\.com/projects/(?:[^/` +
			`]*)/.*)|(?:https?://(?:www\.)?lci\.fr/[^/]+/[\w-]+-(?:\d+)\.html)|(?:https?://(?` +
			`:www\.)?loc\.gov/(?:item/|today/cyberlc/feature_wdesc\.php\?.*\brec=)(?:[0-9a-z_` +
			`.]+))|(?:(?:https?://html5-player\.libsyn\.com/embed/episode/id/(?:[0-9]+)))|(?:` +
			`(?:limelight:media:|https?://(?:link\.videoplatform\.limelight\.com/media/|asset` +
			`s\.delvenetworks\.com/player/loader\.swf)\?.*?\bmediaId=)(?:[a-z0-9]{32}))|(?:ht` +
			`tps?://(?:new\.)?livestream\.com/(?:accounts/(?:\d+)|(?:[^/]+))/(?:events/(?:\d+` +
			`)|(?:[^/]+))(?:/videos/(?:\d+))?)|(?:https?://(?:www\.)?lovehomeporn\.com/video/` +
			`(?:\d+)(?:/(?:[^/?#&]+))?)|(?:https?://(?:www\.)?(?:lynda\.com|educourse\.ga)/(?` +
			`:(?:[^/]+/){2,3}(?:\d+)|player/embed)/(?:\d+))|(?:https?://my\.mail\.ru/music/so` +
			`ngs/[^/?#&]+-(?:[\da-f]+))|(?:(?i)https?://(?:www\.)?manyvids\.com/video/(?:\d+)` +
			`)|(?:https?://(?:www\.)?massengeschmack\.tv/play/(?:[^?&#]+))|(?:(?:mediaset:|ht` +
			`tps?://(?:(?:www|static3)\.)?mediasetplay\.mediaset\.it/(?:(?:video|on-demand)/(` +
			`?:[^/]+/)+[^/]+_|player/index\.html\?.*?\bprogramGuid=))(?:[0-9A-Z]{16}))|(?:htt` +
			`ps://player\.megaphone\.fm/(?:[A-Z0-9]+))|(?:https?://(?:www\.)?metacritic\.com/` +
			`.+?/trailers/(?:\d+))|(?:(?:mva:|https?://(?:mva\.microsoft|(?:www\.)?microsoftv` +
			`irtualacademy)\.com/[^/]+/training-courses/[^/?#&]+-)(?:\d+)(?::|\?l=)(?:[\da-zA` +
			`-Z]+_\d+))|(?:https?://(?:[\da-z_-]+\.)*(?:mlb)\.com/(?:(?:(?:[^/]+/)*c-|(?:shar` +
			`ed/video/embed/(?:embed|m-internal-embed)\.html|(?:[^/]+/)+(?:play|index)\.jsp|)` +
			`\?.*?\bcontent_id=)(?:\d+)))|(?:https?://(?:www\.)?mofosex\.com/videos/(?:\d+)/(` +
			`?:[^/?#&.]+)\.html)|(?:https?://myspace\.com/[^/]+/(?:video/[^/]+/(?:\d+)|music/` +
			`song/[^/?#&]+-(?:\d+)-\d+(?:[/?#&]|$)))|(?:(?:https?://(?:www\.)?myvi\.(?:(?:ru/` +
			`player|tv)/(?:(?:embed/html|flash|api/Video/Get)/|content/preloader\.swf\?.*\bid` +
			`=)|ru/watch/)|myvi:)(?:[\da-zA-Z_-]+))|(?:https?://video\.nationalgeographic\.co` +
			`m/.*?)|(?:https?://(?:m\.)?tv(?:cast)?\.naver\.com/v/(?:\d+))|(?:https?://(?:wat` +
			`ch\.|www\.)?nba\.com/(?:(?:[^/]+/)+(?:[^?]*?))/?(?:/index\.html)?(?:\?.*)?$)|(?:` +
			`https?://(?:www\.)?csnne\.com/video/(?:[0-9a-z-]+))|(?:https?://(?:www\.)?(?:nbc` +
			`news|today|msnbc)\.com/([^/]+/)*(?:.*-)?(?:[^/?]+))|(?:https?://(?:www\.)?nbcspo` +
			`rts\.com//?(?:[^/]+/)+(?:[0-9a-z-]+))|(?:https?://vplayer\.nbcsports\.com/(?:[^/` +
			`]+/)+(?:[0-9a-zA-Z_]+))|(?:https?://(?:www\.)?ndr\.de/(?:[^/]+/)*(?:[^/?#]+),[\d` +
			`a-z]+\.html)|(?:https?://(?:www\.)?ndr\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|ext` +
			`ernalPlayer)\.html)|(?:https?://(?:www\.)?n-joy\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:p` +
			`layer|externalPlayer)_[^/]+\.html)|(?:https?://(?:[^/]+\.)?ndtv\.com/(?:[^/]+/)*` +
			`videos?/?(?:[^/]+/)*[^/?^&]+-(?:\d+))|(?:https?://(?:www\.)?netzkino\.de/\#!/(?:` +
			`[^/]+)/(?:[^/]+))|(?:https?://(?:www\.)?newgrounds\.com/(?:audio/listen|portal/v` +
			`iew)/(?:[0-9]+))|(?:https?://(?:www\.)?nexttv\.com\.tw/(?:[^/]+/)+(?:\d+))|(?:(?` +
			`:https?://api\.nexx(?:\.cloud|cdn\.com)/v3/(?:\d+)/videos/byid/|nexx:(?:(?:\d+):` +
			`)?|https?://arc\.nexx\.cloud/api/video/)(?:\d+))|(?:https?://(?:www\.)?(?:nhl|wc` +
			`h2016)\.com/(?:[^/]+/)*c-(?:\d+))|(?:https?://(?:[^/]+\.)?noovo\.ca/videos/(?:[^` +
			`/]+/[^/?#&]+))|(?:https?://media\.cms\.nova\.cz/embed/(?:[^/?#&]+))|(?:https?://` +
			`(?:[^.]+\.)?(?:tv(?:noviny)?|tn|novaplus|vymena|fanda|krasna|doma|prask)\.nova\.` +
			`cz/(?:[^/]+/)+(?:[^/]+?)(?:\.html|/|$))|(?:https?://(?:(?:www|cn)\.)?nowness\.co` +
			`m/(?:story|(?:series|category)/[^/]+)/(?:[^/]+?)(?:$|[?#]))|(?:https?://(?:www\.` +
			`)?ntv\.ru/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:(?:www|m|mobile)\.)?(?:odnokla` +
			`ssniki|ok)\.ru/(?:video(?:embed)?/|web-api/video/moviePlayer/|live/|dk\?.*?st\.m` +
			`vId=)(?:[\d-]+))|(?:https?://(?:www\.)?onet\.tv/[a-z]/[a-z]+/(?:[0-9a-z-]+)/(?:[` +
			`0-9a-z]+))|(?:https?://(?:www\.)?onet\.tv/[a-z]/(?:[a-z]+)(?:[?#]|$))|(?:(?:ooya` +
			`la:|https?://.+?\.ooyala\.com/.*?(?:embedCode|ec)=)(?:.+?)(&|$))|(?:https?://(?:` +
			`(?:www\.)?(?:verystream\.com))/(?:stream|e)/(?:[a-zA-Z0-9-_]+))|(?:https?://(?:w` +
			`ww\.)?outsidetv\.com/(?:[^/]+/)*?play/[a-zA-Z0-9]{8}/\d+/\d+/(?:[a-zA-Z0-9]{8}))` +
			`|(?:https?://(?:(?:www\.)?packtpub\.com/mapt|subscription\.packtpub\.com)/video/` +
			`[^/]+/(?:\d+)/(?:[^/]+)/(?:[^/]+)(?:/(?:[^/?&#]+))?)|(?:https?://(?:(?:www\.)?pa` +
			`ndora\.tv/view/(?:[^/]+)/(?:\d+)|(?:.+?\.)?channel\.pandora\.tv/channel/video\.p` +
			`tv\?|m\.pandora\.tv/?\?))|(?:(?i)https?://(?:www\.)?parliamentlive\.tv/Event/Ind` +
			`ex/(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}))|(?:https?://(?:(` +
			`?:(?:video|www|player)\.pbs\.org|video\.aptv\.org|video\.gpb\.org|video\.mpbonli` +
			`ne\.org|video\.wnpt\.org|video\.wfsu\.org|video\.wsre\.org|video\.wtcitv\.org|vi` +
			`deo\.pba\.org|video\.alaskapublic\.org|video\.azpbs\.org|portal\.knme\.org|video` +
			`\.vegaspbs\.org|watch\.aetn\.org|video\.ket\.org|video\.wkno\.org|video\.lpb\.or` +
			`g|videos\.oeta\.tv|video\.optv\.org|watch\.wsiu\.org|video\.keet\.org|pbs\.kixe\` +
			`.org|video\.kpbs\.org|video\.kqed\.org|vids\.kvie\.org|video\.pbssocal\.org|vide` +
			`o\.valleypbs\.org|video\.cptv\.org|watch\.knpb\.org|video\.soptv\.org|video\.rmp` +
			`bs\.org|video\.kenw\.org|video\.kued\.org|video\.wyomingpbs\.org|video\.cpt12\.o` +
			`rg|video\.kbyueleven\.org|video\.thirteen\.org|video\.wgbh\.org|video\.wgby\.org` +
			`|watch\.njtvonline\.org|watch\.wliw\.org|video\.mpt\.tv|watch\.weta\.org|video\.` +
			`whyy\.org|video\.wlvt\.org|video\.wvpt\.net|video\.whut\.org|video\.wedu\.org|vi` +
			`deo\.wgcu\.org|video\.wpbt2\.org|video\.wucftv\.org|video\.wuft\.org|watch\.wxel` +
			`\.org|video\.wlrn\.org|video\.wusf\.usf\.edu|video\.scetv\.org|video\.unctv\.org` +
			`|video\.pbshawaii\.org|video\.idahoptv\.org|video\.ksps\.org|watch\.opb\.org|wat` +
			`ch\.nwptv\.org|video\.will\.illinois\.edu|video\.networkknowledge\.tv|video\.wtt` +
			`w\.com|video\.iptv\.org|video\.ninenet\.org|video\.wfwa\.org|video\.wfyi\.org|vi` +
			`deo\.mptv\.org|video\.wnin\.org|video\.wnit\.org|video\.wpt\.org|video\.wvut\.or` +
			`g|video\.weiu\.net|video\.wqpt\.org|video\.wycc\.org|video\.wipb\.org|video\.ind` +
			`ianapublicmedia\.org|watch\.cetconnect\.org|video\.thinktv\.org|video\.wbgu\.org` +
			`|video\.wgvu\.org|video\.netnebraska\.org|video\.pioneer\.org|watch\.sdpb\.org|v` +
			`ideo\.tpt\.org|watch\.ksmq\.org|watch\.kpts\.org|watch\.ktwu\.org|watch\.eastten` +
			`nesseepbs\.org|video\.wcte\.tv|video\.wljt\.org|video\.wosu\.org|video\.woub\.or` +
			`g|video\.wvpublic\.org|video\.wkyupbs\.org|video\.kera\.org|video\.mpbn\.net|vid` +
			`eo\.mountainlake\.org|video\.nhptv\.org|video\.vpt\.org|video\.witf\.org|watch\.` +
			`wqed\.org|video\.wmht\.org|video\.deltabroadcasting\.org|video\.dptv\.org|video\` +
			`.wcmu\.org|video\.wkar\.org|wnmuvideo\.nmu\.edu|video\.wdse\.org|video\.wgte\.or` +
			`g|video\.lptv\.org|video\.kmos\.org|watch\.montanapbs\.org|video\.krwg\.org|vide` +
			`o\.kacvtv\.org|video\.kcostv\.org|video\.wcny\.org|video\.wned\.org|watch\.wpbst` +
			`v\.org|video\.wskg\.org|video\.wxxi\.org|video\.wpsu\.org|on-demand\.wvia\.org|v` +
			`ideo\.wtvi\.org|video\.westernreservepublicmedia\.org|video\.ideastream\.org|vid` +
			`eo\.kcts9\.org|video\.basinpbs\.org|video\.houstonpbs\.org|video\.klrn\.org|vide` +
			`o\.klru\.tv|video\.wtjx\.org|video\.ideastations\.org|video\.kbtc\.org)/(?:(?:vi` +
			`r|port)alplayer|video)/(?:[0-9]+)(?:[?/]|$)|(?:www\.)?pbs\.org/(?:[^/]+/){1,5}(?` +
			`:[^/]+?)(?:\.html)?/?(?:$|[?\#])|(?:video|player)\.pbs\.org/(?:widget/)?partnerp` +
			`layer/(?:[^/]+)/))|(?:https?://(?:www\.)?pearvideo\.com/video_(?:\d+))|(?:https?` +
			`://(?:www\.)?people\.com/people/videos/0,,(?:\d+),00\.html)|(?:https?://player\.` +
			`performgroup\.com/eplayer(?:/eplayer\.html|\.js)#/?(?:[0-9a-f]{26})\.(?:[0-9a-z]` +
			`{26}))|(?:https?://(?:[a-z0-9]+\.)?photobucket\.com/.*(([\?\&]current=)|_)(?:.*)` +
			`\.(?:(flv)|(mp4)))|(?:https?://player\.piksel\.com/v/(?:[a-z0-9]+))|(?:https?://` +
			`(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+))/?(?:$|[?#]))|(?:https?://(?:www\.)?p` +
			`lays\.tv/(?:video|embeds)/(?:[0-9a-f]{18}))|(?:https?://(?:www\.)?playvid\.com/w` +
			`atch(\?v=|/)(?:.+?)(?:#|$))|(?:(?:https?)://(?:(?:[^.]+)\.podomatic\.com/entry|(` +
			`?:www\.)?podomatic\.com/podcasts/(?:[^/]+)/episodes)/(?:[^/?#&]+))|(?:https?://(` +
			`?:www\.)?pokemon\.com/[a-z]{2}(?:.*?play=(?:[a-z0-9]{32})|/(?:[^/]+/)+(?:[^/?#&]` +
			`+)))|(?:https?://(?:[a-zA-Z]+\.)?porn\.com/videos/(?:(?:[^/]+)-)?(?:\d+))|(?:htt` +
			`ps?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos/(?:\d+)(?:/(?:.+))?)|(?:https` +
			`?://(?:(?:[^/]+\.)?(?:pornhub\.(?:com|net))/(?:(?:view_video\.php|video/show)\?v` +
			`iewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(?:[\da-z]+))|(?:https?://(?:\w` +
			`+\.)?pornotube\.com/(?:[^?#]*?)/video/(?:[0-9]+))|(?:https?://(?:www\.)?presstv\` +
			`.ir/[^/]+/(?:\d+)/(?:\d+)/(?:\d+)/(?:\d+)/(?:[^/]+)?)|(?:https?://(?:.+?)\.(?:ra` +
			`dio\.(?:de|at|fr|pt|es|pl|it)|rad\.io))|(?:https?://(?:www\.)?radiojavan\.com/vi` +
			`deos/video/(?:[^/]+)/?)|(?:https?://[^/]+\.(?:rai\.(?:it|tv)|rainews\.it)/.+?-(?` +
			`:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12})(?:-.+?)?\.html)|(?:htt` +
			`ps?://(?:videos\.raywenderlich\.com/courses|(?:www\.)?raywenderlich\.com)/(?:[^/` +
			`]+)/lessons/(?:\d+))|(?:https?://(?:(?:www\.)?redtube\.com/|embed\.redtube\.com/` +
			`\?.*?\bid=)(?:[0-9]+))|(?:(?:rentv:|https?://(?:www\.)?ren\.tv/(?:player|video/e` +
			`pizod)/)(?:\d+))|(?:^https?://(?:www\.)?reverbnation\.com/.*?/song/(?:\d+).*?$)|` +
			`(?:https?://(?:www\.)?rockstargames\.com/videos(?:/video/|#?/?\?.*\bvideo=)(?:\d` +
			`+))|(?:https?://(?:www\.)?prehravac\.rozhlas\.cz/audio/(?:[0-9]+))|(?:https?://(` +
			`?:www\.)?rtp\.pt/play/p(?:[0-9]+)/(?:[^/?#]+)/?)|(?:rts:(?:\d+)|https?://(?:.+?\` +
			`.)?rts\.ch/(?:[^/]+/){2,}(?:[0-9]+)-(?:.+?)\.html)|(?:https?://(?:test)?player\.` +
			`(?:rutv\.ru|vgtrk\.com)/(?:flash\d+v/container\.swf\?id=|iframe/(?:swf|video|liv` +
			`e)/id/|index/iframe/cast_id/)(?:\d+))|(?:https?://(?:www\.)?(?:ruutu|supla)\.fi/` +
			`(?:video|supla)/(?:\d+))|(?:https?://(?:www\.)?(?:safaribooksonline|(?:learning\` +
			`.)?oreilly)\.com/(?:library/view/[^/]+/(?:[^/]+)/(?:[^/?\#&]+)\.html|videos/[^/]` +
			`+/[^/]+/(?:[^-]+-[^/?\#&]+)))|(?:https?://(?:(?:v2|www)\.)?videos\.sapo\.(?:pt|c` +
			`v|ao|mz|tl)/(?:[\da-zA-Z]{20}))|(?:https?://(?:www\.)?sbs\.com\.au/(?:ondemand|n` +
			`ews)/video/(?:single/)?(?:[0-9]+))|(?:https?://(?:www\.)?screencast\.com/t/(?:[a` +
			`-zA-Z0-9]+))|(?:https?://screencast-o-matic\.com/watch/(?:[0-9a-zA-Z]+))|(?:http` +
			`s?://(?:www\.)?senate\.gov/isvp/?\?(?:.+))|(?:https?://(?:www\.)?seznamzpravy\.c` +
			`z/iframe/player\?.*\bsrc=)|(?:https?://vivo\.sx/(?:[\da-z]{10}))|(?:https?://(?:` +
			`.*?\.)?video\.sina\.com\.cn/(?:(?:view/|.*\#)(?:\d+)|.+?/(?:[^/?#]+)(?:\.s?html)` +
			`|api/sinawebApi/outplay.php/(?:.+?)\.swf))|(?:https?://news\.sky\.com/video/[0-9` +
			`a-z-]+-(?:[0-9]+))|(?:https?://(?:www\.)?slideshare\.net/[^/]+?/(?:.+?)($|\?))|(` +
			`?:https?://slideslive\.com/(?:[0-9]+))|(?:https?://(?:www\.)?sonyliv\.com/detail` +
			`s/[^/]+/(?:\d+))|(?:^(?:https?://)?(?:(?:(?:www\.|m\.)?soundcloud\.com/(?!statio` +
			`ns/track)(?:[\w\d-]+)/(?!(?:tracks|albums|sets(?:/.+?)?|reposts|likes|spotlight)` +
			`/?(?:$|[?#]))(?:[\w\d-]+)/?(?:[^?]+?)?(?:[?].*)?$)|(?:api\.soundcloud\.com/track` +
			`s/(?:\d+)(?:/?\?secret_token=(?:[^&]+))?)|(?:(?:w|player|p.)\.soundcloud\.com/pl` +
			`ayer/?.*?url=.*)))|(?:https?://(?:www\.)?soundgasm\.net/u/(?:[0-9a-zA-Z_-]+)/(?:` +
			`[0-9a-zA-Z_-]+))|(?:https?://(?:[^/]+\.)?spankbang\.com/(?:[\da-z]+)/(?:video|pl` +
			`ay|embed)\b)|(?:https?://(?:www\.)?spiegel\.de/video/[^/]*-(?:[0-9]+)(?:-embed|-` +
			`iframe)?(?:\.html)?(?:#.*)?$)|(?:https?://(?:www\.)?stitcher\.com/podcast/(?:[^/` +
			`]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+)(?:[/#?&]|$))|(?:https?://(?:(?:www|play)\.)?(?` +
			`:srf|rts|rsi|rtr|swissinfo)\.ch/play/(?:tv|radio)/(?:[^/]+/(?:video|audio)/[^?]+` +
			`|popup(?:video|audio)player)\?id=(?:[0-9a-f\-]{36}|\d+))|(?:https?://openclassro` +
			`om\.stanford\.edu(?:/?|(/MainFolder/(?:HomePage|CoursePage|VideoPage)\.php([?]co` +
			`urse=(?:[^&]+)(&video=(?:[^&]+))?(&.*)?)?))$)|(?:https?://streamable\.com/(?:[es` +
			`]/)?(?:\w+))|(?:https?://(?:(?:www\.)?sunporno\.com/videos|embeds\.sunporno\.com` +
			`/embed)/(?:\d+))|(?:https?://(?:www\.)?sverigesradio\.se/(?:sida/)?avsnitt/(?:[0` +
			`-9]+))|(?:https?://(?:www\.)?sverigesradio\.se/sida/(?:artikel|gruppsida)\.aspx\` +
			`?.*?\bartikel=(?:[0-9]+))|(?:https?://(?:www\.)?tagesschau\.de/(?:[^/]+/(?:[^/]+` +
			`/)*?(?:[^/#?]+?(?:-?[0-9]+)?))(?:~_?[^/#?]+?)?\.html)|(?:https?://(?:www\.)?tast` +
			`ytrade\.com/tt/shows/[^/]+/episodes/(?:[^/?#&]+))|(?:https?://tds\.lifeway\.com/` +
			`v1/trainingdeliverysystem/courses/(?:\d+)/index\.html)|(?:https?://(?:www\.)?tea` +
			`chertube\.com/(viewVideo\.php\?video_id=|music\.php\?music_id=|video/(?:[\da-z-]` +
			`+-)?|audio/)(?:\d+))|(?:https?://(?:\w+\.)?teamcoco\.com/(?:([^/]+/)*[^/?#]+))|(` +
			`?:https?://(?:www\.)?teamtreehouse\.com/library/(?:[^/]+))|(?:(?:https?://)(?:ww` +
			`w|embed(?:-ssl)?)(?:\.ted\.com/((?:playlists(?:/(?:\d+))?)|((?:talks))|(?:watch)` +
			`/[^/]+/[^/]+)(/lang/(.*?))?/(?:[\w-]+).*)$)|(?:^https?://(?:www\.)?t13\.cl/video` +
			`s(?:/[^/]+)+/(?:[\w-]+))|(?:https?://(?:www\.)?tenta\.com/how-to-download-videos` +
			`)|(?:https?://(?:www\.)?tfo\.org/(?:en|fr)/(?:[^/]+/){2}(?:\d+))|(?:(?:https?://` +
			`(?:link|player)\.theplatform\.com/[sp]/(?:[^/]+)/(?:(?:(?:[^/]+/)+select/)?(?:me` +
			`dia/(?:guid/\d+/)?)?|(?:(?:[^/\?]+/(?:swf|config)|onsite)/select/))?|theplatform` +
			`:)(?:[^/\?&]+))|(?:https?://thescene\.com/watch/[^/]+/(?:[^/#?]+))|(?:https?://(` +
			`?:www\.)?thisoldhouse\.com/(?:watch|how-to|tv-episode)/(?:[^/?#]+))|(?:https?://` +
			`(?:.+?\.)?tinypic\.com/player\.php\?v=(?:[^&]+)&s=\d+)|(?:https?://(?:www\.)?tmz` +
			`\.com/videos/(?:[^/?#]+))|(?:https?://player\.(?:tna|emp)flix\.com/video/(?:\d+)` +
			`)|(?:https?://(?:www\.)?tnaflix\.com/[^/]+/(?:[^/]+)/video(?:\d+))|(?:https?://(` +
			`?:www\.)?empflix\.com/(?:videos/(?:.+?)-|[^/]+/(?:[^/]+)/video)(?:[0-9]+))|(?:ht` +
			`tps?://(?:www\.)?moviefap\.com/videos/(?:[0-9a-f]+)/(?:[^/]+)\.html)|(?:https?:/` +
			`/(?:www\.)?toongoggles\.com/shows/(?:\d+)(?:/[^/]+/episodes/(?:\d+))?)|(?:https?` +
			`://videos\.toypics\.net/view/(?:[0-9]+))|(?:https?://(?:(?:www|m)\.)?trilulilu\.` +
			`ro/(?:[^/]+/)?(?:[^/#\?]+))|(?:https?://(?:www\.)?tube8\.com/(?:[^/]+/)+(?:[^/]+` +
			`)/(?:\d+))|(?:https?://(?:[^/?#&]+)\.tumblr\.com/(?:post|video)/(?:[0-9]+)(?:$|[` +
			`/?#]))|(?:https?://(?:(?:www\.)?tune\.pk/(?:video/|player/embed_player.php?.*?\b` +
			`vid=)|embed\.tune\.pk/play/)(?:\d+))|(?:https?://(?:www\.)?tvanouvelles\.ca/vide` +
			`os/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/video/iframe/id/(?:\d+))|(?:https?://(` +
			`?:www\.)?tvc\.ru/(?!video/iframe/id/)(?:[^?#]+))|(?:https?://(?:(?:[^/]+)\.)?tvn` +
			`24(?:bis)?\.pl/(?:[^/]+/)*(?:[^/]+))|(?:(?:tvp:|https?://[^/]+\.tvp\.(?:pl|info)` +
			`/sess/tvplayer\.php\?.*?object_id=)(?:\d+))|(?:https?://[^/]+\.tvp\.(?:pl|info)/` +
			`(?:video/(?:[^,\s]*,)*|(?:(?!\d+/)[^/]+/)*)(?:\d+))|(?:https?://(?:www\.)?20min\` +
			`.ch/(?:videotv/*\?.*?\bvid=|videoplayer/videoplayer\.html\?.*?\bvideoId@)(?:\d+)` +
			`)|(?:https?://video\.(?:twentythree\.net|23video\.com|filmweb\.no)/v\.ihtml/play` +
			`er\.html\?(?:.*?\bphoto(?:_|%5f)id=(?:\d+).*))|(?:https?://(?:clips\.twitch\.tv/` +
			`(?:[^/]+/)*|(?:www\.)?twitch\.tv/[^/]+/clip/)(?:[^/?#&]+))|(?:https?://(?:www\.)` +
			`?twitter\.com/i/(?:cards/tfw/v1|videos(?:/tweet)?)/(?:\d+))|(?:https?://amp\.twi` +
			`mg\.com/v/(?:[0-9a-f\-]{36}))|(?:https?://video\.udn\.com/(?:embed|play)/news/(?` +
			`:\d+))|(?:https?://(?:www\.)?(?:digiteka\.net|ultimedia\.com)/(?:deliver/(?:gene` +
			`ric|musique)(?:/[^/]+)*/(?:src|article)|default/index/video(?:generic|music)/id)` +
			`/(?:[\d+a-z]+))|(?:https?://utv\.unistra\.fr/(?:index|video)\.php\?id_video\=(?:` +
			`\d+))|(?:https?://(?:(?:app|embed)\.)?ustudio\.com/embed/(?:[^/]+)/(?:[^/]+))|(?` +
			`:https?://(?:www\.)?video\.varzesh3\.com/(?:[^/]+/)+(?:[^/]+)/?)|(?:https?://(?:` +
			`[^/]+\.)?vbox7\.com/(?:play:|(?:emb/external\.php|player/ext\.swf)\?.*?\bvid=)(?` +
			`:[\da-fA-F]+))|(?:https?://veehd\.com/video/(?:\d+))|(?:https?://(?:www\.)?veoh\` +
			`.com/(?:watch|embed|iphone/#_Watch)/(?:(?:v|e|yapi-)[\da-zA-Z]+))|(?:https?://(?` +
			`:.+?\.)?vesti\.ru/(?:.+))|(?:(?:https?://(?:www\.)?vevo\.com/watch/(?!playlist|g` +
			`enre)(?:[^/]+/(?:[^/]+/)?)?|https?://cache\.vevo\.com/m/html/embed\.html\?video=` +
			`|https?://videoplayer\.vevo\.com/embed/embedded\?videoId=|https?://embed\.vevo\.` +
			`com/.*?[?&]isrc=|vevo:)(?:[^&?#]+))|(?:(?:https?://(?:www\.)?(?:vgtv.no|bt.no/tv` +
			`|aftenbladet.no/tv|fvn.no/fvntv|aftenposten.no/webtv|ap.vgtv.no/webtv|tv.aftonbl` +
			`adet.se/abtv|www.aftonbladet.se/tv)/?(?:(?:\#!/)?(?:video|live)/|embed?.*id=|a(?` +
			`:rticles)?/)|(?:vgtv|bttv|satv|fvntv|aptv|abtv):)(?:\d+))|(?:https://www\.vice\.` +
			`com/[^/]+/article/(?:[^?#]+))|(?:https?://(?:www\.)?viddler\.com/(?:v|embed|play` +
			`er)/(?:[a-z0-9]+)(?:.+?\bsecret=(\d+))?)|(?:https?://videa(?:kid)?\.hu/(?:videok` +
			`/(?:[^/]+/)*[^?#&]+-|player\?.*?\bv=|player/v/)(?:[^?#&]+))|(?:https?://videopre` +
			`ss\.com/embed/(?:[\da-zA-Z]+))|(?:https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-` +
			`z]+/(?:[0-9]+)_)|(?:https?://(?:www\.)?vidlii\.com/(?:watch|embed)\?.*?\bv=(?:[0` +
			`-9A-Za-z_-]{11}))|(?:https?://(?:(?:www|(?:player))\.)?vimeo(?:pro)?\.com/(?!(?:` +
			`channels|album|showcase)/[^/?#]+/?(?:$|[?#])|[^/]+/review/|ondemand/)(?:.*?/)?(?` +
			`:(?:play_redirect_hls|moogaloop\.swf)\?clip_id=)?(?:videos?/)?(?:[0-9]+)(?:/[\da` +
			`-f]+)?/?(?:[?&].*)?(?:[#].*)?$)|(?:https?://(?:www\.)?vimeo\.com/ondemand/(?:[^/` +
			`?#&]+))|(?:(?:https://vimeo\.com/[^/]+/review/(?:[^/]+)/[0-9a-f]{10}))|(?:https?` +
			`://(?:player\.vimple\.(?:ru|co)/iframe|vimple\.(?:ru|co))/(?:[\da-f-]{32,36}))|(` +
			`?:https?://(?:www\.)?vine\.co/(?:v|oembed)/(?:\w+))|(?:(?:viqeo:|https?://cdn\.v` +
			`iqeo\.tv/embed/*\?.*?\bvid=|https?://api\.viqeo\.tv/v\d+/data/startup?.*?\bvideo` +
			`(?:%5B%5D|\[\])=)(?:[\da-f]+))|(?:https?://(?:(?:(?:(?:m|new)\.)?vk\.com/video_|` +
			`(?:www\.)?daxab.com/)ext\.php\?(?:.*?\boid=(?:-?\d+).*?\bid=(?:\d+).*)|(?:(?:(?:` +
			`m|new)\.)?vk\.com/(?:.+?\?.*?z=)?video|(?:www\.)?daxab.com/embed/)(?:-?\d+_\d+)(` +
			`?:.*\blist=(?:[\da-f]+))?))|(?:https?://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]` +
			`+))|(?:https?://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+)/playlist/(?:[0-9]+))|` +
			`(?:https?://(?:m\.)?vuclip\.com/w\?.*?cid=(?:[0-9]+))|(?:https?://(?:(?:www|view` +
			`)\.)?vzaar\.com/(?:videos/)?(?:\d+))|(?:(?:washingtonpost:|https?://(?:www\.)?wa` +
			`shingtonpost\.com/video/(?:[^/]+/)*)(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{` +
			`4}-[\da-f]{12}))|(?:(?:wat:|https?://(?:www\.)?wat\.tv/video/.*-)(?:[0-9a-z]+))|` +
			`(?:https?://m\.weibo\.cn/status/(?:[0-9]+)(\?.+)?)|(?:https?://(?:www|m)\.worlds` +
			`tar(?:candy|hiphop)\.com/(?:videos|android)/video\.php\?.*?\bv=(?:[^&]+))|(?:(?:` +
			`https?://video-api\.wsj\.com/api-video/player/iframe\.html\?.*?\bguid=|https?://` +
			`(?:www\.)?(?:wsj|barrons)\.com/video/(?:[^/]+/)+|wsj:)(?:[a-fA-F0-9-]{36}))|(?:(` +
			`?i)https?://(?:www\.)?wsj\.com/articles/(?:[^/?#&]+))|(?:https?://(?:.+?\.)?xham` +
			`ster\.(?:com|one)/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(` +
			`?:https?://(?:.+?\.)?xhamster\.com/xembed\.php\?video=(?:\d+))|(?:https?://(?:vi` +
			`deo|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)|(?:(?:xtube:|https?://(?:www\.)?xtube` +
			`\.com/(?:watch\.php\?.*\bv=|video-watch/(?:embedded/)?(?:[^/]+)-))(?:[^/?&#]+))|` +
			`(?:https?://(?:(?:www\.)?xvideos\.com/video|flashservice\.xvideos\.com/embedfram` +
			`e/|static-hw\.xvideos\.com/swf/xv-player\.swf\?.*?\bid_video=)(?:[0-9]+))|(?:htt` +
			`ps?://(?:www\.)?xxxymovies\.com/videos/(?:\d+)/(?:[^/]+))|(?:(?:https?://(?:(?:[` +
			`a-zA-Z]{2})\.)?[\da-zA-Z_-]+\.yahoo\.com)/(?:[^/]+/)*(?:(?:.+)?-)?(?:[0-9]+)(?:-` +
			`[a-z]+)?(?:\.html)?)|(?:https?://v\.yinyuetai\.com/video(?:/h5)?/(?:[0-9]+))|(?:` +
			`https?://(?:\w+\.)?youjizz\.com/videos/(?:[^/#?]*-(?:\d+)\.html|embed/(?:\d+)))|` +
			`(?:https?://(?:www\.)?youporn\.com/watch/(?:\d+)/(?:[^/?#&]+))|(?:https?://(?:ww` +
			`w\.)?(?:yourupload\.com/(?:watch|embed)|embed\.yourupload\.com)/(?:[A-Za-z0-9]+)` +
			`)|(?:^((?:https?://|//)(?:(?:(?:(?:\w+\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocook` +
			`ie)?\.com/|(?:www\.)?deturl\.com/www\.youtube\.com/|(?:www\.)?pwnyoutube\.com/|(` +
			`?:www\.)?hooktube\.com/|(?:www\.)?yourepeat\.com/|tube\.majestyc\.net/|(?:(?:www` +
			`|dev)\.)?invidio\.us/|(?:(?:www|no)\.)?invidiou\.sh/|(?:(?:www|fi|de)\.)?invidio` +
			`us\.snopyta\.org/|(?:www\.)?invidious\.kabi\.tk/|(?:www\.)?invidious\.enkirton\.` +
			`net/|(?:www\.)?invidious\.13ad\.de/|(?:www\.)?invidious\.mastodon\.host/|(?:www\` +
			`.)?tube\.poal\.co/|(?:www\.)?vid\.wxzm\.sx/|youtube\.googleapis\.com/)(?:.*?\#/)` +
			`?(?:(?:(?:v|embed|e)/(?!videoseries))|(?:(?:(?:watch|movie)(?:_popup)?(?:\.php)?` +
			`/?)?(?:\?|\#!?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.plus|zwearz\.com/watch|)/|(?` +
			`:www\.)?cleanvideosearch\.com/media/action/yt/watch\?videoId=))?([0-9A-Za-z_-]{1` +
			`1})(?!.*?\blist=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,}|WL)))|` +
			`(?:(?:(?:https?://)?(?:\w+\.)?(?:(?:youtube\.com|invidio\.us)/(?:(?:course|view_` +
			`play_list|my_playlists|artist|playlist|watch|embed/(?:videoseries|[0-9A-Za-z_-]{` +
			`11}))\?(?:.*?[&;])*?(?:p|a|list)=|p/)|youtu\.be/[0-9A-Za-z_-]{11}\?.*?\blist=)((` +
			`?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)?[0-9A-Za-z-_]{10,}|(?:MC)[\w\.]*).*|((?:PL|L` +
			`L|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,})))`,
	}
}
