package go_cache

type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
