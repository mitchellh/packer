package rpc

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// MuxConn is a connection that can be used bi-directionally for RPC. Normally,
// Go RPC only allows client-to-server connections. This allows the client
// to actually act as a server as well.
//
// MuxConn works using a fairly dumb multiplexing technique of simply
// framing every piece of data sent into a prefix + data format. Streams
// are established using a subset of the TCP protocol. Only a subset is
// necessary since we assume ordering on the underlying RWC.
type MuxConn struct {
	rwc     io.ReadWriteCloser
	streams map[byte]*Stream
	mu      sync.RWMutex
	wlock   sync.Mutex
}

type muxPacketType byte

const (
	muxPacketSyn muxPacketType = iota
	muxPacketAck
	muxPacketFin
	muxPacketData
)

func NewMuxConn(rwc io.ReadWriteCloser) *MuxConn {
	m := &MuxConn{
		rwc:     rwc,
		streams: make(map[byte]*Stream),
	}

	go m.loop()

	return m
}

// Close closes the underlying io.ReadWriteCloser. This will also close
// all streams that are open.
func (m *MuxConn) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Close all the streams
	for _, w := range m.streams {
		w.Close()
	}
	m.streams = make(map[byte]*Stream)

	return m.rwc.Close()
}

// Accept accepts a multiplexed connection with the given ID. This
// will block until a request is made to connect.
func (m *MuxConn) Accept(id byte) (io.ReadWriteCloser, error) {
	stream, err := m.openStream(id)
	if err != nil {
		return nil, err
	}

	// If the stream isn't closed, then it is already open somehow
	stream.mu.Lock()
	if stream.state != streamStateSynRecv && stream.state != streamStateClosed {
		stream.mu.Unlock()
		return nil, fmt.Errorf("Stream already open in bad state: %d", stream.state)
	}

	if stream.state == streamStateSynRecv {
		// Fast track establishing since we already got the syn
		stream.setState(streamStateEstablished)
		stream.mu.Unlock()
	}

	if stream.state != streamStateEstablished {
		// Go into the listening state
		stream.setState(streamStateListen)
		stream.mu.Unlock()

		// Wait for the connection to establish
	ACCEPT_ESTABLISH_LOOP:
		for {
			time.Sleep(50 * time.Millisecond)
			stream.mu.Lock()
			switch stream.state {
			case streamStateListen:
				stream.mu.Unlock()
			case streamStateEstablished:
				stream.mu.Unlock()
				break ACCEPT_ESTABLISH_LOOP
			default:
				defer stream.mu.Unlock()
				return nil, fmt.Errorf("Stream went to bad state: %d", stream.state)
			}
		}
	}

	// Send the ack down
	if _, err := m.write(stream.id, muxPacketAck, nil); err != nil {
		return nil, err
	}

	return stream, nil
}

// Dial opens a connection to the remote end using the given stream ID.
// An Accept on the remote end will only work with if the IDs match.
func (m *MuxConn) Dial(id byte) (io.ReadWriteCloser, error) {
	stream, err := m.openStream(id)
	if err != nil {
		return nil, err
	}

	// If the stream isn't closed, then it is already open somehow
	stream.mu.Lock()
	if stream.state != streamStateClosed {
		stream.mu.Unlock()
		return nil, fmt.Errorf("Stream already open in bad state: %d", stream.state)
	}

	// Open a connection
	if _, err := m.write(stream.id, muxPacketSyn, nil); err != nil {
		return nil, err
	}
	stream.setState(streamStateSynSent)
	stream.mu.Unlock()

	for {
		time.Sleep(50 * time.Millisecond)
		stream.mu.Lock()
		switch stream.state {
		case streamStateSynSent:
			stream.mu.Unlock()
		case streamStateEstablished:
			stream.mu.Unlock()
			return stream, nil
		default:
			defer stream.mu.Unlock()
			return nil, fmt.Errorf("Stream went to bad state: %d", stream.state)
		}
	}
}

