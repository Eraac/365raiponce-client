package raiponce

type Refresher interface {
	Self() string
}

type Entitable interface {
	Refresher
	create(*Client) error
	update(*Client) error
	remove(*Client) error
}

type Collectionnable interface {
	Refresher
	IsFirstPage() bool
	IsLastPage()  bool
	Previous()    (string, bool)
	Next()        (string, bool)
}

type entity struct {
	ID    int        `json:"id,omitempty"`
	Links entityLink `json:"_links,omitempty"`
}

// entityLink represent _links available into an Entity
type entityLink struct {
	Self Link `json:"self,omitempty"`
}

// Collection is abstract type for all collections
type Collection struct {
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
	Pages int            `json:"pages"`
	Total int            `json:"total"`
	Links CollectionLink `json:"_links"`
}

// CollectionLink represent _links available to a collection
type CollectionLink struct {
	Self     Link `json:"self"`
	Previous Link `json:"previous"`
	Next     Link `json:"next"`
	First    Link `json:"first"`
	Last     Link `json:"last"`
}

// Link represent an url
type Link struct {
	Href string `json:"href,omitempty"`
}

func (collection *Collection) IsFirstPage() bool {
	return bool(collection.Page == 1)
}

func (collection *Collection) IsLastPage() bool {
	return bool(collection.Page == collection.Pages)
}

func (collection *Collection) Previous() (string, bool) {
	if collection.IsFirstPage() {
		return "", false
	}

	return collection.Links.Previous.Href, true
}

func (collection *Collection) Next() (string, bool) {
	if collection.IsLastPage() {
		return "", false
	}

	return collection.Links.Next.Href, true
}

func (collection *Collection) Self() string {
	return collection.Links.Self.Href
}

func (entity *entity) Self() string {
	return entity.Links.Self.Href
}

func (entity *entity) remove(client *Client) error {
	return client.remove(entity.Self())
}

func (client *Client) Previous(collection Collectionnable) bool {
	uri, ok := collection.Previous()

	if !ok {
		return false
	}

	client.get(uri, &collection)

	return true
}

func (client *Client) Next(collection Collectionnable) bool {
	uri, ok := collection.Next()

	if !ok {
		return false
	}

	client.get(uri, &collection)

	return true
}

func (client *Client) Refresh(entity Refresher) {
	uri := entity.Self()

	client.get(uri, &entity)
}

func (client *Client) Create(entity Entitable) error {
	return entity.create(client)
}

func (client *Client) Update(entity Entitable) error {
	return entity.update(client)
}

func (client *Client) Remove(entity Entitable) error {
	return entity.remove(client)
}
