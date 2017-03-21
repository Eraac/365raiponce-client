package raiponce

import "strconv"

const emotionsURI = "/emotions"

// === Struct ===

type formEmotion struct {
	Name string `json:"name"`
}

type Emotion struct {
	entity
	Name string `json:"name"`
}

type emotions struct {
	Emotions []Emotion `json:"items"`
}

type CollectionEmotion struct {
	Collection
	Embedded emotions `json:"_embedded"`
}

// === Provider ===

func (client *Client) CGetEmotions(query *QueryFilter) *CollectionEmotion {
	collection := &CollectionEmotion{}

	client.cget(emotionsURI, collection, query)

	return collection
}

func (client *Client) GetEmotion(id int) *Emotion {
	emotion := &Emotion{}

	client.get(buildURI(emotionsURI, strconv.Itoa(id)), emotion)

	return emotion
}

func (emotion *Emotion) form() formEmotion {
	return formEmotion{
		Name: emotion.Name,
	}
}

func (emotion *Emotion) create(client *Client) error {
	return client.post(emotionsURI, emotion.form(), &emotion)
}

func (emotion *Emotion) update(client *Client) error {
	return client.patch(emotion.Self(), emotion.form(), &emotion)
}
