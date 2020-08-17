package dota2api

import "errors"

const uint32Max = uint64(^uint32(0))

type SteamId struct {
	id     uint64
	isId64 bool
}

func (s SteamId) SteamId64() (uint64, error) {
	if s.isId64 {
		return s.id, nil
	}
	return s.id, errors.New("expected 64bit steamId")
}

func (s SteamId) SteamId32() uint32 {
	return uint32(s.id & uint32Max)
}

func (s *SteamId) SetSteamId32(id uint32) {
	s.isId64 = false
	s.id = uint64(id)
}

func (s *SteamId) SetSteamId64(id uint64) {
	s.isId64 = true
	s.id = id
}

func NewSteamIdFrom64(id uint64) SteamId {
	return SteamId{
		id:     id,
		isId64: true,
	}
}

func NewSteamIdFrom32(id uint32) SteamId {
	return SteamId{
		id:     uint64(id),
		isId64: false,
	}
}
