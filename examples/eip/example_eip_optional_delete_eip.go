/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package eipexamples

import (
	"fmt"
)

func OptionalDeleteEip() {
	Init()

	// 被删除的EIP
	eip := "x.x.x.x"
	// 放到回收站内
	releaseToRecycle := true
	// 删除EIP
	err := eipClient.OptionalDeleteEip(eip, "", releaseToRecycle)
	if err != nil {
		panic(err)
	}

	fmt.Println("OptionalDeleteEip success")
}
