/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package fieldpath

import (
	"github.com/containerd/containerd/protobuf/plugin"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// Enabled returns true if E_Fieldpath is enabled
func Enabled(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, plugin.E_Fieldpath, proto.GetBoolExtension(file.Options, plugin.E_FieldpathAll, false))
}
