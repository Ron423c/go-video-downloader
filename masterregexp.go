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
		`(?:https?://abcnews\.go\.com/(?:[^/]+/video/(?:[0-9a-z-]+)-|video/embed\?.*?\bid` +
			`=)(?:\d+))|(?:https?://(?:(?:(?:embed|www)\.)?acast\.com/|play\.acast\.com/s/)(?` +
			`:[^/]+)/(?:[^/#?]+))|(?:https?://tv\.adobe\.com/(?:(?:fr|de|es|jp)/)?watch/(?:[^` +
			`/]+)/(?:[^/]+))|(?:https?://video\.tv\.adobe\.com/v/(?:\d+))|(?:(?:aol-video:|ht` +
			`tps?://(?:www\.)?aol\.(?:com|ca|co\.uk|de|jp)/video/(?:[^/]+/)*)(?:[0-9a-f]+))|(` +
			`?:https?://(?:www\.)?allocine\.fr/(?:article|video|film)/(?:fichearticle_gen_car` +
			`ticle=|player_gen_cmedia=|fichefilm_gen_cfilm=|video-)(?:[0-9]+)(?:\.html)?)|(?:` +
			`https?://(?:www\.)?aparat\.com/(?:v/|video/video/embed/videohash/)(?:[a-zA-Z0-9]` +
			`+))|(?:https?://(?:www\.)?archive\.org/(?:details|embed)/(?:[^/?#]+)(?:[?].*)?$)` +
			`|(?:^https?://(?:(?:(?:www|classic)\.)?ardmediathek\.de|mediathek\.(?:daserste|r` +
			`bb-online)\.de|one\.ard\.de)/(?:.*/)(?:[0-9]+|[^0-9][^/\?]+)[^/\?]*(?:\?.*)?)|(?` +
			`:https?://(?:www\.)?asiancrush\.com/video/(?:[^/]+/)?0+(?:\d+)v\b)|(?:https?://(` +
			`?:www\.)?audiomack\.com/song/(?:[\w/-]+))|(?:https?://(?:www\.)?(?:telezueri\.ch` +
			`|telebaern\.tv|telem1\.ch)/[^/]+/(?:[^/]+-(?:\d+))(?:\#video=(?:[_0-9a-z]+))?)|(` +
			`?:https?://[^/]+\.bandcamp\.com/track/(?:[^/?#&]+))|(?:https?://(?:www\.)?bandca` +
			`mp\.com/?\?(?:.*?&)?show=(?:\d+))|(?:https?://(?:www\.)?beeg\.(?:com|porn(?:/vid` +
			`eo)?)/(?:\d+))|(?:https?://(?:www\.|pro\.)?beatport\.com/track/(?:[^/]+)/(?:[0-9` +
			`]+))|(?:https?://(?:www\.)?bigflix\.com/.+/(?:[0-9]+))|(?:https?://(?:www\.)?bil` +
			`d\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?:https?://(?:ww` +
			`w\.)?bitchute\.com/(?:video|embed|torrent/[^/]+)/(?:[^/?#&]+))|(?:https?://(?:ww` +
			`w\.)?biqle\.(?:com|org|ru)/watch/(?:-?\d+_\d+))|(?:https?://(?:www\.)?bleacherre` +
			`port\.com/articles/(?:\d+))|(?:https?://(?:www\.)?bleacherreport\.com/video_embe` +
			`d\?id=(?:[0-9a-f-]{36}))|(?:https?://(?:www\.)?bpb\.de/mediathek/(?:[0-9]+)/)|(?` +
			`:(?:https?://(?:www\.)?br(?:-klassik)?\.de)/(?:[a-z0-9\-_]+/)+(?:[a-z0-9\-_]+)\.` +
			`html)|(?:https?://(?:www\.)?bravotv\.com/(?:[^/]+/)+(?:[^/?#]+))|(?:(?:https?://` +
			`.*brightcove\.com/(services|viewer).*?\?|brightcove:)(?:.*))|(?:https?://(?:www\` +
			`.)?byutv\.org/(?:watch|player)/(?!event/)(?:[0-9a-f-]+)(?:/(?:[^/?#&]+))?)|(?:ht` +
			`tps?://(?:www\.)?camdemy\.com/media/(?:\d+))|(?:https?://(?:www\.)?(?:mycanal|pi` +
			`wiplus)\.fr/(?:[^/]+/)*(?:[^?/]+)(?:\.html\?.*\bvid=|/p/)(?:\d+))|(?:https?://(?` +
			`:(?:www\.)?canalc2\.tv/video/|archives-canalc2\.u-strasbg\.fr/video\.asp\?.*\bid` +
			`Video=)(?:\d+))|(?:(?:cbcplayer:|https?://(?:www\.)?cbc\.ca/(?:player/play/|i/ca` +
			`ffeine/syndicate/\?mediaId=))(?:\d+))|(?:https?://(?:www\.)?media\.ccc\.de/v/(?:` +
			`[^/?#&]+))|(?:https?://(?:www\.)?ccma\.cat/(?:[^/]+/)*?(?:video|audio)/(?:\d+))|` +
			`(?:https?://(?:(?:[^/]+)\.(?:cntv|cctv)\.(?:com|cn)|(?:www\.)?ncpa-classic\.com)` +
			`/(?:[^/]+/)*?(?:[^/?#&]+?)(?:/index)?(?:\.s?html|[?#&]|$))|(?:https?://(?:www\.)` +
			`?charlierose\.com/(?:video|episode)(?:s|/player)/(?:\d+))|(?:https?://(?:www\.)?` +
			`chirb\.it/(?:(?:wp|pl)/|fb_chirbit_player\.swf\?key=)?(?:[\da-zA-Z]+))|(?:https?` +
			`://player\.cinchcast\.com/.*?(?:assetId|show_id)=(?:[0-9]+))|(?:https?://(?:www\` +
			`.)?cinemax\.com/(?:[^/]+/video/[0-9a-z-]+-(?:\d+)))|(?:https?://(?:www\.)?ciscol` +
			`ive(?:\.cisco)?\.com/[^#]*#/session/(?:[^/?&]+))|(?:https?://(?:www\.)?clippitus` +
			`er\.tv/c/(?:[a-z]+))|(?:https?://(?:www\.)?clip\.rs/(?:[^/]+)/\d+)|(?:https?://(` +
			`?:chic|www)\.clipsyndicate\.com/video/play(list/\d+)?/(?:\d+))|(?:https?://(?:ww` +
			`w\.)?closertotruth\.com/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?clyp\.it/` +
			`(?:[a-z0-9]+))|(?:https?://(?:(?:edition|www|money)\.)?cnn\.com/(?:video/(?:data` +
			`/.+?|\?)/)?videos?/(?:.+?/(?:[^/]+?)(?:\.(?:[a-z\-]+)|(?=&))))|(?:https?://[^\.]` +
			`+\.blogs\.cnn\.com/.+)|(?:https?://(?:(?:edition|www)\.)?cnn\.com/(?!videos?/))|` +
			`(?:(?:coub:|https?://(?:coub\.com/(?:view|embed|coubs)/|c-cdn\.coub\.com/fb-play` +
			`er\.swf\?.*\bcoub(?:ID|id)=))(?:[\da-z]+))|(?:https?://(?:video|www|player(?:-ba` +
			`ckend)?)\.(?:allure|architecturaldigest|arstechnica|bonappetit|brides|cnevids|cn` +
			`traveler|details|epicurious|glamour|golfdigest|gq|newyorker|self|teenvogue|vanit` +
			`yfair|vogue|wired|wmagazine)\.com/(?:(?:embed(?:js)?|(?:script|inline)/video)/(?` +
			`:[0-9a-f]{24})(?:/(?:[0-9a-f]{24}))?(?:.+?\btarget=(?:[^&]+))?|(?:watch|series|v` +
			`ideo)/(?:[^/?#]+)))|(?:https?://(?:www\.)?cracked\.com/video_(?:\d+)_[\da-z-]+\.` +
			`html)|(?:https?://(?:www\.)?c-span\.org/video/\?(?:[0-9a-f]+))|(?:https?://news\` +
			`.cts\.com\.tw/[a-z]+/[a-z]+/\d+/(?:\d+)\.html)|(?:https?://(?:app\.)?curiosityst` +
			`ream\.com/video/(?:\d+))|(?:https?://(?:www\.)?dailymail\.co\.uk/(?:video/[^/]+/` +
			`video-|embed/video/)(?:[0-9]+))|(?:(?i)https?://(?:(www|touch)\.)?dailymotion\.[` +
			`a-z]{2,3}/(?:(?:(?:embed|swf|#)/)?video|swf)/(?:[^/?_]+))|(?:https?://(?:www\.)?` +
			`digg\.com/video/(?:[^/?#&]+))|(?:https?://(?:www\.)?dotsub\.com/view/(?:[^/]+))|` +
			`(?:https?://(?:(?:www|m)\.)?drtuber\.com/(?:video|embed)/(?:\d+)(?:/(?:[\w-]+))?` +
			`)|(?:https?://video\.aktualne\.cz/(?:[^/]+/)+r~(?:[0-9a-f]{32}))|(?:https?://(?:` +
			`www\.)?(?:discovery|tlc|animalplanet|dmax)\.de/(?:.*\#(?:\d+)|(?:[^/]+/)*videos/` +
			`(?:[^/?#]+)|programme/(?:[^/]+)/video/(?:[^/]+)))|(?:https?://(?:(?:[^/]+\.)?(?:` +
			`disney\.[a-z]{2,3}(?:\.[a-z]{2})?|disney(?:(?:me|latino)\.com|turkiye\.com\.tr|c` +
			`hannel\.de)|(?:starwars|marvelkids)\.com))/(?:(?:embed/|(?:[^/]+/)+[\w-]+-)(?:[a` +
			`-z0-9]{24})|(?:[^/]+/)?(?:[^/?#]+)))|(?:https?://(?:s?evt\.dispeak|events\.digit` +
			`allyspeaking)\.com/(?:[^/]+/)+xml/(?:[^.]+)\.xml)|(?:https?://(?:www\.)?dw\.com/` +
			`(?:[^/]+/)+(?:av|e)-(?:\d+))|(?:(?:eagleplatform:(?:[^/]+):|https?://(?:.+?\.med` +
			`ia\.eagleplatform\.com)/index/player\?.*\brecord_id=)(?:\d+))|(?:https?://(?:www` +
			`\.)?ebaumsworld\.com/videos/[^/]+/(?:\d+))|(?:https?://(?:www\.)?echo\.msk\.ru/s` +
			`ounds/(?:\d+))|(?:https?://(?:[^.]+\.)?elpais\.com/.*/(?:[^/#?]+)\.html(?:$|[?#]` +
			`))|(?:https?://(?:www\.)?engadget\.com/video/(?:[^/?#]+))|(?:https?://(?:www\.)?` +
			`eporner\.com/(?:hd-porn|embed)/(?:\w+)(?:/(?:[\w-]+))?)|(?:https?://?(?:(?:www|v` +
			`1)\.)?escapistmagazine\.com/videos/view/[^/]+/(?:[0-9]+))|(?:https?://(?:www\.)?` +
			`extremetube\.com/(?:[^/]+/)?video/(?:[^/#?&]+))|(?:(?:https?://(?:[\w-]+\.)?(?:f` +
			`acebook\.com|facebookcorewwwi\.onion)/(?:[^#]*?\#!/)?(?:(?:video/video\.php|phot` +
			`o\.php|video\.php|video/embed|story\.php)\?(?:.*?)(?:v|video_id|story_fbid)=|[^/` +
			`]+/videos/(?:[^/]+/)?|[^/]+/posts/|groups/[^/]+/permalink/)|facebook:)(?:[0-9]+)` +
			`)|(?:https?://(?:[\w-]+\.)?facebook\.com/plugins/video\.php\?.*?\bhref=(?:https.` +
			`+))|(?:https?://(?:www\.)?faz\.net/(?:[^/]+/)*.*?-(?:\d+)\.html)|(?:https?://(?:` +
			`www\.)?fc-zenit\.ru/video/(?:[0-9]+))|(?:https?://(?:www\.)?filmweb\.no/(?:trail` +
			`ere|filmnytt)/article(?:\d+)\.ece)|(?:(?:5min:|https?://(?:[^/]*?5min\.com/|deli` +
			`very\.vidible\.tv/aol)(?:(?:Scripts/PlayerSeed\.js|playerseed/?)?\?.*?playList=)` +
			`?)(?:\d+))|(?:https?://(?:www\.|secure\.)?flickr\.com/photos/[\w\-_@]+/(?:\d+))|` +
			`(?:https?://(?:www\.)?formula1\.com/(?:content/fom-website/)?en/video/\d{4}/\d{1` +
			`,2}/(?:.+?)\.html)|(?:https?://(?:video\.(?:insider\.)?fox(?:news|business)\.com` +
			`)/v/(?:video-embed\.html\?video_id=)?(?:\d+))|(?:https?://(?:www\.)?(?:insider\.` +
			`)?foxnews\.com/(?!v)([^/]+/)+(?:[a-z-]+))|(?:https?://(?:www\.)?franceculture\.f` +
			`r/emissions/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?franceinter\.fr/emiss` +
			`ions/(?:[^?#]+))|(?:https?://generation-what\.francetv\.fr/[^/]+/video/(?:[^/?#&` +
			`]+))|(?:https?://(?:www\.)?freesound\.org/people/[^/]+/sounds/(?:[^/]+))|(?:http` +
			`s?://(?:www\.)?(?:fxnetworks|simpsonsworld)\.com/video/(?:\d+))|(?:https?://(?:w` +
			`ww\.)?gamespot\.com/(?:video|article|review)s/(?:[^/]+/\d+-|embed/)(?:\d+))|(?:h` +
			`ttps?://(?:www\.)?gaskrank\.tv/tv/(?:[^/]+)/(?:[^/]+)\.htm)|(?:(?:https?://(?:ww` +
			`w\.)?gazeta\.ru/(?:[^/]+/)?video/(?:main/)*(?:\d{4}/\d{2}/\d{2}/)?(?:[A-Za-z0-9-` +
			`_.]+)\.s?html))|(?:https?://(?:www\.)?gdcvault\.com/play/(?:\d+)(?:/(?:[\w-]+))?` +
			`)|(?:https?://(?:www\.)?gfycat\.com/(?:ifr/|gifs/detail/)?(?:[^-/?#]+))|(?:https` +
			`?://(?:www\.)?godtube\.com/watch/\?v=(?:[\da-zA-Z]+))|(?:^https?://video\.golem\` +
			`.de/.+?/(?:.+?)/)|(?:https?://(?:(?:docs|drive)\.google\.com/(?:(?:uc|open)\?.*?` +
			`id=|file/d/)|video\.google\.com/get_player\?.*?docid=)(?:[a-zA-Z0-9_-]{28,}))|(?` +
			`:https?://on-demand\.gputechconf\.com/gtc/2015/video/S(?:\d+)\.html)|(?:https?:/` +
			`/(?:www\.)?hbo\.com/(?:video|embed)(?:/[^/]+)*/(?:[^/?#]+))|(?:https?://(?:www\.` +
			`)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?:https?://(?:www\.)?hitrecord\.or` +
			`g/records/(?:\d+))|(?:http?://(?:www\.)?hornbunny\.com/videos/(?:[a-z-]+)-(?:\d+` +
			`)\.html)|(?:https?://[\da-z-]+\.(?:howstuffworks|stuff(?:(?:youshould|theydontwa` +
			`ntyouto)know|toblowyourmind|momnevertoldyou)|(?:brain|car)stuffshow|fwthinking|g` +
			`eniusstuff)\.com/(?:[^/]+/)*(?:\d+-)?(?:.+?)-video\.htm)|(?:https?://(?:www\.)?h` +
			`ypem\.com/track/(?:[0-9a-z]{5}))|(?:https?://.+?\.ign\.com/(?:[^/]+/)?(?:videos|` +
			`show_videos|articles|feature|(?:[^/]+/\d+/video))(/.+)?/(?:.+))|(?:https?://(?:w` +
			`ww|m)\.imdb\.com/(?:video|title|list).+?[/-]vi(?:\d+))|(?:https?://(?:i\.)?imgur` +
			`\.com/(?!(?:a|gallery|(?:t(?:opic)?|r)/[^/]+)/)(?:[a-zA-Z0-9]+))|(?:https?://(?:` +
			`i\.)?imgur\.com/(?:gallery|(?:t(?:opic)?|r)/[^/]+)/(?:[a-zA-Z0-9]+))|(?:https?:/` +
			`/(?:www\.)?ina\.fr/(?:video|audio)/(?:[A-Z0-9_]+))|(?:https?://(?:www\.)?infoq\.` +
			`com/(?:[^/]+/)+(?:[^/]+))|(?:(?:https?://(?:www\.)?instagram\.com/p/(?:[^/?#&]+)` +
			`))|(?:https?://(?:www\.)?90tv\.ir/video/(?:[0-9]+)/.*)|(?:https?://(?:www\.)?ivi` +
			`deon\.com/tv/(?:[^/]+/)*camera/(?:\d+-[\da-f]+)/(?:\d+))|(?:https?://(?:(?:www|m` +
			`)\.)?izlesene\.com/(?:video|embedplayer)/(?:[^/]+/)?(?:[0-9]+))|(?:https?://(?:l` +
			`icensing\.jamendo\.com/[^/]+|(?:www\.)?jamendo\.com)/track/(?:[0-9]+)/(?:[^/?#&]` +
			`+))|(?:https?://.*?\.jeuxvideo\.com/.*/(.*?)\.htm)|(?:(?:joj:|https?://media\.jo` +
			`j\.sk/embed/)(?:[^/?#^]+))|(?:(?:https?://(?:content\.jwplatform|cdn\.jwplayer)\` +
			`.com/(?:(?:feed|player|thumb|preview|video)s|jw6|v2/media)/|jwplatform:)(?:[a-zA` +
			`-Z0-9]{8}))|(?:https?://tv\.kakao\.com/channel/(?:\d+)/cliplink/(?:\d+))|(?:(?:k` +
			`altura:(?:\d+):(?:[0-9a-z_]+)|https?://(:?(?:www|cdnapi(?:sec)?)\.)?kaltura\.com` +
			`(?::\d+)?/(?:(?:index\.php/(?:kwidget|extwidget/preview)|html5/html5lib/[^/]+/mw` +
			`EmbedFrame\.php))(?:/(?:[^?]+))?(?:\?(?:.*))?))|(?:https?://(?:www\.)?keezmovies` +
			`\.com/video/(?:(?:[^/]+)-)?(?:\d+))|(?:^https?://(?:(?:www|api)\.)?khanacademy\.` +
			`org/(?:[^/]+)/(?:[^/]+/){,2}(?:[^?#/]+)(?:$|[?#]))|(?:https?://(?:www\.)?kicksta` +
			`rter\.com/projects/(?:[^/]*)/.*)|(?:https?://(?:www\.)?kuwo\.cn/yinyue/(?:\d+))|` +
			`(?:https?://(?:www\.)?kuwo\.cn/mv/(?:\d+?)/)|(?:(https?://)?(?:(?:www\.)?la7\.it` +
			`/([^/]+)/(?:rivedila7|video)/|tg\.la7\.it/repliche-tgla7\?id=)(?:.+))|(?:https?:` +
			`//(?:www\.)?lci\.fr/[^/]+/[\w-]+-(?:\d+)\.html)|(?:https?://(?:www\.)?loc\.gov/(` +
			`?:item/|today/cyberlc/feature_wdesc\.php\?.*\brec=)(?:[0-9a-z_.]+))|(?:(?:https?` +
			`://html5-player\.libsyn\.com/embed/episode/id/(?:[0-9]+)))|(?:(?:limelight:media` +
			`:|https?://(?:link\.videoplatform\.limelight\.com/media/|assets\.delvenetworks\.` +
			`com/player/loader\.swf)\?.*?\bmediaId=)(?:[a-z0-9]{32}))|(?:https?://(?:new\.)?l` +
			`ivestream\.com/(?:accounts/(?:\d+)|(?:[^/]+))/(?:events/(?:\d+)|(?:[^/]+))(?:/vi` +
			`deos/(?:\d+))?)|(?:https?://(?:www\.)?lovehomeporn\.com/video/(?:\d+)(?:/(?:[^/?` +
			`#&]+))?)|(?:https?://(?:www\.)?(?:lynda\.com|educourse\.ga)/(?:(?:[^/]+/){2,3}(?` +
			`:\d+)|player/embed)/(?:\d+))|(?:https?://(?:www\.)?macgamestore\.com/mediaviewer` +
			`\.php\?trailer=(?:\d+))|(?:https?://my\.mail\.ru/music/songs/[^/?#&]+-(?:[\da-f]` +
			`+))|(?:(?i)https?://(?:www\.)?manyvids\.com/video/(?:\d+))|(?:https?://(?:www\.)` +
			`?massengeschmack\.tv/play/(?:[^?&#]+))|(?:(?:mediaset:|https?://(?:(?:www|static` +
			`3)\.)?mediasetplay\.mediaset\.it/(?:(?:video|on-demand)/(?:[^/]+/)+[^/]+_|player` +
			`/index\.html\?.*?\bprogramGuid=))(?:[0-9A-Z]{16}))|(?:https://player\.megaphone\` +
			`.fm/(?:[A-Z0-9]+))|(?:https?://(?:www\.)?metacritic\.com/.+?/trailers/(?:\d+))|(` +
			`?:(?:mva:|https?://(?:mva\.microsoft|(?:www\.)?microsoftvirtualacademy)\.com/[^/` +
			`]+/training-courses/[^/?#&]+-)(?:\d+)(?::|\?l=)(?:[\da-zA-Z]+_\d+))|(?:https?://` +
			`(?:[\da-z_-]+\.)*(?:mlb)\.com/(?:(?:(?:[^/]+/)*c-|(?:shared/video/embed/(?:embed` +
			`|m-internal-embed)\.html|(?:[^/]+/)+(?:play|index)\.jsp|)\?.*?\bcontent_id=)(?:\` +
			`d+)))|(?:https?://(?:www\.)?mofosex\.com/videos/(?:\d+)/(?:[^/?#&.]+)\.html)|(?:` +
			`https?://(?:www\.)?mojvideo\.com/video-(?:[^/]+)/(?:[a-f0-9]+))|(?:https?://mysp` +
			`ace\.com/[^/]+/(?:video/[^/]+/(?:\d+)|music/song/[^/?#&]+-(?:\d+)-\d+(?:[/?#&]|$` +
			`)))|(?:(?:https?://(?:www\.)?myvi\.(?:(?:ru/player|tv)/(?:(?:embed/html|flash|ap` +
			`i/Video/Get)/|content/preloader\.swf\?.*\bid=)|ru/watch/)|myvi:)(?:[\da-zA-Z_-]+` +
			`))|(?:https?://video\.nationalgeographic\.com/.*?)|(?:https?://(?:m\.)?tv(?:cast` +
			`)?\.naver\.com/v/(?:\d+))|(?:https?://(?:watch\.|www\.)?nba\.com/(?:(?:[^/]+/)+(` +
			`?:[^?]*?))/?(?:/index\.html)?(?:\?.*)?$)|(?:https?://(?:www\.)?csnne\.com/video/` +
			`(?:[0-9a-z-]+))|(?:https?://(?:www\.)?(?:nbcnews|today|msnbc)\.com/([^/]+/)*(?:.` +
			`*-)?(?:[^/?]+))|(?:https?://(?:www\.)?nbcsports\.com//?(?:[^/]+/)+(?:[0-9a-z-]+)` +
			`)|(?:https?://vplayer\.nbcsports\.com/(?:[^/]+/)+(?:[0-9a-zA-Z_]+))|(?:https?://` +
			`(?:www\.)?ndr\.de/(?:[^/]+/)*(?:[^/?#]+),[\da-z]+\.html)|(?:https?://(?:www\.)?n` +
			`dr\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)\.html)|(?:https?://(?:w` +
			`ww\.)?n-joy\.de/(?:[^/]+/)*(?:[\da-z]+)-(?:player|externalPlayer)_[^/]+\.html)|(` +
			`?:https?://(?:[^/]+\.)?ndtv\.com/(?:[^/]+/)*videos?/?(?:[^/]+/)*[^/?^&]+-(?:\d+)` +
			`)|(?:https?://(?:www\.)?netzkino\.de/\#!/(?:[^/]+)/(?:[^/]+))|(?:https?://(?:www` +
			`\.)?newgrounds\.com/(?:audio/listen|portal/view)/(?:[0-9]+))|(?:https?://(?:www\` +
			`.)?nexttv\.com\.tw/(?:[^/]+/)+(?:\d+))|(?:(?:https?://api\.nexx(?:\.cloud|cdn\.c` +
			`om)/v3/(?:\d+)/videos/byid/|nexx:(?:(?:\d+):)?|https?://arc\.nexx\.cloud/api/vid` +
			`eo/)(?:\d+))|(?:https?://(?:www\.)?(?:nhl|wch2016)\.com/(?:[^/]+/)*c-(?:\d+))|(?` +
			`:https?://(?:[^/]+\.)?noovo\.ca/videos/(?:[^/]+/[^/?#&]+))|(?:https?://media\.cm` +
			`s\.nova\.cz/embed/(?:[^/?#&]+))|(?:https?://(?:[^.]+\.)?(?:tv(?:noviny)?|tn|nova` +
			`plus|vymena|fanda|krasna|doma|prask)\.nova\.cz/(?:[^/]+/)+(?:[^/]+?)(?:\.html|/|` +
			`$))|(?:https?://(?:(?:www|cn)\.)?nowness\.com/(?:story|(?:series|category)/[^/]+` +
			`)/(?:[^/]+?)(?:$|[?#]))|(?:https?://(?:www\.)?ntv\.ru/(?:[^/]+/)*(?:[^/?#&]+))|(` +
			`?:https?://(?:(?:www|m|mobile)\.)?(?:odnoklassniki|ok)\.ru/(?:video(?:embed)?/|w` +
			`eb-api/video/moviePlayer/|live/|dk\?.*?st\.mvId=)(?:[\d-]+))|(?:https?://(?:www\` +
			`.)?onet\.tv/[a-z]/[a-z]+/(?:[0-9a-z-]+)/(?:[0-9a-z]+))|(?:https?://(?:www\.)?one` +
			`t\.tv/[a-z]/(?:[a-z]+)(?:[?#]|$))|(?:(?:ooyala:|https?://.+?\.ooyala\.com/.*?(?:` +
			`embedCode|ec)=)(?:.+?)(&|$))|(?:https?://(?:(?:www\.)?(?:verystream\.com))/(?:st` +
			`ream|e)/(?:[a-zA-Z0-9-_]+))|(?:https?://(?:www\.)?outsidetv\.com/(?:[^/]+/)*?pla` +
			`y/[a-zA-Z0-9]{8}/\d+/\d+/(?:[a-zA-Z0-9]{8}))|(?:https?://(?:(?:www\.)?packtpub\.` +
			`com/mapt|subscription\.packtpub\.com)/video/[^/]+/(?:\d+)/(?:\d+)/(?:\d+))|(?:ht` +
			`tps?://(?:(?:www\.)?pandora\.tv/view/(?:[^/]+)/(?:\d+)|(?:.+?\.)?channel\.pandor` +
			`a\.tv/channel/video\.ptv\?|m\.pandora\.tv/?\?))|(?:(?i)https?://(?:www\.)?parlia` +
			`mentlive\.tv/Event/Index/(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{` +
			`12}))|(?:https?://(?:www\.)?patreon\.com/(?:creation\?hid=|posts/(?:[\w-]+-)?)(?` +
			`:\d+))|(?:https?://(?:(?:(?:video|www|player)\.pbs\.org|video\.aptv\.org|video\.` +
			`gpb\.org|video\.mpbonline\.org|video\.wnpt\.org|video\.wfsu\.org|video\.wsre\.or` +
			`g|video\.wtcitv\.org|video\.pba\.org|video\.alaskapublic\.org|video\.azpbs\.org|` +
			`portal\.knme\.org|video\.vegaspbs\.org|watch\.aetn\.org|video\.ket\.org|video\.w` +
			`kno\.org|video\.lpb\.org|videos\.oeta\.tv|video\.optv\.org|watch\.wsiu\.org|vide` +
			`o\.keet\.org|pbs\.kixe\.org|video\.kpbs\.org|video\.kqed\.org|vids\.kvie\.org|vi` +
			`deo\.pbssocal\.org|video\.valleypbs\.org|video\.cptv\.org|watch\.knpb\.org|video` +
			`\.soptv\.org|video\.rmpbs\.org|video\.kenw\.org|video\.kued\.org|video\.wyomingp` +
			`bs\.org|video\.cpt12\.org|video\.kbyueleven\.org|video\.thirteen\.org|video\.wgb` +
			`h\.org|video\.wgby\.org|watch\.njtvonline\.org|watch\.wliw\.org|video\.mpt\.tv|w` +
			`atch\.weta\.org|video\.whyy\.org|video\.wlvt\.org|video\.wvpt\.net|video\.whut\.` +
			`org|video\.wedu\.org|video\.wgcu\.org|video\.wpbt2\.org|video\.wucftv\.org|video` +
			`\.wuft\.org|watch\.wxel\.org|video\.wlrn\.org|video\.wusf\.usf\.edu|video\.scetv` +
			`\.org|video\.unctv\.org|video\.pbshawaii\.org|video\.idahoptv\.org|video\.ksps\.` +
			`org|watch\.opb\.org|watch\.nwptv\.org|video\.will\.illinois\.edu|video\.networkk` +
			`nowledge\.tv|video\.wttw\.com|video\.iptv\.org|video\.ninenet\.org|video\.wfwa\.` +
			`org|video\.wfyi\.org|video\.mptv\.org|video\.wnin\.org|video\.wnit\.org|video\.w` +
			`pt\.org|video\.wvut\.org|video\.weiu\.net|video\.wqpt\.org|video\.wycc\.org|vide` +
			`o\.wipb\.org|video\.indianapublicmedia\.org|watch\.cetconnect\.org|video\.thinkt` +
			`v\.org|video\.wbgu\.org|video\.wgvu\.org|video\.netnebraska\.org|video\.pioneer\` +
			`.org|watch\.sdpb\.org|video\.tpt\.org|watch\.ksmq\.org|watch\.kpts\.org|watch\.k` +
			`twu\.org|watch\.easttennesseepbs\.org|video\.wcte\.tv|video\.wljt\.org|video\.wo` +
			`su\.org|video\.woub\.org|video\.wvpublic\.org|video\.wkyupbs\.org|video\.kera\.o` +
			`rg|video\.mpbn\.net|video\.mountainlake\.org|video\.nhptv\.org|video\.vpt\.org|v` +
			`ideo\.witf\.org|watch\.wqed\.org|video\.wmht\.org|video\.deltabroadcasting\.org|` +
			`video\.dptv\.org|video\.wcmu\.org|video\.wkar\.org|wnmuvideo\.nmu\.edu|video\.wd` +
			`se\.org|video\.wgte\.org|video\.lptv\.org|video\.kmos\.org|watch\.montanapbs\.or` +
			`g|video\.krwg\.org|video\.kacvtv\.org|video\.kcostv\.org|video\.wcny\.org|video\` +
			`.wned\.org|watch\.wpbstv\.org|video\.wskg\.org|video\.wxxi\.org|video\.wpsu\.org` +
			`|on-demand\.wvia\.org|video\.wtvi\.org|video\.westernreservepublicmedia\.org|vid` +
			`eo\.ideastream\.org|video\.kcts9\.org|video\.basinpbs\.org|video\.houstonpbs\.or` +
			`g|video\.klrn\.org|video\.klru\.tv|video\.wtjx\.org|video\.ideastations\.org|vid` +
			`eo\.kbtc\.org)/(?:(?:vir|port)alplayer|video)/(?:[0-9]+)(?:[?/]|$)|(?:www\.)?pbs` +
			`\.org/(?:[^/]+/){1,5}(?:[^/]+?)(?:\.html)?/?(?:$|[?\#])|(?:video|player)\.pbs\.o` +
			`rg/(?:widget/)?partnerplayer/(?:[^/]+)/))|(?:https?://(?:www\.)?pearvideo\.com/v` +
			`ideo_(?:\d+))|(?:https?://(?:www\.)?people\.com/people/videos/0,,(?:\d+),00\.htm` +
			`l)|(?:https?://player\.performgroup\.com/eplayer(?:/eplayer\.html|\.js)#/?(?:[0-` +
			`9a-f]{26})\.(?:[0-9a-z]{26}))|(?:https?://(?:[a-z0-9]+\.)?photobucket\.com/.*(([` +
			`\?\&]current=)|_)(?:.*)\.(?:(flv)|(mp4)))|(?:https?://player\.piksel\.com/v/(?:[` +
			`a-z0-9]+))|(?:https?://(?:www\.)?play\.fm/(?:(?:[^/]+/)+(?:[^/]+))/?(?:$|[?#]))|` +
			`(?:https?://(?:www\.)?plays\.tv/(?:video|embeds)/(?:[0-9a-f]{18}))|(?:https?://(` +
			`?:www\.)?playvid\.com/watch(\?v=|/)(?:.+?)(?:#|$))|(?:(?:https?)://(?:(?:[^.]+)\` +
			`.podomatic\.com/entry|(?:www\.)?podomatic\.com/podcasts/(?:[^/]+)/episodes)/(?:[` +
			`^/?#&]+))|(?:https?://(?:www\.)?pokemon\.com/[a-z]{2}(?:.*?play=(?:[a-z0-9]{32})` +
			`|/(?:[^/]+/)+(?:[^/?#&]+)))|(?:https?://(?:[a-zA-Z]+\.)?porn\.com/videos/(?:(?:[` +
			`^/]+)-)?(?:\d+))|(?:https?://(?:www\.)?pornhd\.com/(?:[a-z]{2,4}/)?videos/(?:\d+` +
			`)(?:/(?:.+))?)|(?:https?://(?:(?:[^/]+\.)?(?:pornhub\.(?:com|net))/(?:(?:view_vi` +
			`deo\.php|video/show)\?viewkey=|embed/)|(?:www\.)?thumbzilla\.com/video/)(?:[\da-` +
			`z]+))|(?:https?://(?:\w+\.)?pornotube\.com/(?:[^?#]*?)/video/(?:[0-9]+))|(?:http` +
			`s?://(?:www\.)?puhutv\.com/(?:[^/?#&]+)-izle)|(?:https?://(?:www\.)?presstv\.ir/` +
			`[^/]+/(?:\d+)/(?:\d+)/(?:\d+)/(?:\d+)/(?:[^/]+)?)|(?:https?://(?:.+?)\.(?:radio\` +
			`.(?:de|at|fr|pt|es|pl|it)|rad\.io))|(?:https?://(?:www\.)?radiojavan\.com/videos` +
			`/video/(?:[^/]+)/?)|(?:https?://[^/]+\.(?:rai\.(?:it|tv)|rainews\.it)/.+?-(?:[\d` +
			`a-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12})(?:-.+?)?\.html)|(?:https?:` +
			`//(?:videos\.raywenderlich\.com/courses|(?:www\.)?raywenderlich\.com)/(?:[^/]+)/` +
			`lessons/(?:\d+))|(?:https?://(?:(?:www\.)?redtube\.com/|embed\.redtube\.com/\?.*` +
			`?\bid=)(?:[0-9]+))|(?:(?:rentv:|https?://(?:www\.)?ren\.tv/(?:player|video/epizo` +
			`d)/)(?:\d+))|(?:^https?://(?:www\.)?reverbnation\.com/.*?/song/(?:\d+).*?$)|(?:h` +
			`ttps?://(?:www\.)?rockstargames\.com/videos(?:/video/|#?/?\?.*\bvideo=)(?:\d+))|` +
			`(?:https?://(?:www\.)?prehravac\.rozhlas\.cz/audio/(?:[0-9]+))|(?:rts:(?:\d+)|ht` +
			`tps?://(?:.+?\.)?rts\.ch/(?:[^/]+/){2,}(?:[0-9]+)-(?:.+?)\.html)|(?:https?://(?:` +
			`www\.)?rtvs\.sk/(?:radio|televizia)/archiv/\d+/(?:\d+))|(?:https?://(?:test)?pla` +
			`yer\.(?:rutv\.ru|vgtrk\.com)/(?:flash\d+v/container\.swf\?id=|iframe/(?:swf|vide` +
			`o|live)/id/|index/iframe/cast_id/)(?:\d+))|(?:https?://(?:www\.)?(?:ruutu|supla)` +
			`\.fi/(?:video|supla)/(?:\d+))|(?:https?://(?:www\.)?(?:safaribooksonline|(?:lear` +
			`ning\.)?oreilly)\.com/(?:library/view/[^/]+/(?:[^/]+)/(?:[^/?\#&]+)\.html|videos` +
			`/[^/]+/[^/]+/(?:[^-]+-[^/?\#&]+)))|(?:https?://(?:(?:v2|www)\.)?videos\.sapo\.(?` +
			`:pt|cv|ao|mz|tl)/(?:[\da-zA-Z]{20}))|(?:https?://(?:www\.)?sbs\.com\.au/(?:ondem` +
			`and|news)/video/(?:single/)?(?:[0-9]+))|(?:https?://(?:www\.)?screencast\.com/t/` +
			`(?:[a-zA-Z0-9]+))|(?:https?://(?:www\.)?senate\.gov/isvp/?\?(?:.+))|(?:https?://` +
			`(?:www\.)?seznamzpravy\.cz/iframe/player\?.*\bsrc=)|(?:https?://(?:.*?\.)?video\` +
			`.sina\.com\.cn/(?:(?:view/|.*\#)(?:\d+)|.+?/(?:[^/?#]+)(?:\.s?html)|api/sinawebA` +
			`pi/outplay.php/(?:.+?)\.swf))|(?:https?://news\.sky\.com/video/[0-9a-z-]+-(?:[0-` +
			`9]+))|(?:https?://(?:www\.)?skysports\.com/watch/video/(?:[0-9]+))|(?:https?://(` +
			`?:www\.)?slideshare\.net/[^/]+?/(?:.+?)($|\?))|(?:https?://slideslive\.com/(?:[0` +
			`-9]+))|(?:https?://(?:www\.)?sonyliv\.com/details/[^/]+/(?:\d+))|(?:^(?:https?:/` +
			`/)?(?:(?:(?:www\.|m\.)?soundcloud\.com/(?!stations/track)(?:[\w\d-]+)/(?!(?:trac` +
			`ks|albums|sets(?:/.+?)?|reposts|likes|spotlight)/?(?:$|[?#]))(?:[\w\d-]+)/?(?:[^` +
			`?]+?)?(?:[?].*)?$)|(?:api\.soundcloud\.com/tracks/(?:\d+)(?:/?\?secret_token=(?:` +
			`[^&]+))?)|(?:(?:w|player|p.)\.soundcloud\.com/player/?.*?url=.*)))|(?:https?://(` +
			`?:www\.)?soundgasm\.net/u/(?:[0-9a-zA-Z_-]+)/(?:[0-9a-zA-Z_-]+))|(?:https?://(?:` +
			`[^/]+\.)?spankbang\.com/(?:[\da-z]+)/(?:video|play|embed)\b)|(?:https?://(?:www\` +
			`.)?spiegel\.de/video/[^/]*-(?:[0-9]+)(?:-embed|-iframe)?(?:\.html)?(?:#.*)?$)|(?` +
			`:https?://(?:www\.)?stitcher\.com/podcast/(?:[^/]+/)+e/(?:(?:[^/#?&]+?)-)?(?:\d+` +
			`)(?:[/#?&]|$))|(?:https?://(?:(?:www|play)\.)?(?:srf|rts|rsi|rtr|swissinfo)\.ch/` +
			`play/(?:tv|radio)/[^/]+/(?:video|audio)/[^?]+\?id=(?:[0-9a-f\-]{36}|\d+))|(?:htt` +
			`ps?://openclassroom\.stanford\.edu(?:/?|(/MainFolder/(?:HomePage|CoursePage|Vide` +
			`oPage)\.php([?]course=(?:[^&]+)(&video=(?:[^&]+))?(&.*)?)?))$)|(?:https?://strea` +
			`mable\.com/(?:[es]/)?(?:\w+))|(?:https?://(?:www\.)?(?:streamango\.com|fruithost` +
			`s\.net|streamcherry\.com)/(?:f|embed)/(?:[^/?#&]+))|(?:https?://(?:(?:www\.)?sun` +
			`porno\.com/videos|embeds\.sunporno\.com/embed)/(?:\d+))|(?:https?://(?:www\.)?sv` +
			`erigesradio\.se/(?:sida/)?avsnitt/(?:[0-9]+))|(?:https?://(?:www\.)?sverigesradi` +
			`o\.se/sida/(?:artikel|gruppsida)\.aspx\?.*?\bartikel=(?:[0-9]+))|(?:https?://(?:` +
			`www\.)?tagesschau\.de/(?:[^/]+/(?:[^/]+/)*?(?:[^/#?]+?(?:-?[0-9]+)?))(?:~_?[^/#?` +
			`]+?)?\.html)|(?:https?://(?:www\.)?tastytrade\.com/tt/shows/[^/]+/episodes/(?:[^` +
			`/?#&]+))|(?:https?://tds\.lifeway\.com/v1/trainingdeliverysystem/courses/(?:\d+)` +
			`/index\.html)|(?:https?://(?:www\.)?teachertube\.com/(viewVideo\.php\?video_id=|` +
			`music\.php\?music_id=|video/(?:[\da-z-]+-)?|audio/)(?:\d+))|(?:https?://(?:\w+\.` +
			`)?teamcoco\.com/(?:([^/]+/)*[^/?#]+))|(?:https?://(?:www\.)?teamtreehouse\.com/l` +
			`ibrary/(?:[^/]+))|(?:(?:https?://)(?:www|embed(?:-ssl)?)(?:\.ted\.com/((?:playli` +
			`sts(?:/\d+)?)|((?:talks))|(?:watch)/[^/]+/[^/]+)(/lang/(.*?))?/(?:[\w-]+).*)$)|(` +
			`?:^https?://(?:www\.)?t13\.cl/videos(?:/[^/]+)+/(?:[\w-]+))|(?:https?://(?:www\.` +
			`)?tenta\.com/how-to-download-videos)|(?:https?://(?:www\.)?tfo\.org/(?:en|fr)/(?` +
			`:[^/]+/){2}(?:\d+))|(?:(?:https?://(?:link|player)\.theplatform\.com/[sp]/(?:[^/` +
			`]+)/(?:(?:(?:[^/]+/)+select/)?(?:media/(?:guid/\d+/)?)?|(?:(?:[^/\?]+/(?:swf|con` +
			`fig)|onsite)/select/))?|theplatform:)(?:[^/\?&]+))|(?:https?://feed\.theplatform` +
			`\.com/f/(?:[^/]+)/(?:[^?/]+)\?(?:[^&]+&)*(?:by(?:Gui|I)d=(?:[^&]+)))|(?:https?:/` +
			`/thescene\.com/watch/[^/]+/(?:[^/#?]+))|(?:https?://(?:www\.)?thisoldhouse\.com/` +
			`(?:watch|how-to|tv-episode)/(?:[^/?#]+))|(?:https?://(?:.+?\.)?tinypic\.com/play` +
			`er\.php\?v=(?:[^&]+)&s=\d+)|(?:https?://(?:www\.)?tmz\.com/videos/(?:[^/?#]+))|(` +
			`?:https?://(?:www\.)?tmz\.com/\d{4}/\d{2}/\d{2}/(?:[^/]+)/?)|(?:https?://player\` +
			`.(?:tna|emp)flix\.com/video/(?:\d+))|(?:https?://(?:www\.)?tnaflix\.com/[^/]+/(?` +
			`:[^/]+)/video(?:\d+))|(?:https?://(?:www\.)?empflix\.com/(?:videos/(?:.+?)-|[^/]` +
			`+/(?:[^/]+)/video)(?:[0-9]+))|(?:https?://(?:www\.)?moviefap\.com/videos/(?:[0-9` +
			`a-f]+)/(?:[^/]+)\.html)|(?:https?://(?:www\.)?toongoggles\.com/shows/(?:\d+)(?:/` +
			`[^/]+/episodes/(?:\d+))?)|(?:https?://videos\.toypics\.net/view/(?:[0-9]+))|(?:h` +
			`ttps?://(?:(?:www|m)\.)?trilulilu\.ro/(?:[^/]+/)?(?:[^/#\?]+))|(?:https?://(?:ww` +
			`w\.)?tube8\.com/(?:[^/]+/)+(?:[^/]+)/(?:\d+))|(?:https?://(?:[^/?#&]+)\.tumblr\.` +
			`com/(?:post|video)/(?:[0-9]+)(?:$|[/?#]))|(?:https?://(?:(?:www\.)?tune\.pk/(?:v` +
			`ideo/|player/embed_player.php?.*?\bvid=)|embed\.tune\.pk/play/)(?:\d+))|(?:https` +
			`?://(?:www\.)?tvanouvelles\.ca/videos/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/vid` +
			`eo/iframe/id/(?:\d+))|(?:https?://(?:www\.)?tvc\.ru/(?!video/iframe/id/)(?:[^?#]` +
			`+))|(?:https?://(?:(?:[^/]+)\.)?tvn24(?:bis)?\.pl/(?:[^/]+/)*(?:[^/]+))|(?:(?:tv` +
			`p:|https?://[^/]+\.tvp\.(?:pl|info)/sess/tvplayer\.php\?.*?object_id=)(?:\d+))|(` +
			`?:https?://[^/]+\.tvp\.(?:pl|info)/(?:video/(?:[^,\s]*,)*|(?:(?!\d+/)[^/]+/)*)(?` +
			`:\d+))|(?:(?:mtg:|https?://(?:www\.)?(?:tvplay(?:\.skaties)?\.lv(?:/parraides)?|` +
			`(?:tv3play|play\.tv3)\.lt(?:/programos)?|tv3play(?:\.tv3)?\.ee/sisu|(?:tv(?:3|6|` +
			`8|10)play|viafree)\.se/program|(?:(?:tv3play|viasat4play|tv6play|viafree)\.no|(?` +
			`:tv3play|viafree)\.dk)/programmer|play\.nova(?:tv)?\.bg/programi)/(?:[^/]+/)+)(?` +
			`:\d+))|(?:https?://(?:www\.)?20min\.ch/(?:videotv/*\?.*?\bvid=|videoplayer/video` +
			`player\.html\?.*?\bvideoId@)(?:\d+))|(?:https?://video\.(?:twentythree\.net|23vi` +
			`deo\.com|filmweb\.no)/v\.ihtml/player\.html\?(?:.*?\bphoto(?:_|%5f)id=(?:\d+).*)` +
			`)|(?:https?://(?:clips\.twitch\.tv/(?:[^/]+/)*|(?:www\.)?twitch\.tv/[^/]+/clip/)` +
			`(?:[^/?#&]+))|(?:https?://(?:www\.)?twitter\.com/i/(?:cards/tfw/v1|videos(?:/twe` +
			`et)?)/(?:\d+))|(?:https?://amp\.twimg\.com/v/(?:[0-9a-f\-]{36}))|(?:https?://vid` +
			`eo\.udn\.com/(?:embed|play)/news/(?:\d+))|(?:https?://(?:www\.)?(?:digiteka\.net` +
			`|ultimedia\.com)/(?:deliver/(?:generic|musique)(?:/[^/]+)*/(?:src|article)|defau` +
			`lt/index/video(?:generic|music)/id)/(?:[\d+a-z]+))|(?:https?://utv\.unistra\.fr/` +
			`(?:index|video)\.php\?id_video\=(?:\d+))|(?:https?://(?:www\.)?unity3d\.com/lear` +
			`n/tutorials/(?:[^/]+/)*(?:[^/?#&]+))|(?:https?://(?:www\.)?usatoday\.com/(?:[^/]` +
			`+/)*(?:[^?/#]+))|(?:https?://(?:(?:app|embed)\.)?ustudio\.com/embed/(?:[^/]+)/(?` +
			`:[^/]+))|(?:https?://(?:www\.)?video\.varzesh3\.com/(?:[^/]+/)+(?:[^/]+)/?)|(?:h` +
			`ttps?://(?:[^/]+\.)?vbox7\.com/(?:play:|(?:emb/external\.php|player/ext\.swf)\?.` +
			`*?\bvid=)(?:[\da-fA-F]+))|(?:https?://veehd\.com/video/(?:\d+))|(?:https?://(?:w` +
			`ww\.)?veoh\.com/(?:watch|embed|iphone/#_Watch)/(?:(?:v|e|yapi-)[\da-zA-Z]+))|(?:` +
			`https?://(?:.+?\.)?vesti\.ru/(?:.+))|(?:(?:https?://(?:www\.)?vevo\.com/watch/(?` +
			`!playlist|genre)(?:[^/]+/(?:[^/]+/)?)?|https?://cache\.vevo\.com/m/html/embed\.h` +
			`tml\?video=|https?://videoplayer\.vevo\.com/embed/embedded\?videoId=|vevo:)(?:[^` +
			`&?#]+))|(?:(?:https?://(?:www\.)?(?:vgtv.no|bt.no/tv|aftenbladet.no/tv|fvn.no/fv` +
			`ntv|aftenposten.no/webtv|ap.vgtv.no/webtv|tv.aftonbladet.se/abtv|www.aftonbladet` +
			`.se/tv)/?(?:(?:\#!/)?(?:video|live)/|embed?.*id=|a(?:rticles)?/)|(?:vgtv|bttv|sa` +
			`tv|fvntv|aptv|abtv):)(?:\d+))|(?:https://www\.vice\.com/[^/]+/article/(?:[^?#]+)` +
			`)|(?:https?://(?:www\.)?viddler\.com/(?:v|embed|player)/(?:[a-z0-9]+)(?:.+?\bsec` +
			`ret=(\d+))?)|(?:https?://videa(?:kid)?\.hu/(?:videok/(?:[^/]+/)*[^?#&]+-|player\` +
			`?.*?\bv=|player/v/)(?:[^?#&]+))|(?:https?://videopress\.com/embed/(?:[\da-zA-Z]+` +
			`))|(?:https?://(?:www\.|m\.|)videosz\.com/[a-z]+/[a-z]+/(?:[0-9]+)_)|(?:https?:/` +
			`/(?:www\.)?vidlii\.com/(?:watch|embed)\?.*?\bv=(?:[0-9A-Za-z_-]{11}))|(?:https?:` +
			`//(?:(?:www|(?:player))\.)?vimeo(?:pro)?\.com/(?!(?:channels|album)/[^/?#]+/?(?:` +
			`$|[?#])|[^/]+/review/|ondemand/)(?:.*?/)?(?:(?:play_redirect_hls|moogaloop\.swf)` +
			`\?clip_id=)?(?:videos?/)?(?:[0-9]+)(?:/[\da-f]+)?/?(?:[?&].*)?(?:[#].*)?$)|(?:ht` +
			`tps?://(?:www\.)?vimeo\.com/ondemand/(?:[^/?#&]+))|(?:(?:https://vimeo\.com/[^/]` +
			`+/review/(?:[^/]+)/[0-9a-f]{10}))|(?:https?://(?:player\.vimple\.(?:ru|co)/ifram` +
			`e|vimple\.(?:ru|co))/(?:[\da-f-]{32,36}))|(?:https?://(?:www\.)?vine\.co/(?:v|oe` +
			`mbed)/(?:\w+))|(?:(?:viqeo:|https?://cdn\.viqeo\.tv/embed/*\?.*?\bvid=|https?://` +
			`api\.viqeo\.tv/v\d+/data/startup?.*?\bvideo(?:%5B%5D|\[\])=)(?:[\da-f]+))|(?:htt` +
			`ps?://(?:(?:(?:(?:m|new)\.)?vk\.com/video_|(?:www\.)?daxab.com/)ext\.php\?(?:.*?` +
			`\boid=(?:-?\d+).*?\bid=(?:\d+).*)|(?:(?:(?:m|new)\.)?vk\.com/(?:.+?\?.*?z=)?vide` +
			`o|(?:www\.)?daxab.com/embed/)(?:-?\d+_\d+)(?:.*\blist=(?:[\da-f]+))?))|(?:https?` +
			`://(?:(?:www|m)\.)?vlive\.tv/video/(?:[0-9]+))|(?:https?://(?:(?:www|m)\.)?vlive` +
			`\.tv/video/(?:[0-9]+)/playlist/(?:[0-9]+))|(?:https?://(?:(?:www|view)\.)?vzaar\` +
			`.com/(?:videos/)?(?:\d+))|(?:(?:washingtonpost:|https?://(?:www\.)?washingtonpos` +
			`t\.com/video/(?:[^/]+/)*)(?:[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{` +
			`12}))|(?:(?:wat:|https?://(?:www\.)?wat\.tv/video/.*-)(?:[0-9a-z]+))|(?:https?:/` +
			`/m\.weibo\.cn/status/(?:[0-9]+)(\?.+)?)|(?:https?://(?:www|m)\.worldstar(?:candy` +
			`|hiphop)\.com/(?:videos|android)/video\.php\?.*?\bv=(?:[^&]+))|(?:(?:https?://vi` +
			`deo-api\.wsj\.com/api-video/player/iframe\.html\?.*?\bguid=|https?://(?:www\.)?(` +
			`?:wsj|barrons)\.com/video/(?:[^/]+/)+|wsj:)(?:[a-fA-F0-9-]{36}))|(?:(?i)https?:/` +
			`/(?:www\.)?wsj\.com/articles/(?:[^/?#&]+))|(?:https?://(?:.+?\.)?xhamster\.(?:co` +
			`m|one)/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?:https?://` +
			`(?:.+?\.)?xhamster\.com/xembed\.php\?video=(?:\d+))|(?:https?://(?:video|www)\.x` +
			`nxx\.com/video-?(?:[0-9a-z]+)/)|(?:(?:xtube:|https?://(?:www\.)?xtube\.com/(?:wa` +
			`tch\.php\?.*\bv=|video-watch/(?:embedded/)?(?:[^/]+)-))(?:[^/?&#]+))|(?:https?:/` +
			`/(?:(?:www\.)?xvideos\.com/video|flashservice\.xvideos\.com/embedframe/|static-h` +
			`w\.xvideos\.com/swf/xv-player\.swf\?.*?\bid_video=)(?:[0-9]+))|(?:https?://(?:ww` +
			`w\.)?xxxymovies\.com/videos/(?:\d+)/(?:[^/]+))|(?:(?:https?://(?:(?:[a-zA-Z]{2})` +
			`\.)?[\da-zA-Z_-]+\.yahoo\.com)/(?:[^/]+/)*(?:(?:.+)?-)?(?:[0-9]+)(?:-[a-z]+)?(?:` +
			`\.html)?)|(?:https?://v\.yinyuetai\.com/video(?:/h5)?/(?:[0-9]+))|(?:https?://(?` +
			`:\w+\.)?youjizz\.com/videos/(?:[^/#?]*-(?:\d+)\.html|embed/(?:\d+)))|(?:https?:/` +
			`/(?:www\.)?youporn\.com/watch/(?:\d+)/(?:[^/?#&]+))|(?:https?://(?:www\.)?(?:you` +
			`rupload\.com/(?:watch|embed)|embed\.yourupload\.com)/(?:[A-Za-z0-9]+))|(?:^((?:h` +
			`ttps?://|//)(?:(?:(?:(?:\w+\.)?[yY][oO][uU][tT][uU][bB][eE](?:-nocookie)?\.com/|` +
			`(?:www\.)?deturl\.com/www\.youtube\.com/|(?:www\.)?pwnyoutube\.com/|(?:www\.)?ho` +
			`oktube\.com/|(?:www\.)?yourepeat\.com/|tube\.majestyc\.net/|(?:(?:www|dev)\.)?in` +
			`vidio\.us/|(?:www\.)?invidiou\.sh/|(?:www\.)?invidious\.snopyta\.org/|(?:www\.)?` +
			`invidious\.kabi\.tk/|(?:www\.)?vid\.wxzm\.sx/|youtube\.googleapis\.com/)(?:.*?\#` +
			`/)?(?:(?:(?:v|embed|e)/(?!videoseries))|(?:(?:(?:watch|movie)(?:_popup)?(?:\.php` +
			`)?/?)?(?:\?|\#!?)(?:.*?[&;])??v=)))|(?:youtu\.be|vid\.plus|zwearz\.com/watch|)/|` +
			`(?:www\.)?cleanvideosearch\.com/media/action/yt/watch\?videoId=))?([0-9A-Za-z_-]` +
			`{11})(?!.*?\blist=(?:(?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,}|WL))` +
			`)|(?:(?:(?:https?://)?(?:\w+\.)?(?:(?:youtube\.com|invidio\.us)/(?:(?:course|vie` +
			`w_play_list|my_playlists|artist|playlist|watch|embed/(?:videoseries|[0-9A-Za-z_-` +
			`]{11}))\?(?:.*?[&;])*?(?:p|a|list)=|p/)|youtu\.be/[0-9A-Za-z_-]{11}\?.*?\blist=)` +
			`((?:PL|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)?[0-9A-Za-z-_]{10,}|(?:MC)[\w\.]*).*|((?:PL` +
			`|LL|EC|UU|FL|RD|UL|TL|OLAK5uy_)[0-9A-Za-z-_]{10,})))`,
	}
}
