package core

import (
  "net"
  "strings"
  "errors"
  "context"
  "sync"
	"fmt"
  "time"
  "math/rand"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
  "github.com/libp2p/go-libp2p-core/host"
  "github.com/libp2p/go-libp2p-core/event"
  "github.com/libp2p/go-libp2p-core/protocol"
  "github.com/libp2p/go-libp2p-core/peerstore"
  "github.com/libp2p/go-libp2p-core/peer"
  "github.com/libp2p/go-libp2p-core/network"
  "github.com/libp2p/go-libp2p-core/connmgr"
  "github.com/libp2p/go-libp2p-peerstore/pstoremem"
  "github.com/libp2p/go-libp2p-discovery"
  "github.com/libp2p/go-libp2p-core/helpers"
  "github.com/ipfs/go-log"
	dht "github.com/libp2p/go-libp2p-kad-dht"
  mdns_discovery "github.com/libp2p/go-libp2p/p2p/discovery"

  maddr "github.com/multiformats/go-multiaddr"
)

var (
  HostHeader = "Host"
  HostLogger = log.Logger(HostHeader)

  LookUpInterval = 45 * time.Second
)

//this part was stollen from [https://github.com/libp2p/go-libp2p-examples/blob/master/chat-with-mdns/mdns.go]
type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

//interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

//Initialize the MDNS service
func initMDNS(ctx context.Context, peerhost host.Host, rendezvous string) chan peer.AddrInfo {
	// An hour might be a long long period in practical applications. But this is fine for us
	ser, err := mdns_discovery.NewMdnsService(ctx, peerhost, LookUpInterval, rendezvous)
	if err != nil {
		panic(err)
	}

	//register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	ser.RegisterNotifee(n)
	return n.PeerChan
}

func ListIpAdresses() ([]maddr.Multiaddr, error) {
  returnAddr := []maddr.Multiaddr{}
	addr, err := maddr.NewMultiaddr("/ip4/127.0.0.1/tcp/0")
  if err != nil {
  	return returnAddr, err
  }

	returnAddr = append(returnAddr, addr)

	addrs, err := net.InterfaceAddrs()
  if err != nil {
    return returnAddr, err
  }

  for _, a := range addrs {
    if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      block := strings.Split(a.String(), "/")
      if len(block) > 2 {
        return returnAddr, errors.New("Ip adress with too many slash")
      }

      if ipnet.IP.To4() != nil {
        addr, err := maddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/0", block[0]))
        if err != nil {
          return returnAddr, err
        }
        returnAddr = append(returnAddr, addr)
      } else {
        addr, err := maddr.NewMultiaddr(fmt.Sprintf("/ip6/%s/tcp/0", block[0]))
        if err != nil {
          return returnAddr, err
        }
        returnAddr = append(returnAddr, addr)
      }
    }
  }

	return returnAddr, nil
}

type StreamHandlerMatcher struct {
  Match func(string) bool
  Handler network.StreamHandler
}

func NewHost(ctx context.Context, bootstrapPeers ...maddr.Multiaddr) (ExtHost, error) {
  listenAddresses, err := ListIpAdresses()
  if err != nil {
    return nil, err
  }

  priv, _, err := crypto.GenerateKeyPair(
	   crypto.Ed25519, // Select your key type. Ed25519 are nice short
	    -1,             // Select key length when possible (i.e. RSA).
  )
  if err != nil {
    return nil, err
  }

  h, err := libp2p.New(ctx,
  	libp2p.Identity(priv),

  	libp2p.ListenAddrs(
      listenAddresses...
  	),
  )

  kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		return nil, err
	}

	err = kademliaDHT.Bootstrap(ctx)
  if err != nil {
		return nil, err
	}

  var wg sync.WaitGroup
  for _, peerAddr := range bootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
      h.Connect(ctx, *peerinfo)
		}()
	}

  routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)

  var streamHandlers sync.Map

  return &BasicExtHost {
    Ctx: ctx,
    Host: h,
    StreamHandlers: streamHandlers,
    Routing: routingDiscovery,
    PeerStores: make(map[protocol.ID] peerstore.Peerstore),
    Standard: NewStandardInterface(h.Close),
  }, nil
}

type BasicExtHost struct {
  Ctx context.Context
  Host host.Host
  StreamHandlers sync.Map
  Routing *discovery.RoutingDiscovery
  PeerStores map[protocol.ID] peerstore.Peerstore
  Standard standardFunctionsCloser
}

func (h *BasicExtHost)Close() error {
  return h.Standard.Close()
}

func (h *BasicExtHost)SetCloseHandler(handler func()) {
  h.Standard.SetCloseHandler(handler)
}

func (h *BasicExtHost)SetErrorHandler(handler func(error)) {
  h.Standard.SetErrorHandler(handler)
}

func (h *BasicExtHost)Raise(err error) {
  h.Standard.Raise(err)
}

func (h *BasicExtHost) Check() bool {
  return h.Standard.Check()
}

func (h *BasicExtHost)ID() peer.ID {
  return h.Host.ID()
}