func (m *MuxConn) openStream(id byte) (*Stream, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if stream, ok := m.streams[id]; ok {
		return stream, nil
	}

	// Create the stream object and channel where data will be sent to
	dataR, dataW := io.Pipe()

	// Set the data channel so we can write to it.
	stream := &Stream{
		id:     id,
		mux:    m,
		reader: dataR,
		writer: dataW,
	}
	stream.setState(streamStateClosed)

	m.streams[id] = stream
	return m.streams[id], nil
}

func (m *MuxConn) loop() {
	defer m.Close()

	var id byte
	var packetType muxPacketType
	var length int32
	for {
		if err := binary.Read(m.rwc, binary.BigEndian, &id); err != nil {
			log.Printf("[ERR] Error reading stream ID: %s", err)
			return
		}
		if err := binary.Read(m.rwc, binary.BigEndian, &packetType); err != nil {
			log.Printf("[ERR] Error reading packet type: %s", err)
			return
		}
		if err := binary.Read(m.rwc, binary.BigEndian, &length); err != nil {
			log.Printf("[ERR] Error reading length: %s", err)
			return
		}

		// TODO(mitchellh): probably would be better to re-use a buffer...
		data := make([]byte, length)
		if length > 0 {
			if _, err := m.rwc.Read(data); err != nil {
				log.Printf("[ERR] Error reading data: %s", err)
				return
			}
		}

		stream, err := m.openStream(id)
		if err != nil {
			log.Printf("[ERR] Error opening stream %d: %s", id, err)
			return
		}

		switch packetType {
		case muxPacketAck:
			stream.mu.Lock()
			if stream.state == streamStateSynSent {
				stream.setState(streamStateEstablished)
			} else {
				log.Printf("[ERR] Ack received for stream in state: %d", stream.state)
			}
			stream.mu.Unlock()
		case muxPacketSyn:
			stream.mu.Lock()
			switch stream.state {
			case streamStateClosed:
				stream.setState(streamStateSynRecv)
			case streamStateListen:
				stream.setState(streamStateEstablished)
			default:
				log.Printf("[ERR] Syn received for stream in state: %d", stream.state)
			}
			stream.mu.Unlock()
		case muxPacketFin:
			stream.mu.Lock()
			stream.setState(streamStateClosed)
			stream.writer.Close()
			stream.mu.Unlock()

			m.mu.Lock()
			delete(m.streams, stream.id)
			m.mu.Unlock()
		case muxPacketData:
			stream.mu.Lock()
			if stream.state == streamStateEstablished {
				stream.writer.Write(data)
			} else {
				log.Printf("[ERR] Data received for stream in state: %d", stream.state)
			}
			stream.mu.Unlock()
		}
	}
}

func (m *MuxConn) write(id byte, dataType muxPacketType, p []byte) (int, error) {
	m.wlock.Lock()
	defer m.wlock.Unlock()

	if err := binary.Write(m.rwc, binary.BigEndian, id); err != nil {
		return 0, err
	}
	if err := binary.Write(m.rwc, binary.BigEndian, byte(dataType)); err != nil {
		return 0, err
	}
	if err := binary.Write(m.rwc, binary.BigEndian, int32(len(p))); err != nil {
		return 0, err
	}
	if len(p) == 0 {
		return 0, nil
	}
	return m.rwc.Write(p)
}

// Stream is a single stream of data and implements io.ReadWriteCloser
type Stream struct {
	id           byte
	mux          *MuxConn
	reader       io.Reader
	writer       io.WriteCloser
	state        streamState
	stateUpdated time.Time
	mu           sync.Mutex
}

type streamState byte

const (
	streamStateClosed streamState = iota
	streamStateListen
	streamStateSynRecv
	streamStateSynSent
	streamStateEstablished
	streamStateFinWait
)

func (s *Stream) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state != streamStateEstablished {
		return fmt.Errorf("Stream in bad state: %d", s.state)
	}

	if _, err := s.mux.write(s.id, muxPacketFin, nil); err != nil {
		return err
	}

	s.setState(streamStateClosed)
	s.writer.Close()
	return nil
}

func (s *Stream) Read(p []byte) (int, error) {
	return s.reader.Read(p)
}

func (s *Stream) Write(p []byte) (int, error) {
	return s.mux.write(s.id, muxPacketData, p)
}

func (s *Stream) setState(state streamState) {
	s.state = state
	s.stateUpdated = time.Now().UTC()
}
