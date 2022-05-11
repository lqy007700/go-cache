package go_cache

import pb "go-cache/gocachepb"

type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
