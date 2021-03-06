// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/create-index.asciidoc#L190>
//
// --------------------------------------------------------------------------------
// PUT /test
// {
//     "settings": {
//         "index.write.wait_for_active_shards": "2"
//     }
// }
// --------------------------------------------------------------------------------

func Test_indices_create_index_4d46dbb96125b27f46299547de9d8709(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:4d46dbb96125b27f46299547de9d8709[]
	res, err := es.Indices.Create(
		"test",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "settings": {
		    "index.write.wait_for_active_shards": "2"
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:4d46dbb96125b27f46299547de9d8709[]
}
