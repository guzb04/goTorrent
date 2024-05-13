package torrentfile

import (
	"io"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

const Port uint16 = 6842

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}


type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}


func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)

	if err != nil {
		return nil, err
	}
	return &bto, nil
}

