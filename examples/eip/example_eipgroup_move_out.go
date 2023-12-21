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

	EIP "github.com/baidubce/bce-sdk-go/services/eip"
)

func EipGroupMoveOut() {
	Init()

	// 移出EIP的eipgroup id
	id := "eg-xxxxxxxx"
	eipGroupMoveOutArgs := &EIP.EipGroupMoveOutArgs{
		MoveOutEips: []EIP.MoveOutEip{
			{
				Eip: "x.x.x.x", // 非原生EIP的eip
			},
			{
				Eip:             "x.x.x.x", //原生EIP的eip，需要指定计费方式
				BandWidthInMbps: 10,        //eip的带宽
				Billing: &EIP.Billing{
					PaymentTiming: "PostPaid",    //后付费
					BillingMethod: "ByBandwidth", //计费方式
				},
			},
		},
	}
	err := eipClient.EipGroupMoveOut(id, eipGroupMoveOutArgs)
	if err != nil {
		panic(err)
	}

	fmt.Println("EipGroupMoveOut success")
}
