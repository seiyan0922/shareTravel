package sessions // セッション管理パッケージ

import (
	"net/http"
	"net/url"
)

// /* ******************* *
// * 定数値
// * ****** */
// const (
// 	DefaultSessionName = "default-session" // デフォルトセッション名
// 	DefaultCookieName  = "default-cookie"  // デフォルトCookie名
// )

/*
 セッション構造体
*/
type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

// /*
// 新規セッション生成
// */
// func NewSession(manager *Manager, cookieName string) *Session {
// 	return &Session{
// 		cookieName: cookieName,
// 		manager:    manager,
// 		Values:     map[string]interface{}{},
// 	}
// }

// /*
// セッションの開始
// */
// func StartSession(sessionName, cookieName string, manager *Manager) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var session *Session
// 		var err error
// 		session, err = manager.Get(ctx.Request, cookieName)
// 		if err != nil {
// 			session, err = manager.New(ctx.Request, cookieName)
// 			if err != nil {
// 				println(err.Error())
// 				ctx.Abort()
// 			}
// 		}
// 		session.writer = ctx.Writer
// 		ctx.Set(sessionName, session)
// 		defer context.Clear(ctx.Request)
// 		ctx.Next()
// 	}
// }

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// /*
// デフォルトセッションの開始
// */
// func StartDefaultSession(manager *Manager) gin.HandlerFunc {
// 	return StartSession(DefaultSessionName, DefaultCookieName, manager)
// }

// /*
// セッションの取得
// */
// func GetSession(c *gin.Context, sessionName string) *Session {
// 	return c.MustGet(sessionName).(*Session)
// }

// /*
// デフォルトセッションの取得
// */
// func GetDefaultSession(c *gin.Context) *Session {
// 	return GetSession(c, DefaultSessionName)
// }

// /*
// セッションの保存
// */
// func (s *Session) Save() error {
// 	return s.manager.Save(s.request, s.writer, s)
// }

// /*
// セッション名の取得
// */
// func (s *Session) Name() string {
// 	return s.cookieName
// }

// /*
// セッション変数値の取得
// */
// func (s *Session) Get(key string) (interface{}, bool) {
// 	ret, exists := s.Values[key]
// 	return ret, exists
// }

// /*
// セッション変数値のセット
// */
// func (s *Session) Set(key string, val interface{}) {
// 	s.Values[key] = val
// }

// /*
// セッション変数の削除
// */
// func (s *Session) Delete(key string) {
// 	delete(s.Values, key)
// }

// /*
// セッションの削除
// */
// func (s *Session) Terminate() {
// 	s.manager.Destroy(s.ID)
// }
