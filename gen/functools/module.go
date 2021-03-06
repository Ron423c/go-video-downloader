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
 * functools/module.go: transpiled from functools.py
 */

package functools

import (
	λ "github.com/tenta-browser/go-video-downloader/runtime"
)

var (
	ϒreduce λ.Object
)

func init() {
	λ.InitModule(func() {
		ϒreduce = λ.NewFunction("reduce",
			[]λ.Param{
				{Name: "function"},
				{Name: "iterable"},
				{Name: "initializer", Def: λ.None},
			},
			0, false, false,
			func(λargs []λ.Object) λ.Object {
				var (
					ϒelement     λ.Object
					ϒfunction    = λargs[0]
					ϒinitializer = λargs[2]
					ϒit          λ.Object
					ϒiterable    = λargs[1]
					ϒvalue       λ.Object
					τmp0         λ.Object
					τmp1         λ.Object
				)
				ϒit = λ.Cal(λ.BuiltinIter, ϒiterable)
				if ϒinitializer == λ.None {
					ϒvalue = λ.Cal(λ.BuiltinNext, ϒit)
				} else {
					ϒvalue = ϒinitializer
				}
				τmp0 = λ.Cal(λ.BuiltinIter, ϒit)
				for {
					if τmp1 = λ.NextDefault(τmp0, λ.AfterLast); τmp1 == λ.AfterLast {
						break
					}
					ϒelement = τmp1
					ϒvalue = λ.Cal(ϒfunction, ϒvalue, ϒelement)
				}
				return ϒvalue
			})
	})
}
