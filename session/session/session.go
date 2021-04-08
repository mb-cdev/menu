package session

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         string
	Created_at int64
	Expire_at  int64
	Data       map[string]interface{}
}

var SessionCache sync.Map
var SessionPath string

func init() {
	var once sync.Once

	once.Do(func() {
		rootDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		SessionPath = path.Join(rootDir, "sessions")

		os.Mkdir(SessionPath, 0700)
	})
}

func New() *Session {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	s := &Session{
		ID:         uid.String(),
		Created_at: time.Now().Unix(),
		Expire_at:  time.Now().Add(time.Hour).Unix(),
		Data:       make(map[string]interface{}),
	}

	s.Save()
	SessionCache.Store(s.ID, s)

	return s
}

func Open(ID string) (*Session, error) {

	if sCached, ok := SessionCache.Load(ID); ok {
		if sCached.(*Session).IsExpired() {
			return nil, errors.New("Session expired!")
		}
		return sCached.(*Session), nil
	}

	f, err := os.Open(path.Join(SessionPath, ID))

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(f)

	s := &Session{}
	errDecode := decoder.Decode(s)

	if errDecode != nil {
		return nil, errDecode
	}

	if s.IsExpired() {
		return nil, errors.New("Session expired!")
	}

	SessionCache.Store(s.ID, s)

	return s, nil
}

func (s *Session) IsExpired() bool {
	return s.Expire_at < time.Now().Unix()
}

func (s *Session) Save() (int64, error) {
	f, err := os.OpenFile(path.Join(SessionPath, s.ID), os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		panic(err)
	}

	data, err2 := json.Marshal(s)
	if err2 != nil {
		panic(err2)
	}
	buffer := bytes.NewBuffer(data)

	return io.Copy(f, buffer)
}

func (s *Session) Drop() bool {
	res := os.Remove(path.Join(SessionPath, s.ID))
	SessionCache.Delete(s.ID)

	if res != nil {
		return false
	}

	return true
}

func (s *Session) Set(key string, value interface{}) {
	s.Data[key] = value
	s.Save()
}

func (s *Session) Get(key string) (interface{}, bool) {
	return s.Data[key]
}
