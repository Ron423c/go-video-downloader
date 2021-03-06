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
 * error/module.go: transpiled from urllib/error.py
 */

package error

import (
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	HTTPError λ.Object
	URLError  λ.Object
)

func init() {
	λ.InitModule(func() {
		URLError = λ.Cal(λ.TypeType, λ.StrLiteral("URLError"), λ.NewTuple(λ.OSErrorType), func() λ.Dict {
			var (
				URLError___init__ λ.Object
				URLError___str__  λ.Object
			)
			URLError___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "reason"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒreason = λargs[1]
						ϒself   = λargs[0]
					)
					λ.SetAttr(ϒself, "args", λ.NewTuple(ϒreason))
					λ.SetAttr(ϒself, "reason", ϒreason)
					return λ.None
				})
			URLError___str__ = λ.NewFunction("__str__",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.Mod(λ.StrLiteral("<urlopen error %s>"), λ.GetAttr(ϒself, "reason", nil))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"__init__": URLError___init__,
				"__str__":  URLError___str__,
			})
		}())
		HTTPError = λ.Cal(λ.TypeType, λ.StrLiteral("HTTPError"), λ.NewTuple(URLError), func() λ.Dict {
			var (
				HTTPError___init__ λ.Object
				HTTPError___repr__ λ.Object
				HTTPError___str__  λ.Object
			)
			HTTPError___init__ = λ.NewFunction("__init__",
				[]λ.Param{
					{Name: "self"},
					{Name: "url"},
					{Name: "code"},
					{Name: "msg"},
					{Name: "hdrs"},
					{Name: "fp"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒcode = λargs[2]
						ϒfp   = λargs[5]
						ϒhdrs = λargs[4]
						ϒmsg  = λargs[3]
						ϒself = λargs[0]
						ϒurl  = λargs[1]
					)
					λ.SetAttr(ϒself, "code", ϒcode)
					λ.SetAttr(ϒself, "msg", ϒmsg)
					λ.SetAttr(ϒself, "hdrs", ϒhdrs)
					λ.SetAttr(ϒself, "fp", ϒfp)
					λ.SetAttr(ϒself, "filename", ϒurl)
					return λ.None
				})
			HTTPError___str__ = λ.NewFunction("__str__",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.Mod(λ.StrLiteral("HTTP Error %s: %s"), λ.NewTuple(
						λ.GetAttr(ϒself, "code", nil),
						λ.GetAttr(ϒself, "msg", nil),
					))
				})
			HTTPError___repr__ = λ.NewFunction("__repr__",
				[]λ.Param{
					{Name: "self"},
				},
				0, false, false,
				func(λargs []λ.Object) λ.Object {
					var (
						ϒself = λargs[0]
					)
					return λ.Mod(λ.StrLiteral("<HTTPError %s: %r>"), λ.NewTuple(
						λ.GetAttr(ϒself, "code", nil),
						λ.GetAttr(ϒself, "msg", nil),
					))
				})
			return λ.ClassDictLiteral(map[string]λ.Object{
				"__init__": HTTPError___init__,
				"__repr__": HTTPError___repr__,
				"__str__":  HTTPError___str__,
			})
		}())
	})
}
