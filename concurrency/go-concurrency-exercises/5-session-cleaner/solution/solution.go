package solution

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

const timeLimit = 5 * time.Second

type SessionManager struct {
	sessions map[string]*Session
	mu       sync.Mutex
}

type Session struct {
	Data       map[string]interface{}
	updateTime time.Time
	mu         sync.Mutex
}

func sessionCleaner(ctx context.Context, manager *SessionManager) {
	ticker := time.NewTicker(timeLimit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			manager.mu.Lock()
			now := time.Now()

			for key, session := range manager.sessions {
				session.mu.Lock()
				if now.Sub(session.updateTime) >= timeLimit {
					delete(manager.sessions, key)
				}
				session.mu.Unlock()
			}
			manager.mu.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func NewSessionManager() *SessionManager {
	ctx := context.Context(context.Background())

	m := &SessionManager{
		sessions: make(map[string]*Session),
	}

	go sessionCleaner(ctx, m)
	return m
}

func (m *SessionManager) CreateSession() (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

	m.sessions[sessionID] = &Session{
		Data:       make(map[string]interface{}),
		updateTime: time.Now(),
	}

	return sessionID, nil
}

var ErrSessionNotFound = errors.New("SessionID does not exists")

func (m *SessionManager) GetSessionData(sessionID string) (map[string]interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, ok := m.sessions[sessionID]
	if !ok {
		return nil, ErrSessionNotFound
	}
	return session.Data, nil
}

func (m *SessionManager) UpdateSessionData(sessionID string, data map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.sessions[sessionID]
	if !ok {
		return ErrSessionNotFound
	}

	m.sessions[sessionID] = &Session{
		Data:       data,
		updateTime: time.Now(),
	}

	return nil
}

func main() {
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created new session with ID", sID)

	data := make(map[string]interface{})
	data["website"] = "longhoang.de"

	err = m.UpdateSessionData(sID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update session data, set website to longhoang.de")

	updatedData, err := m.GetSessionData(sID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get session data:", updatedData)
}
