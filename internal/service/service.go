package service

// Feed load and persist information about advert content.
type Feed interface {
	Load() error
}

// Statistics provide statistics data about advert content.
type Statistics interface {
}

// Tracker provide events.
type Tracker interface {
	Click() error
	Load() error
	Show() error
	Postback() error
}
