package algolia

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func initAlgoliaClient() {
	client := search.NewClient("I6SY0AOGJK", "fa3d7ccef2f3f58c977bc0cdca6e02cb")
	index := client.InitIndex("vdi_search_index")

}
