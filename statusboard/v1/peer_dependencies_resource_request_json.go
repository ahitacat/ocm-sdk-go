/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/statusboard/v1

import (
	"io"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

func writePeerDependenciesAddRequest(request *PeerDependenciesAddRequest, writer io.Writer) error {
	return MarshalPeerDependency(request.body, writer)
}
func readPeerDependenciesAddResponse(response *PeerDependenciesAddResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalPeerDependency(reader)
	return err
}
func writePeerDependenciesListRequest(request *PeerDependenciesListRequest, writer io.Writer) error {
	return nil
}
func readPeerDependenciesListResponse(response *PeerDependenciesListResponse, reader io.Reader) error {
	iterator, err := helpers.NewIterator(reader)
	if err != nil {
		return err
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "page":
			value := iterator.ReadInt()
			response.page = &value
		case "size":
			value := iterator.ReadInt()
			response.size = &value
		case "total":
			value := iterator.ReadInt()
			response.total = &value
		case "items":
			items := ReadPeerDependencyList(iterator)
			responseItems := &PeerDependencyList{}
			responseItems.SetItems(items)
			response.items = responseItems
		default:
			iterator.ReadAny()
		}
	}
	return iterator.Error
}
