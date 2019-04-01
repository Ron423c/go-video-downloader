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
 * cookies/module.go: transpiled from http/cookies.py
 */

package cookies

import (
	Ωnet "github.com/tenta-browser/go-video-downloader/lib/net"
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	BaseCookie        λ.Object
	Morsel            λ.Object
	ParseCookieString λ.Object
	SimpleCookie      λ.Object
)

func init() {
	λ.InitModule(func() {
		ParseCookieString = Ωnet.ParseCookieString
		Morsel = λ.Cal(λ.TypeType, λ.NewStr("Morsel"), λ.NewTuple(), func() λ.Dict {
			var (
				Morsel___init__ λ.Object
			)
			Morsel___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "key"},
					{Name: "value"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒkey   = λargs[1]
						ϒself  = λargs[0]
						ϒvalue = λargs[2]
					)
					λ.SetAttr(ϒself, "key", ϒkey)
					λ.SetAttr(ϒself, "value", ϒvalue)
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"): Morsel___init__,
			})
		}())
		BaseCookie = λ.Cal(λ.TypeType, λ.NewStr("BaseCookie"), λ.NewTuple(λ.DictType), func() λ.Dict {
			var (
				BaseCookie___init__ λ.Object
				BaseCookie_load     λ.Object
			)
			BaseCookie___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "input", Def: λ.None},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒinput = λargs[1]
						ϒself  = λargs[0]
					)
					λ.Cal(λ.GetAttr(λ.Cal(λ.SuperType, BaseCookie, ϒself), "__init__", nil))
					if λ.IsTrue(ϒinput) {
						λ.Cal(λ.GetAttr(ϒself, "load", nil), ϒinput)
					}
					return λ.None
				})
			BaseCookie_load = λ.NewFunction("load",
				[]λ.Param{
					{Name: "self"},
					{Name: "rawdata"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒname    λ.Object
						ϒrawdata = λargs[1]
						ϒself    = λargs[0]
						ϒvalue   λ.Object
						τmp0     λ.Object
						τmp1     λ.Object
						τmp2     λ.Object
					)
					τmp0 = λ.Cal(λ.BuiltinIter, λ.Cal(λ.GetAttr(λ.Cal(ParseCookieString, ϒrawdata), "items", nil)))
					for {
						if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
							break
						}
						τmp2 = τmp1
						ϒname = λ.GetItem(τmp2, λ.NewInt(0))
						ϒvalue = λ.GetItem(τmp2, λ.NewInt(1))
						λ.SetItem(ϒself, ϒname, λ.Cal(Morsel, ϒname, ϒvalue))
					}
					return λ.None
				})
			return λ.NewDictWithTable(map[λ.Object]λ.Object{
				λ.NewStr("__init__"): BaseCookie___init__,
				λ.NewStr("load"):     BaseCookie_load,
			})
		}())
		SimpleCookie = λ.Cal(λ.TypeType, λ.NewStr("SimpleCookie"), λ.NewTuple(BaseCookie), func() λ.Dict {
			// pass
			return λ.NewDictWithTable(map[λ.Object]λ.Object{})
		}())
	})
}