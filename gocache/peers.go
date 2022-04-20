package gocache

import pb "github.com/tianye0718/go-cache/gocachepb"

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(in *pb.Request) (*pb.Response, error)
}
