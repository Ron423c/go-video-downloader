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
 * cookiejar/module.go: transpiled from http/cookiejar.py
 */

package cookiejar

import (
	Ωnet "github.com/tenta-browser/go-video-downloader/lib/net"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	Cookie           λ.Object
	CookieJar        λ.Object
	GetCookieHeader  λ.Object
	MozillaCookieJar λ.Object
	SetCookie        λ.Object
)

func init() {
	λ.InitModule(func() {
		SetCookie = Ωnet.SetCookie
		GetCookieHeader = Ωnet.GetCookieHeader
		CookieJar = λ.Cal(λ.TypeType, λ.NewStr("CookieJar"), λ.NewTuple(), func() λ.Dict {
			var (
				CookieJar___init__          λ.Object
				CookieJar_add_cookie_header λ.Object
				CookieJar_set_cookie        λ.Object
			)
			CookieJar___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "jar"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒjar  = λargs[1]
						ϒself = λargs[0]
					)
					λ.SetAttr(ϒself, "jar", ϒjar)
					return λ.None
				})
			CookieJar_add_cookie_header = λ.NewFunction("add_cookie_header",
				[]λ.Param{
					{Name: "self"},
					{Name: "request"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒheader  λ.Object
						ϒrequest = λargs[1]
						ϒself    = λargs[0]
					)
					ϒheader = λ.Cal(GetCookieHeader, λ.GetAttr(ϒself, "jar", nil), λ.GetAttr(ϒrequest, "url", nil))
					λ.Cal(λ.GetAttr(ϒrequest, "add_header", nil), λ.NewStr("Cookie"), ϒheader)
					return λ.None
				})
			CookieJar_set_cookie = λ.NewFunction("set_cookie",
				[]λ.Param{
					{Name: "self"},
					{Name: "cookie"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcookie = λargs[1]
						ϒself   = λargs[0]
					)
					λ.Cal(SetCookie, λ.GetAttr(ϒself, "jar", nil), λ.GetAttr(ϒcookie, "domain", nil), λ.GetAttr(ϒcookie, "name", nil), λ.GetAttr(ϒcookie, "value", nil))
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"):          CookieJar___init__,
				λ.NewStr("add_cookie_header"): CookieJar_add_cookie_header,
				λ.NewStr("set_cookie"):        CookieJar_set_cookie,
			})
		}())
		MozillaCookieJar = λ.Cal(λ.TypeType, λ.NewStr("MozillaCookieJar"), λ.NewTuple(CookieJar), func() λ.Dict {
			// pass
			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
		Cookie = λ.Cal(λ.TypeType, λ.NewStr("Cookie"), λ.NewTuple(), func() λ.Dict {
			var (
				Cookie___init__ λ.Object
			)
			Cookie___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "version"},
					{Name: "name"},
					{Name: "value"},
					{Name: "port"},
					{Name: "port_specified"},
					{Name: "domain"},
					{Name: "domain_specified"},
					{Name: "domain_initial_dot"},
					{Name: "path"},
					{Name: "path_specified"},
					{Name: "secure"},
					{Name: "expires"},
					{Name: "discard"},
					{Name: "comment"},
					{Name: "comment_url"},
					{Name: "rest"},
					{Name: "rfc2109", Def: λ.False},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcomment            = λargs[14]
						ϒcomment_url        = λargs[15]
						ϒdiscard            = λargs[13]
						ϒdomain             = λargs[6]
						ϒdomain_initial_dot = λargs[8]
						ϒdomain_specified   = λargs[7]
						ϒexpires            = λargs[12]
						ϒname               = λargs[2]
						ϒpath               = λargs[9]
						ϒpath_specified     = λargs[10]
						ϒport               = λargs[4]
						ϒport_specified     = λargs[5]
						ϒrest               = λargs[16]
						ϒrfc2109            = λargs[17]
						ϒsecure             = λargs[11]
						ϒself               = λargs[0]
						ϒvalue              = λargs[3]
						ϒversion            = λargs[1]
					)
					_ = ϒcomment
					_ = ϒcomment_url
					_ = ϒdiscard
					_ = ϒdomain_initial_dot
					_ = ϒdomain_specified
					_ = ϒexpires
					_ = ϒpath
					_ = ϒpath_specified
					_ = ϒport
					_ = ϒport_specified
					_ = ϒrest
					_ = ϒrfc2109
					_ = ϒsecure
					_ = ϒversion
					λ.SetAttr(ϒself, "domain", λ.Cal(λ.GetAttr(ϒdomain, "lower", nil)))
					λ.SetAttr(ϒself, "name", ϒname)
					λ.SetAttr(ϒself, "value", ϒvalue)
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"): Cookie___init__,
			})
		}())
	})
}
