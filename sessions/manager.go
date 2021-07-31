package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sync"
)

/*
セッションマネージャ構造体
*/
type Manager struct {
	// database map[string]interface{}
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

// var mg Manager

// /*
// 新規マネージャ生成
// */
// func NewManager() *Manager {
// 	return &mg
// }

// /*
// セッションIDの発行
// */
// func (m *Manager) NewSessionID() string {
// 	b := make([]byte, 64)
// 	if _, err := io.ReadFull(rand.Reader, b); err != nil {
// 		return ""
// 	}
// 	return base64.URLEncoding.EncodeToString(b)
// }

// /*
// 新規セッションの生成
// */
// func (m *Manager) New(r *http.Request, cookieName string) (*Session, error) {
// 	cookie, err := r.Cookie(cookieName)
// 	if err == nil && m.Exists(cookie.Value) {
// 		return nil, errors.New("sessionIDはすでに発行されています")
// 	}

// 	session := NewSession(m, cookieName)
// 	session.ID = m.NewSessionID()
// 	session.request = r

// 	return session, nil
// }

// /*
// セッション情報の保存
// */
// func (m *Manager) Save(r *http.Request, w http.ResponseWriter, session *Session) error {
// 	m.database[session.ID] = session

// 	c := &http.Cookie{
// 		Name:  session.Name(),
// 		Value: session.ID,
// 		Path:  "/",
// 	}

// 	http.SetCookie(session.writer, c)
// 	return nil
// }

// /*
// 既存セッションの存在チェック
// */
// func (m *Manager) Exists(sessionID string) bool {
// 	_, r := m.database[sessionID]
// 	return r
// }

// /*
// 既存セッションの取得
// */
// func (m *Manager) Get(r *http.Request, cookieName string) (*Session, error) {
// 	cookie, err := r.Cookie(cookieName)
// 	if err != nil {
// 		fmt.Println("くっきー取得できてないよ〜")
// 		// リクエストからcookie情報を取得できない場合
// 		return nil, err
// 	}

// 	sessionID := cookie.Value
// 	// cookie情報からセッション情報を取得
// 	buffer, exists := m.database[sessionID]
// 	if !exists {
// 		return nil, errors.New("無効なセッションIDです")
// 	}

// 	session := buffer.(*Session)
// 	session.request = r
// 	return session, nil
// }

// /*
// セッションの破棄
// */
// func (m *Manager) Destroy(sessionID string) {
// 	delete(m.database, sessionID)
// }
