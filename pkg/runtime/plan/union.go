// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package plan

import (
	"context"
)

import (
	"github.com/dubbogo/arana/pkg/proto"
)

type UnionPlan struct {
	Plans []proto.Plan
}

func (u UnionPlan) Type() proto.PlanType {
	return proto.PlanTypeQuery
}

func (u UnionPlan) ExecIn(ctx context.Context, conn proto.VConn) (proto.MixinResult, error) {
	//TODO lazy union result sets
	panic("implement me")
}
