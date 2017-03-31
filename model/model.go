package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/runner-mei/radiusd/config"
)

type User struct {
	ID              int64
	Username        string
	Password        string
	ActiveUntil     *string // Account active until YYYY-MM-DD
	BlockRemain     *int64  // Remaining bandwidth
	SimultaneousUse uint32  // Max conns allowed
	DedicatedIP     *string
	Ratelimit       *string
	DnsOne          *string
	DnsTwo          *string
	Ok              bool
}

func UserID(db *sql.DB, user string) (int64, error) {
	var userID int64
	e := db.QueryRow(
		config.PlaceholderFormat(`SELECT
			id,
		FROM
			users
		WHERE
			username = ?`),
		user,
	).Scan(&userID)
	if e == config.ErrNoRows {
		return userID, nil
	}
	return userID, e
}

func Auth(db *sql.DB, user string) (User, error) {
	u := User{}
	e := db.QueryRow(
		config.PlaceholderFormat(`SELECT
			id,
			username,
			password,
			block_remaining,
			active_until,
			1,
			max_sessions,
			dedicated_address,
			CONCAT(products.ratelimit_up, products.ratelimit_unit, '/', products.ratelimit_down, products.ratelimit_unit),
			dns.one, dns.two
		FROM
			users
		JOIN
			product
		ON
			users.product_id = products.id
		LEFT JOIN
			dns
		ON
			users.dns_id = dns.id
		WHERE
			users.username = ?`),
		user,
	).Scan(&u.ID, &u.Username,
		&u.Password, &u.BlockRemain, &u.ActiveUntil, &u.Ok,
		&u.SimultaneousUse, &u.DedicatedIP, &u.Ratelimit,
		&u.DnsOne, &u.DnsTwo,
	)
	if e == config.ErrNoRows {
		return u, nil
	}
	return u, e
}

type Session struct {
	BytesIn     uint32
	BytesOut    uint32
	PacketsIn   uint32
	PacketsOut  uint32
	SessionID   string
	User        int64
	NasIP       string
	SessionTime uint32
}
type UserLimits struct {
	ID     int64
	Exists bool
}

var ErrNoRows = sql.ErrNoRows

func Begin() (*sql.Tx, error) {
	return config.DB.Begin()
}

func SessionCount(db *sql.DB, user string) (uint32, error) {
	var count uint32 = 0
	e := db.QueryRow(
		config.PlaceholderFormat(`SELECT
			COUNT(*)
		FROM
			sessions
		WHERE
		  EXISTS (
		  	SELECT * 
		  	FROM users 
		  	WHERE 
		  		  sessions.user_id == users.id 
		  		AND 
		  		  users.username = ?
		  )
			`),
		user,
	).Scan(&count)
	return count, e
}

func Limits(db *sql.DB, user string) (UserLimits, error) {
	u := UserLimits{}
	e := db.QueryRow(
		config.PlaceholderFormat(`SELECT
			users.id,
			1
		FROM
			users
		JOIN
			products
		ON
			users.product_id = products.id
		WHERE
			username = ?`),
		user,
	).Scan(&u.ID, &u.Exists)
	return u, e
}

func affectCheck(res sql.Result, expect int64, errMsg error) error {
	affect, e := res.RowsAffected()
	if e != nil {
		return e
	}
	if affect != expect {
		return errMsg
	}
	return nil
}

func SessionAdd(sessionID string, user int64, nasIp, assignedIp, clientIp string) error {
	exists := false
	e := config.DB.QueryRow(
		config.PlaceholderFormat(`SELECT
			1
		FROM
			sessions
		WHERE
			user_id = ?
		AND
			session_id = ?
		AND
			nas_address = ?`),
		user, sessionID, nasIp,
	).Scan(&exists)
	if e != nil && e != sql.ErrNoRows {
		return e
	}
	if exists {
		// Session already stored
		return nil
	}

	res, e := config.DB.Exec(
		config.PlaceholderFormat(`INSERT INTO
			sessions
		(session_id, user, time_added, nas_address, assigned_address, client_address, bytes_in, bytes_out, packets_in, packets_out, session_time)
		VALUES
		(?, ?, ?, ?, ?, ?, 0, 0, 0, 0, 0)`),
		sessionID, user, time.Now().Unix(), nasIp, assignedIp, clientIp,
	)
	if e != nil {
		return e
	}
	return affectCheck(res, 1, fmt.Errorf(
		"session.add fail for sess=%s user=%s",
		sessionID, user,
	))
}