func (h *BasicExtHost)Peerstore() peerstore.Peerstore {
  return h.Host.Peerstore()
}

func (h *BasicExtHost)Addrs() []maddr.Multiaddr {
  return h.Host.Addrs()
}

func (h *BasicExtHost)Network() network.Network {
  return h.Host.Network()
}

func (h *BasicExtHost)Mux() protocol.Switch {
  return h.Host.Mux()
}

func (h *BasicExtHost)Connect(ctx context.Context, pi peer.AddrInfo) error {
  return h.Host.Connect(ctx, pi)
}

func (h *BasicExtHost)ConnManager() connmgr.ConnManager {
  return h.Host.ConnManager()
}

func (h *BasicExtHost)EventBus() event.Bus {
  return h.Host.EventBus()
}

func (h *BasicExtHost)NewStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (network.Stream, error) {
  if p == h.ID() {
    return h.SelfStream(pids...)
  }
  return h.Host.NewStream(ctx, p, pids...)
}
func (h *BasicExtHost)SetStreamHandler(pid protocol.ID, handler network.StreamHandler) {
  match, err := helpers.MultistreamSemverMatcher(pid)
  if err != nil {
    panic(err) //shouldn't happend
  }
  h.StreamHandlers.Store(pid, &StreamHandlerMatcher {
    Match: match,
    Handler: handler,
  })
  h.Host.SetStreamHandler(pid, handler)
}

func (h *BasicExtHost)SetStreamHandlerMatch(pid protocol.ID, match func(string) bool, handler network.StreamHandler) {
  h.StreamHandlers.Store(pid, &StreamHandlerMatcher {
    Match: match,
    Handler: handler,
  })
  h.Host.SetStreamHandlerMatch(pid, match, handler)
}

func (h *BasicExtHost)RemoveStreamHandler(pid protocol.ID) {
  h.StreamHandlers.Delete(pid)
  h.Host.RemoveStreamHandler(pid)
}

func (h *BasicExtHost)PeerstoreProtocol(base protocol.ID) (peerstore.Peerstore, error) {
  pstore, ok := h.PeerStores[base]
  if !ok {
    return pstore, errors.New("no such protocol")
  }

  return pstore, nil
}

func (h *BasicExtHost)Listen(pid protocol.ID, rendezvous string) {
  if !h.Check() {
    return
  }

  h.PeerStores[pid] = pstoremem.NewPeerstore()
  h.PeerStores[pid].AddAddrs(h.ID(), h.Addrs(), peerstore.TempAddrTTL)
  discovery.Advertise(h.Ctx, h.Routing, rendezvous)
  mdnsPeerChan := initMDNS(h.Ctx, h, rendezvous)

  discoveryHandler := func(peer peer.AddrInfo) {
    if peer.ID != h.ID() {
      go func(){
        err := h.Connect(h.Ctx, peer)

        if err == nil {
          h.PeerStores[pid].AddAddrs(peer.ID, peer.Addrs, peerstore.TempAddrTTL)
          MasterLogger.Debugf("Discovered peer %s for protocol %q with rendezvous %q", peer.ID, pid, rendezvous) //--------------------------
        }
      }()
    }
  }

  go func() {
    defer recover()

    for h.Check() {
      peer := <- mdnsPeerChan
      discoveryHandler(peer)
    }
  }()

  go func() {
    defer recover()

    for h.Check() {
      peerChan, err := h.Routing.FindPeers(h.Ctx, rendezvous)
      if err != nil {
        return
      }
      for peer := range peerChan {
        discoveryHandler(peer)
      }
    }
  }()
}

func (h *BasicExtHost)NewPeer(base protocol.ID) (peer.ID, error) {
  defer func() {
    if err := recover(); err != nil {
      h.Raise(err.(error))
    }
  }()

  var nilPeerId peer.ID

  pstore, err := h.PeerstoreProtocol(base)
  if err != nil {
    return nilPeerId, err
  }

  peers := pstore.Peers()
  if len(peers) == 0 {
    return nilPeerId, errors.New("No peers supporting this protocol")
  }

  n := rand.Intn(len(peers))
  return peers[n], nil
}

func (h *BasicExtHost)SelfStream(pid ...protocol.ID) (SelfStream, error) {
  defer func() {
    if err := recover(); err != nil {
      h.Raise(err.(error))
    }
  }()

  if len(pid) == 0 {
    return nil, errors.New("no protocol given")
  }

  if len(pid) > 1 {
    return nil, errors.New("too many protocol given")
  }

  stream := NewStream(pid[0])
  reversed_stream, err := stream.Reverse()
  if err != nil {
    return nil, err
  }

  h.StreamHandlers.Range(func(key interface{}, value interface{}) bool {
    streamHandlerMatcher, ok := value.(*StreamHandlerMatcher)

    if !ok {
      return true
    }

    if !(*streamHandlerMatcher).Match(string(pid[0])) {
      return true
    }

    go (*streamHandlerMatcher).Handler(reversed_stream)
    return false
  })

  return stream, nil
}
