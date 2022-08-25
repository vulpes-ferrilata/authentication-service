package documents

type DocumentRoot struct {
	Document `bson:",inline"`
	Version  int `bson:"version"`
}