func SessionUpdate(txn *sql.Tx, s Session) error {
	res, e := txn.Exec(
		config.PlaceholderFormat(`UPDATE
			sessions
		SET
			bytes_in = bytes_in + ?,
			bytes_out = bytes_out + ?,
			packets_in = packets_in + ?,
			packets_out = packets_out + ?,
			session_time = ?
		WHERE
			session_id = ?
		AND
			user_id = ?
		AND
			nas_address = ?`),
		s.BytesIn, s.BytesOut, s.PacketsIn, s.PacketsOut, s.SessionTime,
		s.SessionID, s.User, s.NasIP,
	)
	if e != nil {
		return e
	}
	return affectCheck(res, 1, fmt.Errorf(
		"session.update fail for sess=%s user=%s",
		s.SessionID, s.User,
	))

}

func SessionRemove(txn *sql.Tx, sessionId, user, nasIp string) error {
	res, e := txn.Exec(
		config.PlaceholderFormat(`DELETE FROM
			session
		WHERE
			session_id = ?
		AND
			user = ?
		AND
			nas_ip = ?`),
		sessionId, user, nasIp,
	)
	if e != nil {
		return e
	}
	return affectCheck(res, 1, fmt.Errorf(
		"session.remove fail for sess=%s",
		sessionId,
	))
	return nil
}

// Copy session to log
func SessionLog(txn *sql.Tx, sessionId string, user string, nasIp string) error {
	res, e := txn.Exec(
		config.PlaceholderFormat(`INSERT INTO
			session_log
			(assigned_ip, bytes_in, bytes_out, client_ip,
			nas_ip, packets_in, packets_out, session_id,
			session_time, user, time_added)
		SELECT
			assigned_ip, bytes_in, bytes_out, client_ip,
			nas_ip, packets_in, packets_out, session_id,
			session_time, user, time_added
		FROM
			session
		WHERE
			session_id = ?
		AND
			user = ?
		AND
			nas_ip = ?`),
		sessionId, user, nasIp,
	)
	if e != nil {
		return e
	}
	return affectCheck(res, 1, fmt.Errorf(
		"session.log fail for sess=%s",
		sessionId,
	))
}

func SessionAcct(db *sql.DB, user int64, hostname string, octetIn uint32, octetOut uint32, packetIn uint32, packetOut uint32, date time.Time) error {
	res, e := db.Exec(config.PlaceholderFormat(`INSERT INTO
			accounting
		(user_id, hostname, bytes_in, bytes_out, packets_in, packets_out, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`),
		user, hostname, octetIn, octetOut, packetIn, packetOut, date)
	if e != nil {
		return e
	}
	affect, e := res.RowsAffected()
	if e != nil {
		return e
	}
	if affect != 1 {
		return fmt.Errorf(
			"Affect fail for user=%s",
			user,
		)
	}
	return nil
}

func UpdateRemaining(db *sql.DB, user int64, remain uint32) error {
	if remain == 0 {
		return nil
	}

	res, e := db.Exec(config.PlaceholderFormat(`UPDATE
			users
		SET
			block_remaining = CASE WHEN block_remaining - ? < 0 THEN 0 ELSE block_remaining - ? END
		WHERE
			id = ?`), remain, remain, user)
	if e != nil {
		return e
	}
	affect, e := res.RowsAffected()
	if e != nil {
		return e
	}
	if affect != 1 {
		// Nothing changed, check if this behaviour is correct
		remain, e := checkRemain(db, user)
		if e != nil {
			return e
		}
		if !remain {
			return fmt.Errorf(
				"Affect fail for user=%s",
				user,
			)
		}
	}
	return nil
}

func checkRemain(db *sql.DB, user int64) (bool, error) {
	var remain *int64

	e := db.QueryRow(
		config.PlaceholderFormat(`SELECT
			block_remaining
		FROM
			users
		WHERE
			username = ?`),
		user,
	).Scan(&remain)
	if remain == nil || *remain == 0 {
		return true, e
	}
	return false, e
}
