// Copyright 2021 CloudWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

// Templates returns all templates defined in this package.
func Templates() []string {
	return []string{
		File, Enum, Typedef, Constant,
		StructLike,
		StructLikeDefault,
		StructLikeRead,
		StructLikeReadField,
		StructLikeWrite,
		StructLikeWriteField,
		FieldGetOrSet,
		FieldIsSet,
		FieldRead,
		FieldReadStructLike,
		FieldReadBaseType,
		FieldReadContainer,
		FieldReadMap,
		FieldReadSet,
		FieldReadList,
		FieldWrite,
		FieldWriteStructLike,
		FieldWriteBaseType,
		FieldWriteContainer,
		FieldWriteMap,
		FieldWriteSet,
		FieldWriteList,
		FunctionSignature, Service, Processor,
		StructLikeDeepEqual,
		StructLikeDeepEqualField,
		FieldDeepEqual,
		FieldDeepEqualBase,
		FieldDeepEqualContainer,
		FieldDeepEqualList,
		FieldDeepEqualSet,
		FieldDeepEqualMap,
		FieldDeepEqualStructLike,
	}
}